package main

func StringShortener(sentence string) []string {

	result := []string{}

	for _, v := range sentence {
		// jika tidak kosong dan huruf terakhir = huruf sekarang
		if len(result) > 0 && result[len(result)-1] == string(v) {
			// hapus yang terakhir
			result = result[:len(result)-1]
		} else {
			// tambah huruf ke result
			result = append(result, string(v))
			return result
		}

	}
	return []string{}

}
