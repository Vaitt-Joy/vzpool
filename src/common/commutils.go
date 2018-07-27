package common

import "strconv"

func toInt(val string) int {
	i, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return i
}

func toInt8(val string) int8 {
	i, err := strconv.ParseInt(val, 10, 8)
	if err != nil {
		return int8(0)
	}
	return int8(i)
}

func toInt32(val string) int32 {
	i, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return int32(0)
	}
	return int32(i)
}

func toInt64(val string) int64 {
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0
	}
	return int64(i)
}

func toUint8(val string) uint8 {
	i, err := strconv.ParseUint(val, 10, 8)
	if err != nil {
		return uint8(0)
	}
	return uint8(i)
}

/* 16 进制*/
func toUint8Hex(val string) uint8 {
	i, err := strconv.ParseUint(val, 16, 8)
	if err != nil {
		return uint8(0)
	}
	return uint8(i)
}

func toUint32(val string) uint32 {
	i, err := strconv.ParseUint(val, 10, 32)
	if err != nil {
		return uint32(0)
	}
	return uint32(i)
}

func toUint64(val string) uint64 {
	i, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return 0
	}
	return uint64(i)
}

func toFloat32(val string) float32 {
	i, err := strconv.ParseFloat(val, 32)
	if err != nil {
		return float32(0)
	}
	return float32(i)
}

func toFloat64(val string) float64 {
	i, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return float64(0)
	}
	return float64(i)
}

func toBool(val string) bool {

}
