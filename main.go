package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/sys/windows"
	"os"
	"os/exec"
	"unsafe"
)

type configuration struct {
	IP       string
	Port     int
	Nprotect int
	Locale   string
	Country  string
}

func main() {
	var config configuration

	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println(err.Error())
	}

	var securityAttrib windows.SecurityAttributes
	securityAttrib.InheritHandle = 1
	securityAttrib.SecurityDescriptor = 0
	securityAttrib.Length = uint32(unsafe.Sizeof(securityAttrib))

	eventHandle, err := windows.CreateEvent(&securityAttrib, 0, 0, nil)

	if err != nil {
		fmt.Println(err)
	}
	if err == nil {
		args := []string{fmt.Sprintf("/auth_ip:%s", config.IP), fmt.Sprintf("/auth_port:%d", config.Port),
			fmt.Sprintf("/use_nprotect:%d", config.Nprotect), "/help_url_w:611", "/help_url_h:625",
			fmt.Sprintf("/locale:%s", config.Locale), fmt.Sprintf("/country:%s", config.Country),
			"/cash", "/commercial_shop", "/layout_dir:6", "/layout_auto:0", "/cash_url_w:800", "/cash_url_h:631"}
		sframe := exec.Command("sframe.exe")
		sframeParent := "sframe.exe_PARENT=Launcher.exe"
		sframeRunner := fmt.Sprintf("sframe.exe_RUNNER=%d", eventHandle)
		sframe.Env = os.Environ()
		sframe.Env = append(sframe.Env, sframeParent)
		sframe.Env = append(sframe.Env, sframeRunner)
		sframe.Args = args
		sframe.Start()

		windows.WaitForSingleObject(eventHandle, 10*1000)
	}
}
