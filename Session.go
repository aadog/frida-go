package frida_go

import (
"fmt"
"frida-go/cfrida"
"unsafe"
)

const (
	FRIDA_SCRIPT_RUNTIME_DEFAULT = iota
	FRIDA_SCRIPT_RUNTIME_QJS
	FRIDA_SCRIPT_RUNTIME_V8
)

type FridaScriptRuntime int

type Session struct {
	CObj
}

func (p *Session) Pid() int {
	return cfrida.Frida_session_get_pid(p.instance)
}

func (p *Session) PersistTimeout() int {
	return cfrida.Frida_session_get_persist_timeout(p.instance)
}

func (p *Session) IsDetached() bool {
	return cfrida.Frida_session_is_detached(p.instance)
}

func (p *Session) Detach() error {
	var err GError
	cfrida.Frida_session_detach_sync(p.instance,0,err.ErrInput())
	if err.IsError(){
		return err.ToError()
	}
	return nil
}

func (p *Session) CreateScript(source string,ops *ScriptOptions)(*Script,error){
	if ops!=nil{
		if ops.Name!=""{
			cfrida.Frida_script_options_set_name(ops.instance,ops.Name)
		}
		cfrida.Frida_script_options_set_runtime(ops.instance, int(ops.Runtime))
	}
	var err GError
	sc:=cfrida.Frida_session_create_script_sync(p.instance,source,ops.instance,0,err.ErrInput())
	if err.IsError(){
		return nil,err.ToError()
	}
	return ScriptFromInst(sc),nil
}
func (p *Session) CreateScriptFormBytes(source string,ops *ScriptOptions)(*Script,error){
	if ops!=nil{
		if ops.Name!=""{
			cfrida.Frida_script_options_set_name(ops.instance,ops.Name)
		}
		cfrida.Frida_script_options_set_runtime(ops.instance, int(ops.Runtime))
	}
	var err GError
	sc:=cfrida.Frida_session_create_script_from_bytes_sync(p.instance,source,ops.instance,0,err.ErrInput())
	if err.IsError(){
		return nil,err.ToError()
	}
	return ScriptFromInst(sc),nil
}
func (p *Session) CompileScript(source string,ops *ScriptOptions)([]byte,error){
	if ops!=nil{
		if ops.Name!=""{
			cfrida.Frida_script_options_set_name(ops.instance,ops.Name)
		}
		cfrida.Frida_script_options_set_runtime(ops.instance, int(ops.Runtime))
	}
	var err GError
	sc:=cfrida.Frida_session_compile_script_sync(p.instance,source,ops.instance,0,err.ErrInput())
	if err.IsError(){
		return nil,err.ToError()
	}
	return []byte(cfrida.CStrToStr(sc)),nil
}


func (p *Session) EnableDebugger(port int)error{
	var err GError
	cfrida.Frida_session_enable_debugger_sync(p.instance,port,0,err.ErrInput())
	if err.IsError(){
		return err.ToError()
	}
	return nil
}
func (p *Session) DisableDebugger()error{
	var err GError
	cfrida.Frida_session_disable_debugger_sync(p.instance,0,err.ErrInput())
	if err.IsError(){
		return err.ToError()
	}
	return nil
}


func (p *Session) SetupPeerConnection(ops *FridaPeerOptions)error{
	if ops!=nil{
		cfrida.Frida_peer_options_set_stun_server(ops.instance, ops.StunServer)
		for _, relay := range ops.Relays {
			cfrida.Frida_peer_options_add_relay(ops.instance, relay.instance)
		}
	}
	var err GError
	cfrida.Frida_session_setup_peer_connection_sync(p.instance,ops.instance,0,err.ErrInput())
	if err.IsError(){
		return err.ToError()
	}
	return nil
}


func (p *Session) Description() string {
	if p.instance==0{
		return ""
	}
	return fmt.Sprintf("Frida.Session(pid: %d)",p.Pid())
}

func (p *Session) Free() {
	cfrida.G_object_unref(p.instance)
}


// SessionFromInst
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func SessionFromInst(inst uintptr) *Session {
	dl:=new(Session)
	dl.instance=inst
	dl.ptr= unsafe.Pointer(inst)
	setFinalizer(dl, (*Session).Free)
	return dl
}

