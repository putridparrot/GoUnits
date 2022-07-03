// <auto-generated>
// This code was generated by the UnitCodeGenerator tool
//
// Changes to this file will be lost if the code is regenerated
// </auto-generated>

package pressure

// Atmospheres (Atmospheres) conversion functions
type Atmospheres float64

// ToBars converts the supplied Atmospheres value to Bars
func (value Atmospheres) ToBars() float64 { 
	return float64(value * 1.01325)
}
// ToPascals converts the supplied Atmospheres value to Pascals
func (value Atmospheres) ToPascals() float64 { 
	return float64(value * 101325.0)
}
// ToTorrs converts the supplied Atmospheres value to Torrs
func (value Atmospheres) ToTorrs() float64 { 
	return float64(value * 760.0)
}
// ToPsi converts the supplied Atmospheres value to Psi
func (value Atmospheres) ToPsi() float64 { 
	return float64(value * 14.69596432068)
}

// Bars (Bars) conversion functions
type Bars float64

// ToAtmospheres converts the supplied Bars value to Atmospheres
func (value Bars) ToAtmospheres() float64 { 
	return float64(value / 1.01325)
}
// ToPascals converts the supplied Bars value to Pascals
func (value Bars) ToPascals() float64 { 
	return float64(value / 0.00001)
}
// ToTorrs converts the supplied Bars value to Torrs
func (value Bars) ToTorrs() float64 { 
	return float64(value * 750.0616827042)
}
// ToPsi converts the supplied Bars value to Psi
func (value Bars) ToPsi() float64 { 
	return float64(value * 14.50378911491)
}

// Pascals (Pascals) conversion functions
type Pascals float64

// ToAtmospheres converts the supplied Pascals value to Atmospheres
func (value Pascals) ToAtmospheres() float64 { 
	return float64(value / 101325.0)
}
// ToBars converts the supplied Pascals value to Bars
func (value Pascals) ToBars() float64 { 
	return float64(value * 0.00001)
}
// ToTorrs converts the supplied Pascals value to Torrs
func (value Pascals) ToTorrs() float64 { 
	return float64(value * 0.007500616827042)
}
// ToPsi converts the supplied Pascals value to Psi
func (value Pascals) ToPsi() float64 { 
	return float64(value * 0.0001450378911491)
}

// Psi (Psi) conversion functions
type Psi float64

// ToBars converts the supplied Psi value to Bars
func (value Psi) ToBars() float64 { 
	return float64(value / 14.50378911491)
}
// ToPascals converts the supplied Psi value to Pascals
func (value Psi) ToPascals() float64 { 
	return float64(value / 0.0001450378911491)
}
// ToAtmospheres converts the supplied Psi value to Atmospheres
func (value Psi) ToAtmospheres() float64 { 
	return float64(value / 14.69596432068)
}
// ToTorrs converts the supplied Psi value to Torrs
func (value Psi) ToTorrs() float64 { 
	return float64(value / 0.01933679515879)
}

// Torrs (Torrs) conversion functions
type Torrs float64

// ToBars converts the supplied Torrs value to Bars
func (value Torrs) ToBars() float64 { 
	return float64(value / 750.0616827042)
}
// ToPascals converts the supplied Torrs value to Pascals
func (value Torrs) ToPascals() float64 { 
	return float64(value / 0.007500616827042)
}
// ToAtmospheres converts the supplied Torrs value to Atmospheres
func (value Torrs) ToAtmospheres() float64 { 
	return float64(value / 760.0)
}
// ToPsi converts the supplied Torrs value to Psi
func (value Torrs) ToPsi() float64 { 
	return float64(value * 0.01933679515879)
}
