package loop

import "fmt"

func Loopfor() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("value = %d \n", i)
	}
}

func Loopwhile() {
	i := 0
	for i <= 5 {
		fmt.Printf("value = %d \n", i)
		i++
	}
}

func LoopforEage() {
	courses := []string{"a", "b", "c"}

	for index, item := range courses {
		fmt.Printf("%d. %s\n", index+1, item)
	}

	//อีกเเบบ forr _, เหมือน gitignore
	for _, v := range courses {
		fmt.Printf("%s\n", v)
	}
}

func Loopforbreak() {
	i := 0
	for i <= 10 {
		fmt.Printf("value = %d \n", i)
		if i == 8 {
			break
		}
		i++
	}
}

// ใน golang ไม่มี whileloop
