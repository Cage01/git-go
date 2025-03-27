package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
)

// Usage: your_program.sh <command> <arg1> <arg2> ...
func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintf(os.Stderr, "Logs from your program will appear here!\n")

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: mygit <command> [<args>...]\n")
		os.Exit(1)
	}

	switch command := os.Args[1]; command {
	case "init":
		// Uncomment this block to pass the first stage!

		for _, dir := range []string{".git", ".git/objects", ".git/refs"} {
			if err := os.MkdirAll(dir, 0755); err != nil {
				fmt.Fprintf(os.Stderr, "Error creating directory: %s\n", err)
			}
		}

		headFileContents := []byte("ref: refs/heads/main\n")
		if err := os.WriteFile(".git/HEAD", headFileContents, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing file: %s\n", err)
		}

		fmt.Println("Initialized git directory")

	case "cat-file":
		// fmt.Printf("checking args: %s\n", os.Args[1:])
		sha := os.Args[len(os.Args)-1]
		prefix := sha[0:2]
		suffix := sha[3:]

		if slices.Contains(os.Args[1:], "-p") {
			path := ".git/objects/" + prefix + "/" + suffix
			fmt.Println(readFile(path))
		}
	case "test":
		fmt.Println("This is a test")
	default:
		fmt.Fprintf(os.Stderr, "Unknown command %s\n", command)
		os.Exit(1)
	}

}

func readFile(file string) []byte {
	b, err := os.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	return decompress(b)
}

func decompress(compressedData []byte) []byte {
	b := bytes.NewReader(compressedData)

	r, err := zlib.NewReader(b)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	decompressedData, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	return decompressedData
}
