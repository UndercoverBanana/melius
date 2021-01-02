package gl

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"melius/app"
)

// Initializes GLFW and sets up the Canvas window.
// The Canvas' job is to serve as a transparent background window for artificial windows to be created on
func CreateWindowCanvas() {
	err := glfw.Init()
	if err != nil { panic(err) }
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Maximized, glfw.True)
	glfw.WindowHint(glfw.Resizable, glfw.False)

	glfw.WindowHint(glfw.TransparentFramebuffer, glfw.True)
	glfw.WindowHint(glfw.Decorated, glfw.False)

	videoMode := glfw.GetPrimaryMonitor().GetVideoMode()

	win, err := glfw.CreateWindow(videoMode.Width, videoMode.Height, app.TITLE, nil, nil)
	if err != nil { panic(err) }

	win.MakeContextCurrent()

	for !win.ShouldClose() {
		win.SwapBuffers()
		glfw.WaitEvents()
	}
}