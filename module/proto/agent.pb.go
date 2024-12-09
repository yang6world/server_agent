// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0--rc2
// source: proto/agent.proto

package agent

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ResourceRequest) Reset() {
	*x = ResourceRequest{}
	mi := &file_proto_agent_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceRequest) ProtoMessage() {}

func (x *ResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_agent_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceRequest.ProtoReflect.Descriptor instead.
func (*ResourceRequest) Descriptor() ([]byte, []int) {
	return file_proto_agent_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname          string           `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Os                string           `protobuf:"bytes,2,opt,name=os,proto3" json:"os,omitempty"`
	KernelVersion     string           `protobuf:"bytes,3,opt,name=kernel_version,json=kernelVersion,proto3" json:"kernel_version,omitempty"`
	CpuUsage          float64          `protobuf:"fixed64,4,opt,name=cpu_usage,json=cpuUsage,proto3" json:"cpu_usage,omitempty"`
	MemoryUsage       float64          `protobuf:"fixed64,5,opt,name=memory_usage,json=memoryUsage,proto3" json:"memory_usage,omitempty"`
	SwapUsage         float64          `protobuf:"fixed64,6,opt,name=swap_usage,json=swapUsage,proto3" json:"swap_usage,omitempty"`
	DiskUsage         string           `protobuf:"bytes,7,opt,name=disk_usage,json=diskUsage,proto3" json:"disk_usage,omitempty"`
	LoadAverage       float64          `protobuf:"fixed64,8,opt,name=load_average,json=loadAverage,proto3" json:"load_average,omitempty"`
	WebshellSupported bool             `protobuf:"varint,9,opt,name=webshell_supported,json=webshellSupported,proto3" json:"webshell_supported,omitempty"`
	StartTime         string           `protobuf:"bytes,10,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	IpAddresses       []string         `protobuf:"bytes,11,rep,name=ip_addresses,json=ipAddresses,proto3" json:"ip_addresses,omitempty"`
	NetUploadSpeed    float64          `protobuf:"fixed64,12,opt,name=net_upload_speed,json=netUploadSpeed,proto3" json:"net_upload_speed,omitempty"`
	NetDownloadSpeed  float64          `protobuf:"fixed64,13,opt,name=net_download_speed,json=netDownloadSpeed,proto3" json:"net_download_speed,omitempty"`
	DockerAvailable   bool             `protobuf:"varint,14,opt,name=docker_available,json=dockerAvailable,proto3" json:"docker_available,omitempty"`
	Containers        []*ContainerInfo `protobuf:"bytes,15,rep,name=containers,proto3" json:"containers,omitempty"`
}

func (x *ResourceResponse) Reset() {
	*x = ResourceResponse{}
	mi := &file_proto_agent_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceResponse) ProtoMessage() {}

func (x *ResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_agent_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceResponse.ProtoReflect.Descriptor instead.
func (*ResourceResponse) Descriptor() ([]byte, []int) {
	return file_proto_agent_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceResponse) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *ResourceResponse) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *ResourceResponse) GetKernelVersion() string {
	if x != nil {
		return x.KernelVersion
	}
	return ""
}

func (x *ResourceResponse) GetCpuUsage() float64 {
	if x != nil {
		return x.CpuUsage
	}
	return 0
}

func (x *ResourceResponse) GetMemoryUsage() float64 {
	if x != nil {
		return x.MemoryUsage
	}
	return 0
}

func (x *ResourceResponse) GetSwapUsage() float64 {
	if x != nil {
		return x.SwapUsage
	}
	return 0
}

func (x *ResourceResponse) GetDiskUsage() string {
	if x != nil {
		return x.DiskUsage
	}
	return ""
}

func (x *ResourceResponse) GetLoadAverage() float64 {
	if x != nil {
		return x.LoadAverage
	}
	return 0
}

func (x *ResourceResponse) GetWebshellSupported() bool {
	if x != nil {
		return x.WebshellSupported
	}
	return false
}

func (x *ResourceResponse) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *ResourceResponse) GetIpAddresses() []string {
	if x != nil {
		return x.IpAddresses
	}
	return nil
}

func (x *ResourceResponse) GetNetUploadSpeed() float64 {
	if x != nil {
		return x.NetUploadSpeed
	}
	return 0
}

func (x *ResourceResponse) GetNetDownloadSpeed() float64 {
	if x != nil {
		return x.NetDownloadSpeed
	}
	return 0
}

func (x *ResourceResponse) GetDockerAvailable() bool {
	if x != nil {
		return x.DockerAvailable
	}
	return false
}

func (x *ResourceResponse) GetContainers() []*ContainerInfo {
	if x != nil {
		return x.Containers
	}
	return nil
}

type ShellRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token   string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Command string `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"` // 客户端发送的命令
}

func (x *ShellRequest) Reset() {
	*x = ShellRequest{}
	mi := &file_proto_agent_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShellRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShellRequest) ProtoMessage() {}

