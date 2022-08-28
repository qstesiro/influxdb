// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: internal/internal.proto

package query

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

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          *string        `protobuf:"bytes,1,req,name=Name" json:"Name,omitempty"`
	Tags          *string        `protobuf:"bytes,2,req,name=Tags" json:"Tags,omitempty"`
	Time          *int64         `protobuf:"varint,3,req,name=Time" json:"Time,omitempty"`
	Nil           *bool          `protobuf:"varint,4,req,name=Nil" json:"Nil,omitempty"`
	Aux           []*Aux         `protobuf:"bytes,5,rep,name=Aux" json:"Aux,omitempty"`
	Aggregated    *uint32        `protobuf:"varint,6,opt,name=Aggregated" json:"Aggregated,omitempty"`
	FloatValue    *float64       `protobuf:"fixed64,7,opt,name=FloatValue" json:"FloatValue,omitempty"`
	IntegerValue  *int64         `protobuf:"varint,8,opt,name=IntegerValue" json:"IntegerValue,omitempty"`
	StringValue   *string        `protobuf:"bytes,9,opt,name=StringValue" json:"StringValue,omitempty"`
	BooleanValue  *bool          `protobuf:"varint,10,opt,name=BooleanValue" json:"BooleanValue,omitempty"`
	UnsignedValue *uint64        `protobuf:"varint,12,opt,name=UnsignedValue" json:"UnsignedValue,omitempty"`
	Stats         *IteratorStats `protobuf:"bytes,11,opt,name=Stats" json:"Stats,omitempty"`
	Trace         []byte         `protobuf:"bytes,13,opt,name=Trace" json:"Trace,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_internal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_internal_internal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_internal_internal_proto_rawDescGZIP(), []int{0}
}

func (x *Point) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Point) GetTags() string {
	if x != nil && x.Tags != nil {
		return *x.Tags
	}
	return ""
}

func (x *Point) GetTime() int64 {
	if x != nil && x.Time != nil {
		return *x.Time
	}
	return 0
}

func (x *Point) GetNil() bool {
	if x != nil && x.Nil != nil {
		return *x.Nil
	}
	return false
}

func (x *Point) GetAux() []*Aux {
	if x != nil {
		return x.Aux
	}
	return nil
}

func (x *Point) GetAggregated() uint32 {
	if x != nil && x.Aggregated != nil {
		return *x.Aggregated
	}
	return 0
}

func (x *Point) GetFloatValue() float64 {
	if x != nil && x.FloatValue != nil {
		return *x.FloatValue
	}
	return 0
}

func (x *Point) GetIntegerValue() int64 {
	if x != nil && x.IntegerValue != nil {
		return *x.IntegerValue
	}
	return 0
}

func (x *Point) GetStringValue() string {
	if x != nil && x.StringValue != nil {
		return *x.StringValue
	}
	return ""
}

func (x *Point) GetBooleanValue() bool {
	if x != nil && x.BooleanValue != nil {
		return *x.BooleanValue
	}
	return false
}

func (x *Point) GetUnsignedValue() uint64 {
	if x != nil && x.UnsignedValue != nil {
		return *x.UnsignedValue
	}
	return 0
}

func (x *Point) GetStats() *IteratorStats {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *Point) GetTrace() []byte {
	if x != nil {
		return x.Trace
	}
	return nil
}

type Aux struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataType      *int32   `protobuf:"varint,1,req,name=DataType" json:"DataType,omitempty"`
	FloatValue    *float64 `protobuf:"fixed64,2,opt,name=FloatValue" json:"FloatValue,omitempty"`
	IntegerValue  *int64   `protobuf:"varint,3,opt,name=IntegerValue" json:"IntegerValue,omitempty"`
	StringValue   *string  `protobuf:"bytes,4,opt,name=StringValue" json:"StringValue,omitempty"`
	BooleanValue  *bool    `protobuf:"varint,5,opt,name=BooleanValue" json:"BooleanValue,omitempty"`
	UnsignedValue *uint64  `protobuf:"varint,6,opt,name=UnsignedValue" json:"UnsignedValue,omitempty"`
}

func (x *Aux) Reset() {
	*x = Aux{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_internal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Aux) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Aux) ProtoMessage() {}

func (x *Aux) ProtoReflect() protoreflect.Message {
	mi := &file_internal_internal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Aux.ProtoReflect.Descriptor instead.
func (*Aux) Descriptor() ([]byte, []int) {
	return file_internal_internal_proto_rawDescGZIP(), []int{1}
}

func (x *Aux) GetDataType() int32 {
	if x != nil && x.DataType != nil {
		return *x.DataType
	}
	return 0
}

func (x *Aux) GetFloatValue() float64 {
	if x != nil && x.FloatValue != nil {
		return *x.FloatValue
	}
	return 0
}

func (x *Aux) GetIntegerValue() int64 {
	if x != nil && x.IntegerValue != nil {
		return *x.IntegerValue
	}
	return 0
}

func (x *Aux) GetStringValue() string {
	if x != nil && x.StringValue != nil {
		return *x.StringValue
	}
	return ""
}

func (x *Aux) GetBooleanValue() bool {
	if x != nil && x.BooleanValue != nil {
		return *x.BooleanValue
	}
	return false
}

func (x *Aux) GetUnsignedValue() uint64 {
	if x != nil && x.UnsignedValue != nil {
		return *x.UnsignedValue
	}
	return 0
}

type IteratorOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expr       *string        `protobuf:"bytes,1,opt,name=Expr" json:"Expr,omitempty"`
	Aux        []string       `protobuf:"bytes,2,rep,name=Aux" json:"Aux,omitempty"`
	Fields     []*VarRef      `protobuf:"bytes,17,rep,name=Fields" json:"Fields,omitempty"`
	Sources    []*Measurement `protobuf:"bytes,3,rep,name=Sources" json:"Sources,omitempty"`
	Interval   *Interval      `protobuf:"bytes,4,opt,name=Interval" json:"Interval,omitempty"`
	Dimensions []string       `protobuf:"bytes,5,rep,name=Dimensions" json:"Dimensions,omitempty"`
	GroupBy    []string       `protobuf:"bytes,19,rep,name=GroupBy" json:"GroupBy,omitempty"`
	Fill       *int32         `protobuf:"varint,6,opt,name=Fill" json:"Fill,omitempty"`
	FillValue  *float64       `protobuf:"fixed64,7,opt,name=FillValue" json:"FillValue,omitempty"`
	Condition  *string        `protobuf:"bytes,8,opt,name=Condition" json:"Condition,omitempty"`
	StartTime  *int64         `protobuf:"varint,9,opt,name=StartTime" json:"StartTime,omitempty"`
	EndTime    *int64         `protobuf:"varint,10,opt,name=EndTime" json:"EndTime,omitempty"`
	Location   *string        `protobuf:"bytes,21,opt,name=Location" json:"Location,omitempty"`
	Ascending  *bool          `protobuf:"varint,11,opt,name=Ascending" json:"Ascending,omitempty"`
	Limit      *int64         `protobuf:"varint,12,opt,name=Limit" json:"Limit,omitempty"`
	Offset     *int64         `protobuf:"varint,13,opt,name=Offset" json:"Offset,omitempty"`
	SLimit     *int64         `protobuf:"varint,14,opt,name=SLimit" json:"SLimit,omitempty"`
	SOffset    *int64         `protobuf:"varint,15,opt,name=SOffset" json:"SOffset,omitempty"`
	StripName  *bool          `protobuf:"varint,22,opt,name=StripName" json:"StripName,omitempty"`
	Dedupe     *bool          `protobuf:"varint,16,opt,name=Dedupe" json:"Dedupe,omitempty"`
	MaxSeriesN *int64         `protobuf:"varint,18,opt,name=MaxSeriesN" json:"MaxSeriesN,omitempty"`
	Ordered    *bool          `protobuf:"varint,20,opt,name=Ordered" json:"Ordered,omitempty"`
}

func (x *IteratorOptions) Reset() {
	*x = IteratorOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_internal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IteratorOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IteratorOptions) ProtoMessage() {}

func (x *IteratorOptions) ProtoReflect() protoreflect.Message {
	mi := &file_internal_internal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IteratorOptions.ProtoReflect.Descriptor instead.
func (*IteratorOptions) Descriptor() ([]byte, []int) {
	return file_internal_internal_proto_rawDescGZIP(), []int{2}
}

func (x *IteratorOptions) GetExpr() string {
	if x != nil && x.Expr != nil {
		return *x.Expr
	}
	return ""
}

func (x *IteratorOptions) GetAux() []string {
	if x != nil {
		return x.Aux
	}
	return nil
}

func (x *IteratorOptions) GetFields() []*VarRef {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *IteratorOptions) GetSources() []*Measurement {
	if x != nil {
		return x.Sources
	}
	return nil
}

func (x *IteratorOptions) GetInterval() *Interval {
	if x != nil {
		return x.Interval
	}
	return nil
}

func (x *IteratorOptions) GetDimensions() []string {
	if x != nil {
		return x.Dimensions
	}
	return nil
}

func (x *IteratorOptions) GetGroupBy() []string {
	if x != nil {
		return x.GroupBy
	}
	return nil
}

func (x *IteratorOptions) GetFill() int32 {
	if x != nil && x.Fill != nil {
		return *x.Fill
	}
	return 0
}

func (x *IteratorOptions) GetFillValue() float64 {
	if x != nil && x.FillValue != nil {
		return *x.FillValue
	}
	return 0
}

func (x *IteratorOptions) GetCondition() string {
	if x != nil && x.Condition != nil {
		return *x.Condition
	}
	return ""
}

func (x *IteratorOptions) GetStartTime() int64 {
	if x != nil && x.StartTime != nil {
		return *x.StartTime
	}
	return 0
}

func (x *IteratorOptions) GetEndTime() int64 {
	if x != nil && x.EndTime != nil {
		return *x.EndTime
	}
	return 0
}

func (x *IteratorOptions) GetLocation() string {
	if x != nil && x.Location != nil {
		return *x.Location
	}
	return ""
}

func (x *IteratorOptions) GetAscending() bool {
	if x != nil && x.Ascending != nil {
		return *x.Ascending
	}
	return false
}

func (x *IteratorOptions) GetLimit() int64 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *IteratorOptions) GetOffset() int64 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

func (x *IteratorOptions) GetSLimit() int64 {
	if x != nil && x.SLimit != nil {
		return *x.SLimit
	}
	return 0
}

func (x *IteratorOptions) GetSOffset() int64 {
	if x != nil && x.SOffset != nil {
		return *x.SOffset
	}
	return 0
}

func (x *IteratorOptions) GetStripName() bool {
	if x != nil && x.StripName != nil {
		return *x.StripName
	}
	return false
}

func (x *IteratorOptions) GetDedupe() bool {
	if x != nil && x.Dedupe != nil {
		return *x.Dedupe
	}
	return false
}

func (x *IteratorOptions) GetMaxSeriesN() int64 {
	if x != nil && x.MaxSeriesN != nil {
		return *x.MaxSeriesN
	}
	return 0
}

func (x *IteratorOptions) GetOrdered() bool {
	if x != nil && x.Ordered != nil {
		return *x.Ordered
	}
	return false
}

type Measurements struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Measurement `protobuf:"bytes,1,rep,name=Items" json:"Items,omitempty"`
}

func (x *Measurements) Reset() {
	*x = Measurements{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_internal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Measurements) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Measurements) ProtoMessage() {}

func (x *Measurements) ProtoReflect() protoreflect.Message {
	mi := &file_internal_internal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Measurements.ProtoReflect.Descriptor instead.
func (*Measurements) Descriptor() ([]byte, []int) {
	return file_internal_internal_proto_rawDescGZIP(), []int{3}
}

func (x *Measurements) GetItems() []*Measurement {
	if x != nil {
		return x.Items
	}
	return nil
}

type Measurement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database        *string `protobuf:"bytes,1,opt,name=Database" json:"Database,omitempty"`
	RetentionPolicy *string `protobuf:"bytes,2,opt,name=RetentionPolicy" json:"RetentionPolicy,omitempty"`
	Name            *string `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	Regex           *string `protobuf:"bytes,4,opt,name=Regex" json:"Regex,omitempty"`
	IsTarget        *bool   `protobuf:"varint,5,opt,name=IsTarget" json:"IsTarget,omitempty"`
	SystemIterator  *string `protobuf:"bytes,6,opt,name=SystemIterator" json:"SystemIterator,omitempty"`
}

