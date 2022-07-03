// <auto-generated>
// This code was generated by the UnitCodeGenerator tool
//
// Changes to this file will be lost if the code is regenerated
// </auto-generated>

package area

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
func TestConvertKnownAcresToSquareKilometres(t * testing.T) {
    if !cmp.Equal(Acres.ToSquareKilometres(100.0), 0.404686, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.404686, Acres.ToSquareKilometres(100.0));
    }
    if !cmp.Equal(Acres.ToSquareKilometres(90.0), 0.364217, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.364217, Acres.ToSquareKilometres(90.0));
    }
    if !cmp.Equal(Acres.ToSquareKilometres(1800.0), 7.284342, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 7.284342, Acres.ToSquareKilometres(1800.0));
    }
}

func TestConvertKnownAcresToSquareMetres(t * testing.T) {
    if !cmp.Equal(Acres.ToSquareMetres(1.2), 4856.23, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 4856.23, Acres.ToSquareMetres(1.2));
    }
    if !cmp.Equal(Acres.ToSquareMetres(0.8), 3237.49, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 3237.49, Acres.ToSquareMetres(0.8));
    }
    if !cmp.Equal(Acres.ToSquareMetres(5.6), 22662.416, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 22662.416, Acres.ToSquareMetres(5.6));
    }
}

func TestConvertKnownAcresToSquareMiles(t * testing.T) {
    if !cmp.Equal(Acres.ToSquareMiles(1009.0), 1.576562, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1.576562, Acres.ToSquareMiles(1009.0));
    }
    if !cmp.Equal(Acres.ToSquareMiles(90.0), 0.140625, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.140625, Acres.ToSquareMiles(90.0));
    }
    if !cmp.Equal(Acres.ToSquareMiles(765.0), 1.19531, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1.19531, Acres.ToSquareMiles(765.0));
    }
}

func TestConvertKnownAcresToSquareYards(t * testing.T) {
    if !cmp.Equal(Acres.ToSquareYards(3.4), 16456.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 16456.0, Acres.ToSquareYards(3.4));
    }
    if !cmp.Equal(Acres.ToSquareYards(0.7), 3388.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 3388.0, Acres.ToSquareYards(0.7));
    }
    if !cmp.Equal(Acres.ToSquareYards(0.01), 48.4, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 48.4, Acres.ToSquareYards(0.01));
    }
}

func TestConvertKnownAcresToSquareFeet(t * testing.T) {
    if !cmp.Equal(Acres.ToSquareFeet(0.01), 435.6, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 435.6, Acres.ToSquareFeet(0.01));
    }
    if !cmp.Equal(Acres.ToSquareFeet(12.0), 522720.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 522720.0, Acres.ToSquareFeet(12.0));
    }
    if !cmp.Equal(Acres.ToSquareFeet(0.67), 29185.2, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 29185.2, Acres.ToSquareFeet(0.67));
    }
}

func TestConvertKnownAcresToSquareInches(t * testing.T) {
    if !cmp.Equal(Acres.ToSquareInches(0.09), 564537.6, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 564537.6, Acres.ToSquareInches(0.09));
    }
    if !cmp.Equal(Acres.ToSquareInches(0.005), 31363.2, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 31363.2, Acres.ToSquareInches(0.005));
    }
    if !cmp.Equal(Acres.ToSquareInches(0.012), 75271.68, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 75271.68, Acres.ToSquareInches(0.012));
    }
}

func TestConvertKnownAcresToHectares(t * testing.T) {
    if !cmp.Equal(Acres.ToHectares(1109.0), 448.7964, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 448.7964, Acres.ToHectares(1109.0));
    }
    if !cmp.Equal(Acres.ToHectares(5.6), 2.26624, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2.26624, Acres.ToHectares(5.6));
    }
    if !cmp.Equal(Acres.ToHectares(1.23), 0.4977633, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.4977633, Acres.ToHectares(1.23));
    }
}

