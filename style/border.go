package style

// An instance of a widget's border attribute.
type Border struct {
	// The radius of the border (how much it curves).
	Radius string

	// The level of detail in the radius.
	// This value is ignored if the Radius is empty.
	RadiusSubdivisions int

	// The thickness of the border.
	Size string

	// The color of the border
	Color Dye
}