func (x *Measurement) Reset() {
	*x = Measurement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_internal_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Measurement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Measurement) ProtoMessage() {}

func (x *Measurement) ProtoReflect() protoreflect.Message {
	mi := &file_internal_internal_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Measurement.ProtoReflect.Descriptor instead.
func (*Measurement) Descriptor() ([]byte, []int) {
	return file_internal_internal_proto_rawDescGZIP(), []int{4}
}

func (x *Measurement) GetDatabase() string {
	if x != nil && x.Database != nil {
		return *x.Database
	}
	return ""
}

func (x *Measurement) GetRetentionPolicy() string {
	if x != nil && x.RetentionPolicy != nil {
		return *x.RetentionPolicy
	}
	return ""
}

func (x *Measurement) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Measurement) GetRegex() string {
	if x != nil && x.Regex != nil {
		return *x.Regex
	}
	return ""
}

func (x *Measurement) GetIsTarget() bool {
	if x != nil && x.IsTarget != nil {
		return *x.IsTarget
	}
	return false
}

func (x *Measurement) GetSystemIterator() string {
	if x != nil && x.SystemIterator != nil {
		return *x.SystemIterator
	}
	return ""
}

type Interval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Duration *int64 `protobuf:"varint,1,opt,name=Duration" json:"Duration,omitempty"`
	Offset   *int64 `protobuf:"varint,2,opt,name=Offset" json:"Offset,omitempty"`
}

