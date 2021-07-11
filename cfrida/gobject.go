package cfrida

import (
	"errors"
	"fmt"
	"runtime"
	"unsafe"
)

var (
	GvariantStringType=G_variant_type_new("s")
	GvariantInt64Type=G_variant_type_new("x")
	GvariantBooleanType=G_variant_type_new("b")
	GvariantVariantType=G_variant_type_new("v")
	GvariantByteArrayType=G_variant_type_new("ay")
	GvariantVarDictType=G_variant_type_new("{sv}")
	GvariantArrayType=G_variant_type_new("a*")
)



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
func G_object_ref(obj uintptr)uintptr{
	r,_,_:=g_object_ref.Call(obj)
	return r
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
	return CStrToGoStr(r)
}
func G_error_get_code(obj uintptr)int{
	r,_,_:=g_error_get_code.Call(obj)
	return int(r)
}
func G_error_free(obj uintptr)int{
	r,_,_:=g_error_free.Call(obj)
	return int(r)
}

func G_io_stream_close(obj uintptr,cancellable uintptr)(bool,error){
	gerr:=MakeGError()
	r,_,_:=g_io_stream_close.Call(obj,cancellable,gerr.Input())
	return int(r)!=0,gerr.ToError()
}
func G_output_stream_close(obj uintptr,cancellable uintptr)error{
	gerr:=MakeGError()
	g_output_stream_close.Call(obj, cancellable,gerr.Input())
	return gerr.ToError()
}
func G_input_stream_close(obj uintptr,cancellable uintptr)error{
	gerr:=MakeGError()
	g_input_stream_close.Call(obj, cancellable,gerr.Input())
	return gerr.ToError()
}

func G_input_stream_read_bytes(obj uintptr,count int,cancellable uintptr)([]byte,error){
	gerr:=MakeGError()
	r,_,_:=g_input_stream_read_bytes.Call(obj, uintptr(count),cancellable,gerr.Input())
	return G_bytes_to_bytes_and_unref(r),gerr.ToError()
}
func G_output_stream_write_bytes(obj uintptr,data []byte,cancellable uintptr)(int,error){
	gerr:=MakeGError()
	bt:=G_bytes_new(data)
	defer G_bytes_unref(bt)
	r,_,_:=g_output_stream_write_bytes.Call(obj, bt,cancellable,gerr.Input())
	return int(r),gerr.ToError()
}
func G_output_stream_write_all(obj uintptr,buf []byte,bytes_writen *int,cancellable uintptr)(bool,error){
	gerr:=MakeGError()
	bufptr:=GoByteToCPtr(buf)
	r,_,_:=g_output_stream_write_all.Call(obj, bufptr, uintptr(len(buf)),uintptr(unsafe.Pointer(&bytes_writen)),cancellable,gerr.Input())
	return r!=0,gerr.ToError()
}
func G_input_stream_read_all(obj uintptr,buffer []byte,count int,bytes_read *int,cancellable uintptr)(bool,error){
	gerr:=MakeGError()
	bufptr:=GoByteToCPtr(buffer)
	r,_,_:=g_input_stream_read_all.Call(obj,bufptr, uintptr(count), uintptr(unsafe.Pointer(&bytes_read)),cancellable,gerr.Input())
	return r!=0,gerr.ToError()
}

