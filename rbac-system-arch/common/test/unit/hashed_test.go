package unit_test

import (
	"testing"
	"urmi-arch/common"
)

func Test(t *testing.T) {
	enc := common.Encrypt("MyJit")
	if enc == "" {
		t.Errorf("Test failed: %v", enc)
	}
	matcked, err := common.Decrypt(enc, "MyJit")
	if err != nil {
		t.Errorf("Test failed: got %v, want %v", err, nil)
	}
	if !matcked {
		t.Errorf("Test failed: ")
	}
}