func (x *Interval) Reset() {
	*x = Interval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_internal_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Interval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Interval) ProtoMessage() {}

func (x *Interval) ProtoReflect() protoreflect.Message {
	mi := &file_internal_internal_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Interval.ProtoReflect.Descriptor instead.
func (*Interval) Descriptor() ([]byte, []int) {
	return file_internal_internal_proto_rawDescGZIP(), []int{5}
}

func (x *Interval) GetDuration() int64 {
	if x != nil && x.Duration != nil {
		return *x.Duration
	}
	return 0
}

func (x *Interval) GetOffset() int64 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

type IteratorStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeriesN *int64 `protobuf:"varint,1,opt,name=SeriesN" json:"SeriesN,omitempty"`
	PointN  *int64 `protobuf:"varint,2,opt,name=PointN" json:"PointN,omitempty"`
}

func (x *IteratorStats) Reset() {
	*x = IteratorStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_internal_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IteratorStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IteratorStats) ProtoMessage() {}

func (x *IteratorStats) ProtoReflect() protoreflect.Message {
	mi := &file_internal_internal_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IteratorStats.ProtoReflect.Descriptor instead.
func (*IteratorStats) Descriptor() ([]byte, []int) {
	return file_internal_internal_proto_rawDescGZIP(), []int{6}
}

func (x *IteratorStats) GetSeriesN() int64 {
	if x != nil && x.SeriesN != nil {
		return *x.SeriesN
	}
	return 0
}

func (x *IteratorStats) GetPointN() int64 {
	if x != nil && x.PointN != nil {
		return *x.PointN
	}
	return 0
}

type VarRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val  *string `protobuf:"bytes,1,req,name=Val" json:"Val,omitempty"`
	Type *int32  `protobuf:"varint,2,opt,name=Type" json:"Type,omitempty"`
}

