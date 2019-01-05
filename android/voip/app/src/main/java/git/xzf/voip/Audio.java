package git.xzf.voip;

import android.media.AudioFormat;
import android.media.AudioRecord;
import android.media.MediaRecorder;

public class Audio {
    int BufferElements2Rec = 1024; // want to play 2048 (2K) since 2 bytes we use only 1024
    int BytesPerElement = 2; // 2 bytes in 16bit format
    private static final int RECORDER_SAMPLERATE = 8000;
    private static final int RECORDER_CHANNELS = AudioFormat.CHANNEL_IN_MONO;
    private static final int RECORDER_AUDIO_ENCODING = AudioFormat.ENCODING_PCM_16BIT;
    AudioRecord record;
    TcpClient tcpClient;

    public Audio() {
        record = new AudioRecord(MediaRecorder.AudioSource.MIC,
                RECORDER_SAMPLERATE, RECORDER_CHANNELS,
                RECORDER_AUDIO_ENCODING, BufferElements2Rec * BytesPerElement);
        tcpClient = new TcpClient(Config.Address, Config.Port);
    }

    boolean recording;

    public void Record() {
        Log.P("enter");
        tcpClient.Connect();
        record.startRecording();
        recording = true;
        while (recording) {
            try {
                byte sData[] = new byte[BufferElements2Rec];
                record.read(sData, 0, BufferElements2Rec);
                tcpClient.Send(sData);

//                for (int i = 0; i < BufferElements2Rec; i++) {
//                    String str = "";
//                    for (int j = 0; j < (int) sData[i]; j++) {
//                        str += "-";
//                    }
//                    Log.P(str);
//                }
            } catch (Exception e) {
                Log.P("error", e.toString());

            }
        }
        Log.P("out");
    }

    public void StopRecord() {
        if (null != record) {
            recording = false;
            record.stop();
            record.release();
        }
    }
}
