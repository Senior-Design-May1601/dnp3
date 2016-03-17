package main
import(
	"fmt"
)
type AppLayer_t struct {
	function_code string
}

func AppRead(data []byte) AppLayer_t{
	var func_code string
	if val, ok := f_codes[data[1]]; ok {
		func_code = val
	}else{
		func_code = fmt.Sprintf("Reserverd code %x",data[1])
	}

	return AppLayer_t{func_code}

}
