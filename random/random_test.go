package random_test

import (
	"crypto/rand"
	"math/big"
	mRand "math/rand"
	"testing"
	"time"

	"github.com/asabya/go-utils/random"
)

var (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func randomStringMathRand(length int) string {
	mRand.Seed(time.Now().UnixNano())
	str := make([]byte, length)
	for i := range str {
		idx := mRand.Intn(len(charset))
		str[i] = charset[idx]
	}
	return string(str)
}

func randomStringCryptoRand(n int) (string, error) {
	b := make([]byte, n)
	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		b[i] = charset[num.Int64()]
	}
	return string(b), nil
}

func randomBytesMathRand(length int) []byte {
	mRand.Seed(time.Now().UnixNano())
	str := make([]byte, length)
	for i := range str {
		idx := mRand.Intn(len(charset))
		str[i] = charset[idx]
	}
	return str
}

func randomBytesCryptoRand(n int) ([]byte, error) {
	b := make([]byte, n)
	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return nil, err
		}
		b[i] = charset[num.Int64()]
	}
	return b, nil
}

func BenchmarkRandomHex(b *testing.B) {
	b.Run("math_rand", func(b *testing.B) {
		for i := 0; i < 1000; i++ {
			randomStringMathRand(1000)
		}
	})
	b.Run("crypto_rand", func(b *testing.B) {
		for i := 0; i < 1000; i++ {
			randomStringCryptoRand(1000)
		}
	})
	b.Run("random_hex", func(b *testing.B) {
		for i := 0; i < 1000; i++ {
			random.Hex(1000)
		}
	})
}

func BenchmarkRandomBytes(b *testing.B) {
	b.Run("math_rand", func(b *testing.B) {
		for i := 0; i < 1000; i++ {
			randomBytesMathRand(1000)
		}
	})
	b.Run("crypto_rand", func(b *testing.B) {
		for i := 0; i < 1000; i++ {
			randomBytesCryptoRand(1000)
		}
	})
	b.Run("random_hex", func(b *testing.B) {
		for i := 0; i < 1000; i++ {
			random.Bytes(1000)
		}
	})
}
