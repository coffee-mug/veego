package editor

import (
	"testing"
)

type mockStdin struct {
	buffer []byte
}

func (s mockStdin) Read(p []byte) (n int, err error) {
	for _, b := range s.buffer {
		p = append(p, b)
	}

	return len(s.buffer), nil
}

func (s mockStdin) Write(p []byte) (n int, err error) {
	for _, b := range p {
		s.buffer = append(s.buffer, b)
	}

	return len(p), nil
}

func TestEditor(t *testing.T) {
	t.Run("Typing a character displays it on screen", func(t *testing.T) {
		var editor Editor
		character := "a"

		stdin := mockStdin{[]byte{}}

		editor.ListenForInputs(stdin)

		stdin.Write([]byte("a"))

		got := editor.buffer

		if string(got) != character {
			t.Errorf("Got %v Want %v", got, character)
		}
	})
}
