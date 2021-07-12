package cfrida

import (
	"runtime"
	"syscall"
)

var(
	libfrida=loadUILib()
)

// 加载库
func loadUILib() *syscall.LazyDLL {
	libName := getDLLName()
	// 如果支持运行时释放，则使用此种方法
	if support, newDLLPath := checkAndReleaseDLL(); support {
		libName = newDLLPath
	} else {
		libName=libName
	}
	lib := syscall.NewLazyDLL(libName)
	err := lib.Load()
	if err != nil {
		panic(err)
	}

	return lib
}

// 获取dll库实例，用于在外扩展第三方组件的。移动来自dfuncs.go
func GetLibFrida() *syscall.LazyDLL {
	return libfrida
}


func init(){
	Frida_init()
}

var (
	platformExtNames = map[string]string{
		"windows": ".dll",
		"linux":   ".so",
		"darwin":  ".dylib",
	}
)

func getDLLName() string {
	libName := "frida_shared"
	if ext, ok := platformExtNames[runtime.GOOS]; ok {
		return libName + ext
	}
	return libName
}