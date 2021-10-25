package cpro

func IntContain(var1 int, var2 []int) bool {

	for _, a := range var2 {
		if a == var1 {
			return true
		}
	}
	return false

}
