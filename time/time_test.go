// <auto-generated>
// This code was generated by the UnitCodeGenerator tool
//
// Changes to this file will be lost if the code is regenerated
// </auto-generated>

package time

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
func TestConvertKnownDaysToMilliseconds(t * testing.T) {
    if !cmp.Equal(Days.ToMilliseconds(0.0009), 77760.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 77760.0, Days.ToMilliseconds(0.0009));
    }
    if !cmp.Equal(Days.ToMilliseconds(0.03), 2592000.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2592000.0, Days.ToMilliseconds(0.03));
    }
    if !cmp.Equal(Days.ToMilliseconds(0.006), 518400.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 518400.0, Days.ToMilliseconds(0.006));
    }
}

func TestConvertKnownDaysToSeconds(t * testing.T) {
    if !cmp.Equal(Days.ToSeconds(1.4), 120960.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 120960.0, Days.ToSeconds(1.4));
    }
    if !cmp.Equal(Days.ToSeconds(0.06), 5184.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 5184.0, Days.ToSeconds(0.06));
    }
    if !cmp.Equal(Days.ToSeconds(0.2), 17280.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 17280.0, Days.ToSeconds(0.2));
    }
}

func TestConvertKnownDaysToMinutes(t * testing.T) {
    if !cmp.Equal(Days.ToMinutes(34.0), 48960.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 48960.0, Days.ToMinutes(34.0));
    }
    if !cmp.Equal(Days.ToMinutes(0.06), 86.4, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 86.4, Days.ToMinutes(0.06));
    }
    if !cmp.Equal(Days.ToMinutes(8.1), 11664.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 11664.0, Days.ToMinutes(8.1));
    }
}

func TestConvertKnownDaysToHours(t * testing.T) {
    if !cmp.Equal(Days.ToHours(8.1), 194.4, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 194.4, Days.ToHours(8.1));
    }
    if !cmp.Equal(Days.ToHours(0.2), 4.8, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 4.8, Days.ToHours(0.2));
    }
    if !cmp.Equal(Days.ToHours(121.0), 2904.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2904.0, Days.ToHours(121.0));
    }
}

func TestConvertKnownDaysToWeeks(t * testing.T) {
    if !cmp.Equal(Days.ToWeeks(121.0), 17.2857, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 17.2857, Days.ToWeeks(121.0));
    }
    if !cmp.Equal(Days.ToWeeks(7.2), 1.02857, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1.02857, Days.ToWeeks(7.2));
    }
    if !cmp.Equal(Days.ToWeeks(0.9), 0.128571, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.128571, Days.ToWeeks(0.9));
    }
}

func TestConvertKnownDaysToMonths(t * testing.T) {
    if !cmp.Equal(Days.ToMonths(0.9), 0.029589, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.029589, Days.ToMonths(0.9));
    }
    if !cmp.Equal(Days.ToMonths(108.0), 3.55068, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 3.55068, Days.ToMonths(108.0));
    }
    if !cmp.Equal(Days.ToMonths(55.0), 1.80822, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1.80822, Days.ToMonths(55.0));
    }
}

func TestConvertKnownDaysToYears(t * testing.T) {
    if !cmp.Equal(Days.ToYears(55.0), 0.150685, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.150685, Days.ToYears(55.0));
    }
    if !cmp.Equal(Days.ToYears(123.0), 0.336986, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.336986, Days.ToYears(123.0));
    }
    if !cmp.Equal(Days.ToYears(900.0), 2.46575, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2.46575, Days.ToYears(900.0));
    }
}

func TestConvertKnownHoursToMilliseconds(t * testing.T) {
    if !cmp.Equal(Hours.ToMilliseconds(0.006), 21600.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 21600.0, Hours.ToMilliseconds(0.006));
    }
    if !cmp.Equal(Hours.ToMilliseconds(0.09), 324000.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 324000.0, Hours.ToMilliseconds(0.09));
    }
    if !cmp.Equal(Hours.ToMilliseconds(0.007), 25200.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 25200.0, Hours.ToMilliseconds(0.007));
    }
}

