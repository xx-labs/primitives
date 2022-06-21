////////////////////////////////////////////////////////////////////////////////
// Copyright © 2020 xx network SEZC                                           //
//                                                                            //
// Use of this source code is governed by a license that can be found in the  //
// LICENSE file                                                               //
////////////////////////////////////////////////////////////////////////////////

package exponential

import (
	"github.com/pkg/errors"
	"sync"
)

const (
	defaultA0              = 0.15
	defaultSmoothingFactor = 2
	defaultNumberOfEvents  = 100
)

// MovingAvg tracks the exponential moving average across a number of events and
// reports when it has surpassed the set cutoff.
type MovingAvg struct {
	cutoff float32 // Maximum limit for the average
	aN     float32 // A(n), current average (initialize to A(0))
	s      float32 // Exponential smoothing factor
	e      uint32  // Number of events to average over
	sync.Mutex
}

// NewMovingAvg creates a new MovingAvg with the given cutoff, initial average,
// smoothing factor, and number of events to average.
func NewMovingAvg(cutoff, a0, s float32, e uint32) *MovingAvg {
	return &MovingAvg{
		cutoff: cutoff,
		aN:     a0,
		s:      s,
		e:      e,
	}
}

// Intake takes in the current average and calculates the exponential average
// returning true if it is over the cutoff and false otherwise.
func (m *MovingAvg) Intake(a float32) error {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	m.aN = (a * (m.s / float32(m.e))) + (m.aN * (1 - (m.s / float32(m.e))))

	if m.aN > m.cutoff {
		return errors.Errorf("exponential average for the last %d events of "+
			"%.2f%% went over cutoff %.2f%%", m.e, m.aN*100, m.cutoff*100)
	}
	return nil
}

// IsOverCutoff returns true if the average has reached the cutoff and false if
// it has not
func (m *MovingAvg) IsOverCutoff() bool {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	return m.aN > m.cutoff
}

// boolToFloat returns 1 if true and 0 if false.
func boolToFloat(b bool) float32 {
	if b {
		return 1
	}
	return 0
}
