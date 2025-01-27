package main

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
	return
}
func (v *Video) SetUrl(url string) (newUrl string, err error) {
}

func (v *Video) GetId() int64 {
	//return 123456789
	return v.Id
}
