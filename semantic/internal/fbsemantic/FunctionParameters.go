// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbsemantic

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type FunctionParameters struct {
	_tab flatbuffers.Table
}

func GetRootAsFunctionParameters(buf []byte, offset flatbuffers.UOffsetT) *FunctionParameters {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FunctionParameters{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *FunctionParameters) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FunctionParameters) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *FunctionParameters) Loc(obj *SourceLocation) *SourceLocation {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(SourceLocation)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *FunctionParameters) List(obj *FunctionParameter, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *FunctionParameters) ListLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *FunctionParameters) Pipe(obj *Identifier) *Identifier {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Identifier)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func FunctionParametersStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func FunctionParametersAddLoc(builder *flatbuffers.Builder, loc flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(loc), 0)
}
func FunctionParametersAddList(builder *flatbuffers.Builder, list flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(list), 0)
}
func FunctionParametersStartListVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func FunctionParametersAddPipe(builder *flatbuffers.Builder, pipe flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(pipe), 0)
}
func FunctionParametersEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}