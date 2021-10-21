package main
import "fmt"

func main() {
    parts := []string{"Hello", "my", "name", "is", "turki"}
    for i := len(parts) - 1; i >= 0; i-- {
        subParts := parts[i:]
        fmt.Println(subParts)
    }
}
