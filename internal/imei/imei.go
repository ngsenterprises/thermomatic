// Package imei implements an IMEI decoder.
package imei

// NOTE: for more information about IMEI codes and their structure you may
// consult with:
//
// https://en.wikipedia.org/wiki/International_Mobile_Equipment_Identity.

import (
	//"fmt"
	"errors"
	"thermomatic/internal/common"
)

const (
	IMEI_CODE_LEN = 15
)

var (
	ErrInvalid  = errors.New("imei: invalid ")
	ErrChecksum = errors.New("imei: invalid checksum")
)

// type ImeiLoginChan struct {
// 	LoginCodes map[uint64]bool   
// 	CAddCode chan uint64
// 	CDeleteCode chan uint64
// 	COut chan bool
// }

// Decode returns the IMEI code contained in the first 15 bytes of b.
//
// In case b isn't strictly composed of digits, the returned error will be
// ErrInvalid.
//
// In case b's checksum is wrong, the returned error will be ErrChecksum.
//
// Decode does NOT allocate under any condition. Additionally, it panics if b
// isn't at least 15 bytes long.

func ReverseImei(data []byte) []byte {
	if !common.HasLength( data, 15 ) {
		return data
	}
	for i, j := 0, 14; i < j; i, j = i +1, j -1 {
		data[i], data[j] = data[j], data[i]
	}
	return data
} 

func encode(code uint64, index int, ac []byte) (data []byte) {
	if code == 0 {
		return ReverseImei(ac)
	} 
	ac[index] = byte(code % uint64(10))
	return encode(code / uint64(10), index +1, ac)
}

// convert 15-digit imei code to byte array 
func Encode(code uint64) (data []byte) {
	return encode(code, 0, make([]byte, 15))
}

func decode( data []byte, index uint, sum uint, ac uint64 ) ( uint64, error) {
	if IMEI_CODE_LEN <= index {
		if sum % 10 == 0 {
			return  ac, nil
		} 
		return  0, ErrInvalid
	}
	if data[index] < 0 || 9 < data[index] {
		return 0, ErrInvalid
	} 
	if index % 2 == 1 { 
		return decode( data,
			index +1, 
			sum +( ( uint(2 * data[index]) / uint(10)) +( uint(2 * data[index]) % uint(10)) ),
			uint64(10)*ac +uint64(data[index]) )
	} 

	return decode( data,
		index +1, 
		sum +uint(data[index]),
		uint64(10)*ac +uint64(data[index]) )
	
}

// convert byte array to 15-digit imei code
func Decode(data []byte) (code uint64, err error) {
	
	if !common.HasLength( data, IMEI_CODE_LEN ) {
		return 0, ErrInvalid
	}

	code, err = decode( data, uint(0), uint(0), uint64(0) )

	if err != nil {
		return uint64(0), err 
	}		

	return code, err
}


