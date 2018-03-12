package main

import (
	"fmt"
	"strconv"
	"syscall"
	"unsafe"
)

type OSVERSIONINFO struct {
	dwOSVersionInfoSize int32
	dwMajorVersion      int32
	dwMinorVersion      int32
	dwBuildNumber       int32
	dwPlatformId        int32
	szCSDVersion        [128]byte
}

func main() {
	fmt.Printf(strconv.Itoa(GetOsVersion()))

}

func GetOsVersion() int {
	var version int = 0
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	var os OSVERSIONINFO
	os.dwOSVersionInfoSize = int32(unsafe.Sizeof(os))
	GetVersionExA := kernel32.NewProc("GetVersionExA")
	rt, _, _ := GetVersionExA.Call(uintptr(unsafe.Pointer(&os)))
	if int(rt) == 1 {
		switch {

		case os.dwMajorVersion == 5 && os.dwMinorVersion == 1:
			version = 5
			break
		case os.dwMajorVersion == 5 && os.dwMinorVersion == 2:
			version = 3
			break
		case os.dwMajorVersion == 6 && os.dwMinorVersion == 0:
			version = 6
			break
		case os.dwMajorVersion == 6 && os.dwMinorVersion == 2:
			version = 8
			break
		case os.dwMajorVersion == 6 && os.dwMinorVersion == 3:
			version = 8
			break
		case os.dwMajorVersion == 10 && os.dwMinorVersion == 3:
			version = 10
			break
		default:
			version = 7
			break
		}
		//ersion = version + " Build(" + strconv.FormatInt(int64(os.dwBuildNumber), 10) + ") " + string(os.szCSDVersion[0:])
	}
	return version
}
