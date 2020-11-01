package gomobilelib

import (
	"github.com/kataras/iris/v12"
)

type Server struct {
	Router *iris.Application
	Jc     JavaCallBack
	//config *argument.Config
	//store  store.Store
}

func (s *Server) ServerAddJavaCallBack(jc JavaCallBack) {
	s.Jc = jc
}

func NewServer() *Server {
	s := &Server{
		Router: iris.New(),
		//config: config,
	}
	return s
}

func Start(serv *Server) (err error) {
	return serv.Router.Run(
		iris.Addr(":8484"),
		iris.WithoutPathCorrection,
		iris.WithoutServerError(iris.ErrServerClosed),
	)
}

func Run() error {
	s := NewServer()
	s.ConfigureRouter()
	return Start(s)
}
