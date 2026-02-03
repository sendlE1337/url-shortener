package shortener

import "strings"

//shortener генерирует короткий идентификатор из числа
const alphabet = "5GgDq8LAaPpHhTtYyC56xZzNn3iRrWwUuKkB7oFfSsJjMm2cVvEe1"

var alphabetLen = uint32(len(alphabet))

func Shortener(id uint32) string {
	var (
		nums    []uint32
		num     = id
		builder strings.Builder
	)

	for num > 0 {
		nums = append(nums, num%alphabetLen)
		num /= alphabetLen
	}

	//Этот шаг нужен всегда, когда конвертируем число в любую другую систему счисления вручную, потому что остатки берутся с конца, а выводить нужно с начала.
	Reverse(nums)

	for _, num := range nums {
		builder.WriteString(string(alphabet[num]))
	}

	return builder.String()
}

//Reverse переворачивает слайс чисел.
func Reverse(s []uint32) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
