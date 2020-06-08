package main

import (
	"fmt"
	"os"

	"github.com/fhirt/go-go-gadget/ascii"
)

const (
	multiplier   = 2
	fileLocation = "C:\\Users\\fabio\\go\\src\\github.com\\fhirt\\go-go-gadget\\ascii\\resources\\"
)

func main() {
	processImage(fileLocation + os.Args[1])
}

func processImage(path string) {
	img, err := ascii.OpenImage(path)
	if err != nil {
		fmt.Println("error while processing image: ", err)
		return
	}

	fmt.Println("---------- Average Colors ---------")
	avgArtwork := ascii.MakeAsciiArt(img, ascii.AverageConverter, ascii.StandardConverter)
	printImage(avgArtwork)

	fmt.Println("---------- MinMax Colors ---------")
	minMaxArtwork := ascii.MakeAsciiArt(img, ascii.MinMaxConverter, ascii.StandardConverter)
	printImage(minMaxArtwork)

	fmt.Println("---------- Luminosity Colors ---------")
	luminosityArtwork := ascii.MakeAsciiArt(img, ascii.LuminosityConverter, ascii.StandardConverter)
	printImage(luminosityArtwork)

}

func printImage(artwork [][]string) {
	for _, line := range artwork {
		for _, pixel := range line {
			for i := 0; i < multiplier; i++ {
				fmt.Print(pixel)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
