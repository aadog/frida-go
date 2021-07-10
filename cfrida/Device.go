package cfrida


func Frida_device_get_id(obj uintptr) string {
	r, _, _ := frida_device_get_id.Call(obj)
	defer G_free(r)
	return CStrToStr(r)
}

func Frida_device_get_name(obj uintptr) string {
	r, _, _ := frida_device_get_name.Call(obj)
	defer G_free(r)
	return CStrToStr(r)
}

func Frida_device_get_dtype(obj uintptr) int {
	r, _, _ := frida_device_get_dtype.Call(obj)
	return int(r)
}
func Frida_device_is_lost(obj uintptr) int {
	r, _, _ := frida_device_is_lost.Call(obj)
	return int(r)
}
func Frida_application_query_options_select_identifier(obj uintptr,identifier string)  {
	frida_application_query_options_select_identifier.Call(obj,GoStrToCStr(identifier))
}
func Frida_application_query_options_set_scope(obj uintptr,scope int)  {
	frida_application_query_options_set_scope.Call(obj, uintptr(scope))
}
func Frida_frontmost_query_options_set_scope(obj uintptr,scope int)  {
	frida_frontmost_query_options_set_scope.Call(obj, uintptr(scope))
}
func Frida_device_enumerate_applications_sync(obj uintptr,ops uintptr, _cancellable uintptr, _error uintptr) uintptr {
	r, _, _ := frida_device_enumerate_applications_sync.Call(obj,ops, _cancellable, _error)
	return r
}
func Frida_device_enumerate_processes_sync(obj uintptr,ops uintptr, _cancellable uintptr, _error uintptr) uintptr {
	r, _, _ := frida_device_enumerate_processes_sync.Call(obj,ops, _cancellable, _error)
	return r
}
func Frida_application_list_size(obj uintptr) int{
	r,_,_:=frida_application_list_size.Call(obj)
	return int(r)
}
func Frida_process_list_size(obj uintptr) int{
	r,_,_:=frida_process_list_size.Call(obj)
	return int(r)
}
func Frida_application_list_get(obj uintptr,index int) uintptr{
	r,_,_:=frida_application_list_get.Call(obj, uintptr(index))
	return r
}
func Frida_process_list_get(obj uintptr,index int) uintptr{
	r,_,_:=frida_process_list_get.Call(obj, uintptr(index))
	return r
}
func Frida_application_query_options_new() uintptr {
	r, _, _ := frida_application_query_options_new.Call()
	return r
}

func Frida_frontmost_query_options_new() uintptr {
	r, _, _ := frida_frontmost_query_options_new.Call()
	return r
}

func Frida_spawn_options_new() uintptr {
	r, _, _ := frida_spawn_options_new.Call()
	return r
}

func Frida_device_spawn_sync(obj uintptr,program string,ops uintptr,cancellable uintptr,error uintptr) int {
	r, _, _ := frida_device_spawn_sync.Call(obj,GoStrToCStr(program),ops,cancellable,error)
	return int(r)
}

func Frida_device_attach_sync(obj uintptr,pid int,ops uintptr,cancellable uintptr,error uintptr) uintptr {
	r, _, _ := frida_device_attach_sync.Call(obj, uintptr(pid),ops,cancellable,error)
	return r
}

func Frida_device_get_frontmost_application_sync(obj uintptr,ops uintptr,cancellable uintptr,error uintptr) uintptr {
	r, _, _ := frida_device_get_frontmost_application_sync.Call(obj,ops,cancellable,error)
	return r
}


func Frida_process_query_options_new() uintptr {
	r, _, _ := frida_process_query_options_new.Call()
	return r
}
func Frida_session_options_new() uintptr {
	r, _, _ := frida_session_options_new.Call()
	return r
}
func Frida_script_options_new() uintptr {
	r, _, _ := frida_script_options_new.Call()
	return r
}
func Frida_peer_options_new() uintptr {
	r, _, _ := frida_peer_options_new.Call()
	return r
}
func Frida_script_options_set_name(obj uintptr,name string) uintptr {
	r, _, _ := frida_script_options_set_name.Call(obj,GoStrToCStr(name))
	return r
}
func Frida_script_options_set_runtime(obj uintptr,value int) uintptr {
	r, _, _ := frida_script_options_set_runtime.Call(obj, uintptr(value))
	return r
}
func Frida_peer_options_set_stun_server(obj uintptr,value string) {
	frida_peer_options_set_stun_server.Call(obj,GoStrToCStr(value))
}
func Frida_peer_options_add_relay(obj uintptr,value uintptr) {
	frida_peer_options_add_relay.Call(obj,value)
}

func Frida_device_resume_sync(obj uintptr,pid int,cancellable uintptr,error uintptr) {
	frida_device_resume_sync.Call(obj, uintptr(pid),cancellable,error)
}
func Frida_device_kill_sync(obj uintptr,pid int,cancellable uintptr,error uintptr) {
	frida_device_kill_sync.Call(obj, uintptr(pid),cancellable,error)
}

func Frida_device_get_host_session_sync(obj uintptr,cancellable uintptr,error uintptr)uintptr{
	r,_,_:=frida_device_get_host_session_sync.Call(obj, cancellable,error)
	return r
}

func Frida_device_open_channel_sync(obj uintptr,address string,cancellable uintptr,error uintptr)uintptr{
	r,_,_:=frida_device_open_channel_sync.Call(obj, GoStrToCStr(address),cancellable,error)
	return r
}

func Frida_device_inject_library_file_sync(obj uintptr,pid int,path string,entrypoint string,data []byte,cancellable uintptr,error uintptr)int{
	r,_,_:=frida_device_inject_library_file_sync.Call(obj, uintptr(pid),GoStrToCStr(path),GoStrToCStr(entrypoint),GetBuffPtr(data),cancellable,error)
	return int(r)
}
func Frida_device_inject_library_blob(obj uintptr,pid int,blbo uintptr,entrypoint string,data []byte,cancellable uintptr,error uintptr)int{
	r,_,_:=frida_device_inject_library_blob_sync.Call(obj, uintptr(pid),blbo,GoStrToCStr(entrypoint),GetBuffPtr(data),cancellable,error)
	return int(r)
}


func Frida_process_query_options_select_pid(obj uintptr,pid int)  {
	frida_process_query_options_select_pid.Call(obj, uintptr(pid))
}
func Frida_process_query_options_set_scope(obj uintptr,scope int)  {
	frida_process_query_options_set_scope.Call(obj, uintptr(scope))
}

func Frida_session_options_set_realm(obj uintptr,value int)  {
	frida_session_options_set_realm.Call(obj, uintptr(value))
}

func Frida_session_options_set_persist_timeout(obj uintptr,value int)  {
	frida_session_options_set_persist_timeout.Call(obj, uintptr(value))
}

func Frida_spawn_options_set_argv(obj uintptr,value uintptr,value_length int)  {
	frida_spawn_options_set_argv.Call(obj, value, uintptr(value_length))
}
func Frida_spawn_options_set_envp(obj uintptr,value uintptr,value_length int)  {
	frida_spawn_options_set_envp.Call(obj, value, uintptr(value_length))
}
func Frida_spawn_options_set_env(obj uintptr,value uintptr,value_length int)  {
	frida_spawn_options_set_env.Call(obj, value, uintptr(value_length))
}
func Frida_spawn_options_set_cwd(obj uintptr,cwd string)  {
	frida_spawn_options_set_cwd.Call(obj, GoStrToCStr(cwd))
}
func Frida_spawn_options_set_stdio(obj uintptr,value int32)  {
	frida_spawn_options_set_stdio.Call(obj, uintptr(value))
}