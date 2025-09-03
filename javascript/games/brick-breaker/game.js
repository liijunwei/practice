let board;
let boardheight = 500;
let boardwidth = 500;
let context;

let playerwidth = 80;
playerwidth = boardwidth; // for testing
let playerheight = 10;

let playervelocity = 10
let player = {
  x: boardwidth/2-playerwidth/2,
  y: boardheight-playerheight-5,
  width: playerwidth,
  height: playerheight,
  velx: playervelocity,
}

let ballwidth = 11;
let ballheight = 10;
let ballvelicityX = getRandomInt(15, 20);
let ballvelicityY = getRandomInt(10, 15);

let ball = {
  x: boardwidth/2,
  y: boardheight/2,
  width: ballwidth,
  height: ballheight,
  velx: ballvelicityX,
  vely: ballvelicityY,
}

// starting top left block
let blockX = 15;
let blockY = 45;

let blocks = [];
let blockwidth = 50;
let blockheight = 10;
let blockcolumns = 8;
let blockrows = 6;
let blockmaxrows = 10;
let blockcount = 0;

let score = 0;
let gameover = false;

window.onload = function() {
  board = document.getElementById("board");
  board.height = boardheight
  board.width = boardwidth
  context = board.getContext("2d")
  context.fillStyle = "pink"
  context.fillRect(player.x, player.y, player.width, player.height)

  // https://developer.mozilla.org/en-US/docs/Web/API/DedicatedWorkerGlobalScope/requestAnimationFrame
  requestAnimationFrame(update)

  document.addEventListener("keydown", movePlayer)

  createblocks()
}

function update() {
  requestAnimationFrame(update)
  if(gameover) {
    return;
  }
  context.clearRect(0, 0, board.width, board.height)

  context.fillStyle = "pink"
  context.fillRect(player.x, player.y, player.width, player.height)

  context.fillStyle = "white"
  ball.x += ball.velx
  ball.y += ball.vely
  context.fillRect(ball.x, ball.y, ball.width, ball.height)

  // bounce ball off walls
  // if(ball.y <= 0 || (ball.y+ball.height)>=boardheight) {
  if(ball.y <= 0) {
    ball.vely *= -1
  }
  if(ball.x <=0 || (ball.x+ball.width)>=boardwidth){
    ball.velx *= -1
  }

  if((ball.y+ball.height)>=boardheight) {
    context.font = "20px sans-serif"
    context.fillText("game over: press 'space' to restart", 80, 250)
    gameover = true;
  }

  if(topcollision(ball, player) || bottomcollision(ball, player)) {
    ball.vely *= -1
  } else if(leftcollision(ball, player) || rightcollision(ball, player)) {
    ball.velx *= -1
  }

  context.fillStyle = "skyblue"
  for(let i = 0; i < blocks.length; i++) {
    let block = blocks[i]
    if(!block.break) {
      if(topcollision(ball, block) || bottomcollision(ball, block)) {
        block.break = true
        ball.vely *= -1
        blockcount--
        score += 1
      } else if(leftcollision(ball, block) || rightcollision(ball, block)) {
        block.break = true
        ball.velx *= -1
        blockcount--
        score += 1
      }
      context.fillRect(block.x, block.y, block.width, block.height)
    }
  }

  if(blockcount == 0) {
    score += 100
    blockrows = Math.min(blockrows+1, blockmaxrows)
    createblocks();
  }

  context.font = "20px sans-serrf"
  context.fillText(score, 10, 20)
}

function outofbound(xpos) {
  return xpos < 0 || xpos + playerwidth > boardwidth
}

function movePlayer(event) {
  if(gameover && event.code == "Space") {
    resetGame();
  }

  if (event.code == "Space") {
    console.log("intervene ball.velx...", ball)
    if (ball.velx > 0) {
      ball.velx++
    } else {
      ball.velx--
    }
  }

  if (event.code == "ArrowLeft") {
    let nextPlayerX = player.x - player.velx
    if(outofbound(nextPlayerX)) {
      console.warn("out of left boundary")
    } else {
      player.x = nextPlayerX
    }
  }
  if (event.code == "ArrowRight") {
    let nextPlayerX = player.x + player.velx
    if(outofbound(nextPlayerX)) {
      console.warn("out of right boundary")
    } else {
      player.x = nextPlayerX
    }
  }
}

// check interctions between two rectangles
function detectCollision(a, b) {
  return a.x < b.x + b.width  && // a's top left corner doesn't reach b's top right corner
         a.x + a.width > b.x  && // a's top right corner passes b's top left corner
         a.y < b.y + b.height && // a's top left corner doesn't reach b's bottom left corner
         a.y + a.height > b.y;   // a's bottom left corner passes b's top left corner
}

function topcollision(ball, block) {
  return detectCollision(ball, block) && (ball.y + ball.height) >= block.y;
}

function bottomcollision(ball, block) {
  return detectCollision(ball, block) && (block.y+block.height) >= ball.y;
}

function leftcollision(ball, block) {
  return detectCollision(ball, block) && (ball.x + ball.width) >= block.x;
}

function rightcollision(ball, block) {
  return detectCollision(ball, block) && (block.x+block.width) >= ball.x;
}

function createblocks() {
  blocks = []
  for(let c = 0; c < blockcolumns; c++) {
    for(let r = 0; r < blockrows; r++) {
      let block = {
        x: blockX+c*blockwidth + c*10,
        y: blockY+r*blockheight + r*10,
        width: blockwidth,
        height: blockheight,
        break: false,
      }
      blocks.push(block)
    }
  }
  blockcount = blocks.length
}

function resetGame() {
  gameover = false;
  player = {
    x: boardwidth/2-playerwidth/2,
    y: boardheight-playerheight-5,
    width: playerwidth,
    height: playerheight,
    velx: playervelocity,
  }
  ball = {
    x: boardwidth/2,
    y: boardheight/2,
    width: ballwidth,
    height: ballheight,
    velx: ballvelicityX,
    vely: ballvelicityY,
  }
  blocks = []
  blockrows = 3
  score = 0
  createblocks()
  console.log("reset game done")
}

function getRandomInt(min, max) {
    min = Math.ceil(min);
    max = Math.floor(max);
    return Math.floor(Math.random() * (max - min + 1)) + min;
}