func TestConvertKnownHectaresToSquareKilometres(t * testing.T) {
    if !cmp.Equal(Hectares.ToSquareKilometres(15000.0), 150.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 150.0, Hectares.ToSquareKilometres(15000.0));
    }
    if !cmp.Equal(Hectares.ToSquareKilometres(300.0), 3.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 3.0, Hectares.ToSquareKilometres(300.0));
    }
    if !cmp.Equal(Hectares.ToSquareKilometres(45.6), 0.456, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.456, Hectares.ToSquareKilometres(45.6));
    }
}

func TestConvertKnownHectaresToSquareMetres(t * testing.T) {
    if !cmp.Equal(Hectares.ToSquareMetres(1.4), 14000.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 14000.0, Hectares.ToSquareMetres(1.4));
    }
    if !cmp.Equal(Hectares.ToSquareMetres(0.9), 9000.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 9000.0, Hectares.ToSquareMetres(0.9));
    }
    if !cmp.Equal(Hectares.ToSquareMetres(0.012), 120.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 120.0, Hectares.ToSquareMetres(0.012));
    }
}

func TestConvertKnownHectaresToSquareMiles(t * testing.T) {
    if !cmp.Equal(Hectares.ToSquareMiles(1102.0), 4.254846, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 4.254846, Hectares.ToSquareMiles(1102.0));
    }
    if !cmp.Equal(Hectares.ToSquareMiles(4500.0), 17.3746, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 17.3746, Hectares.ToSquareMiles(4500.0));
    }
    if !cmp.Equal(Hectares.ToSquareMiles(90.0), 0.347492, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.347492, Hectares.ToSquareMiles(90.0));
    }
}

func TestConvertKnownHectaresToSquareYards(t * testing.T) {
    if !cmp.Equal(Hectares.ToSquareYards(1.2), 14351.8805556, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 14351.8805556, Hectares.ToSquareYards(1.2));
    }
    if !cmp.Equal(Hectares.ToSquareYards(0.8), 9567.92, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 9567.92, Hectares.ToSquareYards(0.8));
    }
    if !cmp.Equal(Hectares.ToSquareYards(34.0), 406636.615, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 406636.615, Hectares.ToSquareYards(34.0));
    }
}

func TestConvertKnownHectaresToSquareFeet(t * testing.T) {
    if !cmp.Equal(Hectares.ToSquareFeet(3.4), 365972.599, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 365972.599, Hectares.ToSquareFeet(3.4));
    }
    if !cmp.Equal(Hectares.ToSquareFeet(0.09), 9687.519, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 9687.519, Hectares.ToSquareFeet(0.09));
    }
    if !cmp.Equal(Hectares.ToSquareFeet(1.2), 129166.7999, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 129166.7999, Hectares.ToSquareFeet(1.2));
    }
}

func TestConvertKnownHectaresToSquareInches(t * testing.T) {
    if !cmp.Equal(Hectares.ToSquareInches(0.009), 139500.28, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 139500.28, Hectares.ToSquareInches(0.009));
    }
    if !cmp.Equal(Hectares.ToSquareInches(0.01), 155000.31, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 155000.31, Hectares.ToSquareInches(0.01));
    }
    if !cmp.Equal(Hectares.ToSquareInches(0.0061), 94550.1891, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 94550.1891, Hectares.ToSquareInches(0.0061));
    }
}

func TestConvertKnownHectaresToAcres(t * testing.T) {
    if !cmp.Equal(Hectares.ToAcres(1.2), 2.96526, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2.96526, Hectares.ToAcres(1.2));
    }
    if !cmp.Equal(Hectares.ToAcres(0.8), 1.97684, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1.97684, Hectares.ToAcres(0.8));
    }
    if !cmp.Equal(Hectares.ToAcres(4.2), 10.3784, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 10.3784, Hectares.ToAcres(4.2));
    }
}

func TestConvertKnownSquareFeetToSquareKilometres(t * testing.T) {
    if !cmp.Equal(SquareFeet.ToSquareKilometres(88997766.0), 8.2681630146, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 8.2681630146, SquareFeet.ToSquareKilometres(88997766.0));
    }
    if !cmp.Equal(SquareFeet.ToSquareKilometres(10009002.0), 0.92986671317, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.92986671317, SquareFeet.ToSquareKilometres(10009002.0));
    }
    if !cmp.Equal(SquareFeet.ToSquareKilometres(987654321.0), 91.75608889, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 91.75608889, SquareFeet.ToSquareKilometres(987654321.0));
    }
}

