package randomers

// Just wrapper fo exports
func GetRandInt(min, max int64) int64 {
	val, _ := randomInt(min, max)
	return val
}
