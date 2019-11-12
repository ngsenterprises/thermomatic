package client

import(
	"math"
	"math/rand"
	"errors"
	"encoding/binary"
	"thermomatic/internal/common"
)


// Reading is the set of device readings.
type Reading struct {
	// Temperature denotes the temperature reading of the message.
	Temperature float64

	// Altitude denotes the altitude reading of the message.
	Altitude float64

	// Latitude denotes the latitude reading of the message.
	Latitude float64

	// Longitude denotes the longitude reading of the message.
	Longitude float64

	// BatteryLevel denotes the battery level reading of the message.
	BatteryLevel float64
}

//convert Reading into byte-array
func (r *Reading)Encode(b []byte) bool {

	if !common.HasLength(b, 40) {
		return false	
	}

	binary.BigEndian.PutUint64(b[0:8], math.Float64bits(r.Temperature))
	binary.BigEndian.PutUint64(b[8:16], math.Float64bits(r.Altitude))
	binary.BigEndian.PutUint64(b[16:24], math.Float64bits(r.Latitude))
	binary.BigEndian.PutUint64(b[24:32], math.Float64bits(r.Longitude))
	binary.BigEndian.PutUint64(b[32:40], math.Float64bits(r.BatteryLevel))
	
	return true
}


// Decode decodes the reading message payload in the given b into r.
//
// If any of the fields are outside their valid min/max ranges ok will be unset.
//
// Decode does NOT allocate under any condition. Additionally, it panics if b
// isn't at least 40 bytes long.
func (r *Reading) Decode(b []byte) (ok bool) {
	
	if !common.HasLength(b, 40) {
		panic(errors.New("Reading Decode: bad length"))
	}

    r.Temperature = math.Float64frombits(binary.BigEndian.Uint64(b[0:8]))
    r.Altitude = math.Float64frombits(binary.BigEndian.Uint64(b[8:16]))
    r.Latitude = math.Float64frombits(binary.BigEndian.Uint64(b[16:24]))
    r.Longitude = math.Float64frombits(binary.BigEndian.Uint64(b[24:32]))
	r.BatteryLevel = math.Float64frombits(binary.BigEndian.Uint64(b[32:40]))

	return true
}

// create a random Reading object
func CreateRandomReading( rand *rand.Rand ) Reading {
	r :=  Reading{}
	r.Temperature = math.Min(float64(600.0001)*rand.Float64() -float64(300.0), float64(300)) // [0.0,1.0) -> [-300, 300]
	r.Altitude = math.Min(float64(40000.0001)*rand.Float64() -float64(20000.0), float64(20000.0))  // [0.0,1.0) -> [-20000, 20000]
	r.Latitude = math.Min(float64(180.0001)*rand.Float64() -float64(90.0), float64(90.0)) //  [0.0,1.0) -> [-90, 90]
	r.Longitude = math.Min(float64(360.0001)*rand.Float64() -float64(180.0), float64(180.0)) // [0.0,1.0) -> [-180, 180]
	r.BatteryLevel = math.Max(float64(0.0001), math.Min(float64(100.0001)*rand.Float64(), float64(100.0))) // [0.0,1.0) -> (0, 100]
	
	return r
}

