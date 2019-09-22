package editor

import (
	"testing"
)

type mockStdin struct {
	buffer []byte
}

func (s *mockStdin) Read(p []byte) (n int, err error) {

}

func TestEditor(t *testing.T) {
	t.Run("Typing a character displays it on screen", func(t *testing.T) {
		var editor Editor
		character := 'a'

		editor.ListenForInputs()

		got := editor.buffer

		if got[0] != character {
			t.Errorf("Got %v Want %v", got, character)
		}
	})
}
