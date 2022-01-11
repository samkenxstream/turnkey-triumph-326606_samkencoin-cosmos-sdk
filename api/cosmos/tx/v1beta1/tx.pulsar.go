package txv1beta1

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	v1beta12 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	v1beta11 "github.com/cosmos/cosmos-sdk/api/cosmos/crypto/multisig/v1beta1"
	v1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/tx/signing/v1beta1"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var _ protoreflect.List = (*_Tx_3_list)(nil)

type _Tx_3_list struct {
	list *[][]byte
}

func (x *_Tx_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_Tx_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfBytes((*x.list)[i])
}

func (x *_Tx_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Bytes()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_Tx_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Bytes()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_Tx_3_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message Tx at list field Signatures as it is not of Message kind"))
}

func (x *_Tx_3_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_Tx_3_list) NewElement() protoreflect.Value {
	var v []byte
	return protoreflect.ValueOfBytes(v)
}

func (x *_Tx_3_list) IsValid() bool {
	return x.list != nil
}

var (
	md_Tx            protoreflect.MessageDescriptor
	fd_Tx_body       protoreflect.FieldDescriptor
	fd_Tx_auth_info  protoreflect.FieldDescriptor
	fd_Tx_signatures protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_tx_v1beta1_tx_proto_init()
	md_Tx = File_cosmos_tx_v1beta1_tx_proto.Messages().ByName("Tx")
	fd_Tx_body = md_Tx.Fields().ByName("body")
	fd_Tx_auth_info = md_Tx.Fields().ByName("auth_info")
	fd_Tx_signatures = md_Tx.Fields().ByName("signatures")
}

var _ protoreflect.Message = (*fastReflection_Tx)(nil)

type fastReflection_Tx Tx

func (x *Tx) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Tx)(x)
}

func (x *Tx) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_tx_v1beta1_tx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Tx_messageType fastReflection_Tx_messageType
var _ protoreflect.MessageType = fastReflection_Tx_messageType{}

type fastReflection_Tx_messageType struct{}

