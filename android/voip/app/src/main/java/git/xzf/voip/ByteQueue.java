package git.xzf.voip;

import java.util.ArrayList;
import java.util.List;

public class ByteQueue {
    private Lock lock;
    private List<byte[]> queue;

    public ByteQueue() {
        queue = new ArrayList<>();
        lock = new Lock();
    }

    public void In(byte[] data) {
        try {
            lock.lock();
            queue.add(data);
            lock.unlock();
        } catch (Exception e) {
            Log.P("ByteQueue.In", e.toString());
        }
    }

    public byte[] Out() {
        if (queue.size() != 0) {
            try {
                lock.lock();
                byte[] tmp = queue.get(0).clone();
                queue.remove(0);
                lock.unlock();
                return tmp;
            } catch (Exception e) {
                Log.P("ByteQueue.Out", e.toString());
            }
        }
        return new byte[0];
    }

    public void Clear() {
        try {
            lock.lock();
            for (int i = 0; i < queue.size(); i++) {
                queue.remove(0);
            }
            lock.unlock();
        } catch (Exception e) {
            Log.P("ByteQueue.Out", e.toString());
        }

    }
}
