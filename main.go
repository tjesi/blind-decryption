package main

import (
	"crypto/elliptic"
	"fmt"
)

// Ellipcic curve P256. Reference:
// https://golang.org/src/crypto/elliptic/elliptic.go
var curve = elliptic.P256()

/*
Public parameters:
P       *big.Int // the order of the underlying field
N       *big.Int // the order of the base point
*/
var params = curve.Params()

// B public constant setting
// randomness to 32 Bytes
const B = 32

func main() {

	// Generate private key k,
	// and public key K = (Kx,Ky).
	k, Kx, Ky := KeyGen()

	// Encode input message
	var InputMessage [B]byte
	copy(InputMessage[:], "{Tjerand Silde}")
	fmt.Println("Input: ", string(InputMessage[:]))

	// Encrypt the input message
  X, Y, Z := Encrypt(Kx, Ky, InputMessage)

	// Blind ciphertext
	Rx, Ry, r := Blind(X, Y)

	// Blind decryption of ciphertext
	Dx, Dy := BlindDecrypt(Rx, Ry, k)

	// Unblind and decrypt message
	OutputMessage := Unblind(Dx, Dy, Z, r)

	// Decode output message
	fmt.Println("Output: ", string(OutputMessage[:]))
}
