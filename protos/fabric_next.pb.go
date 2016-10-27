// Code generated by protoc-gen-go.
// source: fabric_next.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Message2_Type int32

const (
	// Undefined exists to prevent invalid message construction.
	Message2_UNDEFINED Message2_Type = 0
	// Handshake messages.
	Message2_DISCOVERY Message2_Type = 1
	// Sent to catch up with existing peers.
	Message2_SYNC Message2_Type = 2
	// Sent from SDK to endorser. Payload is a Proposal.
	Message2_PROPOSAL Message2_Type = 3
	// Reserved for future use.
	Message2_PROPOSAL_SET Message2_Type = 4
	// Sent from endorser to SDK. Payload is a ProposalResponse.
	Message2_PROPOSAL_RESULT Message2_Type = 5
	// Reserved for future use.
	Message2_PROPOSAL_SET_RESULT Message2_Type = 6
	// Sent from SDK to peer for relay or ordering service. Payload is a
	// Transaction2.
	Message2_TRANSACTION Message2_Type = 7
)

var Message2_Type_name = map[int32]string{
	0: "UNDEFINED",
	1: "DISCOVERY",
	2: "SYNC",
	3: "PROPOSAL",
	4: "PROPOSAL_SET",
	5: "PROPOSAL_RESULT",
	6: "PROPOSAL_SET_RESULT",
	7: "TRANSACTION",
}
var Message2_Type_value = map[string]int32{
	"UNDEFINED":           0,
	"DISCOVERY":           1,
	"SYNC":                2,
	"PROPOSAL":            3,
	"PROPOSAL_SET":        4,
	"PROPOSAL_RESULT":     5,
	"PROPOSAL_SET_RESULT": 6,
	"TRANSACTION":         7,
}

func (x Message2_Type) String() string {
	return proto.EnumName(Message2_Type_name, int32(x))
}
func (Message2_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{1, 0} }

type Proposal_Type int32

const (
	// Undefined exists to prevent invalid message construction.
	Proposal_UNDEFINED Proposal_Type = 0
	// A chaincode. The payload is a ChaincodeSpec.
	Proposal_CHAINCODE Proposal_Type = 1
)

var Proposal_Type_name = map[int32]string{
	0: "UNDEFINED",
	1: "CHAINCODE",
}
var Proposal_Type_value = map[string]int32{
	"UNDEFINED": 0,
	"CHAINCODE": 1,
}

func (x Proposal_Type) String() string {
	return proto.EnumName(Proposal_Type_name, int32(x))
}
func (Proposal_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{2, 0} }

type InvalidTransaction_Cause int32

const (
	InvalidTransaction_TxIdAlreadyExists      InvalidTransaction_Cause = 0
	InvalidTransaction_RWConflictDuringCommit InvalidTransaction_Cause = 1
)

var InvalidTransaction_Cause_name = map[int32]string{
	0: "TxIdAlreadyExists",
	1: "RWConflictDuringCommit",
}
var InvalidTransaction_Cause_value = map[string]int32{
	"TxIdAlreadyExists":      0,
	"RWConflictDuringCommit": 1,
}

func (x InvalidTransaction_Cause) String() string {
	return proto.EnumName(InvalidTransaction_Cause_name, int32(x))
}
func (InvalidTransaction_Cause) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{10, 0} }

// Envelope is used to deliver a message
type Envelope struct {
	// Signature of the message.
	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// The message.
	Message *Message2 `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Envelope) Reset()                    { *m = Envelope{} }
func (m *Envelope) String() string            { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()               {}
func (*Envelope) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *Envelope) GetMessage() *Message2 {
	if m != nil {
		return m.Message
	}
	return nil
}

// A Message2 encapsulates a payload of the indicated type in this message.
type Message2 struct {
	// Type of this message.
	Type Message2_Type `protobuf:"varint,1,opt,name=type,enum=protos.Message2_Type" json:"type,omitempty"`
	// Version indicates message protocol version.
	Version int32 `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	// Timestamp is the time that the message was created as defined by the
	// sender.
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
	// The payload in this message.
	Payload []byte `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Message2) Reset()                    { *m = Message2{} }
func (m *Message2) String() string            { return proto.CompactTextString(m) }
func (*Message2) ProtoMessage()               {}
func (*Message2) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *Message2) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

// A Proposal is sent to an endorser for endorsement. The proposal contains
// a payload (such as a ChaincodeSpec) based on the type field.
type Proposal struct {
	// Type of this message.
	Type Proposal_Type `protobuf:"varint,1,opt,name=type,enum=protos.Proposal_Type" json:"type,omitempty"`
	// Unique ID corresponding to this proposal
	Id string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	// The payload of the proposal as defined by the proposal type.
	Payload []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Proposal) Reset()                    { *m = Proposal{} }
func (m *Proposal) String() string            { return proto.CompactTextString(m) }
func (*Proposal) ProtoMessage()               {}
func (*Proposal) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

// A response with a representation similar to an HTTP response that can
// be used within another message.
type Response2 struct {
	// A status code that should follow the HTTP status codes.
	Status int32 `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	// A message associated with the response code.
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	// A payload that can be used to include metadata with this response.
	Payload []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Response2) Reset()                    { *m = Response2{} }
