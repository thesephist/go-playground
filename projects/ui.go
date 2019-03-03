package main

import (
	"fmt"
	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/golang-ui/nuklear/nk"
	"log"
)

const (
	winWidth  = 500
	winHeight = 500
)

func main() {

	fmt.Println("HI")

	if err := glfw.Init(); err != nil {
		fmt.Println("Init failed")
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	win, err := glfw.CreateWindow(winWidth, winHeight, "Nuklear Demo", nil, nil)
	if err != nil {
		fmt.Println("Error creating window")
	}
	win.MakeContextCurrent()

	width, height := win.GetSize()
	log.Printf("glfw: created window %dx%d", width, height)

	if err := gl.Init(); err != nil {
		fmt.Println("Opengl: init failed", err)
	}
	gl.Viewport(0, 0, int32(width), int32(height))

	_ = nk.NkPlatformInit(win, nk.PlatformInstallCallbacks)

	ch := make(chan int)
	ch <- 3
}
