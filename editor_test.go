package editor

import (
	"io"
	"log"
	"testing"
)

type mockStdin struct {
	s []byte
	i int64
}

func (m *mockStdin) Read(p []byte) (n int, err error) {
	if m.i >= int64(len(m.s)) {
		return 0, io.EOF
	}

	n = copy(p, m.s[m.i:])
	m.i += int64(n)

	return
}

func (m *mockStdin) Write(p []byte) (n int, err error) {
	for _, letter := range p {
		m.s = append(m.s, letter)
		n++
	}
	return
}

func TestEditor(t *testing.T) {
	t.Run("mockStdin implements Reader interface", func(t *testing.T) {
		greeting := "Hello World !"

		stdin := mockStdin{[]byte(greeting), 0}
		buf := make([]byte, len(greeting))
		n, err := stdin.Read(buf)

		if err != nil {
			log.Fatal(err)
		}

		got := string(buf)

		if got != greeting {
			t.Errorf("Got %s Expected %s", got, greeting)
		}

		if n != len(greeting) {
			t.Errorf("Wrong number of chars read, Got %d Expected %d", n, len(greeting))
		}
	})

	t.Run("mockStdin implements Writer interface", func(t *testing.T) {
		greeting := "Hello"
		guy := "Lucas"

		stdin := mockStdin{[]byte(greeting), 0}
		buf := []byte(guy)
		n, err := stdin.Write(buf)

		if err != nil {
			log.Fatal(err)
		}

		got := string(stdin.s)
		expect := greeting + guy

		if got != expect {
			t.Errorf("Got %s Expected %s", got, expect)
		}

		if n != len(guy) {
			t.Errorf("Wrong number of chars read, Got %d Expected %d", n, len(guy))
		}
	})
}
