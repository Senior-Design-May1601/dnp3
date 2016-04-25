package main
import(
	"fmt"
)
type AppLayer_t struct {
	function_code string
}

func AppRead(data []byte) (AppLayer_t,int){
	var func_code string
	control := data[0]
	final:= control & 0x80 >> 7

	if val, ok := f_codes[data[1]]; ok {
		func_code = val
	}else{
		func_code = fmt.Sprintf("Reserverd code %x",data[1])
	}

	return AppLayer_t{func_code}, int(final)

}

func ApplicationResponse(){


}

func makeG120()[]byte {
 	var object []byte
	object[0] = 120 //group number
	object[1] = 1 // variation
	object[2] = 0 // no prefix
	
	//begin arbitrary sequence num
	object[3] = 34
	object[4] = 3
	object[5] = 100
	object[6] = 64
	//end sequenc num
	
	object[7] = 0 //usr number unknown because outstation
	object[8] = 1 // sha-1  4 bytes
	object[9] = 1 // reason for challenge = Critical adsu

	//begin challenge data 4 bytes sha-1 could make these random
	object[10] = 0xFE
	object[11] = 0x00
	object[12] = 0x33
	object[13] = 0xAB
	//end challenge	
	return object
}

