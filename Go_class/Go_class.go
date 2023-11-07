// ภาษา go ไม่มี class
package goclass

import "fmt"

type product struct {
	name  string
	price int
	stock int
}

func Print_value() {
	var p1 product
	p1.name = "Arduing"
	p1.price = 100
	p1.stock = 200
	fmt.Println(p1)      // {Arduing  100 200}
	fmt.Println(p1.name) // print แบบระบุตัว

}
