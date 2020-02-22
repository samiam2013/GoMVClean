package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"
)

func genRandHash(timeStamp int64) string {
	rand.Seed(time.Now().UnixNano())
	// make a random 64 bit number
	randN := rand.Int63n(timeStamp)
	randB := make([]byte, 8)
	binary.LittleEndian.PutUint64(randB, uint64(randN))
	// generate the byte form random hash
	hashN := sha256.Sum256(randB)
	// return string from the byte form
	return fmt.Sprintf("%x", hashN)
}
