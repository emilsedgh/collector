// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.14.0
// source: bloat_report.proto

package pganalyze_collector

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

type BloatLookupMethod int32

const (
	BloatLookupMethod_ESTIMATE_FAST BloatLookupMethod = 0 // Estimation queries
	BloatLookupMethod_ESTIMATE_SLOW BloatLookupMethod = 1 // pgstattuple_approx
	BloatLookupMethod_FULL_SCAN     BloatLookupMethod = 2 // pgstattuple
)

// Enum value maps for BloatLookupMethod.
var (
	BloatLookupMethod_name = map[int32]string{
		0: "ESTIMATE_FAST",
		1: "ESTIMATE_SLOW",
		2: "FULL_SCAN",
	}
	BloatLookupMethod_value = map[string]int32{
		"ESTIMATE_FAST": 0,
		"ESTIMATE_SLOW": 1,
		"FULL_SCAN":     2,
	}
)

func (x BloatLookupMethod) Enum() *BloatLookupMethod {
	p := new(BloatLookupMethod)
	*p = x
	return p
}

func (x BloatLookupMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BloatLookupMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_bloat_report_proto_enumTypes[0].Descriptor()
}

func (BloatLookupMethod) Type() protoreflect.EnumType {
	return &file_bloat_report_proto_enumTypes[0]
}

func (x BloatLookupMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BloatLookupMethod.Descriptor instead.
func (BloatLookupMethod) EnumDescriptor() ([]byte, []int) {
	return file_bloat_report_proto_rawDescGZIP(), []int{0}
}

type BloatReportData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatabaseReferences      []*DatabaseReference      `protobuf:"bytes,10,rep,name=database_references,json=databaseReferences,proto3" json:"database_references,omitempty"`
	RelationReferences      []*RelationReference      `protobuf:"bytes,11,rep,name=relation_references,json=relationReferences,proto3" json:"relation_references,omitempty"`
	IndexReferences         []*IndexReference         `protobuf:"bytes,12,rep,name=index_references,json=indexReferences,proto3" json:"index_references,omitempty"`
	RelationBloatStatistics []*RelationBloatStatistic `protobuf:"bytes,20,rep,name=relation_bloat_statistics,json=relationBloatStatistics,proto3" json:"relation_bloat_statistics,omitempty"`
	IndexBloatStatistics    []*IndexBloatStatistic    `protobuf:"bytes,21,rep,name=index_bloat_statistics,json=indexBloatStatistics,proto3" json:"index_bloat_statistics,omitempty"`
}

func (x *BloatReportData) Reset() {
	*x = BloatReportData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bloat_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BloatReportData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BloatReportData) ProtoMessage() {}

func (x *BloatReportData) ProtoReflect() protoreflect.Message {
	mi := &file_bloat_report_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BloatReportData.ProtoReflect.Descriptor instead.
func (*BloatReportData) Descriptor() ([]byte, []int) {
	return file_bloat_report_proto_rawDescGZIP(), []int{0}
}

func (x *BloatReportData) GetDatabaseReferences() []*DatabaseReference {
	if x != nil {
		return x.DatabaseReferences
	}
	return nil
}

func (x *BloatReportData) GetRelationReferences() []*RelationReference {
	if x != nil {
		return x.RelationReferences
	}
	return nil
}

func (x *BloatReportData) GetIndexReferences() []*IndexReference {
	if x != nil {
		return x.IndexReferences
	}
	return nil
}

func (x *BloatReportData) GetRelationBloatStatistics() []*RelationBloatStatistic {
	if x != nil {
		return x.RelationBloatStatistics
	}
	return nil
}

func (x *BloatReportData) GetIndexBloatStatistics() []*IndexBloatStatistic {
	if x != nil {
		return x.IndexBloatStatistics
	}
	return nil
}

type RelationBloatStatistic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RelationIdx       int32             `protobuf:"varint,1,opt,name=relation_idx,json=relationIdx,proto3" json:"relation_idx,omitempty"`
	BloatLookupMethod BloatLookupMethod `protobuf:"varint,2,opt,name=bloat_lookup_method,json=bloatLookupMethod,proto3,enum=pganalyze.collector.BloatLookupMethod" json:"bloat_lookup_method,omitempty"`
	TotalBytes        int64             `protobuf:"varint,3,opt,name=total_bytes,json=totalBytes,proto3" json:"total_bytes,omitempty"`               // Total bytes of the relation, including bloat
	BloatBytes        int64             `protobuf:"varint,4,opt,name=bloat_bytes,json=bloatBytes,proto3" json:"bloat_bytes,omitempty"`               // Wasted or "Free" bytes which can only be reclaimed by a VACUUM FULL
	LiveTupleBytes    int64             `protobuf:"varint,5,opt,name=live_tuple_bytes,json=liveTupleBytes,proto3" json:"live_tuple_bytes,omitempty"` // Optional, "tuple_len" from pgstattuple
	LiveTupleCount    int64             `protobuf:"varint,6,opt,name=live_tuple_count,json=liveTupleCount,proto3" json:"live_tuple_count,omitempty"` // Optional, "tuple_count" from pgstattuple
	DeadTupleBytes    int64             `protobuf:"varint,7,opt,name=dead_tuple_bytes,json=deadTupleBytes,proto3" json:"dead_tuple_bytes,omitempty"` // Optional, "dead_tuple_len" from pgstattuple
	DeadTupleCount    int64             `protobuf:"varint,8,opt,name=dead_tuple_count,json=deadTupleCount,proto3" json:"dead_tuple_count,omitempty"` // Optional, "dead_tuple_count" from pgstattuple
}

