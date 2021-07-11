package cfrida


func Frida_spawn_get_pid(obj uintptr) uint {
	r, _, _ := frida_spawn_get_pid.Call(obj)
	return uint(r)
}

func Frida_spawn_get_identifier(obj uintptr) string {
	r, _, _ := frida_spawn_get_identifier.Call(obj)
	return CStrToGoStr(r)
}