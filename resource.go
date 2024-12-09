package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net"
	"os/exec"
	"time"

	pb "server_agent/module/proto"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	gopsutilNet "github.com/shirou/gopsutil/net"
)

func roundToTwoDecimalPlaces(value float64) float64 {
	return math.Round(value*100) / 100
}

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
		netUploadSpeed = float64(netIO[0].BytesSent) / 1024 / 1024 / 1024   // 转换为 GB
		netDownloadSpeed = float64(netIO[0].BytesRecv) / 1024 / 1024 / 1024 // 转换为 GB
	}

	// 获取实时网络速度
	netSpeed, err := getRealTimeNetSpeed()
	if err != nil {
		return nil, fmt.Errorf("获取实时网络速度失败: %v", err)
	}

	// 获取 CPU 核数
	cpuCount, err := cpu.Counts(true)
	if err != nil {
		return nil, fmt.Errorf("获取 CPU 核数失败: %v", err)
	}

	// 获取机器工作时间（以天为单位）
	uptime := hostInfo.Uptime / (60 * 60 * 24)

	return map[string]interface{}{
		"hostname":            hostInfo.Hostname,
		"os":                  hostInfo.OS,
		"kernel_version":      hostInfo.KernelVersion,
		"cpu_usage":           roundToTwoDecimalPlaces(cpuUsage),
		"memory_usage":        roundToTwoDecimalPlaces(memStat.UsedPercent),
		"swap_usage":          roundToTwoDecimalPlaces(swapStat.UsedPercent),
		"disk_usage":          roundToTwoDecimalPlaces(diskStat.UsedPercent),
		"disk_total":          roundToTwoDecimalPlaces(float64(diskStat.Total) / 1e9), // 转换为 GB
		"load_average":        roundToTwoDecimalPlaces(loadStat.Load1),
		"net_upload":          roundToTwoDecimalPlaces(netUploadSpeed),
		"net_download":        roundToTwoDecimalPlaces(netDownloadSpeed),
		"real_time_net_speed": netSpeed,
		"cpu_count":           cpuCount,
		"memory_total":        roundToTwoDecimalPlaces(float64((memStat.Total / (1024 * 1024 * 1024)))), // 转换为 GB
		"uptime_days":         uptime,
		"webshell_supported":  checkWebShellSupport(),
	}, nil
}

// 获取实时网络速度
func getRealTimeNetSpeed() (map[string]float64, error) {
	netIO1, err := gopsutilNet.IOCounters(false)
	if err != nil {
		return nil, err
	}
	time.Sleep(1 * time.Second)
	netIO2, err := gopsutilNet.IOCounters(false)
	if err != nil {
		return nil, err
	}

	if len(netIO1) == 0 || len(netIO2) == 0 {
		return nil, fmt.Errorf("无法获取网络 IO 数据")
	}

	uploadSpeed := float64(netIO2[0].BytesSent-netIO1[0].BytesSent) / 1024 / 1024   // 转换为 MB/s
	downloadSpeed := float64(netIO2[0].BytesRecv-netIO1[0].BytesRecv) / 1024 / 1024 // 转换为 MB/s

	return map[string]float64{
		"upload_speed":   roundToTwoDecimalPlaces(uploadSpeed),
		"download_speed": roundToTwoDecimalPlaces(downloadSpeed),
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
