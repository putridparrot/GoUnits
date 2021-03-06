// <auto-generated>
// This code was generated by the UnitCodeGenerator tool
//
// Changes to this file will be lost if the code is regenerated
// </auto-generated>

package frequency

// Gigahertz (Gigahertz) conversion functions
type Gigahertz float64

// ToHertz converts the supplied Gigahertz value to Hertz
func (value Gigahertz) ToHertz() float64 { 
	return float64(value * 1e+9)
}
// ToKilohertz converts the supplied Gigahertz value to Kilohertz
func (value Gigahertz) ToKilohertz() float64 { 
	return float64(value * 1e+6)
}
// ToMegahertz converts the supplied Gigahertz value to Megahertz
func (value Gigahertz) ToMegahertz() float64 { 
	return float64(value * 1000.0)
}

// Hertz (Hertz) conversion functions
type Hertz float64

// ToKilohertz converts the supplied Hertz value to Kilohertz
func (value Hertz) ToKilohertz() float64 { 
	return float64(value / 1000.0)
}
// ToMegahertz converts the supplied Hertz value to Megahertz
func (value Hertz) ToMegahertz() float64 { 
	return float64(value / 1e+6)
}
// ToGigahertz converts the supplied Hertz value to Gigahertz
func (value Hertz) ToGigahertz() float64 { 
	return float64(value / 1e+9)
}

// Kilohertz (Kilohertz) conversion functions
type Kilohertz float64

// ToHertz converts the supplied Kilohertz value to Hertz
func (value Kilohertz) ToHertz() float64 { 
	return float64(value * 1000.0)
}
// ToMegahertz converts the supplied Kilohertz value to Megahertz
func (value Kilohertz) ToMegahertz() float64 { 
	return float64(value / 1000.0)
}
// ToGigahertz converts the supplied Kilohertz value to Gigahertz
func (value Kilohertz) ToGigahertz() float64 { 
	return float64(value / 1e+6)
}

// Megahertz (Megahertz) conversion functions
type Megahertz float64

// ToHertz converts the supplied Megahertz value to Hertz
func (value Megahertz) ToHertz() float64 { 
	return float64(value * 1e+6)
}
// ToKilohertz converts the supplied Megahertz value to Kilohertz
func (value Megahertz) ToKilohertz() float64 { 
	return float64(value * 1000.0)
}
// ToGigahertz converts the supplied Megahertz value to Gigahertz
func (value Megahertz) ToGigahertz() float64 { 
	return float64(value / 1000.0)
}
