package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {
	log.Println("standard logger")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	// custom logger
	mylog := log.New(os.Stdout, "mylog ", log.LstdFlags)
	mylog.Println("from mylog")

	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	// other io.Writer
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)
	buflog.Println("hello buf")
	fmt.Println("from buflog", buf.String())

	// slog
	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")
	myslog.Info("hello again", "key", "val", "age", 25)
}
