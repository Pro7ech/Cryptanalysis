package linear

import (
	"fmt"
	"testing"
)

func printDebug(x [8]uint8) {
	for i := 0; i < 8; i++ {
		fmt.Printf("%02x ", x[i])
	}
	fmt.Printf("\n")
}

func equalSlice(x, y [8]uint8) (v bool) {
	v = true
	for i := 0; i < 8; i++ {
		v = v && (x[i] == y[i])
	}
	return
}

func Test_SPN64(t *testing.T) {

	t.Run(fmt.Sprintf("Encrpyt/Decrypt"), func(t *testing.T) {

		var plaintext, key, ciphertext [8]uint8

		plaintext = [8]uint8{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
		key = [8]uint8{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}

		keys := keySchedule(key)

		Encrypt(&plaintext, keys, &ciphertext)

		if !equalSlice(ciphertext, [8]uint8{0x0c, 0x3d, 0x14, 0x86, 0x99, 0x86, 0xb6, 0xa5}) {
			t.Errorf("Encrypt")
		}

		Decrypt(&ciphertext, keys, &plaintext)

		if !equalSlice(plaintext, [8]uint8{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}) {
			t.Errorf("Decrypt")
		}

	})
}

func Test_RecoverKey(t *testing.T) {

	if !RecoverKey([8]uint8{0xf8, 0x9e, 0xa8, 0xc4, 0x09, 0xde, 0x89, 0xbb}) {
		t.Errorf("Bad Recover Key")
	}

}
