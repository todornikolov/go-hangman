### Go game of Hangman

Put a letter to guess a word in an amount of turns.

#### Todo

- [x] Read and parse words from txt file, every word on a new line containing a-z letters.
- [x] Show guessed chars or show missing letters as dashes.
- [x] Win or loose game based on left letters or turns.
- [x] Draw hangman after each guess turn.

#### How to run

Pull the repository and inside the project directory run:

`docker build --tag go-hangman:0.1 -f ./docker/app/Dockerfile --rm .`

Run the container with the following:

`docker run -i -t go-hangman:0.1 scripts/run-app.sh`
