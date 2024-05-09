package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIntegration(t *testing.T) {
	Convey("Given two microservices for addition and multiplication", t, func() {
		Convey("When adding 2 and 3, and multiplying the result by 4", func() {
			addURL := "http://localhost:8080/add"
			multiplyURL := "http://localhost:8080/multiply"

			// Data to be sent in the request body
			data := map[string]int{"a": 2, "b": 3}
			jsonData, err := json.Marshal(data)
			So(err, ShouldBeNil)

			// Send the POST request to add
			addResp, err := http.Post(addURL, "application/json", bytes.NewBuffer(jsonData))
			So(err, ShouldBeNil)
			defer addResp.Body.Close()

			So(addResp.StatusCode, ShouldEqual, http.StatusOK)

			// Read the response body
			addBody, err := ioutil.ReadAll(addResp.Body)
			So(err, ShouldBeNil)

			// Unmarshal the response JSON
			var addResult map[string]int
			err = json.Unmarshal(addBody, &addResult)
			So(err, ShouldBeNil)

			// Verify the addition result
			So(addResult["result"], ShouldEqual, 5)

			// Send the POST request to multiply
			multiplyResp, err := http.Post(multiplyURL, "application/json", bytes.NewBuffer(jsonData))
			So(err, ShouldBeNil)
			defer multiplyResp.Body.Close()

			So(multiplyResp.StatusCode, ShouldEqual, http.StatusOK)

			// Read the response body
			multiplyBody, err := ioutil.ReadAll(multiplyResp.Body)
			So(err, ShouldBeNil)

			// Unmarshal the response JSON
			var multiplyResult map[string]int
			err = json.Unmarshal(multiplyBody, &multiplyResult)
			So(err, ShouldBeNil)

			// Verify the multiplication result
			So(multiplyResult["result"], ShouldEqual, 6)

			total := addResult["result"] + multiplyResult["result"]
			So(total, ShouldEqual, 11)
		})
	})
}
