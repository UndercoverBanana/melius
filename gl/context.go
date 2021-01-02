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
	glfw.WindowHint(glfw.TransparentFramebuffer, glfw.True)
	glfw.WindowHint(glfw.Decorated, glfw.False)

	// videoMode := glfw.GetPrimaryMonitor().GetVideoMode()

	win, err := glfw.CreateWindow(600, 800, title, nil, nil)
	if err != nil { panic(err) }

	win.SetPos(100, 100)

	win.MakeContextCurrent()

	for !win.ShouldClose() {
		win.SwapBuffers()
		glfw.WaitEvents()
	}
}