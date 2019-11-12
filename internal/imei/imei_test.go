package imei

import (
	"testing"
	//"fmt"
	"bytes"
	//"thermomatic/internal/common"
)

const (
	aimei = 490154203237518
	imei1 = 358265010779665
	imei2 = 358265016799402
	imei3 = 358265015739151
	imei4 = 358265012799638
	imei5 = 358265017779692
	imei6 = 358265016779230
	imei7 = 358265010789508
)

func TestReverse(t *testing.T) {
	data := []byte{1,1,2,2,3,3,4,4,5,5,6,6,7,8,9,}
	res := ReverseImei( data )
	expected := []byte{9,8,7,6,6,5,5,4,4,3,3,2,2,1,1,}
	if bytes.Compare(res, expected) != 0 {
		t.Fatalf("Results %v", expected)
	}
}

func TestEncode(t *testing.T) {
	res := Encode(imei1)
	expected := []byte{9,1,1,3,2,3,0,0,0,4,1,3,2,2,1,}
	if bytes.Compare(res, expected) != 0 {
		t.Fatalf("Results %v", expected)
	}
}

func TestDecode(t *testing.T) {
	data := Encode(imei1)
	res, err := Decode(data)
	if err != nil {
		t.Fatalf("TestDecode: err %s", err.Error())
	}
	if res != imei1 {
		t.Fatalf("TestDecode: res %v", res)
	}
}



/*
func TestDecode(t *testing.T) {
	panic(common.ErrNotImplemented)
}

func TestDecodeAllocations(t *testing.T) {
	panic(common.ErrNotImplemented)
}

func TestDecodePanics(t *testing.T) {
	panic(common.ErrNotImplemented)
}

func BenchmarkDecode(b *testing.B) {
	panic(common.ErrNotImplemented)
}
*/
