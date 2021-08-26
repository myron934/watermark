package main

import (
	"github.com/fogleman/gg"
	"github.com/golang/freetype"
	"golang.org/x/image/font"
	"image"
	"image/color"
	"image/draw"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//readJpeg("./avatar.jpeg")
	//test1()
	testWriteString()
}

func writeStringWaterMark(src image.Image, markString string) (image.Image, error) {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)
	draw.Draw(dst, bounds, src, image.Point{}, draw.Src)

	return nil, nil
}

func testWriteString() {
	fontBytes, err := ioutil.ReadFile("SIMKAI.TTF")
	if err != nil {
		log.Fatal(err)
		return
	}
	ft, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Fatal(err)
		return
	}
	bounds := image.Rect(0, 0, 200, 200)
	dst := image.NewRGBA(bounds)
	blue := color.RGBA{B: 255, A: 255}
	//dst中绘制src的区域大小
	rect := image.Rect(0, 0, bounds.Max.X, bounds.Max.Y)
	// sp: 从src的sp位置开始绘制到 dst中. (0,0) 则为全部绘制
	draw.Draw(dst, rect, &image.Uniform{C: blue}, image.Point{}, draw.Src)

	f := freetype.NewContext()
	//设置分辨率
	f.SetDPI(72)
	//设置字体
	f.SetFont(ft)
	//设置尺寸
	f.SetFontSize(26)
	//绘画区域大小
	f.SetClip(bounds)
	//设置输出的图片
	f.SetDst(dst)
	//设置字体颜色(红色)
	f.SetSrc(image.NewUniform(color.RGBA{R: 255, A: 255}))
	f.SetHinting(font.HintingNone)
	//设置字体的位置

	pt := freetype.Pt(0, int(f.PointToFixed(26))>>6)
	//pt := freetype.Pt(0, 0)
	_, err = f.DrawString("hello,水印\n11111111111111111111111111\n1111111111111111111", pt)
	if err != nil {
		log.Fatal(err)
	}
	writePng("test4.png", dst)
}

func test1() {
	src, err := readImage("./avatar.jpeg")
	if err != nil {
		log.Fatal(err)
		return
	}
	bounds := src.Bounds()
	m := image.NewRGBA(image.Rect(0, 0, bounds.Max.X+100, bounds.Max.Y+100))
	writePng("./test1.png", m)

	draw.Draw(m, m.Bounds(), src, image.Point{}, draw.Src)
	writePng("./test2.png", m)

	blue := color.RGBA{B: 255, A: 255}
	//dst中绘制src的区域大小
	rect := image.Rect(50, 50, bounds.Max.X, bounds.Max.Y)
	// sp: 从src的sp位置开始绘制到 dst中. (0,0) 则为全部绘制
	draw.Draw(m, rect, &image.Uniform{C: blue}, image.Point{}, draw.Src)
	writePng("./test3.png", m)

}

func writePng(path string, m image.Image) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	return png.Encode(out, m)
}

func readImage(path string) (m image.Image, err error) {
	reader, err := os.Open(path)
	if err != nil {
		return
	}
	defer reader.Close()
	m, _, err = image.Decode(reader)

	return
}

func rotateImage() {
	dc := gg.NewContext(1000, 1000)
	dc.DrawCircle(500, 500, 400)
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	dc.SavePNG("out.png")
}
