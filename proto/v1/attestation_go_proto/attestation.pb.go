// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/v1/attestation.proto

package attestation_go_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Type (for example schema) of the attestation payload that was signed.
type PgpSignedAttestation_ContentType int32

const (
	// `ContentType` is not set.
	PgpSignedAttestation_CONTENT_TYPE_UNSPECIFIED PgpSignedAttestation_ContentType = 0
	// Atomic format attestation signature. See
	// https://github.com/containers/image/blob/8a5d2f82a6e3263290c8e0276c3e0f64e77723e7/docs/atomic-signature.md
	// The payload extracted from `signature` is a JSON blob conforming to the
	// linked schema.
	PgpSignedAttestation_SIMPLE_SIGNING_JSON PgpSignedAttestation_ContentType = 1
)

var PgpSignedAttestation_ContentType_name = map[int32]string{
	0: "CONTENT_TYPE_UNSPECIFIED",
	1: "SIMPLE_SIGNING_JSON",
}

var PgpSignedAttestation_ContentType_value = map[string]int32{
	"CONTENT_TYPE_UNSPECIFIED": 0,
	"SIMPLE_SIGNING_JSON":      1,
}

func (x PgpSignedAttestation_ContentType) String() string {
	return proto.EnumName(PgpSignedAttestation_ContentType_name, int32(x))
}

func (PgpSignedAttestation_ContentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_665b62322cc527a8, []int{0, 0}
}

// DO NOT USE: UNDER HEAVY DEVELOPMENT.
// TODO(aysylu): finalize this.
//
// An attestation wrapper with a PGP-compatible signature. This message only
// supports `ATTACHED` signatures, where the payload that is signed is included
// alongside the signature itself in the same file.
type PgpSignedAttestation struct {
	// Required. The raw content of the signature, as output by GNU Privacy Guard
	// (GPG) or equivalent. Since this message only supports attached signatures,
	// the payload that was signed must be attached. While the signature format
	// supported is dependent on the verification implementation, currently only
	// ASCII-armored (`--armor` to gpg), non-clearsigned (`--sign` rather than
	// `--clearsign` to gpg) are supported. Concretely, `gpg --sign --armor
	// --output=signature.gpg payload.json` will create the signature content
	// expected in this field in `signature.gpg` for the `payload.json`
	// attestation payload.
	Signature string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// Type (for example schema) of the attestation payload that was signed.
	// The verifier must ensure that the provided type is one that the verifier
	// supports, and that the attestation payload is a valid instantiation of that
	// type (for example by validating a JSON schema).
	ContentType PgpSignedAttestation_ContentType `protobuf:"varint,3,opt,name=content_type,json=contentType,proto3,enum=grafeas.v1.attestation.PgpSignedAttestation_ContentType" json:"content_type,omitempty"`
	// This field is used by verifiers to select the public key used to validate
	// the signature. Note that the policy of the verifier ultimately determines
	// which public keys verify a signature based on the context of the
	// verification. There is no guarantee validation will succeed if the
	// verifier has no key matching this ID, even if it has a key under a
	// different ID that would verify the signature. Note that this ID should also
	// be present in the signature content above, but that is not expected to be
	// used by the verifier.
	//
	// Types that are valid to be assigned to KeyId:
	//	*PgpSignedAttestation_PgpKeyId
	KeyId                isPgpSignedAttestation_KeyId `protobuf_oneof:"key_id"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *PgpSignedAttestation) Reset()         { *m = PgpSignedAttestation{} }
func (m *PgpSignedAttestation) String() string { return proto.CompactTextString(m) }
func (*PgpSignedAttestation) ProtoMessage()    {}
func (*PgpSignedAttestation) Descriptor() ([]byte, []int) {
	return fileDescriptor_665b62322cc527a8, []int{0}
}

func (m *PgpSignedAttestation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PgpSignedAttestation.Unmarshal(m, b)
}
func (m *PgpSignedAttestation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PgpSignedAttestation.Marshal(b, m, deterministic)
}
func (m *PgpSignedAttestation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PgpSignedAttestation.Merge(m, src)
}
func (m *PgpSignedAttestation) XXX_Size() int {
	return xxx_messageInfo_PgpSignedAttestation.Size(m)
}
func (m *PgpSignedAttestation) XXX_DiscardUnknown() {
	xxx_messageInfo_PgpSignedAttestation.DiscardUnknown(m)
}

var xxx_messageInfo_PgpSignedAttestation proto.InternalMessageInfo

func (m *PgpSignedAttestation) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *PgpSignedAttestation) GetContentType() PgpSignedAttestation_ContentType {
	if m != nil {
		return m.ContentType
	}
	return PgpSignedAttestation_CONTENT_TYPE_UNSPECIFIED
}

type isPgpSignedAttestation_KeyId interface {
	isPgpSignedAttestation_KeyId()
}

type PgpSignedAttestation_PgpKeyId struct {
	PgpKeyId string `protobuf:"bytes,2,opt,name=pgp_key_id,json=pgpKeyId,proto3,oneof"`
}

func (*PgpSignedAttestation_PgpKeyId) isPgpSignedAttestation_KeyId() {}

func (m *PgpSignedAttestation) GetKeyId() isPgpSignedAttestation_KeyId {
	if m != nil {
		return m.KeyId
	}
	return nil
}

func (m *PgpSignedAttestation) GetPgpKeyId() string {
	if x, ok := m.GetKeyId().(*PgpSignedAttestation_PgpKeyId); ok {
		return x.PgpKeyId
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PgpSignedAttestation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PgpSignedAttestation_PgpKeyId)(nil),
	}
}

// Note kind that represents a logical attestation "role" or "authority". For
// example, an organization might have one `Authority` for "QA" and one for
// "build". This note is intended to act strictly as a grouping mechanism for
// the attached occurrences (Attestations). This grouping mechanism also
// provides a security boundary, since IAM ACLs gate the ability for a principle
// to attach an occurrence to a given note. It also provides a single point of
// lookup to find all attached attestation occurrences, even if they don't all
// live in the same project.
type Authority struct {
	// Hint hints at the purpose of the attestation authority.
	Hint                 *Authority_Hint `protobuf:"bytes,1,opt,name=hint,proto3" json:"hint,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Authority) Reset()         { *m = Authority{} }
func (m *Authority) String() string { return proto.CompactTextString(m) }
func (*Authority) ProtoMessage()    {}
func (*Authority) Descriptor() ([]byte, []int) {
	return fileDescriptor_665b62322cc527a8, []int{1}
}

func (m *Authority) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Authority.Unmarshal(m, b)
}
func (m *Authority) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Authority.Marshal(b, m, deterministic)
}
func (m *Authority) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Authority.Merge(m, src)
}
func (m *Authority) XXX_Size() int {
	return xxx_messageInfo_Authority.Size(m)
}
func (m *Authority) XXX_DiscardUnknown() {
	xxx_messageInfo_Authority.DiscardUnknown(m)
}

