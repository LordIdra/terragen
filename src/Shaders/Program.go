package Shaders

import "github.com/go-gl/gl/v4.1-core/gl"

type Program struct {
	ID uint32
}

func NewProgram() Program {
	programID := gl.CreateProgram()
	return Program{programID}
}
