package cfrida



func Frida_script_load_sync(obj uintptr,cancellable uintptr)error{
	gerr:=MakeGError()
	frida_script_load_sync.Call(obj,cancellable,gerr.Input())
	return gerr.ToError()
}

func Frida_script_unload_sync(obj uintptr,cancellable uintptr) error{
	gerr:=MakeGError()
	frida_script_unload_sync.Call(obj,cancellable,gerr.Input())
	return gerr.ToError()
}

func Frida_script_eternalize_sync(obj uintptr,cancellable uintptr) error{
	gerr:=MakeGError()
	frida_script_eternalize_sync.Call(obj,cancellable,gerr.Input())
	return gerr.ToError()
}

func Frida_script_post(obj uintptr,sjson string,data []byte) {
	ptr:=G_bytes_new(data)
	defer G_bytes_unref(ptr)
	frida_script_post.Call(obj,GoStrToCStr(sjson),ptr)
}


func Frida_script_is_destroyed(obj uintptr)bool{
	r,_,_:=frida_script_is_destroyed.Call(obj)
	return r!=0
}