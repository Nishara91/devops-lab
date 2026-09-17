package main

import "testing"

func TestHealthMessage(t *testing.T) {
	expected := "System is healthy and operational!"
	actual := HealthMessage()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}