func (x fastReflection_Tx_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Tx)(nil)
}
func (x fastReflection_Tx_messageType) New() protoreflect.Message {
	return new(fastReflection_Tx)
}
func (x fastReflection_Tx_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Tx
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Tx) Descriptor() protoreflect.MessageDescriptor {
	return md_Tx
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Tx) Type() protoreflect.MessageType {
	return _fastReflection_Tx_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Tx) New() protoreflect.Message {
	return new(fastReflection_Tx)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Tx) Interface() protoreflect.ProtoMessage {
	return (*Tx)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Tx) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Body != nil {
		value := protoreflect.ValueOfMessage(x.Body.ProtoReflect())
		if !f(fd_Tx_body, value) {
			return
		}
	}
	if x.AuthInfo != nil {
		value := protoreflect.ValueOfMessage(x.AuthInfo.ProtoReflect())
		if !f(fd_Tx_auth_info, value) {
			return
		}
	}
	if len(x.Signatures) != 0 {
		value := protoreflect.ValueOfList(&_Tx_3_list{list: &x.Signatures})
		if !f(fd_Tx_signatures, value) {
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
func (x *fastReflection_Tx) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.Tx.body":
		return x.Body != nil
	case "cosmos.tx.v1beta1.Tx.auth_info":
		return x.AuthInfo != nil
	case "cosmos.tx.v1beta1.Tx.signatures":
		return len(x.Signatures) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.Tx"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.Tx does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Tx) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.Tx.body":
		x.Body = nil
	case "cosmos.tx.v1beta1.Tx.auth_info":
		x.AuthInfo = nil
	case "cosmos.tx.v1beta1.Tx.signatures":
		x.Signatures = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.Tx"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.Tx does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Tx) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.tx.v1beta1.Tx.body":
		value := x.Body
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.tx.v1beta1.Tx.auth_info":
		value := x.AuthInfo
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.tx.v1beta1.Tx.signatures":
		if len(x.Signatures) == 0 {
			return protoreflect.ValueOfList(&_Tx_3_list{})
		}
		listValue := &_Tx_3_list{list: &x.Signatures}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.Tx"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.Tx does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Tx) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.Tx.body":
		x.Body = value.Message().Interface().(*TxBody)
	case "cosmos.tx.v1beta1.Tx.auth_info":
		x.AuthInfo = value.Message().Interface().(*AuthInfo)
	case "cosmos.tx.v1beta1.Tx.signatures":
		lv := value.List()
		clv := lv.(*_Tx_3_list)
		x.Signatures = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.Tx"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.Tx does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Tx) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.Tx.body":
		if x.Body == nil {
			x.Body = new(TxBody)
		}
		return protoreflect.ValueOfMessage(x.Body.ProtoReflect())
	case "cosmos.tx.v1beta1.Tx.auth_info":
		if x.AuthInfo == nil {
			x.AuthInfo = new(AuthInfo)
		}
		return protoreflect.ValueOfMessage(x.AuthInfo.ProtoReflect())
	case "cosmos.tx.v1beta1.Tx.signatures":
		if x.Signatures == nil {
			x.Signatures = [][]byte{}
		}
		value := &_Tx_3_list{list: &x.Signatures}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.Tx"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.Tx does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Tx) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.Tx.body":
		m := new(TxBody)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.tx.v1beta1.Tx.auth_info":
		m := new(AuthInfo)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.tx.v1beta1.Tx.signatures":
		list := [][]byte{}
		return protoreflect.ValueOfList(&_Tx_3_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.Tx"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.Tx does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Tx) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.tx.v1beta1.Tx", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Tx) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Tx) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Tx) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Tx) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Tx)
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
		if x.Body != nil {
			l = options.Size(x.Body)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.AuthInfo != nil {
			l = options.Size(x.AuthInfo)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Signatures) > 0 {
			for _, b := range x.Signatures {
				l = len(b)
				n += 1 + l + runtime.Sov(uint64(l))
			}
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
		x := input.Message.Interface().(*Tx)
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
		if len(x.Signatures) > 0 {
			for iNdEx := len(x.Signatures) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.Signatures[iNdEx])
				copy(dAtA[i:], x.Signatures[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Signatures[iNdEx])))
				i--
				dAtA[i] = 0x1a
			}
		}
		if x.AuthInfo != nil {
			encoded, err := options.Marshal(x.AuthInfo)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if x.Body != nil {
			encoded, err := options.Marshal(x.Body)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
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
		x := input.Message.Interface().(*Tx)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Tx: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Tx: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Body == nil {
					x.Body = &TxBody{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Body); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AuthInfo", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.AuthInfo == nil {
					x.AuthInfo = &AuthInfo{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.AuthInfo); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Signatures", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Signatures = append(x.Signatures, make([]byte, postIndex-iNdEx))
				copy(x.Signatures[len(x.Signatures)-1], dAtA[iNdEx:postIndex])
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

var _ protoreflect.List = (*_TxRaw_3_list)(nil)

type _TxRaw_3_list struct {
	list *[][]byte
}

func (x *_TxRaw_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TxRaw_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfBytes((*x.list)[i])
}

func (x *_TxRaw_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Bytes()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_TxRaw_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Bytes()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_TxRaw_3_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message TxRaw at list field Signatures as it is not of Message kind"))
}

func (x *_TxRaw_3_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_TxRaw_3_list) NewElement() protoreflect.Value {
	var v []byte
	return protoreflect.ValueOfBytes(v)
}

func (x *_TxRaw_3_list) IsValid() bool {
	return x.list != nil
}

var (
	md_TxRaw                 protoreflect.MessageDescriptor
	fd_TxRaw_body_bytes      protoreflect.FieldDescriptor
	fd_TxRaw_auth_info_bytes protoreflect.FieldDescriptor
	fd_TxRaw_signatures      protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_tx_v1beta1_tx_proto_init()
	md_TxRaw = File_cosmos_tx_v1beta1_tx_proto.Messages().ByName("TxRaw")
	fd_TxRaw_body_bytes = md_TxRaw.Fields().ByName("body_bytes")
	fd_TxRaw_auth_info_bytes = md_TxRaw.Fields().ByName("auth_info_bytes")
	fd_TxRaw_signatures = md_TxRaw.Fields().ByName("signatures")
}

var _ protoreflect.Message = (*fastReflection_TxRaw)(nil)

type fastReflection_TxRaw TxRaw

func (x *TxRaw) ProtoReflect() protoreflect.Message {
	return (*fastReflection_TxRaw)(x)
}

func (x *TxRaw) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_tx_v1beta1_tx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_TxRaw_messageType fastReflection_TxRaw_messageType
var _ protoreflect.MessageType = fastReflection_TxRaw_messageType{}

type fastReflection_TxRaw_messageType struct{}

func (x fastReflection_TxRaw_messageType) Zero() protoreflect.Message {
	return (*fastReflection_TxRaw)(nil)
}
func (x fastReflection_TxRaw_messageType) New() protoreflect.Message {
	return new(fastReflection_TxRaw)
}
func (x fastReflection_TxRaw_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_TxRaw
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_TxRaw) Descriptor() protoreflect.MessageDescriptor {
	return md_TxRaw
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_TxRaw) Type() protoreflect.MessageType {
	return _fastReflection_TxRaw_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_TxRaw) New() protoreflect.Message {
	return new(fastReflection_TxRaw)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_TxRaw) Interface() protoreflect.ProtoMessage {
	return (*TxRaw)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_TxRaw) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.BodyBytes) != 0 {
		value := protoreflect.ValueOfBytes(x.BodyBytes)
		if !f(fd_TxRaw_body_bytes, value) {
			return
		}
	}
	if len(x.AuthInfoBytes) != 0 {
		value := protoreflect.ValueOfBytes(x.AuthInfoBytes)
		if !f(fd_TxRaw_auth_info_bytes, value) {
			return
		}
	}
	if len(x.Signatures) != 0 {
		value := protoreflect.ValueOfList(&_TxRaw_3_list{list: &x.Signatures})
		if !f(fd_TxRaw_signatures, value) {
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
func (x *fastReflection_TxRaw) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.TxRaw.body_bytes":
		return len(x.BodyBytes) != 0
	case "cosmos.tx.v1beta1.TxRaw.auth_info_bytes":
		return len(x.AuthInfoBytes) != 0
	case "cosmos.tx.v1beta1.TxRaw.signatures":
		return len(x.Signatures) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.TxRaw"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.TxRaw does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TxRaw) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.TxRaw.body_bytes":
		x.BodyBytes = nil
	case "cosmos.tx.v1beta1.TxRaw.auth_info_bytes":
		x.AuthInfoBytes = nil
	case "cosmos.tx.v1beta1.TxRaw.signatures":
		x.Signatures = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.TxRaw"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.TxRaw does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_TxRaw) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.tx.v1beta1.TxRaw.body_bytes":
		value := x.BodyBytes
		return protoreflect.ValueOfBytes(value)
	case "cosmos.tx.v1beta1.TxRaw.auth_info_bytes":
		value := x.AuthInfoBytes
		return protoreflect.ValueOfBytes(value)
	case "cosmos.tx.v1beta1.TxRaw.signatures":
		if len(x.Signatures) == 0 {
			return protoreflect.ValueOfList(&_TxRaw_3_list{})
		}
		listValue := &_TxRaw_3_list{list: &x.Signatures}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.TxRaw"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.TxRaw does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_TxRaw) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.TxRaw.body_bytes":
		x.BodyBytes = value.Bytes()
	case "cosmos.tx.v1beta1.TxRaw.auth_info_bytes":
		x.AuthInfoBytes = value.Bytes()
	case "cosmos.tx.v1beta1.TxRaw.signatures":
		lv := value.List()
		clv := lv.(*_TxRaw_3_list)
		x.Signatures = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.TxRaw"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.TxRaw does not contain field %s", fd.FullName()))
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
func (x *fastReflection_TxRaw) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.TxRaw.signatures":
		if x.Signatures == nil {
			x.Signatures = [][]byte{}
		}
		value := &_TxRaw_3_list{list: &x.Signatures}
		return protoreflect.ValueOfList(value)
	case "cosmos.tx.v1beta1.TxRaw.body_bytes":
		panic(fmt.Errorf("field body_bytes of message cosmos.tx.v1beta1.TxRaw is not mutable"))
	case "cosmos.tx.v1beta1.TxRaw.auth_info_bytes":
		panic(fmt.Errorf("field auth_info_bytes of message cosmos.tx.v1beta1.TxRaw is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.TxRaw"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.TxRaw does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_TxRaw) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.TxRaw.body_bytes":
		return protoreflect.ValueOfBytes(nil)
	case "cosmos.tx.v1beta1.TxRaw.auth_info_bytes":
		return protoreflect.ValueOfBytes(nil)
	case "cosmos.tx.v1beta1.TxRaw.signatures":
		list := [][]byte{}
		return protoreflect.ValueOfList(&_TxRaw_3_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.TxRaw"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.TxRaw does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_TxRaw) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.tx.v1beta1.TxRaw", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_TxRaw) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TxRaw) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_TxRaw) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_TxRaw) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*TxRaw)
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
		l = len(x.BodyBytes)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.AuthInfoBytes)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Signatures) > 0 {
			for _, b := range x.Signatures {
				l = len(b)
				n += 1 + l + runtime.Sov(uint64(l))
			}
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
		x := input.Message.Interface().(*TxRaw)
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
		if len(x.Signatures) > 0 {
			for iNdEx := len(x.Signatures) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.Signatures[iNdEx])
				copy(dAtA[i:], x.Signatures[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Signatures[iNdEx])))
				i--
				dAtA[i] = 0x1a
			}
		}
		if len(x.AuthInfoBytes) > 0 {
			i -= len(x.AuthInfoBytes)
			copy(dAtA[i:], x.AuthInfoBytes)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.AuthInfoBytes)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.BodyBytes) > 0 {
			i -= len(x.BodyBytes)
			copy(dAtA[i:], x.BodyBytes)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BodyBytes)))
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
		x := input.Message.Interface().(*TxRaw)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: TxRaw: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: TxRaw: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BodyBytes", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.BodyBytes = append(x.BodyBytes[:0], dAtA[iNdEx:postIndex]...)
				if x.BodyBytes == nil {
					x.BodyBytes = []byte{}
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AuthInfoBytes", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.AuthInfoBytes = append(x.AuthInfoBytes[:0], dAtA[iNdEx:postIndex]...)
				if x.AuthInfoBytes == nil {
					x.AuthInfoBytes = []byte{}
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Signatures", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Signatures = append(x.Signatures, make([]byte, postIndex-iNdEx))
				copy(x.Signatures[len(x.Signatures)-1], dAtA[iNdEx:postIndex])
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

var (
	md_SignDoc                 protoreflect.MessageDescriptor
	fd_SignDoc_body_bytes      protoreflect.FieldDescriptor
	fd_SignDoc_auth_info_bytes protoreflect.FieldDescriptor
	fd_SignDoc_chain_id        protoreflect.FieldDescriptor
	fd_SignDoc_account_number  protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_tx_v1beta1_tx_proto_init()
	md_SignDoc = File_cosmos_tx_v1beta1_tx_proto.Messages().ByName("SignDoc")
	fd_SignDoc_body_bytes = md_SignDoc.Fields().ByName("body_bytes")
	fd_SignDoc_auth_info_bytes = md_SignDoc.Fields().ByName("auth_info_bytes")
	fd_SignDoc_chain_id = md_SignDoc.Fields().ByName("chain_id")
	fd_SignDoc_account_number = md_SignDoc.Fields().ByName("account_number")
}

var _ protoreflect.Message = (*fastReflection_SignDoc)(nil)

type fastReflection_SignDoc SignDoc

func (x *SignDoc) ProtoReflect() protoreflect.Message {
	return (*fastReflection_SignDoc)(x)
}

func (x *SignDoc) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_tx_v1beta1_tx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_SignDoc_messageType fastReflection_SignDoc_messageType
var _ protoreflect.MessageType = fastReflection_SignDoc_messageType{}

type fastReflection_SignDoc_messageType struct{}

func (x fastReflection_SignDoc_messageType) Zero() protoreflect.Message {
	return (*fastReflection_SignDoc)(nil)
}
func (x fastReflection_SignDoc_messageType) New() protoreflect.Message {
	return new(fastReflection_SignDoc)
}
func (x fastReflection_SignDoc_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_SignDoc
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_SignDoc) Descriptor() protoreflect.MessageDescriptor {
	return md_SignDoc
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_SignDoc) Type() protoreflect.MessageType {
	return _fastReflection_SignDoc_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_SignDoc) New() protoreflect.Message {
	return new(fastReflection_SignDoc)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_SignDoc) Interface() protoreflect.ProtoMessage {
	return (*SignDoc)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_SignDoc) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.BodyBytes) != 0 {
		value := protoreflect.ValueOfBytes(x.BodyBytes)
		if !f(fd_SignDoc_body_bytes, value) {
			return
		}
	}
	if len(x.AuthInfoBytes) != 0 {
		value := protoreflect.ValueOfBytes(x.AuthInfoBytes)
		if !f(fd_SignDoc_auth_info_bytes, value) {
			return
		}
	}
	if x.ChainId != "" {
		value := protoreflect.ValueOfString(x.ChainId)
		if !f(fd_SignDoc_chain_id, value) {
			return
		}
	}
	if x.AccountNumber != uint64(0) {
		value := protoreflect.ValueOfUint64(x.AccountNumber)
		if !f(fd_SignDoc_account_number, value) {
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
func (x *fastReflection_SignDoc) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.SignDoc.body_bytes":
		return len(x.BodyBytes) != 0
	case "cosmos.tx.v1beta1.SignDoc.auth_info_bytes":
		return len(x.AuthInfoBytes) != 0
	case "cosmos.tx.v1beta1.SignDoc.chain_id":
		return x.ChainId != ""
	case "cosmos.tx.v1beta1.SignDoc.account_number":
		return x.AccountNumber != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.SignDoc"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.SignDoc does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SignDoc) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.SignDoc.body_bytes":
		x.BodyBytes = nil
	case "cosmos.tx.v1beta1.SignDoc.auth_info_bytes":
		x.AuthInfoBytes = nil
	case "cosmos.tx.v1beta1.SignDoc.chain_id":
		x.ChainId = ""
	case "cosmos.tx.v1beta1.SignDoc.account_number":
		x.AccountNumber = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.SignDoc"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.SignDoc does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_SignDoc) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.tx.v1beta1.SignDoc.body_bytes":
		value := x.BodyBytes
		return protoreflect.ValueOfBytes(value)
	case "cosmos.tx.v1beta1.SignDoc.auth_info_bytes":
		value := x.AuthInfoBytes
		return protoreflect.ValueOfBytes(value)
	case "cosmos.tx.v1beta1.SignDoc.chain_id":
		value := x.ChainId
		return protoreflect.ValueOfString(value)
	case "cosmos.tx.v1beta1.SignDoc.account_number":
		value := x.AccountNumber
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.SignDoc"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.SignDoc does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_SignDoc) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.SignDoc.body_bytes":
		x.BodyBytes = value.Bytes()
	case "cosmos.tx.v1beta1.SignDoc.auth_info_bytes":
		x.AuthInfoBytes = value.Bytes()
	case "cosmos.tx.v1beta1.SignDoc.chain_id":
		x.ChainId = value.Interface().(string)
	case "cosmos.tx.v1beta1.SignDoc.account_number":
		x.AccountNumber = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.SignDoc"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.SignDoc does not contain field %s", fd.FullName()))
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
func (x *fastReflection_SignDoc) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.SignDoc.body_bytes":
		panic(fmt.Errorf("field body_bytes of message cosmos.tx.v1beta1.SignDoc is not mutable"))
	case "cosmos.tx.v1beta1.SignDoc.auth_info_bytes":
		panic(fmt.Errorf("field auth_info_bytes of message cosmos.tx.v1beta1.SignDoc is not mutable"))
	case "cosmos.tx.v1beta1.SignDoc.chain_id":
		panic(fmt.Errorf("field chain_id of message cosmos.tx.v1beta1.SignDoc is not mutable"))
	case "cosmos.tx.v1beta1.SignDoc.account_number":
		panic(fmt.Errorf("field account_number of message cosmos.tx.v1beta1.SignDoc is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.SignDoc"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.SignDoc does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_SignDoc) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.SignDoc.body_bytes":
		return protoreflect.ValueOfBytes(nil)
	case "cosmos.tx.v1beta1.SignDoc.auth_info_bytes":
		return protoreflect.ValueOfBytes(nil)
	case "cosmos.tx.v1beta1.SignDoc.chain_id":
		return protoreflect.ValueOfString("")
	case "cosmos.tx.v1beta1.SignDoc.account_number":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.SignDoc"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.SignDoc does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_SignDoc) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.tx.v1beta1.SignDoc", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_SignDoc) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SignDoc) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_SignDoc) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_SignDoc) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*SignDoc)
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
		l = len(x.BodyBytes)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.AuthInfoBytes)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ChainId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.AccountNumber != 0 {
			n += 1 + runtime.Sov(uint64(x.AccountNumber))
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
		x := input.Message.Interface().(*SignDoc)
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
		if x.AccountNumber != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.AccountNumber))
			i--
			dAtA[i] = 0x20
		}
		if len(x.ChainId) > 0 {
			i -= len(x.ChainId)
			copy(dAtA[i:], x.ChainId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ChainId)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.AuthInfoBytes) > 0 {
			i -= len(x.AuthInfoBytes)
			copy(dAtA[i:], x.AuthInfoBytes)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.AuthInfoBytes)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.BodyBytes) > 0 {
			i -= len(x.BodyBytes)
			copy(dAtA[i:], x.BodyBytes)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BodyBytes)))
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
		x := input.Message.Interface().(*SignDoc)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SignDoc: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SignDoc: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BodyBytes", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.BodyBytes = append(x.BodyBytes[:0], dAtA[iNdEx:postIndex]...)
				if x.BodyBytes == nil {
					x.BodyBytes = []byte{}
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AuthInfoBytes", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.AuthInfoBytes = append(x.AuthInfoBytes[:0], dAtA[iNdEx:postIndex]...)
				if x.AuthInfoBytes == nil {
					x.AuthInfoBytes = []byte{}
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
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
				x.ChainId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AccountNumber", wireType)
				}
				x.AccountNumber = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.AccountNumber |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
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

var (
	md_SignDocDirectAux                protoreflect.MessageDescriptor
	fd_SignDocDirectAux_body_bytes     protoreflect.FieldDescriptor
	fd_SignDocDirectAux_public_key     protoreflect.FieldDescriptor
	fd_SignDocDirectAux_chain_id       protoreflect.FieldDescriptor
	fd_SignDocDirectAux_account_number protoreflect.FieldDescriptor
	fd_SignDocDirectAux_sequence       protoreflect.FieldDescriptor
	fd_SignDocDirectAux_tip            protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_tx_v1beta1_tx_proto_init()
	md_SignDocDirectAux = File_cosmos_tx_v1beta1_tx_proto.Messages().ByName("SignDocDirectAux")
	fd_SignDocDirectAux_body_bytes = md_SignDocDirectAux.Fields().ByName("body_bytes")
	fd_SignDocDirectAux_public_key = md_SignDocDirectAux.Fields().ByName("public_key")
	fd_SignDocDirectAux_chain_id = md_SignDocDirectAux.Fields().ByName("chain_id")
	fd_SignDocDirectAux_account_number = md_SignDocDirectAux.Fields().ByName("account_number")
	fd_SignDocDirectAux_sequence = md_SignDocDirectAux.Fields().ByName("sequence")
	fd_SignDocDirectAux_tip = md_SignDocDirectAux.Fields().ByName("tip")
}

var _ protoreflect.Message = (*fastReflection_SignDocDirectAux)(nil)

type fastReflection_SignDocDirectAux SignDocDirectAux

func (x *SignDocDirectAux) ProtoReflect() protoreflect.Message {
	return (*fastReflection_SignDocDirectAux)(x)
}

func (x *SignDocDirectAux) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_tx_v1beta1_tx_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_SignDocDirectAux_messageType fastReflection_SignDocDirectAux_messageType
var _ protoreflect.MessageType = fastReflection_SignDocDirectAux_messageType{}

type fastReflection_SignDocDirectAux_messageType struct{}

func (x fastReflection_SignDocDirectAux_messageType) Zero() protoreflect.Message {
	return (*fastReflection_SignDocDirectAux)(nil)
}
func (x fastReflection_SignDocDirectAux_messageType) New() protoreflect.Message {
	return new(fastReflection_SignDocDirectAux)
}
func (x fastReflection_SignDocDirectAux_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_SignDocDirectAux
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_SignDocDirectAux) Descriptor() protoreflect.MessageDescriptor {
	return md_SignDocDirectAux
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_SignDocDirectAux) Type() protoreflect.MessageType {
	return _fastReflection_SignDocDirectAux_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_SignDocDirectAux) New() protoreflect.Message {
	return new(fastReflection_SignDocDirectAux)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_SignDocDirectAux) Interface() protoreflect.ProtoMessage {
	return (*SignDocDirectAux)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_SignDocDirectAux) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.BodyBytes) != 0 {
		value := protoreflect.ValueOfBytes(x.BodyBytes)
		if !f(fd_SignDocDirectAux_body_bytes, value) {
			return
		}
	}
	if x.PublicKey != nil {
		value := protoreflect.ValueOfMessage(x.PublicKey.ProtoReflect())
		if !f(fd_SignDocDirectAux_public_key, value) {
			return
		}
	}
	if x.ChainId != "" {
		value := protoreflect.ValueOfString(x.ChainId)
		if !f(fd_SignDocDirectAux_chain_id, value) {
			return
		}
	}
	if x.AccountNumber != uint64(0) {
		value := protoreflect.ValueOfUint64(x.AccountNumber)
		if !f(fd_SignDocDirectAux_account_number, value) {
			return
		}
	}
	if x.Sequence != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Sequence)
		if !f(fd_SignDocDirectAux_sequence, value) {
			return
		}
	}
	if x.Tip != nil {
		value := protoreflect.ValueOfMessage(x.Tip.ProtoReflect())
		if !f(fd_SignDocDirectAux_tip, value) {
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
func (x *fastReflection_SignDocDirectAux) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.SignDocDirectAux.body_bytes":
		return len(x.BodyBytes) != 0
	case "cosmos.tx.v1beta1.SignDocDirectAux.public_key":
		return x.PublicKey != nil
	case "cosmos.tx.v1beta1.SignDocDirectAux.chain_id":
		return x.ChainId != ""
	case "cosmos.tx.v1beta1.SignDocDirectAux.account_number":
		return x.AccountNumber != uint64(0)
	case "cosmos.tx.v1beta1.SignDocDirectAux.sequence":
		return x.Sequence != uint64(0)
	case "cosmos.tx.v1beta1.SignDocDirectAux.tip":
		return x.Tip != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.SignDocDirectAux"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.SignDocDirectAux does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SignDocDirectAux) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.SignDocDirectAux.body_bytes":
		x.BodyBytes = nil
	case "cosmos.tx.v1beta1.SignDocDirectAux.public_key":
		x.PublicKey = nil
	case "cosmos.tx.v1beta1.SignDocDirectAux.chain_id":
		x.ChainId = ""
	case "cosmos.tx.v1beta1.SignDocDirectAux.account_number":
		x.AccountNumber = uint64(0)
	case "cosmos.tx.v1beta1.SignDocDirectAux.sequence":
		x.Sequence = uint64(0)
	case "cosmos.tx.v1beta1.SignDocDirectAux.tip":
		x.Tip = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.SignDocDirectAux"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.SignDocDirectAux does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_SignDocDirectAux) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.tx.v1beta1.SignDocDirectAux.body_bytes":
		value := x.BodyBytes
		return protoreflect.ValueOfBytes(value)
	case "cosmos.tx.v1beta1.SignDocDirectAux.public_key":
		value := x.PublicKey
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.tx.v1beta1.SignDocDirectAux.chain_id":
		value := x.ChainId
		return protoreflect.ValueOfString(value)
	case "cosmos.tx.v1beta1.SignDocDirectAux.account_number":
		value := x.AccountNumber
		return protoreflect.ValueOfUint64(value)
	case "cosmos.tx.v1beta1.SignDocDirectAux.sequence":
		value := x.Sequence
		return protoreflect.ValueOfUint64(value)
	case "cosmos.tx.v1beta1.SignDocDirectAux.tip":
		value := x.Tip
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.SignDocDirectAux"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.SignDocDirectAux does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_SignDocDirectAux) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.SignDocDirectAux.body_bytes":
		x.BodyBytes = value.Bytes()
	case "cosmos.tx.v1beta1.SignDocDirectAux.public_key":
		x.PublicKey = value.Message().Interface().(*anypb.Any)
	case "cosmos.tx.v1beta1.SignDocDirectAux.chain_id":
		x.ChainId = value.Interface().(string)
	case "cosmos.tx.v1beta1.SignDocDirectAux.account_number":
		x.AccountNumber = value.Uint()
	case "cosmos.tx.v1beta1.SignDocDirectAux.sequence":
		x.Sequence = value.Uint()
	case "cosmos.tx.v1beta1.SignDocDirectAux.tip":
		x.Tip = value.Message().Interface().(*Tip)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.SignDocDirectAux"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.SignDocDirectAux does not contain field %s", fd.FullName()))
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
func (x *fastReflection_SignDocDirectAux) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.SignDocDirectAux.public_key":
		if x.PublicKey == nil {
			x.PublicKey = new(anypb.Any)
		}
		return protoreflect.ValueOfMessage(x.PublicKey.ProtoReflect())
	case "cosmos.tx.v1beta1.SignDocDirectAux.tip":
		if x.Tip == nil {
			x.Tip = new(Tip)
		}
		return protoreflect.ValueOfMessage(x.Tip.ProtoReflect())
	case "cosmos.tx.v1beta1.SignDocDirectAux.body_bytes":
		panic(fmt.Errorf("field body_bytes of message cosmos.tx.v1beta1.SignDocDirectAux is not mutable"))
	case "cosmos.tx.v1beta1.SignDocDirectAux.chain_id":
		panic(fmt.Errorf("field chain_id of message cosmos.tx.v1beta1.SignDocDirectAux is not mutable"))
	case "cosmos.tx.v1beta1.SignDocDirectAux.account_number":
		panic(fmt.Errorf("field account_number of message cosmos.tx.v1beta1.SignDocDirectAux is not mutable"))
	case "cosmos.tx.v1beta1.SignDocDirectAux.sequence":
		panic(fmt.Errorf("field sequence of message cosmos.tx.v1beta1.SignDocDirectAux is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.SignDocDirectAux"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.SignDocDirectAux does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_SignDocDirectAux) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.SignDocDirectAux.body_bytes":
		return protoreflect.ValueOfBytes(nil)
	case "cosmos.tx.v1beta1.SignDocDirectAux.public_key":
		m := new(anypb.Any)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.tx.v1beta1.SignDocDirectAux.chain_id":
		return protoreflect.ValueOfString("")
	case "cosmos.tx.v1beta1.SignDocDirectAux.account_number":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.tx.v1beta1.SignDocDirectAux.sequence":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.tx.v1beta1.SignDocDirectAux.tip":
		m := new(Tip)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.SignDocDirectAux"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.SignDocDirectAux does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_SignDocDirectAux) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.tx.v1beta1.SignDocDirectAux", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_SignDocDirectAux) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SignDocDirectAux) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_SignDocDirectAux) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_SignDocDirectAux) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*SignDocDirectAux)
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
		l = len(x.BodyBytes)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.PublicKey != nil {
			l = options.Size(x.PublicKey)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ChainId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.AccountNumber != 0 {
			n += 1 + runtime.Sov(uint64(x.AccountNumber))
		}
		if x.Sequence != 0 {
			n += 1 + runtime.Sov(uint64(x.Sequence))
		}
		if x.Tip != nil {
			l = options.Size(x.Tip)
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
		x := input.Message.Interface().(*SignDocDirectAux)
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
		if x.Tip != nil {
			encoded, err := options.Marshal(x.Tip)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x32
		}
		if x.Sequence != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Sequence))
			i--
			dAtA[i] = 0x28
		}
		if x.AccountNumber != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.AccountNumber))
			i--
			dAtA[i] = 0x20
		}
		if len(x.ChainId) > 0 {
			i -= len(x.ChainId)
			copy(dAtA[i:], x.ChainId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ChainId)))
			i--
			dAtA[i] = 0x1a
		}
		if x.PublicKey != nil {
			encoded, err := options.Marshal(x.PublicKey)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.BodyBytes) > 0 {
			i -= len(x.BodyBytes)
			copy(dAtA[i:], x.BodyBytes)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BodyBytes)))
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
		x := input.Message.Interface().(*SignDocDirectAux)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SignDocDirectAux: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SignDocDirectAux: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BodyBytes", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.BodyBytes = append(x.BodyBytes[:0], dAtA[iNdEx:postIndex]...)
				if x.BodyBytes == nil {
					x.BodyBytes = []byte{}
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.PublicKey == nil {
					x.PublicKey = &anypb.Any{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.PublicKey); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
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
				x.ChainId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AccountNumber", wireType)
				}
				x.AccountNumber = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.AccountNumber |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
				}
				x.Sequence = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Sequence |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Tip", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Tip == nil {
					x.Tip = &Tip{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Tip); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
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

