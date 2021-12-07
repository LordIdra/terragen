package Init

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"log"
)

func InitGlfw() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
}

func InitOpenGL() {
	err := gl.Init()
	if err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)
}
