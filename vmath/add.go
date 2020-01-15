package vmath

import "errors"

func Add(a int,b int)(ret int ,err error)  {
	if a < 0 ||b < 0 {
		err = errors.New(" both parameter should be non-negative numbers")
		return
	}

	return a+b,nil
}
