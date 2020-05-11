package Stream

import (
	"strconv"
	"fmt"
)

func stringToFloat(num string) float64 {
	flt, err := strconv.ParseFloat(num, 64)
	if err != nil {
		fmt.Printf("error converting string to float64 %v",  err)
	}
		
	return flt
} 

func marketMakerSide(isBuyer bool) string {
	if isBuyer {
		return "buy"
	} else {
		return "sell"
	}
}