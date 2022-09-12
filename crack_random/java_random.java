import java.util.Random;
import java.util.Scanner;


public class GuestNumber {
    public static void main(String args[]) {
        Random random = new Random();
        Scanner scan = new Scanner(System.in);
        System.out.println("Welcom to this random number guessing");

        while(true) {
            System.out.print("Guest my secret number : ");
            int guess = scan.nextInt();
            int secret = random.nextInt();
            if (secret == guess) {
                System.out.println("YES !!!!!!!! You devine my number");
            } else {
                System.out.println("Nope, This is the secret number " + secret);
            }
        }
    }
}
