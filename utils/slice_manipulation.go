package utils

func TransposeStrings(rows []string) []string {
	cols := len(rows[0])
	out := make([]string, cols)

	for i := range cols {
		b := make([]byte, len(rows))
		for j := range len(rows) {
			b[j] = rows[j][i]
		}

		out[i] = string(b)
	}

	return out
}

func Transpose[T any](a [][]T) (result [][]T) {
	rows, cols := len(a), len(a[0])
	for range cols {
		result = append(result, make([]T, rows))
	}

	for i := range rows {
		for j := range cols {
			result[j][i] = a[i][j]
		}
	}

	return
}
