package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

func DrawImage(data [][]int, scale int, fc, fcc, bc color.Color) {
	file, err := os.Create("avator.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	m := image.NewRGBA(image.Rect(0, 0, len(data)*scale, len(data)*scale))
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			v := data[x][y]
			for sx := 0; sx < scale; sx++ {
				for sy := 0; sy < scale; sy++ {
					if v > 0 {
						if x < len(data)/2 {
							m.Set((x*scale)+sx, (y*scale)+sy, fc)
						} else {
							m.Set((x*scale)+sx, (y*scale)+sy, fcc)
						}
					} else {
						m.Set((x*scale)+sx, (y*scale)+sy, bc)
					}
				}
			}
		}
	}

	err = jpeg.Encode(file, m, &jpeg.Options{100})
	if err != nil {
		fmt.Println(err)
	}

}
