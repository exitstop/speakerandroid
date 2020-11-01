package gomobilelib

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
