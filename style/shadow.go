package style

import "github.com/PutterBeanut/melius"

type Shadow struct {
	// The position of the shadow relative to the center of the widget.
	Offset melius.Vector2

	// This variable ultimately controls how hard the shadow is.
	// Zero is complete hardness and higher values progressively increase the shadow softness.
	Height float32
}