func TestConvertKnownSquareFeetToSquareMetres(t * testing.T) {
    if !cmp.Equal(SquareFeet.ToSquareMetres(900.0), 83.6127, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 83.6127, SquareFeet.ToSquareMetres(900.0));
    }
    if !cmp.Equal(SquareFeet.ToSquareMetres(799.123), 74.24095603, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 74.24095603, SquareFeet.ToSquareMetres(799.123));
    }
    if !cmp.Equal(SquareFeet.ToSquareMetres(500.0), 46.4515, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 46.4515, SquareFeet.ToSquareMetres(500.0));
    }
}

func TestConvertKnownSquareFeetToSquareMiles(t * testing.T) {
    if !cmp.Equal(SquareFeet.ToSquareMiles(12000000.0), 0.43044077135, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.43044077135, SquareFeet.ToSquareMiles(12000000.0));
    }
    if !cmp.Equal(SquareFeet.ToSquareMiles(987654321.0), 35.4272239799, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 35.4272239799, SquareFeet.ToSquareMiles(987654321.0));
    }
    if !cmp.Equal(SquareFeet.ToSquareMiles(900800700.0), 32.3117790117, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 32.3117790117, SquareFeet.ToSquareMiles(900800700.0));
    }
}

func TestConvertKnownSquareFeetToSquareYards(t * testing.T) {
    if !cmp.Equal(SquareFeet.ToSquareYards(909.0), 101.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 101.0, SquareFeet.ToSquareYards(909.0));
    }
    if !cmp.Equal(SquareFeet.ToSquareYards(123.456), 13.71733333, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 13.71733333, SquareFeet.ToSquareYards(123.456));
    }
    if !cmp.Equal(SquareFeet.ToSquareYards(8009.0), 889.8889, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 889.8889, SquareFeet.ToSquareYards(8009.0));
    }
}

func TestConvertKnownSquareFeetToSquareInches(t * testing.T) {
    if !cmp.Equal(SquareFeet.ToSquareInches(6.7), 964.8, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 964.8, SquareFeet.ToSquareInches(6.7));
    }
    if !cmp.Equal(SquareFeet.ToSquareInches(1.5), 216.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 216.0, SquareFeet.ToSquareInches(1.5));
    }
    if !cmp.Equal(SquareFeet.ToSquareInches(0.9), 129.6, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 129.6, SquareFeet.ToSquareInches(0.9));
    }
}

func TestConvertKnownSquareFeetToHectares(t * testing.T) {
    if !cmp.Equal(SquareFeet.ToHectares(90000.0), 0.83612736, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.83612736, SquareFeet.ToHectares(90000.0));
    }
    if !cmp.Equal(SquareFeet.ToHectares(120120.0), 1.11595132, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1.11595132, SquareFeet.ToHectares(120120.0));
    }
    if !cmp.Equal(SquareFeet.ToHectares(90071.0), 0.83678697, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.83678697, SquareFeet.ToHectares(90071.0));
    }
}

func TestConvertKnownSquareFeetToAcres(t * testing.T) {
    if !cmp.Equal(SquareFeet.ToAcres(90000.0), 2.0661157, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2.0661157, SquareFeet.ToAcres(90000.0));
    }
    if !cmp.Equal(SquareFeet.ToAcres(123456.0), 2.83415978, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2.83415978, SquareFeet.ToAcres(123456.0));
    }
    if !cmp.Equal(SquareFeet.ToAcres(8809.0), 0.2022268, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.2022268, SquareFeet.ToAcres(8809.0));
    }
}

