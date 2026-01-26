package random

import (
	"crypto/rand"
	"errors"
	"math"
	"math/big"
	"strconv"
	"strings"
)

// createRandValue generates a cryptographically secure random integer within [min, max].
// It handles the big.Int complexity required by the crypto/rand package.
func createRandValue(min, max int64) (int64, error) {
	if min == max {
		return min, nil
	}
	if min > max {
		min, max = max, min
	}

	// Calculate range size: max - min + 1
	bgMax := big.NewInt(max)
	bgMin := big.NewInt(min)
	diff := new(big.Int).Sub(bgMax, bgMin)
	plusOne := big.NewInt(1)
	rangeBig := new(big.Int).Add(diff, plusOne)

	// Generate random number in [0, rangeBig)
	num, err := rand.Int(rand.Reader, rangeBig)
	if err != nil {
		return 0, err
	}

	// Shift back to [min, max]
	res := new(big.Int).Add(num, bgMin)
	return res.Int64(), nil
}

// CreateRandInt is a public wrapper to generate a random integer between min and max.
func CreateRandInt(min, max int64) (int64, error) {
	return createRandValue(min, max)
}

// CreateRandFloat generates a random float with specific decimal precision based on input scale.
// Note: For pure statistical generation (like Weibull), consider using GenerateCanonicalFloat.
func CreateRandFloat(min, max float64) (float64, error) {
	// Helper to determine decimal places and convert float to int64.
	convertFromFloatToInt := func(num float64) (int64, int64) {
		str := strconv.FormatFloat(num, 'f', -1, 64)
		idx := strings.Index(str, ".")
		decimalPlaces := 0
		if idx != -1 {
			decimalPlaces = len(str) - idx - 1
		}
		multiplier := int64(math.Pow10(decimalPlaces))
		res := int64(math.Round(num * float64(multiplier)))
		return res, multiplier
	}

	// Scale inputs to integers to use the secure random integer generator.
	start, startMult := convertFromFloatToInt(min)
	end, endMult := convertFromFloatToInt(max)

	// Normalize multipliers.
	var commonMult int64
	if startMult > endMult {
		commonMult = startMult
	} else {
		commonMult = endMult
	}

	startScaled := start * (commonMult / startMult)
	endScaled := end * (commonMult / endMult)

	randInt, err := createRandValue(startScaled, endScaled)
	if err != nil {
		return 0, err
	}

	result := float64(randInt) / float64(commonMult)
	return result, nil
}

// GetRandArrVal returns a random element from a generic slice.
func GetRandArrVal[T any](slice []T) (T, error) {
	if len(slice) == 0 {
		var zero T
		return zero, errors.New("slice is empty")
	}
	randIdx, err := createRandValue(0, int64(len(slice)-1))
	if err != nil {
		var zero T
		return zero, err
	}
	return slice[randIdx], nil
}

// GenerateCanonicalFloat generates a random float64 in the range [0.0, 1.0).
// It mimics math/rand.Float64() but uses crypto/rand for higher entropy.
// This is essential for accurate physical simulations (e.g., Weibull distribution).
func GenerateCanonicalFloat() (float64, error) {
	// 1 << 53 is the size of the significand for float64, ensuring max precision.
	const maxInt = int64(1 << 53)
	randInt, err := createRandValue(0, maxInt)
	if err != nil {
		return 0.0, err
	}
	return float64(randInt) / float64(maxInt), nil
}
