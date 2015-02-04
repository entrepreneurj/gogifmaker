package main

import (
	"fmt"
	"image/gif"
	_ "image/png"
	_ "image/jpeg"
	"image/draw"
	"image/color/palette"
	"image"
	"os"
	"path/filepath"
	"log"	
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
	
	var my_gif gif.GIF
	my_gif.LoopCount =1

	var opt gif.Options
	opt.NumColors = 256

	for _, file := range files {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ".png" || filepath.Ext(file.Name()) == ".jpeg" || filepath.Ext(file.Name()) == ".jpg" {
				fmt.Println(file.Name(), file.Size(), "bytes")
				reader, err := os.Open(file.Name())
				m, _, err := image.Decode(reader)
				new_img := image.NewPaletted(m.Bounds(), palette.Plan9)
				draw.Draw(new_img, m.Bounds(), m, image.ZP, draw.Src)
				my_gif.Image = append(my_gif.Image, new_img)
				my_gif.Delay = append(my_gif.Delay, 1)
				if err != nil {
					log.Fatal(err)
				}
			}
		}

	}

	out, err := os.Create("./finalGif.gif")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = gif.EncodeAll(out, &my_gif)
	if err != nil {
		log.Fatal(err)
	}
	out.Close()
}

