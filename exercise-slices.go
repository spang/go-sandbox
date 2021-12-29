package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	s_of_s := [][]uint8{}
	for y := 0; y < dy; y++ {
		s := make([]uint8, dx)
		for x, _ := range s {
			s[x] = uint8(x ^ y)
		}
		s_of_s = append(s_of_s, s)
	}

	return s_of_s
}

func main() {
	pic.Show(Pic)
}
