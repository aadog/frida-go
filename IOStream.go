package frida_go

import (
"fmt"
"github.com/a97077088/frida-go/cfrida"
"unsafe"
)

type IOStream struct {
	CObj
	Input uintptr
	OutPut uintptr
}

func (i *IOStream) IsClosed() bool {
	return cfrida.G_io_stream_is_closed(i.instance)
}

func (i *IOStream) Close() (bool,error) {
	err:=cfrida.G_input_stream_close(i.Input,0)
	if err!=nil{
	    return false,err
	}
	err=cfrida.G_output_stream_close(i.OutPut,0)
	if err!=nil{
	    return false,err
	}
	b,err:=cfrida.G_io_stream_close(i.instance,0)
	if err!=nil{
	    return false,err
	}
	return b,nil
}
func (i *IOStream) Read(count int) ([]byte,error) {
	bt,err:=cfrida.G_input_stream_read_bytes(i.instance,count,0)
	if err!=nil{
	    return nil,err
	}
	return bt,nil
}
func (i *IOStream) ReadAll(count int) ([]byte,int,error) {
	buf:=make([]byte,count)
	bytes_read:=0
	_,err:=cfrida.G_input_stream_read_all(i.instance,buf,count,&bytes_read,0)
	if err!=nil{
	    return nil,0,err
	}
	return buf,bytes_read,nil
}

func (i *IOStream) Write(data []byte) (int,error) {
	n,err:=cfrida.G_output_stream_write_bytes(i.OutPut,data,0)
	if err!=nil{
	    return 0,err
	}
	return n,nil
}
func (i *IOStream) WriteAll(data []byte) (int,error) {
	outsize:=0
	_,err:=cfrida.G_output_stream_write_all(i.OutPut,data,&outsize,0,)
	if err!=nil{
	    return 0,err
	}
	return outsize,nil
}

func (p *IOStream) Description() string {
	if p.instance==0{
		return ""
	}
	return fmt.Sprintf("Frida.IOStream()")
}

func (p *IOStream) Free() {
	cfrida.G_object_unref(p.instance)
}


// IOStreamFromInst
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func IOStreamFromInst(inst uintptr) *IOStream {
	dl:=new(IOStream)
	dl.instance=inst
	dl.ptr= unsafe.Pointer(inst)
	dl.Input=cfrida.G_io_stream_get_input_stream(dl.instance)
	dl.OutPut=cfrida.G_io_stream_get_output_stream(dl.instance)
	setFinalizer(dl, (*IOStream).Free)
	return dl
}
