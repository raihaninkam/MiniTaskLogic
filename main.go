package main

func StringShortener(sentence string) []string {

	result := []rune{}

	for _, v := range sentence {
		// jika tidak kosong dan huruf terakhir = huruf sekarang
		if len(result) > 0 && result[len(result)-1] == v {
			// hapus yang terakhir
			result = result[:len(result)-1]
		} else {
			// tambah huruf ke result
			result = append(result, v)

		}

	}
	return []string{string(result)}
}
