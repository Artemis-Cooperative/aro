package packageA

import (
	"testing"

	B "github.com/Artemis-Cooperation/arrow/packageB"
)

func TestHello(t *testing.T) {
	B.greet()
}