func TestConvertKnownHoursToSeconds(t * testing.T) {
    if !cmp.Equal(Hours.ToSeconds(12.0), 43200.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 43200.0, Hours.ToSeconds(12.0));
    }
    if !cmp.Equal(Hours.ToSeconds(3.2), 11520.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 11520.0, Hours.ToSeconds(3.2));
    }
    if !cmp.Equal(Hours.ToSeconds(0.3), 1080.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1080.0, Hours.ToSeconds(0.3));
    }
}

func TestConvertKnownHoursToMinutes(t * testing.T) {
    if !cmp.Equal(Hours.ToMinutes(0.3), 18.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 18.0, Hours.ToMinutes(0.3));
    }
    if !cmp.Equal(Hours.ToMinutes(700.0), 42000.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 42000.0, Hours.ToMinutes(700.0));
    }
    if !cmp.Equal(Hours.ToMinutes(4.8), 288.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 288.0, Hours.ToMinutes(4.8));
    }
}

func TestConvertKnownHoursToDays(t * testing.T) {
    if !cmp.Equal(Hours.ToDays(4.8), 0.2, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.2, Hours.ToDays(4.8));
    }
    if !cmp.Equal(Hours.ToDays(190.0), 7.91667, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 7.91667, Hours.ToDays(190.0));
    }
    if !cmp.Equal(Hours.ToDays(8.5), 0.354167, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.354167, Hours.ToDays(8.5));
    }
}

func TestConvertKnownHoursToWeeks(t * testing.T) {
    if !cmp.Equal(Hours.ToWeeks(800.0), 4.7619, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 4.7619, Hours.ToWeeks(800.0));
    }
    if !cmp.Equal(Hours.ToWeeks(90.0), 0.535714, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.535714, Hours.ToWeeks(90.0));
    }
    if !cmp.Equal(Hours.ToWeeks(102.0), 0.607143, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.607143, Hours.ToWeeks(102.0));
    }
}

func TestConvertKnownHoursToMonths(t * testing.T) {
    if !cmp.Equal(Hours.ToMonths(102.0), 0.139726, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.139726, Hours.ToMonths(102.0));
    }
    if !cmp.Equal(Hours.ToMonths(9876.0), 13.52875, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 13.52875, Hours.ToMonths(9876.0));
    }
    if !cmp.Equal(Hours.ToMonths(100.8), 0.13808204, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.13808204, Hours.ToMonths(100.8));
    }
}

func TestConvertKnownHoursToYears(t * testing.T) {
    if !cmp.Equal(Hours.ToYears(9000.0), 1.027397, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1.027397, Hours.ToYears(9000.0));
    }
    if !cmp.Equal(Hours.ToYears(1003.0), 0.1144977, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.1144977, Hours.ToYears(1003.0));
    }
    if !cmp.Equal(Hours.ToYears(809.0), 0.0923516, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.0923516, Hours.ToYears(809.0));
    }
}

func TestConvertKnownMillisecondsToSeconds(t * testing.T) {
    if !cmp.Equal(Milliseconds.ToSeconds(8.0), 0.008, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.008, Milliseconds.ToSeconds(8.0));
    }
    if !cmp.Equal(Milliseconds.ToSeconds(780.0), 0.78, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.78, Milliseconds.ToSeconds(780.0));
    }
    if !cmp.Equal(Milliseconds.ToSeconds(900.0), 0.9, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.9, Milliseconds.ToSeconds(900.0));
    }
}

func TestConvertKnownMillisecondsToMinutes(t * testing.T) {
    if !cmp.Equal(Milliseconds.ToMinutes(900.0), 0.015, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.015, Milliseconds.ToMinutes(900.0));
    }
    if !cmp.Equal(Milliseconds.ToMinutes(67000.0), 1.1166667, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1.1166667, Milliseconds.ToMinutes(67000.0));
    }
    if !cmp.Equal(Milliseconds.ToMinutes(1234567.0), 20.57611667, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 20.57611667, Milliseconds.ToMinutes(1234567.0));
    }
}

