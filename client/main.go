package main

import (
	"bufio"
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	pb "server_agent/module/proto" // 替换为您的 proto 文件生成的包路径

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func main() {
	// 加载 TLS 凭据
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("server.crt")
	if err != nil {
		log.Fatalf("无法读取 CA 证书: %v", err)
	}
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		RootCAs: certPool,
	})

	// 创建 gRPC 客户端连接
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("连接服务器失败: %v", err)
	}
	defer conn.Close()

	client := pb.NewResourceCheckerClient(conn)

	// 设置上下文和鉴权 Token
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "my-secure-token")
	// 输入需要执行的 Shell 命令

	// 打印命令以进行验证
	var command string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入需要执行的 Shell 命令: ")
	command, _ = reader.ReadString('\n') // 修改为赋值操作
	command = strings.TrimSpace(command)

	// 打印命令以进行验证
	fmt.Printf("发送的命令: %s\n", command)

	// 调用 RunShell 方法执行 Shell 命令
	shellReq := &pb.ShellRequest{
		Command: command,
	}

	resp, err := client.RunShell(ctx, shellReq)
	if err != nil {
		log.Fatalf("执行 Shell 命令失败: %v", err)
	}

	// 输出服务器返回的结果
	if resp.Error != "" {
		fmt.Printf("执行错误: %s\n", resp.Error)
	} else {
		fmt.Printf("执行结果:\n%s\n", resp.Output)
	}
}
