package gomobilelib

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strconv"
	"strings"

	"compress/gzip"
	"os"
)

var indexServer int = 0

var flagInitServer bool = false

var jc JavaCallBack

type JavaCallBack interface {
	PlayOnAndroid(string)
	SpeakAdd(string)
	WriteInFile(string, string) string
	FlagRam(int)
	WaitMutex(int)
	GetVoice() string
	GetEngine() string
	GetLocal() string
	SetEngine(string) int
	SetVoice(string) int
	SetLanguage(string, string) int
	SetSpeechRate(float64) int
	SetPitch(float64) int
}

func Greetings(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func RegisterJavaCallBack(c JavaCallBack) {
	jc = c
}

type ConnScruct struct {
	//Conn    net.Conn
	flagZip int
}

//var Listener net.Conn

func Zreetings(name string) string {
	return "send ok"
}

func (s *ConnScruct) FlagZip(in int, reply *string) error {
	s.flagZip = in
	log.Println("s.flagZip = ", s.flagZip)
	return nil
}
func (s *ConnScruct) PlayOnAndroid(in string, reply *string) error {
	jc.PlayOnAndroid(in)
	*reply = "PlayOnAndroid Ok"
	return nil
}

func (s *ConnScruct) SpeakAdd(in string, reply *string) error {
	jc.SpeakAdd(in)
	*reply = "SpeakAdd Ok"
	return nil
}

var fileNum int

func (s *ConnScruct) FlagRam(in int, reply *string) error {
	jc.FlagRam(in)
	return nil
}

func (s *ConnScruct) GetLocal(in int, reply *string) error {
	*reply = jc.GetLocal()
	return nil
}
func (s *ConnScruct) GetEngine(in int, reply *string) error {
	*reply = jc.GetEngine()
	return nil
}

func (s *ConnScruct) GetVoice(in int, reply *string) error {
	*reply = jc.GetVoice()
	return nil
}

func (s *ConnScruct) SetVoice(in string, reply *string) error {
	_ = jc.SetVoice(in)
	return nil
}

type DoubleString struct {
	one string
	two string
}

func (s *ConnScruct) SetLanguage(in string, reply *string) error {
	ar := strings.Split(in, "-")
	if len(ar) == 2 {
		//log.Println("one ", ar[0], " two ", ar[1])
		if len(ar) == 2 {
			_ = jc.SetLanguage(ar[0], ar[1])
		} else {
			_ = jc.SetLanguage(in, "")
		}
	}
	return nil
}

func (s *ConnScruct) SetSpeechRate(in float64, reply *string) error {
	log.Println("SetSpeechRate: ", in)
	ret := jc.SetSpeechRate(in)
	log.Println("SetSpeechRate ret = ", ret)
	return nil
}

func (s *ConnScruct) SetPitch(in float64, reply *string) error {
	log.Println("SetPitch: ", in)
	_ = jc.SetPitch(in)
	return nil
}

func (s *ConnScruct) HelloAndroid(in string, reply *string) error {
	*reply = "Hello i'am Android"
	return nil
}

// После изменения движка нужен Sleep секунды на 2
func (s *ConnScruct) SetEngine(in string, reply *int) error {
	log.Println("SetEngine: ", in)
	*reply = jc.SetEngine(in)
	return nil
}

func (s *ConnScruct) WriteInFile(in string, reply *[]byte) error {
	var err error
	var bufferByte []byte
	fileName := strconv.Itoa(fileNum) + ".mp3"
	fileNum++
	if fileNum > 5 {
		fileNum = 0
	}
	allPathFile := jc.WriteInFile(in, fileName)
	log.Println("go WaitMutex")
	jc.WaitMutex(0)

	if s.flagZip == 1 {
		log.Println("go zip")
		bufferByte, err = ioutil.ReadFile(allPathFile + fileName) // just pass the file name
		if err != nil {
			log.Println("Error: ", err)
			os.Remove(allPathFile + fileName)
			return errors.New("No such file")
		}
		var buf bytes.Buffer
		w, _ := gzip.NewWriterLevel(&buf, gzip.BestSpeed)
		//w := gzip.NewWriter(&buf)
		defer w.Close()
		w.Write(bufferByte)

		*reply = buf.Bytes()
	} else {
		log.Println("go send")
		*reply, err = ioutil.ReadFile(allPathFile + fileName) // just pass the file name
		if err != nil {
			log.Println("Error: ", err)
			os.Remove(allPathFile + fileName)
			return errors.New("No such file")
		}
	}

	os.Remove(allPathFile + fileName)
	//*reply = "WriteInFile Ok byte = " + strconv.Itoa(len(b))
	//*reply
	return nil
}

func (s *ConnScruct) WaitMutex(in int, reply *string) error {
	jc.WaitMutex(in)
	*reply = "Wait Ok"
	return nil
}

func StartServer() {
	log.Println("indexServer: ", indexServer)
	if flagInitServer == false {
		go asynServ()
		flagInitServer = true
	}
	indexServer++
}

func asynServ() error {
	return nil
}

type ServerCodec struct {
	//Gateway rpc.ServerCodec
}

func NewServerCodec(conn net.Conn) *ServerCodec {
	return &ServerCodec{}
}

//func (s *ServerCodec) ReadRequestHeader(req *rpc.Request) error {
////log.Println("ReadRequestHeader(): ", req)
//return s.Gateway.ReadRequestHeader(req)
//}

//func (s *ServerCodec) ReadRequestBody(obj interface{}) error {
////log.Println("ReadRequestBody()")
//return s.Gateway.ReadRequestBody(obj)
//}

//func (s *ServerCodec) WriteResponse(resp *rpc.Response, obj interface{}) error {
////log.Println("WriteResponse()")
//return s.Gateway.WriteResponse(resp, obj)
//}

//// Close closes the underlying conneciton.
//func (s *ServerCodec) Close() error {
//log.Println("Close()")
//return s.Gateway.Close()
//}
