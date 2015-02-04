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

// Converts the opened image to the required form of *Image.Paletted by
// drawing into a new *Image.Paletted instance
func convertImage(img image.Image) *image.Paletted {

	new_img := image.NewPaletted(img.Bounds(), palette.Plan9)
	draw.Draw(new_img, img.Bounds(), img, image.ZP, draw.Src)
	return new_img
}


func main() {
	fmt.Printf("Making a gif\n")
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
	
	// Setting up the empty GIF struct
	var my_gif gif.GIF
	my_gif.LoopCount =1

	// Iterating through each file and adding it to the GIF if it is 
	// an accepted image file
	for _, file := range files {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ".png" || filepath.Ext(file.Name()) == ".jpeg" || filepath.Ext(file.Name()) == ".jpg" {
				fmt.Println(file.Name(), file.Size(), "bytes")
				reader, err := os.Open(file.Name())
				fmt.Println("decoding image")
				in_image, _, err := image.Decode(reader)
				fmt.Println("converting image")
				new_img := convertImage(in_image)
				fmt.Println("Adding image to gif")
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

