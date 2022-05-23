package server

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"

	C "locr/constant"
)

// OCR请求参数列表
type PaddleOCR struct {
	Key   []string `json:"key"`
	Value []string `json:"value"`
}

// OCR识别结果结构体
type Result struct {
	ErrNo   int      `json:"err_no"`
	ErrMsg  string   `json:"err_msg"`
	Key     []string `json:"key"`
	Value   []string `json:"value"`
	Tensors []string `json:"tensors"`
}

// RecoBase64请求服务端, 对base64格式的图片进行OCR识别
func RecoBase64(img []byte) (*Result, error) {
	// 对图片base64编码
	body := base64.StdEncoding.EncodeToString(img)
	paddleEncode := PaddleOCR{
		Key:   []string{"image"},
		Value: []string{body},
	}
	byteBody, err := json.Marshal(paddleEncode)
	if err != nil {
		return nil, err
	}

	// 构造request
	req, err := http.NewRequest("POST", C.URL+"ocr/prediction", bytes.NewBuffer(byteBody))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rbody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res Result
	err = json.Unmarshal(rbody, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
