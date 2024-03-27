// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Color struct {
	_tab flatbuffers.Table
}

func GetRootAsColor(buf []byte, offset flatbuffers.UOffsetT) *Color {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Color{}
	x.Init(buf, n+offset)
	return x
}

func FinishColorBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsColor(buf []byte, offset flatbuffers.UOffsetT) *Color {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Color{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedColorBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *Color) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Color) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Color) R() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Color) MutateR(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *Color) G() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Color) MutateG(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func (rcv *Color) B() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Color) MutateB(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func (rcv *Color) Alpha() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Color) MutateAlpha(n uint32) bool {
	return rcv._tab.MutateUint32Slot(10, n)
}

func ColorStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func ColorAddR(builder *flatbuffers.Builder, r uint32) {
	builder.PrependUint32Slot(0, r, 0)
}
func ColorAddG(builder *flatbuffers.Builder, g uint32) {
	builder.PrependUint32Slot(1, g, 0)
}
func ColorAddB(builder *flatbuffers.Builder, b uint32) {
	builder.PrependUint32Slot(2, b, 0)
}
func ColorAddAlpha(builder *flatbuffers.Builder, alpha uint32) {
	builder.PrependUint32Slot(3, alpha, 0)
}
func ColorEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
