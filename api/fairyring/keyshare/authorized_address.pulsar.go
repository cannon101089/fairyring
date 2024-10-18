// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package keyshare

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_AuthorizedAddress               protoreflect.MessageDescriptor
	fd_AuthorizedAddress_target        protoreflect.FieldDescriptor
	fd_AuthorizedAddress_is_authorized protoreflect.FieldDescriptor
	fd_AuthorizedAddress_authorized_by protoreflect.FieldDescriptor
)

func init() {
	file_fairyring_keyshare_authorized_address_proto_init()
	md_AuthorizedAddress = File_fairyring_keyshare_authorized_address_proto.Messages().ByName("AuthorizedAddress")
	fd_AuthorizedAddress_target = md_AuthorizedAddress.Fields().ByName("target")
	fd_AuthorizedAddress_is_authorized = md_AuthorizedAddress.Fields().ByName("is_authorized")
	fd_AuthorizedAddress_authorized_by = md_AuthorizedAddress.Fields().ByName("authorized_by")
}

var _ protoreflect.Message = (*fastReflection_AuthorizedAddress)(nil)

type fastReflection_AuthorizedAddress AuthorizedAddress

func (x *AuthorizedAddress) ProtoReflect() protoreflect.Message {
	return (*fastReflection_AuthorizedAddress)(x)
}

