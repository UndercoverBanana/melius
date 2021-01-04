// Contains all of the widget information and properties for drawing widgets.
// If needed, it is possible to create a custom widget (reference docs).
package widget

import (
	"github.com/PutterBeanut/melius"
	"github.com/PutterBeanut/melius/style"
)

// A struct dedicated purely on how a widget should look.
type Style struct {
	// The overall color of the widget.
	Color style.Dye

	// The overrall color of the text on this widget.
	TextColor style.Dye

	// The border of the widget.
	Border style.Border

	// The shadow of the widget.
	Shadow style.Shadow

	// The font and font settings of the widget.
	Font style.Font
}

// Pay attention to OverrideShader. It might come in handy!
type Widget interface {
	// An overriden function that can be used to draw a widget on a window.
	Draw(width string, height string, x string, y string, ) melius.DrawResult

	// Returns the style of the current widget.
	GetStyle() Style

	// Sets the style of the current widget.
	SetStyle(style Style)

	// Handles the input for that specific widget.
	HandleInput(mask melius.InputMask) melius.InputResult

	// Overrides the generated shader, allowing the user to have more control.
	// This is highly recommended due to the app not having to generate and cache shaders at runtime, giving a slight performance boost.
	OverrideShader(shader melius.Shader) melius.ShaderOverrideResult

	// Gets the overriden shader.
	// If the shader was not override, it will return empty.
	GetOverridenShader() melius.Shader
}

// Allows for custom widgets to be created and registered.
func RegisterWidget(w Widget) {

}