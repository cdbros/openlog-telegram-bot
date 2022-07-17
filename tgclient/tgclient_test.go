package tgclient

import "testing"

func TestCreateCommandResponse(t *testing.T) {
	response := createCommandResponse(GREET)
	if response != "Greetings" {
		t.Errorf("[Create Command Response] return error for command %s", GREET)
	}
}
