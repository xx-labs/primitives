////////////////////////////////////////////////////////////////////////////////////////////
// Copyright © 2020 xx network SEZC                                                       //
//                                                                                        //
// Use of this source code is governed by a license that can be found in the LICENSE file //
////////////////////////////////////////////////////////////////////////////////////////////

// Package hw contains files for identifying the hardware inside of a computer

package hw

import (
	"os/exec"

	jww "github.com/spf13/jwalterweatherman"
)

func LogHardware() error {
	// lscpu
	out, err := exec.Command("lscpu").Output()
    if err != nil {
        return err
    }
	jww.INFO.Printf("[HWINFO] CPU INFO:\r\n%s", out)

	// lspci GPUs
	out, err = exec.Command("lspci -vnnn | perl -lne 'print if /^\\d+\\:.+(\\[\\S+\\:\\S+\\])/' | grep VGA").Output()
    if err != nil {
        return err
    }
	jww.INFO.Printf("[HWINFO] GPU INFO:\r\n%s", out)

	// lsblk
	out, err = exec.Command("lsblk").Output()
    if err != nil {
        return err
    }
	jww.INFO.Printf("[HWINFO] PARTITION INFO:\r\n%s", out)

	// df disk usage
	out, err = exec.Command("df -h").Output()
    if err != nil {
        return err
    }
	jww.INFO.Printf("[HWINFO] DISK USAGE INFO:\r\n%s", out)

	// disk hw info
	out, err = exec.Command("lshw -class disk -class storage").Output()
	if err != nil {
		return err
	}
	jww.INFO.Printf("[HWINFO] DISK HW INFO:\r\n%s", out)

	// RAM info
	out, err = exec.Command("dmidecode --type 17").Output()
    if err != nil {
        return err
    }
	jww.INFO.Printf("[HWINFO] RAM HW INFO:\r\n%s", out)

	// RAM usage
	out, err = exec.Command("free -mt").Output()
	if err != nil {
		return err
	}
	jww.INFO.Printf("[HWINFO] RAM USAGE INFO:\r\n%s", out)

	return nil
}