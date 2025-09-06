package repoutil

func GetListLastPage(total int, pageSize int) int {
	if pageSize <= 1 {
		return 1
	}

	full := total / pageSize
	partial := total % pageSize

	if partial > 0 {
		full += 1
	}

	return full
}
