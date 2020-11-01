package server

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

func (s *Server) handlePlayOnAndroid(ctx iris.Context) {
	type TextStruct struct {
		Text string `json:"text"`
	}

	text := TextStruct{}

	err := ctx.ReadJSON(&text)

	if err != nil {
		fmt.Println(err)
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error()) //nolint:errcheck
		return
	}

	s.Jc.PlayOnAndroid(text.Text)

	ctx.StatusCode(iris.StatusOK)
	return
}

func (s *Server) handleGetLocal(ctx iris.Context) {
}

func (s *Server) handleGetEngine(ctx iris.Context) {
}

func (s *Server) handleGetVoice(ctx iris.Context) {
}

func (s *Server) handleSetVoice(ctx iris.Context) {
}

func (s *Server) handleSetLanguage(ctx iris.Context) {
}

func (s *Server) handleSetSpeechRate(ctx iris.Context) {
}

func (s *Server) handleSetPitch(ctx iris.Context) {
}

func (s *Server) handleHelloAndroid(ctx iris.Context) {
}

func (s *Server) handleSetEngine(ctx iris.Context) {
}

func (s *Server) handleWaitMutex(ctx iris.Context) {
}
