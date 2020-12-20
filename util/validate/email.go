package validate

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

const emailLetters = "!#$%&'*+-/=?^_`{|}~abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.@_-"

func Email(str string) (bool, error) {
	for i, r := range []rune(str) {
		if !strings.ContainsRune(emailLetters, r) {
			return false, errors.New(fmt.Sprintf("invalid: email: %d: %v", i, r))
		}
	}

	return true, nil
}
