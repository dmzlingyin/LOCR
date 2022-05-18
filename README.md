<p align="center">
     <img src="https://github.com/dmzlingyin/LOCR/blob/main/docs/demo.gif" width="506" alt="fx preview">
</p>

# Introduction

English | [简体中文](README_zh-CN.md)

LOCR(Lightweight OCR) is a tool that can detect the image and extract the text automatically when the clipboard changed.

# Features

* Cross platform supports: macOS, Linux (X11), and Windows.
* Chinese and English support(feel free to add your language).

# Dependency

* go 1.16+
* OCR server(not nesessary for offline version):
    - [PaddleOCR](https://github.com/PaddlePaddle/PaddleOCR)(Recommend)
    - [tesseract](https://github.com/tesseract-ocr/tesseract)
    - [EasyOCR](https://github.com/JaidedAI/EasyOCR)
    - [ocrserver](https://github.com/otiai10/ocrserver)

# Installation

```go
git clone https://github.com/dmzlingyin/LOCR.git
cd LOCR
go build
./install.sh
```
then you can type locr on you terminal to use it.

# Examples
If you use locr normaly, just run it. If you just want to extact an image, see follow:
```shell
cat xxx.jpg | locr -f

# save to file
cat xxx.jpg | locr -f > text.txt
```

# TODO

- [ ] hotkey
- [ ] offline version