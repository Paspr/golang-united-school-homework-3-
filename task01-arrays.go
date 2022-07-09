package homework

func average(input [15]float32) (result float32) {
	var sum float32 = 0.0
	for _, element := range input {
		sum = sum + element
	}
	result = sum / float32(len(input))
	return
}
