//
// Copyright 2023 The Chainloop Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: controlplane/v1/org_metrics.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type MetricsTimeWindow int32

const (
	MetricsTimeWindow_METRICS_TIME_WINDOW_UNSPECIFIED  MetricsTimeWindow = 0
	MetricsTimeWindow_METRICS_TIME_WINDOW_LAST_30_DAYS MetricsTimeWindow = 1
	MetricsTimeWindow_METRICS_TIME_WINDOW_LAST_7_DAYS  MetricsTimeWindow = 2
	MetricsTimeWindow_METRICS_TIME_WINDOW_LAST_DAY     MetricsTimeWindow = 3
)

// Enum value maps for MetricsTimeWindow.
var (
	MetricsTimeWindow_name = map[int32]string{
		0: "METRICS_TIME_WINDOW_UNSPECIFIED",
		1: "METRICS_TIME_WINDOW_LAST_30_DAYS",
		2: "METRICS_TIME_WINDOW_LAST_7_DAYS",
		3: "METRICS_TIME_WINDOW_LAST_DAY",
	}
	MetricsTimeWindow_value = map[string]int32{
		"METRICS_TIME_WINDOW_UNSPECIFIED":  0,
		"METRICS_TIME_WINDOW_LAST_30_DAYS": 1,
		"METRICS_TIME_WINDOW_LAST_7_DAYS":  2,
		"METRICS_TIME_WINDOW_LAST_DAY":     3,
	}
)

func (x MetricsTimeWindow) Enum() *MetricsTimeWindow {
	p := new(MetricsTimeWindow)
	*p = x
	return p
}

func (x MetricsTimeWindow) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricsTimeWindow) Descriptor() protoreflect.EnumDescriptor {
	return file_controlplane_v1_org_metrics_proto_enumTypes[0].Descriptor()
}

func (MetricsTimeWindow) Type() protoreflect.EnumType {
	return &file_controlplane_v1_org_metrics_proto_enumTypes[0]
}

func (x MetricsTimeWindow) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricsTimeWindow.Descriptor instead.
func (MetricsTimeWindow) EnumDescriptor() ([]byte, []int) {
	return file_controlplane_v1_org_metrics_proto_rawDescGZIP(), []int{0}
}

type OrgMetricsServiceTotalsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeWindow MetricsTimeWindow `protobuf:"varint,1,opt,name=time_window,json=timeWindow,proto3,enum=controlplane.v1.MetricsTimeWindow" json:"time_window,omitempty"`
}

func (x *OrgMetricsServiceTotalsRequest) Reset() {
	*x = OrgMetricsServiceTotalsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_org_metrics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrgMetricsServiceTotalsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrgMetricsServiceTotalsRequest) ProtoMessage() {}

func (x *OrgMetricsServiceTotalsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_org_metrics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrgMetricsServiceTotalsRequest.ProtoReflect.Descriptor instead.
func (*OrgMetricsServiceTotalsRequest) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_org_metrics_proto_rawDescGZIP(), []int{0}
}

func (x *OrgMetricsServiceTotalsRequest) GetTimeWindow() MetricsTimeWindow {
	if x != nil {
		return x.TimeWindow
	}
	return MetricsTimeWindow_METRICS_TIME_WINDOW_UNSPECIFIED
}

type TopWorkflowsByRunsCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// top x number of runs to return
	NumWorkflows int32             `protobuf:"varint,1,opt,name=num_workflows,json=numWorkflows,proto3" json:"num_workflows,omitempty"`
	TimeWindow   MetricsTimeWindow `protobuf:"varint,2,opt,name=time_window,json=timeWindow,proto3,enum=controlplane.v1.MetricsTimeWindow" json:"time_window,omitempty"`
}

func (x *TopWorkflowsByRunsCountRequest) Reset() {
	*x = TopWorkflowsByRunsCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_org_metrics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopWorkflowsByRunsCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopWorkflowsByRunsCountRequest) ProtoMessage() {}

func (x *TopWorkflowsByRunsCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_org_metrics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopWorkflowsByRunsCountRequest.ProtoReflect.Descriptor instead.
func (*TopWorkflowsByRunsCountRequest) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_org_metrics_proto_rawDescGZIP(), []int{1}
}

func (x *TopWorkflowsByRunsCountRequest) GetNumWorkflows() int32 {
	if x != nil {
		return x.NumWorkflows
	}
	return 0
}

func (x *TopWorkflowsByRunsCountRequest) GetTimeWindow() MetricsTimeWindow {
	if x != nil {
		return x.TimeWindow
	}
	return MetricsTimeWindow_METRICS_TIME_WINDOW_UNSPECIFIED
}

type TopWorkflowsByRunsCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*TopWorkflowsByRunsCountResponse_TotalByStatus `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *TopWorkflowsByRunsCountResponse) Reset() {
	*x = TopWorkflowsByRunsCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_org_metrics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopWorkflowsByRunsCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopWorkflowsByRunsCountResponse) ProtoMessage() {}

func (x *TopWorkflowsByRunsCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_org_metrics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopWorkflowsByRunsCountResponse.ProtoReflect.Descriptor instead.
func (*TopWorkflowsByRunsCountResponse) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_org_metrics_proto_rawDescGZIP(), []int{2}
}

func (x *TopWorkflowsByRunsCountResponse) GetResult() []*TopWorkflowsByRunsCountResponse_TotalByStatus {
	if x != nil {
		return x.Result
	}
	return nil
}

type OrgMetricsServiceTotalsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *OrgMetricsServiceTotalsResponse_Result `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *OrgMetricsServiceTotalsResponse) Reset() {
	*x = OrgMetricsServiceTotalsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_org_metrics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrgMetricsServiceTotalsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrgMetricsServiceTotalsResponse) ProtoMessage() {}

func (x *OrgMetricsServiceTotalsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_org_metrics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrgMetricsServiceTotalsResponse.ProtoReflect.Descriptor instead.
func (*OrgMetricsServiceTotalsResponse) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_org_metrics_proto_rawDescGZIP(), []int{3}
}

func (x *OrgMetricsServiceTotalsResponse) GetResult() *OrgMetricsServiceTotalsResponse_Result {
	if x != nil {
		return x.Result
	}
	return nil
}

