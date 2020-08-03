////////////////////////////////////////////////////////////////////////////////////////////
// Copyright © 2020 xx network SEZC                                                       //
//                                                                                        //
// Use of this source code is governed by a license that can be found in the LICENSE file //
////////////////////////////////////////////////////////////////////////////////////////////

package id

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

// Tests that Bytes() returns the correct byte slice of an ID and that it is
// a copy.
func TestID_Bytes(t *testing.T) {
	// Test values
	expectedBytes := newRandomBytes(ArrIDLen, t)
	testID := NewIdFromBytes(expectedBytes, t)

	// Check for the correct values
	testVal := testID.Bytes()
	if !bytes.Equal(expectedBytes, testVal) {
		t.Errorf("Bytes() returned the incorrect byte slice of the ID"+
			"\n\texpected: %+v\n\treceived: %+v", expectedBytes, testVal)
	}

	// Test if the returned bytes are copies
	if &testID[0] == &testVal[0] {
		t.Errorf("Bytes() did not return a copy when it should have."+
			"\n\texpected: any value except %+v\n\treceived: %+v",
			&testID[0], &testVal[0])
	}
}

// Tests that Cmp() returns the correct value when comparing equal and unequal
// IDs.
func TestID_Cmp(t *testing.T) {
	// Test values
	randomBytes1 := newRandomBytes(ArrIDLen, t)
	randomBytes2 := newRandomBytes(ArrIDLen, t)
	testID1 := NewIdFromBytes(randomBytes1, t)
	testID2 := NewIdFromBytes(randomBytes1, t)
	testID3 := NewIdFromBytes(randomBytes2, t)

	// Compare two equal IDs
	testVal := testID1.Cmp(testID2)
	if !testVal {
		t.Errorf("Cmp() incorrectly determined the two IDs are not equal."+
			"\n\texpected: %+v\n\treceived: %+v", true, testVal)
	}

	// Compare two unequal IDs
	testVal = testID1.Cmp(testID3)
	if testVal {
		t.Errorf("Cmp() incorrectly determined the two IDs are equal."+
			"\n\texpected: %+v\n\treceived: %+v", false, testVal)
	}
}

// Test that DeepCopy() returns a copy with the same contents as the original
// and where the pointers are different.
func TestID_DeepCopy(t *testing.T) {
	// Test values
	expectedID := NewIdFromBytes(newRandomBytes(ArrIDLen, t), t)

	// Test if the contents are equal
	testVal := expectedID.DeepCopy()
	if !reflect.DeepEqual(expectedID, testVal) {
		t.Errorf("DeepCopy() returned a copy with the wrong contents."+
			"\n\texpected: %+v\n\treceived: %+v", expectedID, testVal)
	}

	// Test if the returned bytes are copies
	if &expectedID[0] == &testVal[0] {
		t.Errorf("DeepCopy() did not return a copy when it should have."+
			"\n\texpected: any value except %+v\n\treceived: %+v",
			&expectedID[0], &testVal[0])
	}
}

// Tests that the base64 encoded string returned by String() can be decoded into
// the original ID.
func TestID_String(t *testing.T) {
	// Expected values
	expectedBytes := newRandomBytes(ArrIDLen, t)
	expectedID := NewIdFromBytes(expectedBytes, t)

	// Encode into string
	stringID := expectedID.String()

	// Decode the string and check
	newID, err := base64.StdEncoding.DecodeString(stringID)
	if err != nil {
		t.Fatalf("Failed to decode string returned by String():\n%v", err)
	}

	if !bytes.Equal(expectedBytes, newID) {
		t.Errorf("String() did not encode the string correctly."+
			"The decoded strings differ.\n\texpected: %v\n\treceived: %v",
			expectedBytes, newID)
	}
}

// Tests that GetType() returns the correct type for each ID type.
func TestID_GetType(t *testing.T) {
	// Test values
	testTypes := []Type{Generic, Gateway, Node, User, NumTypes, 6}
	randomBytes := [][]byte{
		newRandomBytes(ArrIDLen-1, t), newRandomBytes(ArrIDLen-1, t),
		newRandomBytes(ArrIDLen-1, t), newRandomBytes(ArrIDLen-1, t),
		newRandomBytes(ArrIDLen-1, t), newRandomBytes(ArrIDLen-1, t),
	}
	testIDs := []*ID{
		NewIdFromBytes(append(randomBytes[0], byte(testTypes[0])), t),
		NewIdFromBytes(append(randomBytes[1], byte(testTypes[1])), t),
		NewIdFromBytes(append(randomBytes[2], byte(testTypes[2])), t),
		NewIdFromBytes(append(randomBytes[3], byte(testTypes[3])), t),
		NewIdFromBytes(append(randomBytes[4], byte(testTypes[4])), t),
		NewIdFromBytes(append(randomBytes[5], byte(testTypes[5])), t),
	}

	for i, testID := range testIDs {
		testVal := testID.GetType()
		if testTypes[i] != testVal {
			t.Errorf("GetType() returned the incorrect type."+
				"\n\texpected: %v\n\treceived: %v", testTypes[i], testVal)
		}
	}
}

