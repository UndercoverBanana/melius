package style

// An instance of a color using the red, green, blue, and alpha channels.
type Dye struct {
	// The Red, Green, Blue, and Alpha style channels.
	// R, G, and B go from 0 to 255 and A goes from 0 to 1.
	R, G, B, A float32
}

// Base colors

// Base style
func RED() Dye { return Dye{R: 255} }

// Base style
func ORANGE() Dye { return Dye{R: 255, G: 127} }

// Base style
func YELLOW() Dye { return Dye{R: 255, G: 255} }

// Base style
func GREEN() Dye { return Dye{G: 255} }

// Base style
func BLUE() Dye { return Dye{B: 255} }

// Base style
func PURPLE() Dye { return Dye{R: 127, B: 255} }

// Base style
func PINK() Dye { return Dye{R: 255, B: 255} }

// Base style
func WHITE() Dye { return Dye{R: 255, G: 255, B: 255} }

// Base style
func BLACK() Dye { return Dye{R: 0, G: 0, B: 0} }