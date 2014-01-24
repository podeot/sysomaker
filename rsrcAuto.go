package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const a = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<assembly xmlns="urn:schemas-microsoft-com:asm.v1" manifestVersion="1.0">
	<assemblyIdentity version="1.0.0.0" processorArchitecture="*" name="SomeFunkyNameHere" type="win32"/>
	<dependency>
		<dependentAssembly>
			<assemblyIdentity type="win32" name="Microsoft.Windows.Common-Controls" version="6.0.0.0" processorArchitecture="*" publicKeyToken="6595b64144ccf1df" language="*"/>
		</dependentAssembly>
	</dependency>
</assembly>
`
const ext = "manifest"

func rsrcExist(rsrcAddr string) bool {
	if _, err := os.Stat(rsrcAddr); err != nil {
		return false
	} else {
		return true
	}
}

func main() {
	read := bufio.NewReader(os.Stdin)
	fmt.Println("type your gofile name (ex) file.go -> type file")
	fileName, _ := read.ReadString('\n')
	fileName = strings.Split(fileName, "\r\n")[0]
	tempPath, _ := filepath.Abs("./")
	tempFileName := strings.Join([]string{fileName, ext}, ".")
	New := tempPath + `\` + tempFileName
	fileCreation, _ := os.Create(New)
	fileCreation.WriteString(a)
	fileCreation.Close()
	rsrcAddr := os.ExpandEnv(`${GOPATH}\src\github.com\akavel\rsrc`)
	if rsrcExist(rsrcAddr) == false {
		fmt.Println(`no rsrc installed correctly. please do 'go get github.com/akavel/rsrc'`)
		os.Exit(1)
	}
	cmd := exec.Command("rsrc", "-manifest", tempFileName, "-o", "rsrc.syso")
	err := cmd.Start()
	if err != nil {
		println("error")
	}
	err = cmd.Wait()
	os.Remove(New)
}
