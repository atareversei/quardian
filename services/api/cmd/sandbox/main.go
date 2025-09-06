package main

import (
	"errors"
	"fmt"

	"github.com/atareversei/quardian/services/api/pkg/richerror"
)

func main() {
	err := errors.New("bare minimum error")

	rerr1 := richerror.New("op0").WithError(err).WithMessage("rich error level 0")
	rerr2 := richerror.New("op1").WithError(rerr1).WithMessage("rich error level 1")
	fmt.Println(rerr2.GetMessage())
}
