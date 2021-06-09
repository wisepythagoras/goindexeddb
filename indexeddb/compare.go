package indexeddb

import (
	"errors"
	"reflect"
)

// Compare compares two values.
func Compare(left interface{}, right interface{}) (int, error) {
	leftRt := reflect.TypeOf(left)
	rightRt := reflect.TypeOf(right)
	leftKind := leftRt.Kind()
	rightKind := rightRt.Kind()

	if leftKind != rightKind {
		return 0, errors.New("Invalid type comparison")
	}

	if leftKind == reflect.Slice || leftKind == reflect.Array {
		leftArr := left.([]interface{})
		rightArr := right.([]interface{})

		if len(leftArr) > len(rightArr) {
			return 1, nil
		} else if len(leftArr) < len(rightArr) {
			return -1, nil
		}

		return 0, nil
	} else if leftKind == reflect.Interface || leftKind == reflect.Map {
		if reflect.DeepEqual(left, right) {
			return 0, nil
		} else {
			return 1, nil
		}
	} else if leftKind == reflect.Bool {
		if left == right {
			return 0, nil
		} else if left.(bool) == true {
			return 1, nil
		}

		return -1, nil
	} else if leftKind == reflect.Int ||
		leftKind == reflect.Int8 ||
		leftKind == reflect.Int16 ||
		leftKind == reflect.Int32 ||
		leftKind == reflect.Int64 {
		var leftInt int64
		var rightInt int64

		if leftKind == reflect.Int {
			leftInt = int64(left.(int))
			rightInt = int64(right.(int))
		} else if leftKind == reflect.Int8 {
			leftInt = int64(left.(int8))
			rightInt = int64(right.(int8))
		} else if leftKind == reflect.Int16 {
			leftInt = int64(left.(int16))
			rightInt = int64(right.(int16))
		} else if leftKind == reflect.Int32 {
			leftInt = int64(left.(int32))
			rightInt = int64(right.(int32))
		} else {
			leftInt = left.(int64)
			rightInt = right.(int64)
		}

		if leftInt > rightInt {
			return 1, nil
		} else if leftInt < rightInt {
			return -1, nil
		}

		return 0, nil
	} else if leftKind == reflect.Uint ||
		leftKind == reflect.Uint8 ||
		leftKind == reflect.Uint16 ||
		leftKind == reflect.Uint32 ||
		leftKind == reflect.Uint64 {
		var leftInt uint64
		var rightInt uint64

		if leftKind == reflect.Int {
			leftInt = uint64(left.(uint))
			rightInt = uint64(right.(uint))
		} else if leftKind == reflect.Int8 {
			leftInt = uint64(left.(uint8))
			rightInt = uint64(right.(uint8))
		} else if leftKind == reflect.Int16 {
			leftInt = uint64(left.(uint16))
			rightInt = uint64(right.(uint16))
		} else if leftKind == reflect.Int32 {
			leftInt = uint64(left.(uint32))
			rightInt = uint64(right.(uint32))
		} else {
			leftInt = left.(uint64)
			rightInt = right.(uint64)
		}

		if leftInt > rightInt {
			return 1, nil
		} else if leftInt < rightInt {
			return -1, nil
		}

		return 0, nil
	} else if leftKind == reflect.Float32 || leftKind == reflect.Float64 {
		var leftFloat float64
		var rightFloat float64

		if leftKind == reflect.Int32 {
			leftFloat = float64(left.(float32))
			rightFloat = float64(right.(float32))
		} else {
			leftFloat = left.(float64)
			rightFloat = right.(float64)
		}

		if leftFloat > rightFloat {
			return 1, nil
		} else if leftFloat < rightFloat {
			return -1, nil
		}

		return 0, nil
	} else if leftKind == reflect.String {
		if left.(string) > right.(string) {
			return 1, nil
		} else if left.(string) < right.(string) {
			return -1, nil
		}

		return 0, nil
	}

	return 0, errors.New("Unknown type")
}
