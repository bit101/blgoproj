package main

import (
	"blgoproj/base"

	"github.com/bit101/blgo"
)

const mode = "image"

func main() {
	switch mode {
	case "image":
		width := 800.0
		height := 800.0
		setup(width, height)
		base.RenderImage(width, height, render)
	case "gif":
		width := 400.0
		height := 400.0
		time := 10
		fps := 30
		setup(width, height)
		base.RenderGif(width, height, time, fps, render)
	case "video":
		width := 1280.0
		height := 720.0
		time := 10
		fps := 30
		base.RenderVideo(width, height, time, fps, render)
	}
}

func setup(width, height float64) {

}

func render(surface *blgo.Surface, width, height, percent float64) {
	surface.ClearRGB(percent, 0, 1-percent)

}
