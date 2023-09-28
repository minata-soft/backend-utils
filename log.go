package backendutils

import (
	"log"
	"reflect"
	"runtime"
)

const (
	redColor    = "\033[31m"
	yellowColor = "\033[33m"
	blueColor   = "\033[34m"
	resetColor  = "\033[0m"
)

func DebugMessage(message string, err ...error) {
	pc, _, _, _ := runtime.Caller(1)
	functionName := runtime.FuncForPC(pc).Name()

	log.Printf("%s== DEBUG ==%s --FUNCTION--: %s --MESSAGE-- : %s%s\n",
		blueColor, resetColor, functionName, message, resetColor)
	if len(err) > 0 {
		log.Printf("%s==> err: %v%s\n", yellowColor, err[0], resetColor)
	}
}

func ErrorMessage(message string, err ...error) {
	pc, _, _, _ := runtime.Caller(1)
	functionName := runtime.FuncForPC(pc).Name()

	log.Printf("%s== ERROR ==%s --FUNCTION--: %s --MESSAGE-- : %s%s\n",
		redColor, resetColor, functionName, message, resetColor)
	if len(err) > 0 {
		log.Printf("%s==> err: %v%s\n", yellowColor, err[0], resetColor)
	}
}

func PrintError(fn interface{}, message string, err ...error) {
	var (
		pc               = reflect.ValueOf(fn).Pointer()
		functionName     = runtime.FuncForPC(pc).Name()
		_, file, line, _ = runtime.Caller(1)
	)

	log.Printf("%s== ERROR ==%s --FILE-- %s:%d --FUNCTION--: %s --MESSAGE-- : %s%s\n",
		redColor, resetColor, file, line, functionName, message, resetColor)

	if len(err) > 0 {
		log.Printf("%s==> err:%v %s\n", redColor, err, resetColor)
	}
}
