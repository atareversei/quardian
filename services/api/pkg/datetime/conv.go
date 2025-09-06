package datetime

import (
	"errors"
	"fmt"
	"time"

	"github.com/atareversei/quardian/services/api/internal/entity/commonentity"
)

var ErrMalformedDateTimeString = errors.New("the datetime string is malformed")

func ToStdDateTime(inputPtr *time.Time) commonentity.DateTime {
	datetime := commonentity.DateTime{}

	if inputPtr == nil {
		return datetime
	}

	Y, M, D := inputPtr.Date()
	date := fmt.Sprintf("%d-%s-%s", Y, padZero(int(M)), padZero(D))
	c, m, s := inputPtr.Clock()
	clock := fmt.Sprintf("%s:%s:%s", padZero(c), padZero(m), padZero(s))

	datetime.Date = &date
	datetime.Time = &clock

	return datetime
}

func padZero(num int) string {
	if num < 10 {
		return fmt.Sprintf("0%d", num)
	}
	return fmt.Sprintf("%d", num)
}
