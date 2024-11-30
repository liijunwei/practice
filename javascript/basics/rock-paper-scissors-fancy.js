<!DOCTYPE html>
<html>

<head>
  <title>rock paper scissors - fancy</title>
</head>

<style>
  body {
    background-color: black;
    color: white;
    font-family: Arial;
  }

  .title {
    font-size: 30px;
    font-weight: bold;
  }

  .move-icon {
    height: 50px;
  }

  .move-button {
    background-color: transparent;
    border: 3px solid white;
    border-radius: 50%;
    width: 120px;
    height: 120px;
    cursor: pointer;
  }

  .result {
    font-size: 25px;
    font-weight: bold;
    margin-top: 50px;
  }

  .score {
    margin-top: 50px;
  }

  .reset-score-button {
    background-color: white;
    border: none;
    padding: 8px 15px;
    font-size: 25px;
    cursor: pointer;
  }
</style>


<p class="title">rock paper scissors</p>

<body>
  <!-- image asset: https://supersimple.dev/projects/rock-paper-scissors/ -->
  <button onclick="playGame('rock')" class="move-button"><img src="./rock-emoji.png" class="move-icon"></button>
  <button onclick="playGame('paper')" class="move-button"><img src="./paper-emoji.png" class="move-icon"></button>
  <button onclick="playGame('scissors')" class="move-button"><img src="./scissors-emoji.png" class="move-icon"></button>
  <p class="js-result result"></p>
  <p class="js-moves"></p>
  <p class="js-scores score"></p>
  <button onclick="resetScores()" class="reset-score-button">reset</button>
  <div id="js-table-container"></div>

</body>

<script>
  updateScoreElement()

  function resetScores() {
    localStorage.removeItem('scores')
    localStorage.removeItem('game_logs')
    updateScoreElement()
  }

  function getComputerMove() {
    const num = Math.random()
    let computerMove = ''

    if (num >= 0 && num < 1 / 3) {
      computerMove = 'rock';
    } else if (num >= 1 / 3 && num < 2 / 3) {
      computerMove = 'paper';
    } else if (num >= 2 / 3 && num < 1) {
      computerMove = 'scissors';
    }

    console.assert(computerMove == 'rock' || computerMove == 'paper' || computerMove == 'scissors')

    return computerMove;
  }

  function playGame(playerMove) {
    console.assert(playerMove == 'rock' || playerMove == 'paper' || playerMove == 'scissors')

    const computerMove = getComputerMove()
    console.assert(computerMove == 'rock' || computerMove == 'paper' || computerMove == 'scissors')

    let result = ''

    if (playerMove == 'rock') {
      if (computerMove == 'rock') {
        result = 'tie';
      } else if (computerMove == 'paper') {
        result = 'lose';
      } else if (computerMove == 'scissors') {
        result = 'win';
      }
    } else if (playerMove == 'paper') {
      if (computerMove == 'rock') {
        result = 'win';
      } else if (computerMove == 'paper') {
        result = 'tie';
      } else if (computerMove == 'scissors') {
        result = 'lose';
      }
    } else if (playerMove == 'scissors') {
      if (computerMove == 'rock') {
        result = 'lose';
      } else if (computerMove == 'paper') {
        result = 'win';
      } else if (computerMove == 'scissors') {
        result = 'tie';
      }
    }

    console.assert(result == 'win' || result == 'lose' || result == 'tie')

    const detail = {
      created_at: new Date().toISOString(),
      playerMove: playerMove,
      computerMove: computerMove,
      result: result,
    }

    updateScores(detail)
  }

  function updateScores(detail) {
    let gameResult = detail.result
    console.assert(gameResult == 'win' || gameResult == 'lose' || gameResult == 'tie')

    let playerMove = detail.playerMove;
    let computerMove = detail.computerMove;

    let scores = getScores()

    if (gameResult == 'win') {
      scores.wins += 1
    } else if (gameResult == 'lose') {
      scores.loses += 1
    } else if (gameResult == 'tie') {
      scores.ties += 1
    }

    let winRate = (scores.wins / (scores.wins + scores.loses + scores.ties)).toFixed(2)
    scores.win_rate = winRate

    data = JSON.stringify(scores)
    localStorage.setItem('scores', data)

    updateScoreElement()
    document.querySelector('.js-result').innerHTML = `You ${gameResult}.`;
    document.querySelector('.js-moves').innerHTML = `
      You
      <img src="./${playerMove}-emoji.png" class="move-icon">
      <img src="./${computerMove}-emoji.png" class="move-icon">
      Computer`;
  }

  function updateScoreElement() {
    document.querySelector('.js-scores').innerHTML = getRawScores()
  }

  // global state... unsafe here
  function getScores() {
    return JSON.parse(getRawScores())
  }

  function getRawScores() {
    return localStorage.getItem('scores') || JSON.stringify(getEmptyScores())
  }

  function getEmptyScores() {
    return {
      wins: 0,
      loses: 0,
      ties: 0,
      win_rate: 0,
    }
  }
</script>

</html>
