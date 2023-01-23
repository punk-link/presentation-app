package templateFunctions

func ConvertToInt(source any) int {
	return int(source.(int32))
}
