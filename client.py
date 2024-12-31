import grpc
import ssl
import os
from concurrent import futures
from grpc import ssl_channel_credentials
import module.proto.agent_pb2 as pb2
import module.proto.agent_pb2_grpc as pb2_grpc

def get_stub():
    # 加载证书
    with open('server.crt', 'rb') as f:
        trusted_certs = f.read()
    credentials = ssl_channel_credentials(root_certificates=trusted_certs)

    # 创建 gRPC 通道
    channel = grpc.secure_channel('localhost:50051', credentials)
    stub = pb2_grpc.ResourceCheckerStub(channel)
    return stub

def check_resources(stub):
    # 创建请求
    request = pb2.ResourceRequest()
    # 调用 CheckResources 方法
    response = stub.CheckResources(request)
    print("CheckResources Response:", response)

def run_shell(stub, command):
    # 创建请求
    request = pb2.ShellRequest(command=command)
    # 调用 RunShell 方法
    response = stub.RunShell(request)
    print("RunShell Response:", response)

def main():
    stub = get_stub()
    check_resources(stub)
    run_shell(stub, "echo Hello, World!")

if __name__ == '__main__':
    main()