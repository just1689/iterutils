package iterutils

func MapTo[T any, U any](input []T, mapper func(T) U) []U {
	output := make([]U, len(input))
	for i, elem := range input {
		output[i] = mapper(elem)
	}
	return output
}
