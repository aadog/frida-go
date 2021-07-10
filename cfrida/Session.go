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

func Frida_session_detach_sync(obj uintptr,cancellable uintptr,error uintptr) {
	frida_session_detach_sync.Call(obj,cancellable,error)
}

func Frida_session_create_script_sync(obj uintptr,source string,ops uintptr,cancellable uintptr,error uintptr)uintptr{
	r,_,_:=frida_session_create_script_sync.Call(obj,GoStrToCStr(source),ops,cancellable,error)
	return r
}
func Frida_session_create_script_from_bytes_sync(obj uintptr,source string,ops uintptr,cancellable uintptr,error uintptr)uintptr{
	r,_,_:=frida_session_create_script_from_bytes_sync.Call(obj,GoStrToCStr(source),ops,cancellable,error)
	return r
}
func Frida_session_compile_script_sync(obj uintptr,source string,ops uintptr,cancellable uintptr,error uintptr)uintptr{
	r,_,_:=frida_session_compile_script_sync.Call(obj,GoStrToCStr(source),ops,cancellable,error)
	return r
}

func Frida_session_enable_debugger_sync(obj uintptr,port int,cancellable uintptr,error uintptr){
	frida_session_enable_debugger_sync.Call(obj, uintptr(port),cancellable,error)
}
func Frida_session_disable_debugger_sync(obj uintptr,cancellable uintptr,error uintptr){
	frida_session_disable_debugger_sync.Call(obj, cancellable,error)
}
func Frida_session_setup_peer_connection_sync(obj uintptr,ops uintptr,cancellable uintptr,error uintptr){
	frida_session_setup_peer_connection_sync.Call(obj,ops, cancellable,error)
}

func Frida_relay_new(address,username,password string,kind int)uintptr{
	r,_,_:=frida_relay_new.Call(GoStrToCStr(address),GoStrToCStr(username),GoStrToCStr(password), uintptr(kind))
	return r
}
func Frida_relay_get_address(obj uintptr)string{
	r,_,_:=frida_relay_get_address.Call(obj)
	defer G_free(r)
	return CStrToStr(r)
}