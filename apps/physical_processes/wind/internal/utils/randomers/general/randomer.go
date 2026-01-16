package general

import (
	"crypto/rand"
	"math/big"
)

// TODO: docs and discription
func randomInt(min, max int64) (int64, error) {
	if min > max {
		min, max = max, min
	}
	rangeSize := big.NewInt(max - min + 1)
	n, err := rand.Int(rand.Reader, rangeSize)
	if err != nil {
		return 0, err
	}
	return n.Int64() + min, nil
}

// TODO: docs and discription
func randomFloat(min, max float64, mathPr int) float64 {
	if mathPr%10 != 0 {
		mathPr = 100
	}
	minInt := int64(min * float64(mathPr))
	maxInt := int64(max * float64(mathPr))
	val, err := randomInt(minInt, maxInt)
	if err != nil {
		/*
			TODO:
			! В рамках симуляции возвращаем min
			! В дальнейшем нужен логгер на данном месте
		*/
		return min
	}
	return float64(val) / float64(mathPr)
	// mathematical precision
}

// TODO: docs and discription
func pickRandomElement(list []string) string {
	idx, _ := randomInt(0, int64(len(list)-1))
	return list[idx]
}
