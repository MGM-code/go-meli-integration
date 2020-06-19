package main

import "testing"
        
func TestAdder(t *testing.T) {
    sum := Add(2, 2)
    expected := 4

    if sum != expected {
        t.Errorf("expected '%d' but got '%d'", expected, sum)
    }
}

func TestRest(t *testing.T) {
    sum := Res(2, 2)
    expected := 0

    if sum != expected {
        t.Errorf("expected '%d' but got '%d'", expected, sum)
    }
}

func TestMult(t *testing.T) {
    sum := Mul(4, 2)
    expected := 8

    if sum != expected {
        t.Errorf("expected '%d' but got '%d'", expected, sum)
    }
}

func TestDiv(t *testing.T) {
    sum := Div(2, 2)
    expected := 1

    if sum != expected {
        t.Errorf("expected '%d' but got '%d'", expected, sum)
    }
}
