package app

import (
	"github.com/PutterBeanut/melius/gl"
	"github.com/PutterBeanut/melius/window"
)

// The App type contains all of the functions and variables needed to edit the application's settings.
//
// NOTE: This struct should only be created once (using the New function).
type App struct {
	// Storage of all of the currently opened windows.
	Windows []window.Window
}

// Creates a window and adds it to the Windows slice.
func(a *App) CreateWindow() *window.Window {
	a.Windows = append(a.Windows, window.Window{})
	return &a.Windows[len(a.Windows) - 1]
}

// Creates an instance of the App type.
// Should only be called once.
func New(title string) App {
	// Creates the "background window" (or in this case, "Canvas") that all of the windows are drawn on.
	gl.CreateWindowCanvas(title)

	return App{}
}