func (x *VarRef) Reset() {
	*x = VarRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_internal_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VarRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VarRef) ProtoMessage() {}

func (x *VarRef) ProtoReflect() protoreflect.Message {
	mi := &file_internal_internal_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VarRef.ProtoReflect.Descriptor instead.
func (*VarRef) Descriptor() ([]byte, []int) {
	return file_internal_internal_proto_rawDescGZIP(), []int{7}
}

func (x *VarRef) GetVal() string {
	if x != nil && x.Val != nil {
		return *x.Val
	}
	return ""
}

func (x *VarRef) GetType() int32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

var File_internal_internal_proto protoreflect.FileDescriptor

var file_internal_internal_proto_rawDesc = []byte{
	0x0a, 0x17, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x22, 0x85, 0x03, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x54, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x54, 0x61,
	0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x03,
	0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4e, 0x69, 0x6c, 0x18, 0x04, 0x20,
	0x02, 0x28, 0x08, 0x52, 0x03, 0x4e, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x03, 0x41, 0x75, 0x78, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x41, 0x75,
	0x78, 0x52, 0x03, 0x41, 0x75, 0x78, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x41, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x46, 0x6c, 0x6f, 0x61,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65,
	0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x49, 0x6e,
	0x74, 0x65, 0x67, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0c, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x24, 0x0a, 0x0d, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x49, 0x74,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x63, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x54, 0x72, 0x61, 0x63, 0x65, 0x22, 0xd1, 0x01, 0x0a, 0x03, 0x41, 0x75, 0x78,
	0x12, 0x1a, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x05, 0x52, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0a, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61,
	0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x55,
	0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x85, 0x05, 0x0a,
	0x0f, 0x49, 0x74, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x45, 0x78, 0x70, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x45, 0x78, 0x70, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x75, 0x78, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x03, 0x41, 0x75, 0x78, 0x12, 0x25, 0x0a, 0x06, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x56,
	0x61, 0x72, 0x52, 0x65, 0x66, 0x52, 0x06, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x2c, 0x0a,
	0x07, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x07, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x08, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x08,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x69, 0x6d, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x69,
	0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x42, 0x79, 0x18, 0x13, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x42, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x46, 0x69, 0x6c, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x6c, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x46, 0x69, 0x6c, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x73, 0x63, 0x65, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x41, 0x73, 0x63, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x53, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x53, 0x4f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x72, 0x69, 0x70, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x53, 0x74, 0x72, 0x69, 0x70, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x65, 0x64, 0x75, 0x70, 0x65, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x44, 0x65, 0x64, 0x75, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x61,
	0x78, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x4e, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x4d, 0x61, 0x78, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x4e, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x65, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x65, 0x64, 0x22, 0x38, 0x0a, 0x0c, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x28, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x61, 0x73,
	0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xc1,
	0x01, 0x0a, 0x0b, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x52, 0x65,
	0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x65, 0x67, 0x65,
	0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x52, 0x65, 0x67, 0x65, 0x78, 0x12, 0x1a,
	0x0a, 0x08, 0x49, 0x73, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x49, 0x73, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x49, 0x74, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x74, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x22, 0x3e, 0x0a, 0x08, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x22, 0x41, 0x0a, 0x0d, 0x49, 0x74, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x4e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x4e, 0x12, 0x16, 0x0a,
	0x06, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x4e, 0x22, 0x2e, 0x0a, 0x06, 0x56, 0x61, 0x72, 0x52, 0x65, 0x66, 0x12,
	0x10, 0x0a, 0x03, 0x56, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x03, 0x56, 0x61,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x71, 0x75, 0x65, 0x72, 0x79,
}

var (
	file_internal_internal_proto_rawDescOnce sync.Once
	file_internal_internal_proto_rawDescData = file_internal_internal_proto_rawDesc
)

func file_internal_internal_proto_rawDescGZIP() []byte {
	file_internal_internal_proto_rawDescOnce.Do(func() {
		file_internal_internal_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_internal_proto_rawDescData)
	})
	return file_internal_internal_proto_rawDescData
}

