package util

import "strconv"

func StringToFloat32(s string) float32 {
	result, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0
	}

	return float32(result)
}

func StringToFloat64(s string) float64 {
	result, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}

	return float64(result)
}
