package main

import (
 "github.com/stretchr/testify/assert"
 "testing"
)


func TestSum(t *testing.T) {
 result := Sum(2, 3)
 assert.Equal(t, 5, result, "Expected 5, but got %d", result)
}
