// <auto-generated>
// This code was generated by the UnitCodeGenerator tool
//
// Changes to this file will be lost if the code is regenerated
// </auto-generated>

package speed

import (
	"testing"
	"math"
	"github.com/google/go-cmp/cmp"
)

func withinTolerance() cmp.Option {
	return cmp.Comparer(func(x, y float64) bool {
		delta := math.Abs(x - y)
		mean := math.Abs(x + y) / 2.0
		return delta / mean < 0.01
	})
}
func TestConvertKnownFeetPerSecondToMilesPerHour(t * testing.T) {
    if !cmp.Equal(FeetPerSecond.ToMilesPerHour(14.5), 9.886364, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 9.886364, FeetPerSecond.ToMilesPerHour(14.5));
    }
    if !cmp.Equal(FeetPerSecond.ToMilesPerHour(100.0), 68.1818, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 68.1818, FeetPerSecond.ToMilesPerHour(100.0));
    }
    if !cmp.Equal(FeetPerSecond.ToMilesPerHour(1.6), 1.09091, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1.09091, FeetPerSecond.ToMilesPerHour(1.6));
    }
}

func TestConvertKnownFeetPerSecondToKilometresPerHour(t * testing.T) {
    if !cmp.Equal(FeetPerSecond.ToKilometresPerHour(1.5), 1.64592, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1.64592, FeetPerSecond.ToKilometresPerHour(1.5));
    }
    if !cmp.Equal(FeetPerSecond.ToKilometresPerHour(67.9), 74.50531, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 74.50531, FeetPerSecond.ToKilometresPerHour(67.9));
    }
    if !cmp.Equal(FeetPerSecond.ToKilometresPerHour(109.0), 119.604, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 119.604, FeetPerSecond.ToKilometresPerHour(109.0));
    }
}

func TestConvertKnownFeetPerSecondToMetresPerSecond(t * testing.T) {
    if !cmp.Equal(FeetPerSecond.ToMetresPerSecond(10.8), 3.29184, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 3.29184, FeetPerSecond.ToMetresPerSecond(10.8));
    }
    if !cmp.Equal(FeetPerSecond.ToMetresPerSecond(66.0), 20.1168, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 20.1168, FeetPerSecond.ToMetresPerSecond(66.0));
    }
    if !cmp.Equal(FeetPerSecond.ToMetresPerSecond(2.3), 0.70104, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.70104, FeetPerSecond.ToMetresPerSecond(2.3));
    }
}

func TestConvertKnownFeetPerSecondToKnots(t * testing.T) {
    if !cmp.Equal(FeetPerSecond.ToKnots(2.3), 1.36271, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1.36271, FeetPerSecond.ToKnots(2.3));
    }
    if !cmp.Equal(FeetPerSecond.ToKnots(666.0), 394.594, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 394.594, FeetPerSecond.ToKnots(666.0));
    }
    if !cmp.Equal(FeetPerSecond.ToKnots(17.1), 10.13147, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 10.13147, FeetPerSecond.ToKnots(17.1));
    }
}

func TestConvertKnownKilometresPerHourToMilesPerHour(t * testing.T) {
    if !cmp.Equal(KilometresPerHour.ToMilesPerHour(67.0), 41.6319, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 41.6319, KilometresPerHour.ToMilesPerHour(67.0));
    }
    if !cmp.Equal(KilometresPerHour.ToMilesPerHour(12.0), 7.45645, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 7.45645, KilometresPerHour.ToMilesPerHour(12.0));
    }
    if !cmp.Equal(KilometresPerHour.ToMilesPerHour(6.3), 3.91464, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 3.91464, KilometresPerHour.ToMilesPerHour(6.3));
    }
}

func TestConvertKnownKilometresPerHourToFeetPerSecond(t * testing.T) {
    if !cmp.Equal(KilometresPerHour.ToFeetPerSecond(5.0), 4.55672, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 4.55672, KilometresPerHour.ToFeetPerSecond(5.0));
    }
    if !cmp.Equal(KilometresPerHour.ToFeetPerSecond(1.5), 1.36702, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1.36702, KilometresPerHour.ToFeetPerSecond(1.5));
    }
    if !cmp.Equal(KilometresPerHour.ToFeetPerSecond(89.8), 81.83873, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 81.83873, KilometresPerHour.ToFeetPerSecond(89.8));
    }
}

