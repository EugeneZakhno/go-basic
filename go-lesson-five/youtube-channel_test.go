package main

import (
	_ "errors"
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

}

func TestSetUrl_InvalidUrl(t *testing.T) {

}

func TestGetUrl_EmptyUrl(t *testing.T) {

}
