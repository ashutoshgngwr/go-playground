package main

import "fmt"

type File struct {
	name string
	size int64
	mode byte
}

// if arg is passed by value, doesn't get modified. (*File)
// moreover methods can be defined for any named type
// this doesn't have to be a struct :)
func (file *File) Rename(name string) {
	file.name = name
}

func (file File) String() string {
	return fmt.Sprintf("file{name=%s, size=%d, mode=%b}", file.name, file.size, file.mode)
}

func main() {
	file := File{"some", 1292, 0}
	fmt.Println("file =", file)
	file.Rename("not some")
	fmt.Println("file =", file)
}