func TestConvertKnownSquareInchesToSquareKilometres(t * testing.T) {
    if !cmp.Equal(SquareInches.ToSquareKilometres(678900000.0), 0.437999124, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.437999124, SquareInches.ToSquareKilometres(678900000.0));
    }
    if !cmp.Equal(SquareInches.ToSquareKilometres(10000000000.0), 6.4516, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 6.4516, SquareInches.ToSquareKilometres(10000000000.0));
    }
    if !cmp.Equal(SquareInches.ToSquareKilometres(9e12), 5806.44, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 5806.44, SquareInches.ToSquareKilometres(9e12));
    }
}

func TestConvertKnownSquareInchesToSquareMetres(t * testing.T) {
    if !cmp.Equal(SquareInches.ToSquareMetres(1090.0), 0.7032244, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.7032244, SquareInches.ToSquareMetres(1090.0));
    }
    if !cmp.Equal(SquareInches.ToSquareMetres(1.3e6), 838.708, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 838.708, SquareInches.ToSquareMetres(1.3e6));
    }
    if !cmp.Equal(SquareInches.ToSquareMetres(9988.0), 6.443858, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 6.443858, SquareInches.ToSquareMetres(9988.0));
    }
}

func TestConvertKnownSquareInchesToSquareMiles(t * testing.T) {
    if !cmp.Equal(SquareInches.ToSquareMiles(1007008000.0), 0.2508433450668, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.2508433450668, SquareInches.ToSquareMiles(1007008000.0));
    }
    if !cmp.Equal(SquareInches.ToSquareMiles(1.2e12), 298.9172023262932, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 298.9172023262932, SquareInches.ToSquareMiles(1.2e12));
    }
    if !cmp.Equal(SquareInches.ToSquareMiles(250e7), 0.6227441715131, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.6227441715131, SquareInches.ToSquareMiles(250e7));
    }
}

func TestConvertKnownSquareInchesToSquareYards(t * testing.T) {
    if !cmp.Equal(SquareInches.ToSquareYards(900.0), 0.694444, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.694444, SquareInches.ToSquareYards(900.0));
    }
    if !cmp.Equal(SquareInches.ToSquareYards(15000.0), 11.574074, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 11.574074, SquareInches.ToSquareYards(15000.0));
    }
    if !cmp.Equal(SquareInches.ToSquareYards(3e6), 2314.814815, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2314.814815, SquareInches.ToSquareYards(3e6));
    }
}

func TestConvertKnownSquareInchesToSquareFeet(t * testing.T) {
    if !cmp.Equal(SquareInches.ToSquareFeet(34.0), 0.236111, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.236111, SquareInches.ToSquareFeet(34.0));
    }
    if !cmp.Equal(SquareInches.ToSquareFeet(1002.0), 6.958333, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 6.958333, SquareInches.ToSquareFeet(1002.0));
    }
    if !cmp.Equal(SquareInches.ToSquareFeet(890.0), 6.18056, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 6.18056, SquareInches.ToSquareFeet(890.0));
    }
}

func TestConvertKnownSquareInchesToHectares(t * testing.T) {
    if !cmp.Equal(SquareInches.ToHectares(9000000.0), 0.580644, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.580644, SquareInches.ToHectares(9000000.0));
    }
    if !cmp.Equal(SquareInches.ToHectares(12345678.0), 0.79649376185, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.79649376185, SquareInches.ToHectares(12345678.0));
    }
    if !cmp.Equal(SquareInches.ToHectares(99887766.0), 6.4443591113, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 6.4443591113, SquareInches.ToHectares(99887766.0));
    }
}

func TestConvertKnownSquareInchesToAcres(t * testing.T) {
    if !cmp.Equal(SquareInches.ToAcres(900100.0), 0.143496199, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.143496199, SquareInches.ToAcres(900100.0));
    }
    if !cmp.Equal(SquareInches.ToAcres(5e6), 0.7971125395, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.7971125395, SquareInches.ToAcres(5e6));
    }
    if !cmp.Equal(SquareInches.ToAcres(12345678.0), 1.9681789486, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1.9681789486, SquareInches.ToAcres(12345678.0));
    }
}

