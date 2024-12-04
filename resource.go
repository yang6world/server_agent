package main

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

// CheckResources retrieves system resource information.
func CheckResources() (map[string]interface{}, error) {
	// Host Info
	hostInfo, err := host.Info()
	if err != nil {
		return nil, fmt.Errorf("failed to get host info: %v", err)
	}
	hostname := hostInfo.Hostname
	os := hostInfo.OS
	kernelVersion := hostInfo.KernelVersion

	// CPU Usage
	cpuUsage, err := getCpuUsage()
	if err != nil {
		return nil, fmt.Errorf("failed to get CPU usage: %v", err)
	}

	// Memory Usage
	memStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, fmt.Errorf("failed to get memory usage: %v", err)
	}
	memoryUsage := memStat.UsedPercent

	// Swap Usage
	swapStat, err := mem.SwapMemory()
	if err != nil {
		return nil, fmt.Errorf("failed to get swap usage: %v", err)
	}
	swapUsage := swapStat.UsedPercent

	// Disk Usage
	diskStat, err := disk.Usage("/")
	if err != nil {
		return nil, fmt.Errorf("failed to get disk usage: %v", err)
	}
	diskUsage := fmt.Sprintf("%.2f%% used of %.2fGB", diskStat.UsedPercent, float64(diskStat.Total)/1e9)

	// Load Average
	loadStat, err := load.Avg()
	if err != nil {
		return nil, fmt.Errorf("failed to get load average: %v", err)
	}
	loadAverage := loadStat.Load1
	// Network Up
	netWorkStat, err := load.Network()
	if err != nil {
		return nil, fm + t.Errorf("failed to get network status: %v", err)
	}
	// Network Down

	// WebShell Support
	webShellSupported := checkWebShellSupport()

	return map[string]interface{}{
		"hostname":           hostname,
		"os":                 os,
		"kernel_version":     kernelVersion,
		"cpu_usage":          cpuUsage,
		"memory_usage":       memoryUsage,
		"swap_usage":         swapUsage,
		"disk_usage":         diskUsage,
		"load_average":       loadAverage,
		"webshell_supported": webShellSupported,
	}, nil
}

// getCpuUsage returns the CPU usage as a percentage.
func getCpuUsage() (float64, error) {
	percentages, err := cpu.Percent(1*time.Second, false)
	if err != nil {
		return 0, err
	}
	if len(percentages) > 0 {
		return percentages[0], nil
	}
	return 0, fmt.Errorf("no CPU usage data available")
}

// checkWebShellSupport checks if a web shell is supported.
func checkWebShellSupport() bool {
	// Example: Check if Bash exists (Linux example)
	cmd := exec.Command("which", "bash")
	err := cmd.Run()
	return err == nil
}
