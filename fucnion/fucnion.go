package fucnion

func Sum(a, b int) int {
	return a + b
}

// retuer หลายค่า
func swap(a, b int) (int, int) {
	return b, a
}

// ์Named reture values
func swapV2(a, b int) (x, y int) {
	x = a
	y = b
	return
	// x, y := fucnion.swapV2(10, 20) 10 20 มารับค่า
	// f.Printf("%d , %d  => %d , %d \n ", 10, 20, x, y)
}
