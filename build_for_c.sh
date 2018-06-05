GOPATH="$(pwd)"
export GOPATH=$GOPATH

go get -v github.com/dq2nd/go_first_dummy_lib
go build -buildmode=c-shared -o go_second_dummy_lib_for_c.a go_second_dummy_lib_for_c.go

