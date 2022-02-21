package fv12

import (
	"homomorphic_encryption/poly"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestEncryptDecrypt(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	ciphertext_modulus := int(math.Pow(2, 19))
	plaintext_modulus := int(math.Pow(2, 12))
	extra_mod := int(math.Pow(2, 40))
	poly_deg := int(math.Pow(2, 4) - 1)

	poly_mod := poly.ZeroPoly(int(poly_deg) + 1)
	poly_mod.SetElem(1, poly_deg+1)
	poly_mod.SetElem(1, 0)

	fv12_instance := InitMods(ciphertext_modulus, plaintext_modulus, extra_mod,
		poly_mod)
	pk, sk := fv12_instance.Keygen()

	val_to_encrypt := 123
	ct1 := fv12_instance.Encrypt(pk, val_to_encrypt)
	decrypted_val := fv12_instance.Decrypt(ct1, sk)

	if decrypted_val != val_to_encrypt {
		t.Errorf("got %d, wanted %d", val_to_encrypt, decrypted_val)
	}
}

func TestAddCipher(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	ciphertext_modulus := int(math.Pow(2, 17))
	plaintext_modulus := int(math.Pow(2, 10))
	extra_mod := int(math.Pow(2, 40))
	poly_deg := int(math.Pow(2, 4) - 1)

	poly_mod := poly.ZeroPoly(int(poly_deg) + 1)
	poly_mod.SetElem(1, poly_deg+1)
	poly_mod.SetElem(1, 0)

	fv12_instance := InitMods(ciphertext_modulus, plaintext_modulus, extra_mod,
		poly_mod)
	pk, sk := fv12_instance.Keygen()

	val1 := 45
	val2 := 33
	ct1 := fv12_instance.Encrypt(pk, val1)
	ct2 := fv12_instance.Encrypt(pk, val2)
	added_ct := fv12_instance.AddCipher(ct1, ct2)
	decrypted_added_val := fv12_instance.Decrypt(added_ct, sk)

	if decrypted_added_val != (val1 + val2) {
		t.Errorf("got %d, wanted %d", decrypted_added_val, val1+val2)
	}
}