func (x *ShellRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_agent_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShellRequest.ProtoReflect.Descriptor instead.
func (*ShellRequest) Descriptor() ([]byte, []int) {
	return file_proto_agent_proto_rawDescGZIP(), []int{2}
}

func (x *ShellRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ShellRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

type ShellResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output string `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"` // 命令执行的输出
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`   // 命令执行的错误信息
}

func (x *ShellResponse) Reset() {
	*x = ShellResponse{}
	mi := &file_proto_agent_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShellResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShellResponse) ProtoMessage() {}

func (x *ShellResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_agent_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShellResponse.ProtoReflect.Descriptor instead.
func (*ShellResponse) Descriptor() ([]byte, []int) {
	return file_proto_agent_proto_rawDescGZIP(), []int{3}
}

func (x *ShellResponse) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

func (x *ShellResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ContainerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Image       string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	Status      string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	MemoryUsage string `protobuf:"bytes,5,opt,name=memory_usage,json=memoryUsage,proto3" json:"memory_usage,omitempty"`
	CpuUsage    string `protobuf:"bytes,6,opt,name=cpu_usage,json=cpuUsage,proto3" json:"cpu_usage,omitempty"`
}

func (x *ContainerInfo) Reset() {
	*x = ContainerInfo{}
	mi := &file_proto_agent_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContainerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerInfo) ProtoMessage() {}

func (x *ContainerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_agent_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerInfo.ProtoReflect.Descriptor instead.
func (*ContainerInfo) Descriptor() ([]byte, []int) {
	return file_proto_agent_proto_rawDescGZIP(), []int{4}
}

func (x *ContainerInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ContainerInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ContainerInfo) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *ContainerInfo) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ContainerInfo) GetMemoryUsage() string {
	if x != nil {
		return x.MemoryUsage
	}
	return ""
}

func (x *ContainerInfo) GetCpuUsage() string {
	if x != nil {
		return x.CpuUsage
	}
	return ""
}

var File_proto_agent_proto protoreflect.FileDescriptor

var file_proto_agent_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x22, 0x27, 0x0a, 0x0f, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0xb0, 0x04, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x6f, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6b, 0x65,
	0x72, 0x6e, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x70, 0x75, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08,
	0x63, 0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x77, 0x61, 0x70, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x09, 0x73, 0x77, 0x61, 0x70, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x69,
	0x73, 0x6b, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x64, 0x69, 0x73, 0x6b, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0b, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x12,
	0x77, 0x65, 0x62, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x77, 0x65, 0x62, 0x73, 0x68, 0x65,
	0x6c, 0x6c, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x70,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x28, 0x0a,
	0x10, 0x6e, 0x65, 0x74, 0x5f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x70, 0x65, 0x65,
	0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x6e, 0x65, 0x74, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x6e, 0x65, 0x74, 0x5f, 0x64,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x10, 0x6e, 0x65, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x5f,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0f, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x34, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x0f,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x22, 0x3e, 0x0a, 0x0c, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x3d, 0x0a, 0x0d, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xa1, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x70, 0x75, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65, 0x32, 0x8b, 0x01, 0x0a, 0x0f, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x41, 0x0a,
	0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12,
	0x16, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x35, 0x0a, 0x08, 0x52, 0x75, 0x6e, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x12, 0x13, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x14, 0x5a, 0x12, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_agent_proto_rawDescOnce sync.Once
	file_proto_agent_proto_rawDescData = file_proto_agent_proto_rawDesc
)

func file_proto_agent_proto_rawDescGZIP() []byte {
	file_proto_agent_proto_rawDescOnce.Do(func() {
		file_proto_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_agent_proto_rawDescData)
	})
	return file_proto_agent_proto_rawDescData
}

var file_proto_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_agent_proto_goTypes = []any{
	(*ResourceRequest)(nil),  // 0: agent.ResourceRequest
	(*ResourceResponse)(nil), // 1: agent.ResourceResponse
	(*ShellRequest)(nil),     // 2: agent.ShellRequest
	(*ShellResponse)(nil),    // 3: agent.ShellResponse
	(*ContainerInfo)(nil),    // 4: agent.ContainerInfo
}
var file_proto_agent_proto_depIdxs = []int32{
	4, // 0: agent.ResourceResponse.containers:type_name -> agent.ContainerInfo
	0, // 1: agent.ResourceChecker.CheckResources:input_type -> agent.ResourceRequest
	2, // 2: agent.ResourceChecker.RunShell:input_type -> agent.ShellRequest
	1, // 3: agent.ResourceChecker.CheckResources:output_type -> agent.ResourceResponse
	3, // 4: agent.ResourceChecker.RunShell:output_type -> agent.ShellResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_agent_proto_init() }
func file_proto_agent_proto_init() {
	if File_proto_agent_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_agent_proto_goTypes,
		DependencyIndexes: file_proto_agent_proto_depIdxs,
		MessageInfos:      file_proto_agent_proto_msgTypes,
	}.Build()
	File_proto_agent_proto = out.File
	file_proto_agent_proto_rawDesc = nil
	file_proto_agent_proto_goTypes = nil
	file_proto_agent_proto_depIdxs = nil
}
