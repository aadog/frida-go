package cfrida

import "unsafe"

func G_signal_connect_data(obj uintptr,event string,c_handle uintptr,data uintptr,destory_data uintptr,connect_flags int)int64{
	r,_,_:=g_signal_connect_data.Call(obj,GoStrToCStr(event),c_handle,data,destory_data, uintptr(connect_flags))
	return int64(r)
}
func G_signal_handler_disconnect(obj uintptr,handle int64){
	g_signal_handler_disconnect.Call(obj, uintptr(handle))
}

func G_object_unref(obj uintptr) {
	g_object_unref.Call(obj)
}
func G_object_ref(obj uintptr) {
	g_object_ref.Call(obj)
}
func G_free(obj uintptr) {
	g_free.Call(obj)
}

func G_strlen(obj uintptr)int{
	r,_,_:=g_strlen.Call(obj)
	return int(r)
}

func G_ref_string_length(obj uintptr)int{
	r,_,_:=g_ref_string_length.Call(obj)
	return int(r)
}

func G_error_get_message(obj uintptr)string{
	r,_,_:=g_error_get_message.Call(obj)
	return CStrToStr(r)
}
func G_error_get_code(obj uintptr)int{
	r,_,_:=g_error_get_code.Call(obj)
	return int(r)
}
func G_error_free(obj uintptr)int{
	r,_,_:=g_error_free.Call(obj)
	return int(r)
}

func G_io_stream_close(obj uintptr,cancellable uintptr,error uintptr)bool{
	r,_,_:=g_io_stream_close.Call(obj,cancellable,error)
	return int(r)!=0
}
func G_input_stream_read_bytes(obj uintptr,count int,cancellable uintptr,error uintptr)uintptr{
	r,_,_:=g_input_stream_read_bytes.Call(obj, uintptr(count),cancellable,error)
	return r
}
func G_output_stream_write_bytes(obj uintptr,gbytes uintptr,cancellable uintptr,error uintptr)int{
	r,_,_:=g_output_stream_write_bytes.Call(obj, gbytes,cancellable,error)
	return int(r)
}
func G_output_stream_write_all(obj uintptr,buf []byte,bytes_writen *int,cancellable uintptr,error uintptr)bool{
	bufptr:=GetBuffPtr(buf)
	r,_,_:=g_output_stream_write_all.Call(obj, bufptr, uintptr(len(buf)),uintptr(unsafe.Pointer(&bytes_writen)),cancellable,error)
	return r!=0
}
func G_input_stream_read_all(obj uintptr,buffer uintptr,count int,bytes_read *int,cancellable uintptr,error uintptr)bool{
	r,_,_:=g_input_stream_read_all.Call(obj,buffer, uintptr(count), uintptr(unsafe.Pointer(&bytes_read)),cancellable,error)
	return r!=0
}

func G_bytes_get_data(obj uintptr,outlen *int)[]byte{
	r,_,_:=g_bytes_get_data.Call(obj,uintptr(unsafe.Pointer(&outlen)))
	return CBytesToBytes(r,*outlen)
}
func G_bytes_get_size(obj uintptr)int{
	r,_,_:=g_bytes_get_size.Call(obj)
	return int(r)
}
func G_bytes_unref(obj uintptr){
	g_bytes_unref.Call(obj)
}
func G_bytes_ref(obj uintptr)uintptr{
	r,_,_:=g_bytes_ref.Call(obj)
	return r
}
func G_bytes_new(buf []byte)uintptr{
	bufptr:=GetBuffPtr(buf)
	r,_,_:=g_bytes_new.Call(bufptr, uintptr(len(buf)))
	return r
}

func G_io_stream_is_closed(obj uintptr)bool{
	r,_,_:=g_io_stream_is_closed.Call(obj)
	return r!=0
}

func G_io_stream_get_input_stream(obj uintptr)uintptr{
	r,_,_:=g_io_stream_get_input_stream.Call(obj)
	return r
}

func G_io_stream_get_output_stream(obj uintptr)uintptr{
	r,_,_:=g_io_stream_get_output_stream.Call(obj)
	return r
}