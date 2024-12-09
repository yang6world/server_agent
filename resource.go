package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os/exec"
	pb "server_agent/module/proto"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	gopsutilNet "github.com/shirou/gopsutil/net"
)

// CheckResources 获取系统资源信息
func CheckResources() (map[string]interface{}, error) {
	hostInfo, err := host.Info()
	if err != nil {
		return nil, fmt.Errorf("获取主机信息失败: %v", err)
	}

	cpuUsage, err := getCpuUsage()
	if err != nil {
		return nil, fmt.Errorf("获取 CPU 使用率失败: %v", err)
	}

	memStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, fmt.Errorf("获取内存使用率失败: %v", err)
	}

	swapStat, err := mem.SwapMemory()
	if err != nil {
		return nil, fmt.Errorf("获取 Swap 使用率失败: %v", err)
	}

	diskStat, err := disk.Usage("/")
	if err != nil {
		return nil, fmt.Errorf("获取磁盘使用率失败: %v", err)
	}

	loadStat, err := load.Avg()
	if err != nil {
		return nil, fmt.Errorf("获取负载平均值失败: %v", err)
	}
	// 获取网络上传和下载速度
	netIO, err := gopsutilNet.IOCounters(false)
	if err != nil {
		return nil, fmt.Errorf("获取网络 IO 计数器失败: %v", err)
	}
	var netUploadSpeed, netDownloadSpeed float64
	if len(netIO) > 0 {
		netUploadSpeed = float64(netIO[0].BytesSent) / 1024 / 1024   // 转换为 MB
		netDownloadSpeed = float64(netIO[0].BytesRecv) / 1024 / 1024 // 转换为 MB
	}

	return map[string]interface{}{
		"hostname":           hostInfo.Hostname,
		"os":                 hostInfo.OS,
		"kernel_version":     hostInfo.KernelVersion,
		"cpu_usage":          cpuUsage,
		"memory_usage":       memStat.UsedPercent,
		"swap_usage":         swapStat.UsedPercent,
		"disk_usage":         fmt.Sprintf("%.2f%% of %.2f GB", diskStat.UsedPercent, float64(diskStat.Total)/1e9),
		"load_average":       loadStat.Load1,
		"net_upload_speed":   netUploadSpeed,
		"net_download_speed": netDownloadSpeed,
		"webshell_supported": checkWebShellSupport(),
	}, nil
}

// 获取 CPU 使用率
func getCpuUsage() (float64, error) {
	percentages, err := cpu.Percent(1*time.Second, false)
	if err != nil {
		return 0, err
	}
	if len(percentages) > 0 {
		return percentages[0], nil
	}
	return 0, fmt.Errorf("没有可用的 CPU 数据")
}

// 检查 WebShell 支持
func checkWebShellSupport() bool {
	cmd := exec.Command("which", "bash")
	err := cmd.Run()
	return err == nil
}

// 获取网络接口 IP 信息
func GetIPAddresses() ([]string, error) {
	var ips []string
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip != nil && ip.To4() != nil {
				ips = append(ips, ip.String())
			}
		}
	}
	return ips, nil
}

func calculateCPUPercent(stats types.StatsJSON) float64 {
	cpuDelta := float64(stats.CPUStats.CPUUsage.TotalUsage) - float64(stats.PreCPUStats.CPUUsage.TotalUsage)
	systemDelta := float64(stats.CPUStats.SystemUsage) - float64(stats.PreCPUStats.SystemUsage)
	if systemDelta > 0.0 && cpuDelta > 0.0 {
		return (cpuDelta / systemDelta) * float64(len(stats.CPUStats.CPUUsage.PercpuUsage)) * 100.0
	}
	return 0.0
}

func GetDockerInfo() ([]*pb.ContainerInfo, bool) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, false
	}
	defer cli.Close()

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		return nil, false
	}

	var containerInfos []*pb.ContainerInfo
	for _, container := range containers {
		stats, err := cli.ContainerStatsOneShot(context.Background(), container.ID)
		if err != nil {
			return nil, false
		}

		var statsJSON types.StatsJSON
		if err := json.NewDecoder(stats.Body).Decode(&statsJSON); err != nil {
			return nil, false
		}

		var memoryUsage float64
		if statsJSON.MemoryStats.Usage != 0 {
			memoryUsage = float64(statsJSON.MemoryStats.Usage) / (1024 * 1024) // 转换为 MB
		}

		cpuUsage := calculateCPUPercent(statsJSON)

		containerInfos = append(containerInfos, &pb.ContainerInfo{
			Id:          container.ID[:12],
			Name:        container.Names[0],
			Image:       container.Image,
			Status:      container.Status,
			MemoryUsage: fmt.Sprintf("%.2f MB", memoryUsage),
			CpuUsage:    fmt.Sprintf("%.2f%%", cpuUsage),
		})
	}
	return containerInfos, true
}
