package main

import (
	"testing"
)

func TestIsSimple(t *testing.T) {
	var simpleDigits = []int{2, 3, 5, 7, 11, 13, 17, 19}
	for _, v := range simpleDigits {
		if IsSimple(v) != true {
			t.Errorf("%d is simple, but function say not", v)
		}
	}
	var notSimpleDigits = []int{0, 1, 4, 6, 8, 9, 10, 12}
	for _, v := range notSimpleDigits {
		if IsSimple(v) != false {
			t.Errorf("%d not simple, but function say that simple", v)
		}
	}
}

func TestNextSimple(t *testing.T) {
	type TestData struct {
		SimpleDigit     int
		NextSimpleDigit int
	}
	var datasForTests = []TestData{
		TestData{
			SimpleDigit:     2,
			NextSimpleDigit: 3,
		},
		TestData{
			SimpleDigit:     5,
			NextSimpleDigit: 7,
		},
	}
	for _, v := range datasForTests {
		if NextSimple(v.SimpleDigit) != v.NextSimpleDigit {
			t.Errorf("For %d next simple digit is %d, but function say %d", v.SimpleDigit, v.NextSimpleDigit, NextSimple(v.SimpleDigit))
		}
	}
}
