package main

import (
	"reflect"
	"testing"
)

func TestReading(t *testing.T) {
	numbers := [][]string{
		{"5+5 10"},
		{"7+3 10"}}

	got := readFile("small.csv")
	want := numbers
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
