package cfrida



func Frida_script_load_sync(obj uintptr,cancellable uintptr,error uintptr) {
	frida_script_load_sync.Call(obj,cancellable,error)
}

func Frida_script_unload_sync(obj uintptr,cancellable uintptr,error uintptr) {
	frida_script_unload_sync.Call(obj,cancellable,error)
}

func Frida_script_eternalize_sync(obj uintptr,cancellable uintptr,error uintptr) {
	frida_script_eternalize_sync.Call(obj,cancellable,error)
}


