package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"

	"github.com/Tnze/go-mc/nbt"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No input files")
		os.Exit(22)
	}
	for i := 0; i < len(os.Args)-1; i++ {
		fname := os.Args[1+i]
		b, err := os.ReadFile(fname)
		if err != nil {
			fmt.Printf("Failed to read file [%s]: %s\n", fname, err)
			continue
		}
		if b[0] == 0x1f && b[1] == 0x8b {
			g, err := gzip.NewReader(bytes.NewReader(b))
			if err != nil {
				fmt.Printf("Failed to decompress nbt [%s]: %s\n", fname, err)
				continue
			}
			c, err := io.ReadAll(g)
			if err != nil {
				fmt.Printf("Failed to decompress nbt [%s]: %s\n", fname, err)
				continue
			}
			g.Close()
			b = c
		}
		var s nbt.StringifiedMessage
		err = nbt.Unmarshal(b, &s)
		if err != nil {
			fmt.Printf("Failed to decode nbt [%s]: %s\n", fname, err)
			continue
		}
		if len(os.Args) > 2 {
			fmt.Printf("%s\t%s\n", fname, s)
		} else {
			fmt.Println(s)
		}
	}

}
