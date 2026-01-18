package main

import "strconv"

func SquareService(nStr string) (int, error) {
	n, err := strconv.Atoi(nStr)
	if err != nil {
		return 0, BizError{Code: 10001, Msg: "invalid parameter: n"}
	}
	return n * n, nil
}

const DemoToken = "abc123"
