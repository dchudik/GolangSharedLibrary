package main

import (
	"testing"
)

type TestData struct {
	FirstDigit      int
	NextSimpleDigit int
}

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
	var datasForTests = []TestData{
		TestData{
			FirstDigit:      2,
			NextSimpleDigit: 3,
		},
		TestData{
			FirstDigit:      5,
			NextSimpleDigit: 7,
		},
	}
	for _, v := range datasForTests {
		if NextSimple(v.FirstDigit) != v.NextSimpleDigit {
			t.Errorf("For %d next simple digit is %d, but function say %d", v.FirstDigit, v.NextSimpleDigit, NextSimple(v.FirstDigit))
		}
	}
}

func TestByTK(t *testing.T) {
	t.Log("Test IsSimple function on simple digit in TK")
	SimpleDigit := 11
	if IsSimple(SimpleDigit) != true {
		t.Errorf("%d simple, but function say that not simple", SimpleDigit)
	}
	t.Log("Test IsSimple function on not simple digit in TK")
	notSimpleDigit := 14
	if IsSimple(notSimpleDigit) != false {
		t.Errorf("%d not simple, but function say that simple", notSimpleDigit)
	}
	t.Log("Test NextSimple function on data in TK")
	var datasForTests = []TestData{
		TestData{
			FirstDigit:      11,
			NextSimpleDigit: 13,
		},
		TestData{
			FirstDigit:      14,
			NextSimpleDigit: 17,
		},
	}
	for _, v := range datasForTests {
		if NextSimple(v.FirstDigit) != v.NextSimpleDigit {
			t.Errorf("For %d next simple digit is %d, but function say %d", v.FirstDigit, v.NextSimpleDigit, NextSimple(v.FirstDigit))
		}
	}
}