func TestConvertKnownSquareKilometresToSquareMetres(t * testing.T) {
    if !cmp.Equal(SquareKilometres.ToSquareMetres(0.12), 120000.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 120000.0, SquareKilometres.ToSquareMetres(0.12));
    }
    if !cmp.Equal(SquareKilometres.ToSquareMetres(0.9), 900000.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 900000.0, SquareKilometres.ToSquareMetres(0.9));
    }
    if !cmp.Equal(SquareKilometres.ToSquareMetres(8.123456), 8123456.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 8123456.0, SquareKilometres.ToSquareMetres(8.123456));
    }
}

func TestConvertKnownSquareKilometresToSquareMiles(t * testing.T) {
    if !cmp.Equal(SquareKilometres.ToSquareMiles(492.098), 190.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 190.0, SquareKilometres.ToSquareMiles(492.098));
    }
    if !cmp.Equal(SquareKilometres.ToSquareMiles(23.3099), 9.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 9.0, SquareKilometres.ToSquareMiles(23.3099));
    }
    if !cmp.Equal(SquareKilometres.ToSquareMiles(4661.979), 1800.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1800.0, SquareKilometres.ToSquareMiles(4661.979));
    }
}

func TestConvertKnownSquareKilometresToSquareYards(t * testing.T) {
    if !cmp.Equal(SquareKilometres.ToSquareYards(0.158028071), 189000.0005, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 189000.0005, SquareKilometres.ToSquareYards(0.158028071));
    }
    if !cmp.Equal(SquareKilometres.ToSquareYards(8.361272764), 9999999.0310, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 9999999.0310, SquareKilometres.ToSquareYards(8.361272764));
    }
    if !cmp.Equal(SquareKilometres.ToSquareYards(0.6), 717594.030, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 717594.030, SquareKilometres.ToSquareYards(0.6));
    }
}

func TestConvertKnownSquareKilometresToSquareFeet(t * testing.T) {
    if !cmp.Equal(SquareKilometres.ToSquareFeet(0.009), 96875.194, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 96875.194, SquareKilometres.ToSquareFeet(0.009));
    }
    if !cmp.Equal(SquareKilometres.ToSquareFeet(0.08), 861112.833, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 861112.833, SquareKilometres.ToSquareFeet(0.08));
    }
    if !cmp.Equal(SquareKilometres.ToSquareFeet(0.123), 1323960.9812553, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1323960.9812553, SquareKilometres.ToSquareFeet(0.123));
    }
}

func TestConvertKnownSquareKilometresToSquareInches(t * testing.T) {
    if !cmp.Equal(SquareKilometres.ToSquareInches(0.0008), 1240002.48, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1240002.48, SquareKilometres.ToSquareInches(0.0008));
    }
    if !cmp.Equal(SquareKilometres.ToSquareInches(0.00123), 1906503.813, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1906503.813, SquareKilometres.ToSquareInches(0.00123));
    }
    if !cmp.Equal(SquareKilometres.ToSquareInches(0.000045), 69750.1395, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 69750.1395, SquareKilometres.ToSquareInches(0.000045));
    }
}

func TestConvertKnownSquareKilometresToHectares(t * testing.T) {
    if !cmp.Equal(SquareKilometres.ToHectares(190.0), 19000.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 19000.0, SquareKilometres.ToHectares(190.0));
    }
    if !cmp.Equal(SquareKilometres.ToHectares(55.67), 5567.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 5567.0, SquareKilometres.ToHectares(55.67));
    }
    if !cmp.Equal(SquareKilometres.ToHectares(0.9), 90.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 90.0, SquareKilometres.ToHectares(0.9));
    }
}

func TestConvertKnownSquareKilometresToAcres(t * testing.T) {
    if !cmp.Equal(SquareKilometres.ToAcres(30.393962), 7510.4999900100, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 7510.4999900100, SquareKilometres.ToAcres(30.393962));
    }
    if !cmp.Equal(SquareKilometres.ToAcres(54.1), 13368.3805, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 13368.3805, SquareKilometres.ToAcres(54.1));
    }
    if !cmp.Equal(SquareKilometres.ToAcres(90.67), 22405.0103, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 22405.0103, SquareKilometres.ToAcres(90.67));
    }
}

