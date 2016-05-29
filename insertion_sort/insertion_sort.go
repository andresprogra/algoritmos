package main
import "fmt"
func main() {
	arreglo_desordenado := []int{123,435,986,456,321,765,7,6,1,900}
	fmt.Println(arreglo_desordenado)
	arreglo_ordenado := insertion_sort(arreglo_desordenado)
	fmt.Println(arreglo_ordenado)
}

func insertion_sort(arreglo []int) []int {
	for i := 1; i<len(arreglo); i++ {
		j := i
		for j > 0 && arreglo[j-1] > arreglo[j] {
			swap(j-1,j,&arreglo)
			j--;
		}
	}
	return arreglo
}
func swap(previo int, actual int, puntero_arreglo *[]int) {
	arreglo := *puntero_arreglo
	copia := arreglo[actual]
	arreglo[actual] = arreglo[previo]
	arreglo[previo] = copia
}