func (x *RelationBloatStatistic) Reset() {
	*x = RelationBloatStatistic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bloat_report_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationBloatStatistic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationBloatStatistic) ProtoMessage() {}

func (x *RelationBloatStatistic) ProtoReflect() protoreflect.Message {
	mi := &file_bloat_report_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationBloatStatistic.ProtoReflect.Descriptor instead.
func (*RelationBloatStatistic) Descriptor() ([]byte, []int) {
	return file_bloat_report_proto_rawDescGZIP(), []int{1}
}

func (x *RelationBloatStatistic) GetRelationIdx() int32 {
	if x != nil {
		return x.RelationIdx
	}
	return 0
}

func (x *RelationBloatStatistic) GetBloatLookupMethod() BloatLookupMethod {
	if x != nil {
		return x.BloatLookupMethod
	}
	return BloatLookupMethod_ESTIMATE_FAST
}

func (x *RelationBloatStatistic) GetTotalBytes() int64 {
	if x != nil {
		return x.TotalBytes
	}
	return 0
}

func (x *RelationBloatStatistic) GetBloatBytes() int64 {
	if x != nil {
		return x.BloatBytes
	}
	return 0
}

func (x *RelationBloatStatistic) GetLiveTupleBytes() int64 {
	if x != nil {
		return x.LiveTupleBytes
	}
	return 0
}

func (x *RelationBloatStatistic) GetLiveTupleCount() int64 {
	if x != nil {
		return x.LiveTupleCount
	}
	return 0
}

func (x *RelationBloatStatistic) GetDeadTupleBytes() int64 {
	if x != nil {
		return x.DeadTupleBytes
	}
	return 0
}

func (x *RelationBloatStatistic) GetDeadTupleCount() int64 {
	if x != nil {
		return x.DeadTupleCount
	}
	return 0
}

type IndexBloatStatistic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IndexIdx          int32             `protobuf:"varint,1,opt,name=index_idx,json=indexIdx,proto3" json:"index_idx,omitempty"`
	BloatLookupMethod BloatLookupMethod `protobuf:"varint,2,opt,name=bloat_lookup_method,json=bloatLookupMethod,proto3,enum=pganalyze.collector.BloatLookupMethod" json:"bloat_lookup_method,omitempty"`
	TotalBytes        int64             `protobuf:"varint,3,opt,name=total_bytes,json=totalBytes,proto3" json:"total_bytes,omitempty"`                // Total bytes of the index, including bloat
	BloatBytes        int64             `protobuf:"varint,4,opt,name=bloat_bytes,json=bloatBytes,proto3" json:"bloat_bytes,omitempty"`                // Wasted or "Free" bytes which can only be reclaimed by a REINDEX
	InternalPages     int64             `protobuf:"varint,5,opt,name=internal_pages,json=internalPages,proto3" json:"internal_pages,omitempty"`       // Optional, "internal_pages" from pgstatindex
	LeafPages         int64             `protobuf:"varint,6,opt,name=leaf_pages,json=leafPages,proto3" json:"leaf_pages,omitempty"`                   // Optional, "leaf_pages" from pgstatindex
	EmptyPages        int64             `protobuf:"varint,7,opt,name=empty_pages,json=emptyPages,proto3" json:"empty_pages,omitempty"`                // Optional, "empty_pages" from pgstatindex
	DeletedPages      int64             `protobuf:"varint,8,opt,name=deleted_pages,json=deletedPages,proto3" json:"deleted_pages,omitempty"`          // Optional, "deleted_pages" from pgstatindex
	AvgLeafDensity    float64           `protobuf:"fixed64,9,opt,name=avg_leaf_density,json=avgLeafDensity,proto3" json:"avg_leaf_density,omitempty"` // Optional, "avg_leaf_density" from pgstatindex
}

