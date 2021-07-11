package cfrida



func Frida_bus_is_detached(obj uintptr)bool{
	r,_,_:=frida_bus_is_detached.Call(obj)
	return r!=0
}

func Frida_bus_attach_sync(obj uintptr,cancellable uintptr)(bool,error){
	gerr:=MakeGError()
	r,_,_:=frida_bus_attach_sync.Call(obj,cancellable,gerr.Input())
	return r!=0,gerr.ToError()
}
func Frida_bus_post(obj uintptr,sjson string,data []byte){
	dataptr:=G_bytes_new(data)
	defer G_bytes_unref(dataptr)
	frida_bus_post.Call(obj,GoStrToCStr(sjson),dataptr)
}