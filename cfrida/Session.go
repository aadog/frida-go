package cfrida

func Frida_session_get_pid(obj uintptr) int {
	r, _, _ := frida_session_get_pid.Call(obj)
	return int(r)
}

func Frida_session_get_persist_timeout(obj uintptr) int {
	r, _, _ := frida_session_get_persist_timeout.Call(obj)
	return int(r)
}
func Frida_session_is_detached(obj uintptr) bool {
	r, _, _ := frida_session_is_detached.Call(obj)
	return r!=0
}

func Frida_session_detach_sync(obj uintptr,cancellable uintptr)error{
	gerr:=MakeGError()
	frida_session_detach_sync.Call(obj,cancellable,gerr.Input())
	return gerr.ToError()
}

func Frida_session_create_script_sync(obj uintptr,source string,ops uintptr,cancellable uintptr)(uintptr,error){
	gerr:=MakeGError()
	r,_,_:=frida_session_create_script_sync.Call(obj,GoStrToCStr(source),ops,cancellable,gerr.Input())
	return r,gerr.ToError()
}
func Frida_session_create_script_from_bytes_sync(obj uintptr,source []byte,ops uintptr,cancellable uintptr)(uintptr,error){
	gerr:=MakeGError()
	bt:=G_bytes_new(source)
	defer G_bytes_unref(bt)
	r,_,_:=frida_session_create_script_from_bytes_sync.Call(obj,bt,ops,cancellable,gerr.Input())
	return r,gerr.ToError()
}
func Frida_session_compile_script_sync(obj uintptr,source string,ops uintptr,cancellable uintptr)([]byte,error){
	gerr:=MakeGError()
	r,_,_:=frida_session_compile_script_sync.Call(obj,GoStrToCStr(source),ops,cancellable,gerr.Input())
	return G_bytes_to_bytes_and_unref(r),gerr.ToError()
}

func Frida_session_enable_debugger_sync(obj uintptr,port int,cancellable uintptr)error{
	gerr:=MakeGError()
	frida_session_enable_debugger_sync.Call(obj, uintptr(port),cancellable,gerr.Input())
	return gerr.ToError()
}
func Frida_session_disable_debugger_sync(obj uintptr,cancellable uintptr)error{
	gerr:=MakeGError()
	frida_session_disable_debugger_sync.Call(obj, cancellable,gerr.Input())
	return gerr.ToError()
}
func Frida_session_setup_peer_connection_sync(obj uintptr,ops uintptr,cancellable uintptr)error{
	gerr:=MakeGError()
	frida_session_setup_peer_connection_sync.Call(obj,ops, cancellable,gerr.Input())
	return gerr.ToError()
}

func Frida_relay_new(address,username,password string,kind int)uintptr{
	r,_,_:=frida_relay_new.Call(GoStrToCStr(address),GoStrToCStr(username),GoStrToCStr(password), uintptr(kind))
	return r
}
func Frida_relay_get_address(obj uintptr)string{
	r,_,_:=frida_relay_get_address.Call(obj)
	return CStrToGoStr(r)
}
