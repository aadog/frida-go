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
	*DeviceSignalConnect
}

func (d *Device) Bus() *Bus {
	bushandle:=cfrida.Frida_device_get_bus(d.instance)
	cfrida.G_object_ref(bushandle)
	return BusFromInst(bushandle)
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
func (d *Device) FrontmostApplication(ops FrontmostQueryOptions)(*ApplicationDetails,error){
	rawops:=cfrida.Frida_frontmost_query_options_new()
	defer cfrida.G_object_unref(rawops)
	cfrida.Frida_frontmost_query_options_set_scope(rawops, int(ops.Scope))
	a,err:=cfrida.Frida_device_get_frontmost_application_sync(d.instance,rawops,0)
	if err!=nil{
	    return nil,err
	}
	return ApplicationFromInst(a),nil
}
func (d *Device) EnumerateApplications(ops ApplicationQueryOptions) ([]*ApplicationDetails, error) {
	rawops:=cfrida.Frida_application_query_options_new()
	defer cfrida.G_object_unref(rawops)
	for _,identifier:=range ops.Identifiers{
		cfrida.Frida_application_query_options_select_identifier(rawops,identifier)
	}
	cfrida.Frida_application_query_options_set_scope(rawops, int(ops.Scope))
	rawls,err:=cfrida.Frida_device_enumerate_applications_sync(d.instance,rawops,0)
	if err!=nil{
	    return nil,err
	}
	defer cfrida.G_object_unref(rawls)
	applications:=make([]*ApplicationDetails,0)
	n:=cfrida.Frida_application_list_size(rawls)
	for i := 0; i <n; i++ {
		applications=append(applications,ApplicationFromInst(cfrida.Frida_application_list_get(rawls,i)))
	}
	return applications,nil
}
func (d *Device) EnumerateProcesses(ops ProcessQueryOptions) ([]*ProcessDetails, error) {
	rawops:=cfrida.Frida_process_query_options_new()
	defer cfrida.G_object_unref(rawops)
	for _, pid := range ops.SelectPids {
		cfrida.Frida_process_query_options_select_pid(rawops,pid)
	}
	cfrida.Frida_process_query_options_set_scope(rawops, int(ops.Scope))

	rawls,err:=cfrida.Frida_device_enumerate_processes_sync(d.instance,rawops,0)
	if err!=nil{
	    return nil,err
	}
	defer cfrida.G_object_unref(rawls)
	processes:=make([]*ProcessDetails,0)
	n:=cfrida.Frida_process_list_size(rawls)
	for i := 0; i <n; i++ {
		processes=append(processes,ProcessFromInst(cfrida.Frida_process_list_get(rawls,i)))
	}
	return processes,nil
}
func (d *Device) EnableSpawnGating(ops ProcessQueryOptions) error {
	err:=cfrida.Frida_device_enable_spawn_gating_sync(d.instance,0)
	if err!=nil{
	    return err
	}
	return nil
}
func (d *Device) DisableSpawnGating(ops ProcessQueryOptions) (error) {
	err:=cfrida.Frida_device_disable_spawn_gating_sync(d.instance,0)
	if err!=nil{
	    return err
	}
	return nil
}
func (d *Device) EnumeratePendingSpawn() ([]*SpawnDetails, error) {

	rawls,err:=cfrida.Frida_device_enumerate_pending_spawn_sync(d.instance,0)
	if err!=nil{
	    return nil,err
	}
	defer cfrida.G_object_unref(rawls)
	spawns:=make([]*SpawnDetails,0)
	n:=cfrida.Frida_spawn_list_size(rawls)
	for i := 0; i <n; i++ {
		spawns=append(spawns,SpawnDetailsFromInst(cfrida.Frida_spawn_list_get(rawls,i)))
	}
	return spawns,nil
}
func (d *Device) EnumeratePendingChildren() ([]*ChildDetails, error) {
	rawls,err:=cfrida.Frida_device_enumerate_pending_children_sync(d.instance,0)
	if err!=nil{
	    return nil,err
	}
	defer cfrida.G_object_unref(rawls)
	children:=make([]*ChildDetails,0)
	n:=cfrida.Frida_child_list_size(rawls)
	for i := 0; i <n; i++ {
		children=append(children,ChildDetailsFromInst(cfrida.Frida_child_list_get(rawls,i)))
	}
	return children,nil
}

func (d *Device) Input(pid uint,data []byte) error {
	err:=cfrida.Frida_device_input_sync(d.instance,pid,data,0)
	return err
}
func (d *Device) Attach(pid uint,ops SessionOptions)(*Session,error){
	rawops:=cfrida.Frida_session_options_new()
	defer cfrida.G_object_unref(rawops)
	cfrida.Frida_session_options_set_realm(rawops, int(ops.Realm))
	if ops.PersistTimeout!=0{
		cfrida.Frida_session_options_set_persist_timeout(rawops, ops.PersistTimeout)
	}
	rawsession,err:=cfrida.Frida_device_attach_sync(d.instance,pid,rawops,0)
	if err!=nil{
	    return nil,err
	}
	return SessionFromInst(rawsession),nil
}
func (d *Device) QuerySystemParameters()(map[string]interface{},error){
	dict,err:=cfrida.Frida_device_query_system_parameters_sync(d.instance,0)
	if err!=nil{
	    return nil,err
	}
	return dict,nil
}
func (d *Device) Spawn(program string, ops SpawnOptions)(uint,error){
	rawops:=cfrida.Frida_spawn_options_new()
	defer cfrida.G_object_unref(rawops)
	if len(ops.Argv)!=0{
		a,alen:=StrvFromArray(ops.Argv)
		cfrida.Frida_spawn_options_set_argv(rawops,a,alen)
	}
	if len(ops.Env)!=0{
		a,alen:=StrvFromArray(ops.Env)
		cfrida.Frida_spawn_options_set_env(rawops,a,alen)
	}
	if len(ops.Envp)!=0{
		a,alen:=StrvFromArray(ops.Envp)
		cfrida.Frida_spawn_options_set_envp(rawops,a,alen)
	}
	cfrida.Frida_spawn_options_set_stdio(rawops, int32(ops.Stdio))
	cfrida.Frida_spawn_options_set_cwd(rawops, ops.Cwd)

	pid,err:=cfrida.Frida_device_spawn_sync(d.instance,program,rawops,0)
	if err!=nil{
	    return 0,err
	}
	return pid,nil
}
func (d *Device) Resume(pid uint) error {

	err:=cfrida.Frida_device_resume_sync(d.instance,pid,0,)
	if err!=nil{
	    return err
	}
	return nil
}

func (d *Device) Kill(pid uint) error {

	err:=cfrida.Frida_device_kill_sync(d.instance,pid,0,)
	if err!=nil{
	    return err
	}
	return nil
}

func (d *Device) GetHostSession()(*Session,error){

	r,err:=cfrida.Frida_device_get_host_session_sync(d.instance,0,)
	if err!=nil{
	    return nil,err
	}
	return SessionFromInst(r),nil
}
func (d *Device) OpenChannel(address string)(*IOStream,error){

	r,err:=cfrida.Frida_device_open_channel_sync(d.instance,address,0,)
	if err!=nil{
	    return nil,err
	}
	return IOStreamFromInst(r),nil
}

func (d *Device) InjectLibraryFileFile(pid uint,path string,entrypoint string,data []byte)(int,error){

	r,err:=cfrida.Frida_device_inject_library_file_sync(d.instance,pid,path,entrypoint,data,0,)
	if err!=nil{
	    return 0,err
	}
	return r,nil
}
func (d *Device) InjectLibraryBlobBlob(pid uint,blob []byte,entrypoint string,data []byte)(int,error){
	blobptr:=cfrida.G_bytes_new(blob)
	defer cfrida.G_bytes_unref(blobptr)

	r,err:=cfrida.Frida_device_inject_library_blob(d.instance,pid,blobptr,entrypoint,data,0)
	if err!=nil{
	    return 0,err
	}
	return r,nil
}

func (d *Device) GetProcessById(pid uint,ops ProcessMatchOptions)(*ProcessDetails,error){
	rawops:=cfrida.Frida_process_match_options_new()
	defer cfrida.G_object_ref(rawops)
	if ops.Timeout!=0{
		cfrida.Frida_process_match_options_set_timeout(rawops,ops.Timeout)
	}
	cfrida.Frida_process_match_options_set_scope(rawops, int(ops.Scope))
	r,err:=cfrida.Frida_device_get_process_by_pid_sync(d.instance,pid,rawops,0)
	if err!=nil{
	    return nil,err
	}
	return ProcessFromInst(r),nil
}
func (d *Device) GetProcessByName(name string,ops ProcessMatchOptions)(*ProcessDetails,error){
	rawops:=cfrida.Frida_process_match_options_new()
	defer cfrida.G_object_ref(rawops)
	if ops.Timeout!=0{
		cfrida.Frida_process_match_options_set_timeout(rawops,ops.Timeout)
	}
	cfrida.Frida_process_match_options_set_scope(rawops, int(ops.Scope))

	r,err:=cfrida.Frida_device_get_process_by_name_sync(d.instance,name,rawops,0)
	if err!=nil{
	    return nil,err
	}
	return ProcessFromInst(r),nil
}

func (d *Device) FindProcessById(pid uint,ops ProcessMatchOptions)(*ProcessDetails,error){
	rawops:=cfrida.Frida_process_match_options_new()
	defer cfrida.G_object_ref(rawops)
	if ops.Timeout!=0{
		cfrida.Frida_process_match_options_set_timeout(rawops,ops.Timeout)
	}
	cfrida.Frida_process_match_options_set_scope(rawops, int(ops.Scope))

	r,err:=cfrida.Frida_device_find_process_by_pid_sync(d.instance,pid,rawops,0)
	if err!=nil{
	    return nil,err
	}
	return ProcessFromInst(r),nil
}
func (d *Device) FindProcessByName(name string,ops ProcessMatchOptions)(*ProcessDetails,error){
	rawops:=cfrida.Frida_process_match_options_new()
	defer cfrida.G_object_ref(rawops)
	if ops.Timeout!=0{
		cfrida.Frida_process_match_options_set_timeout(rawops,ops.Timeout)
	}
	cfrida.Frida_process_match_options_set_scope(rawops, int(ops.Scope))

	r,err:=cfrida.Frida_device_find_process_by_name_sync(d.instance,name,rawops,0)
	if err!=nil{
	    return nil,err
	}
	return ProcessFromInst(r),nil
}



func (d *Device) Free() {
	d.free()
	fmt.Println("device gc")
	cfrida.G_object_unref(d.instance)
}


// DeviceFromInst
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func DeviceFromInst(inst uintptr) *Device {
	dl:=new(Device)
	dl.instance=inst
	dl.ptr= unsafe.Pointer(dl.instance)
	dl.DeviceSignalConnect=NewDeviceSignalConnect(dl.instance)
	setFinalizer(dl, (*Device).Free)
	return dl
}