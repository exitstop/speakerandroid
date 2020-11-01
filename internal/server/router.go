package server

func (s *Server) ConfigureRouter() {
	// PlayOnAndroid
	s.Router.Post("/play_on_android", s.handlePlayOnAndroid)
	// GetLocal
	s.Router.Post("/get_local", s.handleGetLocal)
	// GetEngine
	s.Router.Post("/get_engine", s.handleGetEngine)
	// GetVoice
	s.Router.Post("/get_voice", s.handleGetVoice)
	// SetVoice
	s.Router.Post("/set_voice", s.handleSetVoice)
	// SetLanguage
	s.Router.Post("/set_language", s.handleSetLanguage)
	// SetSpeechRate
	s.Router.Post("/set_speech_rate", s.handleSetSpeechRate)
	// SetPitch
	s.Router.Post("/set_pitch", s.handleSetPitch)
	// HelloAndroid
	s.Router.Post("/hello_android", s.handleHelloAndroid)
	// SetEngine
	s.Router.Post("/set_engine", s.handleSetEngine)
	// WaitMutex
	s.Router.Post("/wait_mutex", s.handleWaitMutex)

	// ------------------------------------
	// WriteInFile
	//s.Router.Post("/writeIn_file", s.handleWriteInFile)
}
