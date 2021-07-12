package cfrida

func Frida_init() {
	frida_init.Call()
}

func Frida_shutdown() {
	frida_shutdown.Call()
}

func Frida_deinit() {
	frida_deinit.Call()
}