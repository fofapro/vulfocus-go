package vulfocus

import (
	"fmt"
	"testing"
)

const (
	addr     = "http://vulfocus.fofa.so"
	username = ""
	licence  = ""
)

func TestClient_GetImages(t *testing.T) {
	client := NewClient(addr, username, licence)
	err, images := client.GetImages()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("get %d images", len(images))
}

func TestClient_Start(t *testing.T) {
	client := NewClient(addr, username, licence)
	err, images := client.GetImages()
	if err != nil {
		t.Error(err)
	}
	err, exposed := client.Start(images[0].Name)
	if err != nil {
		t.Error(err)
	}
	println(exposed.Host, exposed.Port)
}

func TestClient_Stop(t *testing.T) {
	client := NewClient(addr, username, licence)
	err, images := client.GetImages()
	if err != nil {
		t.Error(err)
	}
	err = client.Stop(images[0].Name)
	if err != nil {
		t.Error(err)
	}
}

func TestClient_Delete(t *testing.T) {
	client := NewClient(addr, username, licence)
	err, images := client.GetImages()
	if err != nil {
		t.Error(err)
	}
	err = client.Delete(images[0].Name)
	if err != nil {
		t.Error(err)
	}
}
