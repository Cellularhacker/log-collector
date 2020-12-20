package uTime

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"
)

var loc *time.Location

func init() {
	loc, _ = time.LoadLocation("Asia/Seoul")
}

func GetLoc() *time.Location {
	return loc
}

func GetKST(t *time.Time) time.Time {
	if t == nil {
		return time.Now().In(loc)
	}

	return t.In(loc)
}

func ParseKrTime(str string) (int64, error) {
	var err error
	// Safety
	after := strings.ReplaceAll(str, " ", "")
	after = strings.ReplaceAll(after, "\t", "")
	after = strings.ReplaceAll(after, "\n", "")
	after = strings.ReplaceAll(after, "\r", "")

	h := 0
	m := 0
	s := 0

	if strings.Contains(after, "시간") {
		hStr := strings.Split(after, "시간")
		if len(hStr) > 2 {
			return -1, errors.New(fmt.Sprintf("invalid: format: not correct format: %s", str))
		}
		h, err = strconv.Atoi(hStr[0])
		if err != nil {
			return -1, errors.New(fmt.Sprintf("invalid: hour: not integer: %s", hStr[0]))
		}
		if len(hStr) > 1 {
			after = hStr[1]
		}
	}

	if strings.Contains(after, "분") {
		mStr := strings.Split(after, "분")
		if len(mStr) > 2 {
			return -1, errors.New(fmt.Sprintf("invalid: format: not correct format: %s", str))
		}
		m, err = strconv.Atoi(mStr[0])
		if err != nil {
			return -1, errors.New(fmt.Sprintf("invalid: minute: not integer: %s", mStr[0]))
		}
		if len(mStr) > 1 {
			after = mStr[1]
		}
	}

	if strings.Contains(after, "초") {
		sStr := strings.Split(after, "초")
		if len(sStr) > 2 {
			return -1, errors.New(fmt.Sprintf("invalid: format: not correct format: %s", str))
		}
		s, err = strconv.Atoi(sStr[0])
		if err != nil {
			return -1, errors.New(fmt.Sprintf("invalid: second: not integer: %s", sStr[0]))
		}
	}

	return int64(h*3600 + m*60 + s), nil
}
