package helpers

func MyMod(x, y int) (result int) {
	result = x % y
	if result < 0 {
		result += y
	}

	return result
}
