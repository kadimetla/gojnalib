package main

import "C"
import (
	"fmt"
	"unsafe"
)

//export JNAWrapper
func JNAWrapper(cStrings **C.char, length C.int, apiRawInput *C.char) {

	lengthGo := int(length)
	slice := make([]string, lengthGo)
	for i := 0; i < lengthGo; i++ {
		slice[i] = C.GoString(*(**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(cStrings)) + uintptr(i)*unsafe.Sizeof(cStrings))))
		fmt.Println(slice[i])
	}
	fmt.Println(C.GoString(apiRawInput))
}

func CountStrings(cStrings **C.char, length C.int) C.int {
	// Convert C array of string pointers to Go slice of strings
	lengthGo := int(length)
	slice := make([]string, lengthGo)
	for i := 0; i < lengthGo; i++ {
		slice[i] = C.GoString(*(**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(cStrings)) + uintptr(i)*unsafe.Sizeof(cStrings))))
	}

	// Use the slice as needed
	return C.int(len(slice))
}

func main() {
}
