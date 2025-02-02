package goloadgen

import (
	"testing"
)

// Mock function for rand.Intn
func mockRandIntn(n int) int {
	return 2 // return a fixed value for testing
}

func TestSimpleGeneratePayload(t *testing.T){
	payload := GeneratePayload("{user_id: 1}")
	if payload != "{user_id: 1}" {
		t.Errorf("Expected %v. Acual: %v", "{user_id: 1}", payload)
	}
}

func TestGeneratePayloadWithRandInt(t *testing.T){
	// Save the original rand.Intn function
	originalRandIntn := randIntn
	// Replace it with the mock function
	randIntn = mockRandIntn
	// Restore the original function after the test
	defer func() { randIntn = originalRandIntn }()

	payload := GeneratePayload("{ user_id: {{randInt 5}} }")
	if payload != "{ user_id: 2 }" {
		t.Errorf("Expected %v. Actual: %v", "{ user_id: 2 }", payload)
	}
}