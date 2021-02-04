package gca_bdd

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
 * @author  wjj
 * @date  2020/9/2 1:14 上午
 * @description bdd
 */

func TestSpec(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Given 2 even numbers", t, func() {
		a := 2
		b := 4
		Convey("When add the two numbers", func() {
			c := a + b
			Convey("Then the result is still even", func() {
				So(c%2, ShouldEqual, 0)
			})
		})
	})
}