var _ protoreflect.List = (*_TxBody_1_list)(nil)

type _TxBody_1_list struct {
	list *[]*anypb.Any
}

func (x *_TxBody_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TxBody_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_TxBody_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*anypb.Any)
	(*x.list)[i] = concreteValue
}

func (x *_TxBody_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*anypb.Any)
	*x.list = append(*x.list, concreteValue)
}

func (x *_TxBody_1_list) AppendMutable() protoreflect.Value {
	v := new(anypb.Any)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_TxBody_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_TxBody_1_list) NewElement() protoreflect.Value {
	v := new(anypb.Any)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_TxBody_1_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_TxBody_1023_list)(nil)

type _TxBody_1023_list struct {
	list *[]*anypb.Any
}

func (x *_TxBody_1023_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TxBody_1023_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_TxBody_1023_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*anypb.Any)
	(*x.list)[i] = concreteValue
}

func (x *_TxBody_1023_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*anypb.Any)
	*x.list = append(*x.list, concreteValue)
}

func (x *_TxBody_1023_list) AppendMutable() protoreflect.Value {
	v := new(anypb.Any)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_TxBody_1023_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_TxBody_1023_list) NewElement() protoreflect.Value {
	v := new(anypb.Any)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_TxBody_1023_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_TxBody_2047_list)(nil)

type _TxBody_2047_list struct {
	list *[]*anypb.Any
}

func (x *_TxBody_2047_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TxBody_2047_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_TxBody_2047_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*anypb.Any)
	(*x.list)[i] = concreteValue
}

func (x *_TxBody_2047_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*anypb.Any)
	*x.list = append(*x.list, concreteValue)
}

func (x *_TxBody_2047_list) AppendMutable() protoreflect.Value {
	v := new(anypb.Any)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_TxBody_2047_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_TxBody_2047_list) NewElement() protoreflect.Value {
	v := new(anypb.Any)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_TxBody_2047_list) IsValid() bool {
	return x.list != nil
}

var (
	md_TxBody                                protoreflect.MessageDescriptor
	fd_TxBody_messages                       protoreflect.FieldDescriptor
	fd_TxBody_memo                           protoreflect.FieldDescriptor
	fd_TxBody_timeout_height                 protoreflect.FieldDescriptor
	fd_TxBody_extension_options              protoreflect.FieldDescriptor
	fd_TxBody_non_critical_extension_options protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_tx_v1beta1_tx_proto_init()
	md_TxBody = File_cosmos_tx_v1beta1_tx_proto.Messages().ByName("TxBody")
	fd_TxBody_messages = md_TxBody.Fields().ByName("messages")
	fd_TxBody_memo = md_TxBody.Fields().ByName("memo")
	fd_TxBody_timeout_height = md_TxBody.Fields().ByName("timeout_height")
	fd_TxBody_extension_options = md_TxBody.Fields().ByName("extension_options")
	fd_TxBody_non_critical_extension_options = md_TxBody.Fields().ByName("non_critical_extension_options")
}

var _ protoreflect.Message = (*fastReflection_TxBody)(nil)

type fastReflection_TxBody TxBody

func (x *TxBody) ProtoReflect() protoreflect.Message {
	return (*fastReflection_TxBody)(x)
}

func (x *TxBody) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_tx_v1beta1_tx_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_TxBody_messageType fastReflection_TxBody_messageType
var _ protoreflect.MessageType = fastReflection_TxBody_messageType{}

type fastReflection_TxBody_messageType struct{}

