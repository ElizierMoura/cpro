package cpro

import "regexp"

func OnlyNumber(numeros string) string {
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	numbers := re.FindString(numeros)

	return numbers
}
