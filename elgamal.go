package main

import (
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

// RandomBytes sample B random Bytes.
// Should do rejection sampling here.
func RandomBytes() []byte {
	r := make([]byte, B)
	rand.Read(r)
	return r
}

// KeyGen creates a private key k,
// and a public key K = [k]*G.
func KeyGen() (k []byte, Kx, Ky *big.Int) {
	k = RandomBytes()
	Kx, Ky = curve.ScalarBaseMult(k[:])
	return
}

// Encrypt creates a ciphertext (C = (X,Y),Z) of
// the form ([s]*G, H([s]*K) XOR M) for random s.
func Encrypt(Kx, Ky *big.Int, M [B]byte) (X, Y *big.Int, Z [B]byte) {
	s := RandomBytes()
	X, Y = curve.ScalarBaseMult(s[:])
	Tx, Ty := curve.ScalarMult(Kx, Ky, s[:])
	bytes := append(Tx.Bytes(), Ty.Bytes()...)
	Z = XORBytes(sha256.Sum256(bytes), M, B)
	return
}

// XORBytes compute A XOR B
func XORBytes(a, b [B]byte, n int) (dst [B]byte) {
	for i := 0; i < n; i++ {
		dst[i] = a[i] ^ b[i]
	}
	return
}

// Blind randomise C = (X,Y) into R = [r]*C.
func Blind(X, Y *big.Int) (Rx, Ry *big.Int, r []byte) {
	r = RandomBytes()
	Rx, Ry = curve.ScalarMult(X, Y, r[:])
	return
}

// BlindDecrypt blindly decrypt as D = [k]*R.
func BlindDecrypt(Rx, Ry *big.Int, k []byte) (Dx, Dy *big.Int) {
	Dx, Dy = curve.ScalarMult(Rx, Ry, k[:])
	return
}

// Unblind remove blinding and decrypt as
// M = Z XOR H([rInv]*D)
func Unblind(Dx, Dy *big.Int, Z [B]byte, r []byte) (N [B]byte) {
	rInv := new(big.Int).SetBytes(r[:])
	rInv.ModInverse(rInv, params.N)
	Tx, Ty := curve.ScalarMult(Dx, Dy, rInv.Bytes())

	bytes := append(Tx.Bytes(), Ty.Bytes()...)
	N = XORBytes(sha256.Sum256(bytes), Z, B)
	return
}
