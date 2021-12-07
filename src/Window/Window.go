package Window

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

var width = 500
var height = 500
var Window *glfw.Window

func Init() *glfw.Window {
	var err error
	Window, err = glfw.CreateWindow(width, height, "Test", nil, nil)
	if err != nil {
		panic(err)
	}
	Window.MakeContextCurrent()
	return Window
}

func SetHints() {
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
}

func ShouldClose() bool {
	return Window.ShouldClose()
}

func Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func Terminate() {
	glfw.Terminate()
}

func PollEvents() {
	glfw.PollEvents()
}

func SwapBuffers() {
	Window.SwapBuffers()
}
