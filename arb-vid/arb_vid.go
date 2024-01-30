package arb_vid

/*
#cgo LDFLAGS: -L./lib -lvid
#include "./lib/vid.h"
*/
import "C"

func Test() {
	C.mock_crypto(C.CString("running some Rust crypto from Go"))

}
