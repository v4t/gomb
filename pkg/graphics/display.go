package graphics

import (
	"image/color"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// DefaultPalette represents the available colors.
var DefaultPalette = [4]color.RGBA{
	color.RGBA{0xe0, 0xf0, 0xe7, 0xff}, // White
	color.RGBA{0x8b, 0xa3, 0x94, 0xff}, // Light gray
	color.RGBA{0x55, 0x64, 0x5a, 0xff}, // Dark gray
	color.RGBA{0x34, 0x3d, 0x37, 0xff}, // Black
}

// Display screen dimensions
const (
	ScreenWidth  = 160
	ScreenHeight = 144
)

// Display represents gameboy display.
type Display struct {
	Palette     [4]color.RGBA
	window      *pixelgl.Window
	image       *pixel.PictureData
	enabled     bool
	pixelBuffer [ScreenWidth][ScreenHeight][4]uint8
	offset      int
}

// Init initializes display.
func (display *Display) Init() {
	cfg := pixelgl.WindowConfig{
		Title:  "gomb",
		Bounds: pixel.R(0, 0, ScreenWidth, ScreenHeight),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	display.enabled = true
	display.Palette = DefaultPalette

	win.Clear(colornames.Black)
	display.window = win
	display.image = &pixel.PictureData{
		Pix:    make([]color.RGBA, ScreenWidth*ScreenHeight),
		Stride: ScreenWidth,
		Rect:   pixel.R(0, 0, ScreenWidth, ScreenHeight),
	}

	pos := win.GetPos()
	win.SetPos(pixel.ZV)
	win.SetPos(pos)
	display.updateCamera()
	win.Update()
}

// Run pixel application.
func (display *Display) Run(f func()) {
	pixelgl.Run(f)
}

// Init initializes new display.
func Init() *Display {
	display := &Display{}
	return display
}

// Closed returns true if application is closed.
func (display *Display) Closed() bool {
	return display.window.Closed()
}

// IsEnabled boolean flag.
func (display *Display) IsEnabled() bool {
	return display.enabled
}

// Enable display.
func (display *Display) Enable() {
	display.enabled = true
}

// Disable display.
func (display *Display) Disable() {
	display.enabled = false
}

// Draw adds pixel to display buffer.
func (display *Display) Draw(x byte, y byte, colorID byte) {
	if (x < 0 || x >= ScreenWidth) || (y < 0 || y >= ScreenHeight) {
		return
	}
	if display.enabled {

		color := display.Palette[colorID]
		display.pixelBuffer[x][y][0] = color.R
		display.pixelBuffer[x][y][1] = color.G
		display.pixelBuffer[x][y][2] = color.B
		display.pixelBuffer[x][y][3] = color.A
	}
}

// RenderImage is called when image frame is complete.
func (display *Display) RenderImage() {
	// display.image.Pix = display.imgBuffer
	img := make([]color.RGBA, ScreenWidth*ScreenHeight)
	for y := 0; y < ScreenHeight; y++ {
		for x := 0; x < ScreenWidth; x++ {
			index := (ScreenHeight-1-y)*ScreenWidth + x
			img[index] = color.RGBA{
				display.pixelBuffer[x][y][0],
				display.pixelBuffer[x][y][1],
				display.pixelBuffer[x][y][2],
				display.pixelBuffer[x][y][3]}
		}
	}
	display.image.Pix = img
	bg := color.RGBA{R: 202, G: 220, B: 159, A: 0xff}
	display.window.Clear(bg)

	spr := pixel.NewSprite(pixel.Picture(display.image), pixel.R(0, 0, ScreenWidth, ScreenHeight))
	spr.Draw(display.window, pixel.IM)

	display.updateCamera()
	display.window.Update()
	display.offset = 0
}

func (display *Display) updateCamera() {
	xScale := display.window.Bounds().W() / ScreenWidth
	yScale := display.window.Bounds().H() / ScreenHeight
	scale := math.Min(yScale, xScale)

	shift := display.window.Bounds().Size().Scaled(0.5).Sub(pixel.ZV)
	cam := pixel.IM.Scaled(pixel.ZV, scale).Moved(shift)
	display.window.SetMatrix(cam)
}