var file_internal_internal_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_internal_internal_proto_goTypes = []interface{}{
	(*Point)(nil),           // 0: query.Point
	(*Aux)(nil),             // 1: query.Aux
	(*IteratorOptions)(nil), // 2: query.IteratorOptions
	(*Measurements)(nil),    // 3: query.Measurements
	(*Measurement)(nil),     // 4: query.Measurement
	(*Interval)(nil),        // 5: query.Interval
	(*IteratorStats)(nil),   // 6: query.IteratorStats
	(*VarRef)(nil),          // 7: query.VarRef
}
var file_internal_internal_proto_depIdxs = []int32{
	1, // 0: query.Point.Aux:type_name -> query.Aux
	6, // 1: query.Point.Stats:type_name -> query.IteratorStats
	7, // 2: query.IteratorOptions.Fields:type_name -> query.VarRef
	4, // 3: query.IteratorOptions.Sources:type_name -> query.Measurement
	5, // 4: query.IteratorOptions.Interval:type_name -> query.Interval
	4, // 5: query.Measurements.Items:type_name -> query.Measurement
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_internal_internal_proto_init() }
func file_internal_internal_proto_init() {
	if File_internal_internal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_internal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
		file_internal_internal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Aux); i {
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
		file_internal_internal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IteratorOptions); i {
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
		file_internal_internal_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Measurements); i {
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
		file_internal_internal_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Measurement); i {
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
		file_internal_internal_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Interval); i {
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
		file_internal_internal_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IteratorStats); i {
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
		file_internal_internal_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VarRef); i {
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
			RawDescriptor: file_internal_internal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_internal_proto_goTypes,
		DependencyIndexes: file_internal_internal_proto_depIdxs,
		MessageInfos:      file_internal_internal_proto_msgTypes,
	}.Build()
	File_internal_internal_proto = out.File
	file_internal_internal_proto_rawDesc = nil
	file_internal_internal_proto_goTypes = nil
	file_internal_internal_proto_depIdxs = nil
}
