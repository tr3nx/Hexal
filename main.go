package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	var output strings.Builder
	var hex strings.Builder
	var ascii strings.Builder

	filebytes, err := ioutil.ReadFile(args[0])
	if err != nil {
		panic(err)
	}

	var b bytes.Buffer
	b.Write(filebytes)

	for i := 0; ; i += 16 {
		chunk := b.Next(16)
		chunklen := len(chunk)

		if chunklen <= 0 {
			break
		}

		output.WriteString(fmt.Sprintf("%08x: ", i))

		for j, b := range chunk {
			hex.WriteString(fmt.Sprintf("%02x", b))

			if (j+1)%2 == 0 && j != 15 {
				hex.WriteString(" ")
			}

			if int(b) > 31 && int(b) <= 126 {
				ascii.WriteString(string(b))
			} else {
				ascii.WriteString(".")
			}
		}

		if chunklen < 16 {
			pads := 16 - chunklen
			hex.WriteString(strings.Repeat(" ", pads*2))

			spaces := pads / 2
			hex.WriteString(strings.Repeat(" ", spaces))
		}

		hex.WriteString(strings.Repeat(" ", 2))

		output.WriteString(hex.String())
		hex.Reset()

		output.WriteString(ascii.String())
		ascii.Reset()

		output.WriteString(fmt.Sprintln())

		if chunklen < 16 {
			break
		}
	}

	fmt.Printf(output.String())
}
