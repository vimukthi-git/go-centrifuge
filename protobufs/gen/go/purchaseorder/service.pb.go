// Code generated by protoc-gen-go. DO NOT EDIT.
// source: purchaseorder/service.proto

package purchaseorderpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetRequest struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_393a7cd179a9d71e, []int{0}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

type GetVersionRequest struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVersionRequest) Reset()         { *m = GetVersionRequest{} }
func (m *GetVersionRequest) String() string { return proto.CompactTextString(m) }
func (*GetVersionRequest) ProtoMessage()    {}
func (*GetVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_393a7cd179a9d71e, []int{1}
}
func (m *GetVersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVersionRequest.Unmarshal(m, b)
}
func (m *GetVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVersionRequest.Marshal(b, m, deterministic)
}
func (dst *GetVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVersionRequest.Merge(dst, src)
}
func (m *GetVersionRequest) XXX_Size() int {
	return xxx_messageInfo_GetVersionRequest.Size(m)
}
func (m *GetVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetVersionRequest proto.InternalMessageInfo

func (m *GetVersionRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *GetVersionRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type PurchaseOrderCreatePayload struct {
	Collaborators        []string           `protobuf:"bytes,1,rep,name=collaborators,proto3" json:"collaborators,omitempty"`
	Data                 *PurchaseOrderData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PurchaseOrderCreatePayload) Reset()         { *m = PurchaseOrderCreatePayload{} }
func (m *PurchaseOrderCreatePayload) String() string { return proto.CompactTextString(m) }
func (*PurchaseOrderCreatePayload) ProtoMessage()    {}
func (*PurchaseOrderCreatePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_393a7cd179a9d71e, []int{2}
}
func (m *PurchaseOrderCreatePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurchaseOrderCreatePayload.Unmarshal(m, b)
}
func (m *PurchaseOrderCreatePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurchaseOrderCreatePayload.Marshal(b, m, deterministic)
}
func (dst *PurchaseOrderCreatePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurchaseOrderCreatePayload.Merge(dst, src)
}
func (m *PurchaseOrderCreatePayload) XXX_Size() int {
	return xxx_messageInfo_PurchaseOrderCreatePayload.Size(m)
}
func (m *PurchaseOrderCreatePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_PurchaseOrderCreatePayload.DiscardUnknown(m)
}

var xxx_messageInfo_PurchaseOrderCreatePayload proto.InternalMessageInfo

func (m *PurchaseOrderCreatePayload) GetCollaborators() []string {
	if m != nil {
		return m.Collaborators
	}
	return nil
}

func (m *PurchaseOrderCreatePayload) GetData() *PurchaseOrderData {
	if m != nil {
		return m.Data
	}
	return nil
}

type PurchaseOrderUpdatePayload struct {
	Identifier           string             `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Collaborators        []string           `protobuf:"bytes,2,rep,name=collaborators,proto3" json:"collaborators,omitempty"`
	Data                 *PurchaseOrderData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PurchaseOrderUpdatePayload) Reset()         { *m = PurchaseOrderUpdatePayload{} }
func (m *PurchaseOrderUpdatePayload) String() string { return proto.CompactTextString(m) }
func (*PurchaseOrderUpdatePayload) ProtoMessage()    {}
func (*PurchaseOrderUpdatePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_393a7cd179a9d71e, []int{3}
}
func (m *PurchaseOrderUpdatePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurchaseOrderUpdatePayload.Unmarshal(m, b)
}
func (m *PurchaseOrderUpdatePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurchaseOrderUpdatePayload.Marshal(b, m, deterministic)
}
func (dst *PurchaseOrderUpdatePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurchaseOrderUpdatePayload.Merge(dst, src)
}
func (m *PurchaseOrderUpdatePayload) XXX_Size() int {
	return xxx_messageInfo_PurchaseOrderUpdatePayload.Size(m)
}
func (m *PurchaseOrderUpdatePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_PurchaseOrderUpdatePayload.DiscardUnknown(m)
}

var xxx_messageInfo_PurchaseOrderUpdatePayload proto.InternalMessageInfo

func (m *PurchaseOrderUpdatePayload) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *PurchaseOrderUpdatePayload) GetCollaborators() []string {
	if m != nil {
		return m.Collaborators
	}
	return nil
}

func (m *PurchaseOrderUpdatePayload) GetData() *PurchaseOrderData {
	if m != nil {
		return m.Data
	}
	return nil
}

type PurchaseOrderResponse struct {
	Header               *ResponseHeader    `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data                 *PurchaseOrderData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PurchaseOrderResponse) Reset()         { *m = PurchaseOrderResponse{} }
func (m *PurchaseOrderResponse) String() string { return proto.CompactTextString(m) }
func (*PurchaseOrderResponse) ProtoMessage()    {}
func (*PurchaseOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_393a7cd179a9d71e, []int{4}
}
func (m *PurchaseOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurchaseOrderResponse.Unmarshal(m, b)
}
func (m *PurchaseOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurchaseOrderResponse.Marshal(b, m, deterministic)
}
func (dst *PurchaseOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurchaseOrderResponse.Merge(dst, src)
}
func (m *PurchaseOrderResponse) XXX_Size() int {
	return xxx_messageInfo_PurchaseOrderResponse.Size(m)
}
func (m *PurchaseOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PurchaseOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PurchaseOrderResponse proto.InternalMessageInfo

func (m *PurchaseOrderResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PurchaseOrderResponse) GetData() *PurchaseOrderData {
	if m != nil {
		return m.Data
	}
	return nil
}

// ResponseHeader contains a set of common fields for most documents
type ResponseHeader struct {
	DocumentId           string   `protobuf:"bytes,1,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
	VersionId            string   `protobuf:"bytes,2,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	State                string   `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Collaborators        []string `protobuf:"bytes,4,rep,name=collaborators,proto3" json:"collaborators,omitempty"`
	TransactionId        string   `protobuf:"bytes,5,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseHeader) Reset()         { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()    {}
func (*ResponseHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_393a7cd179a9d71e, []int{5}
}
func (m *ResponseHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseHeader.Unmarshal(m, b)
}
func (m *ResponseHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseHeader.Marshal(b, m, deterministic)
}
func (dst *ResponseHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseHeader.Merge(dst, src)
}
func (m *ResponseHeader) XXX_Size() int {
	return xxx_messageInfo_ResponseHeader.Size(m)
}
func (m *ResponseHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseHeader proto.InternalMessageInfo

func (m *ResponseHeader) GetDocumentId() string {
	if m != nil {
		return m.DocumentId
	}
	return ""
}

func (m *ResponseHeader) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

func (m *ResponseHeader) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *ResponseHeader) GetCollaborators() []string {
	if m != nil {
		return m.Collaborators
	}
	return nil
}

func (m *ResponseHeader) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

type PurchaseOrderData struct {
	PoStatus string `protobuf:"bytes,24,opt,name=po_status,json=poStatus,proto3" json:"po_status,omitempty"`
	// purchase order number or reference number
	PoNumber string `protobuf:"bytes,1,opt,name=po_number,json=poNumber,proto3" json:"po_number,omitempty"`
	// name of the ordering company
	OrderName string `protobuf:"bytes,2,opt,name=order_name,json=orderName,proto3" json:"order_name,omitempty"`
	// street and address details of the ordering company
	OrderStreet  string `protobuf:"bytes,3,opt,name=order_street,json=orderStreet,proto3" json:"order_street,omitempty"`
	OrderCity    string `protobuf:"bytes,4,opt,name=order_city,json=orderCity,proto3" json:"order_city,omitempty"`
	OrderZipcode string `protobuf:"bytes,5,opt,name=order_zipcode,json=orderZipcode,proto3" json:"order_zipcode,omitempty"`
	// country ISO code of the ordering company of this purchase order
	OrderCountry string `protobuf:"bytes,6,opt,name=order_country,json=orderCountry,proto3" json:"order_country,omitempty"`
	// name of the recipient company
	RecipientName    string `protobuf:"bytes,7,opt,name=recipient_name,json=recipientName,proto3" json:"recipient_name,omitempty"`
	RecipientStreet  string `protobuf:"bytes,8,opt,name=recipient_street,json=recipientStreet,proto3" json:"recipient_street,omitempty"`
	RecipientCity    string `protobuf:"bytes,9,opt,name=recipient_city,json=recipientCity,proto3" json:"recipient_city,omitempty"`
	RecipientZipcode string `protobuf:"bytes,10,opt,name=recipient_zipcode,json=recipientZipcode,proto3" json:"recipient_zipcode,omitempty"`
	// country ISO code of the receipient of this purchase order
	RecipientCountry string `protobuf:"bytes,11,opt,name=recipient_country,json=recipientCountry,proto3" json:"recipient_country,omitempty"`
	// ISO currency code
	Currency string `protobuf:"bytes,12,opt,name=currency,proto3" json:"currency,omitempty"`
	// ordering gross amount including tax
	OrderAmount int64 `protobuf:"varint,13,opt,name=order_amount,json=orderAmount,proto3" json:"order_amount,omitempty"`
	// invoice amount excluding tax
	NetAmount int64  `protobuf:"varint,14,opt,name=net_amount,json=netAmount,proto3" json:"net_amount,omitempty"`
	TaxAmount int64  `protobuf:"varint,15,opt,name=tax_amount,json=taxAmount,proto3" json:"tax_amount,omitempty"`
	TaxRate   int64  `protobuf:"varint,16,opt,name=tax_rate,json=taxRate,proto3" json:"tax_rate,omitempty"`
	Recipient string `protobuf:"bytes,17,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Order     string `protobuf:"bytes,18,opt,name=order,proto3" json:"order,omitempty"`
	// contact or requester or purchaser at the ordering company
	OrderContact string `protobuf:"bytes,19,opt,name=order_contact,json=orderContact,proto3" json:"order_contact,omitempty"`
	Comment      string `protobuf:"bytes,20,opt,name=comment,proto3" json:"comment,omitempty"`
	// requested delivery date
	DeliveryDate *timestamp.Timestamp `protobuf:"bytes,21,opt,name=delivery_date,json=deliveryDate,proto3" json:"delivery_date,omitempty"`
	// purchase order date
	DateCreated          *timestamp.Timestamp `protobuf:"bytes,22,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
	ExtraData            string               `protobuf:"bytes,23,opt,name=extra_data,json=extraData,proto3" json:"extra_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PurchaseOrderData) Reset()         { *m = PurchaseOrderData{} }
func (m *PurchaseOrderData) String() string { return proto.CompactTextString(m) }
func (*PurchaseOrderData) ProtoMessage()    {}
func (*PurchaseOrderData) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_393a7cd179a9d71e, []int{6}
}
func (m *PurchaseOrderData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurchaseOrderData.Unmarshal(m, b)
}
func (m *PurchaseOrderData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurchaseOrderData.Marshal(b, m, deterministic)
}
func (dst *PurchaseOrderData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurchaseOrderData.Merge(dst, src)
}
func (m *PurchaseOrderData) XXX_Size() int {
	return xxx_messageInfo_PurchaseOrderData.Size(m)
}
func (m *PurchaseOrderData) XXX_DiscardUnknown() {
	xxx_messageInfo_PurchaseOrderData.DiscardUnknown(m)
}

var xxx_messageInfo_PurchaseOrderData proto.InternalMessageInfo

func (m *PurchaseOrderData) GetPoStatus() string {
	if m != nil {
		return m.PoStatus
	}
	return ""
}

func (m *PurchaseOrderData) GetPoNumber() string {
	if m != nil {
		return m.PoNumber
	}
	return ""
}

func (m *PurchaseOrderData) GetOrderName() string {
	if m != nil {
		return m.OrderName
	}
	return ""
}

func (m *PurchaseOrderData) GetOrderStreet() string {
	if m != nil {
		return m.OrderStreet
	}
	return ""
}

func (m *PurchaseOrderData) GetOrderCity() string {
	if m != nil {
		return m.OrderCity
	}
	return ""
}

func (m *PurchaseOrderData) GetOrderZipcode() string {
	if m != nil {
		return m.OrderZipcode
	}
	return ""
}

func (m *PurchaseOrderData) GetOrderCountry() string {
	if m != nil {
		return m.OrderCountry
	}
	return ""
}

func (m *PurchaseOrderData) GetRecipientName() string {
	if m != nil {
		return m.RecipientName
	}
	return ""
}

func (m *PurchaseOrderData) GetRecipientStreet() string {
	if m != nil {
		return m.RecipientStreet
	}
	return ""
}

func (m *PurchaseOrderData) GetRecipientCity() string {
	if m != nil {
		return m.RecipientCity
	}
	return ""
}

func (m *PurchaseOrderData) GetRecipientZipcode() string {
	if m != nil {
		return m.RecipientZipcode
	}
	return ""
}

func (m *PurchaseOrderData) GetRecipientCountry() string {
	if m != nil {
		return m.RecipientCountry
	}
	return ""
}

func (m *PurchaseOrderData) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *PurchaseOrderData) GetOrderAmount() int64 {
	if m != nil {
		return m.OrderAmount
	}
	return 0
}

func (m *PurchaseOrderData) GetNetAmount() int64 {
	if m != nil {
		return m.NetAmount
	}
	return 0
}

func (m *PurchaseOrderData) GetTaxAmount() int64 {
	if m != nil {
		return m.TaxAmount
	}
	return 0
}

func (m *PurchaseOrderData) GetTaxRate() int64 {
	if m != nil {
		return m.TaxRate
	}
	return 0
}

func (m *PurchaseOrderData) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *PurchaseOrderData) GetOrder() string {
	if m != nil {
		return m.Order
	}
	return ""
}

func (m *PurchaseOrderData) GetOrderContact() string {
	if m != nil {
		return m.OrderContact
	}
	return ""
}

func (m *PurchaseOrderData) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *PurchaseOrderData) GetDeliveryDate() *timestamp.Timestamp {
	if m != nil {
		return m.DeliveryDate
	}
	return nil
}

