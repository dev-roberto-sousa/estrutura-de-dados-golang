package sorting

func InsertionSort(A []int) {
	// for j := range A
	// Funcionaria, porém, tornaria a iteração redundante
	// porque range A começa com 0 e em insertion sort
	// o primeiro elemento aqui já é visto como ordenado
	for j := 1; j < len(A); j++ {
		key := A[j]
		i := j - 1
		for i >= 0 && A[i] > key {
			A[i+1] = A[i]
			i--
		}
		A[i+1] = key
	}

}
