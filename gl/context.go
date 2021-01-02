package gl

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

// Initializes GLFW and sets up the Canvas window.
// The Canvas' job is to serve as a transparent background window for artificial windows to be created on
func CreateWindowCanvas(title string) {
	err := glfw.Init()
	if err != nil { panic(err) }
	defer glfw.Terminate()

	// Window hints
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.Decorated, glfw.False)
	glfw.WindowHint(glfw.TransparentFramebuffer, glfw.True)

	win, err := glfw.CreateWindow(100, 100, title, nil, nil)
	if err != nil { panic(err) }

	win.MakeContextCurrent()

	win.Maximize()

	for !win.ShouldClose() {
		win.SwapBuffers()
		glfw.WaitEvents()
	}
}