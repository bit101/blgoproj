package main

import (
	"blgoproj/base"

	"github.com/bit101/blgo"
)

const mode = "image"

func main() {
	base.RenderImage(800, 800, 1, setup, render)
	// base.RenderGif(400, 400, 30, 30, setup, render)
	// base.RenderVideo(1280, 720, 10, 30, setup, render)
}

func setup(width, height float64) {

}

func render(surface *blgo.Surface, width, height, percent float64) {
	surface.ClearRGB(percent, 0, 1-percent)

}
