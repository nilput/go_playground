package main
import "fmt"

func main() {
    var i int = 42
    var f float32 = float32(i)
    var u uint = uint(f)
    fmt.Printf("var i int = 42 : %d\n"+
                "var f float32 = float32(i) : %g\n"+
                "var u uint = uint(f) : %d\n"+
                "i: %T\n"+
                "f: %T\n"+
                "u: %T\n", i, f, u, i, f, u);
}
