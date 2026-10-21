//go:build rsx

package main

// Crypto backend: Red-SecretsXAI (rsx). Build with: go build -tags rsx
// (requires the private rsx Go SDK; see go.mod replace directive / docs/BUILD.md).

import "redsecretsxai/rsx"

const cryptoBackend = "rsx"

func cbRandom(n int) ([]byte, error)                    { return rsx.RandomBytes(n) }
func cbHMAC(key, msg []byte) [32]byte                   { return rsx.HmacSha256(key, msg) }
func cbSeal(key, nonce, aad, pt []byte) ([]byte, error) { return rsx.Seal(key, nonce, aad, pt) }
func cbOpen(key, nonce, aad, ct []byte) ([]byte, error) { return rsx.Open(key, nonce, aad, ct) }
