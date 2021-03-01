package iaa_pipe_filter

import (
	"errors"
)

/**
 * @author  wjj
 * @date  2020/9/8 1:46 上午
 * @description
 */

var SumFilterWrongFormatError = errors.New("input data should be []int")

type SumFilter struct {
}

func NewSumFilter() *SumFilter {
	return &SumFilter{}
}
func (sf *SumFilter) Process(data Request) (Response, error) {
	//elms, ok := data.([]int)
	elms, ok := data.([]int)
	if !ok {
		return nil, SumFilterWrongFormatError
	}
	ret := 0
	for _, elem := range elms {
		ret += elem
	}
	return ret, nil
}
