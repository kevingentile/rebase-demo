package animal

import "fmt"

// Walker describes a type with legs
type Walker interface {
	Legs() int
}

// Describe generates a string about how a walker walks
func DescribeWalker(walker Walker) string {
	return fmt.Sprintf("Walks on %d legs.", walker.Legs())
}
