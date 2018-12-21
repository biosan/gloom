package gloom

import (
	"sync"
	"hash"
	"hash/fnv"
	"math/rand"
)

// Interface : Declare Bloom Filter basic operations
type Interface interface { 
	Add(item []byte)         // Add a byte array to the set
	Query(item []byte) bool  // Check if a byte array is probably into the set
}

// BTTL : Bitset Type Lenght (in bits)
const BTTL = 8

// BloomFilter : struct that contains Bloom Filter data
type BloomFilter struct { 
	// Constant/Build-time Parameters
	k uint				// Number of different hash functions
	m uint				// Size of the bit array 
	// Data generated at build-time
	seeds  	[]uint64	// The hash functions 'seeds'/'IVs'
	hashfn 	hash.Hash64	// A generic 64bit hash function
	wmutex	*sync.Mutex // Mutex is needed to prevent concurrent writes
	// Live data
	bitset 	[]byte		// The bloom-filter bit array 
	N 		uint		// Number of elements inside the filter
}

func (bf *BloomFilter) genSeedsUint64() []uint64 {
	seeds := make([]uint64, bf.k)

	for i := uint(0); i<bf.k; i++ {
		seeds[i] = rand.Uint64()
	}

	return seeds
}

// New : Returns a new BloomFilter object,
func New(k, m uint) *BloomFilter { 
  	bf := &BloomFilter {  
		bitset: make([]byte, m/BTTL),
		hashfn: fnv.New64a(),
		wmutex:	&sync.Mutex{},
		k: 		k,
		N: 		uint(0),
		m: 		m,
	}
	bf.seeds = bf.genSeedsUint64()

	return bf
}

// Add : Insert and element in the BloomFilter
func (bf *BloomFilter) Add(item []byte) { 
	hashes := bf.hashIt(item)

	for _, hash := range hashes {
		bf.wmutex.Lock()
		bf.bitset[hash/BTTL] |= byte(1 << (hash%BTTL))
		bf.wmutex.Unlock()
	}
	bf.N++
}

// Query : Returns 'true' if 'item' is probably in the BloomFilter
//		   Returns 'false' if 'item' is definitely not in the BloomFilter
func (bf *BloomFilter) Query(item []byte) bool {
	hashes := bf.hashIt(item)
	probablyIn := true

	for _, hash := range hashes {
		bit := bf.bitset[hash/BTTL] & (1 << (hash%BTTL))
		if bit != 0 {
			probablyIn = probablyIn && true
		} else {
			probablyIn = probablyIn && false
		}
	}

	return probablyIn
}

// Calculates all the hash values by applying in the item over the // hash functions
func (bf *BloomFilter) hashIt(item []byte) []uint64 { 
	out := make([]uint64, bf.k)
  
	for i, seed := range bf.seeds {  
	  bf.hashfn.Write(item)
	  out[i] = (seed ^ bf.hashfn.Sum64()) % uint64(bf.m)
	  bf.hashfn.Reset()
	}
  
	return out
}