package main

import "fmt"

// ref: https://draveness.me/golang/basic/golang-interface.html

type RPCError struct {
	Code    int64
	Message string
}

func (e *RPCError) Error() string {
	return fmt.Sprintf("%s, code=%d", e.Message, e.Code)
}

func NewRPCError(code int64, msg string) error {
	return &RPCError{ // type check3
		Code:    code,
		Message: msg,
	}
}

func AsErr(err error) error {
	return err
}

func main() {
	var rpcErr error = NewRPCError(400, "unknown err") // type check1
	err := AsErr(rpcErr)                               // type check2
	println(err)
}
