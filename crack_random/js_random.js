function guessNumber() {
    while(1) {
        const prompt = require("prompt-sync")({sigint: true});
        const secretNumber = Math.random().toString().replace('0.','');
        let guess = parseInt(prompt('Enter your guess '));

        if (guess==secretNumber) {
            console.log("YES !!!!!!!! You devine my number");
            break;
        } else{
            console.log("Nope, This is the secret number " + secretNumber);
        }
    }
}

guessNumber()
