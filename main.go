package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/TheInvader360/jack-compiler/engine"
	"github.com/TheInvader360/jack-compiler/handler"
	"github.com/TheInvader360/jack-compiler/tokenizer"

	"github.com/pkg/errors"
)

func main() {
	path := flag.String("path", "examples/HelloWorld", "jack source path")
	debug := flag.Bool("debug", false, "enable debug terminal output")
	vmComments := flag.Bool("comments", false, "enable comments in vm output")
	xml := flag.Bool("xml", false, "include intermediate xml files in output")
	flag.Parse()

	fileInfo, err := os.Stat(*path)
	handler.FatalError(err)

	if fileInfo.IsDir() {
		if strings.HasSuffix(*path, "/") {
			*path += "*"
		} else {
			*path += "/*"
		}
	} else {
		if !strings.HasSuffix(*path, ".jack") {
			handler.FatalError(errors.New("Expected a jack file (*.jack)"))
		}
	}

	files, err := filepath.Glob(*path)
	handler.FatalError(err)

	for _, file := range files {
		if strings.HasSuffix(file, ".jack") {
			jackData, err := ioutil.ReadFile(file)
			handler.FatalError(errors.Wrap(err, fmt.Sprintf("Can't read file: %s", file)))

			tokens := tokenizer.Tokenize(jackData)
			engine := engine.NewCompilationEngine(tokens, *debug, *vmComments)
			tree, vmCode := engine.CompileClass()

			if *xml {
				writeFile(tokenizer.AsXML(tokens), strings.Replace(file, ".jack", "T.xml", 1))
				writeFile(tree, strings.Replace(file, ".jack", ".xml", 1))
			}

			writeFile(vmCode, strings.Replace(file, ".jack", ".vm", 1))
		}
	}
}

func writeFile(data, path string) {
	output := []byte(data)
	if len(output) == 0 {
		handler.FatalError(errors.New("Zero length output"))
	}
	err := ioutil.WriteFile(path, output, 0777)
	handler.FatalError(errors.Wrap(err, fmt.Sprintf("Can't write file: %s", path)))
}