func (x fastReflection_TxBody_messageType) Zero() protoreflect.Message {
	return (*fastReflection_TxBody)(nil)
}
func (x fastReflection_TxBody_messageType) New() protoreflect.Message {
	return new(fastReflection_TxBody)
}
func (x fastReflection_TxBody_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_TxBody
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_TxBody) Descriptor() protoreflect.MessageDescriptor {
	return md_TxBody
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_TxBody) Type() protoreflect.MessageType {
	return _fastReflection_TxBody_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_TxBody) New() protoreflect.Message {
	return new(fastReflection_TxBody)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_TxBody) Interface() protoreflect.ProtoMessage {
	return (*TxBody)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_TxBody) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Messages) != 0 {
		value := protoreflect.ValueOfList(&_TxBody_1_list{list: &x.Messages})
		if !f(fd_TxBody_messages, value) {
			return
		}
	}
	if x.Memo != "" {
		value := protoreflect.ValueOfString(x.Memo)
		if !f(fd_TxBody_memo, value) {
			return
		}
	}
	if x.TimeoutHeight != uint64(0) {
		value := protoreflect.ValueOfUint64(x.TimeoutHeight)
		if !f(fd_TxBody_timeout_height, value) {
			return
		}
	}
	if len(x.ExtensionOptions) != 0 {
		value := protoreflect.ValueOfList(&_TxBody_1023_list{list: &x.ExtensionOptions})
		if !f(fd_TxBody_extension_options, value) {
			return
		}
	}
	if len(x.NonCriticalExtensionOptions) != 0 {
		value := protoreflect.ValueOfList(&_TxBody_2047_list{list: &x.NonCriticalExtensionOptions})
		if !f(fd_TxBody_non_critical_extension_options, value) {
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
func (x *fastReflection_TxBody) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.TxBody.messages":
		return len(x.Messages) != 0
	case "cosmos.tx.v1beta1.TxBody.memo":
		return x.Memo != ""
	case "cosmos.tx.v1beta1.TxBody.timeout_height":
		return x.TimeoutHeight != uint64(0)
	case "cosmos.tx.v1beta1.TxBody.extension_options":
		return len(x.ExtensionOptions) != 0
	case "cosmos.tx.v1beta1.TxBody.non_critical_extension_options":
		return len(x.NonCriticalExtensionOptions) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.TxBody"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.TxBody does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TxBody) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.TxBody.messages":
		x.Messages = nil
	case "cosmos.tx.v1beta1.TxBody.memo":
		x.Memo = ""
	case "cosmos.tx.v1beta1.TxBody.timeout_height":
		x.TimeoutHeight = uint64(0)
	case "cosmos.tx.v1beta1.TxBody.extension_options":
		x.ExtensionOptions = nil
	case "cosmos.tx.v1beta1.TxBody.non_critical_extension_options":
		x.NonCriticalExtensionOptions = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.TxBody"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.TxBody does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_TxBody) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.tx.v1beta1.TxBody.messages":
		if len(x.Messages) == 0 {
			return protoreflect.ValueOfList(&_TxBody_1_list{})
		}
		listValue := &_TxBody_1_list{list: &x.Messages}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.tx.v1beta1.TxBody.memo":
		value := x.Memo
		return protoreflect.ValueOfString(value)
	case "cosmos.tx.v1beta1.TxBody.timeout_height":
		value := x.TimeoutHeight
		return protoreflect.ValueOfUint64(value)
	case "cosmos.tx.v1beta1.TxBody.extension_options":
		if len(x.ExtensionOptions) == 0 {
			return protoreflect.ValueOfList(&_TxBody_1023_list{})
		}
		listValue := &_TxBody_1023_list{list: &x.ExtensionOptions}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.tx.v1beta1.TxBody.non_critical_extension_options":
		if len(x.NonCriticalExtensionOptions) == 0 {
			return protoreflect.ValueOfList(&_TxBody_2047_list{})
		}
		listValue := &_TxBody_2047_list{list: &x.NonCriticalExtensionOptions}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.TxBody"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.TxBody does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_TxBody) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.TxBody.messages":
		lv := value.List()
		clv := lv.(*_TxBody_1_list)
		x.Messages = *clv.list
	case "cosmos.tx.v1beta1.TxBody.memo":
		x.Memo = value.Interface().(string)
	case "cosmos.tx.v1beta1.TxBody.timeout_height":
		x.TimeoutHeight = value.Uint()
	case "cosmos.tx.v1beta1.TxBody.extension_options":
		lv := value.List()
		clv := lv.(*_TxBody_1023_list)
		x.ExtensionOptions = *clv.list
	case "cosmos.tx.v1beta1.TxBody.non_critical_extension_options":
		lv := value.List()
		clv := lv.(*_TxBody_2047_list)
		x.NonCriticalExtensionOptions = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.TxBody"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.TxBody does not contain field %s", fd.FullName()))
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
func (x *fastReflection_TxBody) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.TxBody.messages":
		if x.Messages == nil {
			x.Messages = []*anypb.Any{}
		}
		value := &_TxBody_1_list{list: &x.Messages}
		return protoreflect.ValueOfList(value)
	case "cosmos.tx.v1beta1.TxBody.extension_options":
		if x.ExtensionOptions == nil {
			x.ExtensionOptions = []*anypb.Any{}
		}
		value := &_TxBody_1023_list{list: &x.ExtensionOptions}
		return protoreflect.ValueOfList(value)
	case "cosmos.tx.v1beta1.TxBody.non_critical_extension_options":
		if x.NonCriticalExtensionOptions == nil {
			x.NonCriticalExtensionOptions = []*anypb.Any{}
		}
		value := &_TxBody_2047_list{list: &x.NonCriticalExtensionOptions}
		return protoreflect.ValueOfList(value)
	case "cosmos.tx.v1beta1.TxBody.memo":
		panic(fmt.Errorf("field memo of message cosmos.tx.v1beta1.TxBody is not mutable"))
	case "cosmos.tx.v1beta1.TxBody.timeout_height":
		panic(fmt.Errorf("field timeout_height of message cosmos.tx.v1beta1.TxBody is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.TxBody"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.TxBody does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_TxBody) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.TxBody.messages":
		list := []*anypb.Any{}
		return protoreflect.ValueOfList(&_TxBody_1_list{list: &list})
	case "cosmos.tx.v1beta1.TxBody.memo":
		return protoreflect.ValueOfString("")
	case "cosmos.tx.v1beta1.TxBody.timeout_height":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.tx.v1beta1.TxBody.extension_options":
		list := []*anypb.Any{}
		return protoreflect.ValueOfList(&_TxBody_1023_list{list: &list})
	case "cosmos.tx.v1beta1.TxBody.non_critical_extension_options":
		list := []*anypb.Any{}
		return protoreflect.ValueOfList(&_TxBody_2047_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.TxBody"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.TxBody does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_TxBody) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.tx.v1beta1.TxBody", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_TxBody) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TxBody) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_TxBody) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_TxBody) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*TxBody)
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
		if len(x.Messages) > 0 {
			for _, e := range x.Messages {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		l = len(x.Memo)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.TimeoutHeight != 0 {
			n += 1 + runtime.Sov(uint64(x.TimeoutHeight))
		}
		if len(x.ExtensionOptions) > 0 {
			for _, e := range x.ExtensionOptions {
				l = options.Size(e)
				n += 2 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.NonCriticalExtensionOptions) > 0 {
			for _, e := range x.NonCriticalExtensionOptions {
				l = options.Size(e)
				n += 2 + l + runtime.Sov(uint64(l))
			}
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
		x := input.Message.Interface().(*TxBody)
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
		if len(x.NonCriticalExtensionOptions) > 0 {
			for iNdEx := len(x.NonCriticalExtensionOptions) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.NonCriticalExtensionOptions[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x7f
				i--
				dAtA[i] = 0xfa
			}
		}
		if len(x.ExtensionOptions) > 0 {
			for iNdEx := len(x.ExtensionOptions) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.ExtensionOptions[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x3f
				i--
				dAtA[i] = 0xfa
			}
		}
		if x.TimeoutHeight != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.TimeoutHeight))
			i--
			dAtA[i] = 0x18
		}
		if len(x.Memo) > 0 {
			i -= len(x.Memo)
			copy(dAtA[i:], x.Memo)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Memo)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Messages) > 0 {
			for iNdEx := len(x.Messages) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Messages[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
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
		x := input.Message.Interface().(*TxBody)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: TxBody: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: TxBody: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Messages", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Messages = append(x.Messages, &anypb.Any{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Messages[len(x.Messages)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Memo", wireType)
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
				x.Memo = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
				}
				x.TimeoutHeight = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.TimeoutHeight |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 1023:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ExtensionOptions", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ExtensionOptions = append(x.ExtensionOptions, &anypb.Any{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.ExtensionOptions[len(x.ExtensionOptions)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2047:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NonCriticalExtensionOptions", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.NonCriticalExtensionOptions = append(x.NonCriticalExtensionOptions, &anypb.Any{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.NonCriticalExtensionOptions[len(x.NonCriticalExtensionOptions)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
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

var _ protoreflect.List = (*_AuthInfo_1_list)(nil)

type _AuthInfo_1_list struct {
	list *[]*SignerInfo
}

func (x *_AuthInfo_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_AuthInfo_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_AuthInfo_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*SignerInfo)
	(*x.list)[i] = concreteValue
}

func (x *_AuthInfo_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*SignerInfo)
	*x.list = append(*x.list, concreteValue)
}

func (x *_AuthInfo_1_list) AppendMutable() protoreflect.Value {
	v := new(SignerInfo)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_AuthInfo_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_AuthInfo_1_list) NewElement() protoreflect.Value {
	v := new(SignerInfo)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_AuthInfo_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_AuthInfo              protoreflect.MessageDescriptor
	fd_AuthInfo_signer_infos protoreflect.FieldDescriptor
	fd_AuthInfo_fee          protoreflect.FieldDescriptor
	fd_AuthInfo_tip          protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_tx_v1beta1_tx_proto_init()
	md_AuthInfo = File_cosmos_tx_v1beta1_tx_proto.Messages().ByName("AuthInfo")
	fd_AuthInfo_signer_infos = md_AuthInfo.Fields().ByName("signer_infos")
	fd_AuthInfo_fee = md_AuthInfo.Fields().ByName("fee")
	fd_AuthInfo_tip = md_AuthInfo.Fields().ByName("tip")
}

var _ protoreflect.Message = (*fastReflection_AuthInfo)(nil)

type fastReflection_AuthInfo AuthInfo

func (x *AuthInfo) ProtoReflect() protoreflect.Message {
	return (*fastReflection_AuthInfo)(x)
}

func (x *AuthInfo) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_tx_v1beta1_tx_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_AuthInfo_messageType fastReflection_AuthInfo_messageType
var _ protoreflect.MessageType = fastReflection_AuthInfo_messageType{}

type fastReflection_AuthInfo_messageType struct{}

func (x fastReflection_AuthInfo_messageType) Zero() protoreflect.Message {
	return (*fastReflection_AuthInfo)(nil)
}
func (x fastReflection_AuthInfo_messageType) New() protoreflect.Message {
	return new(fastReflection_AuthInfo)
}
func (x fastReflection_AuthInfo_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_AuthInfo
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_AuthInfo) Descriptor() protoreflect.MessageDescriptor {
	return md_AuthInfo
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_AuthInfo) Type() protoreflect.MessageType {
	return _fastReflection_AuthInfo_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_AuthInfo) New() protoreflect.Message {
	return new(fastReflection_AuthInfo)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_AuthInfo) Interface() protoreflect.ProtoMessage {
	return (*AuthInfo)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_AuthInfo) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.SignerInfos) != 0 {
		value := protoreflect.ValueOfList(&_AuthInfo_1_list{list: &x.SignerInfos})
		if !f(fd_AuthInfo_signer_infos, value) {
			return
		}
	}
	if x.Fee != nil {
		value := protoreflect.ValueOfMessage(x.Fee.ProtoReflect())
		if !f(fd_AuthInfo_fee, value) {
			return
		}
	}
	if x.Tip != nil {
		value := protoreflect.ValueOfMessage(x.Tip.ProtoReflect())
		if !f(fd_AuthInfo_tip, value) {
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
func (x *fastReflection_AuthInfo) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.AuthInfo.signer_infos":
		return len(x.SignerInfos) != 0
	case "cosmos.tx.v1beta1.AuthInfo.fee":
		return x.Fee != nil
	case "cosmos.tx.v1beta1.AuthInfo.tip":
		return x.Tip != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.AuthInfo"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.AuthInfo does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AuthInfo) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.AuthInfo.signer_infos":
		x.SignerInfos = nil
	case "cosmos.tx.v1beta1.AuthInfo.fee":
		x.Fee = nil
	case "cosmos.tx.v1beta1.AuthInfo.tip":
		x.Tip = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.AuthInfo"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.AuthInfo does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_AuthInfo) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.tx.v1beta1.AuthInfo.signer_infos":
		if len(x.SignerInfos) == 0 {
			return protoreflect.ValueOfList(&_AuthInfo_1_list{})
		}
		listValue := &_AuthInfo_1_list{list: &x.SignerInfos}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.tx.v1beta1.AuthInfo.fee":
		value := x.Fee
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.tx.v1beta1.AuthInfo.tip":
		value := x.Tip
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.AuthInfo"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.AuthInfo does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_AuthInfo) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.AuthInfo.signer_infos":
		lv := value.List()
		clv := lv.(*_AuthInfo_1_list)
		x.SignerInfos = *clv.list
	case "cosmos.tx.v1beta1.AuthInfo.fee":
		x.Fee = value.Message().Interface().(*Fee)
	case "cosmos.tx.v1beta1.AuthInfo.tip":
		x.Tip = value.Message().Interface().(*Tip)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.AuthInfo"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.AuthInfo does not contain field %s", fd.FullName()))
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
func (x *fastReflection_AuthInfo) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.AuthInfo.signer_infos":
		if x.SignerInfos == nil {
			x.SignerInfos = []*SignerInfo{}
		}
		value := &_AuthInfo_1_list{list: &x.SignerInfos}
		return protoreflect.ValueOfList(value)
	case "cosmos.tx.v1beta1.AuthInfo.fee":
		if x.Fee == nil {
			x.Fee = new(Fee)
		}
		return protoreflect.ValueOfMessage(x.Fee.ProtoReflect())
	case "cosmos.tx.v1beta1.AuthInfo.tip":
		if x.Tip == nil {
			x.Tip = new(Tip)
		}
		return protoreflect.ValueOfMessage(x.Tip.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.AuthInfo"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.AuthInfo does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_AuthInfo) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.AuthInfo.signer_infos":
		list := []*SignerInfo{}
		return protoreflect.ValueOfList(&_AuthInfo_1_list{list: &list})
	case "cosmos.tx.v1beta1.AuthInfo.fee":
		m := new(Fee)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.tx.v1beta1.AuthInfo.tip":
		m := new(Tip)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.AuthInfo"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.AuthInfo does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_AuthInfo) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.tx.v1beta1.AuthInfo", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_AuthInfo) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AuthInfo) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_AuthInfo) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_AuthInfo) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*AuthInfo)
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
		if len(x.SignerInfos) > 0 {
			for _, e := range x.SignerInfos {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.Fee != nil {
			l = options.Size(x.Fee)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Tip != nil {
			l = options.Size(x.Tip)
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
		x := input.Message.Interface().(*AuthInfo)
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
		if x.Tip != nil {
			encoded, err := options.Marshal(x.Tip)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x1a
		}
		if x.Fee != nil {
			encoded, err := options.Marshal(x.Fee)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.SignerInfos) > 0 {
			for iNdEx := len(x.SignerInfos) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.SignerInfos[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
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
		x := input.Message.Interface().(*AuthInfo)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: AuthInfo: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: AuthInfo: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SignerInfos", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.SignerInfos = append(x.SignerInfos, &SignerInfo{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.SignerInfos[len(x.SignerInfos)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Fee == nil {
					x.Fee = &Fee{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Fee); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Tip", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Tip == nil {
					x.Tip = &Tip{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Tip); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
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

var (
	md_SignerInfo            protoreflect.MessageDescriptor
	fd_SignerInfo_public_key protoreflect.FieldDescriptor
	fd_SignerInfo_mode_info  protoreflect.FieldDescriptor
	fd_SignerInfo_sequence   protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_tx_v1beta1_tx_proto_init()
	md_SignerInfo = File_cosmos_tx_v1beta1_tx_proto.Messages().ByName("SignerInfo")
	fd_SignerInfo_public_key = md_SignerInfo.Fields().ByName("public_key")
	fd_SignerInfo_mode_info = md_SignerInfo.Fields().ByName("mode_info")
	fd_SignerInfo_sequence = md_SignerInfo.Fields().ByName("sequence")
}

var _ protoreflect.Message = (*fastReflection_SignerInfo)(nil)

type fastReflection_SignerInfo SignerInfo

func (x *SignerInfo) ProtoReflect() protoreflect.Message {
	return (*fastReflection_SignerInfo)(x)
}

func (x *SignerInfo) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_tx_v1beta1_tx_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_SignerInfo_messageType fastReflection_SignerInfo_messageType
var _ protoreflect.MessageType = fastReflection_SignerInfo_messageType{}

type fastReflection_SignerInfo_messageType struct{}

func (x fastReflection_SignerInfo_messageType) Zero() protoreflect.Message {
	return (*fastReflection_SignerInfo)(nil)
}
func (x fastReflection_SignerInfo_messageType) New() protoreflect.Message {
	return new(fastReflection_SignerInfo)
}
func (x fastReflection_SignerInfo_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_SignerInfo
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_SignerInfo) Descriptor() protoreflect.MessageDescriptor {
	return md_SignerInfo
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_SignerInfo) Type() protoreflect.MessageType {
	return _fastReflection_SignerInfo_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_SignerInfo) New() protoreflect.Message {
	return new(fastReflection_SignerInfo)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_SignerInfo) Interface() protoreflect.ProtoMessage {
	return (*SignerInfo)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_SignerInfo) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.PublicKey != nil {
		value := protoreflect.ValueOfMessage(x.PublicKey.ProtoReflect())
		if !f(fd_SignerInfo_public_key, value) {
			return
		}
	}
	if x.ModeInfo != nil {
		value := protoreflect.ValueOfMessage(x.ModeInfo.ProtoReflect())
		if !f(fd_SignerInfo_mode_info, value) {
			return
		}
	}
	if x.Sequence != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Sequence)
		if !f(fd_SignerInfo_sequence, value) {
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
func (x *fastReflection_SignerInfo) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.SignerInfo.public_key":
		return x.PublicKey != nil
	case "cosmos.tx.v1beta1.SignerInfo.mode_info":
		return x.ModeInfo != nil
	case "cosmos.tx.v1beta1.SignerInfo.sequence":
		return x.Sequence != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.SignerInfo"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.SignerInfo does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SignerInfo) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.SignerInfo.public_key":
		x.PublicKey = nil
	case "cosmos.tx.v1beta1.SignerInfo.mode_info":
		x.ModeInfo = nil
	case "cosmos.tx.v1beta1.SignerInfo.sequence":
		x.Sequence = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.SignerInfo"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.SignerInfo does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_SignerInfo) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.tx.v1beta1.SignerInfo.public_key":
		value := x.PublicKey
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.tx.v1beta1.SignerInfo.mode_info":
		value := x.ModeInfo
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.tx.v1beta1.SignerInfo.sequence":
		value := x.Sequence
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.SignerInfo"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.SignerInfo does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_SignerInfo) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.SignerInfo.public_key":
		x.PublicKey = value.Message().Interface().(*anypb.Any)
	case "cosmos.tx.v1beta1.SignerInfo.mode_info":
		x.ModeInfo = value.Message().Interface().(*ModeInfo)
	case "cosmos.tx.v1beta1.SignerInfo.sequence":
		x.Sequence = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.SignerInfo"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.SignerInfo does not contain field %s", fd.FullName()))
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
func (x *fastReflection_SignerInfo) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.SignerInfo.public_key":
		if x.PublicKey == nil {
			x.PublicKey = new(anypb.Any)
		}
		return protoreflect.ValueOfMessage(x.PublicKey.ProtoReflect())
	case "cosmos.tx.v1beta1.SignerInfo.mode_info":
		if x.ModeInfo == nil {
			x.ModeInfo = new(ModeInfo)
		}
		return protoreflect.ValueOfMessage(x.ModeInfo.ProtoReflect())
	case "cosmos.tx.v1beta1.SignerInfo.sequence":
		panic(fmt.Errorf("field sequence of message cosmos.tx.v1beta1.SignerInfo is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.SignerInfo"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.SignerInfo does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_SignerInfo) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.SignerInfo.public_key":
		m := new(anypb.Any)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.tx.v1beta1.SignerInfo.mode_info":
		m := new(ModeInfo)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.tx.v1beta1.SignerInfo.sequence":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.SignerInfo"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.SignerInfo does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_SignerInfo) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.tx.v1beta1.SignerInfo", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_SignerInfo) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SignerInfo) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_SignerInfo) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_SignerInfo) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*SignerInfo)
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
		if x.PublicKey != nil {
			l = options.Size(x.PublicKey)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.ModeInfo != nil {
			l = options.Size(x.ModeInfo)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Sequence != 0 {
			n += 1 + runtime.Sov(uint64(x.Sequence))
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
		x := input.Message.Interface().(*SignerInfo)
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
		if x.Sequence != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Sequence))
			i--
			dAtA[i] = 0x18
		}
		if x.ModeInfo != nil {
			encoded, err := options.Marshal(x.ModeInfo)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if x.PublicKey != nil {
			encoded, err := options.Marshal(x.PublicKey)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
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
		x := input.Message.Interface().(*SignerInfo)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SignerInfo: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SignerInfo: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.PublicKey == nil {
					x.PublicKey = &anypb.Any{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.PublicKey); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ModeInfo", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.ModeInfo == nil {
					x.ModeInfo = &ModeInfo{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.ModeInfo); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
				}
				x.Sequence = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Sequence |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
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

var (
	md_ModeInfo        protoreflect.MessageDescriptor
	fd_ModeInfo_single protoreflect.FieldDescriptor
	fd_ModeInfo_multi  protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_tx_v1beta1_tx_proto_init()
	md_ModeInfo = File_cosmos_tx_v1beta1_tx_proto.Messages().ByName("ModeInfo")
	fd_ModeInfo_single = md_ModeInfo.Fields().ByName("single")
	fd_ModeInfo_multi = md_ModeInfo.Fields().ByName("multi")
}

var _ protoreflect.Message = (*fastReflection_ModeInfo)(nil)

type fastReflection_ModeInfo ModeInfo

func (x *ModeInfo) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ModeInfo)(x)
}

func (x *ModeInfo) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_tx_v1beta1_tx_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ModeInfo_messageType fastReflection_ModeInfo_messageType
var _ protoreflect.MessageType = fastReflection_ModeInfo_messageType{}

type fastReflection_ModeInfo_messageType struct{}

func (x fastReflection_ModeInfo_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ModeInfo)(nil)
}
func (x fastReflection_ModeInfo_messageType) New() protoreflect.Message {
	return new(fastReflection_ModeInfo)
}
func (x fastReflection_ModeInfo_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ModeInfo
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ModeInfo) Descriptor() protoreflect.MessageDescriptor {
	return md_ModeInfo
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ModeInfo) Type() protoreflect.MessageType {
	return _fastReflection_ModeInfo_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ModeInfo) New() protoreflect.Message {
	return new(fastReflection_ModeInfo)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ModeInfo) Interface() protoreflect.ProtoMessage {
	return (*ModeInfo)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ModeInfo) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Sum != nil {
		switch o := x.Sum.(type) {
		case *ModeInfo_Single_:
			v := o.Single
			value := protoreflect.ValueOfMessage(v.ProtoReflect())
			if !f(fd_ModeInfo_single, value) {
				return
			}
		case *ModeInfo_Multi_:
			v := o.Multi
			value := protoreflect.ValueOfMessage(v.ProtoReflect())
			if !f(fd_ModeInfo_multi, value) {
				return
			}
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
func (x *fastReflection_ModeInfo) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.ModeInfo.single":
		if x.Sum == nil {
			return false
		} else if _, ok := x.Sum.(*ModeInfo_Single_); ok {
			return true
		} else {
			return false
		}
	case "cosmos.tx.v1beta1.ModeInfo.multi":
		if x.Sum == nil {
			return false
		} else if _, ok := x.Sum.(*ModeInfo_Multi_); ok {
			return true
		} else {
			return false
		}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.ModeInfo"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.ModeInfo does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ModeInfo) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.ModeInfo.single":
		x.Sum = nil
	case "cosmos.tx.v1beta1.ModeInfo.multi":
		x.Sum = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.ModeInfo"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.ModeInfo does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ModeInfo) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.tx.v1beta1.ModeInfo.single":
		if x.Sum == nil {
			return protoreflect.ValueOfMessage((*ModeInfo_Single)(nil).ProtoReflect())
		} else if v, ok := x.Sum.(*ModeInfo_Single_); ok {
			return protoreflect.ValueOfMessage(v.Single.ProtoReflect())
		} else {
			return protoreflect.ValueOfMessage((*ModeInfo_Single)(nil).ProtoReflect())
		}
	case "cosmos.tx.v1beta1.ModeInfo.multi":
		if x.Sum == nil {
			return protoreflect.ValueOfMessage((*ModeInfo_Multi)(nil).ProtoReflect())
		} else if v, ok := x.Sum.(*ModeInfo_Multi_); ok {
			return protoreflect.ValueOfMessage(v.Multi.ProtoReflect())
		} else {
			return protoreflect.ValueOfMessage((*ModeInfo_Multi)(nil).ProtoReflect())
		}
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.ModeInfo"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.ModeInfo does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_ModeInfo) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.ModeInfo.single":
		cv := value.Message().Interface().(*ModeInfo_Single)
		x.Sum = &ModeInfo_Single_{Single: cv}
	case "cosmos.tx.v1beta1.ModeInfo.multi":
		cv := value.Message().Interface().(*ModeInfo_Multi)
		x.Sum = &ModeInfo_Multi_{Multi: cv}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.ModeInfo"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.ModeInfo does not contain field %s", fd.FullName()))
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
func (x *fastReflection_ModeInfo) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.ModeInfo.single":
		if x.Sum == nil {
			value := &ModeInfo_Single{}
			oneofValue := &ModeInfo_Single_{Single: value}
			x.Sum = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
		switch m := x.Sum.(type) {
		case *ModeInfo_Single_:
			return protoreflect.ValueOfMessage(m.Single.ProtoReflect())
		default:
			value := &ModeInfo_Single{}
			oneofValue := &ModeInfo_Single_{Single: value}
			x.Sum = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
	case "cosmos.tx.v1beta1.ModeInfo.multi":
		if x.Sum == nil {
			value := &ModeInfo_Multi{}
			oneofValue := &ModeInfo_Multi_{Multi: value}
			x.Sum = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
		switch m := x.Sum.(type) {
		case *ModeInfo_Multi_:
			return protoreflect.ValueOfMessage(m.Multi.ProtoReflect())
		default:
			value := &ModeInfo_Multi{}
			oneofValue := &ModeInfo_Multi_{Multi: value}
			x.Sum = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.ModeInfo"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.ModeInfo does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ModeInfo) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.ModeInfo.single":
		value := &ModeInfo_Single{}
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.tx.v1beta1.ModeInfo.multi":
		value := &ModeInfo_Multi{}
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.ModeInfo"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.ModeInfo does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ModeInfo) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	case "cosmos.tx.v1beta1.ModeInfo.sum":
		if x.Sum == nil {
			return nil
		}
		switch x.Sum.(type) {
		case *ModeInfo_Single_:
			return x.Descriptor().Fields().ByName("single")
		case *ModeInfo_Multi_:
			return x.Descriptor().Fields().ByName("multi")
		}
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.tx.v1beta1.ModeInfo", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ModeInfo) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ModeInfo) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_ModeInfo) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ModeInfo) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ModeInfo)
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
		switch x := x.Sum.(type) {
		case *ModeInfo_Single_:
			if x == nil {
				break
			}
			l = options.Size(x.Single)
			n += 1 + l + runtime.Sov(uint64(l))
		case *ModeInfo_Multi_:
			if x == nil {
				break
			}
			l = options.Size(x.Multi)
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
		x := input.Message.Interface().(*ModeInfo)
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
		switch x := x.Sum.(type) {
		case *ModeInfo_Single_:
			encoded, err := options.Marshal(x.Single)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		case *ModeInfo_Multi_:
			encoded, err := options.Marshal(x.Multi)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
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
		x := input.Message.Interface().(*ModeInfo)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ModeInfo: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ModeInfo: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Single", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				v := &ModeInfo_Single{}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				x.Sum = &ModeInfo_Single_{v}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Multi", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				v := &ModeInfo_Multi{}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				x.Sum = &ModeInfo_Multi_{v}
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

var (
	md_ModeInfo_Single      protoreflect.MessageDescriptor
	fd_ModeInfo_Single_mode protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_tx_v1beta1_tx_proto_init()
	md_ModeInfo_Single = File_cosmos_tx_v1beta1_tx_proto.Messages().ByName("ModeInfo").Messages().ByName("Single")
	fd_ModeInfo_Single_mode = md_ModeInfo_Single.Fields().ByName("mode")
}

var _ protoreflect.Message = (*fastReflection_ModeInfo_Single)(nil)

type fastReflection_ModeInfo_Single ModeInfo_Single

func (x *ModeInfo_Single) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ModeInfo_Single)(x)
}

func (x *ModeInfo_Single) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_tx_v1beta1_tx_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ModeInfo_Single_messageType fastReflection_ModeInfo_Single_messageType
var _ protoreflect.MessageType = fastReflection_ModeInfo_Single_messageType{}

type fastReflection_ModeInfo_Single_messageType struct{}

func (x fastReflection_ModeInfo_Single_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ModeInfo_Single)(nil)
}
func (x fastReflection_ModeInfo_Single_messageType) New() protoreflect.Message {
	return new(fastReflection_ModeInfo_Single)
}
func (x fastReflection_ModeInfo_Single_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ModeInfo_Single
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ModeInfo_Single) Descriptor() protoreflect.MessageDescriptor {
	return md_ModeInfo_Single
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ModeInfo_Single) Type() protoreflect.MessageType {
	return _fastReflection_ModeInfo_Single_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ModeInfo_Single) New() protoreflect.Message {
	return new(fastReflection_ModeInfo_Single)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ModeInfo_Single) Interface() protoreflect.ProtoMessage {
	return (*ModeInfo_Single)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ModeInfo_Single) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Mode != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.Mode))
		if !f(fd_ModeInfo_Single_mode, value) {
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
func (x *fastReflection_ModeInfo_Single) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.ModeInfo.Single.mode":
		return x.Mode != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.ModeInfo.Single"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.ModeInfo.Single does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ModeInfo_Single) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.ModeInfo.Single.mode":
		x.Mode = 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.ModeInfo.Single"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.ModeInfo.Single does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ModeInfo_Single) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.tx.v1beta1.ModeInfo.Single.mode":
		value := x.Mode
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.ModeInfo.Single"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.ModeInfo.Single does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_ModeInfo_Single) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.ModeInfo.Single.mode":
		x.Mode = (v1beta1.SignMode)(value.Enum())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.ModeInfo.Single"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.ModeInfo.Single does not contain field %s", fd.FullName()))
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
func (x *fastReflection_ModeInfo_Single) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.ModeInfo.Single.mode":
		panic(fmt.Errorf("field mode of message cosmos.tx.v1beta1.ModeInfo.Single is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.ModeInfo.Single"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.ModeInfo.Single does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ModeInfo_Single) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.ModeInfo.Single.mode":
		return protoreflect.ValueOfEnum(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.ModeInfo.Single"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.ModeInfo.Single does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ModeInfo_Single) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.tx.v1beta1.ModeInfo.Single", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ModeInfo_Single) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ModeInfo_Single) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_ModeInfo_Single) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ModeInfo_Single) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ModeInfo_Single)
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
		if x.Mode != 0 {
			n += 1 + runtime.Sov(uint64(x.Mode))
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
		x := input.Message.Interface().(*ModeInfo_Single)
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
		if x.Mode != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Mode))
			i--
			dAtA[i] = 0x8
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
		x := input.Message.Interface().(*ModeInfo_Single)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ModeInfo_Single: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ModeInfo_Single: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Mode", wireType)
				}
				x.Mode = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Mode |= v1beta1.SignMode(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
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

var _ protoreflect.List = (*_ModeInfo_Multi_2_list)(nil)

type _ModeInfo_Multi_2_list struct {
	list *[]*ModeInfo
}

func (x *_ModeInfo_Multi_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_ModeInfo_Multi_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_ModeInfo_Multi_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*ModeInfo)
	(*x.list)[i] = concreteValue
}

