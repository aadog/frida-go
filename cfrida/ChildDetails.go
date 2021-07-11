package cfrida

func Frida_child_get_pid(obj uintptr) uint {
	r, _, _ := frida_child_get_pid.Call(obj)
	return uint(r)
}
func Frida_child_get_parent_pid(obj uintptr) uint {
	r, _, _ := frida_child_get_parent_pid.Call(obj)
	return uint(r)
}
func Frida_child_get_origin(obj uintptr) int {
	r, _, _ := frida_child_get_origin.Call(obj)
	return int(r)
}

func Frida_child_get_identifier(obj uintptr) string {
	r, _, _ := frida_child_get_identifier.Call(obj)
	return CStrToGoStr(r)
}

func Frida_child_get_path(obj uintptr) string {
	r, _, _ := frida_child_get_path.Call(obj)
	return CStrToGoStr(r)
}

func Frida_child_get_argv(obj uintptr) []string {
	r, _, _ := frida_child_get_argv.Call(obj)
	return G_strv_to_strings(r)
}

func Frida_child_get_envp(obj uintptr) map[string]interface{} {
	r, _, _ := frida_child_get_envp.Call(obj)
	defer G_hash_table_unref(r)
	return G_hash_table_to_Map(r)
}