func TestConvertKnownSquareMetresToSquareKilometres(t * testing.T) {
    if !cmp.Equal(SquareMetres.ToSquareKilometres(19000.0), 0.019, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.019, SquareMetres.ToSquareKilometres(19000.0));
    }
    if !cmp.Equal(SquareMetres.ToSquareKilometres(123456.0), 0.123456, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.123456, SquareMetres.ToSquareKilometres(123456.0));
    }
    if !cmp.Equal(SquareMetres.ToSquareKilometres(900100.0), 0.9001, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.9001, SquareMetres.ToSquareKilometres(900100.0));
    }
}

func TestConvertKnownSquareMetresToSquareMiles(t * testing.T) {
    if !cmp.Equal(SquareMetres.ToSquareMiles(190009.0), 0.073362885, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.073362885, SquareMetres.ToSquareMiles(190009.0));
    }
    if !cmp.Equal(SquareMetres.ToSquareMiles(12345678.0), 4.7666929245, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 4.7666929245, SquareMetres.ToSquareMiles(12345678.0));
    }
    if !cmp.Equal(SquareMetres.ToSquareMiles(777666.0), 0.300258521, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.300258521, SquareMetres.ToSquareMiles(777666.0));
    }
}

func TestConvertKnownSquareMetresToSquareYards(t * testing.T) {
    if !cmp.Equal(SquareMetres.ToSquareYards(5.0), 5.97995, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 5.97995, SquareMetres.ToSquareYards(5.0));
    }
    if !cmp.Equal(SquareMetres.ToSquareYards(1.23), 1.471068, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1.471068, SquareMetres.ToSquareYards(1.23));
    }
    if !cmp.Equal(SquareMetres.ToSquareYards(700.0), 837.193, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 837.193, SquareMetres.ToSquareYards(700.0));
    }
}

func TestConvertKnownSquareMetresToSquareFeet(t * testing.T) {
    if !cmp.Equal(SquareMetres.ToSquareFeet(90.0), 968.7519375, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 968.7519375, SquareMetres.ToSquareFeet(90.0));
    }
    if !cmp.Equal(SquareMetres.ToSquareFeet(180.0), 1937.503875, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1937.503875, SquareMetres.ToSquareFeet(180.0));
    }
    if !cmp.Equal(SquareMetres.ToSquareFeet(123.4), 1328.2665454, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1328.2665454, SquareMetres.ToSquareFeet(123.4));
    }
}

func TestConvertKnownSquareMetresToSquareInches(t * testing.T) {
    if !cmp.Equal(SquareMetres.ToSquareInches(8.0), 12400.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 12400.0, SquareMetres.ToSquareInches(8.0));
    }
    if !cmp.Equal(SquareMetres.ToSquareInches(1.23), 1906.504, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1906.504, SquareMetres.ToSquareInches(1.23));
    }
    if !cmp.Equal(SquareMetres.ToSquareInches(0.9), 1395.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1395.0, SquareMetres.ToSquareInches(0.9));
    }
}

func TestConvertKnownSquareMetresToHectares(t * testing.T) {
    if !cmp.Equal(SquareMetres.ToHectares(1234.0), 0.1234, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.1234, SquareMetres.ToHectares(1234.0));
    }
    if !cmp.Equal(SquareMetres.ToHectares(560.9), 0.05609, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.05609, SquareMetres.ToHectares(560.9));
    }
    if !cmp.Equal(SquareMetres.ToHectares(100900.0), 10.09, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 10.09, SquareMetres.ToHectares(100900.0));
    }
}

func TestConvertKnownSquareMetresToAcres(t * testing.T) {
    if !cmp.Equal(SquareMetres.ToAcres(986.0), 0.243646, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.243646, SquareMetres.ToAcres(986.0));
    }
    if !cmp.Equal(SquareMetres.ToAcres(1020.0), 0.2520475, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.2520475, SquareMetres.ToAcres(1020.0));
    }
    if !cmp.Equal(SquareMetres.ToAcres(666111.0), 164.599613, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 164.599613, SquareMetres.ToAcres(666111.0));
    }
}

