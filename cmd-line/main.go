package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}

	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}

	defer dst.Close()

	written, err = io.Copy(dst, src)
	dst.Close()
	src.Close()
	return
}

func main() {
	var path string
	_, err := fmt.Scanln(path)
	if err != nil {
		fmt.Printf("Error happend %v", err)
	}

	dest, err := os.Create(path)
	if err != nil {
		fmt.Printf("Error happend %v", err)
	}

	written, err := io.WriteString(dest, "hi")

	fmt.Println(written)

	dest.Close()
	return

}
