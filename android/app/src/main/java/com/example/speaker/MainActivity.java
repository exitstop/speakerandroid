package com.example.speaker;

import androidx.appcompat.app.AppCompatActivity;

import android.os.Bundle;
import android.widget.TextView;
import java.util.*;

import gomobilelib.Gomobilelib;


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

//        TextView myAwesomeTextView = (TextView)findViewById(R.id.myAwesomeTextView);
//        myAwesomeTextView.setText("ip: " + Utils.getIPAddress(true));

        mTextView = (TextView) findViewById(R.id.mytextview);
        mTextView.setText("ip: " + Utils.getIPAddress(true));

//        String greetings = Gomobilelib.greetings("Android and Gopher");

    }
}