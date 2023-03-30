package maths

import (
	//"github.com/chilledornaments/learn-go-with-tests/maths/clockface"
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	testCases := []struct {
		tm    time.Time
		angle float64
	}{
		// 1 radian = 1 pi
		// there 2 radians = the full circle
		// 1 radian = 30 seconds = half the clock
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 00), 0},
		// 1 radian = 30 seconds, 1.5 radians = 45 seconds
		{simpleTime(0, 0, 45), (math.Pi) * 1.5},
		// (pi / 30 = 1 second) * 7 = 7 seconds
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}

	for _, tc := range testCases {
		t.Run(testName(tc.tm), func(t *testing.T) {
			got := secondsInRadians(tc.tm)
			if got != tc.angle {
				t.Errorf("wanted %v radians, got %v", tc.angle, got)
			}
		})
	}
}

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := secondHandPoint(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Fatalf("Wanted %v Point, but got %v", c.point, got)
			}
		})
	}
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) &&
		roughlyEqualFloat64(a.Y, b.Y)
}

func simpleTime(i int, i2 int, i3 int) time.Time {
	return time.Date(1000, time.January, 1, i, i2, i3, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
