package Go_array

import "fmt"

// func ต้องเป็นตัวใหญ่เพื่อ เเชร์ให้ ฟังชั่น หรือ ไฟล์อื่นใช่งานได้
func Array_basic() {
	var array1 []int = []int{1, 2, 3, 4}
	var array2 = []int{1, 2, 3, 4}

	for _, v := range array1 {
		fmt.Println(v)
	}

	for _, v := range array2 {
		fmt.Println(v)
	}
}

// Slice
func value_Slice() {
	var numebers1 = make([]int, 3, 5) // 3 len , 5 cap เกี่ยวกับmamayre
	numebers1 = append(numebers1, 1)  //1 ค่า
	numebers1 = append(numebers1, 2)  //2 ค่า
	showSlice(numebers1)

	var numebers2 []int
	numebers2 = append(numebers1, 1)
	numebers2 = append(numebers1, 2)
	showSlice(numebers2)
}

func showSlice(x []int) {
	fmt.Printf("len=%d cp=%d slice=%v\n", len(x), cap(x), x)
}

//----------------
// package Go_array

// import "fmt"

// func main() {
// 	var numbers = []int{1, 2, 3, 4, 5, 6}
// 	showSlice(numbers) //

// 	//leading remove
// 	numbers = numbers[1:len(numbers)] //]ลบข้างหน้า
// 	showSlice(numbers)

// 	//trailing remove
// 	numbers = numbers[0 : len(numbers)-1] //]ลบข้างหลัง ตั้ง:ถึง
// 	showSlice(numbers)
// }

// func showSlice(x []int) {
// 	fmt.Printf("len=%d cp=%d slice=%v\n", len(x), cap(x), x)
// }
