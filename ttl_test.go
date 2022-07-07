package expire

import (
	"fmt"
	"testing"
	"time"
)

func ExampleWithTTL() {
	ttl := WithTTL("any_value", 2*time.Second)

	unwrap := ttl.Unwrap()
	if unwrap != nil {
		fmt.Println(*unwrap) // Output: any_value
	}
}

func TestTTL_Unwrap(t *testing.T) {
	t.Parallel()

	ttl := WithTTL("any_value", 2*time.Second)

	unwrap := ttl.Unwrap()
	if unwrap == nil {
		t.Fatal("expected not nil")
	}
	if *unwrap != "any_value" {
		t.Fatalf("expected any_value, got %q", *unwrap)
	}
}

func TestTTL_Unwrap_ttl(t *testing.T) {
	t.Parallel()

	ttl := WithTTL("any_value", 1*time.Second)

	<-time.After(2 * time.Second)
	unwrap := ttl.Unwrap()
	if unwrap != nil {
		t.Fatal("expected nil")
	}
}
