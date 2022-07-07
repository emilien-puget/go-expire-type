package expire

import (
	"fmt"
	"testing"
	"time"
)

func ExampleWithDeadline() {
	deadline := WithDeadline("any_value", time.Now().Add(1*time.Minute))

	unwrap := deadline.Unwrap()
	if unwrap != nil {
		fmt.Println(*unwrap) // Output: any_value
	}
}

func TestDeadline_Unwrap(t *testing.T) {
	t.Parallel()

	deadline := WithDeadline("any_value", time.Now().Add(1*time.Minute))

	unwrap := deadline.Unwrap()
	if unwrap == nil {
		t.Fatal("expected not nil")
	}
	if *unwrap != "any_value" {
		t.Fatalf("expected any_value, got %q", *unwrap)
	}
}

func TestDeadline_Unwrap_ttl(t *testing.T) {
	t.Parallel()

	deadline := WithDeadline("any_value", time.Now().Add(-1))

	<-time.After(2 * time.Second)
	unwrap := deadline.Unwrap()
	if unwrap != nil {
		t.Fatal("expected nil")
	}
}
