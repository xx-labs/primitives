////////////////////////////////////////////////////////////////////////////////////////////
// Copyright © 2020 xx network SEZC                                                       //
//                                                                                        //
// Use of this source code is governed by a license that can be found in the LICENSE file //
////////////////////////////////////////////////////////////////////////////////////////////

// An implementation of the leaky bucket algorithm:
// https://en.wikipedia.org/wiki/Leaky_bucket
package rateLimiting

import (
	"encoding/json"
	"sync"
	"time"
)

// Bucket structure tracks the capacity and rate at which the remaining buckets
// decrease.
type Bucket struct {
	capacity   uint32  // Maximum number of tokens the bucket can hold
	remaining  uint32  // Current number of tokens in the bucket
	leakRate   float64 // Rate that the bucket leaks tokens at [tokens/ns]
	lastUpdate int64   // Time that the bucket was most recently updated
	locked     bool    // When true, prevents bucket from being deleted when stale
	whitelist  bool    // When true, adding tokens always returns true
	sync.Mutex

	// Updates the remaining amount in database bucket. Leave value as nil if
	// the database is not being used.
	addToDb func(uint32, int64)
}

// CreateBucket generates a new empty bucket.
func CreateBucket(capacity, leaked uint32, leakDuration time.Duration,
	addToDb func(uint32, int64)) *Bucket {

	// Calculate the leak rate [tokens/nanosecond]
	leakRate := calculateLeakRate(leaked, leakDuration)

	return CreateBucketFromLeakRatio(capacity, leakRate, addToDb)
}

func calculateLeakRate(leaked uint32, leakedDuration time.Duration) float64 {
	return float64(leaked) / float64(leakedDuration.Nanoseconds())

}

// CreateBucketFromLeakRatio generates a new empty bucket.
func CreateBucketFromLeakRatio(capacity uint32, leakRate float64,
	addToDb func(uint32, int64)) *Bucket {
	return &Bucket{
		capacity:   capacity,
		remaining:  0,
		leakRate:   leakRate,
		lastUpdate: time.Now().UnixNano(),
		locked:     false,
		whitelist:  false,
		addToDb:    addToDb,
	}
}

// CreateBucketFromLeakRatio generates a new empty bucket.
func CreateBucketFromParams(params *BucketParams,
	addToDb func(uint32, int64)) *Bucket {
	return &Bucket{
		capacity:   params.Capacity,
		remaining:  params.Remaining,
		leakRate:   params.LeakRate,
		lastUpdate: params.LastUpdate,
		locked:     params.Locked,
		whitelist:  params.Whitelist,
		addToDb:    addToDb,
	}
}

// Capacity returns the max number of tokens allowed in the bucket.
func (b *Bucket) Capacity() uint32 {
	return b.capacity
}

// Remaining returns the number of tokens in the bucket.
func (b *Bucket) Remaining() uint32 {
	return b.remaining
}

// IsLocked returns true if the bucket is locked.
func (b *Bucket) IsLocked() bool {
	return b.locked
}

// IsWhitelisted returns true if the bucket is on the whitelist.
func (b *Bucket) IsWhitelisted() bool {
	return b.whitelist
}

// IsFull returns true if the bucket is overflowing (i.e. no remaining capacity
// for additional tokens) .
func (b *Bucket) IsFull() bool {
	b.Lock()
	defer b.Unlock()

	// Update the number of remaining tokens
	b.update(b.leakRate)

	return b.remaining >= b.capacity
}

// IsFullOrWhitelist returns true if the bucket is overflowing (i.e. no remaining capacity
// for additional tokens) or if the bucket is whitelisted.
func (b *Bucket) IsFullOrWhitelist() bool {
	b.Lock()
	defer b.Unlock()

	// Update the number of remaining tokens
	b.update(b.leakRate)

	return b.whitelist || b.remaining >= b.capacity
}

// IsEmpty returns true if the bucket is empty.
func (b *Bucket) IsEmpty() bool {
	b.Lock()
	defer b.Unlock()

	// Update the number of remaining tokens
	b.update(b.leakRate)

	return b.remaining == 0
}

