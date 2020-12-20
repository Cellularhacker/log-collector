package uNum

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"strings"
)

func ParseCommaInt64(str string) (int64, error) {
	str = strings.ReplaceAll(str, ",", "")
	f, err := decimal.NewFromString(str)
	if err != nil {
		return -1, err
	}

	return f.IntPart(), nil
}

func ParseCommaFloat64(str string) (float64, error) {
	str = strings.ReplaceAll(str, ",", "")
	f, err := decimal.NewFromString(str)
	if err != nil {
		return -1, errors.New(fmt.Sprintf("decimal.NewFromString(str): %s", err))
	}
	f64, _ := f.Float64()
	return f64, nil
}
