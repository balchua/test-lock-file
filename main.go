package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func doSomething(file *os.File, isWriting bool) {
	if isWriting {
		w := bufio.NewWriter(file)
		currentTime := time.Now()
		_, err := w.WriteString(currentTime.String() + "\n")
		if err != nil {
			log.Printf("Error writing to the file %s: %s", file.Name(), err)
			return
		}
		w.Flush()
	}
	time.Sleep(time.Second * 5)
}

func main() {
	time.Sleep(time.Second)
	var isWriting = false

	log.Printf("arg size: %d", len(os.Args))
	if len(os.Args) < 2 {
		log.Fatal("Invalid Arguments")
	} else if len(os.Args) > 2 {
		isWriting, _ = strconv.ParseBool(strings.ToLower(os.Args[2]))
		log.Printf("write to file %v", isWriting)
	}

	name := os.Args[1]

	file, err := os.OpenFile(name, syscall.O_CREAT|syscall.O_RDWR|syscall.O_NOCTTY, 0666)
	if err != nil {
		log.Printf("error opening file: %s", err)
		return
	}
	defer file.Close()

	flockT := syscall.Flock_t{
		Type:   syscall.F_WRLCK,
		Whence: io.SeekStart,
		Start:  0,
		Len:    0,
	}

	err = syscall.FcntlFlock(file.Fd(), syscall.F_SETLK, &flockT)
	if err != nil {
		log.Printf("Error locking the file %s: %s", file.Name(), err)
		return
	}

	log.Printf("file %s accessed", file.Name())

	for {
		doSomething(file, isWriting)
	}
}
