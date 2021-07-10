package frida_go

import (
	"fmt"
	"frida-go/cfrida"
	"unsafe"
)

const (
	FRIDA_REALM_NATIVE = iota
	FRIDA_REALM_EMULATED
)
type FridaRealm int

const (
	ScopeType_MINIMAL =iota
	ScopeType_METADATA
	ScopeType_FULL
)
type FridaScope int

const (
	DeviceType_LOCAL = iota
	DeviceType_REMOTE
	DeviceType_USB
)
type DeviceType int

type Device struct {
	CObj
}


func (d *Device) Name() string {
	return cfrida.Frida_device_get_name(d.instance)
}
func (d *Device) Id() string {
	return cfrida.Frida_device_get_id(d.instance)
}

func (d *Device) Type() DeviceType {
	return DeviceType(cfrida.Frida_device_get_dtype(d.instance))
}
func (d *Device) IsLost() bool {
	return cfrida.Frida_device_is_lost(d.instance)!=0
}
func (d *Device) Description() string {
	if d.instance==0{
		return ""
	}
	return fmt.Sprintf(`Frida.Device(id: %s, name: %s, kind: %d)`,d.Id(),d.Name(),d.Type())
}
func (d *Device) FrontmostApplication(ops *FrontmostQueryOptions)(*Application,error){
	if ops!=nil{
		cfrida.Frida_frontmost_query_options_set_scope(ops.instance, int(ops.Scope))
	}
	var err GError
	a:=cfrida.Frida_device_get_frontmost_application_sync(d.instance,ops.instance,0,err.ErrInput())
	if err.IsError(){
		return nil,err.ToError()
	}
	return ApplicationFromInst(a),nil
}
func (d *Device) EnumerateApplications(ops *ApplicationQueryOptions) ([]*Application, error) {
	if ops!=nil{
		for _,identifier:=range ops.Identifiers{
			cfrida.Frida_application_query_options_select_identifier(ops.instance,identifier)
		}
		cfrida.Frida_application_query_options_set_scope(ops.instance, int(ops.Scope))
	}
	var err GError
	rawls:=cfrida.Frida_device_enumerate_applications_sync(d.instance,ops.instance,0,err.ErrInput())
	if err.IsError(){
	    return nil,err.ToError()
	}
	defer cfrida.G_object_unref(rawls)
	applications:=make([]*Application,0)
	n:=cfrida.Frida_application_list_size(rawls)
	for i := 0; i <n; i++ {
		applications=append(applications,ApplicationFromInst(cfrida.Frida_application_list_get(rawls,i)))
	}
	return applications,nil
}

func (d *Device) EnumerateProcesses(ops *ProcessQueryOptions) ([]*Process, error) {
	if ops!=nil{
		for _, pid := range ops.SelectPids {
			cfrida.Frida_process_query_options_select_pid(ops.instance,pid)
		}
		cfrida.Frida_process_query_options_set_scope(ops.instance, int(ops.Scope))
	}
	var err GError
	rawls:=cfrida.Frida_device_enumerate_processes_sync(d.instance,ops.instance,0,err.ErrInput())
	if err.IsError(){
		return nil,err.ToError()
	}
	defer cfrida.G_object_unref(rawls)
	processes:=make([]*Process,0)
	n:=cfrida.Frida_process_list_size(rawls)
	for i := 0; i <n; i++ {
		processes=append(processes,ProcessFromInst(cfrida.Frida_process_list_get(rawls,i)))
	}
	return processes,nil
}

func (d *Device) Attach(pid int,ops *SessionOptions)(*Session,error){
	if ops!=nil{
		cfrida.Frida_session_options_set_realm(d.instance, int(ops.Realm))
		if ops.PersistTimeout!=0{
			cfrida.Frida_session_options_set_persist_timeout(d.instance, ops.PersistTimeout)
		}
	}
	var err GError
	rawsession:=cfrida.Frida_device_attach_sync(d.instance,pid,ops.instance,0,err.ErrInput())
	if err.IsError(){
		return nil,err.ToError()
	}
	return SessionFromInst(rawsession),nil
}

func (d *Device) Spawn(program string, ops *SpawnOptions)(int,error){
	if ops!=nil{
		if len(ops.Argv)!=0{
			a,alen:=StrvFromArray(ops.Argv)
			cfrida.Frida_spawn_options_set_argv(ops.instance,a,alen)
		}
		if len(ops.Env)!=0{
			a,alen:=StrvFromArray(ops.Env)
			cfrida.Frida_spawn_options_set_env(ops.instance,a,alen)
		}
		if len(ops.Envp)!=0{
			a,alen:=StrvFromArray(ops.Envp)
			cfrida.Frida_spawn_options_set_envp(ops.instance,a,alen)
		}
		cfrida.Frida_spawn_options_set_stdio(ops.instance, int32(ops.Stdio))
		cfrida.Frida_spawn_options_set_cwd(ops.instance, ops.Cwd)
	}
	var err GError
	pid:=cfrida.Frida_device_spawn_sync(d.instance,program,ops.instance,0,err.ErrInput())
	if err.IsError(){
		return 0,err.ToError()
	}
	return pid,nil
}
func (p *Device) Resume(pid int) error {
	var err GError
	cfrida.Frida_device_resume_sync(p.instance,pid,0,err.ErrInput())
	if err.IsError(){
		return err.ToError()
	}
	return nil
}

func (p *Device) Kill(pid int) error {
	var err GError
	cfrida.Frida_device_kill_sync(p.instance,pid,0,err.ErrInput())
	if err.IsError(){
		return err.ToError()
	}
	return nil
}

func (p *Device) GetHostSession()(*Session,error){
	var err GError
	r:=cfrida.Frida_device_get_host_session_sync(p.instance,0,err.ErrInput())
	if err.IsError(){
		return nil,err.ToError()
	}
	return SessionFromInst(r),nil
}
func (p *Device) OpenChannel(address string)(*IOStream,error){
	var err GError
	r:=cfrida.Frida_device_open_channel_sync(p.instance,address,0,err.ErrInput())
	if err.IsError(){
		return nil,err.ToError()
	}
	return IOStreamFromInst(r),nil
}

func (p *Device) InjectLibraryFileFile(pid int,path string,entrypoint string,data []byte)(int,error){
	var err GError
	r:=cfrida.Frida_device_inject_library_file_sync(p.instance,pid,path,entrypoint,data,0,err.ErrInput())
	if err.IsError(){
		return 0,err.ToError()
	}
	return r,nil
}
func (p *Device) InjectLibraryBlobBlob(pid int,blob []byte,entrypoint string,data []byte)(int,error){
	blobptr:=cfrida.G_bytes_new(blob)
	defer cfrida.G_bytes_unref(blobptr)
	var err GError
	r:=cfrida.Frida_device_inject_library_blob(p.instance,pid,blobptr,entrypoint,data,0,err.ErrInput())
	if err.IsError(){
		return 0,err.ToError()
	}
	return r,nil
}

func (d *Device) Free() {
	cfrida.G_object_unref(d.instance)
}


// DeviceFromInst
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func DeviceFromInst(inst uintptr) *Device {
	dl:=new(Device)
	dl.instance=inst
	dl.ptr= unsafe.Pointer(inst)
	setFinalizer(dl, (*Device).Free)
	return dl
}