func TestConvertKnownMillisecondsToHours(t * testing.T) {
    if !cmp.Equal(Milliseconds.ToHours(1234567.0), 0.3429352778, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.3429352778, Milliseconds.ToHours(1234567.0));
    }
    if !cmp.Equal(Milliseconds.ToHours(100900.0), 0.0280277778, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.0280277778, Milliseconds.ToHours(100900.0));
    }
    if !cmp.Equal(Milliseconds.ToHours(46000.0), 0.012777778, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.012777778, Milliseconds.ToHours(46000.0));
    }
}

func TestConvertKnownMillisecondsToDays(t * testing.T) {
    if !cmp.Equal(Milliseconds.ToDays(9000000.0), 0.1041666667, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.1041666667, Milliseconds.ToDays(9000000.0));
    }
    if !cmp.Equal(Milliseconds.ToDays(123456789.0), 1.42889802083, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1.42889802083, Milliseconds.ToDays(123456789.0));
    }
    if !cmp.Equal(Milliseconds.ToDays(89008900.0), 1.0301956019, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1.0301956019, Milliseconds.ToDays(89008900.0));
    }
}

func TestConvertKnownMillisecondsToWeeks(t * testing.T) {
    if !cmp.Equal(Milliseconds.ToWeeks(89008900.0), 0.14717080026, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.14717080026, Milliseconds.ToWeeks(89008900.0));
    }
    if !cmp.Equal(Milliseconds.ToWeeks(1234567890.0), 2.041282886905, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2.041282886905, Milliseconds.ToWeeks(1234567890.0));
    }
    if !cmp.Equal(Milliseconds.ToWeeks(100200300.0), 0.165675099206, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.165675099206, Milliseconds.ToWeeks(100200300.0));
    }
}

func TestConvertKnownMillisecondsToMonths(t * testing.T) {
    if !cmp.Equal(Milliseconds.ToMonths(100200300400.0), 38.102653412154631, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 38.102653412154631, Milliseconds.ToMonths(100200300400.0));
    }
    if !cmp.Equal(Milliseconds.ToMonths(90001000.0), 0.034246918329, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.034246918329, Milliseconds.ToMonths(90001000.0));
    }
    if !cmp.Equal(Milliseconds.ToMonths(8888888888.0), 3.382374104552, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 3.382374104552, Milliseconds.ToMonths(8888888888.0));
    }
}

func TestConvertKnownMillisecondsToYears(t * testing.T) {
    if !cmp.Equal(Milliseconds.ToYears(8888888888.0), 0.28167767558793383, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.28167767558793383, Milliseconds.ToYears(8888888888.0));
    }
    if !cmp.Equal(Milliseconds.ToYears(123456789123.0), 3.9121899074093087, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 3.9121899074093087, Milliseconds.ToYears(123456789123.0));
    }
    if !cmp.Equal(Milliseconds.ToYears(900080007000.0), 28.522399977032002, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 28.522399977032002, Milliseconds.ToYears(900080007000.0));
    }
}

func TestConvertKnownMinutesToMilliseconds(t * testing.T) {
    if !cmp.Equal(Minutes.ToMilliseconds(0.007), 420.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 420.0, Minutes.ToMilliseconds(0.007));
    }
    if !cmp.Equal(Minutes.ToMilliseconds(0.8), 48000.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 48000.0, Minutes.ToMilliseconds(0.8));
    }
    if !cmp.Equal(Minutes.ToMilliseconds(9.0), 540000.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 540000.0, Minutes.ToMilliseconds(9.0));
    }
}

func TestConvertKnownMinutesToSeconds(t * testing.T) {
    if !cmp.Equal(Minutes.ToSeconds(90.0), 5400.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 5400.0, Minutes.ToSeconds(90.0));
    }
    if !cmp.Equal(Minutes.ToSeconds(0.9), 54.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 54.0, Minutes.ToSeconds(0.9));
    }
    if !cmp.Equal(Minutes.ToSeconds(123.0), 7380.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 7380.0, Minutes.ToSeconds(123.0));
    }
}

