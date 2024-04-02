package random

import (
  "encoding/binary"
  crand "crypto/rand"
  mrand "math/rand"
)

// https://zenn.dev/spiegel/articles/20211016-crypt-rand-as-a-math-rand

type Source struct{}

// Seed method is dummy function for rand.Source interface.
func (s Source) Seed(seed int64) {}

// Uint64 method generates a random number in the range [0, 1<<64).
func (s Source) Uint64() uint64 {
    b := [8]byte{}
    ct, _ := crand.Read(b[:])
    return binary.BigEndian.Uint64(b[:ct])
}

// Int63 method generates a random number in the range [0, 1<<63).
func (s Source) Int63() int64 {
    return (int64)(s.Uint64() >> 1)
}


func RandomInt(max int) int {
  return mrand.New(Source{}).Intn(max)
}

func RandomFloat64() float64 {
	return mrand.New(Source{}).Float64()
}
