package main

import (
	"flag"
	"log"
	"os/exec"
)

var protoFile, outDir string

func init() {
	flag.StringVar(&protoFile, "proto", "wayland.xml", "protocol specification file")
	flag.StringVar(&outDir, "out", ".", "directory for output")
}
func main() {
	flag.Parse()
	if err := genStubs(protoFile, outDir); err != nil {
		log.Fatal(err)
	}
	exec.Command("gofmt", "-w", outDir).Run()
}