func TestConvertKnownMinutesToHours(t * testing.T) {
    if !cmp.Equal(Minutes.ToHours(123.0), 2.05, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2.05, Minutes.ToHours(123.0));
    }
    if !cmp.Equal(Minutes.ToHours(77.0), 1.28333, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1.28333, Minutes.ToHours(77.0));
    }
    if !cmp.Equal(Minutes.ToHours(0.8), 0.0133333, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.0133333, Minutes.ToHours(0.8));
    }
}

func TestConvertKnownMinutesToDays(t * testing.T) {
    if !cmp.Equal(Minutes.ToDays(800.0), 0.555556, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.555556, Minutes.ToDays(800.0));
    }
    if !cmp.Equal(Minutes.ToDays(190.0), 0.131944, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.131944, Minutes.ToDays(190.0));
    }
    if !cmp.Equal(Minutes.ToDays(55.0), 0.0381944, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.0381944, Minutes.ToDays(55.0));
    }
}

func TestConvertKnownMinutesToWeeks(t * testing.T) {
    if !cmp.Equal(Minutes.ToWeeks(1900.0), 0.1884921, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.1884921, Minutes.ToWeeks(1900.0));
    }
    if !cmp.Equal(Minutes.ToWeeks(800.0), 0.0793651, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.0793651, Minutes.ToWeeks(800.0));
    }
    if !cmp.Equal(Minutes.ToWeeks(12345.0), 1.2247024, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1.2247024, Minutes.ToWeeks(12345.0));
    }
}

func TestConvertKnownMinutesToMonths(t * testing.T) {
    if !cmp.Equal(Minutes.ToMonths(1234.0), 0.02817349, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.02817349, Minutes.ToMonths(1234.0));
    }
    if !cmp.Equal(Minutes.ToMonths(90000.0), 2.0547923, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2.0547923, Minutes.ToMonths(90000.0));
    }
    if !cmp.Equal(Minutes.ToMonths(7100.0), 0.1621003, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.1621003, Minutes.ToMonths(7100.0));
    }
}

func TestConvertKnownMinutesToYears(t * testing.T) {
    if !cmp.Equal(Minutes.ToYears(900800.0), 1.71385084, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1.71385084, Minutes.ToYears(900800.0));
    }
    if !cmp.Equal(Minutes.ToYears(12345.0), 0.023487443, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.023487443, Minutes.ToYears(12345.0));
    }
    if !cmp.Equal(Minutes.ToYears(610910.0), 1.16230974, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1.16230974, Minutes.ToYears(610910.0));
    }
}

func TestConvertKnownMonthsToMilliseconds(t * testing.T) {
    if !cmp.Equal(Months.ToMilliseconds(0.00034), 894113.64, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 894113.64, Months.ToMilliseconds(0.00034));
    }
    if !cmp.Equal(Months.ToMilliseconds(0.001), 2629746.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2629746.0, Months.ToMilliseconds(0.001));
    }
    if !cmp.Equal(Months.ToMilliseconds(0.006), 15778476.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 15778476.0, Months.ToMilliseconds(0.006));
    }
}

func TestConvertKnownMonthsToSeconds(t * testing.T) {
    if !cmp.Equal(Months.ToSeconds(0.03), 78840.00, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 78840.00, Months.ToSeconds(0.03));
    }
    if !cmp.Equal(Months.ToSeconds(0.1), 262800.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 262800.0, Months.ToSeconds(0.1));
    }
    if !cmp.Equal(Months.ToSeconds(0.008), 21024.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 21024.0, Months.ToSeconds(0.008));
    }
}

