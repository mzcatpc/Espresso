package main

// NOTE: There should be NO space between the comments and the `import "C"` line.

/*
#cgo LDFLAGS: -L./lib -lvid
#include "./lib/vid.h"
*/
import "C"

func main() {
	C.mock_crypto(C.CString("running some Rust crypto from Go"))
}
