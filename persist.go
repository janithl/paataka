package main

import (
	"fmt"
	"os"
)

type persist struct {
	filename string
	file     *os.File
}

func (p *persist) OpenFile() {
	file, err := os.Create(p.filename)
	if err != nil {
		panic(err)
	}

	p.file = file
}

func (p *persist) WriteFile(filename string, content string) {
	if p.file == nil || p.filename != filename {
		p.filename = filename
		p.OpenFile()
		defer p.CloseFile()
	}

	if nBytes, err := p.file.WriteString(content); err != nil {
		panic(err)
	} else {
		fmt.Printf("wrote %d bytes\n", nBytes)
	}
}

func (p *persist) CloseFile() {
	p.file.Close()
}
