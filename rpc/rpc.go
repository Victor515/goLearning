package rpcdemo

import (
	errors2 "errors"
)

// Service.Method

type DemoService struct {}

type Args struct {
	A, B int
}

func (DemoService) Div(args Args, result *float64) error{
	if args.B == 0{
		return  errors2.New("divide by zero")
	}

	*result = float64(args.A) / float64(args.B)
	return nil
}