// Tests that SetType() sets the type of the ID correctly by checking if the
// ID's type changed after calling SetType().
func TestID_SetType(t *testing.T) {
	// Test values
	expectedType := Node
	testType := Generic
	testBytes := newRandomBytes(dataLen, t)
	testID := NewIdFromBytes(append(testBytes, byte(testType)), t)

	// Change the ID
	testID.SetType(expectedType)

	// Check the ID
	testVal := testID.GetType()
	if expectedType != testVal {
		t.Errorf("SetType() did not set the ID type correctly."+
			"\n\texpected: %v\n\treceived: %v", expectedType, testVal)
	}
}

// Tests that NewIdFromBytes() creates a new ID with the correct contents.
func TestNewIdFromBytes(t *testing.T) {
	// Expected values
	expectedBytes := newRandomBytes(ArrIDLen, t)

	// Create the ID and check its contents
	newID := NewIdFromBytes(expectedBytes, t)

	if !bytes.Equal(expectedBytes, newID[:]) {
		t.Errorf("NewIdFromBytes() produced an ID with the incorrect bytes."+
			"\n\texpected: %v\n\treceived: %v", expectedBytes, newID[:])
	}
}

// Tests that NewIdFromBytes() creates a new ID from bytes with a length smaller
// than 33. The resulting ID should have the bytes and the rest should be 0.
func TestNewIdFromBytes_Underflow(t *testing.T) {
	// Expected values
	expectedBytes := newRandomBytes(ArrIDLen/2, t)
	expectedArr := [ArrIDLen]byte{}
	copy(expectedArr[:], expectedBytes)

	// Create the ID and check its contents
	newID := NewIdFromBytes(expectedBytes, t)

	if !bytes.Equal(expectedArr[:], newID[:]) {
		t.Errorf("NewIdFromBytes() produced an ID with the incorrect bytes."+
			"\n\texpected: %v\n\treceived: %v", expectedArr, newID[:])
	}
}

// Tests that NewIdFromBytes() creates a new ID from bytes with a length larger
// than 33. The resulting ID should the original bytes truncated to 33 bytes.
func TestNewIdFromBytes_Overflow(t *testing.T) {
	// Expected values
	expectedBytes := newRandomBytes(ArrIDLen*2, t)
	expectedArr := [ArrIDLen]byte{}
	copy(expectedArr[:], expectedBytes)

	// Create the ID and check its contents
	newID := NewIdFromBytes(expectedBytes, t)

	if !bytes.Equal(expectedArr[:], newID[:]) {
		t.Errorf("NewIdFromBytes() produced an ID with the incorrect bytes."+
			"\n\texpected: %v\n\treceived: %v", expectedArr, newID[:])
	}
}

// Tests that NewIdFromBytes() panics when given a nil testing object.
func TestNewIdFromBytes_TestError(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("NewIdFromBytes() did not panic when it received a " +
				"nil testing object when it should have.")
		}
	}()

	// Call function with nil testing object
	_ = NewIdFromBytes(newRandomBytes(ArrIDLen, t), nil)
}

// Tests that NewIdFromString() creates a new ID from string correctly. The new
// ID is created from a string that is 32 bytes long so that no truncation or
// padding is required. The test checks that the original string is still
// present in the data.
func TestNewIdFromString(t *testing.T) {
	// Test values
	expectedIdString := "Test ID string of correct length"
	expectedType := Generic
	expectedID := new(ID)
	copy(expectedID[:], append([]byte(expectedIdString), byte(expectedType)))

	// Create the ID and check its contents
	newID := NewIdFromString(expectedIdString, expectedType, t)

	// Check if the new ID matches the expected ID
	if !expectedID.Cmp(newID) {
		t.Errorf("NewIdFromString() produced an ID with the incorrect data."+
			"\n\texpected: %v\n\treceived: %v", expectedID[:], newID[:])
	}

	// Check if the original string is still in the first 32 bytes
	newIdString := string(newID.Bytes()[:ArrIDLen-1])
	if expectedIdString != newIdString {
		t.Errorf("NewIdFromString() did not correctly convert the original "+
			"string to bytes.\n\texpected string: %#v\n\treceived string: %#v"+
			"\n\texpected bytes: %v\n\treceived bytes: %v",
			expectedIdString, newIdString,
			[]byte(expectedIdString), newID.Bytes()[:ArrIDLen-1])
	}
}

// Tests that NewIdFromString() panics when given a nil testing object.
func TestNewIdFromString_TestError(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("NewIdFromString() did not panic when it received a " +
				"nil testing object when it should have.")
		}
	}()

	// Call function with nil testing object
	_ = NewIdFromString("test", Generic, nil)
}

