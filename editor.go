package editor

import (
	"bufio"
	"io"
	"log"
)

type Editor struct {
	buffer []byte
}

func (e *Editor) ListenForInputs(inp io.Reader) {
	reader := bufio.NewReader(inp)

	b, err := reader.ReadByte()
	if err != nil {
		log.Fatal(err)
	}

	e.buffer = append(e.buffer, b)
}
