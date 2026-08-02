package xerror

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/morikuni/failure"
)

const (
	ServerError failure.StringCode = "E-0-00001"
)

func IsCustomError(err error) bool {
	_, ok := failure.CodeOf(err)
	return ok
}

func GetHTTPErrorCode(err error) int {
	code, ok := failure.CodeOf(err)
	if !ok {
		return http.StatusInternalServerError
	}

	c := code.ErrorCode()
	switch {
	case strings.HasPrefix(c, "A-"):
		return http.StatusBadRequest
	case strings.HasPrefix(c, "C-"):
		return http.StatusUnauthorized
	case strings.HasPrefix(c, "E-"):
		return http.StatusServiceUnavailable
	default:
		return http.StatusTeapot
	}
}

func GetStringCode(err error) string {
	code, ok := failure.CodeOf(err)
	if !ok {
		return "-1"
	}

	return code.ErrorCode()
}

func GetErrorMessage(err error) string {
	if err == nil {
		return ""
	}

	msg, ok := failure.MessageOf(err)
	if !ok {
		return err.Error()
	}
	return msg
}

func UnknownServerErr(err error) error {
	if err == nil {
		return nil
	}

	return failure.New(ServerError,
		failure.Context{"error": err.Error()},
	)
}

func GetErrMetaData(err error) (map[string]string, bool) {
	m := make(map[string]string)
	if err == nil {
		return m, false
	}

	i := failure.NewIterator(err)
	for i.Next() {
		var ctx failure.Context
		if i.As(&ctx) {
			m = mergeMetaData(m, ctx)
		}
	}
	return m, len(m) != 0
}

func AddMetaData(err error, meta ...map[string]string) error {
	list := make([]failure.Wrapper, len(meta))
	for i, v := range meta {
		list[i] = failure.Context(v)
	}
	return failure.Custom(err, list...)
}

func mergeMetaData(c1, c2 map[string]string) failure.Context {
	vv := failure.Context{}
	for k, v := range c1 {
		vv[k] = v
	}
	for k, v := range c2 {
		vv[k] = v
	}
	return vv
}

func getDefaultMetaData(err error, meta map[string]string) failure.Context {
	c := failure.Context{
		"logging": "1",
		"cause":   err.Error(),
	}
	return mergeMetaData(c, meta)
}

func generateErr(err error, text string, code failure.StringCode, meta ...map[string]string) error {
	return generateErrWithStackDepth(err, text, code, 1, meta...)
}

func generateErrWithStackDepth(err error, text string, code failure.StringCode, stackDepth int, meta ...map[string]string) error {
	if err == nil {
		return nil
	}

	c := failure.Context{}
	if len(meta) != 0 {
		c = meta[0]
	}
	c["code"] = string(code)

	// Stack Traceの開始深さを調整する（元関数, generateErrWithStackDepth, getStackTrace で+3）
	const defaultDepth = 3
	stackDepth += defaultDepth
	m, f, l, ok := getStackTrace(stackDepth)
	if ok {
		// m = strings.TrimPrefix(m, xconst.PackagePrefix)
		c["app_pkg"] = m
		c["app_func"] = f
		c["app_line"] = strconv.Itoa(l)
	}

	return failure.Translate(err, code,
		failure.Message(text),
		getDefaultMetaData(err, c),
	)
}

func generateErrFromText(text string, code failure.StringCode, meta ...map[string]string) error {
	err := errors.New(text)
	return generateErrWithStackDepth(err, text, code, 1, meta...)
}
