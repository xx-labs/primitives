////////////////////////////////////////////////////////////////////////////////
// Copyright © 2020 Privategrity Corporation                                   /
//                                                                             /
// All rights reserved.                                                        /
////////////////////////////////////////////////////////////////////////////////

package id

import (
	"testing"
)

// Tests that GetHardCodedIDs() returns all the hard coded IDs in the order that
// they were added.
func TestGetHardCodedIDs(t *testing.T) {
	expectedIDs := []*ID{&Permissioning, &NotificationBot, &TempGateway,
		&ZeroUser, &DummyUser, &UDB}

	for i, testID := range GetHardCodedIDs() {
		if !expectedIDs[i].Cmp(testID) {
			t.Errorf("GetHardCodedIDs() did not return the expected ID at "+
				"index %d.\n\texepcted: %v\n\trecieved: %v",
				i, expectedIDs[i], testID)
		}
	}
}

// Tests that CollidesWithHardCodedID() returns false when none of the test IDs
// collide with the hard coded IDs.
func TestCollidesWithHardCodedID_HappyPath(t *testing.T) {
	testIDs := []*ID{
		NewIdFromString("Test1", Generic, t),
		NewIdFromString("Test2", Gateway, t),
		NewIdFromString("Test3", Node, t),
		NewIdFromString("Test4", User, t),
	}

	for _, testID := range testIDs {
		if CollidesWithHardCodedID(testID) {
			t.Errorf("CollidesWithHardCodedID() found collision when none "+
				"should exist.\n\tcolliding ID: %v", testID)
		}
	}
}

// Tests that CollidesWithHardCodedID() returns true when checking if hard coded
// IDs collide.
func TestCollidesWithHardCodedID_True(t *testing.T) {
	testIDs := []*ID{&Permissioning, &NotificationBot, &TempGateway,
		&ZeroUser, &DummyUser, &UDB}

	for _, testID := range testIDs {
		if !CollidesWithHardCodedID(testID) {
			t.Errorf("CollidesWithHardCodedID() did not find a collision when "+
				"one should exist.\n\tcolliding ID: %v", testID)
		}
	}
}
