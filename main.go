package main

import (
	"fmt"
	"homomorphic_encryption/fv12"
	"homomorphic_encryption/poly"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	ciphertext_modulus := int(math.Pow(2, 19))
	plaintext_modulus := int(math.Pow(2, 12))
	extra_mod := int(math.Pow(2, 40))
	poly_deg := int(math.Pow(2, 5) - 1)

	poly_mod := poly.ZeroPoly(int(poly_deg) + 1)
	poly_mod.SetElem(1, poly_deg+1)
	poly_mod.SetElem(1, 0)

	fv12_instance := fv12.InitMods(ciphertext_modulus, plaintext_modulus, extra_mod, poly_mod)

	plain_val1 := 29
	plain_val2 := 32

	pk, sk := fv12_instance.Keygen()

	mul_pk := fv12_instance.EvalKeygenMul(sk)

	ct1 := fv12_instance.Encrypt(pk, plain_val1)
	ct2 := fv12_instance.Encrypt(pk, plain_val2)
	val := fv12_instance.Decrypt(ct1, sk)

	add_ct := fv12_instance.AddCipher(ct1, ct2)
	added_val := fv12_instance.Decrypt(add_ct, sk)

	mul_ct := fv12_instance.MulCipher(ct1, ct2, mul_pk)
	mul_val := fv12_instance.Decrypt(mul_ct, sk)

	fmt.Println(added_val, (plain_val1+plain_val2)%plaintext_modulus)
	fmt.Println(mul_val, plain_val1*plain_val2)
	fmt.Println(val)

}
