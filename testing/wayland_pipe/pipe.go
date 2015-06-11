package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"runtime"

	"github.com/Pursuit92/goland/gen"
	"github.com/davecgh/go-spew/spew"
)

var msgs chan string

func printMsgs(msgs chan string) {
	for s := range msgs {
		fmt.Println(s)
	}
}

var compositorName = "/run/user/1000/compositor"

func socketTest() {
	runtime.GOMAXPROCS(runtime.NumCPU())

start:
	sock, err := net.ListenUnix("unix", &net.UnixAddr{Name: compositorName, Net: "unix"})
	if err != nil {
		if _, err := net.Dial("unix", compositorName); err != nil {
			os.Remove(compositorName)
			goto start
		} else {
			log.Fatal("socket is in use")
		}
	}
	defer sock.Close()
	msgs = make(chan string)
	go printMsgs(msgs)

	for {
		conn, err := sock.AcceptUnix()
		if err != nil {
			log.Println(err)
		}
		go handleConn(conn)
	}
}

var wAddr = &net.UnixAddr{
	Name: "/run/user/1000/weston",
	Net:  "unix",
}

func unixSockPipe(conn1, conn2 *net.UnixConn) {
	for {
		msg, err := gen.ReadMsg(conn1)
		if err != nil {
			return
		}
		msgs <- fmt.Sprintf("%s:\n%s", conn1.RemoteAddr().String(), spew.Sdump(msg))
		err = gen.SendMsg(conn2, msg)
		if err != nil {
			return
		}
	}
}

func handleConn(conn *net.UnixConn) {
	defer conn.Close()
	westConn, err := net.DialUnix("unix", nil, wAddr)
	if err != nil {
		log.Println(err)
		return
	}
	defer westConn.Close()
	go unixSockPipe(westConn, conn)
	unixSockPipe(conn, westConn)
	fmt.Println("All done!")
}
