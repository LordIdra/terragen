package VertexObjects

import "unsafe"

type VertexAttribute struct {
	index      uint32
	size       int32
	_type      uint32
	normalized bool
	stride     int32
	pointer    unsafe.Pointer //this seems like a great idea
}
