package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type appState struct {
	Timestamp  time.Time
	AppVersion string
}

type persist struct {
	file *os.File
}

func (p *persist) OpenFile() {
	file, err := os.Create("./data/appState.json")
	if err != nil {
		panic(err)
	}

	p.file = file
}

func (p *persist) WriteFile(content appState) {
	jsonContent, _ := json.Marshal(content)
	if nBytes, err := p.file.WriteString(string(jsonContent)); err != nil {
		panic(err)
	} else {
		fmt.Printf("wrote %d bytes\n", nBytes)
	}
}

func (p *persist) CloseFile() {
	p.file.Close()
}
