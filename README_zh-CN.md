<p align="center">
     <img src="https://github.com/dmzlingyin/LOCR/blob/main/docs/demo.gif" width="506" alt="fx preview">
</p>

# 这个项目是什么？

简体中文 | [English](README.md)

LOCR(Lightweight OCR)是一款轻量级的文字识别工具, 结合第三方截图工具, 可以快速的对图片文字进行识别。

# 为什么有这个项目

在日常学习的工作中, 难免会遇到一些文字的复制粘贴任务。但由于一些限制，我们无法复制想要的文字，只能一个字一个字的敲出来。而随着近几年OCR技术的成熟，越来越多的网站和工具开始支持图片的文字识别功能，可以快速的提取出我们想要的文字。
但就我个人而言，目前的OCR工具存在着诸多的弊端：

1. 免费使用次数有限
2. 对于网站类识别, 操作过于繁琐（先截图, 再保存, 再打开网站, 再将图片拖到浏览器, 再选中识别出来的文字, 最后粘贴）
3. 对于客户端类识别, 软件功能过于复杂(当然, 仅仅是个人观点), 而且很少支持跨平台(个人电脑为Arch)

# 本项目的目标

本项目预期实现一个轻量但强大的命令行OCR工具, 简单快速地(选中识别区域->复制到剪贴板->粘贴到你想要的任何地方)(需结合第三方截图软件, e.g. Windows: [Snipaste](https://www.snipaste.com/), Linux: [Flameshot](https://flameshot.org/))实现图片文字提取及识别。

# 依赖
1. go1.16+
2. OCR服务器(离线版本不需要), 多种OCR实现可选：
   1. [PaddleOCR](https://github.com/PaddlePaddle/PaddleOCR)(推荐)
   2. [tesseract](https://github.com/tesseract-ocr/tesseract)
   3. [EasyOCR](https://github.com/JaidedAI/EasyOCR)
   4. [ocrserver](https://github.com/otiai10/ocrserver)

# 使用
```go
git clone https://github.com/dmzlingyin/LOCR.git
cd LOCR
go mod tidy
go build
./locr
```

# 待实现的功能

- [x] 结合Unix管道符实现图片的读入与识别
- [ ] 支持离线识别
- [ ] windows 剪贴板截图无法识别bug
