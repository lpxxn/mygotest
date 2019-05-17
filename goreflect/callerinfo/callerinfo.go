package callerinfo

import (
	"path"
	"runtime"
	"strings"
)

type callInfo struct {
	Path        string
	PackageName string
	FileName    string
	FuncName    string
	line        int
}

//RetrieveCallInfo
func RetrieveCallInfo() *callInfo {
	pc, file, line, _ := runtime.Caller(1)
	_, fileName := path.Split(file)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	packageName := ""
	funcName := parts[pl-1]

	if parts[pl-2][0] == '(' {
		funcName = parts[pl-2] + "." + funcName
		packageName = strings.Join(parts[0:pl-2], ".")
	} else {
		packageName = strings.Join(parts[0:pl-1], ".")
	}

	return &callInfo{
		Path:        file,
		PackageName: packageName,
		FileName:    fileName,
		FuncName:    funcName,
		line:        line,
	}

}
