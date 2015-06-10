package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func genStubs(protoFile, outDir string) error {
	var proto Protocol
	if f, err := os.Open(protoFile); err == nil {
		if err = xml.NewDecoder(f).Decode(&proto); err != nil {
			return err
		}
	} else {
		return err
	}
	clientDir := fmt.Sprintf("%s%cclient", outDir, os.PathSeparator)
	serverDir := fmt.Sprintf("%s%cserver", outDir, os.PathSeparator)

	for _, v := range []string{clientDir, serverDir} {
		if _, err := os.Stat(v); err != nil {
			if err := os.MkdirAll(v, 0755); err != nil {
				return err
			}
		}
	}

	if err := genClient(proto, clientDir); err != nil {
		return err
	}

	return genServer(proto, serverDir)
}

func goify(name string) string {
	subs := strings.Split(name, "_")
	caps := make([]string, 0, len(subs))
	for _, v := range subs {
		caps = append(caps, strings.Title(v))
	}

	return strings.Join(caps, "")
}

func outputDesc(file io.Writer, desc Description) {
	for _, v := range strings.Split(desc.Full, "\n") {
		tr := strings.Trim(v, " \t")
		if len(tr) != 0 {
			fmt.Fprintf(file, "// %s\n", tr)
		}
	}
}

func makeArgs(args []Arg) string {
	goArgs := make([]string, 0, len(args))
	for _, v := range args {
		name := v.Name
		if keywords[name] {
			name = "wl_" + name
		}
		name = goify(name)
		argType := goType(v)

		goArgs = append(goArgs, fmt.Sprintf("%s %s", name, argType))
	}

	return strings.Join(goArgs, ",")
}

func goType(arg Arg) string {
	switch arg.Type {
	case "int":
		return "int32"
	case "uint":
		return "uint32"
	case "fixed":
		return "float32"
	case "fd":
		return "uintptr"
	case "new_id":
		return "uint32"
	case "object":
		return "uint32"
	case "array":
		return "[]byte"
	default:
		return arg.Type
	}
}

func genClient(proto Protocol, clientDir string) error {
	for _, iface := range proto.Interfaces {
		iFile, err := os.Create(fmt.Sprintf("%s%c%s.go", clientDir, os.PathSeparator, iface.Name))
		if err != nil {
			return err
		}
		fmt.Fprintln(iFile, "package client")
		outputDesc(iFile, iface.Description)
		fmt.Fprintf(iFile, "type %s struct{}\n", goify(iface.Name))

		for _, v := range iface.Events {
			outputDesc(iFile, v.Description)
			fmt.Fprintf(iFile, "func (obj *%s) %s(%s) {}\n", goify(iface.Name), goify(v.Name), makeArgs(v.Args))
		}
		for _, v := range iface.Requests {
			outputDesc(iFile, v.Description)
			fmt.Fprintf(iFile, "func (obj *%s) %s(%s) {}\n", goify(iface.Name), goify(v.Name), makeArgs(v.Args))
		}

		for _, v := range iface.Enums {
			outputDesc(iFile, v.Description)
			etype := goify(iface.Name + "_" + v.Name)
			fmt.Fprintf(iFile, "type %s uint32\n", etype)
			fmt.Fprintln(iFile, "const (")
			for _, w := range v.Entries {
				fmt.Fprintf(iFile, "%s %s = %s\n", goify(iface.Name+"_"+w.Name), etype, w.Value)
			}
			fmt.Fprintln(iFile, ")")
		}

		iFile.Close()
	}

	return nil
}

func genServer(proto Protocol, serverDir string) error {
	return genClient(proto, serverDir)
}
