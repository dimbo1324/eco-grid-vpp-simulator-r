package physics

import "math"

// WeibullWindSpeed generates a wind speed sample using the Weibull distribution Inverse CDF.
// This is a standard method for stochastic wind simulation.
//
// Parameters:
//   - u: A random uniform float in the range [0, 1).
//   - shape (k): The shape parameter (typically ~2.0 for wind).
//   - scale (lambda): The scale parameter (related to average wind speed).
func WeibullWindSpeed(u, shape, scale float64) float64 {
	// Validate input to avoid math errors with Log.
	if u <= 0 || u >= 1 {
		return 0.0
	}
	// Formula: v = scale * (-ln(1-u))^(1/shape)
	return scale * math.Pow(-math.Log(1-u), 1.0/shape)
}

// WindTurbinePower calculates the mechanical power captured by the turbine rotor in Watts.
//
// Formula: P = 0.5 * rho * Area * v^3 * Cp
//   - rho: Air density (kg/m^3).
//   - cp: Power Coefficient (efficiency factor, Betz limit is 0.59).
func WindTurbinePower(rho, area, windSpeed, cp float64) float64 {
	return 0.5 * rho * area * math.Pow(windSpeed, 3) * cp
}
