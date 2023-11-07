package pointer

import (
	"fmt"
)

func Pointer_Bic() {
	mag := "some message"
	var magPointer *string = &mag

	fmt.Println(mag)
	fmt.Println(*magPointer) //*คือการเข้าถึง value อ้างถึงตัวแปรที่เก็บ

	changeMessage(&mag, "NewMessage1")
	fmt.Println(mag)

	changeMessage(magPointer, "NewMessage2")
	fmt.Println(mag)
}

func changeMessage(aPointer *string, newMessage string) {
	*aPointer = newMessage
}

func Pointer_Bic1() {
	c := 5
	p := &c //& คืนค่าตำแหน่งในหน่วยความทรงจำ

	fmt.Println("Basic ->", c)
	fmt.Println("Pointer ->", p)
	*p = *p + 10
	fmt.Println("Sum Pointer ->", c)
}

func Pointer_Bic_toArray() {
	pa := &[3]int{0, 1, 2}
	ps := &[]string{"a", "b", "c"}
	pm := &map[string]bool{"foo": true, "bar": false}

	// เป็น Pointer ชี้ไปยังตำแหน่งหน่วยความจำของ ต่างๆตามลำดับ
	fmt.Println("pa is", pa)
	fmt.Println("ps is", ps)
	fmt.Println("pm is", pm)

	//ถ้าเราจะเข้าถึงข้อมูลที่ Pointer ชี้อยู่ ต้องใช้ * เพื่อคืนหน่วยความจำที่ array เเละ map เเล้วเราถึงจะเข้าถึงส่วนข้อได้ดังนี้
	fmt.Println("*pa is", *pa)
	fmt.Println("*ps is", *ps)
	fmt.Println("*pm is", *pm)
	fmt.Println("*pa[index -> 0] is", (*pa)[0])
	fmt.Println("*ps[index -> 0] is", (*ps)[0])
	fmt.Println("*pm[map] key foo", (*pm)["foo"])
}

func Pointer_Bic_vaule(p *[]int) string {
	return fmt.Sprintf("Length is %d", len(*p))
}

func Pointer_printvaule() {
	ps := &[]int{1, 2, 3}
	fmt.Println(Pointer_Bic_vaule(ps))
}