func (x *_ModeInfo_Multi_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*ModeInfo)
	*x.list = append(*x.list, concreteValue)
}

func (x *_ModeInfo_Multi_2_list) AppendMutable() protoreflect.Value {
	v := new(ModeInfo)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_ModeInfo_Multi_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_ModeInfo_Multi_2_list) NewElement() protoreflect.Value {
	v := new(ModeInfo)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_ModeInfo_Multi_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_ModeInfo_Multi            protoreflect.MessageDescriptor
	fd_ModeInfo_Multi_bitarray   protoreflect.FieldDescriptor
	fd_ModeInfo_Multi_mode_infos protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_tx_v1beta1_tx_proto_init()
	md_ModeInfo_Multi = File_cosmos_tx_v1beta1_tx_proto.Messages().ByName("ModeInfo").Messages().ByName("Multi")
	fd_ModeInfo_Multi_bitarray = md_ModeInfo_Multi.Fields().ByName("bitarray")
	fd_ModeInfo_Multi_mode_infos = md_ModeInfo_Multi.Fields().ByName("mode_infos")
}

var _ protoreflect.Message = (*fastReflection_ModeInfo_Multi)(nil)

type fastReflection_ModeInfo_Multi ModeInfo_Multi

func (x *ModeInfo_Multi) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ModeInfo_Multi)(x)
}

func (x *ModeInfo_Multi) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_tx_v1beta1_tx_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ModeInfo_Multi_messageType fastReflection_ModeInfo_Multi_messageType
var _ protoreflect.MessageType = fastReflection_ModeInfo_Multi_messageType{}

type fastReflection_ModeInfo_Multi_messageType struct{}

func (x fastReflection_ModeInfo_Multi_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ModeInfo_Multi)(nil)
}
func (x fastReflection_ModeInfo_Multi_messageType) New() protoreflect.Message {
	return new(fastReflection_ModeInfo_Multi)
}
func (x fastReflection_ModeInfo_Multi_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ModeInfo_Multi
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ModeInfo_Multi) Descriptor() protoreflect.MessageDescriptor {
	return md_ModeInfo_Multi
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ModeInfo_Multi) Type() protoreflect.MessageType {
	return _fastReflection_ModeInfo_Multi_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ModeInfo_Multi) New() protoreflect.Message {
	return new(fastReflection_ModeInfo_Multi)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ModeInfo_Multi) Interface() protoreflect.ProtoMessage {
	return (*ModeInfo_Multi)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ModeInfo_Multi) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Bitarray != nil {
		value := protoreflect.ValueOfMessage(x.Bitarray.ProtoReflect())
		if !f(fd_ModeInfo_Multi_bitarray, value) {
			return
		}
	}
	if len(x.ModeInfos) != 0 {
		value := protoreflect.ValueOfList(&_ModeInfo_Multi_2_list{list: &x.ModeInfos})
		if !f(fd_ModeInfo_Multi_mode_infos, value) {
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
func (x *fastReflection_ModeInfo_Multi) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.ModeInfo.Multi.bitarray":
		return x.Bitarray != nil
	case "cosmos.tx.v1beta1.ModeInfo.Multi.mode_infos":
		return len(x.ModeInfos) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.ModeInfo.Multi"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.ModeInfo.Multi does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ModeInfo_Multi) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.ModeInfo.Multi.bitarray":
		x.Bitarray = nil
	case "cosmos.tx.v1beta1.ModeInfo.Multi.mode_infos":
		x.ModeInfos = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.ModeInfo.Multi"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.ModeInfo.Multi does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ModeInfo_Multi) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.tx.v1beta1.ModeInfo.Multi.bitarray":
		value := x.Bitarray
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.tx.v1beta1.ModeInfo.Multi.mode_infos":
		if len(x.ModeInfos) == 0 {
			return protoreflect.ValueOfList(&_ModeInfo_Multi_2_list{})
		}
		listValue := &_ModeInfo_Multi_2_list{list: &x.ModeInfos}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.ModeInfo.Multi"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.ModeInfo.Multi does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_ModeInfo_Multi) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.ModeInfo.Multi.bitarray":
		x.Bitarray = value.Message().Interface().(*v1beta11.CompactBitArray)
	case "cosmos.tx.v1beta1.ModeInfo.Multi.mode_infos":
		lv := value.List()
		clv := lv.(*_ModeInfo_Multi_2_list)
		x.ModeInfos = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.ModeInfo.Multi"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.ModeInfo.Multi does not contain field %s", fd.FullName()))
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
func (x *fastReflection_ModeInfo_Multi) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.ModeInfo.Multi.bitarray":
		if x.Bitarray == nil {
			x.Bitarray = new(v1beta11.CompactBitArray)
		}
		return protoreflect.ValueOfMessage(x.Bitarray.ProtoReflect())
	case "cosmos.tx.v1beta1.ModeInfo.Multi.mode_infos":
		if x.ModeInfos == nil {
			x.ModeInfos = []*ModeInfo{}
		}
		value := &_ModeInfo_Multi_2_list{list: &x.ModeInfos}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.ModeInfo.Multi"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.ModeInfo.Multi does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ModeInfo_Multi) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.ModeInfo.Multi.bitarray":
		m := new(v1beta11.CompactBitArray)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.tx.v1beta1.ModeInfo.Multi.mode_infos":
		list := []*ModeInfo{}
		return protoreflect.ValueOfList(&_ModeInfo_Multi_2_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.ModeInfo.Multi"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.ModeInfo.Multi does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ModeInfo_Multi) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.tx.v1beta1.ModeInfo.Multi", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ModeInfo_Multi) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ModeInfo_Multi) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_ModeInfo_Multi) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ModeInfo_Multi) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ModeInfo_Multi)
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
		if x.Bitarray != nil {
			l = options.Size(x.Bitarray)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.ModeInfos) > 0 {
			for _, e := range x.ModeInfos {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
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
		x := input.Message.Interface().(*ModeInfo_Multi)
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
		if len(x.ModeInfos) > 0 {
			for iNdEx := len(x.ModeInfos) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.ModeInfos[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x12
			}
		}
		if x.Bitarray != nil {
			encoded, err := options.Marshal(x.Bitarray)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
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
		x := input.Message.Interface().(*ModeInfo_Multi)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ModeInfo_Multi: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ModeInfo_Multi: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Bitarray", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Bitarray == nil {
					x.Bitarray = &v1beta11.CompactBitArray{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Bitarray); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ModeInfos", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ModeInfos = append(x.ModeInfos, &ModeInfo{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.ModeInfos[len(x.ModeInfos)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
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

var _ protoreflect.List = (*_Fee_1_list)(nil)

type _Fee_1_list struct {
	list *[]*v1beta12.Coin
}

func (x *_Fee_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_Fee_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_Fee_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta12.Coin)
	(*x.list)[i] = concreteValue
}

func (x *_Fee_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta12.Coin)
	*x.list = append(*x.list, concreteValue)
}

func (x *_Fee_1_list) AppendMutable() protoreflect.Value {
	v := new(v1beta12.Coin)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Fee_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_Fee_1_list) NewElement() protoreflect.Value {
	v := new(v1beta12.Coin)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Fee_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_Fee           protoreflect.MessageDescriptor
	fd_Fee_amount    protoreflect.FieldDescriptor
	fd_Fee_gas_limit protoreflect.FieldDescriptor
	fd_Fee_payer     protoreflect.FieldDescriptor
	fd_Fee_granter   protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_tx_v1beta1_tx_proto_init()
	md_Fee = File_cosmos_tx_v1beta1_tx_proto.Messages().ByName("Fee")
	fd_Fee_amount = md_Fee.Fields().ByName("amount")
	fd_Fee_gas_limit = md_Fee.Fields().ByName("gas_limit")
	fd_Fee_payer = md_Fee.Fields().ByName("payer")
	fd_Fee_granter = md_Fee.Fields().ByName("granter")
}

var _ protoreflect.Message = (*fastReflection_Fee)(nil)

type fastReflection_Fee Fee

func (x *Fee) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Fee)(x)
}

func (x *Fee) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_tx_v1beta1_tx_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Fee_messageType fastReflection_Fee_messageType
var _ protoreflect.MessageType = fastReflection_Fee_messageType{}

type fastReflection_Fee_messageType struct{}

