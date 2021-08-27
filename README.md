# watermark 证件加水印

## 使用步骤

1. 在 https://github.com/myron934/watermark/releases 下载项目, 并解压

2. 将需要添加水印的图片复制到项目目录. 例如 "src.png"

3. 用文本编辑器打开 config.yml, 修改base_image_path的值为 ./src.png(你的图片路径)

   ```yaml
   # 照片地址
   base_image_path: "./src.png"
   # 水印文字
   text: "仅限用于办理护照"
   # 字体 默认为楷体
   font_path: "./SIMKAI.TTF"
   # 字体大小
   font_size: 14
   # 保存图片地址
   save_path: "./result.png"
   
   ```

4. 运行 watermark 即可

