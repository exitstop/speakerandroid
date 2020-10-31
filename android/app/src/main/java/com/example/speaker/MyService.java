package com.example.speaker;

import android.app.Service;
import android.content.Context;
import android.content.Intent;
import android.net.wifi.WifiManager;
import android.os.Environment;
import android.os.IBinder;
import android.os.PowerManager;
import android.speech.tts.TextToSpeech;
import android.speech.tts.UtteranceProgressListener;
import android.speech.tts.Voice;
import android.util.Log;
import android.widget.Toast;

import java.io.File;
import java.io.IOException;
import java.util.List;
import java.util.Locale;
import java.util.Set;
import java.util.concurrent.Semaphore;

import gomobilelib.Gomobilelib;

public class MyService extends Service {
    //Server server;
    private TextToSpeech mTts = null;
    private TextToSpeech mTts2 = null;
    String text;
    String currentEngine = null;
    File file;
    long startTime = 0;
    int totalPermits = 1;
    static Semaphore semaphore = new Semaphore(1, true);
    int SemCount = 0;
    WifiManager.WifiLock wifiLock = null;
    private PowerManager.WakeLock wakeLock;
    private PowerManager.WakeLock wl;
    private PowerManager pm;
    private String TAG = "MyWakeLock";
    private Context context;

    @Override
    public IBinder onBind(Intent intent) {
        throw new UnsupportedOperationException("Not yet implemented");
        //return null;
    }
    @Override
    public void onCreate() {
        Toast.makeText(this, "Invoke background service onCreate method.", Toast.LENGTH_LONG).show();
        super.onCreate();
    }