func TestConvertKnownMonthsToMinutes(t * testing.T) {
    if !cmp.Equal(Months.ToMinutes(0.7), 30660.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 30660.0, Months.ToMinutes(0.7));
    }
    if !cmp.Equal(Months.ToMinutes(1.4), 61319.99, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 61319.99, Months.ToMinutes(1.4));
    }
    if !cmp.Equal(Months.ToMinutes(5.0), 219000.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 219000.0, Months.ToMinutes(5.0));
    }
}

func TestConvertKnownMonthsToHours(t * testing.T) {
    if !cmp.Equal(Months.ToHours(4.0), 2920.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2920.0, Months.ToHours(4.0));
    }
    if !cmp.Equal(Months.ToHours(0.3), 219.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 219.0, Months.ToHours(0.3));
    }
    if !cmp.Equal(Months.ToHours(4.5), 3285.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 3285.0, Months.ToHours(4.5));
    }
}

func TestConvertKnownMonthsToDays(t * testing.T) {
    if !cmp.Equal(Months.ToDays(4.5), 136.875, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 136.875, Months.ToDays(4.5));
    }
    if !cmp.Equal(Months.ToDays(90.0), 2737.5029, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2737.5029, Months.ToDays(90.0));
    }
    if !cmp.Equal(Months.ToDays(0.4), 12.1667, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 12.1667, Months.ToDays(0.4));
    }
}

func TestConvertKnownMonthsToWeeks(t * testing.T) {
    if !cmp.Equal(Months.ToWeeks(0.5), 2.17262, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2.17262, Months.ToWeeks(0.5));
    }
    if !cmp.Equal(Months.ToWeeks(88.0), 382.381120, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 382.381120, Months.ToWeeks(88.0));
    }
    if !cmp.Equal(Months.ToWeeks(12.6), 54.75006, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 54.75006, Months.ToWeeks(12.6));
    }
}

func TestConvertKnownMonthsToYears(t * testing.T) {
    if !cmp.Equal(Months.ToYears(12.6), 1.050001, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1.050001, Months.ToYears(12.6));
    }
    if !cmp.Equal(Months.ToYears(109.0), 9.08334, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 9.08334, Months.ToYears(109.0));
    }
    if !cmp.Equal(Months.ToYears(23.0), 1.91667, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1.91667, Months.ToYears(23.0));
    }
}

func TestConvertKnownNanosecondsToMicroseconds(t * testing.T) {
    if !cmp.Equal(Nanoseconds.ToMicroseconds(1234.0), 1.234, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1.234, Nanoseconds.ToMicroseconds(1234.0));
    }
    if !cmp.Equal(Nanoseconds.ToMicroseconds(90.9), 0.0909, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.0909, Nanoseconds.ToMicroseconds(90.9));
    }
    if !cmp.Equal(Nanoseconds.ToMicroseconds(70000.0), 70.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 70.0, Nanoseconds.ToMicroseconds(70000.0));
    }
}

func TestConvertKnownNanosecondsToMilliseconds(t * testing.T) {
    if !cmp.Equal(Nanoseconds.ToMilliseconds(70000.0), 0.07, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.07, Nanoseconds.ToMilliseconds(70000.0));
    }
    if !cmp.Equal(Nanoseconds.ToMilliseconds(123456.0), 0.123456, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.123456, Nanoseconds.ToMilliseconds(123456.0));
    }
    if !cmp.Equal(Nanoseconds.ToMilliseconds(900900.0), 0.9009, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.9009, Nanoseconds.ToMilliseconds(900900.0));
    }
}

func TestConvertKnownNanosecondsToSeconds(t * testing.T) {
    if !cmp.Equal(Nanoseconds.ToSeconds(900000000.0), 0.9, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.9, Nanoseconds.ToSeconds(900000000.0));
    }
    if !cmp.Equal(Nanoseconds.ToSeconds(123456789.0), 0.123456789, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.123456789, Nanoseconds.ToSeconds(123456789.0));
    }
    if !cmp.Equal(Nanoseconds.ToSeconds(800400700.0), 0.8004007, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.8004007, Nanoseconds.ToSeconds(800400700.0));
    }
}

