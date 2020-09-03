package log
import (
"os"
"log"
"io"
)

var (
 Debug   *log.Logger
 Info    *log.Logger
 Warning *log.Logger
 Error   *log.Logger
)

func init(){
  file, err := os.OpenFile("/var/log/grafana-exporter.log",
       os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
  if err != nil{
   log.Fatalln("Failed to open grafana-exporter log file: ",err)
  }
  Debug = log.New(io.MultiWriter(file,os.Stderr),
    "Debug: ",
    log.Ldate|log.Ltime|log.Lshortfile)
  Info = log.New(file,
    "Info: ",
    log.Ldate|log.Ltime|log.Lshortfile)
  Warning = log.New(file,
    "Warning: ",
    log.Ldate|log.Ltime|log.Lshortfile)
  Error = log.New(file,
    "Error: ",
    log.Ldate|log.Ltime|log.Lshortfile)
}
