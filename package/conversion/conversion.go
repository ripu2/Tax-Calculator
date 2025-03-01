package conversion

import (
	"errors"
	"strconv"
)

func StringToFloat(str string) (float64, error) {

	floatVal, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, errors.New(err.Error())
	}
	return floatVal, nil
}
