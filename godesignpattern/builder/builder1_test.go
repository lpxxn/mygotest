package builder

import (
	"testing"
	"fmt"
)

func TestBuilder(t *testing.T) {
	builder := New()
	car := builder.TopSpeed(100).Paint(BLUE).Build()

	fmt.Println(car.Drive())
	fmt.Println(car.Stop())
}