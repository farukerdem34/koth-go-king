package main

import (
	"fmt"
	"os"
	"reflect"
	"time"
	"unsafe"

	"github.com/g0rbe/go-chattr"
)

const king_path string = "/root/king.txt"
const king_name string = "lomarkomar"

func main() {
	SetProcessName("systemd")
	for true {
		go full_loop()
	}
}

func full_loop() {
	file, _ := os.OpenFile(king_path, os.O_RDONLY, 0666)
	time.Sleep(500 * time.Millisecond)
	chattr.UnsetAttr(file, chattr.FS_IMMUTABLE_FL)
	chattr.UnsetAttr(file, chattr.FS_APPEND_FL)
	time.Sleep(500 * time.Millisecond)
	write_king()
	time.Sleep(500 * time.Millisecond)
	chattr.SetAttr(file, chattr.FS_APPEND_FL)
	chattr.SetAttr(file, chattr.FS_IMMUTABLE_FL)
	file.Close()
}

func check(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func write_king() {
	file, err := os.OpenFile(king_path, os.O_WRONLY, 0666)
	check(err)
	data := []byte(king_name)
	_, err = file.Write(data)
	file.Close()
}

func SetProcessName(name string) error {
	argv0str := (*reflect.StringHeader)(unsafe.Pointer(&os.Args[0]))
	argv0 := (*[1 << 30]byte)(unsafe.Pointer(argv0str.Data))[:argv0str.Len]

	n := copy(argv0, name)
	if n < len(argv0) {
		argv0[n] = 0
	}

	return nil
}
