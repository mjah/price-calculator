package errors

import (
	"testing"
)

func TestErrorCode(t *testing.T) {
	testMessage := "Test Message"
	err := New(NegDenCappedFees, testMessage)

	if err.code != NegDenCappedFees {
		t.Errorf("Expected error code %d got %d.", NegDenCappedFees, err.code)
	}

	if err.message != testMessage {
		t.Errorf("Expected error message %s got %s.", testMessage, err.message)
	}
}