func TestConvertKnownSquareMilesToSquareKilometres(t * testing.T) {
    if !cmp.Equal(SquareMiles.ToSquareKilometres(3.0), 7.76996, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 7.76996, SquareMiles.ToSquareKilometres(3.0));
    }
    if !cmp.Equal(SquareMiles.ToSquareKilometres(0.9), 2.33099, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2.33099, SquareMiles.ToSquareKilometres(0.9));
    }
    if !cmp.Equal(SquareMiles.ToSquareKilometres(100.0), 258.999, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 258.999, SquareMiles.ToSquareKilometres(100.0));
    }
}

func TestConvertKnownSquareMilesToSquareMetres(t * testing.T) {
    if !cmp.Equal(SquareMiles.ToSquareMetres(0.009), 23309.893, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 23309.893, SquareMiles.ToSquareMetres(0.009));
    }
    if !cmp.Equal(SquareMiles.ToSquareMetres(0.010), 25899.88, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 25899.88, SquareMiles.ToSquareMetres(0.010));
    }
    if !cmp.Equal(SquareMiles.ToSquareMetres(0.0006), 1553.99287, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1553.99287, SquareMiles.ToSquareMetres(0.0006));
    }
}

func TestConvertKnownSquareMilesToSquareYards(t * testing.T) {
    if !cmp.Equal(SquareMiles.ToSquareYards(0.01), 30976.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 30976.0, SquareMiles.ToSquareYards(0.01));
    }
    if !cmp.Equal(SquareMiles.ToSquareYards(0.00123), 3810.048, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 3810.048, SquareMiles.ToSquareYards(0.00123));
    }
    if !cmp.Equal(SquareMiles.ToSquareYards(0.09), 278784.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 278784.0, SquareMiles.ToSquareYards(0.09));
    }
}

func TestConvertKnownSquareMilesToSquareFeet(t * testing.T) {
    if !cmp.Equal(SquareMiles.ToSquareFeet(0.01), 278784.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 278784.0, SquareMiles.ToSquareFeet(0.01));
    }
    if !cmp.Equal(SquareMiles.ToSquareFeet(0.005), 139392.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 139392.0, SquareMiles.ToSquareFeet(0.005));
    }
    if !cmp.Equal(SquareMiles.ToSquareFeet(0.08), 2230272.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2230272.0, SquareMiles.ToSquareFeet(0.08));
    }
}

func TestConvertKnownSquareMilesToSquareInches(t * testing.T) {
    if !cmp.Equal(SquareMiles.ToSquareInches(0.0001), 401448.96, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 401448.96, SquareMiles.ToSquareInches(0.0001));
    }
    if !cmp.Equal(SquareMiles.ToSquareInches(0.00098), 3934199.808, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 3934199.808, SquareMiles.ToSquareInches(0.00098));
    }
    if !cmp.Equal(SquareMiles.ToSquareInches(0.000007), 28101.4272, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 28101.4272, SquareMiles.ToSquareInches(0.000007));
    }
}

func TestConvertKnownSquareMilesToHectares(t * testing.T) {
    if !cmp.Equal(SquareMiles.ToHectares(100.0), 25899.881103, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 25899.881103, SquareMiles.ToHectares(100.0));
    }
    if !cmp.Equal(SquareMiles.ToHectares(2.3), 595.697, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 595.697, SquareMiles.ToHectares(2.3));
    }
    if !cmp.Equal(SquareMiles.ToHectares(0.9), 233.099, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 233.099, SquareMiles.ToHectares(0.9));
    }
}

func TestConvertKnownSquareMilesToAcres(t * testing.T) {
    if !cmp.Equal(SquareMiles.ToAcres(2.0), 1280.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1280.0, SquareMiles.ToAcres(2.0));
    }
    if !cmp.Equal(SquareMiles.ToAcres(0.1), 64.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 64.0, SquareMiles.ToAcres(0.1));
    }
    if !cmp.Equal(SquareMiles.ToAcres(4.6), 2944.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2944.0, SquareMiles.ToAcres(4.6));
    }
}

