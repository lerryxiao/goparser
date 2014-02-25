package parser

import (
	"strconv"
)

func ParseInt(item interface{}) int {
	switch v := item.(type) {
	case int:
		return v
	case int32:
		return int(v)
	case int64:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	case string:
		i, _ := strconv.Atoi(v)
		return i
	default:
		return 0
	}
}

func ParseInt32(item interface{}) int32 {
	switch v := item.(type) {
	case int:
		return int32(v)
	case int32:
		return v
	case int64:
		return int32(v)
	case float32:
		return int32(v)
	case float64:
		return int32(v)
	case string:
		i, _ := strconv.Atoi(v)
		return int32(i)
	default:
		return 0
	}
}

func ParseInt64(item interface{}) int64 {
	switch v := item.(type) {
	case int:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case float32:
		return int64(v)
	case float64:
		return int64(v)
	case string:
		i, _ := strconv.Atoi(v)
		return int64(i)
	default:
		return 0
	}
}

func ParseFloat64(item interface{}) float64 {
	switch v := item.(type) {
	case int:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case float32:
		return float64(v)
	case float64:
		return v
	case string:
		f, _ := strconv.ParseFloat(v, 10)
		return f
	default:
		return 0
	}
}

func ParseString(item interface{}) string {
	switch v := item.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case float64:
		return strconv.FormatFloat(v, 10, 2, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 10, 2, 10)
	default:
		return ""
	}
}
