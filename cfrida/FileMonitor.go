package cfrida

func Frida_file_monitor_new(path string)uintptr{
	r,_,_:=frida_file_monitor_new.Call(GoStrToCStr(path))
	return r
}

func Frida_file_monitor_enable_sync(obj uintptr,cancellable uintptr,error uintptr){
	frida_file_monitor_enable_sync.Call(obj,cancellable,error)
}
func Frida_file_monitor_disable_sync(obj uintptr,cancellable uintptr,error uintptr){
	frida_file_monitor_disable_sync.Call(obj,cancellable,error)
}
