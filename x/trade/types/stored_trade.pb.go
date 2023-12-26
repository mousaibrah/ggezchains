// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ggezchain/trade/stored_trade.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type StoredTrade struct {
	TradeIndex      uint64 `protobuf:"varint,1,opt,name=tradeIndex,proto3" json:"tradeIndex,omitempty"`
	TradeType       string `protobuf:"bytes,2,opt,name=tradeType,proto3" json:"tradeType,omitempty"`
	Coin            string `protobuf:"bytes,3,opt,name=coin,proto3" json:"coin,omitempty"`
	Price           string `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
	Quantity        string `protobuf:"bytes,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	ReceiverAddress string `protobuf:"bytes,6,opt,name=receiverAddress,proto3" json:"receiverAddress,omitempty"`
	Status          string `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	Maker           string `protobuf:"bytes,8,opt,name=maker,proto3" json:"maker,omitempty"`
	Checker         string `protobuf:"bytes,9,opt,name=checker,proto3" json:"checker,omitempty"`
	CreateDate      string `protobuf:"bytes,10,opt,name=createDate,proto3" json:"createDate,omitempty"`
	UpdateDate      string `protobuf:"bytes,11,opt,name=updateDate,proto3" json:"updateDate,omitempty"`
	ProcessDate     string `protobuf:"bytes,12,opt,name=processDate,proto3" json:"processDate,omitempty"`
	TradeData       string `protobuf:"bytes,13,opt,name=tradeData,proto3" json:"tradeData,omitempty"`
	Result          string `protobuf:"bytes,14,opt,name=result,proto3" json:"result,omitempty"`
}

func (m *StoredTrade) Reset()         { *m = StoredTrade{} }
func (m *StoredTrade) String() string { return proto.CompactTextString(m) }
func (*StoredTrade) ProtoMessage()    {}
func (*StoredTrade) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cbf8bf5cc0260cd, []int{0}
}
func (m *StoredTrade) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoredTrade) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoredTrade.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StoredTrade) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoredTrade.Merge(m, src)
}
func (m *StoredTrade) XXX_Size() int {
	return m.Size()
}
func (m *StoredTrade) XXX_DiscardUnknown() {
	xxx_messageInfo_StoredTrade.DiscardUnknown(m)
}

var xxx_messageInfo_StoredTrade proto.InternalMessageInfo

func (m *StoredTrade) GetTradeIndex() uint64 {
	if m != nil {
		return m.TradeIndex
	}
	return 0
}

func (m *StoredTrade) GetTradeType() string {
	if m != nil {
		return m.TradeType
	}
	return ""
}

func (m *StoredTrade) GetCoin() string {
	if m != nil {
		return m.Coin
	}
	return ""
}

func (m *StoredTrade) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *StoredTrade) GetQuantity() string {
	if m != nil {
		return m.Quantity
	}
	return ""
}

func (m *StoredTrade) GetReceiverAddress() string {
	if m != nil {
		return m.ReceiverAddress
	}
	return ""
}

func (m *StoredTrade) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *StoredTrade) GetMaker() string {
	if m != nil {
		return m.Maker
	}
	return ""
}

func (m *StoredTrade) GetChecker() string {
	if m != nil {
		return m.Checker
	}
	return ""
}

func (m *StoredTrade) GetCreateDate() string {
	if m != nil {
		return m.CreateDate
	}
	return ""
}

func (m *StoredTrade) GetUpdateDate() string {
	if m != nil {
		return m.UpdateDate
	}
	return ""
}

func (m *StoredTrade) GetProcessDate() string {
	if m != nil {
		return m.ProcessDate
	}
	return ""
}

func (m *StoredTrade) GetTradeData() string {
	if m != nil {
		return m.TradeData
	}
	return ""
}

func (m *StoredTrade) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*StoredTrade)(nil), "ggezchain.trade.StoredTrade")
}

func init() {
	proto.RegisterFile("ggezchain/trade/stored_trade.proto", fileDescriptor_4cbf8bf5cc0260cd)
}

var fileDescriptor_4cbf8bf5cc0260cd = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xcd, 0x4a, 0x2b, 0x31,
	0x14, 0xc7, 0x3b, 0xb7, 0xdf, 0xe9, 0xbd, 0xb7, 0x10, 0x2e, 0x97, 0x20, 0x12, 0x4a, 0x57, 0xdd,
	0xd8, 0x59, 0xf8, 0x04, 0x4a, 0x41, 0xdc, 0xd6, 0xae, 0xdc, 0x48, 0x9a, 0x39, 0x74, 0x82, 0x76,
	0x32, 0x26, 0x67, 0xa4, 0xf5, 0x29, 0xdc, 0xfb, 0x42, 0x2e, 0xbb, 0x74, 0x29, 0xed, 0x8b, 0x48,
	0xce, 0xf4, 0x63, 0x70, 0x37, 0xbf, 0xdf, 0xff, 0x0f, 0xe1, 0x9c, 0x39, 0x6c, 0xb8, 0x58, 0xc0,
	0xab, 0x4e, 0x95, 0xc9, 0x62, 0x74, 0x2a, 0x81, 0xd8, 0xa3, 0x75, 0x90, 0x3c, 0x10, 0x8c, 0x73,
	0x67, 0xd1, 0xf2, 0xfe, 0xb1, 0x33, 0x26, 0x3d, 0x7c, 0xaf, 0xb3, 0xde, 0x1d, 0xf5, 0x66, 0x81,
	0xb9, 0x64, 0x8c, 0x82, 0xdb, 0x2c, 0x81, 0x95, 0x88, 0x06, 0xd1, 0xa8, 0x31, 0xad, 0x18, 0x7e,
	0xce, 0xba, 0x44, 0xb3, 0x75, 0x0e, 0xe2, 0xd7, 0x20, 0x1a, 0x75, 0xa7, 0x27, 0xc1, 0x39, 0x6b,
	0x68, 0x6b, 0x32, 0x51, 0xa7, 0x80, 0xbe, 0xf9, 0x3f, 0xd6, 0xcc, 0x9d, 0xd1, 0x20, 0x1a, 0x24,
	0x4b, 0xe0, 0x67, 0xac, 0xf3, 0x5c, 0xa8, 0x0c, 0x0d, 0xae, 0x45, 0x93, 0x82, 0x23, 0xf3, 0x11,
	0xeb, 0x3b, 0xd0, 0x60, 0x5e, 0xc0, 0x5d, 0x25, 0x89, 0x03, 0xef, 0x45, 0x8b, 0x2a, 0x3f, 0x35,
	0xff, 0xcf, 0x5a, 0x1e, 0x15, 0x16, 0x5e, 0xb4, 0xa9, 0xb0, 0xa7, 0xf0, 0xe6, 0x52, 0x3d, 0x82,
	0x13, 0x9d, 0xf2, 0x4d, 0x02, 0x2e, 0x58, 0x5b, 0xa7, 0xa0, 0x83, 0xef, 0x92, 0x3f, 0x60, 0x98,
	0x5a, 0x3b, 0x50, 0x08, 0x13, 0x85, 0x20, 0x18, 0x85, 0x15, 0x13, 0xf2, 0x22, 0x4f, 0x0e, 0x79,
	0xaf, 0xcc, 0x4f, 0x86, 0x0f, 0x58, 0x2f, 0x77, 0x56, 0x83, 0xf7, 0x54, 0xf8, 0x4d, 0x85, 0xaa,
	0x3a, 0xee, 0x6d, 0xa2, 0x50, 0x89, 0x3f, 0x95, 0xbd, 0x05, 0x11, 0xe6, 0x70, 0xe0, 0x8b, 0x27,
	0x14, 0x7f, 0xcb, 0x39, 0x4a, 0xba, 0xbe, 0xf9, 0xd8, 0xca, 0x68, 0xb3, 0x95, 0xd1, 0xd7, 0x56,
	0x46, 0x6f, 0x3b, 0x59, 0xdb, 0xec, 0x64, 0xed, 0x73, 0x27, 0x6b, 0xf7, 0x17, 0x0b, 0x83, 0x69,
	0x31, 0x1f, 0x6b, 0xbb, 0x8c, 0x97, 0xb6, 0xf0, 0xca, 0xcc, 0x9d, 0x4a, 0xe3, 0xd3, 0x09, 0xac,
	0xf6, 0x47, 0x80, 0xeb, 0x1c, 0xfc, 0xbc, 0x45, 0xbf, 0xff, 0xf2, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0x25, 0xc7, 0x3c, 0x8f, 0x24, 0x02, 0x00, 0x00,
}

func (m *StoredTrade) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoredTrade) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StoredTrade) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Result) > 0 {
		i -= len(m.Result)
		copy(dAtA[i:], m.Result)
		i = encodeVarintStoredTrade(dAtA, i, uint64(len(m.Result)))
		i--
		dAtA[i] = 0x72
	}
	if len(m.TradeData) > 0 {
		i -= len(m.TradeData)
		copy(dAtA[i:], m.TradeData)
		i = encodeVarintStoredTrade(dAtA, i, uint64(len(m.TradeData)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.ProcessDate) > 0 {
		i -= len(m.ProcessDate)
		copy(dAtA[i:], m.ProcessDate)
		i = encodeVarintStoredTrade(dAtA, i, uint64(len(m.ProcessDate)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.UpdateDate) > 0 {
		i -= len(m.UpdateDate)
		copy(dAtA[i:], m.UpdateDate)
		i = encodeVarintStoredTrade(dAtA, i, uint64(len(m.UpdateDate)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.CreateDate) > 0 {
		i -= len(m.CreateDate)
		copy(dAtA[i:], m.CreateDate)
		i = encodeVarintStoredTrade(dAtA, i, uint64(len(m.CreateDate)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Checker) > 0 {
		i -= len(m.Checker)
		copy(dAtA[i:], m.Checker)
		i = encodeVarintStoredTrade(dAtA, i, uint64(len(m.Checker)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Maker) > 0 {
		i -= len(m.Maker)
		copy(dAtA[i:], m.Maker)
		i = encodeVarintStoredTrade(dAtA, i, uint64(len(m.Maker)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintStoredTrade(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ReceiverAddress) > 0 {
		i -= len(m.ReceiverAddress)
		copy(dAtA[i:], m.ReceiverAddress)
		i = encodeVarintStoredTrade(dAtA, i, uint64(len(m.ReceiverAddress)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Quantity) > 0 {
		i -= len(m.Quantity)
		copy(dAtA[i:], m.Quantity)
		i = encodeVarintStoredTrade(dAtA, i, uint64(len(m.Quantity)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Price) > 0 {
		i -= len(m.Price)
		copy(dAtA[i:], m.Price)
		i = encodeVarintStoredTrade(dAtA, i, uint64(len(m.Price)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Coin) > 0 {
		i -= len(m.Coin)
		copy(dAtA[i:], m.Coin)
		i = encodeVarintStoredTrade(dAtA, i, uint64(len(m.Coin)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TradeType) > 0 {
		i -= len(m.TradeType)
		copy(dAtA[i:], m.TradeType)
		i = encodeVarintStoredTrade(dAtA, i, uint64(len(m.TradeType)))
		i--
		dAtA[i] = 0x12
	}
	if m.TradeIndex != 0 {
		i = encodeVarintStoredTrade(dAtA, i, uint64(m.TradeIndex))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintStoredTrade(dAtA []byte, offset int, v uint64) int {
	offset -= sovStoredTrade(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StoredTrade) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TradeIndex != 0 {
		n += 1 + sovStoredTrade(uint64(m.TradeIndex))
	}
	l = len(m.TradeType)
	if l > 0 {
		n += 1 + l + sovStoredTrade(uint64(l))
	}
	l = len(m.Coin)
	if l > 0 {
		n += 1 + l + sovStoredTrade(uint64(l))
	}
	l = len(m.Price)
	if l > 0 {
		n += 1 + l + sovStoredTrade(uint64(l))
	}
	l = len(m.Quantity)
	if l > 0 {
		n += 1 + l + sovStoredTrade(uint64(l))
	}
	l = len(m.ReceiverAddress)
	if l > 0 {
		n += 1 + l + sovStoredTrade(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovStoredTrade(uint64(l))
	}
	l = len(m.Maker)
	if l > 0 {
		n += 1 + l + sovStoredTrade(uint64(l))
	}
	l = len(m.Checker)
	if l > 0 {
		n += 1 + l + sovStoredTrade(uint64(l))
	}
	l = len(m.CreateDate)
	if l > 0 {
		n += 1 + l + sovStoredTrade(uint64(l))
	}
	l = len(m.UpdateDate)
	if l > 0 {
		n += 1 + l + sovStoredTrade(uint64(l))
	}
	l = len(m.ProcessDate)
	if l > 0 {
		n += 1 + l + sovStoredTrade(uint64(l))
	}
	l = len(m.TradeData)
	if l > 0 {
		n += 1 + l + sovStoredTrade(uint64(l))
	}
	l = len(m.Result)
	if l > 0 {
		n += 1 + l + sovStoredTrade(uint64(l))
	}
	return n
}

func sovStoredTrade(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStoredTrade(x uint64) (n int) {
	return sovStoredTrade(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StoredTrade) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStoredTrade
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
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
			return fmt.Errorf("proto: StoredTrade: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoredTrade: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradeIndex", wireType)
			}
			m.TradeIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredTrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TradeIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradeType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredTrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
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
				return ErrInvalidLengthStoredTrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredTrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TradeType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredTrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
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
				return ErrInvalidLengthStoredTrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredTrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredTrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
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
				return ErrInvalidLengthStoredTrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredTrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Price = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredTrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
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
				return ErrInvalidLengthStoredTrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredTrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Quantity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceiverAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredTrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
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
				return ErrInvalidLengthStoredTrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredTrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReceiverAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredTrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
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
				return ErrInvalidLengthStoredTrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredTrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Maker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredTrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
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
				return ErrInvalidLengthStoredTrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredTrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Maker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredTrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
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
				return ErrInvalidLengthStoredTrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredTrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Checker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredTrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
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
				return ErrInvalidLengthStoredTrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredTrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreateDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredTrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
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
				return ErrInvalidLengthStoredTrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredTrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpdateDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProcessDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredTrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
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
				return ErrInvalidLengthStoredTrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredTrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProcessDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradeData", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredTrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
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
				return ErrInvalidLengthStoredTrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredTrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TradeData = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredTrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
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
				return ErrInvalidLengthStoredTrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredTrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStoredTrade(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStoredTrade
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStoredTrade(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStoredTrade
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStoredTrade
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStoredTrade
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthStoredTrade
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStoredTrade
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStoredTrade
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStoredTrade        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStoredTrade          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStoredTrade = fmt.Errorf("proto: unexpected end of group")
)
