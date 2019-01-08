package git.xzf.voip;

public class VoipClient {
    private String address;
    private int port;
    private TcpClient tcpClient;

    public VoipClient(String addr, int p) {
        address = addr;
        port = p;
        tcpClient = new TcpClient(addr, p);
    }

    public void WaitDail() {

    }

}
