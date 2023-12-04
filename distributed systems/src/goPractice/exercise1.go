package main

func Sqrt(x float64) float64 {
	z := 1.0
	counter := 0
	for counter < 10 {
		z -= (z*z - x) / (2 * z)
		counter++
	}
	return z
}
