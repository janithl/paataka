package common

import (
	"log"
	"os"
)

// Persist struct to store application state in file system storage
type Persist struct {
	filename string
	file     *os.File
}

// OpenFile opens a file to write to
func (p *Persist) OpenFile() {
	file, err := os.Create(p.filename)
	if err != nil {
		panic(err)
	}

	p.file = file
}

// WriteFile writes to the file
func (p *Persist) WriteFile(filename string, content string) {
	if p.file == nil || p.filename != filename {
		p.filename = filename
		p.OpenFile()
		defer p.file.Close()
	}

	if nBytes, err := p.file.WriteString(content); err != nil {
		panic(err)
	} else {
		log.Printf("Wrote %d bytes to file %s", nBytes, filename)
	}
}