func TestConvertKnownNanosecondsToMinutes(t * testing.T) {
    if !cmp.Equal(Nanoseconds.ToMinutes(1234567890.0), 0.0205761315, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.0205761315, Nanoseconds.ToMinutes(1234567890.0));
    }
    if !cmp.Equal(Nanoseconds.ToMinutes(800400700.0), 0.0133400116667, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.0133400116667, Nanoseconds.ToMinutes(800400700.0));
    }
    if !cmp.Equal(Nanoseconds.ToMinutes(800100800.0), 0.0133350133333, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.0133350133333, Nanoseconds.ToMinutes(800100800.0));
    }
}

func TestConvertKnownNanosecondsToHours(t * testing.T) {
    if !cmp.Equal(Nanoseconds.ToHours(1234567890123.0), 0.3429355250341667, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.3429355250341667, Nanoseconds.ToHours(1234567890123.0));
    }
    if !cmp.Equal(Nanoseconds.ToHours(900800700600.0), 0.250222416833333, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.250222416833333, Nanoseconds.ToHours(900800700600.0));
    }
    if !cmp.Equal(Nanoseconds.ToHours(66677788999.0), 0.018521608055278, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.018521608055278, Nanoseconds.ToHours(66677788999.0));
    }
}

func TestConvertKnownSecondsToMilliseconds(t * testing.T) {
    if !cmp.Equal(Seconds.ToMilliseconds(9.0), 9000.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 9000.0, Seconds.ToMilliseconds(9.0));
    }
    if !cmp.Equal(Seconds.ToMilliseconds(3.1), 3100.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 3100.0, Seconds.ToMilliseconds(3.1));
    }
    if !cmp.Equal(Seconds.ToMilliseconds(0.9), 900.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 900.0, Seconds.ToMilliseconds(0.9));
    }
}

func TestConvertKnownSecondsToMinutes(t * testing.T) {
    if !cmp.Equal(Seconds.ToMinutes(100.0), 1.66667, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1.66667, Seconds.ToMinutes(100.0));
    }
    if !cmp.Equal(Seconds.ToMinutes(9000.0), 150.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 150.0, Seconds.ToMinutes(9000.0));
    }
    if !cmp.Equal(Seconds.ToMinutes(2300.0), 38.33333, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 38.33333, Seconds.ToMinutes(2300.0));
    }
}

func TestConvertKnownSecondsToHours(t * testing.T) {
    if !cmp.Equal(Seconds.ToHours(2300.0), 0.6388889, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.6388889, Seconds.ToHours(2300.0));
    }
    if !cmp.Equal(Seconds.ToHours(9999.0), 2.7775, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2.7775, Seconds.ToHours(9999.0));
    }
    if !cmp.Equal(Seconds.ToHours(36000.0), 10.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 10.0, Seconds.ToHours(36000.0));
    }
}

func TestConvertKnownSecondsToDays(t * testing.T) {
    if !cmp.Equal(Seconds.ToDays(36000.0), 0.41666667, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.41666667, Seconds.ToDays(36000.0));
    }
    if !cmp.Equal(Seconds.ToDays(90000.0), 1.0416667, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1.0416667, Seconds.ToDays(90000.0));
    }
    if !cmp.Equal(Seconds.ToDays(190000.0), 2.19907407, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2.19907407, Seconds.ToDays(190000.0));
    }
}

func TestConvertKnownSecondsToWeeks(t * testing.T) {
    if !cmp.Equal(Seconds.ToWeeks(190000.0), 0.314153439, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.314153439, Seconds.ToWeeks(190000.0));
    }
    if !cmp.Equal(Seconds.ToWeeks(1234567.0), 2.041281415, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2.041281415, Seconds.ToWeeks(1234567.0));
    }
    if !cmp.Equal(Seconds.ToWeeks(100200.0), 0.165674603, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.165674603, Seconds.ToWeeks(100200.0));
    }
}

