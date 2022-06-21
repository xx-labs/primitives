////////////////////////////////////////////////////////////////////////////////
// Copyright © 2020 xx network SEZC                                           //
//                                                                            //
// Use of this source code is governed by a license that can be found in the  //
// LICENSE file                                                               //
////////////////////////////////////////////////////////////////////////////////

package exponential

import (
	"reflect"
	"testing"
)

// Tests that DefaultMovingAvgParams returns a MovingAvgParams with all the
// default values.
func TestDefaultMovingAvgParams(t *testing.T) {
	expected := MovingAvgParams{
		Cutoff:          defaultCutoff,
		InitialAverage:  defaultInitialAverage,
		SmoothingFactor: defaultSmoothingFactor,
		NumberOfEvents:  defaultNumberOfEvents,
	}
	p := DefaultMovingAvgParams()

	if !reflect.DeepEqual(expected, p) {
		t.Errorf("Did not received expected default parameters."+
			"\nexpected: %+v\nreceived: %+v", expected, p)
	}
}
