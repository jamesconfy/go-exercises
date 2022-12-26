// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package tempconv

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

//!-
