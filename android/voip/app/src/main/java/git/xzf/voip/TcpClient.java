package git.xzf.voip;

import java.io.ByteArrayOutputStream;
import java.io.InputStream;
import java.io.OutputStream;
import java.net.Socket;

public class TcpClient {
    String Address;
    int port;
    Socket socket;
    Lock lock;
    InputStream inputStream;
    OutputStream outputStream;

    public TcpClient(String addr, int p) {
        Address = addr;
        port = p;
    }

    public void Connect() {
        try {
            this.socket = new Socket(Address, port);
        } catch (Exception e) {
            Log.P("TcpClient", "Send", e.toString());
        }
    }

    public void Send(byte[] data) {
        try {
            if (outputStream == null) {
                outputStream = socket.getOutputStream();
            }
            outputStream.write(data);
        } catch (Exception e) {
            Log.P("TcpClient", "Send", e.toString());
        }
    }

    public byte[] Get(long length) {
        try {
            if (inputStream == null) {
                inputStream = socket.getInputStream();
            }
            ByteArrayOutputStream bo = new ByteArrayOutputStream();// 建立字节流
            byte[] buffer = new byte[1024];//1024长度
            int read = 0;
            while (read < length) {// 循环将需要读取内容写入到bo中
                int cur = inputStream.read(buffer, 0, (int) Math.min(1024, length - read));
                if (cur < 0) {//直到读到的返回标记为-1，表示读到流的结尾
                    break;
                }
                read += cur;//每次读取的长度累加
                bo.write(buffer, 0, cur);
            }
            return bo.toByteArray();//返回内容

        } catch (Exception e) {
            Log.P("TcpClient", "Get", e.toString());
        }
        return null;
    }
}