func G_bytes_get_data(obj uintptr)[]byte{
	outlen:=int64(0)
	r,_,_:=g_bytes_get_data.Call(obj,uintptr(unsafe.Pointer(&outlen)))
	if outlen>0{
		return CBytesToGoBytes(r,int(outlen))
	}
	return nil
}
func G_bytes_get_size(obj uintptr)int64{
	r,_,_:=g_bytes_get_size.Call(obj)
	return int64(r)
}
func G_bytes_unref(obj uintptr){
	g_bytes_unref.Call(obj)
}
func G_bytes_ref(obj uintptr)uintptr{
	r,_,_:=g_bytes_ref.Call(obj)
	return r
}
func G_bytes_new(buf []byte)uintptr{
	bufptr:=GoByteToCPtr(buf)
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


func G_bytes_to_bytes(obj uintptr)[]byte{
	if obj>0{
		return G_bytes_get_data(obj)
	}
	return nil
}
func G_bytes_to_bytes_and_unref(obj uintptr)[]byte{
	if obj>0{
		defer G_bytes_unref(obj)
		return G_bytes_get_data(obj)
	}
	return nil
}

func G_hash_table_unref(obj uintptr){
	g_hash_table_unref.Call(obj)
}
func G_hash_table_ref(obj uintptr)uintptr{
	r,_,_:=g_hash_table_ref.Call(obj)
	return r
}
func G_hash_table_iter_init(iter uintptr,hash_table uintptr){
	g_hash_table_iter_init.Call(iter,hash_table)
}
func G_hash_table_iter_next(iter uintptr,key uintptr,value uintptr)bool{
	r,_,_:=g_hash_table_iter_next.Call(iter,key,value)
	return r!=0
}
func G_variant_is_of_type(obj uintptr,gtype uintptr)bool{
	r,_,_:=g_variant_is_of_type.Call(obj,gtype)
	return r!=0
}
func G_variant_get_type_string(obj uintptr)string{
	r,_,_:=g_variant_get_type_string.Call(obj)
	return CStrToGoStr(r)
}
func G_variant_get_string(obj uintptr)string{
	r,_,_:=g_variant_get_string.Call(obj)
	return CStrToGoStr(r)
}
func G_variant_get_int64(obj uintptr)int64{
	r,_,_:=g_variant_get_int64.Call(obj)
	return int64(r)
}
func G_variant_get_boolean(obj uintptr)bool{
	r,_,_:=g_variant_get_boolean.Call(obj)
	return r!=0
}
func G_variant_get_variant(obj uintptr)uintptr{
	r,_,_:=g_variant_get_variant.Call(obj)
	return r
}
func G_variant_get_fixed_array(obj uintptr)[]byte{
	count:=0
	r,_,_:=g_variant_get_fixed_array.Call(obj,uintptr(unsafe.Pointer(&count)),1)
	return CBytesToGoBytes(r,count)
}
func G_variant_iter_next_value(obj uintptr)uintptr{
	r,_,_:=g_variant_iter_next_value.Call(obj)
	return r
}
func G_variant_get_child_value(obj uintptr,index int)uintptr{
	r,_,_:=g_variant_get_child_value.Call(obj, uintptr(index))
	return r
}
func G_variant_unref(obj uintptr){
	g_variant_unref.Call(obj)
}
func G_variant_iter_init(obj uintptr){
	g_variant_iter_init.Call(obj)
}
func G_hash_table_iter_new()uintptr{
	r,_,_:=g_hash_table_iter_new.Call()
	return r
}
func G_hash_table_iter_free(obj uintptr){
	g_hash_table_iter_free.Call(obj)
}
func G_variant_iter_new(val uintptr)uintptr{
	r,_,_:=g_variant_iter_new.Call(val)
	return r
}
func G_variant_iter_free(obj uintptr){
	g_variant_iter_free.Call(obj)
}
func G_variant_type_new(type_string string)uintptr{
	r,_,_:=g_variant_type_new.Call(GoStrToCStr(type_string))
	return r
}
func G_variant_type_free(obj uintptr){
	g_variant_type_free.Call(obj)
}
func G_hash_table_to_Map(hashTablePtr uintptr)map[string]interface{}{
	result:=make(map[string]interface{})
	iterptr:=G_hash_table_iter_new()
	defer G_hash_table_iter_free(iterptr)
	G_hash_table_iter_init(iterptr,hashTablePtr)
	for{
		var rawKey uintptr
		var rawValue uintptr
		if G_hash_table_iter_next(iterptr,uintptr(unsafe.Pointer(&rawKey)),uintptr(unsafe.Pointer(&rawValue)))==false{
			break
		}
		result[CStrToGoStr(rawKey)]=G_valueFromVariant(rawValue)
	}
	return result
}
func G_valueFromVariant(rawPtr uintptr)interface{}{
	if G_variant_is_of_type(rawPtr,GvariantStringType){
		return G_variant_get_string(rawPtr)
	}
	if G_variant_is_of_type(rawPtr,GvariantInt64Type){
		return G_variant_get_int64(rawPtr)
	}
	if G_variant_is_of_type(rawPtr,GvariantBooleanType){
		return G_variant_get_boolean(rawPtr)
	}
	if G_variant_is_of_type(rawPtr,GvariantVariantType){
		return G_valueFromVariant(G_variant_get_variant(rawPtr))
	}
	if G_variant_is_of_type(rawPtr,GvariantByteArrayType){
		return G_variant_get_fixed_array(rawPtr)
	}
	if G_variant_is_of_type(rawPtr,GvariantVarDictType){
		result:=map[string]interface{}{}
		iter:=G_variant_iter_new(rawPtr)
		defer G_variant_iter_free(iter)
		for{
			entry:=G_variant_iter_next_value(iter)
			if entry==0{
				break
			}
			rawKey:=G_variant_get_child_value(rawPtr,0)
			rawValue:=G_variant_get_child_value(rawPtr,1)
			key:=G_variant_get_string(rawKey)
			value:=G_valueFromVariant(rawValue)
			result[key]=value
			G_variant_unref(rawValue)
			G_variant_unref(rawKey)
			G_variant_unref(entry)
		}
		return result
	}
	if G_variant_is_of_type(rawPtr,GvariantArrayType){
		result:=[]interface{}{}
		iter:=G_variant_iter_new(rawPtr)
		defer G_variant_iter_free(iter)
		for{
			child :=G_variant_iter_next_value(iter)
			if child ==0{
				break
			}
			result=append(result,G_valueFromVariant(child))
			G_variant_unref(child)
		}
		return result
	}
	return nil
}

func G_strv_length(obj uintptr)int{
	r,_,_:=g_strv_length.Call(obj)
	return int(r)
}
func G_strv_to_strings(rawPtr uintptr)[]string{
	var result []string
	ptrs:=*(*[]uintptr)(unsafe.Pointer(rawPtr))
	for _, ptr := range ptrs {
		result=append(result,CStrToGoStr(ptr))
	}
	return result
}

type GError struct {
	instance uintptr
}
func (g *GError) Input()uintptr{
	return uintptr(unsafe.Pointer(&g.instance))
}
func (g *GError) ToError()error{
	if g.instance==0{
		return nil
	}
	return errors.New(G_error_get_message(g.instance))
}
func (g *GError) Free(){
	fmt.Println("gerror gc")
	if g.instance!=0{
		G_error_free(g.instance)
		g.instance=0
	}
}
func MakeGError()*GError{
	dl:=new(GError)
	dl.instance=0
	runtime.SetFinalizer(dl, (*GError).Free)
	return dl
}