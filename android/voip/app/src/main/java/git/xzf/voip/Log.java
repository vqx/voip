package git.xzf.voip;

public class Log {
    public static void P(String... str) {
        System.out.print("53w5tebet9 ");
        for (int i = 0; i < str.length; i++) {
            System.out.print(str[i]);
            if (i != str.length - 1) {
                System.out.print(" ");
            } else {
                System.out.println();
            }
        }
    }
}
