package animal

import "fmt"

// Swimmer describes an animal with fins
type Swimmer interface {
	Fins() int
}

// Describe generates a string about how a swimmer swims
func DescribeSwimmer(swimmer Swimmer) string {
	return fmt.Sprintf("swims with %d fins", swimmer.Fins())
}
