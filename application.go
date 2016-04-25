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

func G120v1(remote string,local string)[]byte {
 	buf := make([]byte,29)
	r := []byte(remote)
	l := []byte(local)
	rand.Read(buf) //fill buffer with random bytes

	buf[0] = 0x05 //Preamble [0],[1]
	buf[1] = 0x64
	buf[2] = 28 // lenght
	buf[3] = 0x00 //ACK
	buf[4] = l[0] // local addr[0],[1]
	buf[5] = l[1]
	buf[6] = r[0] //remote addr[0],[1]
	buf[7] = r[1]
	// 8 & 9 are CRC bits

	//begin transport header
	buf[10] = 0xC1 //FIN, FIR bit Transport layer

	// Begin app header
	buf[11] = 0xC0 //FIN, FIR bit 
	buf[12] = 0x83 // Authenticate Response 3.1.2
	buf[13] = 0x00
	buf[14] = 0x00

	// Begin Object header
	buf[15] = 120 //group number
	buf[16] = 1 //variation
	buf[17] = 0x06 //no prefix no range qualifier

	// 18:21 UINT32 Challenge Sequence number
	// 22:23 UINT16 User number for session keys
	buf[24] = 0x04 //Hmac Sha-256 16 bit
	buf[25] = 0x01 // Reason for challenge-> Critical

	//4 extra bytes for challenge
	
	return buf	

}

