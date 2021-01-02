package app

import "github.com/PutterBeanut/melius/gl"

var TITLE = ""

// The App type contains all of the functions and variables needed to edit the application's settings.
//
// NOTE: This struct should only be created once (using the New function).
type App struct {

}

// Creates an instance of the App type.
// Should only be called once.
func New(title string) App {
	TITLE = title

	// Creates the "background window" (or in this case, "Canvas") that all of the windows are drawn on.
	gl.CreateWindowCanvas()

	return App{}
}