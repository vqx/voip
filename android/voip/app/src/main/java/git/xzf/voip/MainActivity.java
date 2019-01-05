package git.xzf.voip;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.widget.Button;
import android.widget.EditText;
import android.widget.TextView;

import java.io.BufferedReader;
import java.io.PrintWriter;
import java.net.Socket;

public class MainActivity extends AppCompatActivity {
    Socket socket;
    PrintWriter out = null;
    BufferedReader in = null;
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        final EditText addressView = findViewById(R.id.AddressView);
        final String address = addressView.getText().toString();
        Button sendBut = findViewById(R.id.SendBut);
        final EditText sendTextView = findViewById(R.id.SendTextView);
        final TextView showTextView = findViewById(R.id.ShowTextView);
        final String sendText = sendTextView.getText().toString();

        sendBut.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Log.P("onClick sendBut");
                new Thread(new Runnable() {
                    @Override
                    public void run() {
                        try {
                            new Audio().Record();
                        } catch (Exception e) {
                            Log.P("onClick Audio new " + e);
                        }
                    }
                }).start();
            }
        });
    }
}