func (x fastReflection_Fee_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Fee)(nil)
}
func (x fastReflection_Fee_messageType) New() protoreflect.Message {
	return new(fastReflection_Fee)
}
func (x fastReflection_Fee_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Fee
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Fee) Descriptor() protoreflect.MessageDescriptor {
	return md_Fee
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Fee) Type() protoreflect.MessageType {
	return _fastReflection_Fee_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Fee) New() protoreflect.Message {
	return new(fastReflection_Fee)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Fee) Interface() protoreflect.ProtoMessage {
	return (*Fee)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Fee) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Amount) != 0 {
		value := protoreflect.ValueOfList(&_Fee_1_list{list: &x.Amount})
		if !f(fd_Fee_amount, value) {
			return
		}
	}
	if x.GasLimit != uint64(0) {
		value := protoreflect.ValueOfUint64(x.GasLimit)
		if !f(fd_Fee_gas_limit, value) {
			return
		}
	}
	if x.Payer != "" {
		value := protoreflect.ValueOfString(x.Payer)
		if !f(fd_Fee_payer, value) {
			return
		}
	}
	if x.Granter != "" {
		value := protoreflect.ValueOfString(x.Granter)
		if !f(fd_Fee_granter, value) {
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
func (x *fastReflection_Fee) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.Fee.amount":
		return len(x.Amount) != 0
	case "cosmos.tx.v1beta1.Fee.gas_limit":
		return x.GasLimit != uint64(0)
	case "cosmos.tx.v1beta1.Fee.payer":
		return x.Payer != ""
	case "cosmos.tx.v1beta1.Fee.granter":
		return x.Granter != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.Fee"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.Fee does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Fee) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.Fee.amount":
		x.Amount = nil
	case "cosmos.tx.v1beta1.Fee.gas_limit":
		x.GasLimit = uint64(0)
	case "cosmos.tx.v1beta1.Fee.payer":
		x.Payer = ""
	case "cosmos.tx.v1beta1.Fee.granter":
		x.Granter = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.Fee"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.Fee does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Fee) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.tx.v1beta1.Fee.amount":
		if len(x.Amount) == 0 {
			return protoreflect.ValueOfList(&_Fee_1_list{})
		}
		listValue := &_Fee_1_list{list: &x.Amount}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.tx.v1beta1.Fee.gas_limit":
		value := x.GasLimit
		return protoreflect.ValueOfUint64(value)
	case "cosmos.tx.v1beta1.Fee.payer":
		value := x.Payer
		return protoreflect.ValueOfString(value)
	case "cosmos.tx.v1beta1.Fee.granter":
		value := x.Granter
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.Fee"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.Fee does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Fee) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.Fee.amount":
		lv := value.List()
		clv := lv.(*_Fee_1_list)
		x.Amount = *clv.list
	case "cosmos.tx.v1beta1.Fee.gas_limit":
		x.GasLimit = value.Uint()
	case "cosmos.tx.v1beta1.Fee.payer":
		x.Payer = value.Interface().(string)
	case "cosmos.tx.v1beta1.Fee.granter":
		x.Granter = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.Fee"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.Fee does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Fee) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.Fee.amount":
		if x.Amount == nil {
			x.Amount = []*v1beta12.Coin{}
		}
		value := &_Fee_1_list{list: &x.Amount}
		return protoreflect.ValueOfList(value)
	case "cosmos.tx.v1beta1.Fee.gas_limit":
		panic(fmt.Errorf("field gas_limit of message cosmos.tx.v1beta1.Fee is not mutable"))
	case "cosmos.tx.v1beta1.Fee.payer":
		panic(fmt.Errorf("field payer of message cosmos.tx.v1beta1.Fee is not mutable"))
	case "cosmos.tx.v1beta1.Fee.granter":
		panic(fmt.Errorf("field granter of message cosmos.tx.v1beta1.Fee is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.Fee"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.Fee does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Fee) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.Fee.amount":
		list := []*v1beta12.Coin{}
		return protoreflect.ValueOfList(&_Fee_1_list{list: &list})
	case "cosmos.tx.v1beta1.Fee.gas_limit":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.tx.v1beta1.Fee.payer":
		return protoreflect.ValueOfString("")
	case "cosmos.tx.v1beta1.Fee.granter":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.Fee"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.Fee does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Fee) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.tx.v1beta1.Fee", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Fee) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Fee) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Fee) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Fee) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Fee)
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
		if len(x.Amount) > 0 {
			for _, e := range x.Amount {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.GasLimit != 0 {
			n += 1 + runtime.Sov(uint64(x.GasLimit))
		}
		l = len(x.Payer)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Granter)
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
		x := input.Message.Interface().(*Fee)
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
		if len(x.Granter) > 0 {
			i -= len(x.Granter)
			copy(dAtA[i:], x.Granter)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Granter)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Payer) > 0 {
			i -= len(x.Payer)
			copy(dAtA[i:], x.Payer)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Payer)))
			i--
			dAtA[i] = 0x1a
		}
		if x.GasLimit != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.GasLimit))
			i--
			dAtA[i] = 0x10
		}
		if len(x.Amount) > 0 {
			for iNdEx := len(x.Amount) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Amount[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
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
		x := input.Message.Interface().(*Fee)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Fee: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Fee: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Amount = append(x.Amount, &v1beta12.Coin{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Amount[len(x.Amount)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GasLimit", wireType)
				}
				x.GasLimit = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.GasLimit |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Payer", wireType)
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
				x.Payer = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
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
				x.Granter = string(dAtA[iNdEx:postIndex])
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

var _ protoreflect.List = (*_Tip_1_list)(nil)

type _Tip_1_list struct {
	list *[]*v1beta12.Coin
}

func (x *_Tip_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_Tip_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_Tip_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta12.Coin)
	(*x.list)[i] = concreteValue
}

func (x *_Tip_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta12.Coin)
	*x.list = append(*x.list, concreteValue)
}

func (x *_Tip_1_list) AppendMutable() protoreflect.Value {
	v := new(v1beta12.Coin)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Tip_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_Tip_1_list) NewElement() protoreflect.Value {
	v := new(v1beta12.Coin)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Tip_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_Tip        protoreflect.MessageDescriptor
	fd_Tip_amount protoreflect.FieldDescriptor
	fd_Tip_tipper protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_tx_v1beta1_tx_proto_init()
	md_Tip = File_cosmos_tx_v1beta1_tx_proto.Messages().ByName("Tip")
	fd_Tip_amount = md_Tip.Fields().ByName("amount")
	fd_Tip_tipper = md_Tip.Fields().ByName("tipper")
}

var _ protoreflect.Message = (*fastReflection_Tip)(nil)

type fastReflection_Tip Tip

func (x *Tip) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Tip)(x)
}

func (x *Tip) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_tx_v1beta1_tx_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Tip_messageType fastReflection_Tip_messageType
var _ protoreflect.MessageType = fastReflection_Tip_messageType{}

type fastReflection_Tip_messageType struct{}

func (x fastReflection_Tip_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Tip)(nil)
}
func (x fastReflection_Tip_messageType) New() protoreflect.Message {
	return new(fastReflection_Tip)
}
func (x fastReflection_Tip_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Tip
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Tip) Descriptor() protoreflect.MessageDescriptor {
	return md_Tip
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Tip) Type() protoreflect.MessageType {
	return _fastReflection_Tip_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Tip) New() protoreflect.Message {
	return new(fastReflection_Tip)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Tip) Interface() protoreflect.ProtoMessage {
	return (*Tip)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Tip) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Amount) != 0 {
		value := protoreflect.ValueOfList(&_Tip_1_list{list: &x.Amount})
		if !f(fd_Tip_amount, value) {
			return
		}
	}
	if x.Tipper != "" {
		value := protoreflect.ValueOfString(x.Tipper)
		if !f(fd_Tip_tipper, value) {
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
func (x *fastReflection_Tip) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.Tip.amount":
		return len(x.Amount) != 0
	case "cosmos.tx.v1beta1.Tip.tipper":
		return x.Tipper != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.Tip"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.Tip does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Tip) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.Tip.amount":
		x.Amount = nil
	case "cosmos.tx.v1beta1.Tip.tipper":
		x.Tipper = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.Tip"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.Tip does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Tip) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.tx.v1beta1.Tip.amount":
		if len(x.Amount) == 0 {
			return protoreflect.ValueOfList(&_Tip_1_list{})
		}
		listValue := &_Tip_1_list{list: &x.Amount}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.tx.v1beta1.Tip.tipper":
		value := x.Tipper
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.Tip"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.Tip does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Tip) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.Tip.amount":
		lv := value.List()
		clv := lv.(*_Tip_1_list)
		x.Amount = *clv.list
	case "cosmos.tx.v1beta1.Tip.tipper":
		x.Tipper = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.Tip"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.Tip does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Tip) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.Tip.amount":
		if x.Amount == nil {
			x.Amount = []*v1beta12.Coin{}
		}
		value := &_Tip_1_list{list: &x.Amount}
		return protoreflect.ValueOfList(value)
	case "cosmos.tx.v1beta1.Tip.tipper":
		panic(fmt.Errorf("field tipper of message cosmos.tx.v1beta1.Tip is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.Tip"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.Tip does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Tip) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.Tip.amount":
		list := []*v1beta12.Coin{}
		return protoreflect.ValueOfList(&_Tip_1_list{list: &list})
	case "cosmos.tx.v1beta1.Tip.tipper":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.Tip"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.Tip does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Tip) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.tx.v1beta1.Tip", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Tip) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Tip) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Tip) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Tip) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Tip)
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
		if len(x.Amount) > 0 {
			for _, e := range x.Amount {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		l = len(x.Tipper)
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
		x := input.Message.Interface().(*Tip)
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
		if len(x.Tipper) > 0 {
			i -= len(x.Tipper)
			copy(dAtA[i:], x.Tipper)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Tipper)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Amount) > 0 {
			for iNdEx := len(x.Amount) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Amount[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
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
		x := input.Message.Interface().(*Tip)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Tip: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Tip: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Amount = append(x.Amount, &v1beta12.Coin{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Amount[len(x.Amount)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Tipper", wireType)
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
				x.Tipper = string(dAtA[iNdEx:postIndex])
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

var (
	md_AuxSignerData          protoreflect.MessageDescriptor
	fd_AuxSignerData_address  protoreflect.FieldDescriptor
	fd_AuxSignerData_sign_doc protoreflect.FieldDescriptor
	fd_AuxSignerData_mode     protoreflect.FieldDescriptor
	fd_AuxSignerData_sig      protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_tx_v1beta1_tx_proto_init()
	md_AuxSignerData = File_cosmos_tx_v1beta1_tx_proto.Messages().ByName("AuxSignerData")
	fd_AuxSignerData_address = md_AuxSignerData.Fields().ByName("address")
	fd_AuxSignerData_sign_doc = md_AuxSignerData.Fields().ByName("sign_doc")
	fd_AuxSignerData_mode = md_AuxSignerData.Fields().ByName("mode")
	fd_AuxSignerData_sig = md_AuxSignerData.Fields().ByName("sig")
}

var _ protoreflect.Message = (*fastReflection_AuxSignerData)(nil)

type fastReflection_AuxSignerData AuxSignerData

func (x *AuxSignerData) ProtoReflect() protoreflect.Message {
	return (*fastReflection_AuxSignerData)(x)
}

func (x *AuxSignerData) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_tx_v1beta1_tx_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_AuxSignerData_messageType fastReflection_AuxSignerData_messageType
var _ protoreflect.MessageType = fastReflection_AuxSignerData_messageType{}

type fastReflection_AuxSignerData_messageType struct{}

func (x fastReflection_AuxSignerData_messageType) Zero() protoreflect.Message {
	return (*fastReflection_AuxSignerData)(nil)
}
func (x fastReflection_AuxSignerData_messageType) New() protoreflect.Message {
	return new(fastReflection_AuxSignerData)
}
func (x fastReflection_AuxSignerData_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_AuxSignerData
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_AuxSignerData) Descriptor() protoreflect.MessageDescriptor {
	return md_AuxSignerData
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_AuxSignerData) Type() protoreflect.MessageType {
	return _fastReflection_AuxSignerData_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_AuxSignerData) New() protoreflect.Message {
	return new(fastReflection_AuxSignerData)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_AuxSignerData) Interface() protoreflect.ProtoMessage {
	return (*AuxSignerData)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_AuxSignerData) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Address != "" {
		value := protoreflect.ValueOfString(x.Address)
		if !f(fd_AuxSignerData_address, value) {
			return
		}
	}
	if x.SignDoc != nil {
		value := protoreflect.ValueOfMessage(x.SignDoc.ProtoReflect())
		if !f(fd_AuxSignerData_sign_doc, value) {
			return
		}
	}
	if x.Mode != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.Mode))
		if !f(fd_AuxSignerData_mode, value) {
			return
		}
	}
	if len(x.Sig) != 0 {
		value := protoreflect.ValueOfBytes(x.Sig)
		if !f(fd_AuxSignerData_sig, value) {
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
func (x *fastReflection_AuxSignerData) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.AuxSignerData.address":
		return x.Address != ""
	case "cosmos.tx.v1beta1.AuxSignerData.sign_doc":
		return x.SignDoc != nil
	case "cosmos.tx.v1beta1.AuxSignerData.mode":
		return x.Mode != 0
	case "cosmos.tx.v1beta1.AuxSignerData.sig":
		return len(x.Sig) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.AuxSignerData"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.AuxSignerData does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AuxSignerData) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.AuxSignerData.address":
		x.Address = ""
	case "cosmos.tx.v1beta1.AuxSignerData.sign_doc":
		x.SignDoc = nil
	case "cosmos.tx.v1beta1.AuxSignerData.mode":
		x.Mode = 0
	case "cosmos.tx.v1beta1.AuxSignerData.sig":
		x.Sig = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.AuxSignerData"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.AuxSignerData does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_AuxSignerData) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.tx.v1beta1.AuxSignerData.address":
		value := x.Address
		return protoreflect.ValueOfString(value)
	case "cosmos.tx.v1beta1.AuxSignerData.sign_doc":
		value := x.SignDoc
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.tx.v1beta1.AuxSignerData.mode":
		value := x.Mode
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	case "cosmos.tx.v1beta1.AuxSignerData.sig":
		value := x.Sig
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.AuxSignerData"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.AuxSignerData does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_AuxSignerData) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.AuxSignerData.address":
		x.Address = value.Interface().(string)
	case "cosmos.tx.v1beta1.AuxSignerData.sign_doc":
		x.SignDoc = value.Message().Interface().(*SignDocDirectAux)
	case "cosmos.tx.v1beta1.AuxSignerData.mode":
		x.Mode = (v1beta1.SignMode)(value.Enum())
	case "cosmos.tx.v1beta1.AuxSignerData.sig":
		x.Sig = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.AuxSignerData"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.AuxSignerData does not contain field %s", fd.FullName()))
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
func (x *fastReflection_AuxSignerData) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.AuxSignerData.sign_doc":
		if x.SignDoc == nil {
			x.SignDoc = new(SignDocDirectAux)
		}
		return protoreflect.ValueOfMessage(x.SignDoc.ProtoReflect())
	case "cosmos.tx.v1beta1.AuxSignerData.address":
		panic(fmt.Errorf("field address of message cosmos.tx.v1beta1.AuxSignerData is not mutable"))
	case "cosmos.tx.v1beta1.AuxSignerData.mode":
		panic(fmt.Errorf("field mode of message cosmos.tx.v1beta1.AuxSignerData is not mutable"))
	case "cosmos.tx.v1beta1.AuxSignerData.sig":
		panic(fmt.Errorf("field sig of message cosmos.tx.v1beta1.AuxSignerData is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.AuxSignerData"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.AuxSignerData does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_AuxSignerData) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.AuxSignerData.address":
		return protoreflect.ValueOfString("")
	case "cosmos.tx.v1beta1.AuxSignerData.sign_doc":
		m := new(SignDocDirectAux)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.tx.v1beta1.AuxSignerData.mode":
		return protoreflect.ValueOfEnum(0)
	case "cosmos.tx.v1beta1.AuxSignerData.sig":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.AuxSignerData"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.AuxSignerData does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_AuxSignerData) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.tx.v1beta1.AuxSignerData", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_AuxSignerData) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AuxSignerData) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_AuxSignerData) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_AuxSignerData) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*AuxSignerData)
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
		l = len(x.Address)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.SignDoc != nil {
			l = options.Size(x.SignDoc)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Mode != 0 {
			n += 1 + runtime.Sov(uint64(x.Mode))
		}
		l = len(x.Sig)
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
		x := input.Message.Interface().(*AuxSignerData)
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
		if len(x.Sig) > 0 {
			i -= len(x.Sig)
			copy(dAtA[i:], x.Sig)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Sig)))
			i--
			dAtA[i] = 0x22
		}
		if x.Mode != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Mode))
			i--
			dAtA[i] = 0x18
		}
		if x.SignDoc != nil {
			encoded, err := options.Marshal(x.SignDoc)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Address) > 0 {
			i -= len(x.Address)
			copy(dAtA[i:], x.Address)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Address)))
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
		x := input.Message.Interface().(*AuxSignerData)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: AuxSignerData: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: AuxSignerData: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
				x.Address = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SignDoc", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.SignDoc == nil {
					x.SignDoc = &SignDocDirectAux{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.SignDoc); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Mode", wireType)
				}
				x.Mode = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Mode |= v1beta1.SignMode(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Sig", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Sig = append(x.Sig[:0], dAtA[iNdEx:postIndex]...)
				if x.Sig == nil {
					x.Sig = []byte{}
				}
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
// source: cosmos/tx/v1beta1/tx.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Tx is the standard type used for broadcasting transactions.
type Tx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// body is the processable content of the transaction
	Body *TxBody `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// auth_info is the authorization related content of the transaction,
	// specifically signers, signer modes and fee
	AuthInfo *AuthInfo `protobuf:"bytes,2,opt,name=auth_info,json=authInfo,proto3" json:"auth_info,omitempty"`
	// signatures is a list of signatures that matches the length and order of
	// AuthInfo's signer_infos to allow connecting signature meta information like
	// public key and signing mode by position.
	Signatures [][]byte `protobuf:"bytes,3,rep,name=signatures,proto3" json:"signatures,omitempty"`
}

func (x *Tx) Reset() {
	*x = Tx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_tx_v1beta1_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tx) ProtoMessage() {}

// Deprecated: Use Tx.ProtoReflect.Descriptor instead.
func (*Tx) Descriptor() ([]byte, []int) {
	return file_cosmos_tx_v1beta1_tx_proto_rawDescGZIP(), []int{0}
}

func (x *Tx) GetBody() *TxBody {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *Tx) GetAuthInfo() *AuthInfo {
	if x != nil {
		return x.AuthInfo
	}
	return nil
}

func (x *Tx) GetSignatures() [][]byte {
	if x != nil {
		return x.Signatures
	}
	return nil
}

// TxRaw is a variant of Tx that pins the signer's exact binary representation
// of body and auth_info. This is used for signing, broadcasting and
// verification. The binary `serialize(tx: TxRaw)` is stored in Tendermint and
// the hash `sha256(serialize(tx: TxRaw))` becomes the "txhash", commonly used
// as the transaction ID.
type TxRaw struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// body_bytes is a protobuf serialization of a TxBody that matches the
	// representation in SignDoc.
	BodyBytes []byte `protobuf:"bytes,1,opt,name=body_bytes,json=bodyBytes,proto3" json:"body_bytes,omitempty"`
	// auth_info_bytes is a protobuf serialization of an AuthInfo that matches the
	// representation in SignDoc.
	AuthInfoBytes []byte `protobuf:"bytes,2,opt,name=auth_info_bytes,json=authInfoBytes,proto3" json:"auth_info_bytes,omitempty"`
	// signatures is a list of signatures that matches the length and order of
	// AuthInfo's signer_infos to allow connecting signature meta information like
	// public key and signing mode by position.
	Signatures [][]byte `protobuf:"bytes,3,rep,name=signatures,proto3" json:"signatures,omitempty"`
}

func (x *TxRaw) Reset() {
	*x = TxRaw{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_tx_v1beta1_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxRaw) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxRaw) ProtoMessage() {}

// Deprecated: Use TxRaw.ProtoReflect.Descriptor instead.
func (*TxRaw) Descriptor() ([]byte, []int) {
	return file_cosmos_tx_v1beta1_tx_proto_rawDescGZIP(), []int{1}
}

func (x *TxRaw) GetBodyBytes() []byte {
	if x != nil {
		return x.BodyBytes
	}
	return nil
}

func (x *TxRaw) GetAuthInfoBytes() []byte {
	if x != nil {
		return x.AuthInfoBytes
	}
	return nil
}

func (x *TxRaw) GetSignatures() [][]byte {
	if x != nil {
		return x.Signatures
	}
	return nil
}

// SignDoc is the type used for generating sign bytes for SIGN_MODE_DIRECT.
type SignDoc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// body_bytes is protobuf serialization of a TxBody that matches the
	// representation in TxRaw.
	BodyBytes []byte `protobuf:"bytes,1,opt,name=body_bytes,json=bodyBytes,proto3" json:"body_bytes,omitempty"`
	// auth_info_bytes is a protobuf serialization of an AuthInfo that matches the
	// representation in TxRaw.
	AuthInfoBytes []byte `protobuf:"bytes,2,opt,name=auth_info_bytes,json=authInfoBytes,proto3" json:"auth_info_bytes,omitempty"`
	// chain_id is the unique identifier of the chain this transaction targets.
	// It prevents signed transactions from being used on another chain by an
	// attacker
	ChainId string `protobuf:"bytes,3,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// account_number is the account number of the account in state
	AccountNumber uint64 `protobuf:"varint,4,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
}

func (x *SignDoc) Reset() {
	*x = SignDoc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_tx_v1beta1_tx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignDoc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignDoc) ProtoMessage() {}

// Deprecated: Use SignDoc.ProtoReflect.Descriptor instead.
func (*SignDoc) Descriptor() ([]byte, []int) {
	return file_cosmos_tx_v1beta1_tx_proto_rawDescGZIP(), []int{2}
}

func (x *SignDoc) GetBodyBytes() []byte {
	if x != nil {
		return x.BodyBytes
	}
	return nil
}

func (x *SignDoc) GetAuthInfoBytes() []byte {
	if x != nil {
		return x.AuthInfoBytes
	}
	return nil
}

func (x *SignDoc) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

func (x *SignDoc) GetAccountNumber() uint64 {
	if x != nil {
		return x.AccountNumber
	}
	return 0
}

// SignDocDirectAux is the type used for generating sign bytes for
// SIGN_MODE_DIRECT_AUX.
type SignDocDirectAux struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// body_bytes is protobuf serialization of a TxBody that matches the
	// representation in TxRaw.
	BodyBytes []byte `protobuf:"bytes,1,opt,name=body_bytes,json=bodyBytes,proto3" json:"body_bytes,omitempty"`
	// public_key is the public key of the signing account.
	PublicKey *anypb.Any `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// chain_id is the identifier of the chain this transaction targets.
	// It prevents signed transactions from being used on another chain by an
	// attacker.
	ChainId string `protobuf:"bytes,3,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// account_number is the account number of the account in state.
	AccountNumber uint64 `protobuf:"varint,4,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	// sequence is the sequence number of the signing account.
	Sequence uint64 `protobuf:"varint,5,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// Tip is the optional tip used for meta-transactions. It should be left
	// empty if the signer is not the tipper for this transaction.
	Tip *Tip `protobuf:"bytes,6,opt,name=tip,proto3" json:"tip,omitempty"`
}

func (x *SignDocDirectAux) Reset() {
	*x = SignDocDirectAux{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_tx_v1beta1_tx_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignDocDirectAux) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignDocDirectAux) ProtoMessage() {}

// Deprecated: Use SignDocDirectAux.ProtoReflect.Descriptor instead.
func (*SignDocDirectAux) Descriptor() ([]byte, []int) {
	return file_cosmos_tx_v1beta1_tx_proto_rawDescGZIP(), []int{3}
}

func (x *SignDocDirectAux) GetBodyBytes() []byte {
	if x != nil {
		return x.BodyBytes
	}
	return nil
}

func (x *SignDocDirectAux) GetPublicKey() *anypb.Any {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *SignDocDirectAux) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

func (x *SignDocDirectAux) GetAccountNumber() uint64 {
	if x != nil {
		return x.AccountNumber
	}
	return 0
}

func (x *SignDocDirectAux) GetSequence() uint64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *SignDocDirectAux) GetTip() *Tip {
	if x != nil {
		return x.Tip
	}
	return nil
}

// TxBody is the body of a transaction that all signers sign over.
type TxBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// messages is a list of messages to be executed. The required signers of
	// those messages define the number and order of elements in AuthInfo's
	// signer_infos and Tx's signatures. Each required signer address is added to
	// the list only the first time it occurs.
	// By convention, the first required signer (usually from the first message)
	// is referred to as the primary signer and pays the fee for the whole
	// transaction.
	Messages []*anypb.Any `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	// memo is any arbitrary note/comment to be added to the transaction.
	// WARNING: in clients, any publicly exposed text should not be called memo,
	// but should be called `note` instead (see https://github.com/cosmos/cosmos-sdk/issues/9122).
	Memo string `protobuf:"bytes,2,opt,name=memo,proto3" json:"memo,omitempty"`
	// timeout is the block height after which this transaction will not
	// be processed by the chain
	TimeoutHeight uint64 `protobuf:"varint,3,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height,omitempty"`
	// extension_options are arbitrary options that can be added by chains
	// when the default options are not sufficient. If any of these are present
	// and can't be handled, the transaction will be rejected
	ExtensionOptions []*anypb.Any `protobuf:"bytes,1023,rep,name=extension_options,json=extensionOptions,proto3" json:"extension_options,omitempty"`
	// extension_options are arbitrary options that can be added by chains
	// when the default options are not sufficient. If any of these are present
	// and can't be handled, they will be ignored
	NonCriticalExtensionOptions []*anypb.Any `protobuf:"bytes,2047,rep,name=non_critical_extension_options,json=nonCriticalExtensionOptions,proto3" json:"non_critical_extension_options,omitempty"`
}

