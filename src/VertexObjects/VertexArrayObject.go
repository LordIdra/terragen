package VertexObjects

import "github.com/go-gl/gl/v4.1-core/gl"

type VertexArrayObject struct {
	vaoId uint32
	vboId uint32
	size  int
}

func NewVertexArrayObject() VertexArrayObject {
	buffer := VertexArrayObject{}
	gl.GenVertexArrays(1, &buffer.vaoId)
	gl.BindVertexArray(buffer.vaoId)
	gl.EnableVertexAttribArray(buffer.vaoId)
	gl.GenBuffers(1, &buffer.vboId)
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer.vboId)
	return buffer
}

func (buffer *VertexArrayObject) AddVertexAttribute(attribute VertexAttribute) {
	buffer.Bind()
	gl.VertexAttribPointer(
		attribute.index,
		attribute.size,
		attribute._type,
		attribute.normalized,
		attribute.stride,
		attribute.pointer)
}

func (buffer *VertexArrayObject) Data(data []float32) {
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(data), gl.Ptr(data), gl.STATIC_DRAW)
	buffer.size = len(data)
}

func (buffer *VertexArrayObject) Bind() {
	gl.BindVertexArray(buffer.vaoId)
}

func (buffer *VertexArrayObject) Unbind() {
	gl.BindVertexArray(0)
}

func (buffer *VertexArrayObject) Draw() {
	gl.DrawArrays(gl.TRIANGLES, 0, int32(buffer.size/3))
}
