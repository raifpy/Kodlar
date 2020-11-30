package main

import (
	"fmt"
	"os"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func main() {
	if glfw.Init() != nil {
		fmt.Fprintln(os.Stderr, "Init edilemedi! : ")
		return
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, gl.TRUE)

	//glfw.WindowHint(glfw.Resizable, glfw.False) // Bu arada i3 arayüzünde resizable false ise regular floation kuralarına uymuyor

	window, err := glfw.CreateWindow(400, 200, "Merhaba Dünya!", nil, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// opengl kullanabilmek için aşağıdakini yapıyoruz;
	//glfw.GetCurrentContext().MakeContextCurrent()
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return

	}

	/*
		go func() {
			for range time.NewTicker(time.Second * 5).C {
				window.Focus()
				//her 5 saniye'de bir pencereye odaklanmayı sağlıyor :)
			}
		}()
	*/

	for !window.ShouldClose() {

		if window.GetKey(glfw.KeyEscape) == glfw.Press {
			fmt.Println("ESC tuşuna basıldı!")
			return
		}

		window.SwapBuffers()
		glfw.PollEvents()

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	}

	fmt.Println("Program kapatılıyor!")

}
