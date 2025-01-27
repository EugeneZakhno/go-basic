package main

import (
	"fmt"
	"strings"
)

const (
	VIDEO_PREFIX = "www.youtube.com/"
)

type Video struct {
	Id          int64
	Name        string
	Description string
	Url         string
}

// Video methods
func (v *Video) GetUrl() string {
	//return "someUrl"
	return v.Url
}

func (v *Video) SetUrl(url string) (newUrl string, err error) {
	if url == " " {
		return "", fmt.Errorf("Url cannot be empty")
	}
	if !strings.HasPrefix(url, VIDEO_PREFIX) {
		return "", fmt.Errorf("Url [%s] is not valid. Must start with [%s].", url, VIDEO_PREFIX)
	}
	v.Url = url
	return v.Url, nil
}

func (v *Video) GetId() int64 {
	//return 123456789
	return v.Id
}
