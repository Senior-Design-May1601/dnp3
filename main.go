package main
import(
	"net"
	"log"
	"fmt"
	"io"
	"bytes"
	

	"github.com/Senior-Design-May1601/Splunk/alert"
	"github.com/Senior-Design-May1601/projectmain/logger"
)

type Headers struct{
	DataLink DataLayer_t
	AppLayer AppLayer_t
}

func handler(c net.Conn){
	defer c.Close()
	var buf bytes.Buffer
	io.Copy(&buf, c)
	fmt.Println(buf.String())

// END GET APPLICATION DATA

	transport_data,d_struct :=DataLinkRead(buf.Bytes()) // Strip DataLink Header
	app_data:=TransportRead(transport_data)
	app_struct := AppRead(app_data)
	data := Headers{d_struct,app_struct}		
	mylogger.Println(makeAlert(data))	
}
func makeAlert(h Headers) string{
	meta := make(map[string]string)
	meta["source"] = h.DataLink.source
	meta["origin"] = h.DataLink.origin
	meta["dlayer_code"] = h.DataLink.f_code
	meta["alayer_code"] = h.AppLayer.function_code
		
	return alert.NewSplunkAlertMessage(meta)
}

var mylogger *log.Logger
func main(){
	mylogger := logger.NewLogger("",0)
	//Add toml stuff

	l, e := net.Listen("tcp",":9000")
	if e != nil {
		mylogger.Fatal("dnp listen error",e)
	}
	defer l.Close()
	for{
		conn, e := l.Accept()
		if e != nil {
			mylogger.Fatal("dnp failed connection",e)
		}
		go handler(conn)
	}	
}
