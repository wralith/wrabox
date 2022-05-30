package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Surprisingly it worked.
// I guess it creates a mock struct to use in tests?
// I declare it as a variable inside the test function??
// What is going on!
type appMock struct {
	*app
}

func TestShowSnippet(t *testing.T) {
	// Dunno what i am doing, again
	var a = &appMock{}

	Convey("Snippet query by id", t, func() {
		// TODO This is not a real test, make it real :)
		// Convey("Should return 200", func() {
		// 	w := httptest.NewRecorder()
		// 	req, _ := http.NewRequest("GET", "/snippet?id=1", nil)
		// 	a.showSnippet(w, req)
		// 	// res := w.Result()

		// 	// got, _ := ioutil.ReadAll(res.Body)
		// 	// want := "A snippet with ID 3..."
		// 	// So(string(got), ShouldEqual, string(want))
		// 	So(w.Code, ShouldEqual, 200)
		// })

		Convey("Should return 404", func() {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/snippet?id=-1", nil)
			a.showSnippet(w, req)

			So(w.Code, ShouldEqual, 404)
		})
	})

}
