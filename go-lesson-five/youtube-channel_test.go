package main

import (
	_ "errors"
	"fmt"
	_ "fmt"
	"github.com/stretchr/testify/assert" // Самая популярная библиотека для тестирования
	"testing"
)

// import "testing"

func TestGetId_ProvidedId(t *testing.T) {
	var expectedId int64 = 12345
	video := &Video{Id: expectedId}

	actualId := video.GetId()

	assert.True(t, actualId == expectedId)
	assert.Equal(t, expectedId, actualId)
}

func TestGetUrl_ProvidedUrl(t *testing.T) {
	expected := "testUrl"
	video := &Video{Url: expected}
	actual := video.GetUrl()
	// assert.Samef(t, expected, actual, "Expected memory address: [%p), actual memory address: [%p]", &expected, &actual)
	assert.Equal(t, expected, actual)
	assert.True(t, expected == actual)
}

func TestGetUrl_DefaultUrl(t *testing.T) {
	video := &Video{}
	assert.Empty(t, video.GetUrl(), "Default value for Video should be empty")
}

func TestSetUrl_ValidUrl(t *testing.T) {
	expectedUrl := VIDEO_PREFIX + "testUrl"
	video := &Video{}
	urlFromFunc, err := video.SetUrl(expectedUrl)
	assert.NoError(t, err)
	assert.Equal(t, expectedUrl, video.Url, "video.url doesn't match expected")
	assert.Equalf(t, expectedUrl, urlFromFunc, "Url [%s] from SetUrl() doesn't match expected [%s]", urlFromFunc, expectedUrl)
}

func TestSetUrl_InvalidUrl(t *testing.T) {
	url := "testUrl"
	video := &Video{}
	expectedErrorMsg := fmt.Sprintf("Url [%s] is not valid. Must start with [%s].", url, VIDEO_PREFIX)
	actualUrlFromFunc, err := video.SetUrl(url)
	assert.Error(t, err)
	assert.EqualError(t, err, expectedErrorMsg)
	assert.Empty(t, actualUrlFromFunc)
}

func TestGetUrl_EmptyUrl(t *testing.T) {
	video := &Video{}
	urlFromFunc, err := video.SetUrl("")
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Empty(t, urlFromFunc)
}
