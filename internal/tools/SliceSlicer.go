package tools


func SliceSlicer(input [][]string) []string{
	var result = []string{}
	for _, arr := range input {
		for _, item := range arr {
			result = append(result, item)
		}
	}

	return result
}
