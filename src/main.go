package main

import (
	"Terragen/src/Init"
	"Terragen/src/Shaders"
	"Terragen/src/VertexObjects"
	"Terragen/src/Window"
	"github.com/go-gl/gl/v4.1-core/gl"
	"log"
	"runtime"
)

var program Shaders.Program
var vao VertexObjects.VertexArrayObject

var triangle = []float32{
	0, 0.5, 0, // top
	-0.5, -0.5, 0, // left
	0.5, -0.5, 0, // right
}

func MainLoop() {
	for !Window.ShouldClose() {
		Window.Clear()
		program.Use()

		vao.Bind()
		vao.Draw()

		Window.PollEvents()
		Window.SwapBuffers()
	}
	Window.Terminate()
	log.Println("Window terminated")
}

func main() {
	runtime.LockOSThread()
	Init.InitGlfw()
	Window.SetHints()
	Window.Init()
	Init.InitOpenGL()
	log.Println("GLFW, Window, and OpenGL initialized")

	program = Shaders.NewProgram()
	program.AttachShader("hello", gl.VERTEX_SHADER)
	program.AttachShader("hello", gl.FRAGMENT_SHADER)
	program.Link()
	log.Println("Shaders compiled")

	vao = VertexObjects.NewVertexArrayObject()
	vao.Data(triangle)
	log.Println("VAO created")

	MainLoop()
}
