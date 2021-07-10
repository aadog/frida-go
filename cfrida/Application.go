package cfrida


func Frida_application_get_identifier(obj uintptr) string {
	r, _, _ := frida_application_get_identifier.Call(obj)
	defer G_free(r)
	return CStrToStr(r)
}

func Frida_application_get_name(obj uintptr) string {
	r, _, _ := frida_application_get_name.Call(obj)
	defer G_free(r)
	return CStrToStr(r)
}

func Frida_application_get_pid(obj uintptr) int {
	r, _, _ := frida_application_get_pid.Call(obj)
	return int(r)
}

func Frida_application_get_parameters(obj uintptr) int {
	r, _, _ := frida_application_get_parameters.Call(obj)
	return int(r)
}
