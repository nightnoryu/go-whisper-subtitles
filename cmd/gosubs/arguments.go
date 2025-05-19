package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func parseArguments() (*arguments, error) {
	set := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	set.Usage = func() {
		_, executableName := filepath.Split(os.Args[0])
		fmt.Fprintf(set.Output(), "Usage: %s [options] <input> <output>\n", executableName)
	}

	model := set.String("model", "ggml-medium.bin", "path to the whisper ggml model")

	if err := set.Parse(os.Args[1:]); err != nil {
		return nil, err
	}

	inputFilename := set.Arg(0)
	if inputFilename == "" {
		set.Usage()
		return nil, errors.New("no input provided")
	}

	outputFilename := set.Arg(1)
	if outputFilename == "" {
		set.Usage()
		return nil, errors.New("no output provided")
	}

	return &arguments{
		model:          *model,
		inputFilename:  inputFilename,
		outputFilename: outputFilename,
	}, nil
}

type arguments struct {
	model          string
	inputFilename  string
	outputFilename string
}