// Tests that NewIdFromUInt() creates a new ID with the correct contents by
// converting the ID back into a uint and comparing it to the original.
func TestNewIdFromUInt(t *testing.T) {
	// Expected values
	expectedUint := rand.Uint64()

	// Create the ID and check its contents
	newID := NewIdFromUInt(expectedUint, Generic, t)
	idUint := binary.BigEndian.Uint64(newID[:ArrIDLen-1])

	if expectedUint != idUint {
		t.Errorf("NewIdFromUInt() produced an ID from uint incorrectly."+
			"\n\texpected: %v\n\treceived: %v", expectedUint, idUint)
	}
}

// Tests that NewIdFromUInt() panics when given a nil testing object.
func TestNewIdFromUInt_TestError(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("NewIdFromUInt() did not panic when it received a " +
				"nil testing object when it should have.")
		}
	}()

	// Call function with nil testing object
	_ = NewIdFromUInt(rand.Uint64(), Generic, nil)
}

// Tests that NewIdFromUInts() creates a new ID with the correct contents by
// converting the ID back into a slice of uints and comparing it to the
// original.
func TestNewIdFromUInts(t *testing.T) {
	// Expected values
	expectedUints := [4]uint64{rand.Uint64(), rand.Uint64(),
		rand.Uint64(), rand.Uint64()}

	// Create the ID and check its contents
	newID := NewIdFromUInts(expectedUints, Generic, t)
	idUints := [4]uint64{}
	for i := range idUints {
		idUints[i] = binary.BigEndian.Uint64(newID[i*8 : (i+1)*8])
	}

	if !reflect.DeepEqual(expectedUints, idUints) {
		t.Errorf("NewIdFromUInts() produced an ID from uints incorrectly."+
			"\n\texpected: %#v\n\treceived: %#v", expectedUints, idUints)
	}
}

// Tests that NewIdFromUInts() panics when given a nil testing object.
func TestNewIdFromUInts_TestError(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("NewIdFromUInts() did not panic when it received a " +
				"nil testing object when it should have.")
		}
	}()

	// Call function with nil testing object
	newUint64s := [4]uint64{rand.Uint64(), rand.Uint64(),
		rand.Uint64(), rand.Uint64()}
	_ = NewIdFromUInts(newUint64s, Generic, nil)
}

// Tests that Marshal() returns the correct byte slice of an ID and that it is a
// copy. This test ensures duplicates the test for Bytes() to ensure that
// Marshal() uses Bytes().
func TestID_Marshal(t *testing.T) {
	// Test values
	expectedBytes := newRandomBytes(ArrIDLen, t)
	testID := NewIdFromBytes(expectedBytes, t)

	// Check for the correct values
	testVal := testID.Marshal()
	if !bytes.Equal(expectedBytes, testVal) {
		t.Errorf("Marshal() returned the incorrect byte slice of the ID"+
			"\n\texpected: %+v\n\treceived: %+v", expectedBytes, testVal)
	}

	// Test if the returned bytes are copies
	if &testID[0] == &testVal[0] {
		t.Errorf("Marshal() did not return a copy when it should have."+
			"\n\texpected: any value except %+v\n\treceived: %+v",
			&testID[0], &testVal[0])
	}
}

// Tests that Unmarshal() creates a new ID with the correct contents and does
// not return an error.
func TestUnmarshal(t *testing.T) {
	// Test values
	expectedBytes := newRandomBytes(ArrIDLen, t)

	// Unmarshal the bytes into an ID
	newID, err := Unmarshal(expectedBytes)

	// Make sure no error occurred
	if err != nil {
		t.Errorf("Unmarshal() produced an unexpected error."+
			"\n\texpected: %v\n\treceived: %v", nil, err)
	}

	// Make sure the ID contents are correct
	if !bytes.Equal(expectedBytes, newID[:]) {
		t.Errorf("Unmarshal() produced an ID with the incorrect bytes."+
			"\n\texpected: %v\n\treceived: %v", expectedBytes, newID[:])
	}
}

// Tests that Unmarshal() produces an error when the given data length is not
// equal to the length of an ID and that the ID returned is nil.
func TestUnmarshal_DataLengthError(t *testing.T) {
	// Test values
	expectedBytes := newRandomBytes(ArrIDLen+10, t)
	expectedError := fmt.Errorf("could not marshal byte slice to ID: "+
		"length of data must be %d, length received was %d",
		ArrIDLen, len(expectedBytes))

	// Unmarshal the bytes into an ID
	newID, err := Unmarshal(expectedBytes)

	// Make sure an error occurs
	if err == nil {
		t.Errorf("Unmarshal() did not product an expected error."+
			"\n\texpected: %v\n\treceived: %v", expectedError, err)
	}

	// Make sure the returned ID is nil
	if newID != nil {
		t.Errorf("Unmarshal() produced a non-nil ID on error."+
			"\n\texpected: %v\n\treceived: %v", nil, newID)
	}
}

// Generates a byte slice of the specified length containing random numbers.
func newRandomBytes(length int, t *testing.T) []byte {
	// Create new byte slice of the correct size
	idBytes := make([]byte, length)

	// Create random bytes
	_, err := rand.Read(idBytes)
	if err != nil {
		t.Fatalf("Failed to generate random bytes: %v", err)
	}

	return idBytes
}