// Add adds the specified number of tokens to the bucket. Returns true if the
// tokens were added; otherwise, returns false if there was insufficient
// capacity to do so.
func (b *Bucket) Add(tokens uint32) (bool, bool) {
	b.Lock()
	defer b.Unlock()

	// Update the number of remaining tokens in the bucket prior to adding
	b.update(b.leakRate)

	// Add the tokens to the bucket
	b.remaining += tokens

	// If using the database, then update the remaining in the database bucket
	if b.addToDb != nil {
		b.addToDb(b.remaining, b.lastUpdate)
	}

	// If the tokens went over capacity, then return false, unless the bucket is
	// whitelisted
	return b.whitelist || b.remaining <= b.capacity, b.whitelist
}

// AddWithExternalParams adds the specified number of tokens to the bucket given external
// bucket parameters rather than the params specified in the bucket.
// Returns true if the tokens were added; otherwise, returns false if there was insufficient
// capacity to do so.
func (b *Bucket) AddWithExternalParams(tokens, capacity, leakedTokens uint32,
	duration time.Duration) (bool, bool) {
	b.Lock()
	defer b.Unlock()

	// Update the number of remaining tokens in the bucket prior to adding
	b.update(calculateLeakRate(leakedTokens, duration))

	// Add the tokens to the bucket
	b.remaining += tokens

	// If using the database, then update the remaining in the database bucket
	if b.addToDb != nil {
		b.addToDb(b.remaining, b.lastUpdate)
	}

	// If the tokens went over capacity, then return false, unless the bucket is
	// whitelisted
	return b.whitelist || b.remaining <= capacity, b.whitelist
}

func (b *Bucket) AddWithoutOverflow(tokens uint32) (bool, bool) {
	b.Lock()
	defer b.Unlock()

	// Update the number of remaining tokens in the bucket prior to adding
	b.update(b.leakRate)

	addOK := b.remaining <= b.capacity

	// Add the tokens to the bucket
	b.remaining += tokens

	// If using the database, then update the remaining in the database bucket
	if b.addToDb != nil {
		b.addToDb(b.remaining, b.lastUpdate)
	}

	// If the tokens went over capacity, then return false, unless the bucket is
	// whitelisted
	return b.whitelist || addOK, b.whitelist
}

// update updates the number of remaining tokens in the bucket. It subtracts the
// number of leaked tokens since lastUpdate from the remaining number of tokens.
// This function is not thread safe. It must be called with a locked mutex.
func (b *Bucket) update(leakRate float64) {
	updateTime := time.Now().UnixNano()

	// Calculate the time elapsed since the last update, in nanoseconds
	elapsedTime := updateTime - b.lastUpdate

	// Calculate the number of tokens that have leaked over the elapsed time
	tokensLeaked := uint32(float64(elapsedTime) * leakRate)

	// Update the number of remaining tokens in the bucket
	if tokensLeaked > b.remaining {
		b.remaining = 0
	} else {
		b.remaining -= tokensLeaked
	}

	// Update timestamp
	b.lastUpdate = updateTime
}

// AddToDB isn't meaningfully serializable, so if necessary it should be
// populated after the fact
type bucketDisk struct {
	Capacity   uint32  // Maximum number of tokens the bucket can hold
	Remaining  uint32  // Current number of tokens in the bucket
	LeakRate   float64 // Rate that the bucket leaks tokens at [tokens/ns]
	LastUpdate int64   // Time that the bucket was most recently updated
	Locked     bool    // When true, prevents bucket from being deleted when stale
	Whitelist  bool    // When true, adding tokens always returns true
}

func (b *Bucket) MarshalJSON() ([]byte, error) {
	b.Lock()
	defer b.Unlock()
	return json.Marshal(&bucketDisk{
		Capacity:   b.capacity,
		Remaining:  b.remaining,
		LeakRate:   b.leakRate,
		LastUpdate: b.lastUpdate,
		Locked:     b.locked,
		Whitelist:  b.whitelist,
	})
}

// Problem: Doesn't include db func
func (b *Bucket) UnmarshalJSON(data []byte) error {
	var bd bucketDisk
	err := json.Unmarshal(data, &bd)
	if err != nil {
		return err
	}

	b.Lock()
	b.whitelist = bd.Whitelist
	b.locked = bd.Locked
	b.lastUpdate = bd.LastUpdate
	b.leakRate = bd.LeakRate
	b.remaining = bd.Remaining
	b.capacity = bd.Capacity
	b.Unlock()
	return nil
}

// This function should be called after unmarshalling if you need a db function
func (b *Bucket) SetAddToDB(dbFunc func(uint32, int64)) {
	b.Lock()
	b.addToDb = dbFunc
	b.Unlock()
}