func (m *PurchaseOrderData) GetDateCreated() *timestamp.Timestamp {
	if m != nil {
		return m.DateCreated
	}
	return nil
}

func (m *PurchaseOrderData) GetExtraData() string {
	if m != nil {
		return m.ExtraData
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "purchaseorder.GetRequest")
	proto.RegisterType((*GetVersionRequest)(nil), "purchaseorder.GetVersionRequest")
	proto.RegisterType((*PurchaseOrderCreatePayload)(nil), "purchaseorder.PurchaseOrderCreatePayload")
	proto.RegisterType((*PurchaseOrderUpdatePayload)(nil), "purchaseorder.PurchaseOrderUpdatePayload")
	proto.RegisterType((*PurchaseOrderResponse)(nil), "purchaseorder.PurchaseOrderResponse")
	proto.RegisterType((*ResponseHeader)(nil), "purchaseorder.ResponseHeader")
	proto.RegisterType((*PurchaseOrderData)(nil), "purchaseorder.PurchaseOrderData")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DocumentServiceClient is the client API for DocumentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DocumentServiceClient interface {
	Create(ctx context.Context, in *PurchaseOrderCreatePayload, opts ...grpc.CallOption) (*PurchaseOrderResponse, error)
	Update(ctx context.Context, in *PurchaseOrderUpdatePayload, opts ...grpc.CallOption) (*PurchaseOrderResponse, error)
	GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*PurchaseOrderResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*PurchaseOrderResponse, error)
}

