// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbsemantic

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type FunctionExpression struct {
	_tab flatbuffers.Table
}

func GetRootAsFunctionExpression(buf []byte, offset flatbuffers.UOffsetT) *FunctionExpression {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FunctionExpression{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *FunctionExpression) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FunctionExpression) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *FunctionExpression) Loc(obj *SourceLocation) *SourceLocation {
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

func (rcv *FunctionExpression) Defaults(obj *ObjectExpression) *ObjectExpression {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(ObjectExpression)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *FunctionExpression) Block(obj *FunctionBlock) *FunctionBlock {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(FunctionBlock)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func FunctionExpressionStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func FunctionExpressionAddLoc(builder *flatbuffers.Builder, loc flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(loc), 0)
}
func FunctionExpressionAddDefaults(builder *flatbuffers.Builder, defaults flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(defaults), 0)
}
func FunctionExpressionAddBlock(builder *flatbuffers.Builder, Block flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(Block), 0)
}
func FunctionExpressionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}