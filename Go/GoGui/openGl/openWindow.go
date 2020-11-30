package main

import (
	//"fyne.io/fyne/internal/painter/gl"
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.6-core/gl"

	"github.com/raifpy/Go/errHandler"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func main() {
	runtime.LockOSThread()

	defer glfw.Terminate()
	window := win()
	prog := initopengl()
	draw(window, prog)

	for !window.ShouldClose() {

	}

}

func win() *glfw.Window {
	glfw.Init()
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	win, err := glfw.CreateWindow(100, 100, "Hey Gl", nil, nil)
	errHandler.HandlerExit(err)

	win.MakeContextCurrent()
	return win

}

func initopengl() uint32 {
	gl.Init()

	vers := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("Gl Version : ", vers)
	prog := gl.CreateProgram()
	gl.LinkProgram(prog)
	return prog
}
func draw(win *glfw.Window, prog uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(prog)
	glfw.PollEvents()
	win.SwapBuffers()
}
