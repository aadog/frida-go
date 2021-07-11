package cfrida

func Frida_device_get_id(obj uintptr) string {
	r, _, _ := frida_device_get_id.Call(obj)
	return CStrToGoStr(r)
}

func Frida_device_get_bus(obj uintptr) uintptr {
	r, _, _ := frida_device_get_bus.Call(obj)
	return r
}

func Frida_device_get_name(obj uintptr) string {
	r, _, _ := frida_device_get_name.Call(obj)
	return CStrToGoStr(r)
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
func Frida_device_enumerate_applications_sync(obj uintptr,ops uintptr, _cancellable uintptr) (uintptr,error) {
	gerr:=MakeGError()
	r, _, _ := frida_device_enumerate_applications_sync.Call(obj,ops, _cancellable, gerr.Input())
	return r,gerr.ToError()
}
func Frida_device_enumerate_processes_sync(obj uintptr,ops uintptr, _cancellable uintptr) (uintptr,error) {
	gerr:=MakeGError()
	r, _, _ := frida_device_enumerate_processes_sync.Call(obj,ops, _cancellable, gerr.Input())
	return r,gerr.ToError()
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

func Frida_device_spawn_sync(obj uintptr,program string,ops uintptr,cancellable uintptr) (uint,error) {
	gerr:=MakeGError()
	r, _, _ := frida_device_spawn_sync.Call(obj,GoStrToCStr(program),ops,cancellable,gerr.Input())
	return uint(r),gerr.ToError()
}
func Frida_device_input_sync(obj uintptr,pid uint,data []byte,cancellable uintptr) (error) {
	gerr:=MakeGError()
	dataptr:=G_bytes_new(data)
	defer G_bytes_unref(dataptr)
	frida_device_input_sync.Call(obj, uintptr(pid),dataptr,cancellable,gerr.Input())
	return gerr.ToError()
}
func Frida_device_attach_sync(obj uintptr,pid uint,ops uintptr,cancellable uintptr) (uintptr,error) {
	gerr:=MakeGError()
	r, _, _ := frida_device_attach_sync.Call(obj, uintptr(pid),ops,cancellable,gerr.Input())
	return r,gerr.ToError()
}
func Frida_device_query_system_parameters_sync(obj uintptr,cancellable uintptr) (map[string]interface{},error) {
	gerr:=MakeGError()
	r, _, _ := frida_device_query_system_parameters_sync.Call(obj,cancellable,gerr.Input())
	defer G_hash_table_unref(r)
	dict:=G_hash_table_to_Map(r)
	return dict,gerr.ToError()
}
func Frida_device_get_frontmost_application_sync(obj uintptr,ops uintptr,cancellable uintptr) (uintptr,error) {
	gerr:=MakeGError()
	r, _, _ := frida_device_get_frontmost_application_sync.Call(obj,ops,cancellable,gerr.Input())
	return r,gerr.ToError()
}


func Frida_process_query_options_new() uintptr {
	r, _, _ := frida_process_query_options_new.Call()
	return r
}
func Frida_process_match_options_new() uintptr {
	r, _, _ := frida_process_match_options_new.Call()
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

func Frida_device_resume_sync(obj uintptr,pid uint,cancellable uintptr)error{
	gerr:=MakeGError()
	frida_device_resume_sync.Call(obj, uintptr(pid),cancellable,gerr.Input())
	return gerr.ToError()
}
func Frida_device_kill_sync(obj uintptr,pid uint,cancellable uintptr) error {
	gerr:=MakeGError()
	frida_device_kill_sync.Call(obj, uintptr(pid),cancellable,gerr.Input())
	return gerr.ToError()
}

func Frida_device_get_host_session_sync(obj uintptr,cancellable uintptr)(uintptr,error){
	gerr:=MakeGError()
	r,_,_:=frida_device_get_host_session_sync.Call(obj, cancellable,gerr.Input())
	return r,gerr.ToError()
}

func Frida_device_open_channel_sync(obj uintptr,address string,cancellable uintptr)(uintptr,error){
	gerr:=MakeGError()
	r,_,_:=frida_device_open_channel_sync.Call(obj, GoStrToCStr(address),cancellable,gerr.Input())
	return r,gerr.ToError()
}

func Frida_device_inject_library_file_sync(obj uintptr,pid uint,path string,entrypoint string,data []byte,cancellable uintptr)(int,error){
	gerr:=MakeGError()
	r,_,_:=frida_device_inject_library_file_sync.Call(obj, uintptr(pid),GoStrToCStr(path),GoStrToCStr(entrypoint),GoByteToCPtr(data),cancellable,gerr.Input())
	return int(r),gerr.ToError()
}
func Frida_device_inject_library_blob(obj uintptr,pid uint,blbo uintptr,entrypoint string,data []byte,cancellable uintptr)(int,error){
	gerr:=MakeGError()
	r,_,_:=frida_device_inject_library_blob_sync.Call(obj, uintptr(pid),blbo,GoStrToCStr(entrypoint),GoByteToCPtr(data),cancellable,gerr.Input())
	return int(r),gerr.ToError()
}


func Frida_process_query_options_select_pid(obj uintptr,pid uint)  {
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

func Frida_device_get_process_by_pid_sync(obj uintptr,pid uint,ops uintptr,cancellable uintptr)(uintptr,error){
	gerr:=MakeGError()
	r,_,_:=frida_device_get_process_by_pid_sync.Call(obj, uintptr(pid),ops,cancellable,gerr.Input())
	return r,gerr.ToError()
}
func Frida_device_get_process_by_name_sync(obj uintptr,name string,ops uintptr,cancellable uintptr)(uintptr,error){
	gerr:=MakeGError()
	r,_,_:=frida_device_get_process_by_name_sync.Call(obj, GoStrToCStr(name),ops,cancellable,gerr.Input())
	return r,gerr.ToError()
}

func Frida_device_find_process_by_pid_sync(obj uintptr,pid uint,ops uintptr,cancellable uintptr)(uintptr,error){
	gerr:=MakeGError()
	r,_,_:=frida_device_find_process_by_pid_sync.Call(obj, uintptr(pid),ops,cancellable,gerr.Input())
	return r,gerr.ToError()
}
func Frida_device_find_process_by_name_sync(obj uintptr,name string,ops uintptr,cancellable uintptr)(uintptr,error){
	gerr:=MakeGError()
	r,_,_:=frida_device_find_process_by_name_sync.Call(obj, GoStrToCStr(name),ops,cancellable,gerr.Input())
	return r,gerr.ToError()
}
func Frida_process_match_options_set_timeout(obj uintptr,val int)  {
	frida_process_match_options_set_timeout.Call(obj, uintptr(val))
}
func Frida_process_match_options_set_scope(obj uintptr,val int)  {
	frida_process_match_options_set_scope.Call(obj, uintptr(val))
}
func Frida_device_enable_spawn_gating_sync(obj uintptr,cancellable uintptr)error{
	gerr:=MakeGError()
	frida_device_enable_spawn_gating_sync.Call(obj, cancellable,gerr.Input())
	return gerr.ToError()
}
func Frida_device_disable_spawn_gating_sync(obj uintptr,cancellable uintptr,) error{
	gerr:=MakeGError()
	frida_device_disable_spawn_gating_sync.Call(obj, cancellable,gerr.Input())
	return gerr.ToError()
}
func Frida_device_enumerate_pending_spawn_sync(obj uintptr,cancellable uintptr,)(uintptr,error){
	gerr:=MakeGError()
	r,_,_:=frida_device_enumerate_pending_spawn_sync.Call(obj, cancellable,gerr.Input())
	return r,gerr.ToError()
}
func Frida_spawn_list_get(obj uintptr,index int) uintptr{
	r,_,_:=frida_spawn_list_get.Call(obj, uintptr(index))
	return r
}
func Frida_device_enumerate_pending_children_sync(obj uintptr,cancellable uintptr,)(uintptr,error){
	gerr:=MakeGError()
	r,_,_:=frida_device_enumerate_pending_children_sync.Call(obj, cancellable,gerr.Input())
	return r,gerr.ToError()
}
func Frida_spawn_list_size(obj uintptr)int{
	r,_,_:=frida_spawn_list_size.Call(obj)
	return int(r)
}
func Frida_child_list_size(obj uintptr)int{
	r,_,_:=frida_child_list_size.Call(obj)
	return int(r)
}

func Frida_child_list_get(obj uintptr,index int) uintptr{
	r,_,_:=frida_child_list_get.Call(obj, uintptr(index))
	return r
}