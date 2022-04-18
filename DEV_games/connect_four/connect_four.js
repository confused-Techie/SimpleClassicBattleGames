// set the games settings

var rows = 6;
var cols = 7;
var squareSize = 50;
var player = 'player1';

// get the container element

var gameBoardContainer = document.getElementById("gameboard");

// make the grid columns and rows
for (let i = 0; i < cols; i++) {
  for (let j = 0; j < rows; j++) {

    // create the new div HTML element for each grid square and make it the right size
    var square = document.createElement("div");
    gameBoardContainer.appendChild(square);

    // give each element a unique id based on its row & column s${ROW}${COL}
    square.id = 's' + j + i;

    square.classList.add("playableSquare");

    // set value of custom data attribute
    square.dataset.occupant = "blank";

    // set the square coordinates
    var topPosition = j * squareSize;
    var leftPosition = i * squareSize;

    // use CSS absolute positioning to place each grid square on the page
    square.style.top = topPosition + 'px';
    square.style.left = leftPosition + 'px';
  }
}


// set event listener for all elements in gameboard,

gameBoardContainer.addEventListener("click", gameAction, false);


class ConnectFour {
  constructor() {
    // The board contains the game board as an array.
    // with the first dimension containing each column, and the second dimension being the row of the column
    // the rows are top to bottom, to make translation to the game board easier
    // 0 = empty, 1 = player1, 2 = player2
    this.board = [
      [0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0]
    ];
  }

  playSquare(row, col, player) {
    // first check if we can play on this column.
    // -1 is used as the check to ensure we can account for the zero space of the board.
    for (var i = this.board[col].length-1; i > -1; i--) {
      if (this.board[col][i] === 0) {
        return { col: col, row: i, playable: true };
        break;
      }
    }
    return { playable: false };
  }

  occupySquare(row, col, player) {
    // this will be used to modify the gamemasters internal board.
    if (player == 'player1') {
      this.board[col][row] = 1;
    } else if (player == 'player2') {
      this.board[col][row] = 2;
    } else {
      console.log(`Unrecongized player attempting to occupy the board: ${player}`);
    }
  }

  checkWinner() {
    // this will need to check for any horizontal, vertical, or diagnol fours in a row.

    // then for tie we will count all the empty spots, and if none, declare a tie.
    var emptySpaces = 0;

    // first we can check all the rows.
    for (var i = 0; i < this.board.length; i++) {
      for (var y = this.board[i].length-1; y > -1; y--) {
        // the best implementation I can currently think of is a fan out approch.
        // starting here at the bottom of each row, and fanning out in all upper possible directions to check for a winning position.

        // but we also don't want to waste any time on an empty starting block.
        if (this.board[i][y] == 0) {
          emptySpaces++;
          continue;
        }

        // so we will check in a vertical up position for a win.
        var curPlayer = this.board[i][y];
        // we only have to check up to three vertical spaces, since we start on the first space and know the value of that one already.

        // while originally this was a collection of if...if else statements, after adding proper error checking to not fail on an undefined spot check,
        // the nested if else, then if would cause the parent if else to not be run, causing this slightly messy change, to a validity win check, switch statement

        var validWinChecks = function(brd) {
          // so for some reason validWinChecks doesn't retain the context of this, so I have to import the board, or context here.
          var tmpObj = {};
          if (brd[i] && brd[i][y-3]) {
            tmpObj.vertical = true;
          } else {
            tmpObj.vertical = false;
          }
          if (brd[i+3] && brd[i+3][y]) {
            tmpObj.horizontal = true;
          } else {
            tmpObj.horizontal = false;
          }
          if (brd[i+3] && brd[i+3][y+3]) {
            tmpObj.diagleft = true;
          } else {
            tmpObj.diagleft = false;
          }
          if (brd[i+3] && brd[i+3][y-3]) {
            tmpObj.diagright = true;
          } else {
            tmpObj.diagright = false;
          }
          return tmpObj;
        };

        const findWinner = function(i, y) {
          if (curPlayer == 1) {
            return 'player1';
          } else if (curPlayer == 2) {
            return 'player2';
          }
        };

        var winCheckRes = validWinChecks(this.board);

        if (winCheckRes.vertical) {
          if (this.board[i][y-1] == curPlayer && this.board[i][y-2] == curPlayer && this.board[i][y-3] == curPlayer) {
            return { win: true, player: findWinner() };
          }
        }
        if (winCheckRes.horizontal) {
          if (this.board[i+1][y] == curPlayer && this.board[i+2][y] == curPlayer && this.board[i+3][y] == curPlayer) {
            return { win: true, player: findWinner() };
          }
        }
        if (winCheckRes.diagleft) {
          if (this.board[i+1][y+1] == curPlayer && this.board[i+2][y+2] == curPlayer && this.board[i+3][y+3] == curPlayer) {
            return { win: true, player: findWinner() };
          }
        }
        if (winCheckRes.diagright) {
          console.log('diagright running');
          if (this.board[i+1][y-1] == curPlayer && this.board[i+2][y-2] == curPlayer && this.board[i+3][y-3] == curPlayer) {
            return { win: true, player: findWinner() };
          }
        }
      }
    }

    if (emptySpaces == 0) {
      return { win: false, tie: true };
    } else {
      console.log(`emptySpaces: ${emptySpaces}`);
      return { win: false };
    }
  }

}

var gameMaster = new ConnectFour();

function gameAction(e) {
  // if item clicked (e.target) is not the parent element on which the event listener was set
  if (e.target !== e.currentTarget) {
    // extract row and column # from the HTML element's id
    var row = e.target.id.substring(1, 2);
    var col = e.target.id.substring(2, 3);

    //alert(`Clicked row: ${row}; col: ${col}`);

    // now we can use the class to check if we can play here.
    var turn = gameMaster.playSquare(row, col, player);
    console.log(`Input: Col: ${col}, Row: ${row}; Output: `);
    console.log(turn);
    if (turn.playable) {
      gameMaster.occupySquare(turn.row, turn.col, player);

      var square = document.getElementById(`s${turn.row}${turn.col}`);

      var colorToSet;
      if (player == 'player1') { colorToSet = "red"; }
      else if (player == 'player2') { colorToSet = "blue"; }

      square.style.backgroundColor = colorToSet;
      console.log('color chhanged');
      // then to check for any winners

      var result = gameMaster.checkWinner();
      if (result.win) {
        console.log('alerting win');
        document.getElementById('dynamic_text').innerText = `${result.player} WINS!`;
        //alert(`${result.player} WINS!`);
      } else if (result.tie) {
        console.log('alerting tie');
        document.getElementById('dynamic_text').innerText = `Its a Tie!`;
      }

      // now change players
      if (player == 'player1') { player = 'player2'; }
      else if (player == 'player2') { player = 'player1'; }

    } else {
      alert('Cant play on this square.');
    }

  }
}
