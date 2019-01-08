package git.xzf.voip;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.widget.Button;
import android.widget.TextView;

import java.util.Random;

public class MainActivity extends AppCompatActivity {
    Audio audio;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        Button dailBut = findViewById(R.id.DailBut);
        Button stopBut = findViewById(R.id.StopBut);
        final TextView userNameView = findViewById(R.id.UserNameView);
        String tmpUserName = "";
        for (int i = 0; i < 6; i++) {
            tmpUserName += new Random().nextInt(10) + "";
        }
        new Thread(new Runnable() {
            @Override
            public void run() {
                Log.P("init Audio ");
                Log.P("init Audio success");
            }
        }).start();
        userNameView.setText(tmpUserName);
        dailBut.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Log.P("onClick dailBut");
                String userName = userNameView.getText().toString();
                new Thread(new Runnable() {
                    @Override
                    public void run() {
                        try {
                            audio = new Audio();
                            audio.Record();
                        } catch (Exception e) {
                            Log.P("onClick Audio new " + e);
                        }
                    }
                }).start();
            }
        });
        stopBut.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Log.P("onClick stopBut");
                audio.StopRecord();
            }
        });
    }
}
