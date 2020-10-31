package com.example.speaker;

import androidx.appcompat.app.AppCompatActivity;

import android.os.Bundle;
import android.widget.TextView;
import android.content.Intent;

public class MainActivity extends AppCompatActivity {
    private TextView mTextView;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        Utils.getMACAddress("wlan0");
        Utils.getMACAddress("eth0");
        Utils.getIPAddress(true); // IPv4
        Utils.getIPAddress(false); // IPv6

        mTextView = (TextView) findViewById(R.id.mytextview);
        mTextView.setText("ip: " + Utils.getIPAddress(true));

        startService(new Intent(MainActivity.this, MyService.class));
    }
    @Override
    public void onSaveInstanceState(Bundle outState) {
        super.onSaveInstanceState(outState);
    }

    @Override
    protected void onResume() {
        System.out.println("onResume()");
        super.onResume();
    }
    @Override
    protected void onPause() {
        System.out.println("onPause()");
        super.onPause();
    }
    @Override
    protected void onStop() {
        super.onStop();
        System.out.println("onStop()");
    }

    @Override
    protected void onDestroy() {
        stopService(new Intent(MainActivity.this,MyService.class));
        System.out.println("onDestroy()");
        super.onDestroy();
    }
}