func (x *TxBody) Reset() {
	*x = TxBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_tx_v1beta1_tx_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxBody) ProtoMessage() {}

// Deprecated: Use TxBody.ProtoReflect.Descriptor instead.
func (*TxBody) Descriptor() ([]byte, []int) {
	return file_cosmos_tx_v1beta1_tx_proto_rawDescGZIP(), []int{4}
}

func (x *TxBody) GetMessages() []*anypb.Any {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *TxBody) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *TxBody) GetTimeoutHeight() uint64 {
	if x != nil {
		return x.TimeoutHeight
	}
	return 0
}

func (x *TxBody) GetExtensionOptions() []*anypb.Any {
	if x != nil {
		return x.ExtensionOptions
	}
	return nil
}

func (x *TxBody) GetNonCriticalExtensionOptions() []*anypb.Any {
	if x != nil {
		return x.NonCriticalExtensionOptions
	}
	return nil
}

// AuthInfo describes the fee and signer modes that are used to sign a
// transaction.
type AuthInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// signer_infos defines the signing modes for the required signers. The number
	// and order of elements must match the required signers from TxBody's
	// messages. The first element is the primary signer and the one which pays
	// the fee.
	SignerInfos []*SignerInfo `protobuf:"bytes,1,rep,name=signer_infos,json=signerInfos,proto3" json:"signer_infos,omitempty"`
	// Fee is the fee and gas limit for the transaction. The first signer is the
	// primary signer and the one which pays the fee. The fee can be calculated
	// based on the cost of evaluating the body and doing signature verification
	// of the signers. This can be estimated via simulation.
	Fee *Fee `protobuf:"bytes,2,opt,name=fee,proto3" json:"fee,omitempty"`
	// Tip is the optional tip used for meta-transactions.
	//
	// Since: cosmos-sdk 0.45
	Tip *Tip `protobuf:"bytes,3,opt,name=tip,proto3" json:"tip,omitempty"`
}

func (x *AuthInfo) Reset() {
	*x = AuthInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_tx_v1beta1_tx_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthInfo) ProtoMessage() {}

// Deprecated: Use AuthInfo.ProtoReflect.Descriptor instead.
func (*AuthInfo) Descriptor() ([]byte, []int) {
	return file_cosmos_tx_v1beta1_tx_proto_rawDescGZIP(), []int{5}
}

func (x *AuthInfo) GetSignerInfos() []*SignerInfo {
	if x != nil {
		return x.SignerInfos
	}
	return nil
}

func (x *AuthInfo) GetFee() *Fee {
	if x != nil {
		return x.Fee
	}
	return nil
}

func (x *AuthInfo) GetTip() *Tip {
	if x != nil {
		return x.Tip
	}
	return nil
}

// SignerInfo describes the public key and signing mode of a single top-level
// signer.
type SignerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// public_key is the public key of the signer. It is optional for accounts
	// that already exist in state. If unset, the verifier can use the required \
	// signer address for this position and lookup the public key.
	PublicKey *anypb.Any `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// mode_info describes the signing mode of the signer and is a nested
	// structure to support nested multisig pubkey's
	ModeInfo *ModeInfo `protobuf:"bytes,2,opt,name=mode_info,json=modeInfo,proto3" json:"mode_info,omitempty"`
	// sequence is the sequence of the account, which describes the
	// number of committed transactions signed by a given address. It is used to
	// prevent replay attacks.
	Sequence uint64 `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (x *SignerInfo) Reset() {
	*x = SignerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_tx_v1beta1_tx_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignerInfo) ProtoMessage() {}

// Deprecated: Use SignerInfo.ProtoReflect.Descriptor instead.
func (*SignerInfo) Descriptor() ([]byte, []int) {
	return file_cosmos_tx_v1beta1_tx_proto_rawDescGZIP(), []int{6}
}

func (x *SignerInfo) GetPublicKey() *anypb.Any {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *SignerInfo) GetModeInfo() *ModeInfo {
	if x != nil {
		return x.ModeInfo
	}
	return nil
}

func (x *SignerInfo) GetSequence() uint64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

// ModeInfo describes the signing mode of a single or nested multisig signer.
type ModeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// sum is the oneof that specifies whether this represents a single or nested
	// multisig signer
	//
	// Types that are assignable to Sum:
	//	*ModeInfo_Single_
	//	*ModeInfo_Multi_
	Sum isModeInfo_Sum `protobuf_oneof:"sum"`
}

func (x *ModeInfo) Reset() {
	*x = ModeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_tx_v1beta1_tx_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModeInfo) ProtoMessage() {}

// Deprecated: Use ModeInfo.ProtoReflect.Descriptor instead.
func (*ModeInfo) Descriptor() ([]byte, []int) {
	return file_cosmos_tx_v1beta1_tx_proto_rawDescGZIP(), []int{7}
}

func (x *ModeInfo) GetSum() isModeInfo_Sum {
	if x != nil {
		return x.Sum
	}
	return nil
}

func (x *ModeInfo) GetSingle() *ModeInfo_Single {
	if x, ok := x.GetSum().(*ModeInfo_Single_); ok {
		return x.Single
	}
	return nil
}

func (x *ModeInfo) GetMulti() *ModeInfo_Multi {
	if x, ok := x.GetSum().(*ModeInfo_Multi_); ok {
		return x.Multi
	}
	return nil
}

type isModeInfo_Sum interface {
	isModeInfo_Sum()
}

type ModeInfo_Single_ struct {
	// single represents a single signer
	Single *ModeInfo_Single `protobuf:"bytes,1,opt,name=single,proto3,oneof"`
}

type ModeInfo_Multi_ struct {
	// multi represents a nested multisig signer
	Multi *ModeInfo_Multi `protobuf:"bytes,2,opt,name=multi,proto3,oneof"`
}

func (*ModeInfo_Single_) isModeInfo_Sum() {}

func (*ModeInfo_Multi_) isModeInfo_Sum() {}

