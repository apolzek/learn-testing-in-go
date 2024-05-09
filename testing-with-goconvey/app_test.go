package main

import (
 . "github.com/smartystreets/goconvey/convey"
 "testing"
)


func TestSum(t *testing.T) {
 Convey("Sum function should add two numbers", t, func() {
  result := Sum(2, 3)
  Convey("The result should be 5", func() {
   So(result, ShouldEqual, 5)
  })
 })
}
