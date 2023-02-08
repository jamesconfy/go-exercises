// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 180.

// Package tempconv performs Celsius and Fahrenheit temperature computations.
package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KtoC converts a Kelvin temperature to Celsius.
func KtoC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// CtoK converts a Celsius temperature to Kelvin.
func CtoK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// KtoF converts a Kelvin temperature to Fahrenheit.
func KtoF(k Kelvin) Fahrenheit { return Fahrenheit((k-273.15)*1.8 + 32) }

// FtoK converts a Fahrenheit temperature to Kelvin.
func FtoK(f Fahrenheit) Kelvin { return Kelvin((f-32)/1.8 + 273.15) }

// !-
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

/*
//!+flagvalue
package flag

// Value is the interface to the value stored in a flag.
type Value interface {
	String() string
	Set(string) error
}
//!-flagvalue
*/

// !+celsiusFlag
// *celsiusFlag satisfies the flag.Value interface.
type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Celsius = KtoC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

//!-celsiusFlag

//!+CelsiusFlag

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

//!-CelsiusFlag
