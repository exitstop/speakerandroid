package gomobilelib

import (
	"fmt"
	"strings"

	"github.com/kataras/iris/v12"
)

func (s *Server) handlePlayOnAndroid(ctx iris.Context) {
	type TextStruct struct {
		Text string `json:"Text"`
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

	ctx.WriteString("ret PlayOnAndroid: " + text.Text) //nolint:errcheck
	ctx.StatusCode(iris.StatusOK)
	return
}

func (s *Server) handleGetLocal(ctx iris.Context) {
	reply := s.Jc.GetLocal()
	ctx.WriteString(reply) //nolint:errcheck
	ctx.StatusCode(iris.StatusOK)
}

func (s *Server) handleGetEngine(ctx iris.Context) {
	reply := s.Jc.GetEngine()
	ctx.WriteString(reply) //nolint:errcheck
	ctx.StatusCode(iris.StatusOK)
}

func (s *Server) handleGetVoice(ctx iris.Context) {
	reply := s.Jc.GetVoice()
	ctx.WriteString(reply) //nolint:errcheck
	ctx.StatusCode(iris.StatusOK)
}

func (s *Server) handleSetVoice(ctx iris.Context) {
	type TextStruct struct {
		Text string `json:"Text"`
	}

	text := TextStruct{}

	err := ctx.ReadJSON(&text)

	if err != nil {
		fmt.Println(err)
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error()) //nolint:errcheck
		return
	}

	s.Jc.SetVoice(text.Text)

	ctx.WriteString("ret SetVoice: " + text.Text) //nolint:errcheck
	ctx.StatusCode(iris.StatusOK)
	return
}

func (s *Server) handleSetLanguage(ctx iris.Context) {
	type TextStruct struct {
		Text string `json:"Text"`
	}

	text := TextStruct{}

	err := ctx.ReadJSON(&text)

	if err != nil {
		fmt.Println(err)
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error()) //nolint:errcheck
		return
	}

	ar := strings.Split(text.Text, "-")
	if len(ar) == 2 {
		//log.Println("one ", ar[0], " two ", ar[1])
		if len(ar) == 2 {
			_ = s.Jc.SetLanguage(ar[0], ar[1])
		} else {
			_ = s.Jc.SetLanguage(text.Text, "")
		}
	}

	ctx.WriteString("ret SetLanguage: " + text.Text) //nolint:errcheck
	ctx.StatusCode(iris.StatusOK)
	return
}

func (s *Server) handleSetSpeechRate(ctx iris.Context) {
	type TextStruct struct {
		SpeechRate float64 `json:"SpeechRate"`
	}

	text := TextStruct{}

	err := ctx.ReadJSON(&text)

	if err != nil {
		fmt.Println(err)
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error()) //nolint:errcheck
		return
	}

	s.Jc.SetSpeechRate(text.SpeechRate)

	ctx.WriteString("ret SetSpeechRate: ok") //nolint:errcheck
	ctx.StatusCode(iris.StatusOK)
	return
}

func (s *Server) handleSetPitch(ctx iris.Context) {
	type TextStruct struct {
		SetPitch float64 `json:"SetPitch"`
	}

	text := TextStruct{}

	err := ctx.ReadJSON(&text)

	if err != nil {
		fmt.Println(err)
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error()) //nolint:errcheck
		return
	}

	s.Jc.SetPitch(text.SetPitch)

	ctx.WriteString("ret SetSpeechRate: ok") //nolint:errcheck
	ctx.StatusCode(iris.StatusOK)
	return
}

func (s *Server) handleHelloAndroid(ctx iris.Context) {
	ctx.WriteString("Hello i'am Android") //nolint:errcheck
	ctx.StatusCode(iris.StatusOK)
}

func (s *Server) handleSetEngine(ctx iris.Context) {
	type TextStruct struct {
		Text string `json:"Text"`
	}

	text := TextStruct{}

	err := ctx.ReadJSON(&text)

	if err != nil {
		fmt.Println(err)
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error()) //nolint:errcheck
		return
	}

	s.Jc.SetEngine(text.Text)

	ctx.WriteString("ret SetEngine: " + text.Text) //nolint:errcheck
	ctx.StatusCode(iris.StatusOK)
	return
}

func (s *Server) handleWaitMutex(ctx iris.Context) {
	type TextStruct struct {
		WaitMutex int `json:"WaitMutex"`
	}

	text := TextStruct{}

	err := ctx.ReadJSON(&text)

	if err != nil {
		fmt.Println(err)
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error()) //nolint:errcheck
		return
	}

	s.Jc.WaitMutex(text.WaitMutex)

	ctx.WriteString("Wait Ok") //nolint:errcheck
	ctx.StatusCode(iris.StatusOK)
	return
}
