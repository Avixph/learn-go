package main

import "fmt"

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// write your error code here
		err := fmt.Errorf("need more coffee! \nvalue was: %v", f)
		return 0, sqrtError{
			lat:  "50.2289 N",
			long: "99.4656 W",
			err:  err,
		}
	}
	return 42, nil
}
