// Code generated with goa v3.14.4, DO NOT EDIT.
//
// tester protocol buffer definition
//
// Command:
// $ goa gen goa.design/clue/example/weather/services/tester/design -o
// services/tester

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.25.1
// source: goagen_tester_tester.proto

package weather_testerpb

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

type TestAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Tests to run
	Include []string `protobuf:"bytes,1,rep,name=include,proto3" json:"include,omitempty"`
	// Tests to exclude
	Exclude []string `protobuf:"bytes,2,rep,name=exclude,proto3" json:"exclude,omitempty"`
}

func (x *TestAllRequest) Reset() {
	*x = TestAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goagen_tester_tester_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestAllRequest) ProtoMessage() {}

func (x *TestAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_tester_tester_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestAllRequest.ProtoReflect.Descriptor instead.
func (*TestAllRequest) Descriptor() ([]byte, []int) {
	return file_goagen_tester_tester_proto_rawDescGZIP(), []int{0}
}

func (x *TestAllRequest) GetInclude() []string {
	if x != nil {
		return x.Include
	}
	return nil
}

func (x *TestAllRequest) GetExclude() []string {
	if x != nil {
		return x.Exclude
	}
	return nil
}

type TestAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Test collections
	Collections []*TestCollection `protobuf:"bytes,1,rep,name=collections,proto3" json:"collections,omitempty"`
	// Duration of the tests in ms
	Duration int64 `protobuf:"zigzag64,2,opt,name=duration,proto3" json:"duration,omitempty"`
	// Number of tests that passed
	PassCount int32 `protobuf:"zigzag32,3,opt,name=pass_count,json=passCount,proto3" json:"pass_count,omitempty"`
	// Number of tests that failed
	FailCount int32 `protobuf:"zigzag32,4,opt,name=fail_count,json=failCount,proto3" json:"fail_count,omitempty"`
}

func (x *TestAllResponse) Reset() {
	*x = TestAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goagen_tester_tester_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestAllResponse) ProtoMessage() {}

func (x *TestAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_tester_tester_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestAllResponse.ProtoReflect.Descriptor instead.
func (*TestAllResponse) Descriptor() ([]byte, []int) {
	return file_goagen_tester_tester_proto_rawDescGZIP(), []int{1}
}

func (x *TestAllResponse) GetCollections() []*TestCollection {
	if x != nil {
		return x.Collections
	}
	return nil
}

func (x *TestAllResponse) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *TestAllResponse) GetPassCount() int32 {
	if x != nil {
		return x.PassCount
	}
	return 0
}

func (x *TestAllResponse) GetFailCount() int32 {
	if x != nil {
		return x.FailCount
	}
	return 0
}

// Collection of test results for grouping by service
type TestCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the test collection
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Test results
	Results []*TestResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	// Duration of the tests in ms
	Duration int64 `protobuf:"zigzag64,3,opt,name=duration,proto3" json:"duration,omitempty"`
	// Number of tests that passed
	PassCount int32 `protobuf:"zigzag32,4,opt,name=pass_count,json=passCount,proto3" json:"pass_count,omitempty"`
	// Number of tests that failed
	FailCount int32 `protobuf:"zigzag32,5,opt,name=fail_count,json=failCount,proto3" json:"fail_count,omitempty"`
}

func (x *TestCollection) Reset() {
	*x = TestCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goagen_tester_tester_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestCollection) ProtoMessage() {}

func (x *TestCollection) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_tester_tester_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestCollection.ProtoReflect.Descriptor instead.
func (*TestCollection) Descriptor() ([]byte, []int) {
	return file_goagen_tester_tester_proto_rawDescGZIP(), []int{2}
}

func (x *TestCollection) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TestCollection) GetResults() []*TestResult {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *TestCollection) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *TestCollection) GetPassCount() int32 {
	if x != nil {
		return x.PassCount
	}
	return 0
}

func (x *TestCollection) GetFailCount() int32 {
	if x != nil {
		return x.FailCount
	}
	return 0
}

