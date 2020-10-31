package com.example.speaker;

import android.speech.tts.TextToSpeech;

import gomobilelib.JavaCallBack;

public class GoCallBack implements JavaCallBack{
    public TextToSpeech mTts_;
    public long flagRam_ = 0;
    public void playOnAndroid(String p0){
        System.out.println("PlayOnAndroid: " + p0);
    }
    public String writeInFile(String p0, String filenName){
        System.out.println("WriteInFile: " + p0);
        return "";
    }
    public void flagRam(long p0) {}
    public void waitMutex(long in){
        System.out.println("waitMutex: ");
    }
    public String getEngine(){return "";};
    public String getLocal(){return "";};
    public String getVoice(){return "";};
    public long setEngine(String p0) { return 0;};
    public long setLanguage(String p0, String p1){ return 0;};
    public long setSpeechRate(double p0){ return 0;};
    public long setVoice(String p0){ return 0;};
    public long setPitch(double p0) {return 0;};
    public void speakAdd(String p0) {};
}

