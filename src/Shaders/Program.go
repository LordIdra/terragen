package Shaders

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"log"
	"strings"
)

type Program struct {
	id uint32
}

func NewProgram() Program {
	programID := gl.CreateProgram()
	gl.LinkProgram(programID)
	return Program{programID}
}

func (program *Program) Use() {
	gl.UseProgram(program.id)
}

func (program *Program) Link() {
	gl.LinkProgram(program.id)
}

func (program *Program) AttachShader(path string, shaderType uint32) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()

	var status int32
	gl.CompileShader(shader)
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)

	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)
		logText := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(logText))
		log.Fatalf("failed to compile %v: %v", source, logText)
	}

	gl.AttachShader(program.id, shaderType)
}
