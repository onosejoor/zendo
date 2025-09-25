package utils_test

import (
	"main/utils"
	"testing"
)

func TestHexToObjectID(t *testing.T) {
	hex := "507f1f77bcf86cd799439011"
	objID := utils.HexToObjectID(hex)

	if objID.Hex() != hex {
		t.Errorf("expected %s, got %s", hex, objID.Hex())
	}
}
