package gfx

import (
	"image/color"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// DefaultPalette represents the selectable colors in the DMG. We use a greenish
// set of colors and Alpha is always 0xff since we won't use transparency.
var DefaultPalette = [4]color.RGBA{
	color.RGBA{155, 188, 15, 0xff}, // White
	color.RGBA{139, 172, 15, 0xff}, // Light gray
	color.RGBA{48, 98, 48, 0xff},   // Dark gray
	color.RGBA{15, 56, 15, 0xff},   // Black
}

// Screen dimensions that will be used in several different places.
const (
	ScreenWidth  = 160
	ScreenHeight = 144
)

// Display is
type Display struct {
	enabled bool
	window  *pixelgl.Window
	image   *pixel.PictureData

	// Palette will contain our R, G, B and A components for each of the four
	// potential colors the Game Boy can display.
	Palette [4]color.RGBA

	// The texture buffer for the current frame, and the current offset where
	// new pixel data should be written into that buffer.
	buffer    []byte
	imgBuffer []color.RGBA
	offset    int
}

func (pd *Display) updateCamera() {
	xScale := pd.window.Bounds().W() / ScreenWidth
	yScale := pd.window.Bounds().H() / ScreenHeight
	scale := math.Min(yScale, xScale)

	shift := pd.window.Bounds().Size().Scaled(0.5).Sub(pixel.ZV)
	cam := pixel.IM.Scaled(pixel.ZV, scale).Moved(shift)
	pd.window.SetMatrix(cam)
}

// Init initializes display.
func (pd *Display) Init() {
	cfg := pixelgl.WindowConfig{
		Title:  "gomb",
		Bounds: pixel.R(0, 0, ScreenWidth, ScreenHeight),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	bufLen := ScreenWidth * ScreenHeight * 4
	pd.buffer = make([]byte, bufLen)
	pd.imgBuffer = make([]color.RGBA, ScreenWidth*ScreenHeight)
	pd.Palette = DefaultPalette

	win.Clear(colornames.Black)
	pd.window = win
	pd.image = &pixel.PictureData{
		Pix:    make([]color.RGBA, ScreenWidth*ScreenHeight),
		Stride: ScreenWidth,
		Rect:   pixel.R(0, 0, ScreenWidth, ScreenHeight),
	}

	pos := win.GetPos()
	win.SetPos(pixel.ZV)
	win.SetPos(pos)
	pd.updateCamera()
	win.Update()
}

// Run pixel application.
func (pd *Display) Run(f func()) {
	pixelgl.Run(f)
}

// Closed returns true if application is closed.
func (pd *Display) Closed() bool {
	return pd.window.Closed()
}

// Enabled returns true if screen is enabled.
func (pd *Display) Enabled() bool {
	return pd.enabled
}

// Enable screen.
func (pd *Display) Enable() {
	pd.enabled = true
}

// Disable screen.
func (pd *Display) Disable() {
	pd.enabled = false
}

// Write adds a new pixel (a mere index into a palette) to the texture buffer.
func (pd *Display) Write(colorIndex uint8) {
	if pd.enabled {
		y := ScreenHeight - (pd.offset / ScreenWidth) - 1
		pixelIndex := y*ScreenWidth + pd.offset%ScreenWidth
		pixelColor := pd.Palette[colorIndex]
		if pixelIndex >= 0 && pixelIndex < len(pd.imgBuffer) {
			pd.imgBuffer[pixelIndex] = color.RGBA{pixelColor.R, pixelColor.G, pixelColor.B, pixelColor.A}
		}
		pd.offset++
	}
}

// DrawDisplayImage is called when image frame is complete.
func (pd *Display) DrawDisplayImage() {
	pd.image.Pix = pd.imgBuffer

	bg := color.RGBA{R: 202, G: 220, B: 159, A: 0xff}
	pd.window.Clear(bg)

	spr := pixel.NewSprite(pixel.Picture(pd.image), pixel.R(0, 0, ScreenWidth, ScreenHeight))
	spr.Draw(pd.window, pixel.IM)

	pd.updateCamera()
	pd.window.Update()
	pd.offset = 0
}