func (m *Response2) String() string            { return proto.CompactTextString(m) }
func (*Response2) ProtoMessage()               {}
func (*Response2) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

// A SystemChaincode is a chaincode compiled into the peer that cannot
// be modified at runtime. These are used to perform critical system level
// functions, including processing endorsements and validating transactions.
type SystemChaincode struct {
	// The ID used to identify a system chaincode.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *SystemChaincode) Reset()                    { *m = SystemChaincode{} }
func (m *SystemChaincode) String() string            { return proto.CompactTextString(m) }
func (*SystemChaincode) ProtoMessage()               {}
func (*SystemChaincode) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

// An action to be taken against the ledger.
type Action struct {
	// Hash of proposal encoded in the Message2 payload. NDD.
	ProposalHash []byte `protobuf:"bytes,1,opt,name=proposalHash,proto3" json:"proposalHash,omitempty"`
	// Uncommitted state changes (simulated) as calculated by the endorser.
	// This generally would include MVCC and PostImage information for both the
	// read set and write set. This is to be forwarded to the ordering
	// service as part of the transaction and must match the simulationResult
	// returned by other endorsers for the proposal.
	SimulationResult []byte `protobuf:"bytes,2,opt,name=simulationResult,proto3" json:"simulationResult,omitempty"`
	// Events that should be sent by committers after the transaction is written
	// to the ledger. This is to be forwarded to the ordering
	// service as part of the transaction and must match the events
	// returned by other endorsers for the proposal.
	Events [][]byte `protobuf:"bytes,3,rep,name=events,proto3" json:"events,omitempty"`
	// ESCC (Endorser System Chaincode) is logic that is run prior to the
	// ProposalResponse being returned to the SDK. It can manipulate the
	// ProposalResponse as needed.
	Escc *SystemChaincode `protobuf:"bytes,4,opt,name=escc" json:"escc,omitempty"`
	// VSCC (Validaing System Chaincode) is logic that is run to transform the
	// raw ledger into the validated ledger.
	Vscc *SystemChaincode `protobuf:"bytes,5,opt,name=vscc" json:"vscc,omitempty"`
}

func (m *Action) Reset()                    { *m = Action{} }
func (m *Action) String() string            { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()               {}
func (*Action) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *Action) GetEscc() *SystemChaincode {
	if m != nil {
		return m.Escc
	}
	return nil
}

func (m *Action) GetVscc() *SystemChaincode {
	if m != nil {
		return m.Vscc
	}
	return nil
}

// Endorsement is included within a proposal response.
type Endorsement struct {
	// Signature of the actionBytes included in the Endorsement.
	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Endorsement) Reset()                    { *m = Endorsement{} }
func (m *Endorsement) String() string            { return proto.CompactTextString(m) }
func (*Endorsement) ProtoMessage()               {}
func (*Endorsement) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

// A ProposalResponse is returned from an endorser to the proposal submitter.
type ProposalResponse struct {
	// A response message indicating whether the endorsement of the action
	// was successful. Additional metadata can be included. This will not
	// be forwarded from the SDK to the ordering service.
	Response *Response2 `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
	// A serialized Action message. NDD.
	ActionBytes []byte `protobuf:"bytes,2,opt,name=actionBytes,proto3" json:"actionBytes,omitempty"`
	// The endorsement of the action included in the proposal response
	Endorsement *Endorsement `protobuf:"bytes,3,opt,name=endorsement" json:"endorsement,omitempty"`
}

func (m *ProposalResponse) Reset()                    { *m = ProposalResponse{} }
func (m *ProposalResponse) String() string            { return proto.CompactTextString(m) }
func (*ProposalResponse) ProtoMessage()               {}
func (*ProposalResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{7} }

func (m *ProposalResponse) GetResponse() *Response2 {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *ProposalResponse) GetEndorsement() *Endorsement {
	if m != nil {
		return m.Endorsement
	}
	return nil
}

// An EndorsedAction describes a single action endorsed by one or more
// endorsers. Multiple endorsed actions can be included in a single
// transaction. The transaction is atomic meaning that either all
// actions in the transaction will be committed or none will be committed.
type EndorsedAction struct {
	// The action to be taken against the ledger. This is generally constructed
	// by an endorser. NDD.
	ActionBytes []byte `protobuf:"bytes,1,opt,name=actionBytes,proto3" json:"actionBytes,omitempty"`
	// The endorsements of the action.
	Endorsements []*Endorsement `protobuf:"bytes,2,rep,name=endorsements" json:"endorsements,omitempty"`
	// The proposal. This is optional and only needed if the SDK wants to store
	// the Proposal on the ledger as opposed to just the hash. The proposal is
	// not included within the Action because it is the SDK's decision whether
	// or not they would like to include this information in the Transaction2.
	// If it was in the Action and signed, either the Endorsers would be
	// required to make the decision or the SDK would need to provide a hint
	// in the Proposal about whether it should be included in the Action.
	// TODO Revisit this decision.
	ProposalBytes []byte `protobuf:"bytes,3,opt,name=proposalBytes,proto3" json:"proposalBytes,omitempty"`
}

func (m *EndorsedAction) Reset()                    { *m = EndorsedAction{} }
func (m *EndorsedAction) String() string            { return proto.CompactTextString(m) }
func (*EndorsedAction) ProtoMessage()               {}
func (*EndorsedAction) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{8} }

func (m *EndorsedAction) GetEndorsements() []*Endorsement {
	if m != nil {
		return m.Endorsements
	}
	return nil
}

// The transaction to be sent to the ordering service. A transaction contains
// one or more endorsed actions. The transaction is atomic meaning that either
// all actions in the transaction will be committed or none will be committed.
type Transaction2 struct {
	// One or more endorsed actions to be committed to the ledger.
	EndorsedActions []*EndorsedAction `protobuf:"bytes,1,rep,name=endorsedActions" json:"endorsedActions,omitempty"`
}

func (m *Transaction2) Reset()                    { *m = Transaction2{} }
func (m *Transaction2) String() string            { return proto.CompactTextString(m) }
func (*Transaction2) ProtoMessage()               {}
func (*Transaction2) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{9} }

func (m *Transaction2) GetEndorsedActions() []*EndorsedAction {
	if m != nil {
		return m.EndorsedActions
	}
	return nil
}

// This is used to wrap an invalid Transaction with the cause
type InvalidTransaction struct {
	Transaction *Transaction2            `protobuf:"bytes,1,opt,name=transaction" json:"transaction,omitempty"`
	Cause       InvalidTransaction_Cause `protobuf:"varint,2,opt,name=cause,enum=protos.InvalidTransaction_Cause" json:"cause,omitempty"`
}

func (m *InvalidTransaction) Reset()                    { *m = InvalidTransaction{} }
func (m *InvalidTransaction) String() string            { return proto.CompactTextString(m) }
func (*InvalidTransaction) ProtoMessage()               {}
func (*InvalidTransaction) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{10} }

func (m *InvalidTransaction) GetTransaction() *Transaction2 {
	if m != nil {
		return m.Transaction
	}
	return nil
}

// Block contains a list of transactions and the crypto hash of previous block
type Block2 struct {
	PreviousBlockHash []byte `protobuf:"bytes,1,opt,name=PreviousBlockHash,json=previousBlockHash,proto3" json:"PreviousBlockHash,omitempty"`
	// transactions are stored in serialized form so that the concenters can avoid marshaling of transactions
	Transactions [][]byte `protobuf:"bytes,2,rep,name=Transactions,json=transactions,proto3" json:"Transactions,omitempty"`
}

func (m *Block2) Reset()                    { *m = Block2{} }
func (m *Block2) String() string            { return proto.CompactTextString(m) }
func (*Block2) ProtoMessage()               {}
func (*Block2) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{11} }

func init() {
	proto.RegisterType((*Envelope)(nil), "protos.Envelope")
	proto.RegisterType((*Message2)(nil), "protos.Message2")
	proto.RegisterType((*Proposal)(nil), "protos.Proposal")
	proto.RegisterType((*Response2)(nil), "protos.Response2")
	proto.RegisterType((*SystemChaincode)(nil), "protos.SystemChaincode")
	proto.RegisterType((*Action)(nil), "protos.Action")
	proto.RegisterType((*Endorsement)(nil), "protos.Endorsement")
	proto.RegisterType((*ProposalResponse)(nil), "protos.ProposalResponse")
	proto.RegisterType((*EndorsedAction)(nil), "protos.EndorsedAction")
	proto.RegisterType((*Transaction2)(nil), "protos.Transaction2")
	proto.RegisterType((*InvalidTransaction)(nil), "protos.InvalidTransaction")
	proto.RegisterType((*Block2)(nil), "protos.Block2")
	proto.RegisterEnum("protos.Message2_Type", Message2_Type_name, Message2_Type_value)
	proto.RegisterEnum("protos.Proposal_Type", Proposal_Type_name, Proposal_Type_value)
	proto.RegisterEnum("protos.InvalidTransaction_Cause", InvalidTransaction_Cause_name, InvalidTransaction_Cause_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Endorser service

type EndorserClient interface {
	ProcessProposal(ctx context.Context, in *Proposal, opts ...grpc.CallOption) (*ProposalResponse, error)
}

type endorserClient struct {
	cc *grpc.ClientConn
}

func NewEndorserClient(cc *grpc.ClientConn) EndorserClient {
	return &endorserClient{cc}
}

func (c *endorserClient) ProcessProposal(ctx context.Context, in *Proposal, opts ...grpc.CallOption) (*ProposalResponse, error) {
	out := new(ProposalResponse)
	err := grpc.Invoke(ctx, "/protos.Endorser/ProcessProposal", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Endorser service

type EndorserServer interface {
	ProcessProposal(context.Context, *Proposal) (*ProposalResponse, error)
}

func RegisterEndorserServer(s *grpc.Server, srv EndorserServer) {
	s.RegisterService(&_Endorser_serviceDesc, srv)
}

func _Endorser_ProcessProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Proposal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndorserServer).ProcessProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Endorser/ProcessProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndorserServer).ProcessProposal(ctx, req.(*Proposal))
	}
	return interceptor(ctx, in, info, handler)
}

var _Endorser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Endorser",
	HandlerType: (*EndorserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessProposal",
			Handler:    _Endorser_ProcessProposal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor5,
}

func init() { proto.RegisterFile("fabric_next.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 811 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x55, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x0e, 0xad, 0x1f, 0x4b, 0x43, 0xc5, 0xa6, 0xd6, 0x8d, 0x23, 0x18, 0x05, 0xea, 0x12, 0x39,
	0xb4, 0x4d, 0xab, 0x00, 0x2c, 0xfa, 0x83, 0x5e, 0x5a, 0x85, 0x62, 0x11, 0xa1, 0xa9, 0x64, 0xac,
	0x98, 0x06, 0xe9, 0x25, 0xa0, 0xc9, 0xb5, 0x42, 0x94, 0xe2, 0x12, 0xdc, 0xa5, 0x60, 0x3d, 0x40,
	0x8f, 0x3d, 0xf4, 0xda, 0xd7, 0xe8, 0x0b, 0xf4, 0xdc, 0xa7, 0xea, 0xee, 0x92, 0x2b, 0x51, 0x34,
	0x9c, 0x9e, 0xa4, 0x99, 0xf9, 0x76, 0xbe, 0x6f, 0x66, 0x76, 0x96, 0x30, 0xbc, 0x09, 0xae, 0xf3,
	0x38, 0x7c, 0x9b, 0x92, 0x5b, 0x3e, 0xce, 0x72, 0xca, 0x29, 0xea, 0xaa, 0x1f, 0x76, 0xf1, 0xd1,
	0x8a, 0xd2, 0x55, 0x42, 0x9e, 0x29, 0xf3, 0xba, 0xb8, 0x79, 0xc6, 0xe3, 0x35, 0x61, 0x3c, 0x58,
	0x67, 0x25, 0xd0, 0xf6, 0xa1, 0xe7, 0xa5, 0x1b, 0x92, 0xd0, 0x8c, 0xa0, 0x0f, 0xa1, 0xcf, 0xe2,
	0x55, 0x1a, 0xf0, 0x22, 0x27, 0x23, 0xe3, 0xd2, 0xf8, 0x64, 0x80, 0xf7, 0x0e, 0xf4, 0x19, 0x1c,
	0x8b, 0xa3, 0x2c, 0x58, 0x91, 0xd1, 0x91, 0x88, 0x99, 0x8e, 0x55, 0xa6, 0x60, 0xe3, 0x9f, 0x4b,
	0xb7, 0x83, 0x35, 0xc0, 0xfe, 0xfb, 0x08, 0x7a, 0xda, 0x8b, 0x3e, 0x85, 0x36, 0xdf, 0x66, 0x65,
	0xc6, 0x13, 0xe7, 0x51, 0xf3, 0xd4, 0xd8, 0x17, 0x41, 0xac, 0x20, 0x68, 0x04, 0xc7, 0x1b, 0x92,
	0xb3, 0x98, 0xa6, 0x8a, 0xa3, 0x83, 0xb5, 0x89, 0xbe, 0x85, 0xfe, 0x4e, 0xfa, 0xa8, 0xa5, 0xf8,
	0x2f, 0xc6, 0x65, 0x71, 0x63, 0x5d, 0xdc, 0xd8, 0xd7, 0x08, 0xbc, 0x07, 0xcb, 0x9c, 0x59, 0xb0,
	0x4d, 0x68, 0x10, 0x8d, 0xda, 0xaa, 0x26, 0x6d, 0xda, 0x7f, 0x18, 0xd0, 0x96, 0xe4, 0xe8, 0x21,
	0xf4, 0x5f, 0xcd, 0xa7, 0xde, 0x8f, 0xb3, 0xb9, 0x37, 0xb5, 0x1e, 0x48, 0x73, 0x3a, 0x5b, 0xba,
	0x8b, 0x5f, 0x3c, 0xfc, 0xc6, 0x32, 0x50, 0x0f, 0xda, 0xcb, 0x37, 0x73, 0xd7, 0x3a, 0x42, 0x03,
	0xe8, 0x5d, 0xe1, 0xc5, 0xd5, 0x62, 0x39, 0x79, 0x69, 0xb5, 0x90, 0x05, 0x03, 0x6d, 0xbd, 0x5d,
	0x7a, 0xbe, 0xd5, 0x46, 0x67, 0x70, 0xba, 0xf3, 0x60, 0x6f, 0xf9, 0xea, 0xa5, 0x6f, 0x75, 0xd0,
	0x63, 0x38, 0xab, 0xc3, 0x74, 0xa0, 0x8b, 0x4e, 0xc1, 0xf4, 0xf1, 0x64, 0xbe, 0x9c, 0xb8, 0xfe,
	0x6c, 0x31, 0xb7, 0x8e, 0xed, 0xdf, 0x0d, 0x91, 0x3f, 0xa7, 0x19, 0x65, 0x41, 0x72, 0x5f, 0xd7,
	0x74, 0xbc, 0xde, 0xb5, 0x13, 0x38, 0x8a, 0x23, 0xd5, 0xb0, 0x3e, 0x16, 0xff, 0xea, 0x15, 0xb7,
	0x0e, 0x2b, 0x7e, 0x72, 0x6f, 0xc1, 0xee, 0x8b, 0xc9, 0x6c, 0xee, 0x2e, 0xa6, 0x9e, 0x65, 0xd8,
	0xaf, 0xa1, 0x8f, 0x09, 0xcb, 0x68, 0xca, 0xc4, 0xf4, 0xce, 0xa1, 0x2b, 0xfa, 0xc8, 0x0b, 0xa6,
	0x94, 0x74, 0x70, 0x65, 0x49, 0x92, 0xfa, 0x75, 0xe8, 0xef, 0x86, 0xff, 0x1e, 0xfa, 0x8f, 0xe1,
	0x74, 0xb9, 0x65, 0x9c, 0xac, 0xdd, 0x77, 0x41, 0x9c, 0x86, 0x34, 0xd2, 0xda, 0x0d, 0xad, 0xdd,
	0xfe, 0xd7, 0x80, 0xee, 0x24, 0xe4, 0x72, 0xe4, 0x36, 0x0c, 0xb2, 0xaa, 0xda, 0x17, 0x01, 0x7b,
	0x57, 0xdd, 0xc8, 0x03, 0x9f, 0xb8, 0x94, 0x16, 0x8b, 0xd7, 0x45, 0x12, 0xc8, 0x13, 0x42, 0x74,
	0x91, 0x70, 0x25, 0x67, 0x80, 0xef, 0xf8, 0x65, 0x25, 0x64, 0x43, 0x52, 0xce, 0x84, 0xac, 0x96,
	0x40, 0x54, 0x16, 0x7a, 0x0a, 0x6d, 0xc2, 0xc2, 0x50, 0xdd, 0x0e, 0xd3, 0x79, 0xac, 0x3b, 0xdd,
	0x50, 0x8a, 0x15, 0x48, 0x82, 0x37, 0x12, 0xdc, 0xf9, 0x1f, 0xb0, 0x04, 0xd9, 0x4f, 0xc1, 0xf4,
	0xd2, 0x88, 0xe6, 0x8c, 0xac, 0x05, 0xd3, 0xfb, 0xf7, 0xcb, 0xfe, 0xcb, 0x00, 0x4b, 0x4f, 0x57,
	0xb7, 0x1f, 0x7d, 0x01, 0xbd, 0xbc, 0xfa, 0xaf, 0x4e, 0x98, 0xce, 0x50, 0x53, 0xee, 0x46, 0x84,
	0x77, 0x10, 0x74, 0x09, 0x66, 0xa0, 0x9a, 0xf7, 0x7c, 0xcb, 0x09, 0xab, 0x3a, 0x51, 0x77, 0xa1,
	0xaf, 0xc0, 0x24, 0x7b, 0x49, 0xd5, 0x26, 0x9d, 0xe9, 0x9c, 0x35, 0xb5, 0xb8, 0x8e, 0xb3, 0xff,
	0x34, 0xe0, 0xa4, 0x0a, 0x46, 0xd5, 0x78, 0x1a, 0x5c, 0xc6, 0x5d, 0xae, 0x6f, 0x60, 0x50, 0xcb,
	0x21, 0xe5, 0xb4, 0xee, 0x23, 0x3b, 0x00, 0xa2, 0x27, 0xf0, 0x50, 0x4f, 0xb9, 0x4c, 0x5e, 0xde,
	0xa3, 0x43, 0xa7, 0x7d, 0x05, 0x03, 0x3f, 0x0f, 0x52, 0x56, 0x52, 0x3a, 0xe8, 0x07, 0x38, 0x25,
	0x07, 0x12, 0xa5, 0x28, 0xc9, 0x78, 0xde, 0x60, 0xac, 0xc2, 0xb8, 0x09, 0xb7, 0xff, 0x31, 0x00,
	0xcd, 0xd2, 0x4d, 0x90, 0xc4, 0x51, 0x2d, 0x33, 0xfa, 0x1a, 0x4c, 0xbe, 0x37, 0xab, 0x39, 0x7c,
	0xa0, 0x93, 0xd6, 0x35, 0xe0, 0x3a, 0x50, 0x9c, 0xeb, 0x84, 0x41, 0xc1, 0xca, 0x05, 0x39, 0x71,
	0x2e, 0xf5, 0x89, 0xbb, 0x14, 0x63, 0x57, 0xe2, 0x70, 0x09, 0xb7, 0xbf, 0x83, 0x8e, 0xb2, 0xd1,
	0x23, 0x18, 0xfa, 0xb7, 0xb3, 0x68, 0x92, 0xe4, 0x24, 0x88, 0xb6, 0xde, 0x6d, 0xcc, 0x38, 0x13,
	0xeb, 0x7a, 0x01, 0xe7, 0xf8, 0xb5, 0x4b, 0xd3, 0x9b, 0x24, 0x0e, 0xf9, 0xb4, 0xc8, 0xe3, 0x74,
	0xe5, 0xd2, 0xf5, 0x3a, 0xe6, 0x62, 0x77, 0x7f, 0x85, 0xee, 0xf3, 0x84, 0x86, 0xbf, 0x39, 0xe8,
	0x73, 0x18, 0x5e, 0xe5, 0x64, 0x13, 0xd3, 0x82, 0x29, 0x4f, 0x6d, 0x87, 0x86, 0x59, 0x33, 0x20,
	0x97, 0xad, 0xa6, 0xa7, 0x9c, 0x95, 0x58, 0xb6, 0x5a, 0x39, 0xcc, 0xf9, 0x49, 0x7e, 0x2b, 0x54,
	0xc7, 0x72, 0xf4, 0xbd, 0x78, 0xea, 0x72, 0x1a, 0x8a, 0x95, 0xdf, 0xbd, 0x58, 0x56, 0xf3, 0x8d,
	0xba, 0x18, 0x35, 0x3d, 0xfa, 0xce, 0xda, 0x0f, 0xae, 0xcb, 0x2f, 0xd4, 0x97, 0xff, 0x05, 0x00,
	0x00, 0xff, 0xff, 0x1e, 0x94, 0x2e, 0x4c, 0xbd, 0x06, 0x00, 0x00,
}
