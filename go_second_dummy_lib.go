package go_second_dummy_lib

import (
    "github.com/dq2nd/go_first_dummy_lib"
)

func Say_hello() string {
    str := go_first_dummy_lib.Say_hello()
    str += "\n"
    str += "Hello from go_second_dummy_lib!"
    return str
}
