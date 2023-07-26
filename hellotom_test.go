package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func HelloTom() string {
	return "Tom"
}

func TestHelloTom(t *testing.T) {
	output := HelloTom()
	Expectoutput := "Tom"
	assert.Equal(t,output,Expectoutput)
}

