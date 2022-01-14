package funcs

func DeleteItem(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
