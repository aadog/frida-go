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

func (s *Session) Pid() int {
	return cfrida.Frida_session_get_pid(s.instance)
}

func (s *Session) PersistTimeout() int {
	return cfrida.Frida_session_get_persist_timeout(s.instance)
}

func (s *Session) IsDetached() bool {
	return cfrida.Frida_session_is_detached(s.instance)
}

func (s *Session) Detach() error {
	err:=cfrida.Frida_session_detach_sync(s.instance,0,)
	if err!=nil{
	    return err
	}
	return nil
}

func (s *Session) CreateScript(source string,ops ScriptOptions)(*Script,error){
	rawops:=cfrida.Frida_script_options_new()
	defer cfrida.G_object_unref(rawops)
	if ops.Name!=""{
		cfrida.Frida_script_options_set_name(rawops,ops.Name)
	}
	cfrida.Frida_script_options_set_runtime(rawops, int(ops.Runtime))

	sc,err:=cfrida.Frida_session_create_script_sync(s.instance,source,rawops,0,)
	if err!=nil{
	    return nil,err
	}
	return ScriptFromInst(sc),nil
}
func (s *Session) CreateScriptFormBytes(source []byte,ops ScriptOptions)(*Script,error){
	rawops:=cfrida.Frida_script_options_new()
	defer cfrida.G_object_unref(rawops)
	if ops.Name!=""{
		cfrida.Frida_script_options_set_name(rawops,ops.Name)
	}
	cfrida.Frida_script_options_set_runtime(rawops, int(ops.Runtime))

	sc,err:=cfrida.Frida_session_create_script_from_bytes_sync(s.instance,source,rawops,0)
	if err!=nil{
	    return nil,err
	}
	return ScriptFromInst(sc),nil
}
func (s *Session) CompileScript(source string,ops ScriptOptions)([]byte,error){
	rawops:=cfrida.Frida_script_options_new()
	defer cfrida.G_object_unref(rawops)
	if ops.Name!=""{
		cfrida.Frida_script_options_set_name(rawops,ops.Name)
	}
	cfrida.Frida_script_options_set_runtime(rawops, int(ops.Runtime))
	bt,err:=cfrida.Frida_session_compile_script_sync(s.instance,source,rawops,0)
	if err!=nil{
	    return nil,err
	}
	return bt,nil
}


func (s *Session) EnableDebugger(port int)error{
	err:=cfrida.Frida_session_enable_debugger_sync(s.instance,port,0)
	if err!=nil{
	    return err
	}
	return nil
}
func (s *Session) DisableDebugger()error{
	err:=cfrida.Frida_session_disable_debugger_sync(s.instance,0)
	if err!=nil{
	    return err
	}
	return nil
}
func (s *Session) SetupPeerConnection(ops PeerOptions)error{
	rawops:=cfrida.Frida_peer_options_new()
	defer cfrida.G_object_unref(rawops)
	cfrida.Frida_peer_options_set_stun_server(rawops, ops.StunServer)
	for _, relay := range ops.Relays {
		cfrida.Frida_peer_options_add_relay(rawops, relay.instance)
	}
	err:=cfrida.Frida_session_setup_peer_connection_sync(s.instance,rawops,0)
	if err!=nil{
	    return err
	}
	return nil
}


func (s *Session) Description() string {
	if s.instance==0{
		return ""
	}
	return fmt.Sprintf("Frida.Session(pid: %d)",s.Pid())
}

func (s *Session) Free() {
	cfrida.G_object_unref(s.instance)
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