func (x *AuthorizedAddress) slowProtoReflect() protoreflect.Message {
	mi := &file_fairyring_keyshare_authorized_address_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_AuthorizedAddress_messageType fastReflection_AuthorizedAddress_messageType
var _ protoreflect.MessageType = fastReflection_AuthorizedAddress_messageType{}

type fastReflection_AuthorizedAddress_messageType struct{}

func (x fastReflection_AuthorizedAddress_messageType) Zero() protoreflect.Message {
	return (*fastReflection_AuthorizedAddress)(nil)
}
func (x fastReflection_AuthorizedAddress_messageType) New() protoreflect.Message {
	return new(fastReflection_AuthorizedAddress)
}
func (x fastReflection_AuthorizedAddress_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_AuthorizedAddress
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_AuthorizedAddress) Descriptor() protoreflect.MessageDescriptor {
	return md_AuthorizedAddress
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_AuthorizedAddress) Type() protoreflect.MessageType {
	return _fastReflection_AuthorizedAddress_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_AuthorizedAddress) New() protoreflect.Message {
	return new(fastReflection_AuthorizedAddress)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_AuthorizedAddress) Interface() protoreflect.ProtoMessage {
	return (*AuthorizedAddress)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_AuthorizedAddress) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Target != "" {
		value := protoreflect.ValueOfString(x.Target)
		if !f(fd_AuthorizedAddress_target, value) {
			return
		}
	}
	if x.IsAuthorized != false {
		value := protoreflect.ValueOfBool(x.IsAuthorized)
		if !f(fd_AuthorizedAddress_is_authorized, value) {
			return
		}
	}
	if x.AuthorizedBy != "" {
		value := protoreflect.ValueOfString(x.AuthorizedBy)
		if !f(fd_AuthorizedAddress_authorized_by, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_AuthorizedAddress) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "fairyring.keyshare.AuthorizedAddress.target":
		return x.Target != ""
	case "fairyring.keyshare.AuthorizedAddress.is_authorized":
		return x.IsAuthorized != false
	case "fairyring.keyshare.AuthorizedAddress.authorized_by":
		return x.AuthorizedBy != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: fairyring.keyshare.AuthorizedAddress"))
		}
		panic(fmt.Errorf("message fairyring.keyshare.AuthorizedAddress does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AuthorizedAddress) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "fairyring.keyshare.AuthorizedAddress.target":
		x.Target = ""
	case "fairyring.keyshare.AuthorizedAddress.is_authorized":
		x.IsAuthorized = false
	case "fairyring.keyshare.AuthorizedAddress.authorized_by":
		x.AuthorizedBy = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: fairyring.keyshare.AuthorizedAddress"))
		}
		panic(fmt.Errorf("message fairyring.keyshare.AuthorizedAddress does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_AuthorizedAddress) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "fairyring.keyshare.AuthorizedAddress.target":
		value := x.Target
		return protoreflect.ValueOfString(value)
	case "fairyring.keyshare.AuthorizedAddress.is_authorized":
		value := x.IsAuthorized
		return protoreflect.ValueOfBool(value)
	case "fairyring.keyshare.AuthorizedAddress.authorized_by":
		value := x.AuthorizedBy
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: fairyring.keyshare.AuthorizedAddress"))
		}
		panic(fmt.Errorf("message fairyring.keyshare.AuthorizedAddress does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AuthorizedAddress) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "fairyring.keyshare.AuthorizedAddress.target":
		x.Target = value.Interface().(string)
	case "fairyring.keyshare.AuthorizedAddress.is_authorized":
		x.IsAuthorized = value.Bool()
	case "fairyring.keyshare.AuthorizedAddress.authorized_by":
		x.AuthorizedBy = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: fairyring.keyshare.AuthorizedAddress"))
		}
		panic(fmt.Errorf("message fairyring.keyshare.AuthorizedAddress does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AuthorizedAddress) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "fairyring.keyshare.AuthorizedAddress.target":
		panic(fmt.Errorf("field target of message fairyring.keyshare.AuthorizedAddress is not mutable"))
	case "fairyring.keyshare.AuthorizedAddress.is_authorized":
		panic(fmt.Errorf("field is_authorized of message fairyring.keyshare.AuthorizedAddress is not mutable"))
	case "fairyring.keyshare.AuthorizedAddress.authorized_by":
		panic(fmt.Errorf("field authorized_by of message fairyring.keyshare.AuthorizedAddress is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: fairyring.keyshare.AuthorizedAddress"))
		}
		panic(fmt.Errorf("message fairyring.keyshare.AuthorizedAddress does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_AuthorizedAddress) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "fairyring.keyshare.AuthorizedAddress.target":
		return protoreflect.ValueOfString("")
	case "fairyring.keyshare.AuthorizedAddress.is_authorized":
		return protoreflect.ValueOfBool(false)
	case "fairyring.keyshare.AuthorizedAddress.authorized_by":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: fairyring.keyshare.AuthorizedAddress"))
		}
		panic(fmt.Errorf("message fairyring.keyshare.AuthorizedAddress does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_AuthorizedAddress) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in fairyring.keyshare.AuthorizedAddress", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_AuthorizedAddress) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AuthorizedAddress) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_AuthorizedAddress) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_AuthorizedAddress) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*AuthorizedAddress)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Target)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.IsAuthorized {
			n += 2
		}
		l = len(x.AuthorizedBy)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*AuthorizedAddress)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.AuthorizedBy) > 0 {
			i -= len(x.AuthorizedBy)
			copy(dAtA[i:], x.AuthorizedBy)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.AuthorizedBy)))
			i--
			dAtA[i] = 0x1a
		}
		if x.IsAuthorized {
			i--
			if x.IsAuthorized {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x10
		}
		if len(x.Target) > 0 {
			i -= len(x.Target)
			copy(dAtA[i:], x.Target)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Target)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*AuthorizedAddress)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: AuthorizedAddress: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: AuthorizedAddress: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Target", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Target = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field IsAuthorized", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.IsAuthorized = bool(v != 0)
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AuthorizedBy", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.AuthorizedBy = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: fairyring/keyshare/authorized_address.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// AuthorizedAddress defines if an address is authorized to submit keyshares
type AuthorizedAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target       string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	IsAuthorized bool   `protobuf:"varint,2,opt,name=is_authorized,json=isAuthorized,proto3" json:"is_authorized,omitempty"`
	AuthorizedBy string `protobuf:"bytes,3,opt,name=authorized_by,json=authorizedBy,proto3" json:"authorized_by,omitempty"`
}

