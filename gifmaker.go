package main

import (
	"fmt"
	//"image/gif"
	//_ "image/png"
	//_ "image/jpeg"
	//"image"
	"os"
	"path/filepath"
	
)

// Global var holding redis connection

//var (
//	pool *redis.Pool // Redis connection pool
//)




func main() {
	fmt.Printf("Making a %s\n", "gif")
	dirname := "." + string(filepath.Separator)

	d, err := os.Open(dirname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Reading " + dirname)

	for _, file := range files {
		if file.Mode().IsRegular() {
			fmt.Println(file.Name(), file.Size(), "bytes")
		}

	}
}