// Fee includes the amount of coins paid in fees and the maximum
// gas to be used by the transaction. The ratio yields an effective "gasprice",
// which must be above some miminum to be accepted into the mempool.
type Fee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// amount is the amount of coins to be paid as a fee
	Amount []*v1beta12.Coin `protobuf:"bytes,1,rep,name=amount,proto3" json:"amount,omitempty"`
	// gas_limit is the maximum gas that can be used in transaction processing
	// before an out of gas error occurs
	GasLimit uint64 `protobuf:"varint,2,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	// if unset, the first signer is responsible for paying the fees. If set, the specified account must pay the fees.
	// the payer must be a tx signer (and thus have signed this field in AuthInfo).
	// setting this field does *not* change the ordering of required signers for the transaction.
	Payer string `protobuf:"bytes,3,opt,name=payer,proto3" json:"payer,omitempty"`
	// if set, the fee payer (either the first signer or the value of the payer field) requests that a fee grant be used
	// to pay fees instead of the fee payer's own balance. If an appropriate fee grant does not exist or the chain does
	// not support fee grants, this will fail
	Granter string `protobuf:"bytes,4,opt,name=granter,proto3" json:"granter,omitempty"`
}

func (x *Fee) Reset() {
	*x = Fee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_tx_v1beta1_tx_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fee) ProtoMessage() {}

// Deprecated: Use Fee.ProtoReflect.Descriptor instead.
func (*Fee) Descriptor() ([]byte, []int) {
	return file_cosmos_tx_v1beta1_tx_proto_rawDescGZIP(), []int{8}
}

func (x *Fee) GetAmount() []*v1beta12.Coin {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *Fee) GetGasLimit() uint64 {
	if x != nil {
		return x.GasLimit
	}
	return 0
}

func (x *Fee) GetPayer() string {
	if x != nil {
		return x.Payer
	}
	return ""
}

func (x *Fee) GetGranter() string {
	if x != nil {
		return x.Granter
	}
	return ""
}

// Tip is the tip used for meta-transactions.
type Tip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// amount is the amount of the tip
	Amount []*v1beta12.Coin `protobuf:"bytes,1,rep,name=amount,proto3" json:"amount,omitempty"`
	// tipper is the address of the account paying for the tip
	Tipper string `protobuf:"bytes,2,opt,name=tipper,proto3" json:"tipper,omitempty"`
}

func (x *Tip) Reset() {
	*x = Tip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_tx_v1beta1_tx_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tip) ProtoMessage() {}

// Deprecated: Use Tip.ProtoReflect.Descriptor instead.
func (*Tip) Descriptor() ([]byte, []int) {
	return file_cosmos_tx_v1beta1_tx_proto_rawDescGZIP(), []int{9}
}

func (x *Tip) GetAmount() []*v1beta12.Coin {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *Tip) GetTipper() string {
	if x != nil {
		return x.Tipper
	}
	return ""
}

// AuxSignerData is the intermediary format that an auxiliary signer (e.g. a
// tipper) builds and sends to the fee payer (who will build and broadcast the
// actual tx). AuxSignerData is not a valid tx in itself, and will be rejected
// by the node if sent directly as-is.
//
// Since: cosmos-sdk 0.45
type AuxSignerData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// address is the bech32-encoded address of the auxiliary signer. If using
	// AuxSignerData across different chains, the bech32 prefix of the target
	// chain (where the final transaction is broadcasted) should be used.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// sign_doc is the SIGN_MOD_DIRECT_AUX sign doc that the auxiliary signer
	// signs. Note: we use the same sign doc even if we're signing with
	// LEGACY_AMINO_JSON.
	SignDoc *SignDocDirectAux `protobuf:"bytes,2,opt,name=sign_doc,json=signDoc,proto3" json:"sign_doc,omitempty"`
	// mode is the signing mode of the single signer
	Mode v1beta1.SignMode `protobuf:"varint,3,opt,name=mode,proto3,enum=cosmos.tx.signing.v1beta1.SignMode" json:"mode,omitempty"`
	// sig is the signature of the sign doc.
	Sig []byte `protobuf:"bytes,4,opt,name=sig,proto3" json:"sig,omitempty"`
}

func (x *AuxSignerData) Reset() {
	*x = AuxSignerData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_tx_v1beta1_tx_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuxSignerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuxSignerData) ProtoMessage() {}

// Deprecated: Use AuxSignerData.ProtoReflect.Descriptor instead.
func (*AuxSignerData) Descriptor() ([]byte, []int) {
	return file_cosmos_tx_v1beta1_tx_proto_rawDescGZIP(), []int{10}
}

func (x *AuxSignerData) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AuxSignerData) GetSignDoc() *SignDocDirectAux {
	if x != nil {
		return x.SignDoc
	}
	return nil
}

func (x *AuxSignerData) GetMode() v1beta1.SignMode {
	if x != nil {
		return x.Mode
	}
	return v1beta1.SignMode(0)
}

func (x *AuxSignerData) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

// Single is the mode info for a single signer. It is structured as a message
// to allow for additional fields such as locale for SIGN_MODE_TEXTUAL in the
// future
type ModeInfo_Single struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// mode is the signing mode of the single signer
	Mode v1beta1.SignMode `protobuf:"varint,1,opt,name=mode,proto3,enum=cosmos.tx.signing.v1beta1.SignMode" json:"mode,omitempty"`
}

func (x *ModeInfo_Single) Reset() {
	*x = ModeInfo_Single{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_tx_v1beta1_tx_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModeInfo_Single) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModeInfo_Single) ProtoMessage() {}

// Deprecated: Use ModeInfo_Single.ProtoReflect.Descriptor instead.
func (*ModeInfo_Single) Descriptor() ([]byte, []int) {
	return file_cosmos_tx_v1beta1_tx_proto_rawDescGZIP(), []int{7, 0}
}

func (x *ModeInfo_Single) GetMode() v1beta1.SignMode {
	if x != nil {
		return x.Mode
	}
	return v1beta1.SignMode(0)
}

// Multi is the mode info for a multisig public key
type ModeInfo_Multi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// bitarray specifies which keys within the multisig are signing
	Bitarray *v1beta11.CompactBitArray `protobuf:"bytes,1,opt,name=bitarray,proto3" json:"bitarray,omitempty"`
	// mode_infos is the corresponding modes of the signers of the multisig
	// which could include nested multisig public keys
	ModeInfos []*ModeInfo `protobuf:"bytes,2,rep,name=mode_infos,json=modeInfos,proto3" json:"mode_infos,omitempty"`
}

func (x *ModeInfo_Multi) Reset() {
	*x = ModeInfo_Multi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_tx_v1beta1_tx_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModeInfo_Multi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModeInfo_Multi) ProtoMessage() {}

// Deprecated: Use ModeInfo_Multi.ProtoReflect.Descriptor instead.
func (*ModeInfo_Multi) Descriptor() ([]byte, []int) {
	return file_cosmos_tx_v1beta1_tx_proto_rawDescGZIP(), []int{7, 1}
}

func (x *ModeInfo_Multi) GetBitarray() *v1beta11.CompactBitArray {
	if x != nil {
		return x.Bitarray
	}
	return nil
}

func (x *ModeInfo_Multi) GetModeInfos() []*ModeInfo {
	if x != nil {
		return x.ModeInfos
	}
	return nil
}

var File_cosmos_tx_v1beta1_tx_proto protoreflect.FileDescriptor

var file_cosmos_tx_v1beta1_tx_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x74, 0x78, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a,
	0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x74, 0x78, 0x2f,
	0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x02, 0x54, 0x78, 0x12, 0x2d, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x78, 0x42,
	0x6f, 0x64, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x22, 0x6e, 0x0a, 0x05, 0x54, 0x78, 0x52, 0x61, 0x77, 0x12, 0x1d, 0x0a, 0x0a,
	0x62, 0x6f, 0x64, 0x79, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x09, 0x62, 0x6f, 0x64, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x22, 0x92, 0x01, 0x0a, 0x07, 0x53, 0x69, 0x67, 0x6e, 0x44, 0x6f, 0x63, 0x12,
	0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x6f, 0x64, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x26,
	0x0a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66,
	0x6f, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49,
	0x64, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xee, 0x01, 0x0a, 0x10, 0x53, 0x69, 0x67,
	0x6e, 0x44, 0x6f, 0x63, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x41, 0x75, 0x78, 0x12, 0x1d, 0x0a,
	0x0a, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x62, 0x6f, 0x64, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x0a,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x28, 0x0a, 0x03, 0x74, 0x69, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x54, 0x69, 0x70, 0x52, 0x03, 0x74, 0x69, 0x70, 0x22, 0x95, 0x02, 0x0a, 0x06, 0x54, 0x78,
	0x42, 0x6f, 0x64, 0x79, 0x12, 0x30, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x08, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x42, 0x0a, 0x11, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xff, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x10, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x5a, 0x0a, 0x1e, 0x6e, 0x6f, 0x6e, 0x5f, 0x63, 0x72, 0x69,
	0x74, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xff, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x1b, 0x6e, 0x6f, 0x6e, 0x43, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61,
	0x6c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0xa0, 0x01, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x40,
	0x0a, 0x0c, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x78,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x12, 0x28, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x46, 0x65, 0x65, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x28, 0x0a, 0x03, 0x74, 0x69,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x69, 0x70, 0x52,
	0x03, 0x74, 0x69, 0x70, 0x22, 0x97, 0x01, 0x0a, 0x0a, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x33, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x09, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x6d, 0x6f, 0x64, 0x65,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x4d, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x22, 0xe0,
	0x02, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3c, 0x0a, 0x06, 0x73,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x4d, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x48,
	0x00, 0x52, 0x06, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x6f, 0x64,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x48, 0x00, 0x52, 0x05, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x1a, 0x41, 0x0a, 0x06, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x37,
	0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x78, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x4d, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x1a, 0x90, 0x01, 0x0a, 0x05, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x12, 0x4b, 0x0a, 0x08, 0x62, 0x69, 0x74, 0x61, 0x72, 0x72, 0x61, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x42, 0x69, 0x74, 0x41,
	0x72, 0x72, 0x61, 0x79, 0x52, 0x08, 0x62, 0x69, 0x74, 0x61, 0x72, 0x72, 0x61, 0x79, 0x12, 0x3a,
	0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x78, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x09, 0x6d, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x42, 0x05, 0x0a, 0x03, 0x73, 0x75,
	0x6d, 0x22, 0xeb, 0x01, 0x0a, 0x03, 0x46, 0x65, 0x65, 0x12, 0x63, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x43, 0x6f, 0x69, 0x6e, 0x42, 0x30, 0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x28, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x67, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x2e, 0x0a, 0x05, 0x70,
	0x61, 0x79, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x52, 0x05, 0x70, 0x61, 0x79, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x07, 0x67,
	0x72, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4,
	0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x22,
	0x9c, 0x01, 0x0a, 0x03, 0x54, 0x69, 0x70, 0x12, 0x63, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f,
	0x69, 0x6e, 0x42, 0x30, 0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x28, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43,
	0x6f, 0x69, 0x6e, 0x73, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x06,
	0x74, 0x69, 0x70, 0x70, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4,
	0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x74, 0x69, 0x70, 0x70, 0x65, 0x72, 0x22, 0xce,
	0x01, 0x0a, 0x0d, 0x41, 0x75, 0x78, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x32, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x3e, 0x0a, 0x08, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x64, 0x6f, 0x63,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x74, 0x78, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x44,
	0x6f, 0x63, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x41, 0x75, 0x78, 0x52, 0x07, 0x73, 0x69, 0x67,
	0x6e, 0x44, 0x6f, 0x63, 0x12, 0x37, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x78, 0x2e, 0x73,
	0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x73, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x73, 0x69, 0x67, 0x42,
	0xc4, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x74,
	0x78, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x07, 0x54, 0x78, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x74, 0x78,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x74, 0x78, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x54, 0x58, 0xaa, 0x02, 0x11, 0x43, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x54, 0x78, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x11, 0x43,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x54, 0x78, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0xe2, 0x02, 0x1d, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x54, 0x78, 0x5c, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x13, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x54, 0x78, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_tx_v1beta1_tx_proto_rawDescOnce sync.Once
	file_cosmos_tx_v1beta1_tx_proto_rawDescData = file_cosmos_tx_v1beta1_tx_proto_rawDesc
)

func file_cosmos_tx_v1beta1_tx_proto_rawDescGZIP() []byte {
	file_cosmos_tx_v1beta1_tx_proto_rawDescOnce.Do(func() {
		file_cosmos_tx_v1beta1_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_tx_v1beta1_tx_proto_rawDescData)
	})
	return file_cosmos_tx_v1beta1_tx_proto_rawDescData
}

var file_cosmos_tx_v1beta1_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_cosmos_tx_v1beta1_tx_proto_goTypes = []interface{}{
	(*Tx)(nil),                       // 0: cosmos.tx.v1beta1.Tx
	(*TxRaw)(nil),                    // 1: cosmos.tx.v1beta1.TxRaw
	(*SignDoc)(nil),                  // 2: cosmos.tx.v1beta1.SignDoc
	(*SignDocDirectAux)(nil),         // 3: cosmos.tx.v1beta1.SignDocDirectAux
	(*TxBody)(nil),                   // 4: cosmos.tx.v1beta1.TxBody
	(*AuthInfo)(nil),                 // 5: cosmos.tx.v1beta1.AuthInfo
	(*SignerInfo)(nil),               // 6: cosmos.tx.v1beta1.SignerInfo
	(*ModeInfo)(nil),                 // 7: cosmos.tx.v1beta1.ModeInfo
	(*Fee)(nil),                      // 8: cosmos.tx.v1beta1.Fee
	(*Tip)(nil),                      // 9: cosmos.tx.v1beta1.Tip
	(*AuxSignerData)(nil),            // 10: cosmos.tx.v1beta1.AuxSignerData
	(*ModeInfo_Single)(nil),          // 11: cosmos.tx.v1beta1.ModeInfo.Single
	(*ModeInfo_Multi)(nil),           // 12: cosmos.tx.v1beta1.ModeInfo.Multi
	(*anypb.Any)(nil),                // 13: google.protobuf.Any
	(*v1beta12.Coin)(nil),            // 14: cosmos.base.v1beta1.Coin
	(v1beta1.SignMode)(0),            // 15: cosmos.tx.signing.v1beta1.SignMode
	(*v1beta11.CompactBitArray)(nil), // 16: cosmos.crypto.multisig.v1beta1.CompactBitArray
}
var file_cosmos_tx_v1beta1_tx_proto_depIdxs = []int32{
	4,  // 0: cosmos.tx.v1beta1.Tx.body:type_name -> cosmos.tx.v1beta1.TxBody
	5,  // 1: cosmos.tx.v1beta1.Tx.auth_info:type_name -> cosmos.tx.v1beta1.AuthInfo
	13, // 2: cosmos.tx.v1beta1.SignDocDirectAux.public_key:type_name -> google.protobuf.Any
	9,  // 3: cosmos.tx.v1beta1.SignDocDirectAux.tip:type_name -> cosmos.tx.v1beta1.Tip
	13, // 4: cosmos.tx.v1beta1.TxBody.messages:type_name -> google.protobuf.Any
	13, // 5: cosmos.tx.v1beta1.TxBody.extension_options:type_name -> google.protobuf.Any
	13, // 6: cosmos.tx.v1beta1.TxBody.non_critical_extension_options:type_name -> google.protobuf.Any
	6,  // 7: cosmos.tx.v1beta1.AuthInfo.signer_infos:type_name -> cosmos.tx.v1beta1.SignerInfo
	8,  // 8: cosmos.tx.v1beta1.AuthInfo.fee:type_name -> cosmos.tx.v1beta1.Fee
	9,  // 9: cosmos.tx.v1beta1.AuthInfo.tip:type_name -> cosmos.tx.v1beta1.Tip
	13, // 10: cosmos.tx.v1beta1.SignerInfo.public_key:type_name -> google.protobuf.Any
	7,  // 11: cosmos.tx.v1beta1.SignerInfo.mode_info:type_name -> cosmos.tx.v1beta1.ModeInfo
	11, // 12: cosmos.tx.v1beta1.ModeInfo.single:type_name -> cosmos.tx.v1beta1.ModeInfo.Single
	12, // 13: cosmos.tx.v1beta1.ModeInfo.multi:type_name -> cosmos.tx.v1beta1.ModeInfo.Multi
	14, // 14: cosmos.tx.v1beta1.Fee.amount:type_name -> cosmos.base.v1beta1.Coin
	14, // 15: cosmos.tx.v1beta1.Tip.amount:type_name -> cosmos.base.v1beta1.Coin
	3,  // 16: cosmos.tx.v1beta1.AuxSignerData.sign_doc:type_name -> cosmos.tx.v1beta1.SignDocDirectAux
	15, // 17: cosmos.tx.v1beta1.AuxSignerData.mode:type_name -> cosmos.tx.signing.v1beta1.SignMode
	15, // 18: cosmos.tx.v1beta1.ModeInfo.Single.mode:type_name -> cosmos.tx.signing.v1beta1.SignMode
	16, // 19: cosmos.tx.v1beta1.ModeInfo.Multi.bitarray:type_name -> cosmos.crypto.multisig.v1beta1.CompactBitArray
	7,  // 20: cosmos.tx.v1beta1.ModeInfo.Multi.mode_infos:type_name -> cosmos.tx.v1beta1.ModeInfo
	21, // [21:21] is the sub-list for method output_type
	21, // [21:21] is the sub-list for method input_type
	21, // [21:21] is the sub-list for extension type_name
	21, // [21:21] is the sub-list for extension extendee
	0,  // [0:21] is the sub-list for field type_name
}

func init() { file_cosmos_tx_v1beta1_tx_proto_init() }
func file_cosmos_tx_v1beta1_tx_proto_init() {
	if File_cosmos_tx_v1beta1_tx_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_tx_v1beta1_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tx); i {
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
		file_cosmos_tx_v1beta1_tx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxRaw); i {
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
		file_cosmos_tx_v1beta1_tx_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignDoc); i {
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
		file_cosmos_tx_v1beta1_tx_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignDocDirectAux); i {
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
		file_cosmos_tx_v1beta1_tx_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxBody); i {
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
		file_cosmos_tx_v1beta1_tx_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthInfo); i {
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
		file_cosmos_tx_v1beta1_tx_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignerInfo); i {
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
		file_cosmos_tx_v1beta1_tx_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModeInfo); i {
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
		file_cosmos_tx_v1beta1_tx_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fee); i {
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
		file_cosmos_tx_v1beta1_tx_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tip); i {
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
		file_cosmos_tx_v1beta1_tx_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuxSignerData); i {
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
		file_cosmos_tx_v1beta1_tx_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModeInfo_Single); i {
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
		file_cosmos_tx_v1beta1_tx_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModeInfo_Multi); i {
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
	file_cosmos_tx_v1beta1_tx_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*ModeInfo_Single_)(nil),
		(*ModeInfo_Multi_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cosmos_tx_v1beta1_tx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_tx_v1beta1_tx_proto_goTypes,
		DependencyIndexes: file_cosmos_tx_v1beta1_tx_proto_depIdxs,
		MessageInfos:      file_cosmos_tx_v1beta1_tx_proto_msgTypes,
	}.Build()
	File_cosmos_tx_v1beta1_tx_proto = out.File
	file_cosmos_tx_v1beta1_tx_proto_rawDesc = nil
	file_cosmos_tx_v1beta1_tx_proto_goTypes = nil
	file_cosmos_tx_v1beta1_tx_proto_depIdxs = nil
}