func (x *IndexBloatStatistic) Reset() {
	*x = IndexBloatStatistic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bloat_report_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexBloatStatistic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexBloatStatistic) ProtoMessage() {}

func (x *IndexBloatStatistic) ProtoReflect() protoreflect.Message {
	mi := &file_bloat_report_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexBloatStatistic.ProtoReflect.Descriptor instead.
func (*IndexBloatStatistic) Descriptor() ([]byte, []int) {
	return file_bloat_report_proto_rawDescGZIP(), []int{2}
}

func (x *IndexBloatStatistic) GetIndexIdx() int32 {
	if x != nil {
		return x.IndexIdx
	}
	return 0
}

func (x *IndexBloatStatistic) GetBloatLookupMethod() BloatLookupMethod {
	if x != nil {
		return x.BloatLookupMethod
	}
	return BloatLookupMethod_ESTIMATE_FAST
}

func (x *IndexBloatStatistic) GetTotalBytes() int64 {
	if x != nil {
		return x.TotalBytes
	}
	return 0
}

func (x *IndexBloatStatistic) GetBloatBytes() int64 {
	if x != nil {
		return x.BloatBytes
	}
	return 0
}

func (x *IndexBloatStatistic) GetInternalPages() int64 {
	if x != nil {
		return x.InternalPages
	}
	return 0
}

func (x *IndexBloatStatistic) GetLeafPages() int64 {
	if x != nil {
		return x.LeafPages
	}
	return 0
}

func (x *IndexBloatStatistic) GetEmptyPages() int64 {
	if x != nil {
		return x.EmptyPages
	}
	return 0
}

func (x *IndexBloatStatistic) GetDeletedPages() int64 {
	if x != nil {
		return x.DeletedPages
	}
	return 0
}

func (x *IndexBloatStatistic) GetAvgLeafDensity() float64 {
	if x != nil {
		return x.AvgLeafDensity
	}
	return 0
}

var File_bloat_report_proto protoreflect.FileDescriptor

var file_bloat_report_proto_rawDesc = []byte{
	0x0a, 0x12, 0x62, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x70, 0x67, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x2e,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x1a, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x03, 0x0a, 0x0f, 0x42, 0x6c, 0x6f, 0x61,
	0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x57, 0x0a, 0x13, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x67, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x7a, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x12, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x73, 0x12, 0x57, 0x0a, 0x13, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x70, 0x67, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x2e, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x12, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x4e, 0x0a,
	0x10, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x67, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x7a, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x67, 0x0a,
	0x19, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x6c, 0x6f, 0x61, 0x74, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x70, 0x67, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x2e, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x6c, 0x6f, 0x61, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x52, 0x17, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x61, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x5e, 0x0a, 0x16, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f,
	0x62, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x18, 0x15, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x67, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x7a, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x42, 0x6c, 0x6f, 0x61, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x52, 0x14, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x6c, 0x6f, 0x61, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x22, 0xfd, 0x02, 0x0a, 0x16, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x61, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x78, 0x12, 0x56, 0x0a, 0x13, 0x62, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x6c, 0x6f,
	0x6f, 0x6b, 0x75, 0x70, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x26, 0x2e, 0x70, 0x67, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x2e, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x42, 0x6c, 0x6f, 0x61, 0x74, 0x4c, 0x6f, 0x6f,
	0x6b, 0x75, 0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x11, 0x62, 0x6c, 0x6f, 0x61, 0x74,
	0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x62, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x62, 0x6c, 0x6f, 0x61, 0x74, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x28,
	0x0a, 0x10, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x5f, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6c, 0x69, 0x76, 0x65, 0x54, 0x75,
	0x70, 0x6c, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x69, 0x76, 0x65,
	0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0e, 0x6c, 0x69, 0x76, 0x65, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x64, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65,
	0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x64, 0x65,
	0x61, 0x64, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x10,
	0x64, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x64, 0x65, 0x61, 0x64, 0x54, 0x75, 0x70, 0x6c,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x82, 0x03, 0x0a, 0x13, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x42, 0x6c, 0x6f, 0x61, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x69, 0x64, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x49, 0x64, 0x78, 0x12, 0x56, 0x0a, 0x13, 0x62,
	0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x5f, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x70, 0x67, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x7a, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x42,
	0x6c, 0x6f, 0x61, 0x74, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x52, 0x11, 0x62, 0x6c, 0x6f, 0x61, 0x74, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x62, 0x6c, 0x6f, 0x61, 0x74,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x6c, 0x65, 0x61, 0x66, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x6c, 0x65, 0x61, 0x66, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x50, 0x61, 0x67, 0x65,
	0x73, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x76, 0x67, 0x5f, 0x6c, 0x65, 0x61, 0x66, 0x5f, 0x64, 0x65,
	0x6e, 0x73, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x61, 0x76, 0x67,
	0x4c, 0x65, 0x61, 0x66, 0x44, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x2a, 0x48, 0x0a, 0x11, 0x42,
	0x6c, 0x6f, 0x61, 0x74, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x12, 0x11, 0x0a, 0x0d, 0x45, 0x53, 0x54, 0x49, 0x4d, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x53,
	0x54, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x53, 0x54, 0x49, 0x4d, 0x41, 0x54, 0x45, 0x5f,
	0x53, 0x4c, 0x4f, 0x57, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x55, 0x4c, 0x4c, 0x5f, 0x53,
	0x43, 0x41, 0x4e, 0x10, 0x02, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x67, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x2f, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f, 0x70,
	0x67, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bloat_report_proto_rawDescOnce sync.Once
	file_bloat_report_proto_rawDescData = file_bloat_report_proto_rawDesc
)

func file_bloat_report_proto_rawDescGZIP() []byte {
	file_bloat_report_proto_rawDescOnce.Do(func() {
		file_bloat_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_bloat_report_proto_rawDescData)
	})
	return file_bloat_report_proto_rawDescData
}

