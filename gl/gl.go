package gl

import "runtime"

func init() {
	// This is needed to arrange that main() runs on main thread.
	runtime.LockOSThread()
}