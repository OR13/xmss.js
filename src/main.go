package main

import (
	"syscall/js"
)

func add(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(args[0].Int() + args[1].Int())
}

func init() {
	// we have to declare our functions in an init func otherwise they aren't
	// available in JS land at the call time.
	js.Global().Set("add", js.FuncOf(add))
}

func main() {
    <-make(chan bool)
}


