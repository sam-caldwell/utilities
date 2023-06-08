package utilities

func StringListToAnyList(list *[]string) []any {
	result := make([]any, len(*list))
	for i, v := range *list {
		result[i] = v
	}
	return result
}
