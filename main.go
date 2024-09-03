package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Gimme something to mask!!!")
		return
	}

	const (
		linkPattern = "http://"
		nlink       = len(linkPattern)
		mask        = '*'
	)

	var (
		text = args[0]
		size = len(text)
		buf  = make([]byte, 0, size)
		in   bool
	)

	// We use buffer bc strings are immutable.
	// So if we used a string buffer, it would allocate
	// a new string value each time. Small performance hit now,
	// but if we were processing lots of data, it would be a big hit
	for i := 0; i < size; i++ {

		// We want to know when we are inside the link
		if len(text[i:]) >= nlink && text[i:i+nlink] == linkPattern {
			in = true
			// Appends the bytes of link pattern
			buf = append(buf, linkPattern...)
			i += nlink
		}

		c := text[i]
		switch c {
		case ' ', '\t', '\n':
			in = false
		}
		// If in the link, then we want to mask the character
		if in {
			c = mask
		}
		buf = append(buf, c)
	}

	fmt.Println(string(buf))

}
