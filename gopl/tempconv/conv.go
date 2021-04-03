package tempconv

func CToF(f Celsius) Fahrenheit {
	return Fahrenheit((f*5/9 + 32))
}

func CToK(k Celsius) Kelvins {
	return Kelvins(k + -273.15)
}
