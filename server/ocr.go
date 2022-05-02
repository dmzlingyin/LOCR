package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type Content struct {
	Result  string
	Version string
}

func RecoFile(file *os.File) (*Content, error) {
	// 构造form-data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	// 构造request
	req, err := http.NewRequest("POST", "http://172.17.130.166:8086/file", body)
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

func RecoBase64() bool {
	return true
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
