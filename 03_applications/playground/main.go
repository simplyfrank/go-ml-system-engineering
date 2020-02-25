package main

import "fmt"

type Logger interface {
	log()
}

type WebLogger struct {}
func (*WebLogger) log(){}

type FileLogger struct {}
func (*FileLogger) log(){}

func main() {
	a := &WebLogger{}
	b := &FileLogger{}
	a.log()
	b.log()
	loggerDefault(b)
}
func loggerDefault(lg Logger) {
	switch lg.(type) {
	case *WebLogger:
		fmt.Println("It's a weblogger type")
	case *FileLogger:
		fmt.Println("It's a FileLogger type")
	}
}