type TopWorkflowsByRunsCountResponse_TotalByStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workflow *WorkflowItem `protobuf:"bytes,1,opt,name=workflow,proto3" json:"workflow,omitempty"`
	// Status -> [initialized, error, success]
	RunsTotalByStatus map[string]int32 `protobuf:"bytes,2,rep,name=runs_total_by_status,json=runsTotalByStatus,proto3" json:"runs_total_by_status,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *TopWorkflowsByRunsCountResponse_TotalByStatus) Reset() {
	*x = TopWorkflowsByRunsCountResponse_TotalByStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_org_metrics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopWorkflowsByRunsCountResponse_TotalByStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopWorkflowsByRunsCountResponse_TotalByStatus) ProtoMessage() {}

func (x *TopWorkflowsByRunsCountResponse_TotalByStatus) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_org_metrics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopWorkflowsByRunsCountResponse_TotalByStatus.ProtoReflect.Descriptor instead.
func (*TopWorkflowsByRunsCountResponse_TotalByStatus) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_org_metrics_proto_rawDescGZIP(), []int{2, 0}
}

func (x *TopWorkflowsByRunsCountResponse_TotalByStatus) GetWorkflow() *WorkflowItem {
	if x != nil {
		return x.Workflow
	}
	return nil
}

func (x *TopWorkflowsByRunsCountResponse_TotalByStatus) GetRunsTotalByStatus() map[string]int32 {
	if x != nil {
		return x.RunsTotalByStatus
	}
	return nil
}

type OrgMetricsServiceTotalsResponse_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RunsTotal int32 `protobuf:"varint,1,opt,name=runs_total,json=runsTotal,proto3" json:"runs_total,omitempty"`
	// Status -> [initialized, error, success]
	RunsTotalByStatus map[string]int32 `protobuf:"bytes,2,rep,name=runs_total_by_status,json=runsTotalByStatus,proto3" json:"runs_total_by_status,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// runner_type -> [generic, github_action, ...]
	RunsTotalByRunnerType map[string]int32 `protobuf:"bytes,3,rep,name=runs_total_by_runner_type,json=runsTotalByRunnerType,proto3" json:"runs_total_by_runner_type,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *OrgMetricsServiceTotalsResponse_Result) Reset() {
	*x = OrgMetricsServiceTotalsResponse_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_org_metrics_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrgMetricsServiceTotalsResponse_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrgMetricsServiceTotalsResponse_Result) ProtoMessage() {}

func (x *OrgMetricsServiceTotalsResponse_Result) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_org_metrics_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrgMetricsServiceTotalsResponse_Result.ProtoReflect.Descriptor instead.
func (*OrgMetricsServiceTotalsResponse_Result) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_org_metrics_proto_rawDescGZIP(), []int{3, 0}
}

func (x *OrgMetricsServiceTotalsResponse_Result) GetRunsTotal() int32 {
	if x != nil {
		return x.RunsTotal
	}
	return 0
}

func (x *OrgMetricsServiceTotalsResponse_Result) GetRunsTotalByStatus() map[string]int32 {
	if x != nil {
		return x.RunsTotalByStatus
	}
	return nil
}

func (x *OrgMetricsServiceTotalsResponse_Result) GetRunsTotalByRunnerType() map[string]int32 {
	if x != nil {
		return x.RunsTotalByRunnerType
	}
	return nil
}

var File_controlplane_v1_org_metrics_proto protoreflect.FileDescriptor

var file_controlplane_v1_org_metrics_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x72, 0x67, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x1e, 0x4f, 0x72, 0x67, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4d, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f,
	0x77, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x20, 0x00, 0x52, 0x0a, 0x74, 0x69, 0x6d,
	0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x22, 0x9f, 0x01, 0x0a, 0x1e, 0x54, 0x6f, 0x70, 0x57,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x42, 0x79, 0x52, 0x75, 0x6e, 0x73, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x0d, 0x6e, 0x75,
	0x6d, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x1a, 0x04, 0x18, 0x14, 0x28, 0x01, 0x52, 0x0c, 0x6e, 0x75,
	0x6d, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x12, 0x4d, 0x0a, 0x0b, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x22, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e,
	0x64, 0x6f, 0x77, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x20, 0x00, 0x52, 0x0a, 0x74,
	0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x22, 0x95, 0x03, 0x0a, 0x1f, 0x54, 0x6f,
	0x70, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x42, 0x79, 0x52, 0x75, 0x6e, 0x73,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x6f, 0x70, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x42, 0x79, 0x52, 0x75,
	0x6e, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x99, 0x02, 0x0a, 0x0d, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x42,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x12, 0x86, 0x01, 0x0a, 0x14, 0x72, 0x75, 0x6e, 0x73, 0x5f, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x62, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x55, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73,
	0x42, 0x79, 0x52, 0x75, 0x6e, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x52, 0x75, 0x6e, 0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x11, 0x72, 0x75, 0x6e, 0x73, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x42, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x44, 0x0a, 0x16, 0x52,
	0x75, 0x6e, 0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0xbc, 0x04, 0x0a, 0x1f, 0x4f, 0x72, 0x67, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0xc7, 0x03, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x75, 0x6e, 0x73, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x75, 0x6e, 0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x7f, 0x0a, 0x14, 0x72, 0x75, 0x6e, 0x73, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x62,
	0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4e,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x72, 0x67, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x52, 0x75, 0x6e, 0x73, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x42, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x11,
	0x72, 0x75, 0x6e, 0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x8c, 0x01, 0x0a, 0x19, 0x72, 0x75, 0x6e, 0x73, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x62, 0x79, 0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x52, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x52,
	0x75, 0x6e, 0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x79, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x15, 0x72, 0x75, 0x6e, 0x73, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x42, 0x79, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x1a, 0x44, 0x0a, 0x16, 0x52, 0x75, 0x6e, 0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x48, 0x0a, 0x1a, 0x52, 0x75, 0x6e, 0x73, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x42, 0x79, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x2a, 0xa5, 0x01, 0x0a, 0x11, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x54, 0x69, 0x6d, 0x65,
	0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x23, 0x0a, 0x1f, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43,
	0x53, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x24, 0x0a, 0x20, 0x4d,
	0x45, 0x54, 0x52, 0x49, 0x43, 0x53, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44,
	0x4f, 0x57, 0x5f, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x33, 0x30, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10,
	0x01, 0x12, 0x23, 0x0a, 0x1f, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x53, 0x5f, 0x54, 0x49, 0x4d,
	0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x37, 0x5f,
	0x44, 0x41, 0x59, 0x53, 0x10, 0x02, 0x12, 0x20, 0x0a, 0x1c, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43,
	0x53, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x4c, 0x41,
	0x53, 0x54, 0x5f, 0x44, 0x41, 0x59, 0x10, 0x03, 0x32, 0xfe, 0x01, 0x0a, 0x11, 0x4f, 0x72, 0x67,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6b,
	0x0a, 0x06, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x73, 0x12, 0x2f, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7c, 0x0a, 0x17, 0x54,
	0x6f, 0x70, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x42, 0x79, 0x52, 0x75, 0x6e,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x73, 0x42, 0x79, 0x52, 0x75, 0x6e, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x42, 0x79, 0x52, 0x75, 0x6e, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c, 0x6f, 0x6f,
	0x70, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c, 0x6f, 0x6f, 0x70, 0x2f,
	0x61, 0x70, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controlplane_v1_org_metrics_proto_rawDescOnce sync.Once
	file_controlplane_v1_org_metrics_proto_rawDescData = file_controlplane_v1_org_metrics_proto_rawDesc
)

func file_controlplane_v1_org_metrics_proto_rawDescGZIP() []byte {
	file_controlplane_v1_org_metrics_proto_rawDescOnce.Do(func() {
		file_controlplane_v1_org_metrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_controlplane_v1_org_metrics_proto_rawDescData)
	})
	return file_controlplane_v1_org_metrics_proto_rawDescData
}

var file_controlplane_v1_org_metrics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_controlplane_v1_org_metrics_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_controlplane_v1_org_metrics_proto_goTypes = []interface{}{
	(MetricsTimeWindow)(0),                                // 0: controlplane.v1.MetricsTimeWindow
	(*OrgMetricsServiceTotalsRequest)(nil),                // 1: controlplane.v1.OrgMetricsServiceTotalsRequest
	(*TopWorkflowsByRunsCountRequest)(nil),                // 2: controlplane.v1.TopWorkflowsByRunsCountRequest
	(*TopWorkflowsByRunsCountResponse)(nil),               // 3: controlplane.v1.TopWorkflowsByRunsCountResponse
	(*OrgMetricsServiceTotalsResponse)(nil),               // 4: controlplane.v1.OrgMetricsServiceTotalsResponse
	(*TopWorkflowsByRunsCountResponse_TotalByStatus)(nil), // 5: controlplane.v1.TopWorkflowsByRunsCountResponse.TotalByStatus
	nil, // 6: controlplane.v1.TopWorkflowsByRunsCountResponse.TotalByStatus.RunsTotalByStatusEntry
	(*OrgMetricsServiceTotalsResponse_Result)(nil), // 7: controlplane.v1.OrgMetricsServiceTotalsResponse.Result
	nil,                  // 8: controlplane.v1.OrgMetricsServiceTotalsResponse.Result.RunsTotalByStatusEntry
	nil,                  // 9: controlplane.v1.OrgMetricsServiceTotalsResponse.Result.RunsTotalByRunnerTypeEntry
	(*WorkflowItem)(nil), // 10: controlplane.v1.WorkflowItem
}
var file_controlplane_v1_org_metrics_proto_depIdxs = []int32{
	0,  // 0: controlplane.v1.OrgMetricsServiceTotalsRequest.time_window:type_name -> controlplane.v1.MetricsTimeWindow
	0,  // 1: controlplane.v1.TopWorkflowsByRunsCountRequest.time_window:type_name -> controlplane.v1.MetricsTimeWindow
	5,  // 2: controlplane.v1.TopWorkflowsByRunsCountResponse.result:type_name -> controlplane.v1.TopWorkflowsByRunsCountResponse.TotalByStatus
	7,  // 3: controlplane.v1.OrgMetricsServiceTotalsResponse.result:type_name -> controlplane.v1.OrgMetricsServiceTotalsResponse.Result
	10, // 4: controlplane.v1.TopWorkflowsByRunsCountResponse.TotalByStatus.workflow:type_name -> controlplane.v1.WorkflowItem
	6,  // 5: controlplane.v1.TopWorkflowsByRunsCountResponse.TotalByStatus.runs_total_by_status:type_name -> controlplane.v1.TopWorkflowsByRunsCountResponse.TotalByStatus.RunsTotalByStatusEntry
	8,  // 6: controlplane.v1.OrgMetricsServiceTotalsResponse.Result.runs_total_by_status:type_name -> controlplane.v1.OrgMetricsServiceTotalsResponse.Result.RunsTotalByStatusEntry
	9,  // 7: controlplane.v1.OrgMetricsServiceTotalsResponse.Result.runs_total_by_runner_type:type_name -> controlplane.v1.OrgMetricsServiceTotalsResponse.Result.RunsTotalByRunnerTypeEntry
	1,  // 8: controlplane.v1.OrgMetricsService.Totals:input_type -> controlplane.v1.OrgMetricsServiceTotalsRequest
	2,  // 9: controlplane.v1.OrgMetricsService.TopWorkflowsByRunsCount:input_type -> controlplane.v1.TopWorkflowsByRunsCountRequest
	4,  // 10: controlplane.v1.OrgMetricsService.Totals:output_type -> controlplane.v1.OrgMetricsServiceTotalsResponse
	3,  // 11: controlplane.v1.OrgMetricsService.TopWorkflowsByRunsCount:output_type -> controlplane.v1.TopWorkflowsByRunsCountResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_controlplane_v1_org_metrics_proto_init() }
func file_controlplane_v1_org_metrics_proto_init() {
	if File_controlplane_v1_org_metrics_proto != nil {
		return
	}
	file_controlplane_v1_response_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_controlplane_v1_org_metrics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrgMetricsServiceTotalsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_controlplane_v1_org_metrics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopWorkflowsByRunsCountRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_controlplane_v1_org_metrics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopWorkflowsByRunsCountResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_controlplane_v1_org_metrics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrgMetricsServiceTotalsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_controlplane_v1_org_metrics_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopWorkflowsByRunsCountResponse_TotalByStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_controlplane_v1_org_metrics_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrgMetricsServiceTotalsResponse_Result); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_controlplane_v1_org_metrics_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_controlplane_v1_org_metrics_proto_goTypes,
		DependencyIndexes: file_controlplane_v1_org_metrics_proto_depIdxs,
		EnumInfos:         file_controlplane_v1_org_metrics_proto_enumTypes,
		MessageInfos:      file_controlplane_v1_org_metrics_proto_msgTypes,
	}.Build()
	File_controlplane_v1_org_metrics_proto = out.File
	file_controlplane_v1_org_metrics_proto_rawDesc = nil
	file_controlplane_v1_org_metrics_proto_goTypes = nil
	file_controlplane_v1_org_metrics_proto_depIdxs = nil
}