func TestConvertKnownSecondsToMonths(t * testing.T) {
    if !cmp.Equal(Seconds.ToMonths(1000000.0), 0.380517087, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.380517087, Seconds.ToMonths(1000000.0));
    }
    if !cmp.Equal(Seconds.ToMonths(9876543.0), 3.75819337, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 3.75819337, Seconds.ToMonths(9876543.0));
    }
    if !cmp.Equal(Seconds.ToMonths(200900.0), 0.0764458827, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.0764458827, Seconds.ToMonths(200900.0));
    }
}

func TestConvertKnownSecondsToYears(t * testing.T) {
    if !cmp.Equal(Seconds.ToYears(123456789.0), 3.91478909817, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 3.91478909817, Seconds.ToYears(123456789.0));
    }
    if !cmp.Equal(Seconds.ToYears(100900.0), 0.00319951801, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.00319951801, Seconds.ToYears(100900.0));
    }
    if !cmp.Equal(Seconds.ToYears(987654.0), 0.0313183029, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.0313183029, Seconds.ToYears(987654.0));
    }
}

func TestConvertKnownWeeksToMilliseconds(t * testing.T) {
    if !cmp.Equal(Weeks.ToMilliseconds(0.001), 604800.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 604800.0, Weeks.ToMilliseconds(0.001));
    }
    if !cmp.Equal(Weeks.ToMilliseconds(0.005), 3024000.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 3024000.0, Weeks.ToMilliseconds(0.005));
    }
    if !cmp.Equal(Weeks.ToMilliseconds(0.0009), 544320.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 544320.0, Weeks.ToMilliseconds(0.0009));
    }
}

func TestConvertKnownWeeksToSeconds(t * testing.T) {
    if !cmp.Equal(Weeks.ToSeconds(0.1), 60480.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 60480.0, Weeks.ToSeconds(0.1));
    }
    if !cmp.Equal(Weeks.ToSeconds(0.08), 48384.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 48384.0, Weeks.ToSeconds(0.08));
    }
    if !cmp.Equal(Weeks.ToSeconds(3.0), 1814400.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1814400.0, Weeks.ToSeconds(3.0));
    }
}

func TestConvertKnownWeeksToMinutes(t * testing.T) {
    if !cmp.Equal(Weeks.ToMinutes(2.0), 20160.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 20160.0, Weeks.ToMinutes(2.0));
    }
    if !cmp.Equal(Weeks.ToMinutes(0.4), 4032.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 4032.0, Weeks.ToMinutes(0.4));
    }
    if !cmp.Equal(Weeks.ToMinutes(0.02), 201.6, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 201.6, Weeks.ToMinutes(0.02));
    }
}

func TestConvertKnownWeeksToHours(t * testing.T) {
    if !cmp.Equal(Weeks.ToHours(0.02), 3.36, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 3.36, Weeks.ToHours(0.02));
    }
    if !cmp.Equal(Weeks.ToHours(77.0), 12936.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 12936.0, Weeks.ToHours(77.0));
    }
    if !cmp.Equal(Weeks.ToHours(9.2), 1545.6, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1545.6, Weeks.ToHours(9.2));
    }
}

func TestConvertKnownWeeksToDays(t * testing.T) {
    if !cmp.Equal(Weeks.ToDays(9.2), 64.4, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 64.4, Weeks.ToDays(9.2));
    }
    if !cmp.Equal(Weeks.ToDays(169.0), 1183.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1183.0, Weeks.ToDays(169.0));
    }
    if !cmp.Equal(Weeks.ToDays(4.0), 28.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 28.0, Weeks.ToDays(4.0));
    }
}

func TestConvertKnownWeeksToMonths(t * testing.T) {
    if !cmp.Equal(Weeks.ToMonths(4.0), 0.920547, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.920547, Weeks.ToMonths(4.0));
    }
    if !cmp.Equal(Weeks.ToMonths(900.0), 207.12319687, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 207.12319687, Weeks.ToMonths(900.0));
    }
    if !cmp.Equal(Weeks.ToMonths(3.8), 0.87452, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.87452, Weeks.ToMonths(3.8));
    }
}

