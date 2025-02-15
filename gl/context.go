package gl

import (
	"github.com/go-gl/gl/v2.1/gl"
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

	videoMode := glfw.GetPrimaryMonitor().GetVideoMode()

	win, err := glfw.CreateWindow(videoMode.Width, videoMode.Height, title, nil, nil)
	if err != nil { panic(err) }

	win.MakeContextCurrent()

	win.Maximize()

	err = gl.Init()
	if err != nil { panic(err) }

	for !win.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)

		gl.Begin(gl.TRIANGLES)
		gl.Vertex2f(-0.5, -0.5)
		gl.Vertex2f(+0.5, -0.5)
		gl.Vertex2f(+0.0, +0.5)
		gl.End()

		win.SwapBuffers()
		glfw.WaitEvents()
	}
}