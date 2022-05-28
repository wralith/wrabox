package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestShowSnippet(t *testing.T) {
	// Dunno what i am doing, again

	Convey("Snippet query by id", t, func() {
		Convey("Should return query id in body", func() {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/snippet?id=666", nil)
			showSnippet(w, req)
			res := w.Result()

			got, _ := ioutil.ReadAll(res.Body)
			want := "A snippet with ID 666..."
			So(string(got), ShouldEqual, string(want))
			So(w.Code, ShouldEqual, 200)
		})

		Convey("Should return 404", func() {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/snippet?id=-1", nil)
			showSnippet(w, req)

			So(w.Code, ShouldEqual, 404)
		})
	})

}
