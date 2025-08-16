package calculator

// Function TambahAngka akan bisa diakses oleh dari package lain
// Exported
func TambahAngka(a, b int) int {
	return a + b
}

func KurangAngka(a, b int) int {
	return a - b
}
