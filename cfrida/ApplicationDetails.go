package cfrida


func Frida_application_get_identifier(obj uintptr) string {
	r, _, _ := frida_application_get_identifier.Call(obj)
	return CStrToGoStr(r)
}

func Frida_application_get_name(obj uintptr) string {
	r, _, _ := frida_application_get_name.Call(obj)
	return CStrToGoStr(r)
}

func Frida_application_get_pid(obj uintptr) int {
	r, _, _ := frida_application_get_pid.Call(obj)
	return int(r)
}

func Frida_application_get_parameters(obj uintptr) map[string]interface{} {
	r, _, _ := frida_application_get_parameters.Call(obj)
	defer G_hash_table_unref(r)
	return G_hash_table_to_Map(r)
}
