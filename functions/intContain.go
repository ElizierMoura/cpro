package cpro

// IntContain verifica se um int contém algum dos números citados no array.
func IntContain(numeroParaVerificar int, numerosNecessarios []int) bool {
	for _, numero := range numerosNecessarios {
		if numero == numeroParaVerificar {
			return true
		}
	}
	return false

}
