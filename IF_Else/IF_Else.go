package ifelse

import (
	"fmt"
)

func IFv1() {
	number := 4
	if number < 10 {
		fmt.Println("T")
	} else {
		fmt.Println("F")
	}
}

func IFv2() {
	if result := dosomethig(); result == "OK" {
		fmt.Println("OK")
	} else {
		fmt.Println("No OK")
	}
}

func dosomethig() string {
	return "OK"
}

// switch
func IFv3() {
	index := 2
	switch index {
	case 0:
		fmt.Println("0")
		break
	case 1:
		fmt.Println("1")
		break
	case 2:
		fmt.Println("2")
		break
	default:
		fmt.Println(" do some thig")
		break
	}
}
