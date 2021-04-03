package base

import (
	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/util"
)

type RenderFunc func(surface *blgo.Surface, width, height, percent float64)
type SetupFunc func(width, height float64)

var (
	surface    *blgo.Surface
	renderFunc RenderFunc
)

const framesDir = "frames" // Must Exist!!!

func RenderImage(width, height, percent float64, setup SetupFunc, render RenderFunc) {
	outFileName := "out.png"
	surface = blgo.NewSurface(width, height)
	setup(width, height)
	render(surface, width, height, percent)
	surface.WriteToPNG(outFileName)
	util.ViewImage(outFileName)
}

func RenderGif(width, height float64, seconds, fps int, setup SetupFunc, render RenderFunc) {
	outFileName := "out.gif"
	setup(width, height)
	renderAnim(width, height, seconds, fps, render)
	util.ConvertToGIF(framesDir, outFileName, float64(fps))
	util.ViewImage(outFileName)
}

func RenderVideo(width, height float64, seconds, fps int, setup SetupFunc, render RenderFunc) {
	outFileName := "out.mp4"
	setup(width, height)
	renderAnim(width, height, seconds, fps, render)
	util.ConvertToYoutube(framesDir, outFileName, fps)
	util.VLC(outFileName)
}

func renderAnim(width, height float64, seconds, fps int, render RenderFunc) {
	framesDir := "frames" // MUST EXIST!!!
	surface = blgo.NewSurface(width, height)
	renderFunc = render
	animation := anim.NewAnimation(surface, fps*seconds, framesDir)
	animation.Render(internalRender)
}

func internalRender(percent float64) {
	renderFunc(surface, surface.Width, surface.Height, percent)
}
