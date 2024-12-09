package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os/exec"
	"time"

	pb "server_agent/module/proto" // 替换为实际 proto 包路径

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// ResourceCheckerServer 定义服务
type ResourceCheckerServer struct {
	pb.UnimplementedResourceCheckerServer
}

var serverStartTime = time.Now()

// CheckResources 实现资源检查逻辑
func (s *ResourceCheckerServer) CheckResources(ctx context.Context, req *pb.ResourceRequest) (*pb.ResourceResponse, error) {
	if err := AuthInterceptor(ctx); err != nil {
		return nil, err
	}

	// 获取系统信息
	data, err := CheckResources()
	if err != nil {
		return nil, err
	}

	// 获取网络接口信息
	ipAddresses, err := GetIPAddresses()
	if err != nil {
		ipAddresses = []string{"无法获取 IP"}
	}

	// 获取 Docker 信息
	containers, dockerAvailable := GetDockerInfo()

	return &pb.ResourceResponse{
		Hostname:          data["hostname"].(string),
		Os:                data["os"].(string),
		KernelVersion:     data["kernel_version"].(string),
		CpuUsage:          data["cpu_usage"].(float64),
		MemoryUsage:       data["memory_usage"].(float64),
		SwapUsage:         data["swap_usage"].(float64),
		DiskUsage:         data["disk_usage"].(float64),
		DiskTotal:         data["disk_total"].(float64),
		LoadAverage:       data["load_average"].(float64),
		NetUpload:         data["net_upload"].(float64),
		NetDownload:       data["net_download"].(float64),
		RealTimeNetSpeed:  data["real_time_net_speed"].(map[string]float64),
		CpuCount:          int32(data["cpu_count"].(int)),
		MemoryTotal:       data["memory_total"].(float64),
		UptimeDays:        data["uptime_days"].(float64),
		WebshellSupported: data["webshell_supported"].(bool),
		IpAddresses:       ipAddresses,
		DockerAvailable:   dockerAvailable,
		Containers:        containers,
	}, nil
}

// RunShell 执行客户端发来的 Shell 命令
func (s *ResourceCheckerServer) RunShell(ctx context.Context, req *pb.ShellRequest) (*pb.ShellResponse, error) {
	if err := AuthInterceptor(ctx); err != nil {
		return nil, err
	}
	cmd := exec.Command("bash", "-c", req.Command)
	cmd.Args = []string{"bash", "-c", req.Command}
	fmt.Printf("执行命令: %s\n", req.Command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return &pb.ShellResponse{Output: string(output), Error: err.Error()}, nil
	}
	return &pb.ShellResponse{Output: string(output)}, nil
}

func main() {
	// 加载 TLS 配置
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatalf("加载 TLS 密钥失败: %v", err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("server.crt")
	if err != nil {
		log.Fatalf("读取 CA 证书失败: %v", err)
	}
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    certPool,
	})

	// 启动 gRPC 服务
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterResourceCheckerServer(grpcServer, &ResourceCheckerServer{})
	fmt.Println("服务器运行于 :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("启动服务失败: %v", err)
	}
}
