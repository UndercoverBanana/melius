// The GL package handles everything related to drawing in OpenGL, such as creating the window, widgets, etc.
// This package is not meant to be imported, as it could be considered the "backend" of the repository and should be left untouched.
// That being said, the code is full of comments, if for any reason that additional features should be added.
package gl

import "runtime"

func init() {
	// This is needed to arrange that main() runs on main thread.
	runtime.LockOSThread()
}