package simulator

import "fmt"
import "time"
import "thermomatic/internal/common"
import "thermomatic/internal/client"
import "thermomatic/internal/imei"

func StartSimulator() {

	imei_codes := []uint64{	
		358265010779665,
		358265016799402,
		358265015739151,
		358265012799638,
		358265017779692,
		358265016779230,
		358265010789508,
		358265013749921,
		358265017749240,
		358265019759932,
		358265014779323,
		358265012719271,
		358265016729466,
		358265017799013,
		358265010779178,
		358265017739480,
		358265011789853,
		358265018769114,
		358265016799766,
		358265014729559,
		358265017749109,
		358265014779877,
		358265015749291,
		358265017789550,
		358265012769474,
	}	
	
	for index := 0; index < common.SIMULATOR_NODES; index++ {
		fmt.Printf( "StartSimulator: %d \n", index )

		go client.StartClient(imei.Encode(imei_codes[index]))
	
		time.Sleep( 250 * time.Millisecond )	
	} 
}


