package main

import (
	"flag"
	"gopkg.in/yaml.v2"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func init() {
	if flag.Lookup("conf") == nil {
		flag.StringVar(&confPath, "conf", "./config.yml", "配置文件路径")
	}
}

var confPath string

func main() {
	flag.Parse()
	confBytes, err := os.ReadFile(confPath)
	if err != nil {
		log.Fatalf("读取配置文件出错: %s\n", err)
		return
	}
	config := NewDefaultConfig()
	err = yaml.Unmarshal(confBytes, &config)
	if err != nil {
		log.Fatalf("解析配置文件出错: %s\n", err)
		return
	}
	wm, err := NewWaterMarkWithBaseImage(config.BaseImagePath, config.Text)
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = wm.LoadFontFace(config.FontPath, config.FontSize)
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = wm.SavePNG(config.SavePath)
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Printf("图片已保存至 %s\n", config.SavePath)

}
