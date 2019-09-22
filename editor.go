package editor

import (
	"bufio"
	"io"
	"log"
)

type Editor struct {
	buffer []rune
}

func (e *Editor) ListenForInputs(inp io.Reader) {
	reader := bufio.NewReader(inp)

	r, _, err := reader.ReadRune()
	if err != nil {
		log.Fatal(err)
	}

	e.buffer = append(e.buffer, r)
}
