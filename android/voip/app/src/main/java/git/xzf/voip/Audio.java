package git.xzf.voip;

import android.media.AudioFormat;
import android.media.AudioManager;
import android.media.AudioRecord;
import android.media.AudioTrack;
import android.media.MediaRecorder;

public class Audio {
    int BufferElements2Rec = 1024; // want to play 2048 (2K) since 2 bytes we use only 1024
    int BytesPerElement = 2; // 2 bytes in 16bit format
    private static final int RECORDER_SAMPLERATE = 44100;
    private static final int RECORDER_CHANNELS = AudioFormat.CHANNEL_IN_MONO;
    private static final int RECORDER_AUDIO_ENCODING = AudioFormat.ENCODING_PCM_16BIT;
    AudioRecord record;
    TcpClient tcpClient;
    AudioTrack audioTrack;
    ByteQueue queue;

    public Audio() {
        queue = new ByteQueue();
        record = new AudioRecord(MediaRecorder.AudioSource.DEFAULT,
                RECORDER_SAMPLERATE,
                RECORDER_CHANNELS,
                RECORDER_AUDIO_ENCODING,
                BufferElements2Rec * BytesPerElement);
        audioTrack = new AudioTrack(
                AudioManager.STREAM_VOICE_CALL,
                RECORDER_SAMPLERATE,
                AudioFormat.CHANNEL_OUT_MONO,
                AudioFormat.ENCODING_PCM_16BIT,
                BufferElements2Rec * BytesPerElement,
                AudioTrack.MODE_STREAM);
        // tcpClient = new TcpClient(Config.Address, Config.Port);
    }

    boolean recording;

    public void Record() throws Exception {
        Log.P("enter");
        // tcpClient.Connect();
        record.startRecording();
        recording = true;
        new Thread(new Runnable() {
            @Override
            public void run() {
                Log.P("start record thread ");
                while (recording) {
                    byte sData[] = new byte[BufferElements2Rec];
                    record.read(sData, 0, BufferElements2Rec);
                    queue.In(sData);
                    //  tcpClient.Send(sData);

                }
            }
        }).start();
        new Thread(new Runnable() {
            @Override
            public void run() {
                Log.P("start play thread ");
                while (recording) {
                    byte[] tmp = queue.Out();
                    if (tmp.length != 0) {
                        audioTrack.write(tmp, 0, BufferElements2Rec);
                        audioTrack.play();
                    }
                }
            }
        }).start();
        Log.P("end run Thread");
    }

    public void StopRecord() {
        if (null != record) {
            recording = false;
            record.stop();
            record.release();
        }
    }
}
