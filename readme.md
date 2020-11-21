### Go game of Hangman version 0.1

Put a letter to guess a word in an amount of turns.

#### Todo

- [ ] Read and parse words from txt file, every word on a new line containing a-z letters.
- [x] Show guessed chars or show missing letters as dashes.
- [ ] Draw hangman after each guess turn.
- [x] Win or loose game based on left letters or turns.

#### How to run

`docoker-compose build`

`docker run -i -t gogame_app scripts/run-app.sh`