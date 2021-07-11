package cfrida

func Frida_process_get_pid(obj uintptr) uint {
	r, _, _ := frida_process_get_pid.Call(obj)
	return uint(r)
}

func Frida_process_get_name(obj uintptr) string {
	r, _, _ := frida_process_get_name.Call(obj)
	return CStrToGoStr(r)
}
func Frida_process_get_parameters(obj uintptr) map[string]interface{} {
	r, _, _ := frida_process_get_parameters.Call(obj)
	defer G_hash_table_unref(r)
	dict:=G_hash_table_to_Map(r)
	return dict
}