func TestConvertKnownWeeksToYears(t * testing.T) {
    if !cmp.Equal(Weeks.ToYears(3.0), 0.0575342, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.0575342, Weeks.ToYears(3.0));
    }
    if !cmp.Equal(Weeks.ToYears(235.0), 4.50685, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 4.50685, Weeks.ToYears(235.0));
    }
    if !cmp.Equal(Weeks.ToYears(1090.0), 20.90411, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 20.90411, Weeks.ToYears(1090.0));
    }
}

func TestConvertKnownYearsToMilliseconds(t * testing.T) {
    if !cmp.Equal(Years.ToMilliseconds(0.001), 31556952.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 31556952.0, Years.ToMilliseconds(0.001));
    }
    if !cmp.Equal(Years.ToMilliseconds(0.0009), 28401256.8, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 28401256.8, Years.ToMilliseconds(0.0009));
    }
    if !cmp.Equal(Years.ToMilliseconds(0.00034), 10729363.680000002, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 10729363.680000002, Years.ToMilliseconds(0.00034));
    }
}

func TestConvertKnownYearsToSeconds(t * testing.T) {
    if !cmp.Equal(Years.ToSeconds(0.06), 1892160.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1892160.0, Years.ToSeconds(0.06));
    }
    if !cmp.Equal(Years.ToSeconds(0.009), 283824.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 283824.0, Years.ToSeconds(0.009));
    }
    if !cmp.Equal(Years.ToSeconds(0.02), 630720.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 630720.0, Years.ToSeconds(0.02));
    }
}

func TestConvertKnownYearsToMinutes(t * testing.T) {
    if !cmp.Equal(Years.ToMinutes(0.02), 10512.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 10512.0, Years.ToMinutes(0.02));
    }
    if !cmp.Equal(Years.ToMinutes(0.3), 157680.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 157680.0, Years.ToMinutes(0.3));
    }
    if !cmp.Equal(Years.ToMinutes(0.001), 525.6, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 525.6, Years.ToMinutes(0.001));
    }
}

func TestConvertKnownYearsToHours(t * testing.T) {
    if !cmp.Equal(Years.ToHours(0.001), 8.76, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 8.76, Years.ToHours(0.001));
    }
    if !cmp.Equal(Years.ToHours(0.3), 2628.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2628.0, Years.ToHours(0.3));
    }
    if !cmp.Equal(Years.ToHours(2.0), 17520.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 17520.0, Years.ToHours(2.0));
    }
}

func TestConvertKnownYearsToDays(t * testing.T) {
    if !cmp.Equal(Years.ToDays(2.0), 730.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 730.0, Years.ToDays(2.0));
    }
    if !cmp.Equal(Years.ToDays(1009.0), 368285.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 368285.0, Years.ToDays(1009.0));
    }
    if !cmp.Equal(Years.ToDays(7.0), 2555.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2555.0, Years.ToDays(7.0));
    }
}

func TestConvertKnownYearsToWeeks(t * testing.T) {
    if !cmp.Equal(Years.ToWeeks(7.0), 365.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 365.0, Years.ToWeeks(7.0));
    }
    if !cmp.Equal(Years.ToWeeks(1.3), 67.7857, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 67.7857, Years.ToWeeks(1.3));
    }
    if !cmp.Equal(Years.ToWeeks(88.0), 4588.5839, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 4588.5839, Years.ToWeeks(88.0));
    }
}

func TestConvertKnownYearsToMonths(t * testing.T) {
    if !cmp.Equal(Years.ToMonths(6.0), 71.9999, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 71.9999, Years.ToMonths(6.0));
    }
    if !cmp.Equal(Years.ToMonths(12.0), 144.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 144.0, Years.ToMonths(12.0));
    }
    if !cmp.Equal(Years.ToMonths(0.3), 3.6, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 3.6, Years.ToMonths(0.3));
    }
}