var xxx_messageInfo_Authority proto.InternalMessageInfo

func (m *Authority) GetHint() *Authority_Hint {
	if m != nil {
		return m.Hint
	}
	return nil
}

// This submessage provides human-readable hints about the purpose of the
// authority. Because the name of a note acts as its resource reference, it is
// important to disambiguate the canonical name of the Note (which might be a
// UUID for security purposes) from "readable" names more suitable for debug
// output. Note that these hints should not be used to look up authorities in
// security sensitive contexts, such as when looking up attestations to
// verify.
type Authority_Hint struct {
	// Required. The human readable name of this attestation authority, for
	// example "qa".
	HumanReadableName    string   `protobuf:"bytes,1,opt,name=human_readable_name,json=humanReadableName,proto3" json:"human_readable_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Authority_Hint) Reset()         { *m = Authority_Hint{} }
func (m *Authority_Hint) String() string { return proto.CompactTextString(m) }
func (*Authority_Hint) ProtoMessage()    {}
func (*Authority_Hint) Descriptor() ([]byte, []int) {
	return fileDescriptor_665b62322cc527a8, []int{1, 0}
}

func (m *Authority_Hint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Authority_Hint.Unmarshal(m, b)
}
func (m *Authority_Hint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Authority_Hint.Marshal(b, m, deterministic)
}
func (m *Authority_Hint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Authority_Hint.Merge(m, src)
}
func (m *Authority_Hint) XXX_Size() int {
	return xxx_messageInfo_Authority_Hint.Size(m)
}
func (m *Authority_Hint) XXX_DiscardUnknown() {
	xxx_messageInfo_Authority_Hint.DiscardUnknown(m)
}

var xxx_messageInfo_Authority_Hint proto.InternalMessageInfo

func (m *Authority_Hint) GetHumanReadableName() string {
	if m != nil {
		return m.HumanReadableName
	}
	return ""
}

// Details of an attestation occurrence.
type Details struct {
	// Required. Attestation for the resource.
	Attestation          *Attestation `protobuf:"bytes,1,opt,name=attestation,proto3" json:"attestation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Details) Reset()         { *m = Details{} }
func (m *Details) String() string { return proto.CompactTextString(m) }
func (*Details) ProtoMessage()    {}
func (*Details) Descriptor() ([]byte, []int) {
	return fileDescriptor_665b62322cc527a8, []int{2}
}

func (m *Details) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Details.Unmarshal(m, b)
}
func (m *Details) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Details.Marshal(b, m, deterministic)
}
func (m *Details) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Details.Merge(m, src)
}
func (m *Details) XXX_Size() int {
	return xxx_messageInfo_Details.Size(m)
}
func (m *Details) XXX_DiscardUnknown() {
	xxx_messageInfo_Details.DiscardUnknown(m)
}