    @Override
    public int onStartCommand(Intent intent, int flags, int startId) {
        Toast.makeText(this, "Service started by user.", Toast.LENGTH_LONG).show();

        pm = (PowerManager) getSystemService(Context.POWER_SERVICE);
        wl = pm.newWakeLock(PowerManager.PARTIAL_WAKE_LOCK, TAG);
        wl.acquire();

        Log.e("log------", "i'am created");
        System.out.println("onCreate()");

        Gomobilelib.startServer();

        InitTts("");

        final Locale engLang = new Locale("en", "US");
        final Locale rusLang = new Locale("ru", "RU");

        System.out.println("engLang = " + engLang + " rusLang = " + rusLang);

        Gomobilelib.registerJavaCallBack(new GoCallBack() {
            public long setEngine(String p0) {
                return InitTts(p0);
            };
            // "ru" "RU"
            public long setLanguage(String p0, String p1){
                System.out.println("setLanguage: " + p0 + " " + p1);
                if (p1 != "") {
                    return mTts.setLanguage(new Locale(p0, p1));
                } else {
                    return mTts.setLanguage(new Locale(p0));
                }
            };
            public long setSpeechRate(double p0){
                System.out.println("setSpeechRate: " + p0);
                return mTts.setSpeechRate((float)p0);
            };
            public long setVoice(String p0){
                Set<Voice>	voices = mTts.getVoices();
                for(Voice item : voices){
                    if (p0 == item.getName()) {
                        System.out.println("setVoice: " + item.getName());
                        return mTts.setVoice(item);
                    }
                }
                return -1;
            };

            public long setPitch(double p0){
                return mTts.setPitch((float)p0);
            };

            public String getLocal(){
                String retLocal = "";
                Set<Locale> local = mTts.getAvailableLanguages ();
                for(Locale item : local){
                    retLocal += item + ";";
                    //System.out.println("lang: " + item.getDisplayName() + " " + item.getLanguage());
                }
                return retLocal;
            };
            public String getEngine(){
                String retEngine = "";
                List<TextToSpeech.EngineInfo> engine = mTts.getEngines ();
                for(TextToSpeech.EngineInfo item : engine){
                    retEngine += item + ";";
                    //System.out.println("engine: " + item.name + " " + item.label);
                }
                return retEngine;
            };
            public String getVoice(){
                String retVoice = "";

                //System.out.println("SemCount = "+ semaphore.availablePermits());
                //semaphore.acquire();

                Set<Voice>	voices = mTts.getVoices();
                for(Voice item : voices){
                    retVoice += item + ";";
                    //System.out.println("voices: " + item.getName());
                }

                System.out.println("Release SemCount = " + semaphore.availablePermits());
                if(semaphore.availablePermits() < totalPermits ) {
                    semaphore.release(totalPermits - semaphore.availablePermits());
                }

                return retVoice;
            };

            public void playOnAndroid(String p0) {
                System.out.println("PlayOnAndroid: " + p0);
                //mTts.setLanguage(rusLang);
                mTts.speak(p0, TextToSpeech.QUEUE_FLUSH, null, null);
                //mTts.setLanguage(Locale.US);


            }

            public void speakAdd(String p0) {
                System.out.println("speakAdd: " + p0);
                //mTts.setLanguage(engLang);
                //waitMutex();
                mTts.speak(p0, TextToSpeech.QUEUE_ADD, null, null);
                //mTts.setLanguage(rusLang);
            }

            public void waitMutex(long in) {
                System.out.println(" : acquiring lock...");
                System.out.println(" : available Semaphore permits now: "
                        + semaphore.availablePermits());

                if ( in == 0 ) {
                    try {
                        System.out.println("SemCount = " + semaphore.availablePermits());
                        semaphore.acquire();
                    } catch (InterruptedException e) {
                        e.printStackTrace();
                    }
                }

                try{
                    System.out.println("SemCount = "+ semaphore.availablePermits());
                    semaphore.acquire();
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }

                System.out.println(" : got the permit!");

                System.out.println("Release SemCount = " + semaphore.availablePermits());
                if(semaphore.availablePermits() < totalPermits ) {
                    semaphore.release(totalPermits - semaphore.availablePermits());
                }
            }

            public String writeInFile(String p0, String fileName) {

                System.out.println("Release SemCount = " + semaphore.availablePermits());
                if(semaphore.availablePermits() < totalPermits ) {
                    semaphore.release(totalPermits - semaphore.availablePermits());
                }

                class OneShotTask implements Runnable {
                    String str;
                    String p0_;
                    String fileName_;
                    OneShotTask(String p0, String fileName) {
                        p0_ = p0;
                        fileName_ = fileName;
                    }
                    public void run() {
                        System.out.println("run() + ");
                        System.out.println("WriteInFile: " + p0_);

                        file = new File("");
                        if ( flagRam_ == 0 ) {
                            file = new File(getExternalFilesDir(Environment.DIRECTORY_DOWNLOADS), fileName_);
                            file.getParentFile().mkdirs();
                            System.out.println("file " + file);
                        } else {
                            file = new File("/storage/emulated/0/RAM Disk/" + fileName_);
                            System.out.println("file " + file);
                        }

                        //startTime = System.nanoTime();

                        mTts.synthesizeToFile(p0_, null, file, fileName_);
                        System.out.println("synthesizeToFile + ");
                    }
                }

                Thread thread = new Thread(new OneShotTask(p0,fileName));
                thread.start();

                if ( flagRam_ != 0 ) {
                    return "/storage/emulated/0/RAM Disk/";
                }
                return getExternalFilesDir(Environment.DIRECTORY_DOWNLOADS) + "/";
            }

            public void flagRam(long p0) {
                flagRam_ = p0;
            }

        });


        return super.onStartCommand(intent, flags, startId);
        //return START_STICKY;
    }
    private int InitTts(String p0) {

        if( currentEngine != null && currentEngine.equals(p0) ) {
            return -1;
        }
        if (currentEngine != null) {
            //mTts.shutdown();
        }
        if (p0 == "") {
            p0 = "com.google.android.tts";
        }

        if(semaphore.availablePermits() < totalPermits ) {
            System.out.println("----SemCount = "+ semaphore.availablePermits());
            semaphore.release(totalPermits - semaphore.availablePermits());
            System.out.println("----SemCount = "+ semaphore.availablePermits());
        }


        System.out.println("last currentEngine " + currentEngine);
        currentEngine = p0;
        System.out.println("set currentEngine " + currentEngine);

        try{
            System.out.println("SemCount = "+ semaphore.availablePermits());
            semaphore.acquire();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }


        mTts = new TextToSpeech(this, new TextToSpeech.OnInitListener() {
            @Override
            public void onInit(int status) {
                // TODO Auto-generated method stub
                if (status == TextToSpeech.SUCCESS) {
                    //int result=mTts.setLanguage(Locale.US);
                    int result = mTts.setLanguage(new Locale("ru", "RU"));
                    if (result == TextToSpeech.LANG_MISSING_DATA ||
                            result == TextToSpeech.LANG_NOT_SUPPORTED) {
                        Log.e("error", "This Language is not supported");
                        return;
                    } else {
                        System.out.println("--Set tts by default Ok: ");
                    }
                } else {
                    Log.e("error", "Initilization Failed!");
                }

                // com.google.android.tts
                // com.acapelagroup.android.tts

                float speechRate = 1.0f;
                mTts.setSpeechRate(speechRate);
                final String initEngine = mTts.getDefaultEngine();
                System.out.println(initEngine);
                final int maxSpeechInput = mTts.getMaxSpeechInputLength();
                System.out.println("maxSpeechInput: " + maxSpeechInput);
                //mTts.setEngineByPackageName()

                mTts.setOnUtteranceProgressListener(new UtteranceProgressListener() {
                    @Override
                    public void onStart(String utteranceId) {
                        long estimatedTime = (System.nanoTime() - startTime)/1000000;
                        System.out.println("onStart("+utteranceId+") - " + estimatedTime);
                        // Speaking started.

                    }

                    @Override
                    public void onDone(String utteranceId) {
                        long estimatedTime = (System.nanoTime() - startTime)/1000000;
                        // Speaking stopped.
                        System.out.println("onDone("+utteranceId+") - " + estimatedTime);
//                        Hello.zreetings(getExternalFilesDir(Environment.DIRECTORY_DOWNLOADS) + utteranceId);

                        System.out.println("Release SemCount = " + semaphore.availablePermits());
                        if(semaphore.availablePermits() < totalPermits ) {
                            semaphore.release(totalPermits - semaphore.availablePermits());
                        }

                    }

                    @Override
                    public void onError(String utteranceId) {
                        long estimatedTime = (System.nanoTime() - startTime)/1000000;
                        // There was an error.
                        System.out.println("onError - " + estimatedTime);

                        System.out.println("Release SemCount = " + semaphore.availablePermits());
                        if(semaphore.availablePermits() < totalPermits ) {
                            semaphore.release(totalPermits - semaphore.availablePermits());
                        }

                    }
                });

                System.out.println("Release SemCount = " + semaphore.availablePermits());
                if(semaphore.availablePermits() < totalPermits ) {
                    semaphore.release(totalPermits - semaphore.availablePermits());
                }
            }
        }, currentEngine);
        return 0;
    }


    @Override
    public void onDestroy() {
        wl.release();
        System.out.println("Release SemCount = " + semaphore.availablePermits());
        if(semaphore.availablePermits() < totalPermits ) {
            semaphore.release(totalPermits - semaphore.availablePermits());
        }

        mTts.shutdown();
        /*
        if (wakeLock != null) {
            if (wakeLock.isHeld()) {
                wakeLock.release();
            }
        }
        if (wifiLock != null) {
            if (wifiLock.isHeld()) {
                wifiLock.release();
            }
        }
         */
        Toast.makeText(this, "Service destroyed by user.", Toast.LENGTH_LONG).show();
        super.onDestroy();
    }
}
