// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ClipRectangle struct {
	_tab flatbuffers.Table
}

func GetRootAsClipRectangle(buf []byte, offset flatbuffers.UOffsetT) *ClipRectangle {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ClipRectangle{}
	x.Init(buf, n+offset)
	return x
}

func FinishClipRectangleBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsClipRectangle(buf []byte, offset flatbuffers.UOffsetT) *ClipRectangle {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ClipRectangle{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedClipRectangleBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *ClipRectangle) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ClipRectangle) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ClipRectangle) X() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *ClipRectangle) MutateX(n float64) bool {
	return rcv._tab.MutateFloat64Slot(4, n)
}

func (rcv *ClipRectangle) Y() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *ClipRectangle) MutateY(n float64) bool {
	return rcv._tab.MutateFloat64Slot(6, n)
}

func (rcv *ClipRectangle) Width() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *ClipRectangle) MutateWidth(n float64) bool {
	return rcv._tab.MutateFloat64Slot(8, n)
}

func (rcv *ClipRectangle) Height() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *ClipRectangle) MutateHeight(n float64) bool {
	return rcv._tab.MutateFloat64Slot(10, n)
}

func ClipRectangleStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func ClipRectangleAddX(builder *flatbuffers.Builder, x float64) {
	builder.PrependFloat64Slot(0, x, 0.0)
}
func ClipRectangleAddY(builder *flatbuffers.Builder, y float64) {
	builder.PrependFloat64Slot(1, y, 0.0)
}
func ClipRectangleAddWidth(builder *flatbuffers.Builder, width float64) {
	builder.PrependFloat64Slot(2, width, 0.0)
}
func ClipRectangleAddHeight(builder *flatbuffers.Builder, height float64) {
	builder.PrependFloat64Slot(3, height, 0.0)
}
func ClipRectangleEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
