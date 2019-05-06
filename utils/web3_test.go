package utils

import (
	"fmt"
	"testing"
)

func TestPrivateKeyToAddress(t *testing.T) {
	privateKey := "0xe4abcbf75d38cf61c4fde0ade1148f90376616f5233b7c1fef2a78c5992a9a50"
	address := PrivateKeyToAddress(privateKey)
	fmt.Printf("%+v", address)
}

func TestSignString(t *testing.T) {
	sig := SignString("0xe4abcbf75d38cf61c4fde0ade1148f90376616f5233b7c1fef2a78c5992a9a50", "HYDRO-AUTHENTICATION@1524088776656")
	fmt.Printf("%+v", sig)
}