// Test result for a single test
type TestResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the test
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Status of the test
	Passed bool `protobuf:"varint,2,opt,name=passed,proto3" json:"passed,omitempty"`
	// Error message if the test failed
	Error *string `protobuf:"bytes,3,opt,name=error,proto3,oneof" json:"error,omitempty"`
	// Duration of the test in ms
	Duration int64 `protobuf:"zigzag64,4,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *TestResult) Reset() {
	*x = TestResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goagen_tester_tester_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResult) ProtoMessage() {}

func (x *TestResult) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_tester_tester_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResult.ProtoReflect.Descriptor instead.
func (*TestResult) Descriptor() ([]byte, []int) {
	return file_goagen_tester_tester_proto_rawDescGZIP(), []int{3}
}

func (x *TestResult) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TestResult) GetPassed() bool {
	if x != nil {
		return x.Passed
	}
	return false
}

func (x *TestResult) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

func (x *TestResult) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

type TestSmokeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TestSmokeRequest) Reset() {
	*x = TestSmokeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goagen_tester_tester_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestSmokeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestSmokeRequest) ProtoMessage() {}

func (x *TestSmokeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_tester_tester_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestSmokeRequest.ProtoReflect.Descriptor instead.
func (*TestSmokeRequest) Descriptor() ([]byte, []int) {
	return file_goagen_tester_tester_proto_rawDescGZIP(), []int{4}
}

type TestSmokeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Test collections
	Collections []*TestCollection `protobuf:"bytes,1,rep,name=collections,proto3" json:"collections,omitempty"`
	// Duration of the tests in ms
	Duration int64 `protobuf:"zigzag64,2,opt,name=duration,proto3" json:"duration,omitempty"`
	// Number of tests that passed
	PassCount int32 `protobuf:"zigzag32,3,opt,name=pass_count,json=passCount,proto3" json:"pass_count,omitempty"`
	// Number of tests that failed
	FailCount int32 `protobuf:"zigzag32,4,opt,name=fail_count,json=failCount,proto3" json:"fail_count,omitempty"`
}

func (x *TestSmokeResponse) Reset() {
	*x = TestSmokeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goagen_tester_tester_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestSmokeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestSmokeResponse) ProtoMessage() {}

func (x *TestSmokeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_tester_tester_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestSmokeResponse.ProtoReflect.Descriptor instead.
func (*TestSmokeResponse) Descriptor() ([]byte, []int) {
	return file_goagen_tester_tester_proto_rawDescGZIP(), []int{5}
}

func (x *TestSmokeResponse) GetCollections() []*TestCollection {
	if x != nil {
		return x.Collections
	}
	return nil
}

func (x *TestSmokeResponse) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *TestSmokeResponse) GetPassCount() int32 {
	if x != nil {
		return x.PassCount
	}
	return 0
}

func (x *TestSmokeResponse) GetFailCount() int32 {
	if x != nil {
		return x.FailCount
	}
	return 0
}

type TestForecasterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TestForecasterRequest) Reset() {
	*x = TestForecasterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goagen_tester_tester_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestForecasterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestForecasterRequest) ProtoMessage() {}

func (x *TestForecasterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_tester_tester_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestForecasterRequest.ProtoReflect.Descriptor instead.
func (*TestForecasterRequest) Descriptor() ([]byte, []int) {
	return file_goagen_tester_tester_proto_rawDescGZIP(), []int{6}
}

type TestForecasterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Test collections
	Collections []*TestCollection `protobuf:"bytes,1,rep,name=collections,proto3" json:"collections,omitempty"`
	// Duration of the tests in ms
	Duration int64 `protobuf:"zigzag64,2,opt,name=duration,proto3" json:"duration,omitempty"`
	// Number of tests that passed
	PassCount int32 `protobuf:"zigzag32,3,opt,name=pass_count,json=passCount,proto3" json:"pass_count,omitempty"`
	// Number of tests that failed
	FailCount int32 `protobuf:"zigzag32,4,opt,name=fail_count,json=failCount,proto3" json:"fail_count,omitempty"`
}

func (x *TestForecasterResponse) Reset() {
	*x = TestForecasterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goagen_tester_tester_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestForecasterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestForecasterResponse) ProtoMessage() {}

func (x *TestForecasterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_tester_tester_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestForecasterResponse.ProtoReflect.Descriptor instead.
func (*TestForecasterResponse) Descriptor() ([]byte, []int) {
	return file_goagen_tester_tester_proto_rawDescGZIP(), []int{7}
}

func (x *TestForecasterResponse) GetCollections() []*TestCollection {
	if x != nil {
		return x.Collections
	}
	return nil
}

func (x *TestForecasterResponse) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *TestForecasterResponse) GetPassCount() int32 {
	if x != nil {
		return x.PassCount
	}
	return 0
}

func (x *TestForecasterResponse) GetFailCount() int32 {
	if x != nil {
		return x.FailCount
	}
	return 0
}

type TestLocatorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TestLocatorRequest) Reset() {
	*x = TestLocatorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goagen_tester_tester_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestLocatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestLocatorRequest) ProtoMessage() {}

func (x *TestLocatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_tester_tester_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestLocatorRequest.ProtoReflect.Descriptor instead.
func (*TestLocatorRequest) Descriptor() ([]byte, []int) {
	return file_goagen_tester_tester_proto_rawDescGZIP(), []int{8}
}

type TestLocatorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Test collections
	Collections []*TestCollection `protobuf:"bytes,1,rep,name=collections,proto3" json:"collections,omitempty"`
	// Duration of the tests in ms
	Duration int64 `protobuf:"zigzag64,2,opt,name=duration,proto3" json:"duration,omitempty"`
	// Number of tests that passed
	PassCount int32 `protobuf:"zigzag32,3,opt,name=pass_count,json=passCount,proto3" json:"pass_count,omitempty"`
	// Number of tests that failed
	FailCount int32 `protobuf:"zigzag32,4,opt,name=fail_count,json=failCount,proto3" json:"fail_count,omitempty"`
}

func (x *TestLocatorResponse) Reset() {
	*x = TestLocatorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goagen_tester_tester_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestLocatorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestLocatorResponse) ProtoMessage() {}

func (x *TestLocatorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_tester_tester_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestLocatorResponse.ProtoReflect.Descriptor instead.
func (*TestLocatorResponse) Descriptor() ([]byte, []int) {
	return file_goagen_tester_tester_proto_rawDescGZIP(), []int{9}
}

func (x *TestLocatorResponse) GetCollections() []*TestCollection {
	if x != nil {
		return x.Collections
	}
	return nil
}

func (x *TestLocatorResponse) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *TestLocatorResponse) GetPassCount() int32 {
	if x != nil {
		return x.PassCount
	}
	return 0
}

func (x *TestLocatorResponse) GetFailCount() int32 {
	if x != nil {
		return x.FailCount
	}
	return 0
}

var File_goagen_tester_tester_proto protoreflect.FileDescriptor

var file_goagen_tester_tester_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x67, 0x6f, 0x61, 0x67, 0x65, 0x6e, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x74, 0x65, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x77, 0x65,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x65, 0x72, 0x22, 0x44, 0x0a, 0x0e,
	0x54, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x65, 0x78, 0x63, 0x6c, 0x75,
	0x64, 0x65, 0x22, 0xad, 0x01, 0x0a, 0x0f, 0x54, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x77, 0x65,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x12, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x11, 0x52, 0x09, 0x70, 0x61, 0x73, 0x73, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x61, 0x69, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x11, 0x52, 0x09, 0x66, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0xb4, 0x01, 0x0a, 0x0e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x77, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x12, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x11, 0x52,
	0x09, 0x70, 0x61, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x61,
	0x69, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x11, 0x52, 0x09,
	0x66, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x79, 0x0a, 0x0a, 0x54, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x61, 0x73, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x61, 0x73,
	0x73, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x12,
	0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x12, 0x0a, 0x10, 0x54, 0x65, 0x73, 0x74, 0x53, 0x6d, 0x6f, 0x6b,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xaf, 0x01, 0x0a, 0x11, 0x54, 0x65, 0x73,
	0x74, 0x53, 0x6d, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40,
	0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x65,
	0x73, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x12, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x61, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x11,
	0x52, 0x09, 0x70, 0x61, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66,
	0x61, 0x69, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x11, 0x52,
	0x09, 0x66, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x54, 0x65,
	0x73, 0x74, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0xb4, 0x01, 0x0a, 0x16, 0x54, 0x65, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x65,
	0x63, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40,
	0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x65,
	0x73, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x12, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x61, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x11,
	0x52, 0x09, 0x70, 0x61, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66,
	0x61, 0x69, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x11, 0x52,
	0x09, 0x66, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x54, 0x65,
	0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0xb1, 0x01, 0x0a, 0x13, 0x54, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x12, 0x52, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x73, 0x73, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x11, 0x52, 0x09, 0x70, 0x61, 0x73, 0x73,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x61, 0x69, 0x6c, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x11, 0x52, 0x09, 0x66, 0x61, 0x69, 0x6c, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x32, 0xdf, 0x02, 0x0a, 0x06, 0x54, 0x65, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x4a, 0x0a, 0x07, 0x54, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x1e, 0x2e, 0x77, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x77, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x09, 0x54,
	0x65, 0x73, 0x74, 0x53, 0x6d, 0x6f, 0x6b, 0x65, 0x12, 0x20, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x53, 0x6d,
	0x6f, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x77, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x53, 0x6d, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a,
	0x0e, 0x54, 0x65, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x25, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x5f, 0x74, 0x65, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x65,
	0x63, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56,
	0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x22, 0x2e,
	0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x13, 0x5a, 0x11, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_goagen_tester_tester_proto_rawDescOnce sync.Once
	file_goagen_tester_tester_proto_rawDescData = file_goagen_tester_tester_proto_rawDesc
)

func file_goagen_tester_tester_proto_rawDescGZIP() []byte {
	file_goagen_tester_tester_proto_rawDescOnce.Do(func() {
		file_goagen_tester_tester_proto_rawDescData = protoimpl.X.CompressGZIP(file_goagen_tester_tester_proto_rawDescData)
	})
	return file_goagen_tester_tester_proto_rawDescData
}

var file_goagen_tester_tester_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_goagen_tester_tester_proto_goTypes = []interface{}{
	(*TestAllRequest)(nil),         // 0: weather_tester.TestAllRequest
	(*TestAllResponse)(nil),        // 1: weather_tester.TestAllResponse
	(*TestCollection)(nil),         // 2: weather_tester.TestCollection
	(*TestResult)(nil),             // 3: weather_tester.TestResult
	(*TestSmokeRequest)(nil),       // 4: weather_tester.TestSmokeRequest
	(*TestSmokeResponse)(nil),      // 5: weather_tester.TestSmokeResponse
	(*TestForecasterRequest)(nil),  // 6: weather_tester.TestForecasterRequest
	(*TestForecasterResponse)(nil), // 7: weather_tester.TestForecasterResponse
	(*TestLocatorRequest)(nil),     // 8: weather_tester.TestLocatorRequest
	(*TestLocatorResponse)(nil),    // 9: weather_tester.TestLocatorResponse
}
var file_goagen_tester_tester_proto_depIdxs = []int32{
	2, // 0: weather_tester.TestAllResponse.collections:type_name -> weather_tester.TestCollection
	3, // 1: weather_tester.TestCollection.results:type_name -> weather_tester.TestResult
	2, // 2: weather_tester.TestSmokeResponse.collections:type_name -> weather_tester.TestCollection
	2, // 3: weather_tester.TestForecasterResponse.collections:type_name -> weather_tester.TestCollection
	2, // 4: weather_tester.TestLocatorResponse.collections:type_name -> weather_tester.TestCollection
	0, // 5: weather_tester.Tester.TestAll:input_type -> weather_tester.TestAllRequest
	4, // 6: weather_tester.Tester.TestSmoke:input_type -> weather_tester.TestSmokeRequest
	6, // 7: weather_tester.Tester.TestForecaster:input_type -> weather_tester.TestForecasterRequest
	8, // 8: weather_tester.Tester.TestLocator:input_type -> weather_tester.TestLocatorRequest
	1, // 9: weather_tester.Tester.TestAll:output_type -> weather_tester.TestAllResponse
	5, // 10: weather_tester.Tester.TestSmoke:output_type -> weather_tester.TestSmokeResponse
	7, // 11: weather_tester.Tester.TestForecaster:output_type -> weather_tester.TestForecasterResponse
	9, // 12: weather_tester.Tester.TestLocator:output_type -> weather_tester.TestLocatorResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_goagen_tester_tester_proto_init() }
func file_goagen_tester_tester_proto_init() {
	if File_goagen_tester_tester_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goagen_tester_tester_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestAllRequest); i {
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
		file_goagen_tester_tester_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestAllResponse); i {
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
		file_goagen_tester_tester_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestCollection); i {
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
		file_goagen_tester_tester_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestResult); i {
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
		file_goagen_tester_tester_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestSmokeRequest); i {
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
		file_goagen_tester_tester_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestSmokeResponse); i {
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
		file_goagen_tester_tester_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestForecasterRequest); i {
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
		file_goagen_tester_tester_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestForecasterResponse); i {
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
		file_goagen_tester_tester_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestLocatorRequest); i {
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
		file_goagen_tester_tester_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestLocatorResponse); i {
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
	file_goagen_tester_tester_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_goagen_tester_tester_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goagen_tester_tester_proto_goTypes,
		DependencyIndexes: file_goagen_tester_tester_proto_depIdxs,
		MessageInfos:      file_goagen_tester_tester_proto_msgTypes,
	}.Build()
	File_goagen_tester_tester_proto = out.File
	file_goagen_tester_tester_proto_rawDesc = nil
	file_goagen_tester_tester_proto_goTypes = nil
	file_goagen_tester_tester_proto_depIdxs = nil
}
