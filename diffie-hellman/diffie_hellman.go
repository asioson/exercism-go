// Package diffiehellman implements routines to allow secure key exchange
package diffiehellman

import (
    "math/big"
    "crypto/rand"
)

// PrivateKey takes in a large prime number p and returns
// a private key
func PrivateKey(p *big.Int) *big.Int {
    max := big.NewInt(0).Sub(p, big.NewInt(2))
    a, _ := rand.Int(rand.Reader, max)
    a.Add(a, big.NewInt(2))
    return a
}


// PublicKey takes in private key a, prime number p, and g
// and returns g**a mod p as public key
func PublicKey(private, p *big.Int, g int64) *big.Int {
    A := big.NewInt(g)
    A.Exp(A, private, p)
    return A
}

// NewPair takes in prime number p and g and returns
// a pair of keys (private and public)
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
    a := PrivateKey(p)
    return a, PublicKey(a, p, g)
}

// SecretKey computes the secret key using the pair of keys
// (private and public) and prime number p
func SecretKey(private1, public2, p *big.Int) *big.Int {
    s := big.NewInt(0)
    s.Exp(public2, private1, p)
    return s
}

