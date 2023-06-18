# Embedding files in the app

[How to Use go:embed in Go](https://blog.jetbrains.com/go/2021/06/09/how-to-use-go-embed-in-go-1-16/)

[embed](https://pkg.go.dev/embed) docs

[fs](https://pkg.go.dev/io/fs@master)

[example](https://www.reddit.com/r/golang/comments/kiuvbq/is_there_any_way_to_walk_an_embedfs_using_the/)

[fs walk](https://bitfieldconsulting.com/golang/filesystems)

## Discussions

[//go:embed \<path to file in parent directory\> doesn't work #46056](https://github.com/golang/go/issues/46056)

_The best practice here is to create a package in parent level with go:embed var in there; then reference to this var in the child level packages._

## `embed` lessons learned

```
// embedded.go

//go:embed dir1
//go:embed dir2
var Dir12 embed.FS

//go:embed dir3
var Dir3 embed.FS

//go:embed A/B/C
var ABC embed.FS
```

Above will embed the `dir1` and `dir2` in the `Dir12` variable, `dir3`in `Dir3` and `A/B/C` in `ABC`.

Note: dirs are relative to the directory where the file doing the embedding is found. It is not possible to embed dirs or files from the parent directiories (patterns starting with `.` or `..` are not allowed.

In the demo project `goembed-demo`, these dirs are, along with the file `embedded.go`, in the dir `pkg/embedded`.

```
goembed-demo % tree                                                                                                        [main L|✚1…2]
.
├── pkg
│   └── embedded
│       ├── A
│       │   └── B
│       │       └── C
│       │           └── abra.txt
│       ├── dir1
│       │   ├── hello.txt
│       │   └── hello2.txt
│       ├── dir2
│       │   └── hello2.txt
│       ├── dir3
│       │   └── hello3.txt
│       └── embedded.go
├── README.md
├── go.mod
└── main.go
```

In `main.go` we import the package `embedded`...

```
package main

import (
    "fmt"
    "io/fs"
    "log"

    "github.com/rudifa/goembed-demo/pkg/embedded"
)

func main() {
...
    fmt.Println("Walk the embedded directory Dir12")
    walkDir(embedded.Dir12)

    fmt.Println("Walk the embedded directory Dir3")
    walkDir(embedded.Dir3)

    fmt.Println("Walk the embedded directory ABC")
    walkDir(embedded.ABC)
...
}

```

Above `walkDir` calls produce this:

```
Walk the embedded directory Dir12
. d
dir1 d
dir1/hello.txt
dir1/hello2.txt
dir2 d
dir2/hello2.txt
Walk the embedded directory Dir3
. d
dir3 d
dir3/hello3.txt
Walk the embedded directory ABC
. d
A d
A/B d
A/B/C d
A/B/C/abra.txt
```

... with

```
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

```

Finally, we can read the contents of a file

```
    // read the contents of the file dir3/file3.txt

    file3, err := embedded.Dir3.ReadFile("dir3/hello3.txt")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Contents of dir3/file3.txt:")
    fmt.Println(string(file3))

    // read the contents of the file A/B/C/abra.txt
    fileAbra, err := embedded.ABC.ReadFile("A/B/C/abra.txt")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Contents of fileAbra:\n" + string(fileAbra))
```

which yields

```
Contents of dir3/file3.txt:
hello3

Contents of file A/B/C/abra.txt:
Ô flots abracadabrantesques
Prenez mon cœur, qu'il soit sauvé.
Ithyphalliques et pioupiesques
Leurs insultes l'ont dépravé !
                 A.Rimbaud
```
