package div

import "testing"
import "fmt"


func TestDiv(t *testing.T) {
    div := Two(2, 2)
    expected := 1

    if div != expected {
        t.Errorf("expected '%d' but got '%d'", expected, div)
    }
}

func ExampleDiv() {
    div := Two(6, 2)
    fmt.Println(div)
    // Output: 3
}