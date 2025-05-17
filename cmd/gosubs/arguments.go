package main

import (
	"flag"
	"fmt"
	"os"
)

func parseArguments() *arguments {
	model := flag.String("model", "ggml-medium.bin", "path to the whisper ggml model")
	flag.Parse()

	if len(flag.Args()) != 2 {
		fmt.Printf("usage: %s {input filename} {output filename} [-model=ggml-medium.bin]\n", os.Args[0])
		os.Exit(1)
	}

	return &arguments{
		model:          *model,
		inputFilename:  flag.Arg(0),
		outputFilename: flag.Arg(1),
	}
}

type arguments struct {
	model          string
	inputFilename  string
	outputFilename string
}
