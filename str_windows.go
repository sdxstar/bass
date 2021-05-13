package bass

import (
	"github.com/axgle/mahonia"
	"unsafe"
)

func cstr(str string) uintptr {
	return uintptr(unsafe.Pointer(&(utf8ToAnsi(str))[0]))
}

// utf8ToAnsi
func utf8ToAnsi(bs string) []byte {
	decoder := mahonia.NewEncoder("GBK")
	return []byte(decoder.ConvertString(bs))
}