func TestConvertKnownKilometresPerHourToMetresPerSecond(t * testing.T) {
    if !cmp.Equal(KilometresPerHour.ToMetresPerSecond(67.0), 18.6111, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 18.6111, KilometresPerHour.ToMetresPerSecond(67.0));
    }
    if !cmp.Equal(KilometresPerHour.ToMetresPerSecond(5.9), 1.63889, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1.63889, KilometresPerHour.ToMetresPerSecond(5.9));
    }
    if !cmp.Equal(KilometresPerHour.ToMetresPerSecond(900.0), 250.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 250.0, KilometresPerHour.ToMetresPerSecond(900.0));
    }
}

func TestConvertKnownKilometresPerHourToKnots(t * testing.T) {
    if !cmp.Equal(KilometresPerHour.ToKnots(900.0), 485.961, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 485.961, KilometresPerHour.ToKnots(900.0));
    }
    if !cmp.Equal(KilometresPerHour.ToKnots(3.9), 2.10583, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2.10583, KilometresPerHour.ToKnots(3.9));
    }
    if !cmp.Equal(KilometresPerHour.ToKnots(12.0), 6.47948, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 6.47948, KilometresPerHour.ToKnots(12.0));
    }
}

func TestConvertKnownKnotsToMilesPerHour(t * testing.T) {
    if !cmp.Equal(Knots.ToMilesPerHour(8.0), 9.20624, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 9.20624, Knots.ToMilesPerHour(8.0));
    }
    if !cmp.Equal(Knots.ToMilesPerHour(1.2), 1.38094, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1.38094, Knots.ToMilesPerHour(1.2));
    }
    if !cmp.Equal(Knots.ToMilesPerHour(670.0), 771.022, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 771.022, Knots.ToMilesPerHour(670.0));
    }
}

func TestConvertKnownKnotsToKilometresPerHour(t * testing.T) {
    if !cmp.Equal(Knots.ToKilometresPerHour(678.0), 1255.66, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1255.66, Knots.ToKilometresPerHour(678.0));
    }
    if !cmp.Equal(Knots.ToKilometresPerHour(1.8), 3.3336, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 3.3336, Knots.ToKilometresPerHour(1.8));
    }
    if !cmp.Equal(Knots.ToKilometresPerHour(56.0), 103.712, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 103.712, Knots.ToKilometresPerHour(56.0));
    }
}

func TestConvertKnownKnotsToFeetPerSecond(t * testing.T) {
    if !cmp.Equal(Knots.ToFeetPerSecond(56.0), 94.5174, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 94.5174, Knots.ToFeetPerSecond(56.0));
    }
    if !cmp.Equal(Knots.ToFeetPerSecond(4.7), 7.93271, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 7.93271, Knots.ToFeetPerSecond(4.7));
    }
    if !cmp.Equal(Knots.ToFeetPerSecond(3.0), 5.06343, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 5.06343, Knots.ToFeetPerSecond(3.0));
    }
}

func TestConvertKnownKnotsToMetresPerSecond(t * testing.T) {
    if !cmp.Equal(Knots.ToMetresPerSecond(5.0), 2.57222, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2.57222, Knots.ToMetresPerSecond(5.0));
    }
    if !cmp.Equal(Knots.ToMetresPerSecond(9.1), 4.68144, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 4.68144, Knots.ToMetresPerSecond(9.1));
    }
    if !cmp.Equal(Knots.ToMetresPerSecond(190.0), 97.7444, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 97.7444, Knots.ToMetresPerSecond(190.0));
    }
}

func TestConvertKnownMetresPerSecondToMilesPerHour(t * testing.T) {
    if !cmp.Equal(MetresPerSecond.ToMilesPerHour(13.0), 29.0802, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 29.0802, MetresPerSecond.ToMilesPerHour(13.0));
    }
    if !cmp.Equal(MetresPerSecond.ToMilesPerHour(6.7), 14.9875, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 14.9875, MetresPerSecond.ToMilesPerHour(6.7));
    }
    if !cmp.Equal(MetresPerSecond.ToMilesPerHour(140.0), 313.171, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 313.171, MetresPerSecond.ToMilesPerHour(140.0));
    }
}

