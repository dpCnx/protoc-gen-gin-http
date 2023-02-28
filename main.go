package main

import (
	"flag"
	"fmt"

	"github.com/dpCnx/protoc-gen-gin-http/logic"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

/*
	protobuf插件名称需要使用protoc-gen-xxx
	当使用protoc --xxx_out时就会调用proto-gen-xxx插件
*/

var (
	showVersion = flag.Bool("version", false, "print the version and exit")
	omitempty   = flag.Bool("omitempty", true, "omit if google.api is empty")
)

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-gin-http %v\n", logic.Release)
		return
	}
	protogen.Options{
		ParamFunc: flag.CommandLine.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			logic.GenerateFile(gen, f, *omitempty)
		}
		return nil
	})
}
