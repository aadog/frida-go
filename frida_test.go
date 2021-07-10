package frida_go

import (
	"fmt"
	"testing"
	"time"
)

func TestDevice(t *testing.T) {
	dm:=DeviceManager_Create()
	defer dm.Close()


}

func TestDevice_OpenChannel(t *testing.T) {
	dm:=DeviceManager_Create()
	d,err:=dm.GetDeviceByType(DeviceType_USB,1000)
	if err!=nil{
		panic(err)
	}
	fmt.Println(d.Description())
	fmt.Println(d.OpenChannel("::1"))
}

func TestAttach(t *testing.T) {
	dm:=DeviceManager_Create()
	d,err:=dm.GetDeviceByType(DeviceType_USB,1000)
	if err!=nil{
		panic(err)
	}
	a,err:=d.FrontmostApplication(NewFrontmostQueryOptions())
	if err!=nil{
	    panic(err)
	}
	if !a.IsValid(){
		panic("没有前台进程")
	}


	ops:=NewSessionOptions()
	ops.PersistTimeout=2000
	session,err:=d.Attach(a.Pid(),ops)
	if err!=nil{
		panic(err)
	}
	fmt.Println(session)
	defer session.Detach()
	fmt.Println(session.PersistTimeout())
	fmt.Println(session.Description())
	scops:=NewScriptOptions()
	scops.Runtime=FRIDA_SCRIPT_RUNTIME_V8
	sc,err:=session.CreateScript(`
	console.log("aaa")
	`,scops)
	if err!=nil{
	   panic(err)
	}
	err=sc.Load()
	if err!=nil{
	    panic(err)
	}
	defer sc.UnLoad()

	fmt.Println(sc)
	//fmt.Println(session.IsDetached())

}

func TestDevice_Spawn(t *testing.T) {
	dm:=DeviceManager_Create()
	d,err:=dm.GetDeviceByType(DeviceType_USB,1000)
	if err!=nil{
		panic(err)
	}
	fmt.Println(d.OpenChannel("127.0.0.1"))
}

func TestDevice_FrontmostApplication(t *testing.T) {
	dm:=DeviceManager_Create()
	d,err:=dm.GetDeviceByType(DeviceType_USB,1000)
	if err!=nil{
		panic(err)
	}
	a,err:=d.FrontmostApplication(NewFrontmostQueryOptions())
	if err!=nil{
	    panic(err)
	}
	fmt.Println(a.Description())
}

func TestApplication(t *testing.T) {
	dm:=DeviceManager_Create()
	d,err:=dm.GetDeviceByType(DeviceType_USB,1000)
	if err!=nil{
	    panic(err)
	}
	ops:=NewApplicationQueryOptions()
	apps,err:=d.EnumerateApplications(ops)
	if err!=nil{
	    panic(err)
	}
	for _, app := range apps {
		fmt.Println(app.Description())
	}
}

func TestDeviceManager_GetDeviceById(t *testing.T) {
	dm:=DeviceManager_Create()
	ls,err:=dm.EnumerateDevices()
	if err!=nil{
	    panic(err)
	}
	for _, l := range ls{
		fmt.Println(l.Description())
	}
	d,err:=dm.GetDeviceByType(DeviceType_REMOTE,100)
	if err!=nil{
	    t.Fatal(err)
	}
	fmt.Println(d.Description())
}

func TestManager(t *testing.T) {
	dm:=DeviceManager_Create()
	defer dm.Close()
	ls,err:=dm.EnumerateDevices()
	if err!=nil{
	    panic(err)
	}
	for _, l := range ls{
		fmt.Println(l.Description())
	}
	dm.OnChanged(func(manager *DeviceManager) {
		fmt.Println("device change:",)
		ls,err:=manager.EnumerateDevices()
		if err!=nil{
		    panic(err)
		}
		for _, l := range ls{
			fmt.Println(l.Description())
		}
	})
	fmt.Println("wait devices change...")
	for{
		time.Sleep(time.Second)
	}

}

func TestDevice_EnumerateProcesses(t *testing.T) {
	mgr:=DeviceManager_Create()
	d,_:=mgr.GetDeviceByType(DeviceType_USB,1000)
	ops:=NewProcessQueryOptions()
	ops.SelectPids=[]int{1,9125}
	pss,_:=d.EnumerateProcesses(ops)
	for _, process := range pss {
		fmt.Println(process.Description())
	}
}