func TestConvertKnownMetresPerSecondToKilometresPerHour(t * testing.T) {
    if !cmp.Equal(MetresPerSecond.ToKilometresPerHour(800.0), 2880.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2880.0, MetresPerSecond.ToKilometresPerHour(800.0));
    }
    if !cmp.Equal(MetresPerSecond.ToKilometresPerHour(7.8), 28.08, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 28.08, MetresPerSecond.ToKilometresPerHour(7.8));
    }
    if !cmp.Equal(MetresPerSecond.ToKilometresPerHour(9000.0), 32400.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 32400.0, MetresPerSecond.ToKilometresPerHour(9000.0));
    }
}

func TestConvertKnownMetresPerSecondToFeetPerSecond(t * testing.T) {
    if !cmp.Equal(MetresPerSecond.ToFeetPerSecond(9000.0), 29527.56, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 29527.56, MetresPerSecond.ToFeetPerSecond(9000.0));
    }
    if !cmp.Equal(MetresPerSecond.ToFeetPerSecond(3.4), 11.1549, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 11.1549, MetresPerSecond.ToFeetPerSecond(3.4));
    }
    if !cmp.Equal(MetresPerSecond.ToFeetPerSecond(12.0), 39.3701, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 39.3701, MetresPerSecond.ToFeetPerSecond(12.0));
    }
}

func TestConvertKnownMetresPerSecondToKnots(t * testing.T) {
    if !cmp.Equal(MetresPerSecond.ToKnots(12.0), 23.3261, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 23.3261, MetresPerSecond.ToKnots(12.0));
    }
    if !cmp.Equal(MetresPerSecond.ToKnots(6.3), 12.2462, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 12.2462, MetresPerSecond.ToKnots(6.3));
    }
    if !cmp.Equal(MetresPerSecond.ToKnots(9.0), 17.4946, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 17.4946, MetresPerSecond.ToKnots(9.0));
    }
}

func TestConvertKnownMilesPerHourToKilometresPerHour(t * testing.T) {
    if !cmp.Equal(MilesPerHour.ToKilometresPerHour(345.0), 555.224, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 555.224, MilesPerHour.ToKilometresPerHour(345.0));
    }
    if !cmp.Equal(MilesPerHour.ToKilometresPerHour(1.6), 2.57495, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2.57495, MilesPerHour.ToKilometresPerHour(1.6));
    }
    if !cmp.Equal(MilesPerHour.ToKilometresPerHour(0.5), 0.804672, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.804672, MilesPerHour.ToKilometresPerHour(0.5));
    }
}

func TestConvertKnownMilesPerHourToFeetPerSecond(t * testing.T) {
    if !cmp.Equal(MilesPerHour.ToFeetPerSecond(0.5), 0.733333, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.733333, MilesPerHour.ToFeetPerSecond(0.5));
    }
    if !cmp.Equal(MilesPerHour.ToFeetPerSecond(123.0), 180.4, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 180.4, MilesPerHour.ToFeetPerSecond(123.0));
    }
    if !cmp.Equal(MilesPerHour.ToFeetPerSecond(4.5), 6.6, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 6.6, MilesPerHour.ToFeetPerSecond(4.5));
    }
}

func TestConvertKnownMilesPerHourToMetresPerSecond(t * testing.T) {
    if !cmp.Equal(MilesPerHour.ToMetresPerSecond(4.5), 2.01168, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2.01168, MilesPerHour.ToMetresPerSecond(4.5));
    }
    if !cmp.Equal(MilesPerHour.ToMetresPerSecond(100.0), 44.704, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 44.704, MilesPerHour.ToMetresPerSecond(100.0));
    }
    if !cmp.Equal(MilesPerHour.ToMetresPerSecond(40.0), 17.8816, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 17.8816, MilesPerHour.ToMetresPerSecond(40.0));
    }
}

func TestConvertKnownMilesPerHourToKnots(t * testing.T) {
    if !cmp.Equal(MilesPerHour.ToKnots(30.0), 26.0693, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 26.0693, MilesPerHour.ToKnots(30.0));
    }
    if !cmp.Equal(MilesPerHour.ToKnots(4.5), 3.91039, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 3.91039, MilesPerHour.ToKnots(4.5));
    }
    if !cmp.Equal(MilesPerHour.ToKnots(55.0), 47.7937, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 47.7937, MilesPerHour.ToKnots(55.0));
    }
}

