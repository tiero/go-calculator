package sub

import "testing"
import "fmt"


func TestSub(t *testing.T) {
    sub := Sub(2, 2)
    expected := 0

    if sub != expected {
        t.Errorf("expected '%d' but got '%d'", expected, sub)
    }
}

func ExampleSub() {
    sub := Sub(5, 1)
    fmt.Println(sub)
    // Output: 4
}