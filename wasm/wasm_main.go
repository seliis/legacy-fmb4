package main

import (
	"syscall/js"
)

func moveLanding(this js.Value, arg []js.Value) interface{} {
	js.Global().Get("window").Get("location").Set(
		"href", "/",
	)
	return nil
}

func controlNavs(this js.Value, arg []js.Value) interface{} {
	menu := js.Global().Get("document").Call(
		"querySelector", "nav",
	)
	menu.Get("classList").Call(
		"toggle", "visible",
	)
	return nil
}

func addrChanger(this js.Value, arg []js.Value) interface{} {
	js.Global().Get("window").Get("location").Set(
		"href", arg[0],
	)
	return nil
}

func main() {
	js.Global().Set("moveLanding", js.FuncOf(moveLanding))
	js.Global().Set("controlNavs", js.FuncOf(controlNavs))
	js.Global().Set("addrChanger", js.FuncOf(addrChanger))

	println("Miho's Web Assembly Loaded")
	<-make(chan struct{})
	select {}
}
