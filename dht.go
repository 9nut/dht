// This package uses the low level C routines from
// the Adafruit Python DHT library located at:
// https://github.com/adafruit/Adafruit_Python_DHT

package dht

/*
#cgo CFLAGS: -std=gnu99
#cgo LDFLAGS: -lrt

#include "common_dht_read.h"
#include "pi_dht_read.h"

void dhtread(int type, int pin, float* temp, int *humi, int *st) {
	float fhumi;
	int ret;

	*st = pi_dht_read(DHT22, pin, &fhumi, temp);
	if( *st == DHT_SUCCESS )
		*humi = (int)fhumi;
	else
		*st = -*st;
}

*/
import "C"

import (
	"errors"
)

// these correlate to DHT_ERROR_* values in common_dht_read.h
var DHTErrs []error = []error{
	nil,
	errors.New("Timed out"),
	errors.New("Bad Checksum"),
	errors.New("Bad Argument"),
	errors.New("GPIO Error"),
}

func DHTRead(t, pin int) (temp float32, humi int32, err error) {
	var st int32
	C.dhtread(C.int(t), C.int(pin), (*C.float)(&temp), (*C.int)(&humi), (*C.int)(&st))

	err = DHTErrs[st]
	return
}
