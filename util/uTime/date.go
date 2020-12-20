package uTime

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"
)

func GetNowDate() int {
	loc, _ := time.LoadLocation("Asia/Seoul")
	year, month, day := time.Now().In(loc).Date()
	str := fmt.Sprintf("%4d%2d%2d", year, int(month), day)
	i, _ := strconv.Atoi(str)
	return i
}

func GetKSTDate(t *time.Time) int {
	if t == nil {
		tt := time.Now()
		t = &tt
	}
	year, month, day := t.In(loc).Date()
	str := fmt.Sprintf("%d%d%d", year, int(month), day)
	i, _ := strconv.Atoi(str)
	return i
}

func GetKSTDateStrBeautify(t *time.Time) string {
	if t == nil {
		tt := time.Now()
		t = &tt
	}
	loc, _ := time.LoadLocation("Asia/Seoul")
	year, month, day := t.In(loc).Date()
	hour, minute, second := t.In(loc).Clock()

	return fmt.Sprintf("%04d-%02d-%02d_%02d:%02d:%02d", year, int(month), day, hour, minute, second)
}

func ParseKSTISODate(str string) (*time.Time, error) {
	s := strings.Split(str, "-")
	if len(s) != 3 {
		return nil, errors.New(fmt.Sprintf("invalid: str: not ISO Date format: %s", str))
	}

	var err error
	is := make([]int, 3)
	for i, v := range s {
		is[i], err = strconv.Atoi(v)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("invalid: s[%d]: not integer: %s", i, v))
		}
	}

	t := time.Date(is[0], time.Month(is[1]), is[2], 0, 0, 0, 0, loc)

	return &t, nil
}

func GetKSTDateStr(t *time.Time) string {
	if t == nil {
		tt := time.Now()
		t = &tt
	}
	year, month, day := t.In(loc).Date()
	hour, minute, second := t.In(loc).Clock()

	return fmt.Sprintf("%04d%02d%02d-%02d%02d%02d", year, int(month), day, hour, minute, second)
}

func ParseBufTerm(str string) ([]int64, error) {
	s := strings.ReplaceAll(str, " ", "")
	s = strings.ReplaceAll(s, "조회기간", "")
	s = strings.ReplaceAll(s, ":", "")
	s = strings.ReplaceAll(s, "기준", "")
	s = strings.ReplaceAll(s, "(", "")
	s = strings.ReplaceAll(s, ")", "")
	s = strings.ReplaceAll(s, "부가세", "")
	s = strings.ReplaceAll(s, "별도", "")

	dates := strings.Split(s, "~")
	if len(dates) != 2 {
		return nil, errors.New(fmt.Sprintf("invalid: wrong bufTerm format: %s", str))
	}

	from, err := ParseKSTISODate(dates[0])
	if err != nil {
		return nil, errors.New(fmt.Sprintf("invalid: wrong bufTerm format: from: %s", dates[0]))
	}
	to, err := ParseKSTISODate(dates[1])
	if err != nil {
		return nil, errors.New(fmt.Sprintf("invalid: wrong bufTerm format: to: %s", dates[1]))
	}
	to2 := to.Add(24 * time.Hour).Add(-1 * time.Second)

	return []int64{from.Unix(), to2.Unix()}, nil
}

func ParseDotYearMonth(str string) (int, error) {
	s := strings.ReplaceAll(str, ".", "")
	s = strings.ReplaceAll(s, "-", "")
	return strconv.Atoi(s)
}
