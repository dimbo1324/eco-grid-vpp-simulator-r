package physics

import "math"

func WeibullWindSpeed(u, shape, scale float64) float64 {
	if u <= 0 || u >= 1 {
		return 0.0
	}
	return scale * math.Pow(-math.Log(1-u), 1.0/shape)
}

func WindTurbinePower(rho, area, windSpeed, cp float64) float64 {
	return 0.5 * rho * area * math.Pow(windSpeed, 3) * cp
}
