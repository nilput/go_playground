package main
import (
    "fmt"
)

/**
    [7, 3, 2]
*/
func insertionSort(arr []int) []int {
    for i := 0; i < len(arr); i++ {
        var val int = arr[i];
        var j int
        for j = i-1; j >= 0 && arr[j] >= val; j-- {
            arr[j+1] = arr[j];
        }
        arr[j+1] = val
    }
    return arr
}

func swap(arr []int, i int, j int) {
    tmp := arr[i]
    arr[i] = arr[j]
    arr[j] = tmp
}

func selectionSort(arr []int) []int {
    for i := 0; i < len(arr); i++ {
        min_idx := i
        for j := i+1; j < len(arr); j++ {
            if arr[j] < arr[min_idx] {
                min_idx = j;
            }
        }
        swap(arr, min_idx, i);
    }
    return arr;
}

func main() {
    arr := []int{13, 4, 7, 11, 3, 13}
    fmt.Println("Insertion: ", insertionSort(arr))
    fmt.Println("Selection: ", selectionSort(arr))
}