func (x *AuthorizedAddress) Reset() {
	*x = AuthorizedAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fairyring_keyshare_authorized_address_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizedAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizedAddress) ProtoMessage() {}

// Deprecated: Use AuthorizedAddress.ProtoReflect.Descriptor instead.
func (*AuthorizedAddress) Descriptor() ([]byte, []int) {
	return file_fairyring_keyshare_authorized_address_proto_rawDescGZIP(), []int{0}
}

func (x *AuthorizedAddress) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *AuthorizedAddress) GetIsAuthorized() bool {
	if x != nil {
		return x.IsAuthorized
	}
	return false
}

func (x *AuthorizedAddress) GetAuthorizedBy() string {
	if x != nil {
		return x.AuthorizedBy
	}
	return ""
}

var File_fairyring_keyshare_authorized_address_proto protoreflect.FileDescriptor

var file_fairyring_keyshare_authorized_address_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x66, 0x61, 0x69, 0x72, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x6b, 0x65, 0x79, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x66,
	0x61, 0x69, 0x72, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x22, 0x75, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x23,
	0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x64, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x42, 0x79, 0x42, 0xbe, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d,
	0x2e, 0x66, 0x61, 0x69, 0x72, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x42, 0x16, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x23, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x61, 0x69, 0x72, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x68, 0x61,
	0x72, 0x65, 0xa2, 0x02, 0x03, 0x46, 0x4b, 0x58, 0xaa, 0x02, 0x12, 0x46, 0x61, 0x69, 0x72, 0x79,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x4b, 0x65, 0x79, 0x73, 0x68, 0x61, 0x72, 0x65, 0xca, 0x02, 0x12,
	0x46, 0x61, 0x69, 0x72, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x5c, 0x4b, 0x65, 0x79, 0x73, 0x68, 0x61,
	0x72, 0x65, 0xe2, 0x02, 0x1e, 0x46, 0x61, 0x69, 0x72, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x5c, 0x4b,
	0x65, 0x79, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x46, 0x61, 0x69, 0x72, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x3a,
	0x3a, 0x4b, 0x65, 0x79, 0x73, 0x68, 0x61, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_fairyring_keyshare_authorized_address_proto_rawDescOnce sync.Once
	file_fairyring_keyshare_authorized_address_proto_rawDescData = file_fairyring_keyshare_authorized_address_proto_rawDesc
)

func file_fairyring_keyshare_authorized_address_proto_rawDescGZIP() []byte {
	file_fairyring_keyshare_authorized_address_proto_rawDescOnce.Do(func() {
		file_fairyring_keyshare_authorized_address_proto_rawDescData = protoimpl.X.CompressGZIP(file_fairyring_keyshare_authorized_address_proto_rawDescData)
	})
	return file_fairyring_keyshare_authorized_address_proto_rawDescData
}

var file_fairyring_keyshare_authorized_address_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_fairyring_keyshare_authorized_address_proto_goTypes = []interface{}{
	(*AuthorizedAddress)(nil), // 0: fairyring.keyshare.AuthorizedAddress
}
var file_fairyring_keyshare_authorized_address_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fairyring_keyshare_authorized_address_proto_init() }
func file_fairyring_keyshare_authorized_address_proto_init() {
	if File_fairyring_keyshare_authorized_address_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fairyring_keyshare_authorized_address_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizedAddress); i {
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
			RawDescriptor: file_fairyring_keyshare_authorized_address_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fairyring_keyshare_authorized_address_proto_goTypes,
		DependencyIndexes: file_fairyring_keyshare_authorized_address_proto_depIdxs,
		MessageInfos:      file_fairyring_keyshare_authorized_address_proto_msgTypes,
	}.Build()
	File_fairyring_keyshare_authorized_address_proto = out.File
	file_fairyring_keyshare_authorized_address_proto_rawDesc = nil
	file_fairyring_keyshare_authorized_address_proto_goTypes = nil
	file_fairyring_keyshare_authorized_address_proto_depIdxs = nil
}
