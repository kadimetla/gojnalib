package main

import "C"
import (
	"fmt"
	"reflect"
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

func ProcessStrings(strings **C.char, count C.int) {
	// Convert the C array of strings into a Go slice of strings
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(strings)),
		Len:  int(count),
		Cap:  int(count),
	}
	goStrings := *(*[]*C.char)(unsafe.Pointer(&hdr))

	// Now you can range over goStrings and use C.GoString to convert each *C.char to a Go string
	for _, str := range goStrings {
		goStr := C.GoString(str)
		fmt.Println(goStr)
	}
}

func main() {
}
