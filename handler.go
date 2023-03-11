package main

import (
	"context"
	"errors"
	"github.com/zzxwill/kitex-demo/kitex_gen/api"
)

// CalculatorImpl implements the last service interface defined in the IDL.
type CalculatorImpl struct{}

// GetMaximal implements the CalculatorImpl interface.
func (s *CalculatorImpl) GetMaximal(ctx context.Context, request *api.GetMaximalRequest) (resp *api.GetMaximalResponse, err error) {
	if len(request.Numbers) == 0 {
		err = errors.New("empty numbers")
		return
	}
	var max int32
	for i, num := range request.Numbers {
		if i == 0 {
			max = num
			continue
		}
		if num > max {
			max = num
		}
	}
	resp = &api.GetMaximalResponse{Result_: max}
	return
}
