package main

import (
	"fmt"
	"io/fs"
	"log"

	"github.com/rudifa/goembed-demo/pkg/embedded"
)

func main() {
	fmt.Println("Here we go")

	var x = embedded.Dir12

	// println type of x
	fmt.Printf("%T\n", x)

	// Read the files from Dir12Files
	dir12Files, err := embedded.Dir12.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Files in .:") // dir1, dir2
	for _, file := range dir12Files {
		fi, _ := file.Info()
		fmt.Println("  ", file.Name(), file.IsDir(), file.Type(), fi)
	}

	// Read the files from dir1
	dir1Files, err := embedded.Dir12.ReadDir("dir1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Files in dir1:")
	for _, file := range dir1Files {
		fi, _ := file.Info()
		fmt.Println("  ", file.Name(), file.IsDir(), file.Type(), fi)
	}

	// Read the files from dir2
	dir2Files, err := embedded.Dir12.ReadDir("dir2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Files in dir2:")
	for _, file := range dir2Files {
		fi, _ := file.Info()
		fmt.Println("  ", file.Name(), file.IsDir(), file.Type(), fi)
	}

	// countFiles2(embedded.Dir12Files)

	// ------------------------------------------------------------

	// Read the files from dir3
	dir3Files, err := embedded.Dir3.ReadDir("dir3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Files in dir3:")
	for _, file := range dir3Files {
		fi, _ := file.Info()
		fmt.Println("  ", file.Name(), file.IsDir(), file.Type(), fi)
	}

	// read the contents of the file dir3/file3.txt
	file3, err := embedded.Dir3.ReadFile("dir3/hello3.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Contents of dir3/file3.txt:")
	fmt.Println(string(file3))

	// ------------------------------------------------------------

	fmt.Println("Walk the embedded directory Dir12")
	walkDir(embedded.Dir12)

	fmt.Println("Walk the embedded directory Dir3")
	walkDir(embedded.Dir3)

	fmt.Println("Walk the embedded directory ABC")
	walkDir(embedded.ABC)

	// ------------------------------------------------------------

	// List the files from dir ABC/.
	dir, err := embedded.ABC.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Files in ABC/:")
	for _, file := range dir {
		fi, _ := file.Info()
		fmt.Println("  ", file.Name(), file.IsDir(), file.Type(), fi)
	}


	// List the files from dir ABC/A.
	dir, err = embedded.ABC.ReadDir("A")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Files in ABC/A:")
	for _, file := range dir {
		fi, _ := file.Info()
		fmt.Println("  ", file.Name(), file.IsDir(), file.Type(), fi)
	}


	// List the files from dir ABC/A/B.
	dir, err = embedded.ABC.ReadDir("A/B")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Files in ABC/A/B:")
	for _, file := range dir {
		fi, _ := file.Info()
		fmt.Println("  ", file.Name(), file.IsDir(), file.Type(), fi)
	}

	// List the files from dir ABC/.
	dirABC, err := embedded.ABC.ReadDir("A/B/C")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Files in ABC/A/B/C:")
	for _, file := range dirABC {
		fi, _ := file.Info()
		fmt.Println("  ", file.Name(), file.IsDir(), file.Type(), fi)
	}

	// read the contents of the file A/B/C/abra.txt
	fileAbra, err := embedded.ABC.ReadFile("A/B/C/abra.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Contents of file A/B/C/abra.txt:\n" + string(fileAbra))
}

// walkDir walks the file system, possibly embedded and prints the file names
func walkDir(fileSystem fs.FS) {
	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		isDir := ""
		if d.IsDir() {
			isDir = " d"
		}
		fmt.Println(path + isDir)
		return nil
	})
}
