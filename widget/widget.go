// Contains all of the widget information and properties for drawing widgets.
// If needed, it is possible to create a custom widget (reference docs).
package widget

import "github.com/PutterBeanut/melius"

type WidgetStyle struct {

}

// Pay attension to OverrideShader. It might come in handy!
type Widget interface {
	// An overriden function that can be used to draw a widget on a window.
	Draw(width int, height int, x int, y int, ) melius.DrawResult

	// Returns the style of the current widget.
	GetStyle() WidgetStyle

	// Sets the style of the current widget.
	SetStyle(style WidgetStyle)

	// Handles the input for that specific widget.
	HandleInput(mask melius.InputMask) melius.InputResult

	// Overrides the generated shader, allowing the user to have more control.
	// This is highly recommended due to the app not having to generate and cache shaders at runtime, giving a slight performance boost.
	OverrideShader(shader melius.Shader) melius.ShaderOverrideResult

	// Gets the overriden shader.
	// If the shader was not override, it will return empty.
	GetOverridenShader() melius.Shader
}