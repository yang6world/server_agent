package main

import (
	"context"
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

	return map[string]interface{}{
		"hostname":           hostInfo.Hostname,
		"os":                 hostInfo.OS,
		"kernel_version":     hostInfo.KernelVersion,
		"cpu_usage":          cpuUsage,
		"memory_usage":       memStat.UsedPercent,
		"swap_usage":         swapStat.UsedPercent,
		"disk_usage":         fmt.Sprintf("%.2f%% of %.2f GB", diskStat.UsedPercent, float64(diskStat.Total)/1e9),
		"load_average":       loadStat.Load1,
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

// 执行 Shell 命令
func ExecuteShellCommand(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

// 获取 Docker 信息
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
		containerInfos = append(containerInfos, &pb.ContainerInfo{
			Id:     container.ID[:12],
			Name:   container.Names[0],
			Image:  container.Image,
			Status: container.Status,
		})
	}
	return containerInfos, true
}