type documentServiceClient struct {
	cc *grpc.ClientConn
}

func NewDocumentServiceClient(cc *grpc.ClientConn) DocumentServiceClient {
	return &documentServiceClient{cc}
}

func (c *documentServiceClient) Create(ctx context.Context, in *PurchaseOrderCreatePayload, opts ...grpc.CallOption) (*PurchaseOrderResponse, error) {
	out := new(PurchaseOrderResponse)
	err := c.cc.Invoke(ctx, "/purchaseorder.DocumentService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) Update(ctx context.Context, in *PurchaseOrderUpdatePayload, opts ...grpc.CallOption) (*PurchaseOrderResponse, error) {
	out := new(PurchaseOrderResponse)
	err := c.cc.Invoke(ctx, "/purchaseorder.DocumentService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*PurchaseOrderResponse, error) {
	out := new(PurchaseOrderResponse)
	err := c.cc.Invoke(ctx, "/purchaseorder.DocumentService/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*PurchaseOrderResponse, error) {
	out := new(PurchaseOrderResponse)
	err := c.cc.Invoke(ctx, "/purchaseorder.DocumentService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DocumentServiceServer is the server API for DocumentService service.
type DocumentServiceServer interface {
	Create(context.Context, *PurchaseOrderCreatePayload) (*PurchaseOrderResponse, error)
	Update(context.Context, *PurchaseOrderUpdatePayload) (*PurchaseOrderResponse, error)
	GetVersion(context.Context, *GetVersionRequest) (*PurchaseOrderResponse, error)
	Get(context.Context, *GetRequest) (*PurchaseOrderResponse, error)
}

func RegisterDocumentServiceServer(s *grpc.Server, srv DocumentServiceServer) {
	s.RegisterService(&_DocumentService_serviceDesc, srv)
}

func _DocumentService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseOrderCreatePayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/purchaseorder.DocumentService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).Create(ctx, req.(*PurchaseOrderCreatePayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseOrderUpdatePayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/purchaseorder.DocumentService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).Update(ctx, req.(*PurchaseOrderUpdatePayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/purchaseorder.DocumentService/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).GetVersion(ctx, req.(*GetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/purchaseorder.DocumentService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DocumentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "purchaseorder.DocumentService",
	HandlerType: (*DocumentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _DocumentService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DocumentService_Update_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _DocumentService_GetVersion_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _DocumentService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "purchaseorder/service.proto",
}

func init() {
	proto.RegisterFile("purchaseorder/service.proto", fileDescriptor_service_393a7cd179a9d71e)
}

var fileDescriptor_service_393a7cd179a9d71e = []byte{
	// 963 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x5f, 0x6f, 0x1b, 0x45,
	0x10, 0xd7, 0xe5, 0x8f, 0x63, 0xaf, 0xed, 0xa4, 0xde, 0xb6, 0x70, 0xbd, 0x34, 0xf4, 0x30, 0xad,
	0x48, 0xdb, 0xc4, 0x96, 0x42, 0xe1, 0x01, 0x09, 0xa1, 0x34, 0x91, 0x42, 0x1e, 0x28, 0xd1, 0x05,
	0x78, 0xa8, 0x90, 0xac, 0xf5, 0xde, 0xc4, 0x39, 0xc9, 0x77, 0x7b, 0xec, 0xad, 0x83, 0x4d, 0xd5,
	0x17, 0xc4, 0x17, 0x20, 0xbc, 0x80, 0x90, 0xf8, 0x10, 0xf0, 0x51, 0xf8, 0x0a, 0x7c, 0x05, 0xde,
	0xd1, 0xce, 0xee, 0xd9, 0x3e, 0xbb, 0x38, 0x09, 0x4f, 0xa7, 0xfd, 0xcd, 0x6f, 0xe6, 0x7e, 0x33,
	0xb3, 0x3b, 0x43, 0x36, 0xd3, 0x81, 0xe4, 0xe7, 0x2c, 0x03, 0x21, 0x43, 0x90, 0xed, 0x0c, 0xe4,
	0x45, 0xc4, 0xa1, 0x95, 0x4a, 0xa1, 0x04, 0xad, 0x17, 0x8c, 0xde, 0xfd, 0x9e, 0x10, 0xbd, 0x3e,
	0xb4, 0x59, 0x1a, 0xb5, 0x59, 0x92, 0x08, 0xc5, 0x54, 0x24, 0x92, 0xcc, 0x90, 0xbd, 0x07, 0xd6,
	0x8a, 0xa7, 0xee, 0xe0, 0xac, 0xad, 0xa2, 0x18, 0x32, 0xc5, 0xe2, 0xd4, 0x12, 0x76, 0xf0, 0xc3,
	0x77, 0x7b, 0x90, 0xec, 0x66, 0xdf, 0xb1, 0x5e, 0x0f, 0x64, 0x5b, 0xa4, 0x18, 0x62, 0x3e, 0x5c,
	0x73, 0x87, 0x90, 0x23, 0x50, 0x01, 0x7c, 0x3b, 0x80, 0x4c, 0xd1, 0x77, 0x08, 0x89, 0x42, 0x48,
	0x54, 0x74, 0x16, 0x81, 0x74, 0x1d, 0xdf, 0xd9, 0xae, 0x04, 0x53, 0x48, 0xf3, 0x73, 0xd2, 0x38,
	0x02, 0xf5, 0x35, 0xc8, 0x2c, 0x12, 0xc9, 0x35, 0x9d, 0xa8, 0x4b, 0xd6, 0x2e, 0x8c, 0x87, 0xbb,
	0x84, 0xc6, 0xfc, 0xd8, 0x1c, 0x12, 0xef, 0xc4, 0xa6, 0xfe, 0x85, 0x4e, 0xfd, 0x40, 0x02, 0x53,
	0x70, 0xc2, 0x46, 0x7d, 0xc1, 0x42, 0xfa, 0x90, 0xd4, 0xb9, 0xe8, 0xf7, 0x59, 0x57, 0x48, 0xa6,
	0x84, 0xcc, 0x5c, 0xc7, 0x5f, 0xde, 0xae, 0x04, 0x45, 0x90, 0x3e, 0x23, 0x2b, 0x21, 0x53, 0x0c,
	0x43, 0x57, 0xf7, 0xfc, 0x56, 0xa1, 0x96, 0xad, 0x42, 0xf8, 0x43, 0xa6, 0x58, 0x80, 0xec, 0xe6,
	0x2f, 0xce, 0xcc, 0xaf, 0xbf, 0x4a, 0xc3, 0xa9, 0x5f, 0x5f, 0x95, 0xd2, 0x9c, 0xb4, 0xa5, 0x45,
	0xd2, 0x96, 0x6f, 0x24, 0xed, 0x47, 0x87, 0xdc, 0x2d, 0xd8, 0x02, 0xc8, 0x52, 0x91, 0x64, 0x40,
	0x3f, 0x24, 0xa5, 0x73, 0x60, 0xa1, 0x55, 0x54, 0xdd, 0xdb, 0x9a, 0x89, 0x98, 0x13, 0x3f, 0x43,
	0x52, 0x60, 0xc9, 0xff, 0xb3, 0x42, 0x7f, 0x3a, 0x64, 0xbd, 0x18, 0x90, 0x3e, 0x20, 0xd5, 0x50,
	0xf0, 0x41, 0x0c, 0x89, 0xea, 0x44, 0x61, 0x5e, 0x96, 0x1c, 0x3a, 0x0e, 0xe9, 0x16, 0x21, 0xb6,
	0xb5, 0xda, 0x6e, 0x9a, 0x5d, 0xb1, 0xc8, 0x71, 0x48, 0xef, 0x90, 0xd5, 0x4c, 0x31, 0x05, 0x58,
	0x90, 0x4a, 0x60, 0x0e, 0xf3, 0xb5, 0x5c, 0x79, 0x53, 0x2d, 0x1f, 0x91, 0x75, 0x25, 0x59, 0x92,
	0x31, 0xae, 0x6c, 0xf8, 0x55, 0x0c, 0x52, 0x9f, 0x42, 0x8f, 0xc3, 0xe6, 0x3f, 0x25, 0xd2, 0x98,
	0xcb, 0x88, 0x6e, 0x92, 0x4a, 0x2a, 0x3a, 0xfa, 0x77, 0x83, 0xcc, 0x75, 0xd1, 0xaf, 0x9c, 0x8a,
	0x53, 0x3c, 0x5b, 0x63, 0x32, 0x88, 0xbb, 0xe3, 0x56, 0x97, 0x53, 0xf1, 0x02, 0xcf, 0x3a, 0x23,
	0x2c, 0x53, 0x27, 0x61, 0x31, 0xe4, 0x19, 0x21, 0xf2, 0x82, 0xc5, 0x40, 0xdf, 0x25, 0x35, 0x63,
	0xce, 0x94, 0x04, 0x50, 0x36, 0xb1, 0x2a, 0x62, 0xa7, 0x08, 0x4d, 0x22, 0xf0, 0x48, 0x8d, 0xdc,
	0x95, 0xa9, 0x08, 0x07, 0x91, 0x1a, 0xd1, 0xf7, 0x48, 0xdd, 0x98, 0xbf, 0x8f, 0x52, 0x2e, 0x42,
	0xb0, 0x69, 0x99, 0xb0, 0x2f, 0x0d, 0x36, 0x21, 0x71, 0x31, 0x48, 0x94, 0x1c, 0xb9, 0xa5, 0x29,
	0xd2, 0x81, 0xc1, 0x74, 0x85, 0x24, 0xf0, 0x28, 0x8d, 0x74, 0x7b, 0x50, 0xee, 0x9a, 0xa9, 0xd0,
	0x18, 0x45, 0xc9, 0x8f, 0xc9, 0xad, 0x09, 0xcd, 0xca, 0x2e, 0x23, 0x71, 0x63, 0x8c, 0x5b, 0xe9,
	0x85, 0x88, 0x28, 0xbf, 0x32, 0x13, 0x11, 0x53, 0x78, 0x4a, 0x1a, 0x13, 0x5a, 0x9e, 0x06, 0x41,
	0xe6, 0xe4, 0x57, 0x79, 0x2a, 0x05, 0x72, 0x9e, 0x4e, 0x75, 0x86, 0x9c, 0xa7, 0xe4, 0x91, 0x32,
	0x1f, 0x48, 0x09, 0x09, 0x1f, 0xb9, 0x35, 0xd3, 0x99, 0xfc, 0x3c, 0x29, 0x3d, 0x8b, 0x35, 0xdb,
	0xad, 0xfb, 0xce, 0xf6, 0xb2, 0x2d, 0xfd, 0x3e, 0x42, 0xba, 0xf4, 0x09, 0xa8, 0x9c, 0xb0, 0x8e,
	0x84, 0x4a, 0x02, 0x6a, 0x62, 0x56, 0x6c, 0x98, 0x9b, 0x37, 0x8c, 0x59, 0xb1, 0xa1, 0x35, 0xdf,
	0x23, 0x65, 0x6d, 0x96, 0xfa, 0xc2, 0xde, 0x42, 0xe3, 0x9a, 0x62, 0xc3, 0x40, 0x5f, 0xd9, 0xfb,
	0xa4, 0x32, 0xd6, 0xea, 0x36, 0x4c, 0x4b, 0xc7, 0x80, 0xbe, 0xe6, 0xa8, 0xc2, 0xa5, 0xe6, 0x9a,
	0xe3, 0x61, 0xba, 0x87, 0x89, 0x62, 0x5c, 0xb9, 0xb7, 0x0b, 0x3d, 0x44, 0x4c, 0x8f, 0x4a, 0x2e,
	0x62, 0xfd, 0x9a, 0xdc, 0x3b, 0x66, 0x54, 0xda, 0x23, 0xfd, 0x94, 0xd4, 0x43, 0xe8, 0x47, 0x17,
	0x20, 0x47, 0x1d, 0x3d, 0xa9, 0xdc, 0xbb, 0xf8, 0x9a, 0xbd, 0x96, 0x59, 0x07, 0xad, 0x7c, 0x1d,
	0xb4, 0xbe, 0xcc, 0xd7, 0x41, 0x50, 0xcb, 0x1d, 0x0e, 0xb5, 0xe6, 0x4f, 0x48, 0x4d, 0xfb, 0x75,
	0x38, 0xce, 0xd8, 0xd0, 0x7d, 0xeb, 0x4a, 0xff, 0xaa, 0xe6, 0x9b, 0x91, 0x8c, 0x4f, 0x1b, 0x86,
	0x4a, 0xb2, 0x0e, 0x8e, 0x92, 0xb7, 0x4d, 0xce, 0x88, 0xe8, 0x17, 0xb6, 0xf7, 0xeb, 0x2a, 0xd9,
	0x38, 0xb4, 0x83, 0xe0, 0xd4, 0x2c, 0x37, 0xfa, 0x93, 0x43, 0x4a, 0xc6, 0x9d, 0x3e, 0x5e, 0x34,
	0x74, 0x0a, 0x53, 0xdf, 0x7b, 0xb8, 0x88, 0x9a, 0x0f, 0xa4, 0xe6, 0x47, 0x97, 0xfb, 0x9e, 0xe7,
	0x1a, 0xcf, 0xcc, 0x67, 0x7e, 0xee, 0xe3, 0xa3, 0xd3, 0x0f, 0x7f, 0xfd, 0xfd, 0xf3, 0xd2, 0xed,
	0xe6, 0x7a, 0xbb, 0x10, 0xea, 0x63, 0xe7, 0x09, 0xfd, 0xdd, 0x21, 0x25, 0x33, 0xea, 0x17, 0x6b,
	0x2a, 0xac, 0x83, 0x6b, 0x6a, 0x3a, 0x40, 0x4d, 0xc6, 0xf3, 0x3f, 0x34, 0xf9, 0xde, 0x66, 0x51,
	0x53, 0xfb, 0xd5, 0x64, 0xab, 0xbc, 0xd6, 0x02, 0xff, 0x70, 0x70, 0x21, 0xdb, 0x15, 0x4b, 0x67,
	0xa7, 0xf5, 0xdc, 0xf6, 0xbd, 0xa6, 0xb6, 0x6f, 0x2e, 0xf7, 0x77, 0xbc, 0x27, 0x47, 0xa0, 0x7c,
	0xe6, 0x67, 0x29, 0xf0, 0xe8, 0x2c, 0xe2, 0xbe, 0x9d, 0xcc, 0xbe, 0x38, 0x7b, 0xb3, 0xda, 0xf7,
	0xe9, 0xa3, 0x05, 0x6a, 0xdb, 0xaf, 0xac, 0xff, 0x6b, 0xfa, 0x9b, 0x43, 0x96, 0x8f, 0x40, 0xd1,
	0x7b, 0xf3, 0x6a, 0x6f, 0x26, 0xf3, 0xf4, 0x72, 0x7f, 0xd7, 0x7b, 0xaa, 0x65, 0xaa, 0x73, 0xf0,
	0xcd, 0x5b, 0x57, 0x57, 0xea, 0xdc, 0xa2, 0x8b, 0xaa, 0xfa, 0xfc, 0x19, 0x69, 0x70, 0x11, 0x17,
	0xff, 0xff, 0xbc, 0x66, 0x6f, 0xe9, 0x89, 0xbe, 0xf7, 0x27, 0xce, 0xcb, 0x8d, 0x82, 0x39, 0xed,
	0x76, 0x4b, 0xf8, 0x22, 0x3e, 0xf8, 0x37, 0x00, 0x00, 0xff, 0xff, 0xf9, 0xa0, 0xe0, 0x99, 0xba,
	0x09, 0x00, 0x00,
}
