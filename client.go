package vulfocus

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	_start      = "start"
	_stop       = "stop"
	_delete     = "delete"
	successCode = 200
)

type Image struct {
	Name    string `json:"image_name"`
	VulName string `json:"image_vul_name"`
	Desc    string `json:"image_desc"`
}

type Exposed struct {
	Host string
	Port string
}

type Result struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
}

type ImageResult struct {
	Status int     `json:"status"`
	Data   []Image `json:"data"`
	Msg    string  `json:"msg"`
}

type Client struct {
	addr     string
	username string
	licence  string
}

func NewClient(addr, username, licence string) *Client {
	return &Client{
		addr:     addr,
		username: username,
		licence:  licence,
	}
}

func (r Client) GetImages() (err error, images []Image) {
	resp, err := http.Get(fmt.Sprintf("%s/api/imgs/operation?username=%s&licence=%s", r.addr, r.username, r.licence))
	if err != nil {
		return err, nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return errors.New(string(body)), nil
	}
	var result ImageResult
	if err := json.Unmarshal(body, &result); err != nil {
		return err, nil
	}
	if result.Status != successCode {
		return errors.New(result.Msg), nil
	}
	return nil, result.Data
}

func (r Client) operation(imageName, requisition string) (err error, result Result) {
	data := url.Values{}
	data.Set("username", r.username)
	data.Set("licence", r.licence)
	data.Set("image_name", imageName)
	data.Set("requisition", requisition)
	resp, err := http.PostForm(fmt.Sprintf("%s/api/imgs/operation", r.addr), data)
	if err != nil {
		return err, result
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return errors.New(string(body)), result
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return err, result
	}
	if result.Status != successCode {
		return errors.New(result.Msg), result
	}
	return nil, result
}

func (r Client) Start(imageName string) (err error, exposed Exposed) {
	err, result := r.operation(imageName, _start)
	if err != nil {
		return err, exposed
	}
	data := result.Data.(map[string]interface{})
	exposed.Host = data["host"].(string)
	exposed.Port = data["port"].(string)
	return nil, exposed
}

func (r Client) Stop(imageName string) error {
	err, _ := r.operation(imageName, _stop)
	if err != nil {
		return err
	}
	return nil
}

func (r Client) Delete(imageName string) error {
	err, _ := r.operation(imageName, _delete)
	if err != nil {
		return err
	}
	return nil
}
