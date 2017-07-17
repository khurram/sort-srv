package handler

import (
	proto "github.com/mistrq/sort-srv/proto"
	"golang.org/x/net/context"
)

type Sort struct{}

func (s *Sort) InsertionSort(ctx context.Context, req *proto.InputArray, rsp *proto.OutputArray) error {
	input := req.Input
	for j := 1; j < len(input); j++ {
		key := input[j]
		i := j - 1
		for i >= 0 && input[i] > key {
			input[i+1] = input[i]
			i--
		}
		input[i+1] = key
	}
	rsp.Output = input
	return nil
}

func (s *Sort) SelectionSort(ctx context.Context, req *proto.InputArray, rsp *proto.OutputArray) error {
	input := req.Input
	for j := 0; j < len(input)-1; j++ {
		min := j
		for i := j + 1; i < len(input); i++ {
			if input[i] < input[min] {
				min = i
			}
		}

		if min != j {
			tmp := input[j]
			input[j] = input[min]
			input[min] = tmp
		}
	}
	rsp.Output = input
	return nil
}
