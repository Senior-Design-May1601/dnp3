package main
import(
	"fmt"
	"crypto/rand"	
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

func AppResponseHeader()[]byte{
	header := make([]byte,0,7)
	header[0] = 0xC0 //FIN, FIR bit 
	header[1] = 0x83 // Authenticate Response 3.1.2
	header[2] = 0x00
	header[3] = 0x00
	header[4] = 120 //group number
	header[5] = 1 //variation
	header[6] = 0x06 //no prefix no range qualifier


	return header
}

func G120v1()[]byte {
 	buf := make([]byte,19)
	rand.Read(buf)

	buf[0] = 0xC0 //FIN, FIR bit 
	buf[1] = 0x83 // Authenticate Response 3.1.2
	buf[2] = 0x00
	buf[3] = 0x00
	buf[4] = 120 //group number
	buf[5] = 1 //variation
	buf[6] = 0x06 //no prefix no range qualifier
	buf[13] = 0x04 //Hmac Sha-256 16 bit
	buf[14] = 0x01 // Reason for challenge-> Critical
	
	return buf	

}

