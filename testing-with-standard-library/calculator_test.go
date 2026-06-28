package main

import (
 "testing"
)

func TestSum(t *testing.T) {
 result := Sum(2, 3)
 if result != 5 {
  t.Errorf("Expected 5, but got %d", result)
 }
}

func TestAdd(t *testing.T) {
 result := Add(2, 3)
 expected := 5
 if result != expected {
  t.Errorf("Expected %d, but got %d", expected, result)
 }
}

func TestMultiply(t *testing.T) {
 result := Multiply(2, 3)
 expected := 6
 if result != expected {
  t.Errorf("Expected %d, but got %d", expected, result)
 }
}

func TestAddTableDriven(t *testing.T) {
 tests := []struct {
  a, b, expected int
 }{
  {2, 3, 5},
  {0, 0, 0},
  {-1, 1, 0},
  {10, -5, 5},
 }

 for _, test := range tests {
  result := Add(test.a, test.b)
  if result != test.expected {
   t.Errorf("For %d+%d, expected %d, but got %d", test.a, test.b, test.expected, result)
  }
 }
}
