package main

import "testing"

func HelloTom() string {
	return "Jerry"
}

func TestHelloTom(t *testing.T) {
	output := HelloTom()
	Expectoutput := "Tom"
	if output != Expectoutput {
		t.Errorf("Expected %s do not match actual %s", Expectoutput, output)
	}
}

