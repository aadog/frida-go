package cfrida

func Frida_process_get_pid(obj uintptr) int {
	r, _, _ := frida_process_get_pid.Call(obj)
	return int(r)
}

func Frida_process_get_name(obj uintptr) string {
	r, _, _ := frida_process_get_name.Call(obj)
	defer G_free(r)
	return CStrToStr(r)
}