////////////////////////////////////////////////////////////////////////////////////////////
// Copyright © 2020 xx network SEZC                                                       //
//                                                                                        //
// Use of this source code is governed by a license that can be found in the LICENSE file //
////////////////////////////////////////////////////////////////////////////////////////////

package id

import (
	"fmt"
	"testing"
)

// Tests that Type.String() returns the correct string for each Type.
func TestType_String(t *testing.T) {
	testValues := map[Type]string{
		Generic:  "generic",
		Gateway:  "gateway",
		Node:     "node",
		User:     "user",
		NumTypes: "4",
	}

	for idType, expected := range testValues {
		if expected != idType.String() {
			t.Errorf("String() returned incorrect string for type."+
				"\n\texpected: %s\n\treceived: %s", expected, idType.String())
		}
	}
}

// Tests that Type.String() returns an error when given a Type that has not been
// defined.
func TestType_String_Error(t *testing.T) {
	testType := Type(5)
	expectedError := fmt.Sprintf("%s%d", noIDTypeErr, testType)

	// Test stringer error
	testVal := testType.String()
	if expectedError != testVal {
		t.Errorf("String() did not return an error when it should have."+
			"\n\texpected: %s\n\treceived: %s", expectedError, testVal)
	}
}
