package main
import(
	"net"
	"log"
	"io"
	"bytes"
	"os"	
	"flag"
	"strconv"
	
	"github.com/Senior-Design-May1601/Splunk/alert"
	"github.com/Senior-Design-May1601/projectmain/logger"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Address string
	Port int
}

func handler(c net.Conn,mylogger *log.Logger){
	defer c.Close()

	var buf bytes.Buffer

	fileLog := logger.NewLogger("",0)
	f, _ := os.OpenFile("dnp-log-file",  os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)	
	fileLog.SetOutput(f)
	defer f.Close()



	remoteAddr := c.RemoteAddr()
	localAddr := c.LocalAddr()

	for{
		io.Copy(&buf, c)
		fileLog.Println("all data", buf)
		resp,dl :=DataLinkRead(buf.Bytes()) // Strip DataLink header
		resp,final:= TransportRead(resp) //Strip Transport header
		ap,aFinal:=AppRead(resp) //Strip Application header	
		fileLog.Println("finalApp",aFinal)
		fileLog.Println("finalT",final)	
	

		str := makeAlert(dl,ap,remoteAddr,localAddr)
		fileLog.Println(str)
		mylogger.Println(str)

		
		buf.Reset()	
	}

}
func makeAlert(dl DataLayer_t, app AppLayer_t, remoteAddr net.Addr,localAddr net.Addr) string{
	meta := make(map[string]string)
	meta["machine_source"] = dl.source
	meta["origin"] = dl.origin
	meta["ln_function_code"] = dl.f_code
	meta["ap_function_code"] = app.function_code
	meta["remote_address"] = remoteAddr.String()
	meta["local_address"] = localAddr.String()
	resp :=  alert.NewSplunkAlertMessage(meta)
	return resp
}
func makeUnknownAlert(remoteAddr net.Addr,localAddr net.Addr) string{
	meta := make(map[string]string)
	meta["remote_address"] = remoteAddr.String()
	meta["message"] = "An unknown application protocol has connected to this device"
	meta["local_address"] = localAddr.String()
	return alert.NewSplunkAlertMessage(meta)
}

var mylogger *log.Logger
var config Config

func main(){
	mylogger := logger.NewLogger("",0)	

	configPath := flag.String("config","","config file")
	flag.Parse()

	if _, err := toml.DecodeFile(*configPath, &config); err != nil {
		mylogger.Fatal("cannot decode file",err)
	}

	l, e := net.Listen("tcp",config.Address+":"+strconv.Itoa(config.Port))
	if e != nil {
		mylogger.Fatal("dnp listen error",e)
	}
	defer l.Close()
	for{
		conn, e := l.Accept()
		if e != nil {
			mylogger.Fatal("dnp failed connection",e)
		}
		go handler(conn,mylogger)
	}	
}
