package server

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	C "locr/constant"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type Content struct {
	Result  string
	Version string
}

type ShotEncode struct {
	Base64 string `json:"base64"`
	Trim   string `json:"trim"`
}

// RecoFile 根据图片path进行识别
// 适用于图片复制场景(非截图)
func RecoFile(file *os.File) (*Content, error) {
	// 构造form-data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	// 构造request
	req, err := http.NewRequest("POST", C.URL+"file", body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
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

	var content Content
	err = json.Unmarshal(rbody, &content)
	if err != nil {
		return nil, err
	}
	return &content, nil
}

// RecoBase64 识别剪贴板的图片
// 适用于截图场景
func RecoBase64(img []byte) (*Content, error) {
	// 对截图base64编码
	body := base64.StdEncoding.EncodeToString(img)
	shotEncode := ShotEncode{
		Base64: body,
		Trim:   "\n",
	}
	byteBody, err := json.Marshal(shotEncode)
	if err != nil {
		return nil, err
	}

	// 构造request
	req, err := http.NewRequest("POST", C.URL+"base64", bytes.NewBuffer(byteBody))
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

	var content Content
	err = json.Unmarshal(rbody, &content)
	if err != nil {
		return nil, err
	}
	return &content, nil
}

func Status() error {
	resp, err := http.Get("http://lab-server:8086/status")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}