var xxx_messageInfo_Details proto.InternalMessageInfo

func (m *Details) GetAttestation() *Attestation {
	if m != nil {
		return m.Attestation
	}
	return nil
}

// Occurrence that represents a single "attestation". The authenticity of an
// attestation can be verified using the attached signature. If the verifier
// trusts the public key of the signer, then verifying the signature is
// sufficient to establish trust. In this circumstance, the authority to which
// this attestation is attached is primarily useful for look-up (how to find
// this attestation if you already know the authority and artifact to be
// verified) and intent (which authority was this attestation intended to sign
// for).
type Attestation struct {
	// Required. The signature, generally over the `resource_url`, that verifies
	// this attestation. The semantics of the signature veracity are ultimately
	// determined by the verification engine.
	//
	// Types that are valid to be assigned to Signature:
	//	*Attestation_PgpSignedAttestation
	Signature            isAttestation_Signature `protobuf_oneof:"signature"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Attestation) Reset()         { *m = Attestation{} }
func (m *Attestation) String() string { return proto.CompactTextString(m) }
func (*Attestation) ProtoMessage()    {}
func (*Attestation) Descriptor() ([]byte, []int) {
	return fileDescriptor_665b62322cc527a8, []int{3}
}

func (m *Attestation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attestation.Unmarshal(m, b)
}
func (m *Attestation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attestation.Marshal(b, m, deterministic)
}
func (m *Attestation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attestation.Merge(m, src)
}
func (m *Attestation) XXX_Size() int {
	return xxx_messageInfo_Attestation.Size(m)
}
func (m *Attestation) XXX_DiscardUnknown() {
	xxx_messageInfo_Attestation.DiscardUnknown(m)
}

var xxx_messageInfo_Attestation proto.InternalMessageInfo

type isAttestation_Signature interface {
	isAttestation_Signature()
}

type Attestation_PgpSignedAttestation struct {
	PgpSignedAttestation *PgpSignedAttestation `protobuf:"bytes,1,opt,name=pgp_signed_attestation,json=pgpSignedAttestation,proto3,oneof"`
}

func (*Attestation_PgpSignedAttestation) isAttestation_Signature() {}

func (m *Attestation) GetSignature() isAttestation_Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Attestation) GetPgpSignedAttestation() *PgpSignedAttestation {
	if x, ok := m.GetSignature().(*Attestation_PgpSignedAttestation); ok {
		return x.PgpSignedAttestation
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Attestation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Attestation_PgpSignedAttestation)(nil),
	}
}

func init() {
	proto.RegisterEnum("grafeas.v1.attestation.PgpSignedAttestation_ContentType", PgpSignedAttestation_ContentType_name, PgpSignedAttestation_ContentType_value)
	proto.RegisterType((*PgpSignedAttestation)(nil), "grafeas.v1.attestation.PgpSignedAttestation")
	proto.RegisterType((*Authority)(nil), "grafeas.v1.attestation.Authority")
	proto.RegisterType((*Authority_Hint)(nil), "grafeas.v1.attestation.Authority.Hint")
	proto.RegisterType((*Details)(nil), "grafeas.v1.attestation.Details")
	proto.RegisterType((*Attestation)(nil), "grafeas.v1.attestation.Attestation")
}

func init() { proto.RegisterFile("proto/v1/attestation.proto", fileDescriptor_665b62322cc527a8) }

var fileDescriptor_665b62322cc527a8 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x6d, 0xba, 0x69, 0xac, 0x37, 0x08, 0x0d, 0x6f, 0x1a, 0x65, 0x9a, 0xd0, 0x14, 0x24, 0xb4,
	0x07, 0xe4, 0x6a, 0x43, 0x42, 0x13, 0x6f, 0xfd, 0x08, 0x6d, 0xf8, 0xc8, 0x22, 0xa7, 0x3c, 0x00,
	0x42, 0x96, 0xdb, 0x18, 0xd7, 0x62, 0xb5, 0xad, 0xc4, 0x99, 0x94, 0x27, 0xe0, 0xaf, 0xf0, 0x4b,
	0x51, 0xb3, 0xd2, 0x44, 0x22, 0x7d, 0xd8, 0x53, 0x72, 0xcf, 0xf1, 0xb9, 0xe7, 0x7e, 0xc1, 0x89,
	0x49, 0xb5, 0xd5, 0xbd, 0xdb, 0x8b, 0x1e, 0xb3, 0x96, 0x67, 0x96, 0x59, 0xa9, 0x15, 0x2e, 0x41,
	0x74, 0x2c, 0x52, 0xf6, 0x9d, 0xb3, 0x0c, 0xdf, 0x5e, 0xe0, 0x1a, 0xeb, 0xfd, 0x6e, 0xc3, 0x51,
	0x24, 0x4c, 0x2c, 0x85, 0xe2, 0x49, 0xbf, 0x22, 0xd0, 0x29, 0x74, 0x32, 0x29, 0x14, 0xb3, 0x79,
	0xca, 0xbb, 0xce, 0x99, 0x73, 0xde, 0x21, 0x15, 0x80, 0xbe, 0xc2, 0xc3, 0xb9, 0x56, 0x96, 0x2b,
	0x4b, 0x6d, 0x61, 0x78, 0x77, 0xe7, 0xcc, 0x39, 0x7f, 0x74, 0x79, 0x85, 0x9b, 0x5d, 0x70, 0x93,
	0x03, 0x1e, 0xde, 0x25, 0x98, 0x16, 0x86, 0x13, 0x77, 0x5e, 0x05, 0xe8, 0x19, 0x80, 0x11, 0x86,
	0xfe, 0xe0, 0x05, 0x95, 0x49, 0xb7, 0xbd, 0xf2, 0x9e, 0xb4, 0xc8, 0xbe, 0x11, 0xe6, 0x3d, 0x2f,
	0x82, 0xc4, 0x1b, 0x81, 0x5b, 0xd3, 0xa2, 0x53, 0xe8, 0x0e, 0xaf, 0xc3, 0xa9, 0x1f, 0x4e, 0xe9,
	0xf4, 0x73, 0xe4, 0xd3, 0x4f, 0x61, 0x1c, 0xf9, 0xc3, 0xe0, 0x6d, 0xe0, 0x8f, 0x0e, 0x5a, 0xe8,
	0x09, 0x1c, 0xc6, 0xc1, 0xc7, 0xe8, 0x83, 0x4f, 0xe3, 0x60, 0x1c, 0x06, 0xe1, 0x98, 0xbe, 0x8b,
	0xaf, 0xc3, 0x03, 0x67, 0xb0, 0x0f, 0x7b, 0x77, 0x0e, 0xde, 0x4f, 0xe8, 0xf4, 0x73, 0xbb, 0xd0,
	0xa9, 0xb4, 0x05, 0x7a, 0x03, 0xbb, 0x0b, 0xa9, 0x6c, 0xd9, 0xb2, 0x7b, 0xf9, 0x62, 0x5b, 0x47,
	0x1b, 0x01, 0x9e, 0x48, 0x65, 0x49, 0xa9, 0x39, 0x79, 0x0d, 0xbb, 0xab, 0x08, 0x61, 0x38, 0x5c,
	0xe4, 0x4b, 0xa6, 0x68, 0xca, 0x59, 0xc2, 0x66, 0x37, 0x9c, 0x2a, 0xb6, 0xfc, 0x37, 0xc5, 0xc7,
	0x25, 0x45, 0xd6, 0x4c, 0xc8, 0x96, 0xdc, 0x8b, 0xe0, 0xc1, 0x88, 0x5b, 0x26, 0x6f, 0x32, 0xe4,
	0x83, 0x5b, 0xb3, 0x59, 0x57, 0xf1, 0x7c, 0x6b, 0x15, 0xd5, 0x3f, 0xa9, 0xeb, 0xbc, 0x5f, 0x0e,
	0xb8, 0xf5, 0x6d, 0x26, 0x70, 0xbc, 0x1a, 0x69, 0x56, 0x2e, 0x81, 0xfe, 0xef, 0xf0, 0xf2, 0x3e,
	0x9b, 0x9b, 0xb4, 0xc8, 0x91, 0x69, 0xc0, 0x07, 0x6e, 0xed, 0x66, 0x06, 0xdf, 0xe0, 0xa9, 0xd4,
	0x5b, 0xd2, 0x46, 0xce, 0x97, 0x2b, 0x21, 0xed, 0x22, 0x9f, 0xe1, 0xb9, 0x5e, 0xf6, 0xd6, 0x8f,
	0x36, 0xdf, 0xa6, 0x3b, 0xa6, 0x42, 0xd3, 0x12, 0xff, 0xd3, 0xde, 0x19, 0x93, 0xfe, 0x6c, 0xaf,
	0x0c, 0x5e, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x92, 0xbe, 0xf5, 0x79, 0xf5, 0x02, 0x00, 0x00,
}