var file_bloat_report_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_bloat_report_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_bloat_report_proto_goTypes = []interface{}{
	(BloatLookupMethod)(0),         // 0: pganalyze.collector.BloatLookupMethod
	(*BloatReportData)(nil),        // 1: pganalyze.collector.BloatReportData
	(*RelationBloatStatistic)(nil), // 2: pganalyze.collector.RelationBloatStatistic
	(*IndexBloatStatistic)(nil),    // 3: pganalyze.collector.IndexBloatStatistic
	(*DatabaseReference)(nil),      // 4: pganalyze.collector.DatabaseReference
	(*RelationReference)(nil),      // 5: pganalyze.collector.RelationReference
	(*IndexReference)(nil),         // 6: pganalyze.collector.IndexReference
}
var file_bloat_report_proto_depIdxs = []int32{
	4, // 0: pganalyze.collector.BloatReportData.database_references:type_name -> pganalyze.collector.DatabaseReference
	5, // 1: pganalyze.collector.BloatReportData.relation_references:type_name -> pganalyze.collector.RelationReference
	6, // 2: pganalyze.collector.BloatReportData.index_references:type_name -> pganalyze.collector.IndexReference
	2, // 3: pganalyze.collector.BloatReportData.relation_bloat_statistics:type_name -> pganalyze.collector.RelationBloatStatistic
	3, // 4: pganalyze.collector.BloatReportData.index_bloat_statistics:type_name -> pganalyze.collector.IndexBloatStatistic
	0, // 5: pganalyze.collector.RelationBloatStatistic.bloat_lookup_method:type_name -> pganalyze.collector.BloatLookupMethod
	0, // 6: pganalyze.collector.IndexBloatStatistic.bloat_lookup_method:type_name -> pganalyze.collector.BloatLookupMethod
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_bloat_report_proto_init() }
func file_bloat_report_proto_init() {
	if File_bloat_report_proto != nil {
		return
	}
	file_shared_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_bloat_report_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BloatReportData); i {
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
		file_bloat_report_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationBloatStatistic); i {
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
		file_bloat_report_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexBloatStatistic); i {
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
			RawDescriptor: file_bloat_report_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bloat_report_proto_goTypes,
		DependencyIndexes: file_bloat_report_proto_depIdxs,
		EnumInfos:         file_bloat_report_proto_enumTypes,
		MessageInfos:      file_bloat_report_proto_msgTypes,
	}.Build()
	File_bloat_report_proto = out.File
	file_bloat_report_proto_rawDesc = nil
	file_bloat_report_proto_goTypes = nil
	file_bloat_report_proto_depIdxs = nil
}
