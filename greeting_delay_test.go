package main

import "testing"

func TestGreetingWithDelay(t *testing.T) {
	expectedString := "<b>Code.education Rocks!</b>"
	resultString := greeting("Code.education Rocks!")

	if expectedString != resultString {
		t.Errorf("Expected %v but got %v", expectedString, resultString)
	}
}