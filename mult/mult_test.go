package mult

import "testing"
import "fmt"


func TestMult(t *testing.T) {
    mult := Two(2, 2)
    expected := 1

    if mult != expected {
        t.Errorf("expected '%d' but got '%d'", expected, mult)
    }
}

func ExampleMult() {
    mult := Two(6, 2)
    fmt.Println(mult)
    // Output: 3
}