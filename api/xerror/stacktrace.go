package xerror

import (
	"runtime"
	"strings"
)

// getStackTrace はこの処理に必要な内部処理を実行します。
func getStackTrace(depth int) (moduleName, funcName string, line int, ok bool) {
	pt, _, line, ok := runtime.Caller(depth)
	if !ok {
		return "", "", 0, false
	}

	modName, fnName := getFunctionName(pt)
	return modName, fnName, line, true
}

// getFunctionName はこの処理に必要な内部処理を実行します。
func getFunctionName(pt uintptr) (string, string) {
	fn := runtime.FuncForPC(pt)
	if fn == nil {
		return "", ""
	}
	pack := ""
	name := fn.Name()
	if idx := strings.LastIndex(name, "."); idx != -1 {
		pack = name[:idx]
		name = name[idx+1:]
	}
	name = strings.ReplaceAll(name, "·", ".")
	return pack, name
}
