package multiplier
import "testing"

func TestMultiplicate(test *testing.T) {
    var result int = Dup(5)
    if result != 10 {
        test.Error("Expected 10, got: ", result)
    }    
}
