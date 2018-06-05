package main

import (
    "C"
    "github.com/dq2nd/go_first_dummy_lib"
)

//export Say_hello
func Say_hello() *C.char{
    str := go_first_dummy_lib.Say_hello()
    str += "\n"
    str += "Hello from go_second_dummy_lib!"
    return C.CString(str)
}

func main() {
    // We need the main function to make possible
    // CGO compiler to compile the package as C shared library
}

