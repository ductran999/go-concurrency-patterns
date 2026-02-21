package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func sleep(sleepTime int) {
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	fmt.Println("Slept for", sleepTime, "ms")
}

func hello(w http.ResponseWriter, r *http.Request) {
	sleep(5)
	sleep(10)
	io.WriteString(w, "Memory Management Test")
}

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		slog.Error(err.Error())
	}
}
