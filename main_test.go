package main

import (
	"github.com/fogleman/gg"
	"log"
	"testing"
)

func TestDrawText(t *testing.T) {
	im, err := gg.LoadImage("./avatar.jpeg")
	if err != nil {
		t.Error(err)
		return
	}
	bounds := im.Bounds()
	//ft, err := truetype.Parse(goregular.TTF)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//face := truetype.NewFace(ft, &truetype.Options{Size: 48})
	dc := gg.NewContext(bounds.Size().X*2, bounds.Size().Y*2)
	dc.DrawImage(im, 0, 0)

	//dc.DrawCircle(50, 50, 50)
	//dc.SetRGB(0, 0, 1)
	//dc.Fill()
	//dc.SetFontFace(face)

	//dc.DrawRectangle(100, 100, 50,50)
	dc.SetRGBA255(0, 255, 255, 255)
	//dc.Fill()

	if err := dc.LoadFontFace("./SIMKAI.TTF", 12); err != nil {
		panic(err)
	}
	//dc.Rotate(45)
	dc.DrawStringAnchored("Hello, 世界111!", 0, 0, 0, 1)

	dc.DrawStringAnchored("Hello, 世界!", 100, 0, 0, 1)
	dc.DrawStringAnchored("Hello, 世界!", 100, 100, 0, 1)
	dc.SavePNG("test5.png")
}

func TestDrawWaterMark(t *testing.T) {
	wm, err := NewWaterMarkWithBaseImage("./avatar.jpeg", "hello世界啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊")
	if err != nil {
		log.Fatalln(err)
		return
	}
	wm.SavePNG("./test6.png")
	log.Println("绘制完成")
}
