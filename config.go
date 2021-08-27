package main

type Config struct {
	BaseImagePath string  `yaml:"base_image_path"`
	Text          string  `yaml:"text"`
	FontPath      string  `yaml:"font_path"`
	FontSize      float64 `yaml:"font_size"`
	SavePath      string  `yaml:"save_path"`
}

func NewConfig(baseImagePath string, text string, fontPath string, fontSize float64, savePath string) *Config {
	return &Config{BaseImagePath: baseImagePath, Text: text, FontPath: fontPath, FontSize: fontSize, SavePath: savePath}
}

func NewDefaultConfig() *Config {
	return NewConfig("./src.png", "hello,世界", "./SIMKAI.TTF", 14, "./result.png")
}
