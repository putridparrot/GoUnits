# GoUnits

[![Build](https://github.com/putridparrot/GoUnits/actions/workflows/build.yml/badge.svg)](https://github.com/putridparrot/GoUnits/actions/workflows/build.yml)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/putridparrot/GoUnits/blob/master/LICENSE.md)
[![GitHub Releases](https://img.shields.io/github/release/putridparrot/GoUnits.svg)](https://github.com/putridparrot/GoUnits/releases)
[![GitHub Issues](https://img.shields.io/github/issues/putridparrot/GoUnits.svg)](https://github.com/putridparrot/GoUnits/issues)

The main aim for this project is to produce conversion functions for the various units of measurement (listed below). The code and the unit tests
are auto-generated by the UnitCodeGenerator tool. Hence any changes need to made within the definitions file.

### Unit conversions also available for 

[C#](https://github.com/putridparrot/PutridParrot.Units)  
[Dart](https://github.com/putridparrot/DartUnits)  
[F#](https://github.com/putridparrot/FSharp.Units)  
[Java](https://github.com/putridparrot/JavaUnits)  
[Python](https://github.com/putridparrot/PyUnits)  
[Rust](https://github.com/putridparrot/RustUnits)   
[Swift](https://github.com/putridparrot/SwiftUnits)   
[TypeScript](https://github.com/putridparrot/unit-conversions)  

### Example

Example of usage:

```
import 'package:conversion_units/conversion_units.dart';

var fahrenheit = Celsius.toFahrenheit(12.0);
```

Each unit of measure then includes functions to convert to each for example, converting each Temperature unit of measure to each other Temperature.

### Angle

	* Degrees (deg)
	* Gradians (grad)
	* Milliradians (mrad)
	* Minute Of Arc (arcmin)
	* Radians (rad)
	* Seconds Of Arc (arcseconds)

### Area

	* Acres (acre)
	* Hectares (hectare)
	* Square Feet (feet2)
	* Square Inches (inch2)
	* Square Kilometres (km2)
	* Square Metres (m2)
	* Square Miles (mile2)
	* Square Yards (yard2)

### Data Storage

	* Bits (b)
	* Gigabits (Gb)
	* Gigabytes (GB)
	* Kibibits (kibibit)
	* Kilobits (k)
	* Kilobytes (KB)
	* Mebibits (mebibit)
	* Megabits (Mb)
	* Megabytes (MB)
	* Terabits (Tb)
	* Terabytes (TB)

### Data Transfer Rate

	* Bits Per Second (bps)
	* GigaBits Per Second (Gbps)
	* GigaBytes Per Second (GBps)
	* Kibibits Per Second (Kibitps)
	* KiloBits Per Second (kbps)
	* KiloBytes Per Second (kBps)
	* Mebibits Per Second (Mibit)
	* MegaBits Per Second (Mbps)
	* MegaBytes Per Second (MBps)
	* TeraBits Per Second (Tbps)
	* TeraBytes Per Second (TBps)

### Energy

	* Btu (btu)
	* Calories (cal)
	* Electronvolts (eV)
	* Foot Pounds (ftlb)
	* Joules (J)
	* Kilocalories (kCal)
	* Kilojoules (kJ)
	* Kilowatt Hours (kWh)
	* US Therms (ustherm)
	* Watt Hours (Wh)

### Force

	* Dynes (dyn)
	* Kilogram-force (kp)
	* Newtons (N)
	* Poundals (pdl)

### Frequency

	* Gigahertz (GHz)
	* Hertz (Hz)
	* Kilohertz (kHz)
	* Megahertz (MHz)

### Fuel Economy

	* Kilometre Per Litre (kml)
	* Litres Per 100 Kilometres (l100km)
	* Miles Per Gallon (mpg)
	* US Miles Per Gallon (usmpg)

### Length

	* Centimetres (cm)
	* Feet (ft)
	* Inches (inch)
	* Kilometres (km)
	* Metres (m)
	* Miles (mile)
	* Millimetres (mm)
	* Nautical Miles (NM)
	* Yards (yard)

### Luminous Energy

	* LumenHour (lmh)
	* LumenMinute (lmmin)
	* LumenSecond (lms)
	* Talbot (T)

### Magnetomotive Force

	* Ampere-turns (AT)
	* Gilberts (Gi)

### Mass

	* Carats (ct)
	* Grams (g)
	* Kilograms (kg)
	* Milligrams (mg)
	* Ounces (oz)
	* Pounds (lb)
	* Stones (stone)
	* Tonnes (tonne)

### Pressure

	* Atmospheres (atm)
	* Bars (b)
	* Pascals (Pa)
	* Psi (psi)
	* Torrs (Torr)

### Speed

	* Feet Per Second (fps)
	* Kilometres Per Hour (kph)
	* Knots (knot)
	* Metres Per Second (mps)
	* Miles Per Hour (mph)

### Temperature

	* Celsius (C)
	* Fahrenheit (F)
	* Kelvin (K)
	* Rankine (R)

### Time

	* Centuries (century)
	* Days (day)
	* Decades (decade)
	* Hours (hour)
	* Microseconds (microsecond)
	* Milliseconds (millisecond)
	* Minutes (minute)
	* Months (month)
	* Nanoseconds (nanosecond)
	* Seconds (second)
	* Weeks (week)
	* Years (year)

### Volume

	* Fluid Ounces (floz)
	* Gallons (gal)
	* Kilolitres (kl)
	* Litres (l)
	* Millilitres (ml)
	* Pints (pt)
	* Quarts (qt)
	* Tablespoons (tbsp)
	* Teaspoons (tsp)
	* US Cups (uscup)
	* US Fluid Ounces (usfloz)
	* US Gallons (usgal)
	* US Pints (uspt)
	* US Quarts (usqt)
	* US Tablespoons (ustbsp)
	* US Teaspoons (ustsp)

### Apps

The "My Unit Conversions" app. (for Android, iOS and Windows) uses the C# (PutridParrot.Units) version of this library and can be obtained from the following stores/locations.

[![Get it on Google Play](https://apps.putridparrot.com/Images/googleplay153x46.png)](https://play.google.com/store/apps/details?id=com.putridparrot.MyUnitConversion&pcampaignid=pcampaignidMKT-Other-global-all-co-prtnr-py-PartBadge-Mar2515-1)
[![Download on the Apple Store](https://apps.putridparrot.com/Images/applestore153x46.png)](https://apps.apple.com/app/my-unit-conversion/id1600275661)
[![Available on Fire](https://apps.putridparrot.com/Images/fire153x46.png)](https://www.amazon.co.uk/MTBSOFTWARE-LIMITED-My-Unit-Conversion/dp/B09RTBBGNM/)
[![Available on Windows](https://apps.putridparrot.com/Images/MS_864X312.svg)](https://apps.microsoft.com/store/detail/my-unit-conversion/9NK6CTDN0L2L)

The Mac version of the app. uses this Swift package and can be obtained from the following location

[![Download on the Apple Store](https://apps.putridparrot.com/Images/applestore153x46.png)](https://apps.apple.com/app/my-unit-conversion/id1600275661)