func TestConvertKnownSquareYardsToSquareKilometres(t * testing.T) {
    if !cmp.Equal(SquareYards.ToSquareKilometres(900000.0), 0.752514624, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.752514624, SquareYards.ToSquareKilometres(900000.0));
    }
    if !cmp.Equal(SquareYards.ToSquareKilometres(190190.0), 0.159023063, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.159023063, SquareYards.ToSquareKilometres(190190.0));
    }
    if !cmp.Equal(SquareYards.ToSquareKilometres(7000000.0), 5.85289152, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 5.85289152, SquareYards.ToSquareKilometres(7000000.0));
    }
}

func TestConvertKnownSquareYardsToSquareMetres(t * testing.T) {
    if !cmp.Equal(SquareYards.ToSquareMetres(700.0), 585.289, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 585.289, SquareYards.ToSquareMetres(700.0));
    }
    if !cmp.Equal(SquareYards.ToSquareMetres(12.0), 10.0335, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 10.0335, SquareYards.ToSquareMetres(12.0));
    }
    if !cmp.Equal(SquareYards.ToSquareMetres(9.1), 7.60876, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 7.60876, SquareYards.ToSquareMetres(9.1));
    }
}

func TestConvertKnownSquareYardsToSquareMiles(t * testing.T) {
    if !cmp.Equal(SquareYards.ToSquareMiles(98700.0), 0.031863378, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.031863378, SquareYards.ToSquareMiles(98700.0));
    }
    if !cmp.Equal(SquareYards.ToSquareMiles(8888888.0), 2.869604855, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2.869604855, SquareYards.ToSquareMiles(8888888.0));
    }
    if !cmp.Equal(SquareYards.ToSquareMiles(100200300.0), 32.3477208161, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 32.3477208161, SquareYards.ToSquareMiles(100200300.0));
    }
}

func TestConvertKnownSquareYardsToSquareFeet(t * testing.T) {
    if !cmp.Equal(SquareYards.ToSquareFeet(12.0), 108.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 108.0, SquareYards.ToSquareFeet(12.0));
    }
    if !cmp.Equal(SquareYards.ToSquareFeet(5.6), 50.4, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 50.4, SquareYards.ToSquareFeet(5.6));
    }
    if !cmp.Equal(SquareYards.ToSquareFeet(102.5), 922.5, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 922.5, SquareYards.ToSquareFeet(102.5));
    }
}

func TestConvertKnownSquareYardsToSquareInches(t * testing.T) {
    if !cmp.Equal(SquareYards.ToSquareInches(56.7), 73483.2, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 73483.2, SquareYards.ToSquareInches(56.7));
    }
    if !cmp.Equal(SquareYards.ToSquareInches(1.8), 2332.8, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2332.8, SquareYards.ToSquareInches(1.8));
    }
    if !cmp.Equal(SquareYards.ToSquareInches(0.2), 259.2, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 259.2, SquareYards.ToSquareInches(0.2));
    }
}

func TestConvertKnownSquareYardsToHectares(t * testing.T) {
    if !cmp.Equal(SquareYards.ToHectares(10090.0), 0.84365251, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.84365251, SquareYards.ToHectares(10090.0));
    }
    if !cmp.Equal(SquareYards.ToHectares(98765.0), 8.2580119, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 8.2580119, SquareYards.ToHectares(98765.0));
    }
    if !cmp.Equal(SquareYards.ToHectares(3090.09), 0.2583634, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.2583634, SquareYards.ToHectares(3090.09));
    }
}

func TestConvertKnownSquareYardsToAcres(t * testing.T) {
    if !cmp.Equal(SquareYards.ToAcres(1000.0), 0.206612, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.206612, SquareYards.ToAcres(1000.0));
    }
    if !cmp.Equal(SquareYards.ToAcres(899.0), 0.185744, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.185744, SquareYards.ToAcres(899.0));
    }
    if !cmp.Equal(SquareYards.ToAcres(5678.0), 1.17314, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1.17314, SquareYards.ToAcres(5678.0));
    }
}

