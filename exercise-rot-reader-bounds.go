package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot13.r.Read(b)

	if err != nil {
		return n, err
	}

	for i := 0; i < n; i++ {
		// lowercase translation
		if b[i] > 'a' && b[i] <= 'z' {
			new_rune := b[i] + 13
			if new_rune > 'z' {
				new_rune -= ('z' - 'a' + 1)
			}
			b[i] = new_rune
			// uppercase translation
		} else if b[i] > 'A' && b[i] <= 'Z' {
			new_rune := b[i] + 13
			if new_rune > 'Z' {
				new_rune -= ('Z' - 'A' + 1)
			}
			b[i] = new_rune
		}
	}

	return n, nil

}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
