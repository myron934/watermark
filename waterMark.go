package main

import (
	"github.com/fogleman/gg"
	"golang.org/x/image/font"
	"image"
	"image/color"
)

type WaterMark struct {
	width     int
	height    int
	text      string
	color     color.Color
	fontFace  font.Face
	baseImage image.Image
}

//NewWaterMark 新建一个 WaterMark, 采用默认的配置
func NewWaterMark(width int, height int, text string) *WaterMark {
	mark := WaterMark{width: width, height: height, text: text}
	mark.SetRGBA255(162, 162, 162, 180)
	//mark.SetRGBA255(255, 255, 255, 255)
	mark.LoadFontFace("./SIMKAI.TTF", 14)
	return &mark
}

func NewWaterMarkWithBaseImage(baseImagePath string, text string) (*WaterMark, error) {
	baseImage, err := gg.LoadImage(baseImagePath)
	if err != nil {
		return nil, err
	}
	bounds := baseImage.Bounds()

	wm := NewWaterMark(bounds.Size().X, bounds.Size().Y, text)
	wm.SetBaseImage(baseImage)
	return wm, nil
}

func (wm *WaterMark) SetRGBA255(r, g, b, a int) {
	wm.color = color.NRGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
}

func (wm *WaterMark) LoadFontFace(path string, points float64) error {
	face, err := gg.LoadFontFace(path, points)
	if err != nil {
		return err
	}
	wm.fontFace = face
	return nil
}

func (wm *WaterMark) Draw() image.Image {
	dc := gg.NewContext(wm.width, wm.height)
	dc.SetFontFace(wm.fontFace)
	dc.SetColor(wm.color)

	dc.DrawImage(wm.baseImage, 0, 0)
	mark := wm.drawMark()
	dc.Rotate(0.6)

	//dc.DrawImageAnchored(mark, 0, wm.height, 0, 1)
	dc.DrawImageAnchored(mark, 0, -wm.height, 0, 0)
	return dc.Image()
}

func (wm *WaterMark) SavePNG(path string) error {
	return gg.SavePNG(path, wm.Draw())
}

func (wm *WaterMark) drawMark() image.Image {
	width := wm.width * 2
	height := wm.height * 2
	dc := gg.NewContext(wm.width*2, wm.height*2)
	dc.SetFontFace(wm.fontFace)
	dc.SetColor(wm.color)
	textWidth, textHeight := dc.MeasureString(wm.text)
	spanX := float64(width / 20)
	spanY := float64(height / 20)
	var beginX float64 = 0
	flag := false
	for beginX <= float64(width) {
		var beginY float64 = 0
		for beginY <= float64(height) {
			var extraX float64 = 0
			if flag {
				extraX = spanX
			}
			flag = !flag
			dc.DrawStringAnchored(wm.text, beginX+extraX, beginY, 0, 1)
			beginY += spanY + textHeight
		}
		beginX += spanX + textWidth
	}
	return dc.Image()

}

func (wm *WaterMark) Width() int {
	return wm.width
}

func (wm *WaterMark) SetWidth(width int) {
	wm.width = width
}

func (wm *WaterMark) Height() int {
	return wm.height
}

func (wm *WaterMark) SetHeight(height int) {
	wm.height = height
}

func (wm *WaterMark) Text() string {
	return wm.text
}

func (wm *WaterMark) SetText(text string) {
	wm.text = text
}

func (wm *WaterMark) BaseImage() image.Image {
	return wm.baseImage
}

func (wm *WaterMark) SetBaseImage(baseImage image.Image) {
	wm.baseImage = baseImage
}

func (wm *WaterMark) LoadBaseImage(path string) error {
	im, err := gg.LoadImage(path)
	if err != nil {
		return err
	}
	wm.baseImage = im
	return nil
}
