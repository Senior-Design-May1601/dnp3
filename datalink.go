package main
import(
	"fmt"
)

// "-" Reserved or Obsolete
var pf_codes = map[byte]string{
	0 : "RESET_LINK_STATES",	
	1 : "-",
	2 : "TEST_LINK_STATE",
	3 : "CONFIRMED_USER_DATA",
	4 : "UNCONFIRMED_USER_DATA",
	5 : "-",
	6 : "-",
	7 : "-",
	8 : "-",
	9 : "REQUEST_LINK_STATUS",
	10 : "-",
	11 : "-",
	12 : "-",
	13 : "-",
	14 : "-",
	15 : "-",
}
var sf_codes = map[byte]string{
	0 : "ACK",
	1 : "NAK",
	2 : "-",
	3 : "-",
	4 : "-",
	5 : "-",
	6 : "-",
	7 : "-",
	8 : "-",
	9 : "-",
	10 : "-",
	11 : "LINK_STATUS",
	12 : "-",
	13 : "-",
	14 : "-",
	15 : "NOT_SUPPORTED",
}
var origin_code = map[byte]string{
	0 : "Outstation",
	1 : "Master",
}

type DataLayer_t struct {	
	f_code string
	origin string
	source string
}
func DataLinkRead(data []byte) ([]byte,DataLayer_t){
	//upper layer data
	payload := data[10:] 
	
	// Primary bit
	PRM := (data[3] & 0x40) >> 6

	// master or outstation
	origin := (data[3] & 0x80) >> 7

	// Get source Address
	source := fmt.Sprintf("%x%x",data[7],data[6])

	fmt.Println(source)
	func_code := data[3] & 0x0F
	var code string

	if PRM == 0x00 {
		code = sf_codes[func_code]
	}
	
	if PRM == 0x01 {
		code = pf_codes[func_code]
	}
	return payload,DataLayer_t{code,origin_code[origin],source}
}
