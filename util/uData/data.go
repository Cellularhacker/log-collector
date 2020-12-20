package uData

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"strings"
)

func ParseDataSize(str string) (int64, error) {
	b := decimal.New(1024, 0)
	d := decimal.New(1, 0)

	var err error
	var f decimal.Decimal
	// MARK: Safety
	r := strings.ReplaceAll(str, ",", "")
	r = strings.ReplaceAll(r, "\t", "")
	r = strings.ReplaceAll(r, " ", "")

	// MARK: Checking Overuse (ex. 3Mbps, 5Mbps... or 그냥 과금 등)
	if strings.Contains(r, "초과") {
		r = strings.ReplaceAll(r, "초과", "")
		// MARK: Safety for starting decimal with dot(.)
		if []rune(r)[0] == '.' {
			r = "0" + r
		}
		// MARK: Add '-' for indicate used over.
		r = "-" + r
	}

	if strings.Contains(r, "GB") {
		r = strings.ReplaceAll(r, "GB", "")
		d = b.Mul(b).Mul(b)
	} else if strings.Contains(r, "MB") {
		r = strings.ReplaceAll(r, "MB", "")
		d = b.Mul(b)
	} else if strings.Contains(r, "KB") {
		r = strings.ReplaceAll(r, "KB", "")
		d = b
	}
	f, err = decimal.NewFromString(r)
	if err != nil {
		return -1, errors.New(fmt.Sprintf("invalid: data: wrong format: %s: %s", str, err))
	}
	f = f.Mul(d)
	return f.IntPart(), nil
}

func ParseSmsCount(str string) (int64, error) {
	s := strings.ReplaceAll(str, "건", "")
	b, err := decimal.NewFromString(s)
	if err != nil {
		return -1, err
	}

	return b.IntPart(), nil
}
