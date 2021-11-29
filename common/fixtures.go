package common

// ----- JACK -----

const Array_Main_Jack string = `class Main {
    function void main() {
        var Array a;
        var int length;
        var int i, sum;

        let length = Keyboard.readInt("HOW MANY NUMBERS? ");
        let a = Array.new(length);
        let i = 0;

        while (i < length) {
            let a[i] = Keyboard.readInt("ENTER THE NEXT NUMBER: ");
            let i = i + 1;
        }

        let i = 0;
        let sum = 0;

        while (i < length) {
        let sum = sum + a[i];
            let i = i + 1;
        }

        do Output.printString("THE AVERAGE IS: ");
        do Output.printInt(sum / length);
        do Output.println();

        return;
    }
}
`

const Average_Main_Jack string = `class Main {
    function void main() {
        var Array a;
        var int length;
        var int i, sum;

        let length = Keyboard.readInt("How many numbers? ");
        let a = Array.new(length);

        let i = 0;
        while (i < length) {
            let a[i] = Keyboard.readInt("Enter a number: ");
            let sum = sum + a[i];
            let i = i + 1;
        }

        do Output.printString("The average is ");
        do Output.printInt(sum / length);

        return;
    }
}
`

const ComplexArrays_Main_Jack string = `class Main {
    function void main() {
        var Array a, b, c;

        let a = Array.new(10);
        let b = Array.new(5);
        let c = Array.new(1);

        let a[3] = 2;
        let a[4] = 8;
        let a[5] = 4;
        let b[a[3]] = a[3] + 3;
        let a[b[a[3]]] = a[a[5]] * b[((7 - a[3]) - Main.double(2)) + 1];
        let c[0] = null;
        let c = c[0];

        do Output.printString("Test 1: expected result: 5; actual result: ");
        do Output.printInt(b[2]);
        do Output.println();
        do Output.printString("Test 2: expected result: 40; actual result: ");
        do Output.printInt(a[5]);
        do Output.println();
        do Output.printString("Test 3: expected result: 0; actual result: ");
        do Output.printInt(c);
        do Output.println();

        let c = null;

        if (c = null) {
            do Main.fill(a, 10);
            let c = a[3];
            let c[1] = 33;
            let c = a[7];
            let c[1] = 77;
            let b = a[3];
            let b[1] = b[1] + c[1];
        }

        do Output.printString("Test 4: expected result: 77; actual result: ");
        do Output.printInt(c[1]);
        do Output.println();
        do Output.printString("Test 5: expected result: 110; actual result: ");
        do Output.printInt(b[1]);
        do Output.println();
        return;
    }

    function int double(int a) {
        return a * 2;
    }

    function void fill(Array a, int size) {
        while (size > 0) {
            let size = size - 1;
            let a[size] = Array.new(3);
        }
        return;
    }
}
`

const ConvertToBin_Main_Jack string = `/**
* 1) Load the program into the supplied VM emulator
* 2) Put some value in RAM[8000]
* 3) Switch to "no animation"
* 4) Run the program (give it enough time to run)
* 5) Stop the program
* 6) Check that RAM[8001]..RAM[8016] contains the correct binary result, and that none of these memory locations contains -1
*/
class Main {
   function void main() {
       var int value;
       do Main.fillMemory(8001, 16, -1);
       let value = Memory.peek(8000);
       do Main.convert(value);
       return;
   }

   function void convert(int value) {
       var int mask, position;
       var boolean loop;

       let loop = true;
       while (loop) {
           let position = position + 1;
           let mask = Main.nextMask(mask);

           if (~(position > 16)) {
               if (~((value & mask) = 0)) {
                   do Memory.poke(8000 + position, 1);
                  }
               else {
                   do Memory.poke(8000 + position, 0);
               }
           }
           else {
               let loop = false;
           }
       }
       return;
   }

   function int nextMask(int mask) {
       if (mask = 0) {
           return 1;
       }
       else {
           return mask * 2;
       }
   }

   function void fillMemory(int startAddress, int length, int value) {
       while (length > 0) {
           do Memory.poke(startAddress, value);
           let length = length - 1;
           let startAddress = startAddress + 1;
       }
       return;
   }
}
`

const HelloWorld_Main_Jack string = `class Main {
    function void main () {
        // Hello World...
        do Output.printString("Hello World");

        /*
          Hello
          World
        */
        
        return;
    }
}
`

const MultiConditionIf_Main_Jack string = `class Main {
    function void main () {
        var char key;
        let key = Keyboard.keyPressed();
        if ((key = Direction.U()) | (key = Direction.D()) | (key = Direction.L()) | (key = Direction.R())) {
        }
        return;
    }
}
`

const Pong_Ball_Jack string = `class Ball {
    field int x, y;
    field int lengthx, lengthy;
    field int d, straightD, diagonalD;
    field boolean invert, positivex, positivey;
    field int leftWall, rightWall, topWall, bottomWall;
    field int wall;

    constructor Ball new(int Ax, int Ay, int AleftWall, int ArightWall, int AtopWall, int AbottomWall) {
        let x = Ax;
        let y = Ay;
        let leftWall = AleftWall;
        let rightWall = ArightWall - 6;
        let topWall = AtopWall;
        let bottomWall = AbottomWall - 6;
        let wall = 0;
        do show();
        return this;
    }

    method void dispose() {
        do Memory.deAlloc(this);
        return;
    }

    method void show() {
        do Screen.setColor(true);
        do draw();
        return;
    }

    method void hide() {
        do Screen.setColor(false);
        do draw();
        return;
    }

    method void draw() {
        do Screen.drawRectangle(x, y, x + 5, y + 5);
        return;
    }

    method int getLeft() {
        return x;
    }

    method int getRight() {
        return x + 5;
    }

    method void setDestination(int destx, int desty) {
        var int dx, dy, temp;
        let lengthx = destx - x;
        let lengthy = desty - y;
        let dx = Math.abs(lengthx);
        let dy = Math.abs(lengthy);
        let invert = (dx < dy);

        if (invert) {
            let temp = dx;
            let dx = dy;
            let dy = temp;
            let positivex = (y < desty);
            let positivey = (x < destx);
        }
        else {
            let positivex = (x < destx);
            let positivey = (y < desty);
        }

        let d = (2 * dy) - dx;
        let straightD = 2 * dy;
        let diagonalD = 2 * (dy - dx);

        return;
    }

    method int move() {
        do hide();

        if (d < 0) { let d = d + straightD; }
        else {
            let d = d + diagonalD;

            if (positivey) {
                if (invert) { let x = x + 4; }
                else { let y = y + 4; }
            }
            else {
                if (invert) { let x = x - 4; }
                else { let y = y - 4; }
            }
        }

        if (positivex) {
            if (invert) { let y = y + 4; }
            else { let x = x + 4; }
        }
        else {
            if (invert) { let y = y - 4; }
            else { let x = x - 4; }
        }

        if (~(x > leftWall)) {
            let wall = 1;
            let x = leftWall;
        }
        if (~(x < rightWall)) {
            let wall = 2;
            let x = rightWall;
        }
        if (~(y > topWall)) {
            let wall = 3;
            let y = topWall;
        }
        if (~(y < bottomWall)) {
            let wall = 4;
            let y = bottomWall;
        }

        do show();

        return wall;
    }

    method void bounce(int bouncingDirection) {
        var int newx, newy, divLengthx, divLengthy, factor;
        let divLengthx = lengthx / 10;
        let divLengthy = lengthy / 10;
        if (bouncingDirection = 0) { let factor = 10; }
        else {
            if (((~(lengthx < 0)) & (bouncingDirection = 1)) | ((lengthx < 0) & (bouncingDirection = (-1)))) {
                let factor = 20;
            }
            else { let factor = 5; }
        }

        if (wall = 1) {
            let newx = 506;
            let newy = (divLengthy * (-50)) / divLengthx;
            let newy = y + (newy * factor);
        }
        else {
            if (wall = 2) {
                let newx = 0;
                let newy = (divLengthy * 50) / divLengthx;
                let newy = y + (newy * factor);
            }
            else {
                if (wall = 3) {
                    let newy = 250;
                    let newx = (divLengthx * (-25)) / divLengthy;
                    let newx = x + (newx * factor);
                }
                else {
                    let newy = 0;
                    let newx = (divLengthx * 25) / divLengthy;
                    let newx = x + (newx * factor);
                }
            }
        }

        do setDestination(newx, newy);
        return;
    }
}
`

const Pong_Bat_Jack string = `class Bat {
    field int x, y;
    field int width, height;
    field int direction;

    constructor Bat new(int Ax, int Ay, int Awidth, int Aheight) {
        let x = Ax;
        let y = Ay;
        let width = Awidth;
        let height = Aheight;
        let direction = 2;
        do show();
        return this;
    }

    method void dispose() {
        do Memory.deAlloc(this);
        return;
    }

    method void show() {
        do Screen.setColor(true);
        do draw();
        return;
    }

    method void hide() {
        do Screen.setColor(false);
        do draw();
        return;
    }

    method void draw() {
        do Screen.drawRectangle(x, y, x + width, y + height);
        return;
    }

    method void setDirection(int Adirection) {
        let direction = Adirection;
        return;
    }

    method int getLeft() {
        return x;
    }

    method int getRight() {
        return x + width;
    }

    method void setWidth(int Awidth) {
        do hide();
        let width = Awidth;
        do show();
        return;
    }

    method void move() {
        if (direction = 1) {
            let x = x - 4;
            if (x < 0) { let x = 0; }
            do Screen.setColor(false);
            do Screen.drawRectangle((x + width) + 1, y, (x + width) + 4, y + height);
            do Screen.setColor(true);
            do Screen.drawRectangle(x, y, x + 3, y + height);
        }
        else {
            let x = x + 4;
            if ((x + width) > 511) { let x = 511 - width; }
            do Screen.setColor(false);
            do Screen.drawRectangle(x - 4, y, x - 1, y + height);
            do Screen.setColor(true);
            do Screen.drawRectangle((x + width) - 3, y, x + width, y + height);
        }
        return;
    }
}
`

const Pong_Main_Jack string = `class Main {
    function void main() {
        var PongGame game;
        do PongGame.newInstance();
        let game = PongGame.getInstance();
        do game.run();
        do game.dispose();
        return;
    }
}
`

const Pong_PongGame_Jack string = `class PongGame {
    static PongGame instance;
    field Bat bat;
    field Ball ball;
    field int wall;
    field boolean exit;
    field int score;
    field int lastWall;
    field int batWidth;

    constructor PongGame new() {
        do Screen.clearScreen();
        let batWidth = 50;
        let bat = Bat.new(230, 229, batWidth, 7);
        let ball = Ball.new(253, 222, 0, 511, 0, 229);
        do ball.setDestination(400,0);
        do Screen.drawRectangle(0, 238, 511, 240);
        do Output.moveCursor(22,0);
        do Output.printString("Score: 0");

        let exit = false;
        let score = 0;
        let wall = 0;
        let lastWall = 0;

        return this;
    }

    method void dispose() {
        do bat.dispose();
        do ball.dispose();
        do Memory.deAlloc(this);
        return;
    }

    function void newInstance() {
        let instance = PongGame.new();
        return;
    }

    function PongGame getInstance() {
        return instance;
    }

    method void run() {
        var char key;
        while (~exit) {
            while ((key = 0) & (~exit)) {
                let key = Keyboard.keyPressed();
                do bat.move();
                do moveBall();
                do Sys.wait(50);
            }

            if (key = 130) { do bat.setDirection(1); }
            else {
                if (key = 132) { do bat.setDirection(2); }
                else {
                    if (key = 140) { let exit = true; }
                }
            }

            while ((~(key = 0)) & (~exit)) {
                let key = Keyboard.keyPressed();
                do bat.move();
                do moveBall();
                do Sys.wait(50);
            }
        }

        if (exit) {
            do Output.moveCursor(10,27);
            do Output.printString("Game Over");
        }

        return;
    }

    method void moveBall() {
        var int bouncingDirection, batLeft, batRight, ballLeft, ballRight;

        let wall = ball.move();

        if ((wall > 0) & (~(wall = lastWall))) {
            let lastWall = wall;
            let bouncingDirection = 0;
            let batLeft = bat.getLeft();
            let batRight = bat.getRight();
            let ballLeft = ball.getLeft();
            let ballRight = ball.getRight();

            if (wall = 4) {
                let exit = (batLeft > ballRight) | (batRight < ballLeft);
                if (~exit) {
                    if (ballRight < (batLeft + 10)) { let bouncingDirection = -1; }
                    else {
                        if (ballLeft > (batRight - 10)) { let bouncingDirection = 1; }
                    }
                    let batWidth = batWidth - 2;
                    do bat.setWidth(batWidth);
                    let score = score + 1;
                    do Output.moveCursor(22,7);
                    do Output.printInt(score);
                }
            }
            do ball.bounce(bouncingDirection);
        }
        return;
    }
}
`

const Seven_Main_Jack string = `class Main {
    function void main() {
        do Output.printInt(1 + (2 * 3));
        return;
    }
}
`

const SimpleIf_Main_Jack string = `class Main {
    function void main () {
        var int x;
        var int y;
        var int z;
        let x = 0;
        let y = 0;
        let z = 0;

        if (x = 0) {
            if (y = 0) {
                let z = 0;
            }
            else {
                let z = 1;
            }
        }
        else {
            let z = 2;
        }

        if (y = 0) {
            let z = 3;
        }
        else {
            let z = 4;
        }

        if (z = 0) {
            let z = 5;
        }

        return;
    }
}
`

const SimpleWhile_Main_Jack string = `class Main {
    function void main () {
        var int x;
        var int y;
        var int z;
        let x = 0;
        let y = 0;
        let z = 0;

        while (x < 3) {
            while (y < 3) {
                let y = y + 1;
            }
            let x = x + 1;
            let y = 0;
        }

        while (z < 3) {
            let z = z + 1;
        }

        return;
    }
}
`

const Sokoban_Board_Jack string = `class Board {
    field int width, height;
    field Array cells;
    field Player player;

    // creates a board (map data encoding: Player "@", Box "$", Goal ".", Wall "#", Goal+Player "+", Goal+Box "*")
    constructor Board new(int aWidth, int aHeight, String aMapData) {
        var int x, y;
        var char cellChar;

        let width = aWidth;
        let height = aHeight;
        let cells = Array.new(width * height);
        let x = 0;
        let y = 0;

        while (y < height) {
            while (x < width) {
                let cellChar = aMapData.charAt((y * width) + x);
                if (cellChar = 64) { // @ = 64
                    let cells[(y * width) + x] = Cell.new(CellType.CellTypeNone(), false);
                    let player = Player.new(x, y);
                }
                if (cellChar = 36) { // $ = 36
                    let cells[(y * width) + x] = Cell.new(CellType.CellTypeNone(), true);
                }
                if (cellChar = 46) { // . = 46
                    let cells[(y * width) + x] = Cell.new(CellType.CellTypeGoal(), false);
                }
                if (cellChar = 35) { // # = 35
                    let cells[(y * width) + x] = Cell.new(CellType.CellTypeWall(), false);
                }
                if (cellChar = 43) { // + = 43
                    let cells[(y * width) + x] = Cell.new(CellType.CellTypeGoal(), false);
                    let player = Player.new(x, y);
                }
                if (cellChar = 42) { // * = 42
                    let cells[(y * width) + x] = Cell.new(CellType.CellTypeGoal(), true);
                }
                if (cellChar = 32) { // (space) = 32
                    let cells[(y * width) + x] = Cell.new(CellType.CellTypeNone(), false);
                }
                let x = x + 1;
            }
            let y = y + 1;
            let x = 0;
        }

        return this;
    }

    // returns the cell at the given location
    method Cell getCell(int aX, int aY) {
        return cells[(aY * width) + aX];
    }

    // returns true if every goal cell on the board has a box
    method boolean isComplete() {
        var int i, numCells;
        var Cell cell;

        let i = 0;
        let numCells = width * height;

        while (i < numCells) {
            let cell = cells[i];
            if ((cell.getTypeOf() = CellType.CellTypeGoal()) & (~(cell.getHasBox()))) { // if cell is of type goal and cell does not have a box...
                return false;
            }
            let i = i + 1;
        }

        return true;
    }

    method int getWidth() {
        return width;
    }

    method Player getHeight() {
        return height;
    }

    method Player getPlayer() {
        return player;
    }

    method void dispose() {
        var int i, numCells;
        var Cell cell;

        let i = 0;
        let numCells = width * height;

        while (i < numCells) {
            let cell = cells[i];
            do cell.dispose();
            let i = i + 1;
        }
        do cells.dispose();
        do player.dispose();
        do Memory.deAlloc(this);

        return;
    }
}
`

const Sokoban_Cell_Jack string = `class Cell {
    field CellType typeOf;
    field boolean hasBox;

    constructor Cell new(CellType aTypeOf, boolean aHasBox) {
        let typeOf = aTypeOf;
        let hasBox = aHasBox;

        return this;
    }

    method CellType getTypeOf() {
        return typeOf;
    }

    method boolean getHasBox() {
        return hasBox;
    }

    method void setHasBox(boolean aHasBox) {
        let hasBox = aHasBox;

        return;
    }

    method void dispose() {
        do Memory.deAlloc(this);

        return;
    }
}
`

const Sokoban_CellType_Jack string = `class CellType {
    static int CellTypeNone, CellTypeGoal, CellTypeWall;

    function void init() {
        let CellTypeNone = 0;
        let CellTypeGoal = 1;
        let CellTypeWall = 2;

        return;
    }

    function int CellTypeNone() {
        return CellTypeNone;
    }

    function int CellTypeGoal() {
        return CellTypeGoal;
    }

    function int CellTypeWall() {
        return CellTypeWall;
    }
}
`

const Sokoban_Controller_Jack string = `class Controller {
    field Model model;
    field boolean running;
    field char lastTickKey;

    // creates a controller
    constructor Controller new(Model aModel) {
        let model = aModel;
        let running = true;
        let lastTickKey = 0;

        return this;
    }

    // starts a new game at level 1
    method void startNewGame() {
        var LevelManager levelManager;

        let levelManager = model.getLevelManager();

        do levelManager.reset();
        do tryStartNextLevel();

        return;
    }

    // handles user input as appropriate (game state dependent behaviour)
    method void handleInput() {
        var char key;

        let key = Keyboard.keyPressed();

        if (~(key = 0)) {
            if (~((key = Direction.U()) | (key = Direction.D()) | (key = Direction.L()) | (key = Direction.R()) | (key = 32) | (key = 82) | (key = 83) | (key = 140))) {
                let lastTickKey = 0;
                return;
            }
            if (key = 140) { // Escape = 140
                let running = false;
            }
            if (~(lastTickKey = key)) {
                if (model.getState() = State.StateGameComplete()) {
                    if (key = 32) { // (space) = 32
                        do startNewGame();
                    }
                }
                if (model.getState() = State.StateLevelComplete()) {
                    if (key = 32) { // (space) = 32
                        do tryStartNextLevel();
                    }
                }
                if (model.getState() = State.StatePlaying()) {
                    if ((key = Direction.U()) | (key = Direction.D()) | (key = Direction.L()) | (key = Direction.R())) {
                        do tryMovePlayer(key);
                    }
                    if (key = 82) { // R = 82
                        do restartLevel();
                    }
                    if (key = 83) { // S = 83
                        // do tryStartNextLevel(); // uncomment this line to enable cheating :)
                    }
                }
            }
            let lastTickKey = key;
        }
        else {
            let lastTickKey = 0;
        }

        return;
    }

    // move player (and an adjacent box where appropriate) in the specified direction if possible & check for board completion (and handle appropriately) if a box is moved
    method void tryMovePlayer(Direction aDirection) {
        var Board board;
        var Player player;
        var int currentX, currentY, targetX, targetY, nextX, nextY;
        var Cell targetCell;
        var Cell nextCell;

        let board = model.getBoard();
        let player = board.getPlayer();
        let currentX = player.getX();
        let currentY = player.getY();
        let targetX = currentX;
        let targetY = currentY;
        let nextX = targetX;
        let nextY = targetY;

        if (aDirection = Direction.U()) {
            let targetY = targetY - 1;
            let nextY = nextY - 2;
        }
        if (aDirection = Direction.D()) {
            let targetY = targetY + 1;
            let nextY = nextY + 2;
        }
        if (aDirection = Direction.L()) {
            let targetX = targetX - 1;
            let nextX = nextX - 2;
        }
        if (aDirection = Direction.R()) {
            let targetX = targetX + 1;
            let nextX = nextX + 2;
        }

        let targetCell = board.getCell(targetX, targetY);

        if (~(targetCell.getTypeOf() = CellType.CellTypeWall())) {
            if (targetCell.getHasBox()) {
                let nextCell = board.getCell(nextX, nextY);
                if ((~(nextCell.getTypeOf() = CellType.CellTypeWall())) & (~(nextCell.getHasBox()))) { // if next cell is not of type wall and next cell does not have a box...
                    do targetCell.setHasBox(false);
                    do nextCell.setHasBox(true);
                    do player.moveTo(targetX, targetY);
                    if (board.isComplete()) {
                        do model.setState(State.StateLevelComplete());
                    }
                }
            }
            else {
                do player.moveTo(targetX, targetY);
            }
        }

        return;
    }

    // starts the next level if the current one isn't the last, else sets game state to game complete
    method void tryStartNextLevel() {
        var LevelManager levelManager;
        var Level level;
        var Board board;

        let levelManager = model.getLevelManager();

        if (levelManager.hasNextLevel()) {
            do levelManager.progressToNextLevel();

            let level = levelManager.getCurrentLevel();
            let board = Board.new(level.getWidth(), level.getHeight(), level.getMapData());

            do model.setBoard(board);
            do model.setState(State.StatePlaying());
        }
        else {
            do model.setState(State.StateGameComplete());
        }

        return;
    }

    // resets the game board to the current level's starting state
    method void restartLevel() {
        var LevelManager levelManager;
        var Level level;
        var Board board;

        let levelManager = model.getLevelManager();
        let level = levelManager.getCurrentLevel();
        let board = Board.new(level.getWidth(), level.getHeight(), level.getMapData());

        do model.setBoard(board);
        do model.setState(State.StatePlaying());

        return;
    }

    method boolean isRunning() {
        return running;
    }

    method void dispose() {
        do Memory.deAlloc(this);

        return;
    }
}
`

const Sokoban_Direction_Jack string = `class Direction {
    static int U, D, L, R;

    function void init() {
        let U = 131;
        let D = 133;
        let L = 130;
        let R = 132;

        return;
    }

    function int U() {
        return U;
    }

    function int D() {
        return D;
    }

    function int L() {
        return L;
    }

    function int R() {
        return R;
    }
}
`

const Sokoban_Level_Jack string = `class Level {
    field int width, height;
    field string mapData; // Player "@", Box "$", Goal ".", Wall "#", Goal+Player "+", Goal+Box "*", None " "

    constructor Level new(int aWidth, int aHeight, string aMapData) {
        let width = aWidth;
        let height = aHeight;
        let mapData = aMapData;

        return this;
    }

    method int getWidth() {
        return width;
    }

    method int getHeight() {
        return height;
    }

    method string getMapData() {
        return mapData;
    }

    method void dispose() {
        do Memory.deAlloc(this);

        return;
    }
}
`

const Sokoban_LevelManager_Jack string = `class LevelManager {
    field int currentLevelNumber, finalLevelNumber;
    field Array levels;

    // creates a level manager
    constructor LevelManager new() {
        let currentLevelNumber = 0;
        let finalLevelNumber = 10;
        let levels = Array.new(finalLevelNumber+1);
        let levels[0] = Level.new(0, 0, "");
        let levels[1] = Level.new(8, 8, "  ###     #.#     # #######$ $.##. $@#######$#     #.#     ###  ");
        let levels[2] = Level.new(9, 9, "#####    #   #    # $@# #### $$# #.#### ###.# ##    .# #   #  # #   #### #####   ");
        let levels[3] = Level.new(6, 8, " #### ##  # # @$# ##$ #### $ ##.$  ##..*.#######");
        let levels[4] = Level.new(8, 8, " ####    #@ ###  # $  # ### # ###.# #  ##.$  # ##.   $ #########");
        let levels[5] = Level.new(8, 7, "  ######  #    ####$$$ ##@ $.. ## $...######  #    #### ");
        let levels[6] = Level.new(8, 7, "  ##### ###  @# #  $. ###  .$. #### *$ #  #   ##  ##### ");
        let levels[7] = Level.new(8, 8, "  ####    #..#   ## .##  #  $.# ## $  ###  #$$ ##  @   #########");
        let levels[8] = Level.new(8, 7, "#########  #   ##@$..$ ## $.* ### $..$ ##  #   #########");
        let levels[9] = Level.new(9, 7, "######   #    #   # $$$##  #  #..#####  ..$ # # @    # ########");
        let levels[10] = Level.new(7, 8, "########..$..##..#..## $$$ ##  $  ## $$$ ##  #@ ########");

        return this;
    }

    // returns the current level number
    method int getCurrentLevelNumber() {
        return currentLevelNumber;
    }

    // returns the final level number
    method int getFinalLevelNumber() {
        return finalLevelNumber;
    }

    // returns the current level
    method Level getCurrentLevel() {
        return levels[currentLevelNumber];
    }

    // returns true if the current level is not the last
    method boolean hasNextLevel() {
        return currentLevelNumber < finalLevelNumber;
    }

    // increments the current level
    method void progressToNextLevel() {
        let currentLevelNumber = currentLevelNumber + 1;

        return;
    }

    // resets the level manager
    method void reset() {
        let currentLevelNumber = 0;

        return;
    }

    method void dispose() {
        do levels.dispose();
        do Memory.deAlloc(this);

        return;
    }
}
`

const Sokoban_Main_Jack string = `class Main {
    function void main() {
        var Direction direction;
        var Model model;
        var View view;
        var Controller controller;

        do CellType.init();
        do Direction.init();
        do State.init();
        
        let model = Model.new();
        let view = View.new(model);
        let controller = Controller.new(model);

        do controller.startNewGame();

        while (controller.isRunning()) {
            do controller.handleInput();
            do view.draw();
            do model.update();
            do Sys.wait(10);
        }

        do model.dispose();
        do view.dispose();
        do controller.dispose();

        return;
    }
}
`

const Sokoban_Model_Jack string = `class Model {
    field LevelManager levelManager;
    field Board board;
    field State state;
    field int tickAccumulator;
    field boolean screenDirty;

    constructor Model new() {
        let levelManager = LevelManager.new();
        let state = State.StatePlaying();
        let tickAccumulator = 0;
        let screenDirty = false;
        return this;
    }

    method void update() {
        let tickAccumulator = tickAccumulator + 1;
        if (tickAccumulator > 20) {
            let tickAccumulator = 0;
        }
        if (screenDirty) {
            let screenDirty = false;
        }

        return;
    }

    method LevelManager getLevelManager() {
        return levelManager;
    }

    method Board getBoard() {
        return board;
    }

    method State getState() {
        return state;
    }

    method int getTickAccumulator() {
        return tickAccumulator;
    }

    method boolean isScreenDirty() {
        return screenDirty;
    }

    method void setBoard(Board aBoard) {
        if (~(board = null)) {
            do board.dispose();
        }
        let board = aBoard;

        return;
    }

    method void setState(State aState) {
        let state = aState;
        let screenDirty = true; // forces a full screen refresh on state change

        return;
    }

    method void dispose() {
        do levelManager.dispose();
        do board.dispose();
        do Memory.deAlloc(this);

        return;
    }
}
`

const Sokoban_Player_Jack string = `class Player {
    field int x, y;

    constructor Player new(int ax, int ay) {
        let x = ax;
        let y = ay;

        return this;
    }

    method int getX() {
        return x;
    }

    method int getY() {
        return y;
    }

    method void moveTo(int aX, int aY) {
        let x = aX;
        let y = aY;

        return;
    }

    method void dispose() {
        do Memory.deAlloc(this);

        return;
    }
}
`

const Sokoban_Sprites_Jack string = `class Sprites {

    function void drawBox(int location) {
        var int memAddress;
        let memAddress = 16384+location;
        do Memory.poke(memAddress+0, 0);
        do Memory.poke(memAddress+32, 0);
        do Memory.poke(memAddress+64, 0);
        do Memory.poke(memAddress+96, 4080);
        do Memory.poke(memAddress+128, 8184);
        do Memory.poke(memAddress+160, 8184);
        do Memory.poke(memAddress+192, 8184);
        do Memory.poke(memAddress+224, 8184);
        do Memory.poke(memAddress+256, 8184);
        do Memory.poke(memAddress+288, 8184);
        do Memory.poke(memAddress+320, 8184);
        do Memory.poke(memAddress+352, 8184);
        do Memory.poke(memAddress+384, 4080);
        do Memory.poke(memAddress+416, 0);
        do Memory.poke(memAddress+448, 0);
        do Memory.poke(memAddress+480, 0);
        return;
    }

    function void drawClear(int location) {
        var int memAddress;
        let memAddress = 16384+location;
        do Memory.poke(memAddress+0, 0);
        do Memory.poke(memAddress+32, 0);
        do Memory.poke(memAddress+64, 0);
        do Memory.poke(memAddress+96, 0);
        do Memory.poke(memAddress+128, 0);
        do Memory.poke(memAddress+160, 0);
        do Memory.poke(memAddress+192, 0);
        do Memory.poke(memAddress+224, 0);
        do Memory.poke(memAddress+256, 0);
        do Memory.poke(memAddress+288, 0);
        do Memory.poke(memAddress+320, 0);
        do Memory.poke(memAddress+352, 0);
        do Memory.poke(memAddress+384, 0);
        do Memory.poke(memAddress+416, 0);
        do Memory.poke(memAddress+448, 0);
        do Memory.poke(memAddress+480, 0);
        return;
    }

    function void drawGoal(int location) {
        var int memAddress;
        let memAddress = 16384+location;
        do Memory.poke(memAddress+0, 0);
        do Memory.poke(memAddress+32, 16380);
        do Memory.poke(memAddress+64, 24582);
        do Memory.poke(memAddress+96, 16386);
        do Memory.poke(memAddress+128, 16386);
        do Memory.poke(memAddress+160, 16386);
        do Memory.poke(memAddress+192, 16386);
        do Memory.poke(memAddress+224, 16386);
        do Memory.poke(memAddress+256, 16386);
        do Memory.poke(memAddress+288, 16386);
        do Memory.poke(memAddress+320, 16386);
        do Memory.poke(memAddress+352, 16386);
        do Memory.poke(memAddress+384, 16386);
        do Memory.poke(memAddress+416, 24582);
        do Memory.poke(memAddress+448, 16380);
        do Memory.poke(memAddress+480, 0);
        return;
    }

    function void drawGoalBox(int location) {
        var int memAddress;
        let memAddress = 16384+location;
        do Memory.poke(memAddress+0, 0);
        do Memory.poke(memAddress+32, 16380);
        do Memory.poke(memAddress+64, 24582);
        do Memory.poke(memAddress+96, 20466);
        do Memory.poke(memAddress+128, 24570);
        do Memory.poke(memAddress+160, 24570);
        do Memory.poke(memAddress+192, 24570);
        do Memory.poke(memAddress+224, 24570);
        do Memory.poke(memAddress+256, 24570);
        do Memory.poke(memAddress+288, 24570);
        do Memory.poke(memAddress+320, 24570);
        do Memory.poke(memAddress+352, 24570);
        do Memory.poke(memAddress+384, 20466);
        do Memory.poke(memAddress+416, 24582);
        do Memory.poke(memAddress+448, 16380);
        do Memory.poke(memAddress+480, 0);
        return;
    }

    function void drawGoalPlayer(int location) {
        var int memAddress;
        let memAddress = 16384+location;
        do Memory.poke(memAddress+0, 0);
        do Memory.poke(memAddress+32, 16380);
        do Memory.poke(memAddress+64, 24582);
        do Memory.poke(memAddress+96, 16386);
        do Memory.poke(memAddress+128, 18402);
        do Memory.poke(memAddress+160, 20466);
        do Memory.poke(memAddress+192, 19890);
        do Memory.poke(memAddress+224, 23994);
        do Memory.poke(memAddress+256, 24570);
        do Memory.poke(memAddress+288, 19410);
        do Memory.poke(memAddress+320, 19506);
        do Memory.poke(memAddress+352, 18402);
        do Memory.poke(memAddress+384, 16386);
        do Memory.poke(memAddress+416, 24582);
        do Memory.poke(memAddress+448, 16380);
        do Memory.poke(memAddress+480, 0);
        return;
    }

    function void drawLogoA1(int location) {
        var int memAddress;
        let memAddress = 16384+location;
        do Memory.poke(memAddress+0, 0);
        do Memory.poke(memAddress+32, 0);
        do Memory.poke(memAddress+64, 0);
        do Memory.poke(memAddress+96, 0);
        do Memory.poke(memAddress+128, 0);
        do Memory.poke(memAddress+160, 0);
        do Memory.poke(memAddress+192, 0);
        do Memory.poke(memAddress+224, 0);
        do Memory.poke(memAddress+256, 0);
        do Memory.poke(memAddress+288, -4096);
        do Memory.poke(memAddress+320, -1024);
        do Memory.poke(memAddress+352, -512);
        do Memory.poke(memAddress+384, 7936);
        do Memory.poke(memAddress+416, 16128);
        do Memory.poke(memAddress+448, -256);
        do Memory.poke(memAddress+480, -512);
        return;
    }

    function void drawLogoA2(int location) {
        var int memAddress;
        let memAddress = 16384+location;
        do Memory.poke(memAddress+0, -2048);
        do Memory.poke(memAddress+32, -16384);
        do Memory.poke(memAddress+64, -28928);
        do Memory.poke(memAddress+96, -256);
        do Memory.poke(memAddress+128, -512);
        do Memory.poke(memAddress+160, -512);
        do Memory.poke(memAddress+192, 0);
        do Memory.poke(memAddress+224, 0);
        do Memory.poke(memAddress+256, 0);
        do Memory.poke(memAddress+288, 0);
        do Memory.poke(memAddress+320, 0);
        do Memory.poke(memAddress+352, 0);
        do Memory.poke(memAddress+384, 0);
        do Memory.poke(memAddress+416, 0);
        do Memory.poke(memAddress+448, 0);
        do Memory.poke(memAddress+480, 0);
        return;
    }

    function void drawLogoA3(int location) {
        var int memAddress;
        let memAddress = 16384+location;
        do Memory.poke(memAddress+0, 0);
        do Memory.poke(memAddress+32, 0);
        do Memory.poke(memAddress+64, -1024);
        do Memory.poke(memAddress+96, 0);
        do Memory.poke(memAddress+128, 0);
        do Memory.poke(memAddress+160, 16384);
        do Memory.poke(memAddress+192, -8192);
        do Memory.poke(memAddress+224, -20480);
        do Memory.poke(memAddress+256, -2048);
        do Memory.poke(memAddress+288, -6144);
        do Memory.poke(memAddress+320, 10240);
        do Memory.poke(memAddress+352, 16384);
        do Memory.poke(memAddress+384, 0);
        do Memory.poke(memAddress+416, 0);
        do Memory.poke(memAddress+448, -1024);
        do Memory.poke(memAddress+480, 0);
        return;
    }

    function void drawLogoB1(int location) {
        var int memAddress;
        let memAddress = 16384+location;
        do Memory.poke(memAddress+0, 0);
        do Memory.poke(memAddress+32, 0);
        do Memory.poke(memAddress+64, 0);
        do Memory.poke(memAddress+96, 0);
        do Memory.poke(memAddress+128, 0);
        do Memory.poke(memAddress+160, 0);
        do Memory.poke(memAddress+192, 0);
        do Memory.poke(memAddress+224, 0);
        do Memory.poke(memAddress+256, 0);
        do Memory.poke(memAddress+288, 15879);
        do Memory.poke(memAddress+320, -121);
        do Memory.poke(memAddress+352, -49);
        do Memory.poke(memAddress+384, -2097);
        do Memory.poke(memAddress+416, -7200);
        do Memory.poke(memAddress+448, -7199);
        do Memory.poke(memAddress+480, -7193);
        return;
    }

    function void drawLogoB2(int location) {
        var int memAddress;
        let memAddress = 16384+location;
        do Memory.poke(memAddress+0, -7185);
        do Memory.poke(memAddress+32, -7185);
        do Memory.poke(memAddress+64, -2097);
        do Memory.poke(memAddress+96, -57);
        do Memory.poke(memAddress+128, -125);
        do Memory.poke(memAddress+160, 15872);
        do Memory.poke(memAddress+192, 0);
        do Memory.poke(memAddress+224, 0);
        do Memory.poke(memAddress+256, 0);
        do Memory.poke(memAddress+288, 0);
        do Memory.poke(memAddress+320, 0);
        do Memory.poke(memAddress+352, 0);
        do Memory.poke(memAddress+384, 0);
        do Memory.poke(memAddress+416, 0);
        do Memory.poke(memAddress+448, 0);
        do Memory.poke(memAddress+480, 0);
        return;
    }

    function void drawLogoB3(int location) {
        var int memAddress;
        let memAddress = 16384+location;
        do Memory.poke(memAddress+0, 0);
        do Memory.poke(memAddress+32, 0);
        do Memory.poke(memAddress+64, -1);
        do Memory.poke(memAddress+96, 0);
        do Memory.poke(memAddress+128, 24320);
        do Memory.poke(memAddress+160, 17412);
        do Memory.poke(memAddress+192, 17423);
        do Memory.poke(memAddress+224, 17435);
        do Memory.poke(memAddress+256, -15297);
        do Memory.poke(memAddress+288, 17455);
        do Memory.poke(memAddress+320, 17448);
        do Memory.poke(memAddress+352, 17412);
        do Memory.poke(memAddress+384, 17408);
        do Memory.poke(memAddress+416, 0);
        do Memory.poke(memAddress+448, -1);
        do Memory.poke(memAddress+480, 0);
        return;
    }

    function void drawLogoC1(int location) {
        var int memAddress;
        let memAddress = 16384+location;
        do Memory.poke(memAddress+0, 0);
        do Memory.poke(memAddress+32, 0);
        do Memory.poke(memAddress+64, 0);
        do Memory.poke(memAddress+96, 0);
        do Memory.poke(memAddress+128, 0);
        do Memory.poke(memAddress+160, 0);
        do Memory.poke(memAddress+192, 0);
        do Memory.poke(memAddress+224, 0);
        do Memory.poke(memAddress+256, 0);
        do Memory.poke(memAddress+288, 4344);
        do Memory.poke(memAddress+320, 14584);
        do Memory.poke(memAddress+352, -775);
        do Memory.poke(memAddress+384, 32505);
        do Memory.poke(memAddress+416, 16371);
        do Memory.poke(memAddress+448, 8179);
        do Memory.poke(memAddress+480, 4083);
        return;
    }

    function void drawLogoC2(int location) {
        var int memAddress;
        let memAddress = 16384+location;
        do Memory.poke(memAddress+0, 8179);
        do Memory.poke(memAddress+32, 16371);
        do Memory.poke(memAddress+64, 32505);
        do Memory.poke(memAddress+96, -775);
        do Memory.poke(memAddress+128, 30968);
        do Memory.poke(memAddress+160, 6392);
        do Memory.poke(memAddress+192, 0);
        do Memory.poke(memAddress+224, 0);
        do Memory.poke(memAddress+256, 0);
        do Memory.poke(memAddress+288, 0);
        do Memory.poke(memAddress+320, 0);
        do Memory.poke(memAddress+352, 0);
        do Memory.poke(memAddress+384, 0);
        do Memory.poke(memAddress+416, 0);
        do Memory.poke(memAddress+448, 0);
        do Memory.poke(memAddress+480, 0);
        return;
    }

    function void drawLogoC3(int location) {
        var int memAddress;
        let memAddress = 16384+location;
        do Memory.poke(memAddress+0, 0);
        do Memory.poke(memAddress+32, 0);
        do Memory.poke(memAddress+64, -1);
        do Memory.poke(memAddress+96, 0);
        do Memory.poke(memAddress+128, 32244);
        do Memory.poke(memAddress+160, 4116);
        do Memory.poke(memAddress+192, 4116);
        do Memory.poke(memAddress+224, 4116);
        do Memory.poke(memAddress+256, 4215);
        do Memory.poke(memAddress+288, 4116);
        do Memory.poke(memAddress+320, 4116);
        do Memory.poke(memAddress+352, 4116);
        do Memory.poke(memAddress+384, 32244);
        do Memory.poke(memAddress+416, 0);
        do Memory.poke(memAddress+448, -1);
        do Memory.poke(memAddress+480, 0);
        return;
    }

    function void drawLogoD1(int location) {
        var int memAddress;
        let memAddress = 16384+location;
        do Memory.poke(memAddress+0, 0);
        do Memory.poke(memAddress+32, 0);
        do Memory.poke(memAddress+64, 0);
        do Memory.poke(memAddress+96, 0);
        do Memory.poke(memAddress+128, 0);
        do Memory.poke(memAddress+160, 0);
        do Memory.poke(memAddress+192, 0);
        do Memory.poke(memAddress+224, 0);
        do Memory.poke(memAddress+256, 0);
        do Memory.poke(memAddress+288, -31776);
        do Memory.poke(memAddress+320, -28680);
        do Memory.poke(memAddress+352, -24580);
        do Memory.poke(memAddress+384, -24708);
        do Memory.poke(memAddress+416, 15934);
        do Memory.poke(memAddress+448, 15934);
        do Memory.poke(memAddress+480, 15934);
        return;
    }

    function void drawLogoD2(int location) {
        var int memAddress;
        let memAddress = 16384+location;
        do Memory.poke(memAddress+0, 15934);
        do Memory.poke(memAddress+32, 15934);
        do Memory.poke(memAddress+64, -24708);
        do Memory.poke(memAddress+96, -24580);
        do Memory.poke(memAddress+128, -28680);
        do Memory.poke(memAddress+160, -31776);
        do Memory.poke(memAddress+192, 0);
        do Memory.poke(memAddress+224, 0);
        do Memory.poke(memAddress+256, 8);
        do Memory.poke(memAddress+288, 8);
        do Memory.poke(memAddress+320, 2360);
        do Memory.poke(memAddress+352, 2376);
        do Memory.poke(memAddress+384, 2376);
        do Memory.poke(memAddress+416, 3640);
        do Memory.poke(memAddress+448, 2048);
        do Memory.poke(memAddress+480, 1792);
        return;
    }

    function void drawLogoD3(int location) {
        var int memAddress;
        let memAddress = 16384+location;
        do Memory.poke(memAddress+0, 0);
        do Memory.poke(memAddress+32, 0);
        do Memory.poke(memAddress+64, -1);
        do Memory.poke(memAddress+96, 0);
        do Memory.poke(memAddress+128, -7087);
        do Memory.poke(memAddress+160, 5203);
        do Memory.poke(memAddress+192, 5205);
        do Memory.poke(memAddress+224, 5209);
        do Memory.poke(memAddress+256, -2991);
        do Memory.poke(memAddress+288, 5201);
        do Memory.poke(memAddress+320, 5201);
        do Memory.poke(memAddress+352, 4753);
        do Memory.poke(memAddress+384, 4369);
        do Memory.poke(memAddress+416, 0);
        do Memory.poke(memAddress+448, -1);
        do Memory.poke(memAddress+480, 0);
        return;
    }

    function void drawLogoE1(int location) {
        var int memAddress;
        let memAddress = 16384+location;
        do Memory.poke(memAddress+0, 0);
        do Memory.poke(memAddress+32, 0);
        do Memory.poke(memAddress+64, 0);
        do Memory.poke(memAddress+96, 0);
        do Memory.poke(memAddress+128, 0);
        do Memory.poke(memAddress+160, 0);
        do Memory.poke(memAddress+192, 0);
        do Memory.poke(memAddress+224, 0);
        do Memory.poke(memAddress+256, 0);
        do Memory.poke(memAddress+288, 255);
        do Memory.poke(memAddress+320, 511);
        do Memory.poke(memAddress+352, 1023);
        do Memory.poke(memAddress+384, 1999);
        do Memory.poke(memAddress+416, 1999);
        do Memory.poke(memAddress+448, 1023);
        do Memory.poke(memAddress+480, -31745);
        return;
    }

    function void drawLogoE2(int location) {
        var int memAddress;
        let memAddress = 16384+location;
        do Memory.poke(memAddress+0, -30721);
        do Memory.poke(memAddress+32, -28785);
        do Memory.poke(memAddress+64, -12337);
        do Memory.poke(memAddress+96, -12289);
        do Memory.poke(memAddress+128, -6145);
        do Memory.poke(memAddress+160, -7169);
        do Memory.poke(memAddress+192, 0);
        do Memory.poke(memAddress+224, 0);
        do Memory.poke(memAddress+256, 0);
        do Memory.poke(memAddress+288, 0);
        do Memory.poke(memAddress+320, 0);
        do Memory.poke(memAddress+352, 0);
        do Memory.poke(memAddress+384, 0);
        do Memory.poke(memAddress+416, 0);
        do Memory.poke(memAddress+448, 0);
        do Memory.poke(memAddress+480, 0);
        return;
    }

    function void drawLogoE3(int location) {
        var int memAddress;
        let memAddress = 16384+location;
        do Memory.poke(memAddress+0, 0);
        do Memory.poke(memAddress+32, 0);
        do Memory.poke(memAddress+64, -1);
        do Memory.poke(memAddress+96, 0);
        do Memory.poke(memAddress+128, -8388);
        do Memory.poke(memAddress+160, 16709);
        do Memory.poke(memAddress+192, 16709);
        do Memory.poke(memAddress+224, 16709);
        do Memory.poke(memAddress+256, -14523);
        do Memory.poke(memAddress+288, 16709);
        do Memory.poke(memAddress+320, 16709);
        do Memory.poke(memAddress+352, 16709);
        do Memory.poke(memAddress+384, 24381);
        do Memory.poke(memAddress+416, 0);
        do Memory.poke(memAddress+448, -1);
        do Memory.poke(memAddress+480, 0);
        return;
    }

    function void drawLogoF1(int location) {
        var int memAddress;
        let memAddress = 16384+location;
        do Memory.poke(memAddress+0, 0);
        do Memory.poke(memAddress+32, 0);
        do Memory.poke(memAddress+64, 0);
        do Memory.poke(memAddress+96, 0);
        do Memory.poke(memAddress+128, 0);
        do Memory.poke(memAddress+160, 0);
        do Memory.poke(memAddress+192, 0);
        do Memory.poke(memAddress+224, 0);
        do Memory.poke(memAddress+256, 0);
        do Memory.poke(memAddress+288, -1986);
        do Memory.poke(memAddress+320, -1986);
        do Memory.poke(memAddress+352, -1986);
        do Memory.poke(memAddress+384, -1921);
        do Memory.poke(memAddress+416, -3969);
        do Memory.poke(memAddress+448, -3977);
        do Memory.poke(memAddress+480, -3849);
        return;
    }

    function void drawLogoF2(int location) {
        var int memAddress;
        let memAddress = 16384+location;
        do Memory.poke(memAddress+0, -3849);
        do Memory.poke(memAddress+32, -3841);
        do Memory.poke(memAddress+64, -1537);
        do Memory.poke(memAddress+96, -1537);
        do Memory.poke(memAddress+128, -1053);
        do Memory.poke(memAddress+160, -1053);
        do Memory.poke(memAddress+192, 0);
        do Memory.poke(memAddress+224, 0);
        do Memory.poke(memAddress+256, 0);
        do Memory.poke(memAddress+288, 0);
        do Memory.poke(memAddress+320, 0);
        do Memory.poke(memAddress+352, 0);
        do Memory.poke(memAddress+384, 0);
        do Memory.poke(memAddress+416, 0);
        do Memory.poke(memAddress+448, 0);
        do Memory.poke(memAddress+480, 0);
        return;
    }

    function void drawLogoF3(int location) {
        var int memAddress;
        let memAddress = 16384+location;
        do Memory.poke(memAddress+0, 0);
        do Memory.poke(memAddress+32, 0);
        do Memory.poke(memAddress+64, -1);
        do Memory.poke(memAddress+96, 0);
        do Memory.poke(memAddress+128, 30947);
        do Memory.poke(memAddress+160, 1300);
        do Memory.poke(memAddress+192, 1284);
        do Memory.poke(memAddress+224, 1284);
        do Memory.poke(memAddress+256, 15555);
        do Memory.poke(memAddress+288, 17666);
        do Memory.poke(memAddress+320, 17668);
        do Memory.poke(memAddress+352, 17684);
        do Memory.poke(memAddress+384, 14564);
        do Memory.poke(memAddress+416, 0);
        do Memory.poke(memAddress+448, -1);
        do Memory.poke(memAddress+480, 0);
        return;
    }

    function void drawLogoG1(int location) {
        var int memAddress;
        let memAddress = 16384+location;
        do Memory.poke(memAddress+0, 0);
        do Memory.poke(memAddress+32, 0);
        do Memory.poke(memAddress+64, 0);
        do Memory.poke(memAddress+96, 0);
        do Memory.poke(memAddress+128, 0);
        do Memory.poke(memAddress+160, 0);
        do Memory.poke(memAddress+192, 0);
        do Memory.poke(memAddress+224, 0);
        do Memory.poke(memAddress+256, 0);
        do Memory.poke(memAddress+288, 248);
        do Memory.poke(memAddress+320, 248);
        do Memory.poke(memAddress+352, 249);
        do Memory.poke(memAddress+384, 251);
        do Memory.poke(memAddress+416, 123);
        do Memory.poke(memAddress+448, 127);
        do Memory.poke(memAddress+480, 127);
        return;
    }

    function void drawLogoG2(int location) {
        var int memAddress;
        let memAddress = 16384+location;
        do Memory.poke(memAddress+0, 127);
        do Memory.poke(memAddress+32, 126);
        do Memory.poke(memAddress+64, 254);
        do Memory.poke(memAddress+96, 252);
        do Memory.poke(memAddress+128, 248);
        do Memory.poke(memAddress+160, 248);
        do Memory.poke(memAddress+192, 0);
        do Memory.poke(memAddress+224, 0);
        do Memory.poke(memAddress+256, 0);
        do Memory.poke(memAddress+288, 0);
        do Memory.poke(memAddress+320, 0);
        do Memory.poke(memAddress+352, 0);
        do Memory.poke(memAddress+384, 0);
        do Memory.poke(memAddress+416, 0);
        do Memory.poke(memAddress+448, 0);
        do Memory.poke(memAddress+480, 0);
        return;
    }

    function void drawLogoG3(int location) {
        var int memAddress;
        let memAddress = 16384+location;
        do Memory.poke(memAddress+0, 0);
        do Memory.poke(memAddress+32, 0);
        do Memory.poke(memAddress+64, 63);
        do Memory.poke(memAddress+96, 0);
        do Memory.poke(memAddress+128, 14);
        do Memory.poke(memAddress+160, 17);
        do Memory.poke(memAddress+192, 17);
        do Memory.poke(memAddress+224, 25);
        do Memory.poke(memAddress+256, 21);
        do Memory.poke(memAddress+288, 19);
        do Memory.poke(memAddress+320, 17);
        do Memory.poke(memAddress+352, 17);
        do Memory.poke(memAddress+384, 14);
        do Memory.poke(memAddress+416, 0);
        do Memory.poke(memAddress+448, 63);
        do Memory.poke(memAddress+480, 0);
        return;
    }

    function void drawPlayer(int location) {
        var int memAddress;
        let memAddress = 16384+location;
        do Memory.poke(memAddress+0, 0);
        do Memory.poke(memAddress+32, 0);
        do Memory.poke(memAddress+64, 0);
        do Memory.poke(memAddress+96, 0);
        do Memory.poke(memAddress+128, 2016);
        do Memory.poke(memAddress+160, 4080);
        do Memory.poke(memAddress+192, 3504);
        do Memory.poke(memAddress+224, 7608);
        do Memory.poke(memAddress+256, 8184);
        do Memory.poke(memAddress+288, 3024);
        do Memory.poke(memAddress+320, 3120);
        do Memory.poke(memAddress+352, 2016);
        do Memory.poke(memAddress+384, 0);
        do Memory.poke(memAddress+416, 0);
        do Memory.poke(memAddress+448, 0);
        do Memory.poke(memAddress+480, 0);
        return;
    }

    function void drawWall(int location) {
        var int memAddress;
        let memAddress = 16384+location;
        do Memory.poke(memAddress+0, 0);
        do Memory.poke(memAddress+32, 32766);
        do Memory.poke(memAddress+64, 16386);
        do Memory.poke(memAddress+96, 19114);
        do Memory.poke(memAddress+128, 21842);
        do Memory.poke(memAddress+160, 19114);
        do Memory.poke(memAddress+192, 21842);
        do Memory.poke(memAddress+224, 19114);
        do Memory.poke(memAddress+256, 21842);
        do Memory.poke(memAddress+288, 19114);
        do Memory.poke(memAddress+320, 21842);
        do Memory.poke(memAddress+352, 19114);
        do Memory.poke(memAddress+384, 21842);
        do Memory.poke(memAddress+416, 16386);
        do Memory.poke(memAddress+448, 32766);
        do Memory.poke(memAddress+480, 0);
        return;
    }

}
`

const Sokoban_State_Jack string = `class State {
    static int StatePlaying, StateLevelComplete, StateGameComplete;

    function void init() {
        let StatePlaying = 0;
        let StateLevelComplete = 1;
        let StateGameComplete = 2;

        return;
    }

    function int StatePlaying() {
        return StatePlaying;
    }

    function int StateLevelComplete() {
        return StateLevelComplete;
    }

    function int StateGameComplete() {
        return StateGameComplete;
    }
}
`

const Sokoban_Utils_Jack string = `class Utils {
    // the modulo operation returns the remainder of a division e.g. mod(5, 2) would return 1, because 5 divided by 2 has a quotient of 2 and a remainder of 1... mod(a,b) = a-b*(a/b)
    function int mod(int dividend, int divisor) {
        var int quotient;

        let quotient = dividend / divisor;

        return (dividend - (quotient * divisor));
    }
}
`

const Sokoban_View_Jack string = `class View {
    field Model model;
    field String controlInfoTitle;
    field String controlInfoMove;
    field String controlInfoReset;
    field String controlInfoQuit;
    field String controlInfoNext;
    field String controlInfoRestart;
    field String messageLevelComplete;
    field String messageGameComplete;
    field String messageCongratulations;

    constructor View new(Model aModel) {
        let model = aModel;
        let controlInfoTitle = "---Controls---";
        let controlInfoMove = "Cursors:  Move";
        let controlInfoReset = "R:       Reset";
        let controlInfoQuit = "Escape:   Quit";
        let controlInfoNext = "Space:    Next";
        let controlInfoRestart = "Space: Restart";
        let messageLevelComplete = "LEVEL COMPLETE";
        let messageGameComplete = "GAME COMPLETE!";
        let messageCongratulations = "CONGRATULATIONS!";

        return this;
    }

    method void draw() {
        var int x, y;
        var Board board;
        var Player player;

        let board = model.getBoard();
        let player = board.getPlayer();

        // not very mvc, but the more usual "fully clear and redraw the screen on each frame" approach doesn't really suit the hack platform... 
        if (model.isScreenDirty()) {
            do Screen.setColor(false);
            do Screen.drawRectangle(0, 0, 511, 255);
        }

        do drawLogo();

        if (model.getState() = State.StatePlaying()) {
            do drawBoard();

            do drawLevelInfo();

            // draw "playing" control info
            do Output.moveCursor(17, 48);
            do Output.printString(controlInfoTitle);
            do Output.moveCursor(19, 48);
            do Output.printString(controlInfoMove);
            do Output.moveCursor(20, 48);
            do Output.printString(controlInfoReset);
            do Output.moveCursor(21, 48);
            do Output.printString(controlInfoQuit);
        }

        if (model.getState() = State.StateLevelComplete()) {
            do drawBoard();

            do drawLevelInfo();

            // draw flashing level complete message
            if (model.getTickAccumulator() < 10) {
                do Output.moveCursor(12, 48);
                do Output.printString(messageLevelComplete);
            } else {
                do Screen.setColor(false);
                do Screen.drawRectangle(384, 133, 494, 142);
            }

            // draw "level complete" control info
            do Output.moveCursor(17, 48);
            do Output.printString(controlInfoTitle);
            do Output.moveCursor(19, 48);
            do Output.printString(controlInfoNext);
            do Output.moveCursor(21, 48);
            do Output.printString(controlInfoQuit);
        }

        if (model.getState() = State.StateGameComplete()) {
            // draw game complete message
            do Output.moveCursor(10, 16);
            do Output.printString(messageGameComplete);

            // draw alternating congratulations message and smiley face border
            if (model.getTickAccumulator() < 10) {
                do Output.moveCursor(12, 15);
                do Output.printString(messageCongratulations);
                do drawGameCompleteBorder(false);
            } else {
                do Screen.setColor(false);
                do Screen.drawRectangle(120, 133, 244, 141);
                do drawGameCompleteBorder(true);
            }

            // draw "game complete" control info
            do Output.moveCursor(17, 48);
            do Output.printString(controlInfoTitle);
            do Output.moveCursor(19, 48);
            do Output.printString(controlInfoRestart);
            do Output.moveCursor(21, 48);
            do Output.printString(controlInfoQuit);
        }

        return;
    }

    method void drawLogo() {
        do Sprites.drawLogoA1(56 + (0 * 512));
        do Sprites.drawLogoB1(57 + (0 * 512));
        do Sprites.drawLogoC1(58 + (0 * 512));
        do Sprites.drawLogoD1(59 + (0 * 512));
        do Sprites.drawLogoE1(60 + (0 * 512));
        do Sprites.drawLogoF1(61 + (0 * 512));
        do Sprites.drawLogoG1(62 + (0 * 512));
        do Sprites.drawLogoA2(56 + (1 * 512));
        do Sprites.drawLogoB2(57 + (1 * 512));
        do Sprites.drawLogoC2(58 + (1 * 512));
        do Sprites.drawLogoD2(59 + (1 * 512));
        do Sprites.drawLogoE2(60 + (1 * 512));
        do Sprites.drawLogoF2(61 + (1 * 512));
        do Sprites.drawLogoG2(62 + (1 * 512));
        do Sprites.drawLogoA3(56 + (2 * 512));
        do Sprites.drawLogoB3(57 + (2 * 512));
        do Sprites.drawLogoC3(58 + (2 * 512));
        do Sprites.drawLogoD3(59 + (2 * 512));
        do Sprites.drawLogoE3(60 + (2 * 512));
        do Sprites.drawLogoF3(61 + (2 * 512));
        do Sprites.drawLogoG3(62 + (2 * 512));

        return;
    }

    method void drawBoard() {
        var Board board;
        var int x, y, width, height, boardOffsetX, boardOffsetY;
        var Player player;
        var Cell cell;

        let board = model.getBoard();
        let x = 0;
        let y = 0;
        let width = board.getWidth();
        let height = board.getHeight();
        let boardOffsetX = ((22 - width) / 2) + 1;
		let boardOffsetY = ((14 - height) / 2) + 1;
        let player = board.getPlayer();

        while (y < height) {
            while (x < width) {
                let cell = board.getCell(x, y);
                if ((player.getX() = x) & (player.getY() = y)) {
                    if (cell.getTypeOf() = CellType.CellTypeGoal()) {
                        do Sprites.drawGoalPlayer(boardOffsetX + x + ((boardOffsetY + y) * 512));
                    }
                    else {
                        do Sprites.drawPlayer(boardOffsetX + x + ((boardOffsetY + y) * 512));
                    }
                }
                else {
                    if (cell.getTypeOf() = CellType.CellTypeNone()) {
                        if (cell.getHasBox()) {
                            do Sprites.drawBox(boardOffsetX + x + ((boardOffsetY + y) * 512));
                        }
                        else {
                            do Sprites.drawClear(boardOffsetX + x + ((boardOffsetY + y) * 512));
                        }
                    }
                    if (cell.getTypeOf() = CellType.CellTypeGoal()) {
                        if (cell.getHasBox()) {
                            do Sprites.drawGoalBox(boardOffsetX + x + ((boardOffsetY + y) * 512));
                        }
                        else {
                            do Sprites.drawGoal(boardOffsetX + x + ((boardOffsetY + y) * 512));
                        }
                    }
                    if (cell.getTypeOf() = CellType.CellTypeWall()) {
                        do Sprites.drawWall(boardOffsetX + x + ((boardOffsetY + y) * 512));
                    }
                }
                let x = x + 1;
            }
            let y = y + 1;
            let x = 0;
        }

        return;
    }

    method void drawLevelInfo() {
        var LevelManager levelManager;
        var String levelInfo;
        var int currentLevelNumber;
        var Array currentLevelNumberDigits;
        var int finalLevelNumber;
        var Array finalLevelNumberDigits;

        let levelManager = model.getLevelManager();
        let levelInfo = "Level XX of YY";

        let currentLevelNumber = levelManager.getCurrentLevelNumber();
        let currentLevelNumberDigits = Array.new(2);
        let currentLevelNumberDigits[0] = currentLevelNumber/10;
        let currentLevelNumberDigits[1] = Utils.mod(currentLevelNumber, 10);
        do levelInfo.setCharAt(6, 48 + currentLevelNumberDigits[0]); // 0 to 9 = 48 to 57
        do levelInfo.setCharAt(7, 48 + currentLevelNumberDigits[1]); // 0 to 9 = 48 to 57

        let finalLevelNumber = levelManager.getFinalLevelNumber();
        let finalLevelNumberDigits = Array.new(2);
        let finalLevelNumberDigits[0] = finalLevelNumber/10;
        let finalLevelNumberDigits[1] = Utils.mod(finalLevelNumber, 10);
        do levelInfo.setCharAt(12, 48 + finalLevelNumberDigits[0]); // 0 to 9 = 48 to 57
        do levelInfo.setCharAt(13, 48 + finalLevelNumberDigits[1]); // 0 to 9 = 48 to 57

        do Output.moveCursor(7, 48);
        do Output.printString(levelInfo);

        do levelInfo.dispose();
        do currentLevelNumberDigits.dispose();
        do finalLevelNumberDigits.dispose();

        return;
    }

    method void drawGameCompleteBorder(boolean aShowSmileyFaces) {
        var int x, y;

        let x = 1;
        let y = 1;

        while (y < 15) {
            while (x < 23) {
                if ((x = 1) | (x = 22) | (y = 1) | (y = 14)) {
                    if (aShowSmileyFaces) {
                        do Sprites.drawPlayer(x + (y * 512));
                    }
                    else {
                        do Sprites.drawClear(x + (y * 512));
                    }
                }
                let x = x + 1;
            }
            let y = y + 1;
            let x = 1;
        }

        return;
    }

    method void dispose() {
        do Memory.deAlloc(this);

        return;
    }
}
`

const Square_Main_Jack string = `class Main {
    static boolean test;
    function void main() {
        var SquareGame game;
        let game = SquareGame.new();
        do game.run();
        do game.dispose();
        return;
    }

    function void more() {
        var int i, j;
        var String s;
        var Array a;
        if (false) {
            let s = "string constant";
            let s = null;
            let a[1] = a[2];
        }
        else {
            let i = i * (-j);
            let j = j / (-2);
            let i = i | j;
        }
        return;
    }
}
`

const Square_Square_Jack string = `class Square {
    field int x, y;
    field int size;

    constructor Square new(int Ax, int Ay, int Asize) {
        let x = Ax;
        let y = Ay;
        let size = Asize;
        do draw();
        return this;
    }

    method void dispose() {
        do Memory.deAlloc(this);
        return;
    }

    method void draw() {
        do Screen.setColor(true);
        do Screen.drawRectangle(x, y, x + size, y + size);
        return;
    }

    method void erase() {
        do Screen.setColor(false);
        do Screen.drawRectangle(x, y, x + size, y + size);
        return;
    }

    method void incSize() {
        if (((y + size) < 254) & ((x + size) < 510)) {
            do erase();
            let size = size + 2;
            do draw();
        }
        return;
    }

    method void decSize() {
        if (size > 2) {
            do erase();
            let size = size - 2;
            do draw();
        }
        return;
    }

    method void moveUp() {
        if (y > 1) {
            do Screen.setColor(false);
            do Screen.drawRectangle(x, (y + size) - 1, x + size, y + size);
            let y = y - 2;
            do Screen.setColor(true);
            do Screen.drawRectangle(x, y, x + size, y + 1);
        }
        return;
    }

    method void moveDown() {
        if ((y + size) < 254) {
            do Screen.setColor(false);
            do Screen.drawRectangle(x, y, x + size, y + 1);
            let y = y + 2;
            do Screen.setColor(true);
            do Screen.drawRectangle(x, (y + size) - 1, x + size, y + size);
        }
        return;
    }

    method void moveLeft() {
        if (x > 1) {
            do Screen.setColor(false);
            do Screen.drawRectangle((x + size) - 1, y, x + size, y + size);
            let x = x - 2;
            do Screen.setColor(true);
            do Screen.drawRectangle(x, y, x + 1, y + size);
        }
        return;
    }

    method void moveRight() {
        if ((x + size) < 510) {
            do Screen.setColor(false);
            do Screen.drawRectangle(x, y, x + 1, y + size);
            let x = x + 2;
            do Screen.setColor(true);
            do Screen.drawRectangle((x + size) - 1, y, x + size, y + size);
        }
        return;
    }
}
`

const Square_SquareGame_Jack string = `class SquareGame {
    field Square square;
    field int direction;

    constructor SquareGame new() {
        let square = Square.new(0, 0, 30);
        let direction = 0;
        return this;
    }

    method void dispose() {
        do square.dispose();
        do Memory.deAlloc(this);
        return;
    }

    method void moveSquare() {
        if (direction = 1) { do square.moveUp(); }
        if (direction = 2) { do square.moveDown(); }
        if (direction = 3) { do square.moveLeft(); }
        if (direction = 4) { do square.moveRight(); }
        do Sys.wait(5);
        return;
    }

    method void run() {
        var char key;
        var boolean exit;
        let exit = false;

        while (~exit) {
            while (key = 0) {
                let key = Keyboard.keyPressed();
                do moveSquare();
            }
            if (key = 81)  { let exit = true; }
            if (key = 90)  { do square.decSize(); }
            if (key = 88)  { do square.incSize(); }
            if (key = 131) { let direction = 1; }
            if (key = 133) { let direction = 2; }
            if (key = 130) { let direction = 3; }
            if (key = 132) { let direction = 4; }

            while (~(key = 0)) {
                let key = Keyboard.keyPressed();
                do moveSquare();
            }
        }
        return;
    }
}
`

const SquareExpressionless_Main_Jack string = `class Main {
    static boolean test;

    function void main() {
        var SquareGame game;
        let game = game;
        do game.run();
        do game.dispose();
        return;
    }

    function void more() {
        var boolean b;
        if (b) {
        }
        else {
        }
        return;
    }
}
`

const SquareExpressionless_Square_Jack string = `class Square {
    field int x, y;
    field int size;

    constructor Square new(int Ax, int Ay, int Asize) {
        let x = Ax;
        let y = Ay;
        let size = Asize;
        do draw();
        return x;
    }

    method void dispose() {
        do Memory.deAlloc(this);
        return;
    }

    method void draw() {
        do Screen.setColor(x);
        do Screen.drawRectangle(x, y, x, y);
        return;
    }

    method void erase() {
        do Screen.setColor(x);
        do Screen.drawRectangle(x, y, x, y);
        return;
    }

    method void incSize() {
        if (x) {
            do erase();
            let size = size;
            do draw();
        }
        return;
    }

    method void decSize() {
        if (size) {
            do erase();
            let size = size;
            do draw();
        }
        return;
    }

    method void moveUp() {
        if (y) {
            do Screen.setColor(x);
            do Screen.drawRectangle(x, y, x, y);
            let y = y;
            do Screen.setColor(x);
            do Screen.drawRectangle(x, y, x, y);
        }
        return;
    }

    method void moveDown() {
        if (y) {
            do Screen.setColor(x);
            do Screen.drawRectangle(x, y, x, y);
            let y = y;
            do Screen.setColor(x);
            do Screen.drawRectangle(x, y, x, y);
        }
        return;
    }

    method void moveLeft() {
        if (x) {
            do Screen.setColor(x);
            do Screen.drawRectangle(x, y, x, y);
            let x = x;
            do Screen.setColor(x);
            do Screen.drawRectangle(x, y, x, y);
         }
         return;
    }

    method void moveRight() {
        if (x) {
            do Screen.setColor(x);
            do Screen.drawRectangle(x, y, x, y);
            let x = x;
            do Screen.setColor(x);
            do Screen.drawRectangle(x, y, x, y);
        }
        return;
    }
}
`

const SquareExpressionless_SquareGame_Jack string = `class SquareGame {
    field Square square; 
    field int direction; 

    constructor SquareGame new() {
        let square = square;
        let direction = direction;
        return square;
    }

    method void dispose() {
        do square.dispose();
        do Memory.deAlloc(square);
        return;
    }

    method void moveSquare() {
        if (direction) { do square.moveUp(); }
        if (direction) { do square.moveDown(); }
        if (direction) { do square.moveLeft(); }
        if (direction) { do square.moveRight(); }
        do Sys.wait(direction);
        return;
    }

    method void run() {
        var char key;
        var boolean exit;

        let exit = key;
        while (exit) {
            while (key) {
                let key = key;
                do moveSquare();
             }

            if (key) { let exit = exit; }
            if (key) { do square.decSize(); }
            if (key) { do square.incSize(); }
            if (key) { let direction = exit; }
            if (key) { let direction = key; }
            if (key) { let direction = square; }
            if (key) { let direction = direction; }

            while (key) {
                let key = key;
                do moveSquare();
            }
        }
        return;
    }
}
`

// ----- T_XML -----

const Array_MainT_XML string = `<tokens>
<keyword> class </keyword>
<identifier> Main </identifier>
<symbol> { </symbol>
<keyword> function </keyword>
<keyword> void </keyword>
<identifier> main </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> var </keyword>
<identifier> Array </identifier>
<identifier> a </identifier>
<symbol> ; </symbol>
<keyword> var </keyword>
<keyword> int </keyword>
<identifier> length </identifier>
<symbol> ; </symbol>
<keyword> var </keyword>
<keyword> int </keyword>
<identifier> i </identifier>
<symbol> , </symbol>
<identifier> sum </identifier>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> length </identifier>
<symbol> = </symbol>
<identifier> Keyboard </identifier>
<symbol> . </symbol>
<identifier> readInt </identifier>
<symbol> ( </symbol>
<stringConstant> HOW MANY NUMBERS?  </stringConstant>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> a </identifier>
<symbol> = </symbol>
<identifier> Array </identifier>
<symbol> . </symbol>
<identifier> new </identifier>
<symbol> ( </symbol>
<identifier> length </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> i </identifier>
<symbol> = </symbol>
<integerConstant> 0 </integerConstant>
<symbol> ; </symbol>
<keyword> while </keyword>
<symbol> ( </symbol>
<identifier> i </identifier>
<symbol> &lt; </symbol>
<identifier> length </identifier>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> let </keyword>
<identifier> a </identifier>
<symbol> [ </symbol>
<identifier> i </identifier>
<symbol> ] </symbol>
<symbol> = </symbol>
<identifier> Keyboard </identifier>
<symbol> . </symbol>
<identifier> readInt </identifier>
<symbol> ( </symbol>
<stringConstant> ENTER THE NEXT NUMBER:  </stringConstant>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> i </identifier>
<symbol> = </symbol>
<identifier> i </identifier>
<symbol> + </symbol>
<integerConstant> 1 </integerConstant>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> let </keyword>
<identifier> i </identifier>
<symbol> = </symbol>
<integerConstant> 0 </integerConstant>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> sum </identifier>
<symbol> = </symbol>
<integerConstant> 0 </integerConstant>
<symbol> ; </symbol>
<keyword> while </keyword>
<symbol> ( </symbol>
<identifier> i </identifier>
<symbol> &lt; </symbol>
<identifier> length </identifier>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> let </keyword>
<identifier> sum </identifier>
<symbol> = </symbol>
<identifier> sum </identifier>
<symbol> + </symbol>
<identifier> a </identifier>
<symbol> [ </symbol>
<identifier> i </identifier>
<symbol> ] </symbol>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> i </identifier>
<symbol> = </symbol>
<identifier> i </identifier>
<symbol> + </symbol>
<integerConstant> 1 </integerConstant>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> do </keyword>
<identifier> Output </identifier>
<symbol> . </symbol>
<identifier> printString </identifier>
<symbol> ( </symbol>
<stringConstant> THE AVERAGE IS:  </stringConstant>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Output </identifier>
<symbol> . </symbol>
<identifier> printInt </identifier>
<symbol> ( </symbol>
<identifier> sum </identifier>
<symbol> / </symbol>
<identifier> length </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Output </identifier>
<symbol> . </symbol>
<identifier> println </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> return </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<symbol> } </symbol>
</tokens>
`

const HelloWorld_MainT_XML string = `<tokens>
<keyword> class </keyword>
<identifier> Main </identifier>
<symbol> { </symbol>
<keyword> function </keyword>
<keyword> void </keyword>
<identifier> main </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> Output </identifier>
<symbol> . </symbol>
<identifier> printString </identifier>
<symbol> ( </symbol>
<stringConstant> Hello World </stringConstant>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> return </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<symbol> } </symbol>
</tokens>
`

const Square_MainT_XML string = `<tokens>
<keyword> class </keyword>
<identifier> Main </identifier>
<symbol> { </symbol>
<keyword> static </keyword>
<keyword> boolean </keyword>
<identifier> test </identifier>
<symbol> ; </symbol>
<keyword> function </keyword>
<keyword> void </keyword>
<identifier> main </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> var </keyword>
<identifier> SquareGame </identifier>
<identifier> game </identifier>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> game </identifier>
<symbol> = </symbol>
<identifier> SquareGame </identifier>
<symbol> . </symbol>
<identifier> new </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> game </identifier>
<symbol> . </symbol>
<identifier> run </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> game </identifier>
<symbol> . </symbol>
<identifier> dispose </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> return </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> function </keyword>
<keyword> void </keyword>
<identifier> more </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> var </keyword>
<keyword> int </keyword>
<identifier> i </identifier>
<symbol> , </symbol>
<identifier> j </identifier>
<symbol> ; </symbol>
<keyword> var </keyword>
<identifier> String </identifier>
<identifier> s </identifier>
<symbol> ; </symbol>
<keyword> var </keyword>
<identifier> Array </identifier>
<identifier> a </identifier>
<symbol> ; </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<keyword> false </keyword>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> let </keyword>
<identifier> s </identifier>
<symbol> = </symbol>
<stringConstant> string constant </stringConstant>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> s </identifier>
<symbol> = </symbol>
<keyword> null </keyword>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> a </identifier>
<symbol> [ </symbol>
<integerConstant> 1 </integerConstant>
<symbol> ] </symbol>
<symbol> = </symbol>
<identifier> a </identifier>
<symbol> [ </symbol>
<integerConstant> 2 </integerConstant>
<symbol> ] </symbol>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> else </keyword>
<symbol> { </symbol>
<keyword> let </keyword>
<identifier> i </identifier>
<symbol> = </symbol>
<identifier> i </identifier>
<symbol> * </symbol>
<symbol> ( </symbol>
<symbol> - </symbol>
<identifier> j </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> j </identifier>
<symbol> = </symbol>
<identifier> j </identifier>
<symbol> / </symbol>
<symbol> ( </symbol>
<symbol> - </symbol>
<integerConstant> 2 </integerConstant>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> i </identifier>
<symbol> = </symbol>
<identifier> i </identifier>
<symbol> | </symbol>
<identifier> j </identifier>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> return </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<symbol> } </symbol>
</tokens>
`

const Square_SquareT_XML string = `<tokens>
<keyword> class </keyword>
<identifier> Square </identifier>
<symbol> { </symbol>
<keyword> field </keyword>
<keyword> int </keyword>
<identifier> x </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> ; </symbol>
<keyword> field </keyword>
<keyword> int </keyword>
<identifier> size </identifier>
<symbol> ; </symbol>
<keyword> constructor </keyword>
<identifier> Square </identifier>
<identifier> new </identifier>
<symbol> ( </symbol>
<keyword> int </keyword>
<identifier> Ax </identifier>
<symbol> , </symbol>
<keyword> int </keyword>
<identifier> Ay </identifier>
<symbol> , </symbol>
<keyword> int </keyword>
<identifier> Asize </identifier>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> let </keyword>
<identifier> x </identifier>
<symbol> = </symbol>
<identifier> Ax </identifier>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> y </identifier>
<symbol> = </symbol>
<identifier> Ay </identifier>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> size </identifier>
<symbol> = </symbol>
<identifier> Asize </identifier>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> draw </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> return </keyword>
<keyword> this </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> dispose </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> Memory </identifier>
<symbol> . </symbol>
<identifier> deAlloc </identifier>
<symbol> ( </symbol>
<keyword> this </keyword>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> return </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> draw </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<keyword> true </keyword>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> , </symbol>
<identifier> x </identifier>
<symbol> + </symbol>
<identifier> size </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> + </symbol>
<identifier> size </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> return </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> erase </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<keyword> false </keyword>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> , </symbol>
<identifier> x </identifier>
<symbol> + </symbol>
<identifier> size </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> + </symbol>
<identifier> size </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> return </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> incSize </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<symbol> ( </symbol>
<symbol> ( </symbol>
<identifier> y </identifier>
<symbol> + </symbol>
<identifier> size </identifier>
<symbol> ) </symbol>
<symbol> &lt; </symbol>
<integerConstant> 254 </integerConstant>
<symbol> ) </symbol>
<symbol> &amp; </symbol>
<symbol> ( </symbol>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> + </symbol>
<identifier> size </identifier>
<symbol> ) </symbol>
<symbol> &lt; </symbol>
<integerConstant> 510 </integerConstant>
<symbol> ) </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> erase </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> size </identifier>
<symbol> = </symbol>
<identifier> size </identifier>
<symbol> + </symbol>
<integerConstant> 2 </integerConstant>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> draw </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> return </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> decSize </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> size </identifier>
<symbol> &gt; </symbol>
<integerConstant> 2 </integerConstant>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> erase </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> size </identifier>
<symbol> = </symbol>
<identifier> size </identifier>
<symbol> - </symbol>
<integerConstant> 2 </integerConstant>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> draw </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> return </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> moveUp </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> y </identifier>
<symbol> &gt; </symbol>
<integerConstant> 1 </integerConstant>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<keyword> false </keyword>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> , </symbol>
<symbol> ( </symbol>
<identifier> y </identifier>
<symbol> + </symbol>
<identifier> size </identifier>
<symbol> ) </symbol>
<symbol> - </symbol>
<integerConstant> 1 </integerConstant>
<symbol> , </symbol>
<identifier> x </identifier>
<symbol> + </symbol>
<identifier> size </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> + </symbol>
<identifier> size </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> y </identifier>
<symbol> = </symbol>
<identifier> y </identifier>
<symbol> - </symbol>
<integerConstant> 2 </integerConstant>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<keyword> true </keyword>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> , </symbol>
<identifier> x </identifier>
<symbol> + </symbol>
<identifier> size </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> + </symbol>
<integerConstant> 1 </integerConstant>
<symbol> ) </symbol>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> return </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> moveDown </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<symbol> ( </symbol>
<identifier> y </identifier>
<symbol> + </symbol>
<identifier> size </identifier>
<symbol> ) </symbol>
<symbol> &lt; </symbol>
<integerConstant> 254 </integerConstant>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<keyword> false </keyword>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> , </symbol>
<identifier> x </identifier>
<symbol> + </symbol>
<identifier> size </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> + </symbol>
<integerConstant> 1 </integerConstant>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> y </identifier>
<symbol> = </symbol>
<identifier> y </identifier>
<symbol> + </symbol>
<integerConstant> 2 </integerConstant>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<keyword> true </keyword>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> , </symbol>
<symbol> ( </symbol>
<identifier> y </identifier>
<symbol> + </symbol>
<identifier> size </identifier>
<symbol> ) </symbol>
<symbol> - </symbol>
<integerConstant> 1 </integerConstant>
<symbol> , </symbol>
<identifier> x </identifier>
<symbol> + </symbol>
<identifier> size </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> + </symbol>
<identifier> size </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> return </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> moveLeft </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> &gt; </symbol>
<integerConstant> 1 </integerConstant>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<keyword> false </keyword>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> + </symbol>
<identifier> size </identifier>
<symbol> ) </symbol>
<symbol> - </symbol>
<integerConstant> 1 </integerConstant>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> , </symbol>
<identifier> x </identifier>
<symbol> + </symbol>
<identifier> size </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> + </symbol>
<identifier> size </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> x </identifier>
<symbol> = </symbol>
<identifier> x </identifier>
<symbol> - </symbol>
<integerConstant> 2 </integerConstant>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<keyword> true </keyword>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> , </symbol>
<identifier> x </identifier>
<symbol> + </symbol>
<integerConstant> 1 </integerConstant>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> + </symbol>
<identifier> size </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> return </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> moveRight </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> + </symbol>
<identifier> size </identifier>
<symbol> ) </symbol>
<symbol> &lt; </symbol>
<integerConstant> 510 </integerConstant>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<keyword> false </keyword>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> , </symbol>
<identifier> x </identifier>
<symbol> + </symbol>
<integerConstant> 1 </integerConstant>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> + </symbol>
<identifier> size </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> x </identifier>
<symbol> = </symbol>
<identifier> x </identifier>
<symbol> + </symbol>
<integerConstant> 2 </integerConstant>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<keyword> true </keyword>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> + </symbol>
<identifier> size </identifier>
<symbol> ) </symbol>
<symbol> - </symbol>
<integerConstant> 1 </integerConstant>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> , </symbol>
<identifier> x </identifier>
<symbol> + </symbol>
<identifier> size </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> + </symbol>
<identifier> size </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> return </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<symbol> } </symbol>
</tokens>
`

const Square_SquareGameT_XML string = `<tokens>
<keyword> class </keyword>
<identifier> SquareGame </identifier>
<symbol> { </symbol>
<keyword> field </keyword>
<identifier> Square </identifier>
<identifier> square </identifier>
<symbol> ; </symbol>
<keyword> field </keyword>
<keyword> int </keyword>
<identifier> direction </identifier>
<symbol> ; </symbol>
<keyword> constructor </keyword>
<identifier> SquareGame </identifier>
<identifier> new </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> let </keyword>
<identifier> square </identifier>
<symbol> = </symbol>
<identifier> Square </identifier>
<symbol> . </symbol>
<identifier> new </identifier>
<symbol> ( </symbol>
<integerConstant> 0 </integerConstant>
<symbol> , </symbol>
<integerConstant> 0 </integerConstant>
<symbol> , </symbol>
<integerConstant> 30 </integerConstant>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> direction </identifier>
<symbol> = </symbol>
<integerConstant> 0 </integerConstant>
<symbol> ; </symbol>
<keyword> return </keyword>
<keyword> this </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> dispose </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> square </identifier>
<symbol> . </symbol>
<identifier> dispose </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Memory </identifier>
<symbol> . </symbol>
<identifier> deAlloc </identifier>
<symbol> ( </symbol>
<keyword> this </keyword>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> return </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> moveSquare </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> direction </identifier>
<symbol> = </symbol>
<integerConstant> 1 </integerConstant>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> square </identifier>
<symbol> . </symbol>
<identifier> moveUp </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> direction </identifier>
<symbol> = </symbol>
<integerConstant> 2 </integerConstant>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> square </identifier>
<symbol> . </symbol>
<identifier> moveDown </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> direction </identifier>
<symbol> = </symbol>
<integerConstant> 3 </integerConstant>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> square </identifier>
<symbol> . </symbol>
<identifier> moveLeft </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> direction </identifier>
<symbol> = </symbol>
<integerConstant> 4 </integerConstant>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> square </identifier>
<symbol> . </symbol>
<identifier> moveRight </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> do </keyword>
<identifier> Sys </identifier>
<symbol> . </symbol>
<identifier> wait </identifier>
<symbol> ( </symbol>
<integerConstant> 5 </integerConstant>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> return </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> run </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> var </keyword>
<keyword> char </keyword>
<identifier> key </identifier>
<symbol> ; </symbol>
<keyword> var </keyword>
<keyword> boolean </keyword>
<identifier> exit </identifier>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> exit </identifier>
<symbol> = </symbol>
<keyword> false </keyword>
<symbol> ; </symbol>
<keyword> while </keyword>
<symbol> ( </symbol>
<symbol> ~ </symbol>
<identifier> exit </identifier>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> while </keyword>
<symbol> ( </symbol>
<identifier> key </identifier>
<symbol> = </symbol>
<integerConstant> 0 </integerConstant>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> let </keyword>
<identifier> key </identifier>
<symbol> = </symbol>
<identifier> Keyboard </identifier>
<symbol> . </symbol>
<identifier> keyPressed </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> moveSquare </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> key </identifier>
<symbol> = </symbol>
<integerConstant> 81 </integerConstant>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> let </keyword>
<identifier> exit </identifier>
<symbol> = </symbol>
<keyword> true </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> key </identifier>
<symbol> = </symbol>
<integerConstant> 90 </integerConstant>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> square </identifier>
<symbol> . </symbol>
<identifier> decSize </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> key </identifier>
<symbol> = </symbol>
<integerConstant> 88 </integerConstant>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> square </identifier>
<symbol> . </symbol>
<identifier> incSize </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> key </identifier>
<symbol> = </symbol>
<integerConstant> 131 </integerConstant>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> let </keyword>
<identifier> direction </identifier>
<symbol> = </symbol>
<integerConstant> 1 </integerConstant>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> key </identifier>
<symbol> = </symbol>
<integerConstant> 133 </integerConstant>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> let </keyword>
<identifier> direction </identifier>
<symbol> = </symbol>
<integerConstant> 2 </integerConstant>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> key </identifier>
<symbol> = </symbol>
<integerConstant> 130 </integerConstant>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> let </keyword>
<identifier> direction </identifier>
<symbol> = </symbol>
<integerConstant> 3 </integerConstant>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> key </identifier>
<symbol> = </symbol>
<integerConstant> 132 </integerConstant>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> let </keyword>
<identifier> direction </identifier>
<symbol> = </symbol>
<integerConstant> 4 </integerConstant>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> while </keyword>
<symbol> ( </symbol>
<symbol> ~ </symbol>
<symbol> ( </symbol>
<identifier> key </identifier>
<symbol> = </symbol>
<integerConstant> 0 </integerConstant>
<symbol> ) </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> let </keyword>
<identifier> key </identifier>
<symbol> = </symbol>
<identifier> Keyboard </identifier>
<symbol> . </symbol>
<identifier> keyPressed </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> moveSquare </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<symbol> } </symbol>
<symbol> } </symbol>
<keyword> return </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<symbol> } </symbol>
</tokens>
`

const SquareExpressionless_MainT_XML string = `<tokens>
<keyword> class </keyword>
<identifier> Main </identifier>
<symbol> { </symbol>
<keyword> static </keyword>
<keyword> boolean </keyword>
<identifier> test </identifier>
<symbol> ; </symbol>
<keyword> function </keyword>
<keyword> void </keyword>
<identifier> main </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> var </keyword>
<identifier> SquareGame </identifier>
<identifier> game </identifier>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> game </identifier>
<symbol> = </symbol>
<identifier> game </identifier>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> game </identifier>
<symbol> . </symbol>
<identifier> run </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> game </identifier>
<symbol> . </symbol>
<identifier> dispose </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> return </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> function </keyword>
<keyword> void </keyword>
<identifier> more </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> var </keyword>
<keyword> boolean </keyword>
<identifier> b </identifier>
<symbol> ; </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> b </identifier>
<symbol> ) </symbol>
<symbol> { </symbol>
<symbol> } </symbol>
<keyword> else </keyword>
<symbol> { </symbol>
<symbol> } </symbol>
<keyword> return </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<symbol> } </symbol>
</tokens>
`

const SquareExpressionless_SquareT_XML string = `<tokens>
<keyword> class </keyword>
<identifier> Square </identifier>
<symbol> { </symbol>
<keyword> field </keyword>
<keyword> int </keyword>
<identifier> x </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> ; </symbol>
<keyword> field </keyword>
<keyword> int </keyword>
<identifier> size </identifier>
<symbol> ; </symbol>
<keyword> constructor </keyword>
<identifier> Square </identifier>
<identifier> new </identifier>
<symbol> ( </symbol>
<keyword> int </keyword>
<identifier> Ax </identifier>
<symbol> , </symbol>
<keyword> int </keyword>
<identifier> Ay </identifier>
<symbol> , </symbol>
<keyword> int </keyword>
<identifier> Asize </identifier>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> let </keyword>
<identifier> x </identifier>
<symbol> = </symbol>
<identifier> Ax </identifier>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> y </identifier>
<symbol> = </symbol>
<identifier> Ay </identifier>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> size </identifier>
<symbol> = </symbol>
<identifier> Asize </identifier>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> draw </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> return </keyword>
<identifier> x </identifier>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> dispose </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> Memory </identifier>
<symbol> . </symbol>
<identifier> deAlloc </identifier>
<symbol> ( </symbol>
<keyword> this </keyword>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> return </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> draw </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> , </symbol>
<identifier> x </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> return </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> erase </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> , </symbol>
<identifier> x </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> return </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> incSize </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> erase </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> size </identifier>
<symbol> = </symbol>
<identifier> size </identifier>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> draw </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> return </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> decSize </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> size </identifier>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> erase </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> size </identifier>
<symbol> = </symbol>
<identifier> size </identifier>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> draw </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> return </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> moveUp </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> y </identifier>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> , </symbol>
<identifier> x </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> y </identifier>
<symbol> = </symbol>
<identifier> y </identifier>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> , </symbol>
<identifier> x </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> return </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> moveDown </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> y </identifier>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> , </symbol>
<identifier> x </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> y </identifier>
<symbol> = </symbol>
<identifier> y </identifier>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> , </symbol>
<identifier> x </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> return </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> moveLeft </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> , </symbol>
<identifier> x </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> x </identifier>
<symbol> = </symbol>
<identifier> x </identifier>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> , </symbol>
<identifier> x </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> return </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> moveRight </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> , </symbol>
<identifier> x </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> x </identifier>
<symbol> = </symbol>
<identifier> x </identifier>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<identifier> x </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> , </symbol>
<identifier> x </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> return </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<symbol> } </symbol>
</tokens>
`

const SquareExpressionless_SquareGameT_XML string = `<tokens>
<keyword> class </keyword>
<identifier> SquareGame </identifier>
<symbol> { </symbol>
<keyword> field </keyword>
<identifier> Square </identifier>
<identifier> square </identifier>
<symbol> ; </symbol>
<keyword> field </keyword>
<keyword> int </keyword>
<identifier> direction </identifier>
<symbol> ; </symbol>
<keyword> constructor </keyword>
<identifier> SquareGame </identifier>
<identifier> new </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> let </keyword>
<identifier> square </identifier>
<symbol> = </symbol>
<identifier> square </identifier>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> direction </identifier>
<symbol> = </symbol>
<identifier> direction </identifier>
<symbol> ; </symbol>
<keyword> return </keyword>
<identifier> square </identifier>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> dispose </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> square </identifier>
<symbol> . </symbol>
<identifier> dispose </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> Memory </identifier>
<symbol> . </symbol>
<identifier> deAlloc </identifier>
<symbol> ( </symbol>
<identifier> square </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> return </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> moveSquare </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> direction </identifier>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> square </identifier>
<symbol> . </symbol>
<identifier> moveUp </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> direction </identifier>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> square </identifier>
<symbol> . </symbol>
<identifier> moveDown </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> direction </identifier>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> square </identifier>
<symbol> . </symbol>
<identifier> moveLeft </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> direction </identifier>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> square </identifier>
<symbol> . </symbol>
<identifier> moveRight </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> do </keyword>
<identifier> Sys </identifier>
<symbol> . </symbol>
<identifier> wait </identifier>
<symbol> ( </symbol>
<identifier> direction </identifier>
<symbol> ) </symbol>
<symbol> ; </symbol>
<keyword> return </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> run </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> var </keyword>
<keyword> char </keyword>
<identifier> key </identifier>
<symbol> ; </symbol>
<keyword> var </keyword>
<keyword> boolean </keyword>
<identifier> exit </identifier>
<symbol> ; </symbol>
<keyword> let </keyword>
<identifier> exit </identifier>
<symbol> = </symbol>
<identifier> key </identifier>
<symbol> ; </symbol>
<keyword> while </keyword>
<symbol> ( </symbol>
<identifier> exit </identifier>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> while </keyword>
<symbol> ( </symbol>
<identifier> key </identifier>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> let </keyword>
<identifier> key </identifier>
<symbol> = </symbol>
<identifier> key </identifier>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> moveSquare </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> key </identifier>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> let </keyword>
<identifier> exit </identifier>
<symbol> = </symbol>
<identifier> exit </identifier>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> key </identifier>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> square </identifier>
<symbol> . </symbol>
<identifier> decSize </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> key </identifier>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> do </keyword>
<identifier> square </identifier>
<symbol> . </symbol>
<identifier> incSize </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> key </identifier>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> let </keyword>
<identifier> direction </identifier>
<symbol> = </symbol>
<identifier> exit </identifier>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> key </identifier>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> let </keyword>
<identifier> direction </identifier>
<symbol> = </symbol>
<identifier> key </identifier>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> key </identifier>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> let </keyword>
<identifier> direction </identifier>
<symbol> = </symbol>
<identifier> square </identifier>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> if </keyword>
<symbol> ( </symbol>
<identifier> key </identifier>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> let </keyword>
<identifier> direction </identifier>
<symbol> = </symbol>
<identifier> direction </identifier>
<symbol> ; </symbol>
<symbol> } </symbol>
<keyword> while </keyword>
<symbol> ( </symbol>
<identifier> key </identifier>
<symbol> ) </symbol>
<symbol> { </symbol>
<keyword> let </keyword>
<identifier> key </identifier>
<symbol> = </symbol>
<identifier> key </identifier>
<symbol> ; </symbol>
<keyword> do </keyword>
<identifier> moveSquare </identifier>
<symbol> ( </symbol>
<symbol> ) </symbol>
<symbol> ; </symbol>
<symbol> } </symbol>
<symbol> } </symbol>
<keyword> return </keyword>
<symbol> ; </symbol>
<symbol> } </symbol>
<symbol> } </symbol>
</tokens>
`

// ----- XML -----

const Array_Main_XML string = `<class>
<keyword> class </keyword>
<identifier> Main </identifier>
<symbol> { </symbol>
<subroutineDec>
<keyword> function </keyword>
<keyword> void </keyword>
<identifier> main </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<varDec>
<keyword> var </keyword>
<identifier> Array </identifier>
<identifier> a </identifier>
<symbol> ; </symbol>
</varDec>
<varDec>
<keyword> var </keyword>
<keyword> int </keyword>
<identifier> length </identifier>
<symbol> ; </symbol>
</varDec>
<varDec>
<keyword> var </keyword>
<keyword> int </keyword>
<identifier> i </identifier>
<symbol> , </symbol>
<identifier> sum </identifier>
<symbol> ; </symbol>
</varDec>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> length </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> Keyboard </identifier>
<symbol> . </symbol>
<identifier> readInt </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<stringConstant> HOW MANY NUMBERS?  </stringConstant>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<letStatement>
<keyword> let </keyword>
<identifier> a </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> Array </identifier>
<symbol> . </symbol>
<identifier> new </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> length </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<letStatement>
<keyword> let </keyword>
<identifier> i </identifier>
<symbol> = </symbol>
<expression>
<term>
<integerConstant> 0 </integerConstant>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<whileStatement>
<keyword> while </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> i </identifier>
</term>
<symbol> &lt; </symbol>
<term>
<identifier> length </identifier>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> a </identifier>
<symbol> [ </symbol>
<expression>
<term>
<identifier> i </identifier>
</term>
</expression>
<symbol> ] </symbol>
<symbol> = </symbol>
<expression>
<term>
<identifier> Keyboard </identifier>
<symbol> . </symbol>
<identifier> readInt </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<stringConstant> ENTER THE NEXT NUMBER:  </stringConstant>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<letStatement>
<keyword> let </keyword>
<identifier> i </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> i </identifier>
</term>
<symbol> + </symbol>
<term>
<integerConstant> 1 </integerConstant>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
</statements>
<symbol> } </symbol>
</whileStatement>
<letStatement>
<keyword> let </keyword>
<identifier> i </identifier>
<symbol> = </symbol>
<expression>
<term>
<integerConstant> 0 </integerConstant>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<letStatement>
<keyword> let </keyword>
<identifier> sum </identifier>
<symbol> = </symbol>
<expression>
<term>
<integerConstant> 0 </integerConstant>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<whileStatement>
<keyword> while </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> i </identifier>
</term>
<symbol> &lt; </symbol>
<term>
<identifier> length </identifier>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> sum </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> sum </identifier>
</term>
<symbol> + </symbol>
<term>
<identifier> a </identifier>
<symbol> [ </symbol>
<expression>
<term>
<identifier> i </identifier>
</term>
</expression>
<symbol> ] </symbol>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<letStatement>
<keyword> let </keyword>
<identifier> i </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> i </identifier>
</term>
<symbol> + </symbol>
<term>
<integerConstant> 1 </integerConstant>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
</statements>
<symbol> } </symbol>
</whileStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Output </identifier>
<symbol> . </symbol>
<identifier> printString </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<stringConstant> THE AVERAGE IS:  </stringConstant>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Output </identifier>
<symbol> . </symbol>
<identifier> printInt </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> sum </identifier>
</term>
<symbol> / </symbol>
<term>
<identifier> length </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Output </identifier>
<symbol> . </symbol>
<identifier> println </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<symbol> } </symbol>
</class>
`

const HelloWorld_Main_XML string = `<class>
<keyword> class </keyword>
<identifier> Main </identifier>
<symbol> { </symbol>
<subroutineDec>
<keyword> function </keyword>
<keyword> void </keyword>
<identifier> main </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> Output </identifier>
<symbol> . </symbol>
<identifier> printString </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<stringConstant> Hello World </stringConstant>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<symbol> } </symbol>
</class>
`

const Square_Main_XML string = `<class>
<keyword> class </keyword>
<identifier> Main </identifier>
<symbol> { </symbol>
<classVarDec>
<keyword> static </keyword>
<keyword> boolean </keyword>
<identifier> test </identifier>
<symbol> ; </symbol>
</classVarDec>
<subroutineDec>
<keyword> function </keyword>
<keyword> void </keyword>
<identifier> main </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<varDec>
<keyword> var </keyword>
<identifier> SquareGame </identifier>
<identifier> game </identifier>
<symbol> ; </symbol>
</varDec>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> game </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> SquareGame </identifier>
<symbol> . </symbol>
<identifier> new </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<doStatement>
<keyword> do </keyword>
<identifier> game </identifier>
<symbol> . </symbol>
<identifier> run </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<doStatement>
<keyword> do </keyword>
<identifier> game </identifier>
<symbol> . </symbol>
<identifier> dispose </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<subroutineDec>
<keyword> function </keyword>
<keyword> void </keyword>
<identifier> more </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<varDec>
<keyword> var </keyword>
<keyword> int </keyword>
<identifier> i </identifier>
<symbol> , </symbol>
<identifier> j </identifier>
<symbol> ; </symbol>
</varDec>
<varDec>
<keyword> var </keyword>
<identifier> String </identifier>
<identifier> s </identifier>
<symbol> ; </symbol>
</varDec>
<varDec>
<keyword> var </keyword>
<identifier> Array </identifier>
<identifier> a </identifier>
<symbol> ; </symbol>
</varDec>
<statements>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<keyword> false </keyword>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> s </identifier>
<symbol> = </symbol>
<expression>
<term>
<stringConstant> string constant </stringConstant>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<letStatement>
<keyword> let </keyword>
<identifier> s </identifier>
<symbol> = </symbol>
<expression>
<term>
<keyword> null </keyword>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<letStatement>
<keyword> let </keyword>
<identifier> a </identifier>
<symbol> [ </symbol>
<expression>
<term>
<integerConstant> 1 </integerConstant>
</term>
</expression>
<symbol> ] </symbol>
<symbol> = </symbol>
<expression>
<term>
<identifier> a </identifier>
<symbol> [ </symbol>
<expression>
<term>
<integerConstant> 2 </integerConstant>
</term>
</expression>
<symbol> ] </symbol>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
</statements>
<symbol> } </symbol>
<keyword> else </keyword>
<symbol> { </symbol>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> i </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> i </identifier>
</term>
<symbol> * </symbol>
<term>
<symbol> ( </symbol>
<expression>
<term>
<symbol> - </symbol>
<term>
<identifier> j </identifier>
</term>
</term>
</expression>
<symbol> ) </symbol>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<letStatement>
<keyword> let </keyword>
<identifier> j </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> j </identifier>
</term>
<symbol> / </symbol>
<term>
<symbol> ( </symbol>
<expression>
<term>
<symbol> - </symbol>
<term>
<integerConstant> 2 </integerConstant>
</term>
</term>
</expression>
<symbol> ) </symbol>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<letStatement>
<keyword> let </keyword>
<identifier> i </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> i </identifier>
</term>
<symbol> | </symbol>
<term>
<identifier> j </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<symbol> } </symbol>
</class>
`

const Square_Square_XML string = `<class>
<keyword> class </keyword>
<identifier> Square </identifier>
<symbol> { </symbol>
<classVarDec>
<keyword> field </keyword>
<keyword> int </keyword>
<identifier> x </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> ; </symbol>
</classVarDec>
<classVarDec>
<keyword> field </keyword>
<keyword> int </keyword>
<identifier> size </identifier>
<symbol> ; </symbol>
</classVarDec>
<subroutineDec>
<keyword> constructor </keyword>
<identifier> Square </identifier>
<identifier> new </identifier>
<symbol> ( </symbol>
<parameterList>
<keyword> int </keyword>
<identifier> Ax </identifier>
<symbol> , </symbol>
<keyword> int </keyword>
<identifier> Ay </identifier>
<symbol> , </symbol>
<keyword> int </keyword>
<identifier> Asize </identifier>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> x </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> Ax </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<letStatement>
<keyword> let </keyword>
<identifier> y </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> Ay </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<letStatement>
<keyword> let </keyword>
<identifier> size </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> Asize </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<doStatement>
<keyword> do </keyword>
<identifier> draw </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<returnStatement>
<keyword> return </keyword>
<expression>
<term>
<keyword> this </keyword>
</term>
</expression>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<subroutineDec>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> dispose </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> Memory </identifier>
<symbol> . </symbol>
<identifier> deAlloc </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<keyword> this </keyword>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<subroutineDec>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> draw </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<keyword> true </keyword>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
<symbol> + </symbol>
<term>
<identifier> size </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
<symbol> + </symbol>
<term>
<identifier> size </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<subroutineDec>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> erase </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<keyword> false </keyword>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
<symbol> + </symbol>
<term>
<identifier> size </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
<symbol> + </symbol>
<term>
<identifier> size </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<subroutineDec>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> incSize </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<statements>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<symbol> ( </symbol>
<expression>
<term>
<symbol> ( </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
<symbol> + </symbol>
<term>
<identifier> size </identifier>
</term>
</expression>
<symbol> ) </symbol>
</term>
<symbol> &lt; </symbol>
<term>
<integerConstant> 254 </integerConstant>
</term>
</expression>
<symbol> ) </symbol>
</term>
<symbol> &amp; </symbol>
<term>
<symbol> ( </symbol>
<expression>
<term>
<symbol> ( </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
<symbol> + </symbol>
<term>
<identifier> size </identifier>
</term>
</expression>
<symbol> ) </symbol>
</term>
<symbol> &lt; </symbol>
<term>
<integerConstant> 510 </integerConstant>
</term>
</expression>
<symbol> ) </symbol>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> erase </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<letStatement>
<keyword> let </keyword>
<identifier> size </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> size </identifier>
</term>
<symbol> + </symbol>
<term>
<integerConstant> 2 </integerConstant>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<doStatement>
<keyword> do </keyword>
<identifier> draw </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<subroutineDec>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> decSize </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<statements>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> size </identifier>
</term>
<symbol> &gt; </symbol>
<term>
<integerConstant> 2 </integerConstant>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> erase </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<letStatement>
<keyword> let </keyword>
<identifier> size </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> size </identifier>
</term>
<symbol> - </symbol>
<term>
<integerConstant> 2 </integerConstant>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<doStatement>
<keyword> do </keyword>
<identifier> draw </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<subroutineDec>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> moveUp </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<statements>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
<symbol> &gt; </symbol>
<term>
<integerConstant> 1 </integerConstant>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<keyword> false </keyword>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<symbol> ( </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
<symbol> + </symbol>
<term>
<identifier> size </identifier>
</term>
</expression>
<symbol> ) </symbol>
</term>
<symbol> - </symbol>
<term>
<integerConstant> 1 </integerConstant>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
<symbol> + </symbol>
<term>
<identifier> size </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
<symbol> + </symbol>
<term>
<identifier> size </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<letStatement>
<keyword> let </keyword>
<identifier> y </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
<symbol> - </symbol>
<term>
<integerConstant> 2 </integerConstant>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<keyword> true </keyword>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
<symbol> + </symbol>
<term>
<identifier> size </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
<symbol> + </symbol>
<term>
<integerConstant> 1 </integerConstant>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<subroutineDec>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> moveDown </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<statements>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<symbol> ( </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
<symbol> + </symbol>
<term>
<identifier> size </identifier>
</term>
</expression>
<symbol> ) </symbol>
</term>
<symbol> &lt; </symbol>
<term>
<integerConstant> 254 </integerConstant>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<keyword> false </keyword>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
<symbol> + </symbol>
<term>
<identifier> size </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
<symbol> + </symbol>
<term>
<integerConstant> 1 </integerConstant>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<letStatement>
<keyword> let </keyword>
<identifier> y </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
<symbol> + </symbol>
<term>
<integerConstant> 2 </integerConstant>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<keyword> true </keyword>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<symbol> ( </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
<symbol> + </symbol>
<term>
<identifier> size </identifier>
</term>
</expression>
<symbol> ) </symbol>
</term>
<symbol> - </symbol>
<term>
<integerConstant> 1 </integerConstant>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
<symbol> + </symbol>
<term>
<identifier> size </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
<symbol> + </symbol>
<term>
<identifier> size </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<subroutineDec>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> moveLeft </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<statements>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
<symbol> &gt; </symbol>
<term>
<integerConstant> 1 </integerConstant>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<keyword> false </keyword>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<symbol> ( </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
<symbol> + </symbol>
<term>
<identifier> size </identifier>
</term>
</expression>
<symbol> ) </symbol>
</term>
<symbol> - </symbol>
<term>
<integerConstant> 1 </integerConstant>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
<symbol> + </symbol>
<term>
<identifier> size </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
<symbol> + </symbol>
<term>
<identifier> size </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<letStatement>
<keyword> let </keyword>
<identifier> x </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
<symbol> - </symbol>
<term>
<integerConstant> 2 </integerConstant>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<keyword> true </keyword>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
<symbol> + </symbol>
<term>
<integerConstant> 1 </integerConstant>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
<symbol> + </symbol>
<term>
<identifier> size </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<subroutineDec>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> moveRight </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<statements>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<symbol> ( </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
<symbol> + </symbol>
<term>
<identifier> size </identifier>
</term>
</expression>
<symbol> ) </symbol>
</term>
<symbol> &lt; </symbol>
<term>
<integerConstant> 510 </integerConstant>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<keyword> false </keyword>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
<symbol> + </symbol>
<term>
<integerConstant> 1 </integerConstant>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
<symbol> + </symbol>
<term>
<identifier> size </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<letStatement>
<keyword> let </keyword>
<identifier> x </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
<symbol> + </symbol>
<term>
<integerConstant> 2 </integerConstant>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<keyword> true </keyword>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<symbol> ( </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
<symbol> + </symbol>
<term>
<identifier> size </identifier>
</term>
</expression>
<symbol> ) </symbol>
</term>
<symbol> - </symbol>
<term>
<integerConstant> 1 </integerConstant>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
<symbol> + </symbol>
<term>
<identifier> size </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
<symbol> + </symbol>
<term>
<identifier> size </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<symbol> } </symbol>
</class>
`

const Square_SquareGame_XML string = `<class>
<keyword> class </keyword>
<identifier> SquareGame </identifier>
<symbol> { </symbol>
<classVarDec>
<keyword> field </keyword>
<identifier> Square </identifier>
<identifier> square </identifier>
<symbol> ; </symbol>
</classVarDec>
<classVarDec>
<keyword> field </keyword>
<keyword> int </keyword>
<identifier> direction </identifier>
<symbol> ; </symbol>
</classVarDec>
<subroutineDec>
<keyword> constructor </keyword>
<identifier> SquareGame </identifier>
<identifier> new </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> square </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> Square </identifier>
<symbol> . </symbol>
<identifier> new </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<integerConstant> 0 </integerConstant>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<integerConstant> 0 </integerConstant>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<integerConstant> 30 </integerConstant>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<letStatement>
<keyword> let </keyword>
<identifier> direction </identifier>
<symbol> = </symbol>
<expression>
<term>
<integerConstant> 0 </integerConstant>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<returnStatement>
<keyword> return </keyword>
<expression>
<term>
<keyword> this </keyword>
</term>
</expression>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<subroutineDec>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> dispose </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> square </identifier>
<symbol> . </symbol>
<identifier> dispose </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Memory </identifier>
<symbol> . </symbol>
<identifier> deAlloc </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<keyword> this </keyword>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<subroutineDec>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> moveSquare </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<statements>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> direction </identifier>
</term>
<symbol> = </symbol>
<term>
<integerConstant> 1 </integerConstant>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> square </identifier>
<symbol> . </symbol>
<identifier> moveUp </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> direction </identifier>
</term>
<symbol> = </symbol>
<term>
<integerConstant> 2 </integerConstant>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> square </identifier>
<symbol> . </symbol>
<identifier> moveDown </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> direction </identifier>
</term>
<symbol> = </symbol>
<term>
<integerConstant> 3 </integerConstant>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> square </identifier>
<symbol> . </symbol>
<identifier> moveLeft </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> direction </identifier>
</term>
<symbol> = </symbol>
<term>
<integerConstant> 4 </integerConstant>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> square </identifier>
<symbol> . </symbol>
<identifier> moveRight </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Sys </identifier>
<symbol> . </symbol>
<identifier> wait </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<integerConstant> 5 </integerConstant>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<subroutineDec>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> run </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<varDec>
<keyword> var </keyword>
<keyword> char </keyword>
<identifier> key </identifier>
<symbol> ; </symbol>
</varDec>
<varDec>
<keyword> var </keyword>
<keyword> boolean </keyword>
<identifier> exit </identifier>
<symbol> ; </symbol>
</varDec>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> exit </identifier>
<symbol> = </symbol>
<expression>
<term>
<keyword> false </keyword>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<whileStatement>
<keyword> while </keyword>
<symbol> ( </symbol>
<expression>
<term>
<symbol> ~ </symbol>
<term>
<identifier> exit </identifier>
</term>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<whileStatement>
<keyword> while </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> key </identifier>
</term>
<symbol> = </symbol>
<term>
<integerConstant> 0 </integerConstant>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> key </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> Keyboard </identifier>
<symbol> . </symbol>
<identifier> keyPressed </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<doStatement>
<keyword> do </keyword>
<identifier> moveSquare </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
<symbol> } </symbol>
</whileStatement>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> key </identifier>
</term>
<symbol> = </symbol>
<term>
<integerConstant> 81 </integerConstant>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> exit </identifier>
<symbol> = </symbol>
<expression>
<term>
<keyword> true </keyword>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> key </identifier>
</term>
<symbol> = </symbol>
<term>
<integerConstant> 90 </integerConstant>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> square </identifier>
<symbol> . </symbol>
<identifier> decSize </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> key </identifier>
</term>
<symbol> = </symbol>
<term>
<integerConstant> 88 </integerConstant>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> square </identifier>
<symbol> . </symbol>
<identifier> incSize </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> key </identifier>
</term>
<symbol> = </symbol>
<term>
<integerConstant> 131 </integerConstant>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> direction </identifier>
<symbol> = </symbol>
<expression>
<term>
<integerConstant> 1 </integerConstant>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> key </identifier>
</term>
<symbol> = </symbol>
<term>
<integerConstant> 133 </integerConstant>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> direction </identifier>
<symbol> = </symbol>
<expression>
<term>
<integerConstant> 2 </integerConstant>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> key </identifier>
</term>
<symbol> = </symbol>
<term>
<integerConstant> 130 </integerConstant>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> direction </identifier>
<symbol> = </symbol>
<expression>
<term>
<integerConstant> 3 </integerConstant>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> key </identifier>
</term>
<symbol> = </symbol>
<term>
<integerConstant> 132 </integerConstant>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> direction </identifier>
<symbol> = </symbol>
<expression>
<term>
<integerConstant> 4 </integerConstant>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<whileStatement>
<keyword> while </keyword>
<symbol> ( </symbol>
<expression>
<term>
<symbol> ~ </symbol>
<term>
<symbol> ( </symbol>
<expression>
<term>
<identifier> key </identifier>
</term>
<symbol> = </symbol>
<term>
<integerConstant> 0 </integerConstant>
</term>
</expression>
<symbol> ) </symbol>
</term>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> key </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> Keyboard </identifier>
<symbol> . </symbol>
<identifier> keyPressed </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<doStatement>
<keyword> do </keyword>
<identifier> moveSquare </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
<symbol> } </symbol>
</whileStatement>
</statements>
<symbol> } </symbol>
</whileStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<symbol> } </symbol>
</class>
`

const SquareExpressionless_Main_XML string = `<class>
<keyword> class </keyword>
<identifier> Main </identifier>
<symbol> { </symbol>
<classVarDec>
<keyword> static </keyword>
<keyword> boolean </keyword>
<identifier> test </identifier>
<symbol> ; </symbol>
</classVarDec>
<subroutineDec>
<keyword> function </keyword>
<keyword> void </keyword>
<identifier> main </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<varDec>
<keyword> var </keyword>
<identifier> SquareGame </identifier>
<identifier> game </identifier>
<symbol> ; </symbol>
</varDec>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> game </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> game </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<doStatement>
<keyword> do </keyword>
<identifier> game </identifier>
<symbol> . </symbol>
<identifier> run </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<doStatement>
<keyword> do </keyword>
<identifier> game </identifier>
<symbol> . </symbol>
<identifier> dispose </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<subroutineDec>
<keyword> function </keyword>
<keyword> void </keyword>
<identifier> more </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<varDec>
<keyword> var </keyword>
<keyword> boolean </keyword>
<identifier> b </identifier>
<symbol> ; </symbol>
</varDec>
<statements>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> b </identifier>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
</statements>
<symbol> } </symbol>
<keyword> else </keyword>
<symbol> { </symbol>
<statements>
</statements>
<symbol> } </symbol>
</ifStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<symbol> } </symbol>
</class>
`

const SquareExpressionless_Square_XML string = `<class>
<keyword> class </keyword>
<identifier> Square </identifier>
<symbol> { </symbol>
<classVarDec>
<keyword> field </keyword>
<keyword> int </keyword>
<identifier> x </identifier>
<symbol> , </symbol>
<identifier> y </identifier>
<symbol> ; </symbol>
</classVarDec>
<classVarDec>
<keyword> field </keyword>
<keyword> int </keyword>
<identifier> size </identifier>
<symbol> ; </symbol>
</classVarDec>
<subroutineDec>
<keyword> constructor </keyword>
<identifier> Square </identifier>
<identifier> new </identifier>
<symbol> ( </symbol>
<parameterList>
<keyword> int </keyword>
<identifier> Ax </identifier>
<symbol> , </symbol>
<keyword> int </keyword>
<identifier> Ay </identifier>
<symbol> , </symbol>
<keyword> int </keyword>
<identifier> Asize </identifier>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> x </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> Ax </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<letStatement>
<keyword> let </keyword>
<identifier> y </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> Ay </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<letStatement>
<keyword> let </keyword>
<identifier> size </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> Asize </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<doStatement>
<keyword> do </keyword>
<identifier> draw </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<returnStatement>
<keyword> return </keyword>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<subroutineDec>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> dispose </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> Memory </identifier>
<symbol> . </symbol>
<identifier> deAlloc </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<keyword> this </keyword>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<subroutineDec>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> draw </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<subroutineDec>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> erase </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<subroutineDec>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> incSize </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<statements>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> erase </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<letStatement>
<keyword> let </keyword>
<identifier> size </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> size </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<doStatement>
<keyword> do </keyword>
<identifier> draw </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<subroutineDec>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> decSize </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<statements>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> size </identifier>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> erase </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<letStatement>
<keyword> let </keyword>
<identifier> size </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> size </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<doStatement>
<keyword> do </keyword>
<identifier> draw </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<subroutineDec>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> moveUp </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<statements>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<letStatement>
<keyword> let </keyword>
<identifier> y </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<subroutineDec>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> moveDown </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<statements>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<letStatement>
<keyword> let </keyword>
<identifier> y </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<subroutineDec>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> moveLeft </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<statements>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<letStatement>
<keyword> let </keyword>
<identifier> x </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<subroutineDec>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> moveRight </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<statements>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<letStatement>
<keyword> let </keyword>
<identifier> x </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> setColor </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Screen </identifier>
<symbol> . </symbol>
<identifier> drawRectangle </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<symbol> } </symbol>
</class>
`

const SquareExpressionless_SquareGame_XML string = `<class>
<keyword> class </keyword>
<identifier> SquareGame </identifier>
<symbol> { </symbol>
<classVarDec>
<keyword> field </keyword>
<identifier> Square </identifier>
<identifier> square </identifier>
<symbol> ; </symbol>
</classVarDec>
<classVarDec>
<keyword> field </keyword>
<keyword> int </keyword>
<identifier> direction </identifier>
<symbol> ; </symbol>
</classVarDec>
<subroutineDec>
<keyword> constructor </keyword>
<identifier> SquareGame </identifier>
<identifier> new </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> square </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> square </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<letStatement>
<keyword> let </keyword>
<identifier> direction </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> direction </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<returnStatement>
<keyword> return </keyword>
<expression>
<term>
<identifier> square </identifier>
</term>
</expression>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<subroutineDec>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> dispose </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> square </identifier>
<symbol> . </symbol>
<identifier> dispose </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Memory </identifier>
<symbol> . </symbol>
<identifier> deAlloc </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> square </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<subroutineDec>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> moveSquare </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<statements>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> direction </identifier>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> square </identifier>
<symbol> . </symbol>
<identifier> moveUp </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> direction </identifier>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> square </identifier>
<symbol> . </symbol>
<identifier> moveDown </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> direction </identifier>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> square </identifier>
<symbol> . </symbol>
<identifier> moveLeft </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> direction </identifier>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> square </identifier>
<symbol> . </symbol>
<identifier> moveRight </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<doStatement>
<keyword> do </keyword>
<identifier> Sys </identifier>
<symbol> . </symbol>
<identifier> wait </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<identifier> direction </identifier>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<subroutineDec>
<keyword> method </keyword>
<keyword> void </keyword>
<identifier> run </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<varDec>
<keyword> var </keyword>
<keyword> char </keyword>
<identifier> key </identifier>
<symbol> ; </symbol>
</varDec>
<varDec>
<keyword> var </keyword>
<keyword> boolean </keyword>
<identifier> exit </identifier>
<symbol> ; </symbol>
</varDec>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> exit </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> key </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<whileStatement>
<keyword> while </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> exit </identifier>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<whileStatement>
<keyword> while </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> key </identifier>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> key </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> key </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<doStatement>
<keyword> do </keyword>
<identifier> moveSquare </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
<symbol> } </symbol>
</whileStatement>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> key </identifier>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> exit </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> exit </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> key </identifier>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> square </identifier>
<symbol> . </symbol>
<identifier> decSize </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> key </identifier>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier> square </identifier>
<symbol> . </symbol>
<identifier> incSize </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> key </identifier>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> direction </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> exit </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> key </identifier>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> direction </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> key </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> key </identifier>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> direction </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> square </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> key </identifier>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> direction </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> direction </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<whileStatement>
<keyword> while </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> key </identifier>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> key </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> key </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<doStatement>
<keyword> do </keyword>
<identifier> moveSquare </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
<symbol> } </symbol>
</whileStatement>
</statements>
<symbol> } </symbol>
</whileStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<symbol> } </symbol>
</class>
`

// ----- E_XML -----

const HelloWorld_MainE_XML string = `<class>
<keyword> class </keyword>
<identifier category="TODO" action="TODO" kind="TODO" index="TODO"> Main </identifier>
<symbol> { </symbol>
<subroutineDec>
<keyword> function </keyword>
<keyword> void </keyword>
<identifier category="TODO" action="TODO" kind="TODO" index="TODO"> main </identifier>
<symbol> ( </symbol>
<parameterList>
</parameterList>
<symbol> ) </symbol>
<subroutineBody>
<symbol> { </symbol>
<statements>
<doStatement>
<keyword> do </keyword>
<identifier category="TODO" action="TODO" kind="TODO" index="TODO"> Output </identifier>
<symbol> . </symbol>
<identifier category="TODO" action="TODO" kind="TODO" index="TODO"> printString </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<stringConstant> Hello World </stringConstant>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
<symbol> } </symbol>
</subroutineBody>
</subroutineDec>
<symbol> } </symbol>
</class>
`

// ----- VM -----

const Average_Main_VM string = `function Main.main 4
push constant 18
call String.new 1
push constant 72
call String.appendChar 2
push constant 111
call String.appendChar 2
push constant 119
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 109
call String.appendChar 2
push constant 97
call String.appendChar 2
push constant 110
call String.appendChar 2
push constant 121
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 110
call String.appendChar 2
push constant 117
call String.appendChar 2
push constant 109
call String.appendChar 2
push constant 98
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 114
call String.appendChar 2
push constant 115
call String.appendChar 2
push constant 63
call String.appendChar 2
push constant 32
call String.appendChar 2
call Keyboard.readInt 1
pop local 1
push local 1
call Array.new 1
pop local 0
push constant 0
pop local 2
label WHILE_EXP0
push local 2
push local 1
lt
not
if-goto WHILE_END0
push local 2
push local 0
add
push constant 16
call String.new 1
push constant 69
call String.appendChar 2
push constant 110
call String.appendChar 2
push constant 116
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 114
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 97
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 110
call String.appendChar 2
push constant 117
call String.appendChar 2
push constant 109
call String.appendChar 2
push constant 98
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 114
call String.appendChar 2
push constant 58
call String.appendChar 2
push constant 32
call String.appendChar 2
call Keyboard.readInt 1
pop temp 0
pop pointer 1
push temp 0
pop that 0
push local 3
push local 2
push local 0
add
pop pointer 1
push that 0
add
pop local 3
push local 2
push constant 1
add
pop local 2
goto WHILE_EXP0
label WHILE_END0
push constant 15
call String.new 1
push constant 84
call String.appendChar 2
push constant 104
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 97
call String.appendChar 2
push constant 118
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 114
call String.appendChar 2
push constant 97
call String.appendChar 2
push constant 103
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 105
call String.appendChar 2
push constant 115
call String.appendChar 2
push constant 32
call String.appendChar 2
call Output.printString 1
pop temp 0
push local 3
push local 1
call Math.divide 2
call Output.printInt 1
pop temp 0
push constant 0
return
`

const ComplexArrays_Main_VM string = `function Main.main 3
push constant 10
call Array.new 1
pop local 0
push constant 5
call Array.new 1
pop local 1
push constant 1
call Array.new 1
pop local 2
push constant 3
push local 0
add
push constant 2
pop temp 0
pop pointer 1
push temp 0
pop that 0
push constant 4
push local 0
add
push constant 8
pop temp 0
pop pointer 1
push temp 0
pop that 0
push constant 5
push local 0
add
push constant 4
pop temp 0
pop pointer 1
push temp 0
pop that 0
push constant 3
push local 0
add
pop pointer 1
push that 0
push local 1
add
push constant 3
push local 0
add
pop pointer 1
push that 0
push constant 3
add
pop temp 0
pop pointer 1
push temp 0
pop that 0
push constant 3
push local 0
add
pop pointer 1
push that 0
push local 1
add
pop pointer 1
push that 0
push local 0
add
push constant 5
push local 0
add
pop pointer 1
push that 0
push local 0
add
pop pointer 1
push that 0
push constant 7
push constant 3
push local 0
add
pop pointer 1
push that 0
sub
push constant 2
call Main.double 1
sub
push constant 1
add
push local 1
add
pop pointer 1
push that 0
call Math.multiply 2
pop temp 0
pop pointer 1
push temp 0
pop that 0
push constant 0
push local 2
add
push constant 0
pop temp 0
pop pointer 1
push temp 0
pop that 0
push constant 0
push local 2
add
pop pointer 1
push that 0
pop local 2
push constant 43
call String.new 1
push constant 84
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 115
call String.appendChar 2
push constant 116
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 49
call String.appendChar 2
push constant 58
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 120
call String.appendChar 2
push constant 112
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 99
call String.appendChar 2
push constant 116
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 100
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 114
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 115
call String.appendChar 2
push constant 117
call String.appendChar 2
push constant 108
call String.appendChar 2
push constant 116
call String.appendChar 2
push constant 58
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 53
call String.appendChar 2
push constant 59
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 97
call String.appendChar 2
push constant 99
call String.appendChar 2
push constant 116
call String.appendChar 2
push constant 117
call String.appendChar 2
push constant 97
call String.appendChar 2
push constant 108
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 114
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 115
call String.appendChar 2
push constant 117
call String.appendChar 2
push constant 108
call String.appendChar 2
push constant 116
call String.appendChar 2
push constant 58
call String.appendChar 2
push constant 32
call String.appendChar 2
call Output.printString 1
pop temp 0
push constant 2
push local 1
add
pop pointer 1
push that 0
call Output.printInt 1
pop temp 0
call Output.println 0
pop temp 0
push constant 44
call String.new 1
push constant 84
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 115
call String.appendChar 2
push constant 116
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 50
call String.appendChar 2
push constant 58
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 120
call String.appendChar 2
push constant 112
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 99
call String.appendChar 2
push constant 116
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 100
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 114
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 115
call String.appendChar 2
push constant 117
call String.appendChar 2
push constant 108
call String.appendChar 2
push constant 116
call String.appendChar 2
push constant 58
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 52
call String.appendChar 2
push constant 48
call String.appendChar 2
push constant 59
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 97
call String.appendChar 2
push constant 99
call String.appendChar 2
push constant 116
call String.appendChar 2
push constant 117
call String.appendChar 2
push constant 97
call String.appendChar 2
push constant 108
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 114
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 115
call String.appendChar 2
push constant 117
call String.appendChar 2
push constant 108
call String.appendChar 2
push constant 116
call String.appendChar 2
push constant 58
call String.appendChar 2
push constant 32
call String.appendChar 2
call Output.printString 1
pop temp 0
push constant 5
push local 0
add
pop pointer 1
push that 0
call Output.printInt 1
pop temp 0
call Output.println 0
pop temp 0
push constant 43
call String.new 1
push constant 84
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 115
call String.appendChar 2
push constant 116
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 51
call String.appendChar 2
push constant 58
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 120
call String.appendChar 2
push constant 112
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 99
call String.appendChar 2
push constant 116
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 100
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 114
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 115
call String.appendChar 2
push constant 117
call String.appendChar 2
push constant 108
call String.appendChar 2
push constant 116
call String.appendChar 2
push constant 58
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 48
call String.appendChar 2
push constant 59
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 97
call String.appendChar 2
push constant 99
call String.appendChar 2
push constant 116
call String.appendChar 2
push constant 117
call String.appendChar 2
push constant 97
call String.appendChar 2
push constant 108
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 114
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 115
call String.appendChar 2
push constant 117
call String.appendChar 2
push constant 108
call String.appendChar 2
push constant 116
call String.appendChar 2
push constant 58
call String.appendChar 2
push constant 32
call String.appendChar 2
call Output.printString 1
pop temp 0
push local 2
call Output.printInt 1
pop temp 0
call Output.println 0
pop temp 0
push constant 0
pop local 2
push local 2
push constant 0
eq
if-goto IF_TRUE0
goto IF_FALSE0
label IF_TRUE0
push local 0
push constant 10
call Main.fill 2
pop temp 0
push constant 3
push local 0
add
pop pointer 1
push that 0
pop local 2
push constant 1
push local 2
add
push constant 33
pop temp 0
pop pointer 1
push temp 0
pop that 0
push constant 7
push local 0
add
pop pointer 1
push that 0
pop local 2
push constant 1
push local 2
add
push constant 77
pop temp 0
pop pointer 1
push temp 0
pop that 0
push constant 3
push local 0
add
pop pointer 1
push that 0
pop local 1
push constant 1
push local 1
add
push constant 1
push local 1
add
pop pointer 1
push that 0
push constant 1
push local 2
add
pop pointer 1
push that 0
add
pop temp 0
pop pointer 1
push temp 0
pop that 0
label IF_FALSE0
push constant 44
call String.new 1
push constant 84
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 115
call String.appendChar 2
push constant 116
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 52
call String.appendChar 2
push constant 58
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 120
call String.appendChar 2
push constant 112
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 99
call String.appendChar 2
push constant 116
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 100
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 114
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 115
call String.appendChar 2
push constant 117
call String.appendChar 2
push constant 108
call String.appendChar 2
push constant 116
call String.appendChar 2
push constant 58
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 55
call String.appendChar 2
push constant 55
call String.appendChar 2
push constant 59
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 97
call String.appendChar 2
push constant 99
call String.appendChar 2
push constant 116
call String.appendChar 2
push constant 117
call String.appendChar 2
push constant 97
call String.appendChar 2
push constant 108
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 114
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 115
call String.appendChar 2
push constant 117
call String.appendChar 2
push constant 108
call String.appendChar 2
push constant 116
call String.appendChar 2
push constant 58
call String.appendChar 2
push constant 32
call String.appendChar 2
call Output.printString 1
pop temp 0
push constant 1
push local 2
add
pop pointer 1
push that 0
call Output.printInt 1
pop temp 0
call Output.println 0
pop temp 0
push constant 45
call String.new 1
push constant 84
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 115
call String.appendChar 2
push constant 116
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 53
call String.appendChar 2
push constant 58
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 120
call String.appendChar 2
push constant 112
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 99
call String.appendChar 2
push constant 116
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 100
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 114
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 115
call String.appendChar 2
push constant 117
call String.appendChar 2
push constant 108
call String.appendChar 2
push constant 116
call String.appendChar 2
push constant 58
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 49
call String.appendChar 2
push constant 49
call String.appendChar 2
push constant 48
call String.appendChar 2
push constant 59
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 97
call String.appendChar 2
push constant 99
call String.appendChar 2
push constant 116
call String.appendChar 2
push constant 117
call String.appendChar 2
push constant 97
call String.appendChar 2
push constant 108
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 114
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 115
call String.appendChar 2
push constant 117
call String.appendChar 2
push constant 108
call String.appendChar 2
push constant 116
call String.appendChar 2
push constant 58
call String.appendChar 2
push constant 32
call String.appendChar 2
call Output.printString 1
pop temp 0
push constant 1
push local 1
add
pop pointer 1
push that 0
call Output.printInt 1
pop temp 0
call Output.println 0
pop temp 0
push constant 0
return
function Main.double 0
push argument 0
push constant 2
call Math.multiply 2
return
function Main.fill 0
label WHILE_EXP0
push argument 1
push constant 0
gt
not
if-goto WHILE_END0
push argument 1
push constant 1
sub
pop argument 1
push argument 1
push argument 0
add
push constant 3
call Array.new 1
pop temp 0
pop pointer 1
push temp 0
pop that 0
goto WHILE_EXP0
label WHILE_END0
push constant 0
return
`

const ConvertToBin_Main_VM string = `function Main.main 1
push constant 8001
push constant 16
push constant 1
neg
call Main.fillMemory 3
pop temp 0
push constant 8000
call Memory.peek 1
pop local 0
push local 0
call Main.convert 1
pop temp 0
push constant 0
return
function Main.convert 3
push constant 0
not
pop local 2
label WHILE_EXP0
push local 2
not
if-goto WHILE_END0
push local 1
push constant 1
add
pop local 1
push local 0
call Main.nextMask 1
pop local 0
push local 1
push constant 16
gt
not
if-goto IF_TRUE0
goto IF_FALSE0
label IF_TRUE0
push argument 0
push local 0
and
push constant 0
eq
not
if-goto IF_TRUE1
goto IF_FALSE1
label IF_TRUE1
push constant 8000
push local 1
add
push constant 1
call Memory.poke 2
pop temp 0
goto IF_END1
label IF_FALSE1
push constant 8000
push local 1
add
push constant 0
call Memory.poke 2
pop temp 0
label IF_END1
goto IF_END0
label IF_FALSE0
push constant 0
pop local 2
label IF_END0
goto WHILE_EXP0
label WHILE_END0
push constant 0
return
function Main.nextMask 0
push argument 0
push constant 0
eq
if-goto IF_TRUE0
goto IF_FALSE0
label IF_TRUE0
push constant 1
return
goto IF_END0
label IF_FALSE0
push argument 0
push constant 2
call Math.multiply 2
return
label IF_END0
function Main.fillMemory 0
label WHILE_EXP0
push argument 1
push constant 0
gt
not
if-goto WHILE_END0
push argument 0
push argument 2
call Memory.poke 2
pop temp 0
push argument 1
push constant 1
sub
pop argument 1
push argument 0
push constant 1
add
pop argument 0
goto WHILE_EXP0
label WHILE_END0
push constant 0
return
`

const MultiConditionIf_Main_VM string = `function Main.main 1
call Keyboard.keyPressed 0
pop local 0
push local 0
call Direction.U 0
eq
push local 0
call Direction.D 0
eq
or
push local 0
call Direction.L 0
eq
or
push local 0
call Direction.R 0
eq
or
if-goto IF_TRUE0
goto IF_FALSE0
label IF_TRUE0
label IF_FALSE0
push constant 0
return
`

const Pong_Ball_VM string = `function Ball.new 0
push constant 15
call Memory.alloc 1
pop pointer 0
push argument 0
pop this 0
push argument 1
pop this 1
push argument 2
pop this 10
push argument 3
push constant 6
sub
pop this 11
push argument 4
pop this 12
push argument 5
push constant 6
sub
pop this 13
push constant 0
pop this 14
push pointer 0
call Ball.show 1
pop temp 0
push pointer 0
return
function Ball.dispose 0
push argument 0
pop pointer 0
push pointer 0
call Memory.deAlloc 1
pop temp 0
push constant 0
return
function Ball.show 0
push argument 0
pop pointer 0
push constant 0
not
call Screen.setColor 1
pop temp 0
push pointer 0
call Ball.draw 1
pop temp 0
push constant 0
return
function Ball.hide 0
push argument 0
pop pointer 0
push constant 0
call Screen.setColor 1
pop temp 0
push pointer 0
call Ball.draw 1
pop temp 0
push constant 0
return
function Ball.draw 0
push argument 0
pop pointer 0
push this 0
push this 1
push this 0
push constant 5
add
push this 1
push constant 5
add
call Screen.drawRectangle 4
pop temp 0
push constant 0
return
function Ball.getLeft 0
push argument 0
pop pointer 0
push this 0
return
function Ball.getRight 0
push argument 0
pop pointer 0
push this 0
push constant 5
add
return
function Ball.setDestination 3
push argument 0
pop pointer 0
push argument 1
push this 0
sub
pop this 2
push argument 2
push this 1
sub
pop this 3
push this 2
call Math.abs 1
pop local 0
push this 3
call Math.abs 1
pop local 1
push local 0
push local 1
lt
pop this 7
push this 7
if-goto IF_TRUE0
goto IF_FALSE0
label IF_TRUE0
push local 0
pop local 2
push local 1
pop local 0
push local 2
pop local 1
push this 1
push argument 2
lt
pop this 8
push this 0
push argument 1
lt
pop this 9
goto IF_END0
label IF_FALSE0
push this 0
push argument 1
lt
pop this 8
push this 1
push argument 2
lt
pop this 9
label IF_END0
push constant 2
push local 1
call Math.multiply 2
push local 0
sub
pop this 4
push constant 2
push local 1
call Math.multiply 2
pop this 5
push constant 2
push local 1
push local 0
sub
call Math.multiply 2
pop this 6
push constant 0
return
function Ball.move 0
push argument 0
pop pointer 0
push pointer 0
call Ball.hide 1
pop temp 0
push this 4
push constant 0
lt
if-goto IF_TRUE0
goto IF_FALSE0
label IF_TRUE0
push this 4
push this 5
add
pop this 4
goto IF_END0
label IF_FALSE0
push this 4
push this 6
add
pop this 4
push this 9
if-goto IF_TRUE1
goto IF_FALSE1
label IF_TRUE1
push this 7
if-goto IF_TRUE2
goto IF_FALSE2
label IF_TRUE2
push this 0
push constant 4
add
pop this 0
goto IF_END2
label IF_FALSE2
push this 1
push constant 4
add
pop this 1
label IF_END2
goto IF_END1
label IF_FALSE1
push this 7
if-goto IF_TRUE3
goto IF_FALSE3
label IF_TRUE3
push this 0
push constant 4
sub
pop this 0
goto IF_END3
label IF_FALSE3
push this 1
push constant 4
sub
pop this 1
label IF_END3
label IF_END1
label IF_END0
push this 8
if-goto IF_TRUE4
goto IF_FALSE4
label IF_TRUE4
push this 7
if-goto IF_TRUE5
goto IF_FALSE5
label IF_TRUE5
push this 1
push constant 4
add
pop this 1
goto IF_END5
label IF_FALSE5
push this 0
push constant 4
add
pop this 0
label IF_END5
goto IF_END4
label IF_FALSE4
push this 7
if-goto IF_TRUE6
goto IF_FALSE6
label IF_TRUE6
push this 1
push constant 4
sub
pop this 1
goto IF_END6
label IF_FALSE6
push this 0
push constant 4
sub
pop this 0
label IF_END6
label IF_END4
push this 0
push this 10
gt
not
if-goto IF_TRUE7
goto IF_FALSE7
label IF_TRUE7
push constant 1
pop this 14
push this 10
pop this 0
label IF_FALSE7
push this 0
push this 11
lt
not
if-goto IF_TRUE8
goto IF_FALSE8
label IF_TRUE8
push constant 2
pop this 14
push this 11
pop this 0
label IF_FALSE8
push this 1
push this 12
gt
not
if-goto IF_TRUE9
goto IF_FALSE9
label IF_TRUE9
push constant 3
pop this 14
push this 12
pop this 1
label IF_FALSE9
push this 1
push this 13
lt
not
if-goto IF_TRUE10
goto IF_FALSE10
label IF_TRUE10
push constant 4
pop this 14
push this 13
pop this 1
label IF_FALSE10
push pointer 0
call Ball.show 1
pop temp 0
push this 14
return
function Ball.bounce 5
push argument 0
pop pointer 0
push this 2
push constant 10
call Math.divide 2
pop local 2
push this 3
push constant 10
call Math.divide 2
pop local 3
push argument 1
push constant 0
eq
if-goto IF_TRUE0
goto IF_FALSE0
label IF_TRUE0
push constant 10
pop local 4
goto IF_END0
label IF_FALSE0
push this 2
push constant 0
lt
not
push argument 1
push constant 1
eq
and
push this 2
push constant 0
lt
push argument 1
push constant 1
neg
eq
and
or
if-goto IF_TRUE1
goto IF_FALSE1
label IF_TRUE1
push constant 20
pop local 4
goto IF_END1
label IF_FALSE1
push constant 5
pop local 4
label IF_END1
label IF_END0
push this 14
push constant 1
eq
if-goto IF_TRUE2
goto IF_FALSE2
label IF_TRUE2
push constant 506
pop local 0
push local 3
push constant 50
neg
call Math.multiply 2
push local 2
call Math.divide 2
pop local 1
push this 1
push local 1
push local 4
call Math.multiply 2
add
pop local 1
goto IF_END2
label IF_FALSE2
push this 14
push constant 2
eq
if-goto IF_TRUE3
goto IF_FALSE3
label IF_TRUE3
push constant 0
pop local 0
push local 3
push constant 50
call Math.multiply 2
push local 2
call Math.divide 2
pop local 1
push this 1
push local 1
push local 4
call Math.multiply 2
add
pop local 1
goto IF_END3
label IF_FALSE3
push this 14
push constant 3
eq
if-goto IF_TRUE4
goto IF_FALSE4
label IF_TRUE4
push constant 250
pop local 1
push local 2
push constant 25
neg
call Math.multiply 2
push local 3
call Math.divide 2
pop local 0
push this 0
push local 0
push local 4
call Math.multiply 2
add
pop local 0
goto IF_END4
label IF_FALSE4
push constant 0
pop local 1
push local 2
push constant 25
call Math.multiply 2
push local 3
call Math.divide 2
pop local 0
push this 0
push local 0
push local 4
call Math.multiply 2
add
pop local 0
label IF_END4
label IF_END3
label IF_END2
push pointer 0
push local 0
push local 1
call Ball.setDestination 3
pop temp 0
push constant 0
return
`

const Pong_Bat_VM string = `function Bat.new 0
push constant 5
call Memory.alloc 1
pop pointer 0
push argument 0
pop this 0
push argument 1
pop this 1
push argument 2
pop this 2
push argument 3
pop this 3
push constant 2
pop this 4
push pointer 0
call Bat.show 1
pop temp 0
push pointer 0
return
function Bat.dispose 0
push argument 0
pop pointer 0
push pointer 0
call Memory.deAlloc 1
pop temp 0
push constant 0
return
function Bat.show 0
push argument 0
pop pointer 0
push constant 0
not
call Screen.setColor 1
pop temp 0
push pointer 0
call Bat.draw 1
pop temp 0
push constant 0
return
function Bat.hide 0
push argument 0
pop pointer 0
push constant 0
call Screen.setColor 1
pop temp 0
push pointer 0
call Bat.draw 1
pop temp 0
push constant 0
return
function Bat.draw 0
push argument 0
pop pointer 0
push this 0
push this 1
push this 0
push this 2
add
push this 1
push this 3
add
call Screen.drawRectangle 4
pop temp 0
push constant 0
return
function Bat.setDirection 0
push argument 0
pop pointer 0
push argument 1
pop this 4
push constant 0
return
function Bat.getLeft 0
push argument 0
pop pointer 0
push this 0
return
function Bat.getRight 0
push argument 0
pop pointer 0
push this 0
push this 2
add
return
function Bat.setWidth 0
push argument 0
pop pointer 0
push pointer 0
call Bat.hide 1
pop temp 0
push argument 1
pop this 2
push pointer 0
call Bat.show 1
pop temp 0
push constant 0
return
function Bat.move 0
push argument 0
pop pointer 0
push this 4
push constant 1
eq
if-goto IF_TRUE0
goto IF_FALSE0
label IF_TRUE0
push this 0
push constant 4
sub
pop this 0
push this 0
push constant 0
lt
if-goto IF_TRUE1
goto IF_FALSE1
label IF_TRUE1
push constant 0
pop this 0
label IF_FALSE1
push constant 0
call Screen.setColor 1
pop temp 0
push this 0
push this 2
add
push constant 1
add
push this 1
push this 0
push this 2
add
push constant 4
add
push this 1
push this 3
add
call Screen.drawRectangle 4
pop temp 0
push constant 0
not
call Screen.setColor 1
pop temp 0
push this 0
push this 1
push this 0
push constant 3
add
push this 1
push this 3
add
call Screen.drawRectangle 4
pop temp 0
goto IF_END0
label IF_FALSE0
push this 0
push constant 4
add
pop this 0
push this 0
push this 2
add
push constant 511
gt
if-goto IF_TRUE2
goto IF_FALSE2
label IF_TRUE2
push constant 511
push this 2
sub
pop this 0
label IF_FALSE2
push constant 0
call Screen.setColor 1
pop temp 0
push this 0
push constant 4
sub
push this 1
push this 0
push constant 1
sub
push this 1
push this 3
add
call Screen.drawRectangle 4
pop temp 0
push constant 0
not
call Screen.setColor 1
pop temp 0
push this 0
push this 2
add
push constant 3
sub
push this 1
push this 0
push this 2
add
push this 1
push this 3
add
call Screen.drawRectangle 4
pop temp 0
label IF_END0
push constant 0
return
`

const Pong_Main_VM string = `function Main.main 1
call PongGame.newInstance 0
pop temp 0
call PongGame.getInstance 0
pop local 0
push local 0
call PongGame.run 1
pop temp 0
push local 0
call PongGame.dispose 1
pop temp 0
push constant 0
return
`

const Pong_PongGame_VM string = `function PongGame.new 0
push constant 7
call Memory.alloc 1
pop pointer 0
call Screen.clearScreen 0
pop temp 0
push constant 50
pop this 6
push constant 230
push constant 229
push this 6
push constant 7
call Bat.new 4
pop this 0
push constant 253
push constant 222
push constant 0
push constant 511
push constant 0
push constant 229
call Ball.new 6
pop this 1
push this 1
push constant 400
push constant 0
call Ball.setDestination 3
pop temp 0
push constant 0
push constant 238
push constant 511
push constant 240
call Screen.drawRectangle 4
pop temp 0
push constant 22
push constant 0
call Output.moveCursor 2
pop temp 0
push constant 8
call String.new 1
push constant 83
call String.appendChar 2
push constant 99
call String.appendChar 2
push constant 111
call String.appendChar 2
push constant 114
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 58
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 48
call String.appendChar 2
call Output.printString 1
pop temp 0
push constant 0
pop this 3
push constant 0
pop this 4
push constant 0
pop this 2
push constant 0
pop this 5
push pointer 0
return
function PongGame.dispose 0
push argument 0
pop pointer 0
push this 0
call Bat.dispose 1
pop temp 0
push this 1
call Ball.dispose 1
pop temp 0
push pointer 0
call Memory.deAlloc 1
pop temp 0
push constant 0
return
function PongGame.newInstance 0
call PongGame.new 0
pop static 0
push constant 0
return
function PongGame.getInstance 0
push static 0
return
function PongGame.run 1
push argument 0
pop pointer 0
label WHILE_EXP0
push this 3
not
not
if-goto WHILE_END0
label WHILE_EXP1
push local 0
push constant 0
eq
push this 3
not
and
not
if-goto WHILE_END1
call Keyboard.keyPressed 0
pop local 0
push this 0
call Bat.move 1
pop temp 0
push pointer 0
call PongGame.moveBall 1
pop temp 0
push constant 50
call Sys.wait 1
pop temp 0
goto WHILE_EXP1
label WHILE_END1
push local 0
push constant 130
eq
if-goto IF_TRUE0
goto IF_FALSE0
label IF_TRUE0
push this 0
push constant 1
call Bat.setDirection 2
pop temp 0
goto IF_END0
label IF_FALSE0
push local 0
push constant 132
eq
if-goto IF_TRUE1
goto IF_FALSE1
label IF_TRUE1
push this 0
push constant 2
call Bat.setDirection 2
pop temp 0
goto IF_END1
label IF_FALSE1
push local 0
push constant 140
eq
if-goto IF_TRUE2
goto IF_FALSE2
label IF_TRUE2
push constant 0
not
pop this 3
label IF_FALSE2
label IF_END1
label IF_END0
label WHILE_EXP2
push local 0
push constant 0
eq
not
push this 3
not
and
not
if-goto WHILE_END2
call Keyboard.keyPressed 0
pop local 0
push this 0
call Bat.move 1
pop temp 0
push pointer 0
call PongGame.moveBall 1
pop temp 0
push constant 50
call Sys.wait 1
pop temp 0
goto WHILE_EXP2
label WHILE_END2
goto WHILE_EXP0
label WHILE_END0
push this 3
if-goto IF_TRUE3
goto IF_FALSE3
label IF_TRUE3
push constant 10
push constant 27
call Output.moveCursor 2
pop temp 0
push constant 9
call String.new 1
push constant 71
call String.appendChar 2
push constant 97
call String.appendChar 2
push constant 109
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 79
call String.appendChar 2
push constant 118
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 114
call String.appendChar 2
call Output.printString 1
pop temp 0
label IF_FALSE3
push constant 0
return
function PongGame.moveBall 5
push argument 0
pop pointer 0
push this 1
call Ball.move 1
pop this 2
push this 2
push constant 0
gt
push this 2
push this 5
eq
not
and
if-goto IF_TRUE0
goto IF_FALSE0
label IF_TRUE0
push this 2
pop this 5
push constant 0
pop local 0
push this 0
call Bat.getLeft 1
pop local 1
push this 0
call Bat.getRight 1
pop local 2
push this 1
call Ball.getLeft 1
pop local 3
push this 1
call Ball.getRight 1
pop local 4
push this 2
push constant 4
eq
if-goto IF_TRUE1
goto IF_FALSE1
label IF_TRUE1
push local 1
push local 4
gt
push local 2
push local 3
lt
or
pop this 3
push this 3
not
if-goto IF_TRUE2
goto IF_FALSE2
label IF_TRUE2
push local 4
push local 1
push constant 10
add
lt
if-goto IF_TRUE3
goto IF_FALSE3
label IF_TRUE3
push constant 1
neg
pop local 0
goto IF_END3
label IF_FALSE3
push local 3
push local 2
push constant 10
sub
gt
if-goto IF_TRUE4
goto IF_FALSE4
label IF_TRUE4
push constant 1
pop local 0
label IF_FALSE4
label IF_END3
push this 6
push constant 2
sub
pop this 6
push this 0
push this 6
call Bat.setWidth 2
pop temp 0
push this 4
push constant 1
add
pop this 4
push constant 22
push constant 7
call Output.moveCursor 2
pop temp 0
push this 4
call Output.printInt 1
pop temp 0
label IF_FALSE2
label IF_FALSE1
push this 1
push local 0
call Ball.bounce 2
pop temp 0
label IF_FALSE0
push constant 0
return
`

const Seven_Main_VM string = `function Main.main 0
push constant 1
push constant 2
push constant 3
call Math.multiply 2
add
call Output.printInt 1
pop temp 0
push constant 0
return
`

const SimpleIf_Main_VM string = `function Main.main 3
push constant 0
pop local 0
push constant 0
pop local 1
push constant 0
pop local 2
push local 0
push constant 0
eq
if-goto IF_TRUE0
goto IF_FALSE0
label IF_TRUE0
push local 1
push constant 0
eq
if-goto IF_TRUE1
goto IF_FALSE1
label IF_TRUE1
push constant 0
pop local 2
goto IF_END1
label IF_FALSE1
push constant 1
pop local 2
label IF_END1
goto IF_END0
label IF_FALSE0
push constant 2
pop local 2
label IF_END0
push local 1
push constant 0
eq
if-goto IF_TRUE2
goto IF_FALSE2
label IF_TRUE2
push constant 3
pop local 2
goto IF_END2
label IF_FALSE2
push constant 4
pop local 2
label IF_END2
push local 2
push constant 0
eq
if-goto IF_TRUE3
goto IF_FALSE3
label IF_TRUE3
push constant 5
pop local 2
label IF_FALSE3
push constant 0
return
`

const SimpleWhile_Main_VM string = `function Main.main 3
push constant 0
pop local 0
push constant 0
pop local 1
push constant 0
pop local 2
label WHILE_EXP0
push local 0
push constant 3
lt
not
if-goto WHILE_END0
label WHILE_EXP1
push local 1
push constant 3
lt
not
if-goto WHILE_END1
push local 1
push constant 1
add
pop local 1
goto WHILE_EXP1
label WHILE_END1
push local 0
push constant 1
add
pop local 0
push constant 0
pop local 1
goto WHILE_EXP0
label WHILE_END0
label WHILE_EXP2
push local 2
push constant 3
lt
not
if-goto WHILE_END2
push local 2
push constant 1
add
pop local 2
goto WHILE_EXP2
label WHILE_END2
push constant 0
return
`

const Sokoban_Board_VM string = `function Board.new 3
push constant 4
call Memory.alloc 1
pop pointer 0
push argument 0
pop this 0
push argument 1
pop this 1
push this 0
push this 1
call Math.multiply 2
call Array.new 1
pop this 2
push constant 0
pop local 0
push constant 0
pop local 1
label WHILE_EXP0
push local 1
push this 1
lt
not
if-goto WHILE_END0
label WHILE_EXP1
push local 0
push this 0
lt
not
if-goto WHILE_END1
push argument 2
push local 1
push this 0
call Math.multiply 2
push local 0
add
call String.charAt 2
pop local 2
push local 2
push constant 64
eq
if-goto IF_TRUE0
goto IF_FALSE0
label IF_TRUE0
push local 1
push this 0
call Math.multiply 2
push local 0
add
push this 2
add
call CellType.CellTypeNone 0
push constant 0
call Cell.new 2
pop temp 0
pop pointer 1
push temp 0
pop that 0
push local 0
push local 1
call Player.new 2
pop this 3
label IF_FALSE0
push local 2
push constant 36
eq
if-goto IF_TRUE1
goto IF_FALSE1
label IF_TRUE1
push local 1
push this 0
call Math.multiply 2
push local 0
add
push this 2
add
call CellType.CellTypeNone 0
push constant 0
not
call Cell.new 2
pop temp 0
pop pointer 1
push temp 0
pop that 0
label IF_FALSE1
push local 2
push constant 46
eq
if-goto IF_TRUE2
goto IF_FALSE2
label IF_TRUE2
push local 1
push this 0
call Math.multiply 2
push local 0
add
push this 2
add
call CellType.CellTypeGoal 0
push constant 0
call Cell.new 2
pop temp 0
pop pointer 1
push temp 0
pop that 0
label IF_FALSE2
push local 2
push constant 35
eq
if-goto IF_TRUE3
goto IF_FALSE3
label IF_TRUE3
push local 1
push this 0
call Math.multiply 2
push local 0
add
push this 2
add
call CellType.CellTypeWall 0
push constant 0
call Cell.new 2
pop temp 0
pop pointer 1
push temp 0
pop that 0
label IF_FALSE3
push local 2
push constant 43
eq
if-goto IF_TRUE4
goto IF_FALSE4
label IF_TRUE4
push local 1
push this 0
call Math.multiply 2
push local 0
add
push this 2
add
call CellType.CellTypeGoal 0
push constant 0
call Cell.new 2
pop temp 0
pop pointer 1
push temp 0
pop that 0
push local 0
push local 1
call Player.new 2
pop this 3
label IF_FALSE4
push local 2
push constant 42
eq
if-goto IF_TRUE5
goto IF_FALSE5
label IF_TRUE5
push local 1
push this 0
call Math.multiply 2
push local 0
add
push this 2
add
call CellType.CellTypeGoal 0
push constant 0
not
call Cell.new 2
pop temp 0
pop pointer 1
push temp 0
pop that 0
label IF_FALSE5
push local 2
push constant 32
eq
if-goto IF_TRUE6
goto IF_FALSE6
label IF_TRUE6
push local 1
push this 0
call Math.multiply 2
push local 0
add
push this 2
add
call CellType.CellTypeNone 0
push constant 0
call Cell.new 2
pop temp 0
pop pointer 1
push temp 0
pop that 0
label IF_FALSE6
push local 0
push constant 1
add
pop local 0
goto WHILE_EXP1
label WHILE_END1
push local 1
push constant 1
add
pop local 1
push constant 0
pop local 0
goto WHILE_EXP0
label WHILE_END0
push pointer 0
return
function Board.getCell 0
push argument 0
pop pointer 0
push argument 2
push this 0
call Math.multiply 2
push argument 1
add
push this 2
add
pop pointer 1
push that 0
return
function Board.isComplete 3
push argument 0
pop pointer 0
push constant 0
pop local 0
push this 0
push this 1
call Math.multiply 2
pop local 1
label WHILE_EXP0
push local 0
push local 1
lt
not
if-goto WHILE_END0
push local 0
push this 2
add
pop pointer 1
push that 0
pop local 2
push local 2
call Cell.getTypeOf 1
call CellType.CellTypeGoal 0
eq
push local 2
call Cell.getHasBox 1
not
and
if-goto IF_TRUE0
goto IF_FALSE0
label IF_TRUE0
push constant 0
return
label IF_FALSE0
push local 0
push constant 1
add
pop local 0
goto WHILE_EXP0
label WHILE_END0
push constant 0
not
return
function Board.getWidth 0
push argument 0
pop pointer 0
push this 0
return
function Board.getHeight 0
push argument 0
pop pointer 0
push this 1
return
function Board.getPlayer 0
push argument 0
pop pointer 0
push this 3
return
function Board.dispose 3
push argument 0
pop pointer 0
push constant 0
pop local 0
push this 0
push this 1
call Math.multiply 2
pop local 1
label WHILE_EXP0
push local 0
push local 1
lt
not
if-goto WHILE_END0
push local 0
push this 2
add
pop pointer 1
push that 0
pop local 2
push local 2
call Cell.dispose 1
pop temp 0
push local 0
push constant 1
add
pop local 0
goto WHILE_EXP0
label WHILE_END0
push this 2
call Array.dispose 1
pop temp 0
push this 3
call Player.dispose 1
pop temp 0
push pointer 0
call Memory.deAlloc 1
pop temp 0
push constant 0
return
`

const Sokoban_Cell_VM string = `function Cell.new 0
push constant 2
call Memory.alloc 1
pop pointer 0
push argument 0
pop this 0
push argument 1
pop this 1
push pointer 0
return
function Cell.getTypeOf 0
push argument 0
pop pointer 0
push this 0
return
function Cell.getHasBox 0
push argument 0
pop pointer 0
push this 1
return
function Cell.setHasBox 0
push argument 0
pop pointer 0
push argument 1
pop this 1
push constant 0
return
function Cell.dispose 0
push argument 0
pop pointer 0
push pointer 0
call Memory.deAlloc 1
pop temp 0
push constant 0
return
`

const Sokoban_CellType_VM string = `function CellType.init 0
push constant 0
pop static 0
push constant 1
pop static 1
push constant 2
pop static 2
push constant 0
return
function CellType.CellTypeNone 0
push static 0
return
function CellType.CellTypeGoal 0
push static 1
return
function CellType.CellTypeWall 0
push static 2
return
`

const Sokoban_Controller_VM string = `function Controller.new 0
push constant 3
call Memory.alloc 1
pop pointer 0
push argument 0
pop this 0
push constant 0
not
pop this 1
push constant 0
pop this 2
push pointer 0
return
function Controller.startNewGame 1
push argument 0
pop pointer 0
push this 0
call Model.getLevelManager 1
pop local 0
push local 0
call LevelManager.reset 1
pop temp 0
push pointer 0
call Controller.tryStartNextLevel 1
pop temp 0
push constant 0
return
function Controller.handleInput 1
push argument 0
pop pointer 0
call Keyboard.keyPressed 0
pop local 0
push local 0
push constant 0
eq
not
if-goto IF_TRUE0
goto IF_FALSE0
label IF_TRUE0
push local 0
call Direction.U 0
eq
push local 0
call Direction.D 0
eq
or
push local 0
call Direction.L 0
eq
or
push local 0
call Direction.R 0
eq
or
push local 0
push constant 32
eq
or
push local 0
push constant 82
eq
or
push local 0
push constant 83
eq
or
push local 0
push constant 140
eq
or
not
if-goto IF_TRUE1
goto IF_FALSE1
label IF_TRUE1
push constant 0
pop this 2
push constant 0
return
label IF_FALSE1
push local 0
push constant 140
eq
if-goto IF_TRUE2
goto IF_FALSE2
label IF_TRUE2
push constant 0
pop this 1
label IF_FALSE2
push this 2
push local 0
eq
not
if-goto IF_TRUE3
goto IF_FALSE3
label IF_TRUE3
push this 0
call Model.getState 1
call State.StateGameComplete 0
eq
if-goto IF_TRUE4
goto IF_FALSE4
label IF_TRUE4
push local 0
push constant 32
eq
if-goto IF_TRUE5
goto IF_FALSE5
label IF_TRUE5
push pointer 0
call Controller.startNewGame 1
pop temp 0
label IF_FALSE5
label IF_FALSE4
push this 0
call Model.getState 1
call State.StateLevelComplete 0
eq
if-goto IF_TRUE6
goto IF_FALSE6
label IF_TRUE6
push local 0
push constant 32
eq
if-goto IF_TRUE7
goto IF_FALSE7
label IF_TRUE7
push pointer 0
call Controller.tryStartNextLevel 1
pop temp 0
label IF_FALSE7
label IF_FALSE6
push this 0
call Model.getState 1
call State.StatePlaying 0
eq
if-goto IF_TRUE8
goto IF_FALSE8
label IF_TRUE8
push local 0
call Direction.U 0
eq
push local 0
call Direction.D 0
eq
or
push local 0
call Direction.L 0
eq
or
push local 0
call Direction.R 0
eq
or
if-goto IF_TRUE9
goto IF_FALSE9
label IF_TRUE9
push pointer 0
push local 0
call Controller.tryMovePlayer 2
pop temp 0
label IF_FALSE9
push local 0
push constant 82
eq
if-goto IF_TRUE10
goto IF_FALSE10
label IF_TRUE10
push pointer 0
call Controller.restartLevel 1
pop temp 0
label IF_FALSE10
push local 0
push constant 83
eq
if-goto IF_TRUE11
goto IF_FALSE11
label IF_TRUE11
label IF_FALSE11
label IF_FALSE8
label IF_FALSE3
push local 0
pop this 2
goto IF_END0
label IF_FALSE0
push constant 0
pop this 2
label IF_END0
push constant 0
return
function Controller.tryMovePlayer 10
push argument 0
pop pointer 0
push this 0
call Model.getBoard 1
pop local 0
push local 0
call Board.getPlayer 1
pop local 1
push local 1
call Player.getX 1
pop local 2
push local 1
call Player.getY 1
pop local 3
push local 2
pop local 4
push local 3
pop local 5
push local 4
pop local 6
push local 5
pop local 7
push argument 1
call Direction.U 0
eq
if-goto IF_TRUE0
goto IF_FALSE0
label IF_TRUE0
push local 5
push constant 1
sub
pop local 5
push local 7
push constant 2
sub
pop local 7
label IF_FALSE0
push argument 1
call Direction.D 0
eq
if-goto IF_TRUE1
goto IF_FALSE1
label IF_TRUE1
push local 5
push constant 1
add
pop local 5
push local 7
push constant 2
add
pop local 7
label IF_FALSE1
push argument 1
call Direction.L 0
eq
if-goto IF_TRUE2
goto IF_FALSE2
label IF_TRUE2
push local 4
push constant 1
sub
pop local 4
push local 6
push constant 2
sub
pop local 6
label IF_FALSE2
push argument 1
call Direction.R 0
eq
if-goto IF_TRUE3
goto IF_FALSE3
label IF_TRUE3
push local 4
push constant 1
add
pop local 4
push local 6
push constant 2
add
pop local 6
label IF_FALSE3
push local 0
push local 4
push local 5
call Board.getCell 3
pop local 8
push local 8
call Cell.getTypeOf 1
call CellType.CellTypeWall 0
eq
not
if-goto IF_TRUE4
goto IF_FALSE4
label IF_TRUE4
push local 8
call Cell.getHasBox 1
if-goto IF_TRUE5
goto IF_FALSE5
label IF_TRUE5
push local 0
push local 6
push local 7
call Board.getCell 3
pop local 9
push local 9
call Cell.getTypeOf 1
call CellType.CellTypeWall 0
eq
not
push local 9
call Cell.getHasBox 1
not
and
if-goto IF_TRUE6
goto IF_FALSE6
label IF_TRUE6
push local 8
push constant 0
call Cell.setHasBox 2
pop temp 0
push local 9
push constant 0
not
call Cell.setHasBox 2
pop temp 0
push local 1
push local 4
push local 5
call Player.moveTo 3
pop temp 0
push local 0
call Board.isComplete 1
if-goto IF_TRUE7
goto IF_FALSE7
label IF_TRUE7
push this 0
call State.StateLevelComplete 0
call Model.setState 2
pop temp 0
label IF_FALSE7
label IF_FALSE6
goto IF_END5
label IF_FALSE5
push local 1
push local 4
push local 5
call Player.moveTo 3
pop temp 0
label IF_END5
label IF_FALSE4
push constant 0
return
function Controller.tryStartNextLevel 3
push argument 0
pop pointer 0
push this 0
call Model.getLevelManager 1
pop local 0
push local 0
call LevelManager.hasNextLevel 1
if-goto IF_TRUE0
goto IF_FALSE0
label IF_TRUE0
push local 0
call LevelManager.progressToNextLevel 1
pop temp 0
push local 0
call LevelManager.getCurrentLevel 1
pop local 1
push local 1
call Level.getWidth 1
push local 1
call Level.getHeight 1
push local 1
call Level.getMapData 1
call Board.new 3
pop local 2
push this 0
push local 2
call Model.setBoard 2
pop temp 0
push this 0
call State.StatePlaying 0
call Model.setState 2
pop temp 0
goto IF_END0
label IF_FALSE0
push this 0
call State.StateGameComplete 0
call Model.setState 2
pop temp 0
label IF_END0
push constant 0
return
function Controller.restartLevel 3
push argument 0
pop pointer 0
push this 0
call Model.getLevelManager 1
pop local 0
push local 0
call LevelManager.getCurrentLevel 1
pop local 1
push local 1
call Level.getWidth 1
push local 1
call Level.getHeight 1
push local 1
call Level.getMapData 1
call Board.new 3
pop local 2
push this 0
push local 2
call Model.setBoard 2
pop temp 0
push this 0
call State.StatePlaying 0
call Model.setState 2
pop temp 0
push constant 0
return
function Controller.isRunning 0
push argument 0
pop pointer 0
push this 1
return
function Controller.dispose 0
push argument 0
pop pointer 0
push pointer 0
call Memory.deAlloc 1
pop temp 0
push constant 0
return
`

const Sokoban_Direction_VM string = `function Direction.init 0
push constant 131
pop static 0
push constant 133
pop static 1
push constant 130
pop static 2
push constant 132
pop static 3
push constant 0
return
function Direction.U 0
push static 0
return
function Direction.D 0
push static 1
return
function Direction.L 0
push static 2
return
function Direction.R 0
push static 3
return
`

const Sokoban_Level_VM string = `function Level.new 0
push constant 3
call Memory.alloc 1
pop pointer 0
push argument 0
pop this 0
push argument 1
pop this 1
push argument 2
pop this 2
push pointer 0
return
function Level.getWidth 0
push argument 0
pop pointer 0
push this 0
return
function Level.getHeight 0
push argument 0
pop pointer 0
push this 1
return
function Level.getMapData 0
push argument 0
pop pointer 0
push this 2
return
function Level.dispose 0
push argument 0
pop pointer 0
push pointer 0
call Memory.deAlloc 1
pop temp 0
push constant 0
return
`

const Sokoban_LevelManager_VM string = `function LevelManager.new 0
push constant 3
call Memory.alloc 1
pop pointer 0
push constant 0
pop this 0
push constant 10
pop this 1
push this 1
push constant 1
add
call Array.new 1
pop this 2
push constant 0
push this 2
add
push constant 0
push constant 0
push constant 0
call String.new 1
call Level.new 3
pop temp 0
pop pointer 1
push temp 0
pop that 0
push constant 1
push this 2
add
push constant 8
push constant 8
push constant 64
call String.new 1
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 64
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
call Level.new 3
pop temp 0
pop pointer 1
push temp 0
pop that 0
push constant 2
push this 2
add
push constant 9
push constant 9
push constant 81
call String.new 1
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 64
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
call Level.new 3
pop temp 0
pop pointer 1
push temp 0
pop that 0
push constant 3
push this 2
add
push constant 6
push constant 8
push constant 48
call String.new 1
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 64
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 42
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
call Level.new 3
pop temp 0
pop pointer 1
push temp 0
pop that 0
push constant 4
push this 2
add
push constant 8
push constant 8
push constant 64
call String.new 1
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 64
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
call Level.new 3
pop temp 0
pop pointer 1
push temp 0
pop that 0
push constant 5
push this 2
add
push constant 8
push constant 7
push constant 56
call String.new 1
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 64
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
call Level.new 3
pop temp 0
pop pointer 1
push temp 0
pop that 0
push constant 6
push this 2
add
push constant 8
push constant 7
push constant 56
call String.new 1
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 64
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 42
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
call Level.new 3
pop temp 0
pop pointer 1
push temp 0
pop that 0
push constant 7
push this 2
add
push constant 8
push constant 8
push constant 64
call String.new 1
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 64
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
call Level.new 3
pop temp 0
pop pointer 1
push temp 0
pop that 0
push constant 8
push this 2
add
push constant 8
push constant 7
push constant 56
call String.new 1
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 64
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 42
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
call Level.new 3
pop temp 0
pop pointer 1
push temp 0
pop that 0
push constant 9
push this 2
add
push constant 9
push constant 7
push constant 63
call String.new 1
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 64
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
call Level.new 3
pop temp 0
pop pointer 1
push temp 0
pop that 0
push constant 10
push this 2
add
push constant 7
push constant 8
push constant 56
call String.new 1
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 46
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 36
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 64
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
push constant 35
call String.appendChar 2
call Level.new 3
pop temp 0
pop pointer 1
push temp 0
pop that 0
push pointer 0
return
function LevelManager.getCurrentLevelNumber 0
push argument 0
pop pointer 0
push this 0
return
function LevelManager.getFinalLevelNumber 0
push argument 0
pop pointer 0
push this 1
return
function LevelManager.getCurrentLevel 0
push argument 0
pop pointer 0
push this 0
push this 2
add
pop pointer 1
push that 0
return
function LevelManager.hasNextLevel 0
push argument 0
pop pointer 0
push this 0
push this 1
lt
return
function LevelManager.progressToNextLevel 0
push argument 0
pop pointer 0
push this 0
push constant 1
add
pop this 0
push constant 0
return
function LevelManager.reset 0
push argument 0
pop pointer 0
push constant 0
pop this 0
push constant 0
return
function LevelManager.dispose 0
push argument 0
pop pointer 0
push this 2
call Array.dispose 1
pop temp 0
push pointer 0
call Memory.deAlloc 1
pop temp 0
push constant 0
return
`

const Sokoban_Main_VM string = `function Main.main 4
call CellType.init 0
pop temp 0
call Direction.init 0
pop temp 0
call State.init 0
pop temp 0
call Model.new 0
pop local 1
push local 1
call View.new 1
pop local 2
push local 1
call Controller.new 1
pop local 3
push local 3
call Controller.startNewGame 1
pop temp 0
label WHILE_EXP0
push local 3
call Controller.isRunning 1
not
if-goto WHILE_END0
push local 3
call Controller.handleInput 1
pop temp 0
push local 2
call View.draw 1
pop temp 0
push local 1
call Model.update 1
pop temp 0
push constant 10
call Sys.wait 1
pop temp 0
goto WHILE_EXP0
label WHILE_END0
push local 1
call Model.dispose 1
pop temp 0
push local 2
call View.dispose 1
pop temp 0
push local 3
call Controller.dispose 1
pop temp 0
push constant 0
return
`

const Sokoban_Model_VM string = `function Model.new 0
push constant 5
call Memory.alloc 1
pop pointer 0
call LevelManager.new 0
pop this 0
call State.StatePlaying 0
pop this 2
push constant 0
pop this 3
push constant 0
pop this 4
push pointer 0
return
function Model.update 0
push argument 0
pop pointer 0
push this 3
push constant 1
add
pop this 3
push this 3
push constant 20
gt
if-goto IF_TRUE0
goto IF_FALSE0
label IF_TRUE0
push constant 0
pop this 3
label IF_FALSE0
push this 4
if-goto IF_TRUE1
goto IF_FALSE1
label IF_TRUE1
push constant 0
pop this 4
label IF_FALSE1
push constant 0
return
function Model.getLevelManager 0
push argument 0
pop pointer 0
push this 0
return
function Model.getBoard 0
push argument 0
pop pointer 0
push this 1
return
function Model.getState 0
push argument 0
pop pointer 0
push this 2
return
function Model.getTickAccumulator 0
push argument 0
pop pointer 0
push this 3
return
function Model.isScreenDirty 0
push argument 0
pop pointer 0
push this 4
return
function Model.setBoard 0
push argument 0
pop pointer 0
push this 1
push constant 0
eq
not
if-goto IF_TRUE0
goto IF_FALSE0
label IF_TRUE0
push this 1
call Board.dispose 1
pop temp 0
label IF_FALSE0
push argument 1
pop this 1
push constant 0
return
function Model.setState 0
push argument 0
pop pointer 0
push argument 1
pop this 2
push constant 0
not
pop this 4
push constant 0
return
function Model.dispose 0
push argument 0
pop pointer 0
push this 0
call LevelManager.dispose 1
pop temp 0
push this 1
call Board.dispose 1
pop temp 0
push pointer 0
call Memory.deAlloc 1
pop temp 0
push constant 0
return
`

const Sokoban_Player_VM string = `function Player.new 0
push constant 2
call Memory.alloc 1
pop pointer 0
push argument 0
pop this 0
push argument 1
pop this 1
push pointer 0
return
function Player.getX 0
push argument 0
pop pointer 0
push this 0
return
function Player.getY 0
push argument 0
pop pointer 0
push this 1
return
function Player.moveTo 0
push argument 0
pop pointer 0
push argument 1
pop this 0
push argument 2
pop this 1
push constant 0
return
function Player.dispose 0
push argument 0
pop pointer 0
push pointer 0
call Memory.deAlloc 1
pop temp 0
push constant 0
return
`

const Sokoban_Sprites_VM string = `function Sprites.drawBox 1
push constant 16384
push argument 0
add
pop local 0
push local 0
push constant 0
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 32
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 64
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 96
add
push constant 4080
call Memory.poke 2
pop temp 0
push local 0
push constant 128
add
push constant 8184
call Memory.poke 2
pop temp 0
push local 0
push constant 160
add
push constant 8184
call Memory.poke 2
pop temp 0
push local 0
push constant 192
add
push constant 8184
call Memory.poke 2
pop temp 0
push local 0
push constant 224
add
push constant 8184
call Memory.poke 2
pop temp 0
push local 0
push constant 256
add
push constant 8184
call Memory.poke 2
pop temp 0
push local 0
push constant 288
add
push constant 8184
call Memory.poke 2
pop temp 0
push local 0
push constant 320
add
push constant 8184
call Memory.poke 2
pop temp 0
push local 0
push constant 352
add
push constant 8184
call Memory.poke 2
pop temp 0
push local 0
push constant 384
add
push constant 4080
call Memory.poke 2
pop temp 0
push local 0
push constant 416
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 448
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 480
add
push constant 0
call Memory.poke 2
pop temp 0
push constant 0
return
function Sprites.drawClear 1
push constant 16384
push argument 0
add
pop local 0
push local 0
push constant 0
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 32
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 64
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 96
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 128
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 160
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 192
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 224
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 256
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 288
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 320
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 352
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 384
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 416
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 448
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 480
add
push constant 0
call Memory.poke 2
pop temp 0
push constant 0
return
function Sprites.drawGoal 1
push constant 16384
push argument 0
add
pop local 0
push local 0
push constant 0
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 32
add
push constant 16380
call Memory.poke 2
pop temp 0
push local 0
push constant 64
add
push constant 24582
call Memory.poke 2
pop temp 0
push local 0
push constant 96
add
push constant 16386
call Memory.poke 2
pop temp 0
push local 0
push constant 128
add
push constant 16386
call Memory.poke 2
pop temp 0
push local 0
push constant 160
add
push constant 16386
call Memory.poke 2
pop temp 0
push local 0
push constant 192
add
push constant 16386
call Memory.poke 2
pop temp 0
push local 0
push constant 224
add
push constant 16386
call Memory.poke 2
pop temp 0
push local 0
push constant 256
add
push constant 16386
call Memory.poke 2
pop temp 0
push local 0
push constant 288
add
push constant 16386
call Memory.poke 2
pop temp 0
push local 0
push constant 320
add
push constant 16386
call Memory.poke 2
pop temp 0
push local 0
push constant 352
add
push constant 16386
call Memory.poke 2
pop temp 0
push local 0
push constant 384
add
push constant 16386
call Memory.poke 2
pop temp 0
push local 0
push constant 416
add
push constant 24582
call Memory.poke 2
pop temp 0
push local 0
push constant 448
add
push constant 16380
call Memory.poke 2
pop temp 0
push local 0
push constant 480
add
push constant 0
call Memory.poke 2
pop temp 0
push constant 0
return
function Sprites.drawGoalBox 1
push constant 16384
push argument 0
add
pop local 0
push local 0
push constant 0
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 32
add
push constant 16380
call Memory.poke 2
pop temp 0
push local 0
push constant 64
add
push constant 24582
call Memory.poke 2
pop temp 0
push local 0
push constant 96
add
push constant 20466
call Memory.poke 2
pop temp 0
push local 0
push constant 128
add
push constant 24570
call Memory.poke 2
pop temp 0
push local 0
push constant 160
add
push constant 24570
call Memory.poke 2
pop temp 0
push local 0
push constant 192
add
push constant 24570
call Memory.poke 2
pop temp 0
push local 0
push constant 224
add
push constant 24570
call Memory.poke 2
pop temp 0
push local 0
push constant 256
add
push constant 24570
call Memory.poke 2
pop temp 0
push local 0
push constant 288
add
push constant 24570
call Memory.poke 2
pop temp 0
push local 0
push constant 320
add
push constant 24570
call Memory.poke 2
pop temp 0
push local 0
push constant 352
add
push constant 24570
call Memory.poke 2
pop temp 0
push local 0
push constant 384
add
push constant 20466
call Memory.poke 2
pop temp 0
push local 0
push constant 416
add
push constant 24582
call Memory.poke 2
pop temp 0
push local 0
push constant 448
add
push constant 16380
call Memory.poke 2
pop temp 0
push local 0
push constant 480
add
push constant 0
call Memory.poke 2
pop temp 0
push constant 0
return
function Sprites.drawGoalPlayer 1
push constant 16384
push argument 0
add
pop local 0
push local 0
push constant 0
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 32
add
push constant 16380
call Memory.poke 2
pop temp 0
push local 0
push constant 64
add
push constant 24582
call Memory.poke 2
pop temp 0
push local 0
push constant 96
add
push constant 16386
call Memory.poke 2
pop temp 0
push local 0
push constant 128
add
push constant 18402
call Memory.poke 2
pop temp 0
push local 0
push constant 160
add
push constant 20466
call Memory.poke 2
pop temp 0
push local 0
push constant 192
add
push constant 19890
call Memory.poke 2
pop temp 0
push local 0
push constant 224
add
push constant 23994
call Memory.poke 2
pop temp 0
push local 0
push constant 256
add
push constant 24570
call Memory.poke 2
pop temp 0
push local 0
push constant 288
add
push constant 19410
call Memory.poke 2
pop temp 0
push local 0
push constant 320
add
push constant 19506
call Memory.poke 2
pop temp 0
push local 0
push constant 352
add
push constant 18402
call Memory.poke 2
pop temp 0
push local 0
push constant 384
add
push constant 16386
call Memory.poke 2
pop temp 0
push local 0
push constant 416
add
push constant 24582
call Memory.poke 2
pop temp 0
push local 0
push constant 448
add
push constant 16380
call Memory.poke 2
pop temp 0
push local 0
push constant 480
add
push constant 0
call Memory.poke 2
pop temp 0
push constant 0
return
function Sprites.drawLogoA1 1
push constant 16384
push argument 0
add
pop local 0
push local 0
push constant 0
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 32
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 64
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 96
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 128
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 160
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 192
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 224
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 256
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 288
add
push constant 4096
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 320
add
push constant 1024
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 352
add
push constant 512
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 384
add
push constant 7936
call Memory.poke 2
pop temp 0
push local 0
push constant 416
add
push constant 16128
call Memory.poke 2
pop temp 0
push local 0
push constant 448
add
push constant 256
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 480
add
push constant 512
neg
call Memory.poke 2
pop temp 0
push constant 0
return
function Sprites.drawLogoA2 1
push constant 16384
push argument 0
add
pop local 0
push local 0
push constant 0
add
push constant 2048
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 32
add
push constant 16384
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 64
add
push constant 28928
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 96
add
push constant 256
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 128
add
push constant 512
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 160
add
push constant 512
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 192
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 224
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 256
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 288
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 320
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 352
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 384
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 416
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 448
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 480
add
push constant 0
call Memory.poke 2
pop temp 0
push constant 0
return
function Sprites.drawLogoA3 1
push constant 16384
push argument 0
add
pop local 0
push local 0
push constant 0
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 32
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 64
add
push constant 1024
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 96
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 128
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 160
add
push constant 16384
call Memory.poke 2
pop temp 0
push local 0
push constant 192
add
push constant 8192
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 224
add
push constant 20480
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 256
add
push constant 2048
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 288
add
push constant 6144
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 320
add
push constant 10240
call Memory.poke 2
pop temp 0
push local 0
push constant 352
add
push constant 16384
call Memory.poke 2
pop temp 0
push local 0
push constant 384
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 416
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 448
add
push constant 1024
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 480
add
push constant 0
call Memory.poke 2
pop temp 0
push constant 0
return
function Sprites.drawLogoB1 1
push constant 16384
push argument 0
add
pop local 0
push local 0
push constant 0
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 32
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 64
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 96
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 128
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 160
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 192
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 224
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 256
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 288
add
push constant 15879
call Memory.poke 2
pop temp 0
push local 0
push constant 320
add
push constant 121
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 352
add
push constant 49
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 384
add
push constant 2097
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 416
add
push constant 7200
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 448
add
push constant 7199
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 480
add
push constant 7193
neg
call Memory.poke 2
pop temp 0
push constant 0
return
function Sprites.drawLogoB2 1
push constant 16384
push argument 0
add
pop local 0
push local 0
push constant 0
add
push constant 7185
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 32
add
push constant 7185
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 64
add
push constant 2097
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 96
add
push constant 57
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 128
add
push constant 125
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 160
add
push constant 15872
call Memory.poke 2
pop temp 0
push local 0
push constant 192
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 224
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 256
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 288
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 320
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 352
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 384
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 416
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 448
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 480
add
push constant 0
call Memory.poke 2
pop temp 0
push constant 0
return
function Sprites.drawLogoB3 1
push constant 16384
push argument 0
add
pop local 0
push local 0
push constant 0
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 32
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 64
add
push constant 1
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 96
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 128
add
push constant 24320
call Memory.poke 2
pop temp 0
push local 0
push constant 160
add
push constant 17412
call Memory.poke 2
pop temp 0
push local 0
push constant 192
add
push constant 17423
call Memory.poke 2
pop temp 0
push local 0
push constant 224
add
push constant 17435
call Memory.poke 2
pop temp 0
push local 0
push constant 256
add
push constant 15297
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 288
add
push constant 17455
call Memory.poke 2
pop temp 0
push local 0
push constant 320
add
push constant 17448
call Memory.poke 2
pop temp 0
push local 0
push constant 352
add
push constant 17412
call Memory.poke 2
pop temp 0
push local 0
push constant 384
add
push constant 17408
call Memory.poke 2
pop temp 0
push local 0
push constant 416
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 448
add
push constant 1
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 480
add
push constant 0
call Memory.poke 2
pop temp 0
push constant 0
return
function Sprites.drawLogoC1 1
push constant 16384
push argument 0
add
pop local 0
push local 0
push constant 0
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 32
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 64
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 96
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 128
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 160
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 192
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 224
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 256
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 288
add
push constant 4344
call Memory.poke 2
pop temp 0
push local 0
push constant 320
add
push constant 14584
call Memory.poke 2
pop temp 0
push local 0
push constant 352
add
push constant 775
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 384
add
push constant 32505
call Memory.poke 2
pop temp 0
push local 0
push constant 416
add
push constant 16371
call Memory.poke 2
pop temp 0
push local 0
push constant 448
add
push constant 8179
call Memory.poke 2
pop temp 0
push local 0
push constant 480
add
push constant 4083
call Memory.poke 2
pop temp 0
push constant 0
return
function Sprites.drawLogoC2 1
push constant 16384
push argument 0
add
pop local 0
push local 0
push constant 0
add
push constant 8179
call Memory.poke 2
pop temp 0
push local 0
push constant 32
add
push constant 16371
call Memory.poke 2
pop temp 0
push local 0
push constant 64
add
push constant 32505
call Memory.poke 2
pop temp 0
push local 0
push constant 96
add
push constant 775
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 128
add
push constant 30968
call Memory.poke 2
pop temp 0
push local 0
push constant 160
add
push constant 6392
call Memory.poke 2
pop temp 0
push local 0
push constant 192
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 224
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 256
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 288
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 320
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 352
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 384
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 416
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 448
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 480
add
push constant 0
call Memory.poke 2
pop temp 0
push constant 0
return
function Sprites.drawLogoC3 1
push constant 16384
push argument 0
add
pop local 0
push local 0
push constant 0
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 32
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 64
add
push constant 1
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 96
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 128
add
push constant 32244
call Memory.poke 2
pop temp 0
push local 0
push constant 160
add
push constant 4116
call Memory.poke 2
pop temp 0
push local 0
push constant 192
add
push constant 4116
call Memory.poke 2
pop temp 0
push local 0
push constant 224
add
push constant 4116
call Memory.poke 2
pop temp 0
push local 0
push constant 256
add
push constant 4215
call Memory.poke 2
pop temp 0
push local 0
push constant 288
add
push constant 4116
call Memory.poke 2
pop temp 0
push local 0
push constant 320
add
push constant 4116
call Memory.poke 2
pop temp 0
push local 0
push constant 352
add
push constant 4116
call Memory.poke 2
pop temp 0
push local 0
push constant 384
add
push constant 32244
call Memory.poke 2
pop temp 0
push local 0
push constant 416
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 448
add
push constant 1
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 480
add
push constant 0
call Memory.poke 2
pop temp 0
push constant 0
return
function Sprites.drawLogoD1 1
push constant 16384
push argument 0
add
pop local 0
push local 0
push constant 0
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 32
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 64
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 96
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 128
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 160
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 192
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 224
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 256
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 288
add
push constant 31776
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 320
add
push constant 28680
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 352
add
push constant 24580
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 384
add
push constant 24708
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 416
add
push constant 15934
call Memory.poke 2
pop temp 0
push local 0
push constant 448
add
push constant 15934
call Memory.poke 2
pop temp 0
push local 0
push constant 480
add
push constant 15934
call Memory.poke 2
pop temp 0
push constant 0
return
function Sprites.drawLogoD2 1
push constant 16384
push argument 0
add
pop local 0
push local 0
push constant 0
add
push constant 15934
call Memory.poke 2
pop temp 0
push local 0
push constant 32
add
push constant 15934
call Memory.poke 2
pop temp 0
push local 0
push constant 64
add
push constant 24708
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 96
add
push constant 24580
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 128
add
push constant 28680
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 160
add
push constant 31776
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 192
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 224
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 256
add
push constant 8
call Memory.poke 2
pop temp 0
push local 0
push constant 288
add
push constant 8
call Memory.poke 2
pop temp 0
push local 0
push constant 320
add
push constant 2360
call Memory.poke 2
pop temp 0
push local 0
push constant 352
add
push constant 2376
call Memory.poke 2
pop temp 0
push local 0
push constant 384
add
push constant 2376
call Memory.poke 2
pop temp 0
push local 0
push constant 416
add
push constant 3640
call Memory.poke 2
pop temp 0
push local 0
push constant 448
add
push constant 2048
call Memory.poke 2
pop temp 0
push local 0
push constant 480
add
push constant 1792
call Memory.poke 2
pop temp 0
push constant 0
return
function Sprites.drawLogoD3 1
push constant 16384
push argument 0
add
pop local 0
push local 0
push constant 0
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 32
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 64
add
push constant 1
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 96
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 128
add
push constant 7087
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 160
add
push constant 5203
call Memory.poke 2
pop temp 0
push local 0
push constant 192
add
push constant 5205
call Memory.poke 2
pop temp 0
push local 0
push constant 224
add
push constant 5209
call Memory.poke 2
pop temp 0
push local 0
push constant 256
add
push constant 2991
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 288
add
push constant 5201
call Memory.poke 2
pop temp 0
push local 0
push constant 320
add
push constant 5201
call Memory.poke 2
pop temp 0
push local 0
push constant 352
add
push constant 4753
call Memory.poke 2
pop temp 0
push local 0
push constant 384
add
push constant 4369
call Memory.poke 2
pop temp 0
push local 0
push constant 416
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 448
add
push constant 1
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 480
add
push constant 0
call Memory.poke 2
pop temp 0
push constant 0
return
function Sprites.drawLogoE1 1
push constant 16384
push argument 0
add
pop local 0
push local 0
push constant 0
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 32
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 64
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 96
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 128
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 160
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 192
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 224
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 256
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 288
add
push constant 255
call Memory.poke 2
pop temp 0
push local 0
push constant 320
add
push constant 511
call Memory.poke 2
pop temp 0
push local 0
push constant 352
add
push constant 1023
call Memory.poke 2
pop temp 0
push local 0
push constant 384
add
push constant 1999
call Memory.poke 2
pop temp 0
push local 0
push constant 416
add
push constant 1999
call Memory.poke 2
pop temp 0
push local 0
push constant 448
add
push constant 1023
call Memory.poke 2
pop temp 0
push local 0
push constant 480
add
push constant 31745
neg
call Memory.poke 2
pop temp 0
push constant 0
return
function Sprites.drawLogoE2 1
push constant 16384
push argument 0
add
pop local 0
push local 0
push constant 0
add
push constant 30721
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 32
add
push constant 28785
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 64
add
push constant 12337
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 96
add
push constant 12289
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 128
add
push constant 6145
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 160
add
push constant 7169
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 192
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 224
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 256
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 288
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 320
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 352
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 384
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 416
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 448
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 480
add
push constant 0
call Memory.poke 2
pop temp 0
push constant 0
return
function Sprites.drawLogoE3 1
push constant 16384
push argument 0
add
pop local 0
push local 0
push constant 0
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 32
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 64
add
push constant 1
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 96
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 128
add
push constant 8388
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 160
add
push constant 16709
call Memory.poke 2
pop temp 0
push local 0
push constant 192
add
push constant 16709
call Memory.poke 2
pop temp 0
push local 0
push constant 224
add
push constant 16709
call Memory.poke 2
pop temp 0
push local 0
push constant 256
add
push constant 14523
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 288
add
push constant 16709
call Memory.poke 2
pop temp 0
push local 0
push constant 320
add
push constant 16709
call Memory.poke 2
pop temp 0
push local 0
push constant 352
add
push constant 16709
call Memory.poke 2
pop temp 0
push local 0
push constant 384
add
push constant 24381
call Memory.poke 2
pop temp 0
push local 0
push constant 416
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 448
add
push constant 1
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 480
add
push constant 0
call Memory.poke 2
pop temp 0
push constant 0
return
function Sprites.drawLogoF1 1
push constant 16384
push argument 0
add
pop local 0
push local 0
push constant 0
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 32
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 64
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 96
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 128
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 160
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 192
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 224
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 256
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 288
add
push constant 1986
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 320
add
push constant 1986
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 352
add
push constant 1986
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 384
add
push constant 1921
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 416
add
push constant 3969
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 448
add
push constant 3977
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 480
add
push constant 3849
neg
call Memory.poke 2
pop temp 0
push constant 0
return
function Sprites.drawLogoF2 1
push constant 16384
push argument 0
add
pop local 0
push local 0
push constant 0
add
push constant 3849
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 32
add
push constant 3841
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 64
add
push constant 1537
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 96
add
push constant 1537
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 128
add
push constant 1053
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 160
add
push constant 1053
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 192
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 224
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 256
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 288
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 320
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 352
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 384
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 416
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 448
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 480
add
push constant 0
call Memory.poke 2
pop temp 0
push constant 0
return
function Sprites.drawLogoF3 1
push constant 16384
push argument 0
add
pop local 0
push local 0
push constant 0
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 32
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 64
add
push constant 1
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 96
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 128
add
push constant 30947
call Memory.poke 2
pop temp 0
push local 0
push constant 160
add
push constant 1300
call Memory.poke 2
pop temp 0
push local 0
push constant 192
add
push constant 1284
call Memory.poke 2
pop temp 0
push local 0
push constant 224
add
push constant 1284
call Memory.poke 2
pop temp 0
push local 0
push constant 256
add
push constant 15555
call Memory.poke 2
pop temp 0
push local 0
push constant 288
add
push constant 17666
call Memory.poke 2
pop temp 0
push local 0
push constant 320
add
push constant 17668
call Memory.poke 2
pop temp 0
push local 0
push constant 352
add
push constant 17684
call Memory.poke 2
pop temp 0
push local 0
push constant 384
add
push constant 14564
call Memory.poke 2
pop temp 0
push local 0
push constant 416
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 448
add
push constant 1
neg
call Memory.poke 2
pop temp 0
push local 0
push constant 480
add
push constant 0
call Memory.poke 2
pop temp 0
push constant 0
return
function Sprites.drawLogoG1 1
push constant 16384
push argument 0
add
pop local 0
push local 0
push constant 0
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 32
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 64
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 96
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 128
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 160
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 192
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 224
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 256
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 288
add
push constant 248
call Memory.poke 2
pop temp 0
push local 0
push constant 320
add
push constant 248
call Memory.poke 2
pop temp 0
push local 0
push constant 352
add
push constant 249
call Memory.poke 2
pop temp 0
push local 0
push constant 384
add
push constant 251
call Memory.poke 2
pop temp 0
push local 0
push constant 416
add
push constant 123
call Memory.poke 2
pop temp 0
push local 0
push constant 448
add
push constant 127
call Memory.poke 2
pop temp 0
push local 0
push constant 480
add
push constant 127
call Memory.poke 2
pop temp 0
push constant 0
return
function Sprites.drawLogoG2 1
push constant 16384
push argument 0
add
pop local 0
push local 0
push constant 0
add
push constant 127
call Memory.poke 2
pop temp 0
push local 0
push constant 32
add
push constant 126
call Memory.poke 2
pop temp 0
push local 0
push constant 64
add
push constant 254
call Memory.poke 2
pop temp 0
push local 0
push constant 96
add
push constant 252
call Memory.poke 2
pop temp 0
push local 0
push constant 128
add
push constant 248
call Memory.poke 2
pop temp 0
push local 0
push constant 160
add
push constant 248
call Memory.poke 2
pop temp 0
push local 0
push constant 192
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 224
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 256
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 288
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 320
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 352
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 384
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 416
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 448
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 480
add
push constant 0
call Memory.poke 2
pop temp 0
push constant 0
return
function Sprites.drawLogoG3 1
push constant 16384
push argument 0
add
pop local 0
push local 0
push constant 0
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 32
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 64
add
push constant 63
call Memory.poke 2
pop temp 0
push local 0
push constant 96
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 128
add
push constant 14
call Memory.poke 2
pop temp 0
push local 0
push constant 160
add
push constant 17
call Memory.poke 2
pop temp 0
push local 0
push constant 192
add
push constant 17
call Memory.poke 2
pop temp 0
push local 0
push constant 224
add
push constant 25
call Memory.poke 2
pop temp 0
push local 0
push constant 256
add
push constant 21
call Memory.poke 2
pop temp 0
push local 0
push constant 288
add
push constant 19
call Memory.poke 2
pop temp 0
push local 0
push constant 320
add
push constant 17
call Memory.poke 2
pop temp 0
push local 0
push constant 352
add
push constant 17
call Memory.poke 2
pop temp 0
push local 0
push constant 384
add
push constant 14
call Memory.poke 2
pop temp 0
push local 0
push constant 416
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 448
add
push constant 63
call Memory.poke 2
pop temp 0
push local 0
push constant 480
add
push constant 0
call Memory.poke 2
pop temp 0
push constant 0
return
function Sprites.drawPlayer 1
push constant 16384
push argument 0
add
pop local 0
push local 0
push constant 0
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 32
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 64
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 96
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 128
add
push constant 2016
call Memory.poke 2
pop temp 0
push local 0
push constant 160
add
push constant 4080
call Memory.poke 2
pop temp 0
push local 0
push constant 192
add
push constant 3504
call Memory.poke 2
pop temp 0
push local 0
push constant 224
add
push constant 7608
call Memory.poke 2
pop temp 0
push local 0
push constant 256
add
push constant 8184
call Memory.poke 2
pop temp 0
push local 0
push constant 288
add
push constant 3024
call Memory.poke 2
pop temp 0
push local 0
push constant 320
add
push constant 3120
call Memory.poke 2
pop temp 0
push local 0
push constant 352
add
push constant 2016
call Memory.poke 2
pop temp 0
push local 0
push constant 384
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 416
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 448
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 480
add
push constant 0
call Memory.poke 2
pop temp 0
push constant 0
return
function Sprites.drawWall 1
push constant 16384
push argument 0
add
pop local 0
push local 0
push constant 0
add
push constant 0
call Memory.poke 2
pop temp 0
push local 0
push constant 32
add
push constant 32766
call Memory.poke 2
pop temp 0
push local 0
push constant 64
add
push constant 16386
call Memory.poke 2
pop temp 0
push local 0
push constant 96
add
push constant 19114
call Memory.poke 2
pop temp 0
push local 0
push constant 128
add
push constant 21842
call Memory.poke 2
pop temp 0
push local 0
push constant 160
add
push constant 19114
call Memory.poke 2
pop temp 0
push local 0
push constant 192
add
push constant 21842
call Memory.poke 2
pop temp 0
push local 0
push constant 224
add
push constant 19114
call Memory.poke 2
pop temp 0
push local 0
push constant 256
add
push constant 21842
call Memory.poke 2
pop temp 0
push local 0
push constant 288
add
push constant 19114
call Memory.poke 2
pop temp 0
push local 0
push constant 320
add
push constant 21842
call Memory.poke 2
pop temp 0
push local 0
push constant 352
add
push constant 19114
call Memory.poke 2
pop temp 0
push local 0
push constant 384
add
push constant 21842
call Memory.poke 2
pop temp 0
push local 0
push constant 416
add
push constant 16386
call Memory.poke 2
pop temp 0
push local 0
push constant 448
add
push constant 32766
call Memory.poke 2
pop temp 0
push local 0
push constant 480
add
push constant 0
call Memory.poke 2
pop temp 0
push constant 0
return
`

const Sokoban_State_VM string = `function State.init 0
push constant 0
pop static 0
push constant 1
pop static 1
push constant 2
pop static 2
push constant 0
return
function State.StatePlaying 0
push static 0
return
function State.StateLevelComplete 0
push static 1
return
function State.StateGameComplete 0
push static 2
return
`

const Sokoban_Utils_VM string = `function Utils.mod 1
push argument 0
push argument 1
call Math.divide 2
pop local 0
push argument 0
push local 0
push argument 1
call Math.multiply 2
sub
return
`

const Sokoban_View_VM string = `function View.new 0
push constant 10
call Memory.alloc 1
pop pointer 0
push argument 0
pop this 0
push constant 14
call String.new 1
push constant 45
call String.appendChar 2
push constant 45
call String.appendChar 2
push constant 45
call String.appendChar 2
push constant 67
call String.appendChar 2
push constant 111
call String.appendChar 2
push constant 110
call String.appendChar 2
push constant 116
call String.appendChar 2
push constant 114
call String.appendChar 2
push constant 111
call String.appendChar 2
push constant 108
call String.appendChar 2
push constant 115
call String.appendChar 2
push constant 45
call String.appendChar 2
push constant 45
call String.appendChar 2
push constant 45
call String.appendChar 2
pop this 1
push constant 14
call String.new 1
push constant 67
call String.appendChar 2
push constant 117
call String.appendChar 2
push constant 114
call String.appendChar 2
push constant 115
call String.appendChar 2
push constant 111
call String.appendChar 2
push constant 114
call String.appendChar 2
push constant 115
call String.appendChar 2
push constant 58
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 77
call String.appendChar 2
push constant 111
call String.appendChar 2
push constant 118
call String.appendChar 2
push constant 101
call String.appendChar 2
pop this 2
push constant 14
call String.new 1
push constant 82
call String.appendChar 2
push constant 58
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 82
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 115
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 116
call String.appendChar 2
pop this 3
push constant 14
call String.new 1
push constant 69
call String.appendChar 2
push constant 115
call String.appendChar 2
push constant 99
call String.appendChar 2
push constant 97
call String.appendChar 2
push constant 112
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 58
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 81
call String.appendChar 2
push constant 117
call String.appendChar 2
push constant 105
call String.appendChar 2
push constant 116
call String.appendChar 2
pop this 4
push constant 14
call String.new 1
push constant 83
call String.appendChar 2
push constant 112
call String.appendChar 2
push constant 97
call String.appendChar 2
push constant 99
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 58
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 78
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 120
call String.appendChar 2
push constant 116
call String.appendChar 2
pop this 5
push constant 14
call String.new 1
push constant 83
call String.appendChar 2
push constant 112
call String.appendChar 2
push constant 97
call String.appendChar 2
push constant 99
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 58
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 82
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 115
call String.appendChar 2
push constant 116
call String.appendChar 2
push constant 97
call String.appendChar 2
push constant 114
call String.appendChar 2
push constant 116
call String.appendChar 2
pop this 6
push constant 14
call String.new 1
push constant 76
call String.appendChar 2
push constant 69
call String.appendChar 2
push constant 86
call String.appendChar 2
push constant 69
call String.appendChar 2
push constant 76
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 67
call String.appendChar 2
push constant 79
call String.appendChar 2
push constant 77
call String.appendChar 2
push constant 80
call String.appendChar 2
push constant 76
call String.appendChar 2
push constant 69
call String.appendChar 2
push constant 84
call String.appendChar 2
push constant 69
call String.appendChar 2
pop this 7
push constant 14
call String.new 1
push constant 71
call String.appendChar 2
push constant 65
call String.appendChar 2
push constant 77
call String.appendChar 2
push constant 69
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 67
call String.appendChar 2
push constant 79
call String.appendChar 2
push constant 77
call String.appendChar 2
push constant 80
call String.appendChar 2
push constant 76
call String.appendChar 2
push constant 69
call String.appendChar 2
push constant 84
call String.appendChar 2
push constant 69
call String.appendChar 2
push constant 33
call String.appendChar 2
pop this 8
push constant 16
call String.new 1
push constant 67
call String.appendChar 2
push constant 79
call String.appendChar 2
push constant 78
call String.appendChar 2
push constant 71
call String.appendChar 2
push constant 82
call String.appendChar 2
push constant 65
call String.appendChar 2
push constant 84
call String.appendChar 2
push constant 85
call String.appendChar 2
push constant 76
call String.appendChar 2
push constant 65
call String.appendChar 2
push constant 84
call String.appendChar 2
push constant 73
call String.appendChar 2
push constant 79
call String.appendChar 2
push constant 78
call String.appendChar 2
push constant 83
call String.appendChar 2
push constant 33
call String.appendChar 2
pop this 9
push pointer 0
return
function View.draw 4
push argument 0
pop pointer 0
push this 0
call Model.getBoard 1
pop local 2
push local 2
call Board.getPlayer 1
pop local 3
push this 0
call Model.isScreenDirty 1
if-goto IF_TRUE0
goto IF_FALSE0
label IF_TRUE0
push constant 0
call Screen.setColor 1
pop temp 0
push constant 0
push constant 0
push constant 511
push constant 255
call Screen.drawRectangle 4
pop temp 0
label IF_FALSE0
push pointer 0
call View.drawLogo 1
pop temp 0
push this 0
call Model.getState 1
call State.StatePlaying 0
eq
if-goto IF_TRUE1
goto IF_FALSE1
label IF_TRUE1
push pointer 0
call View.drawBoard 1
pop temp 0
push pointer 0
call View.drawLevelInfo 1
pop temp 0
push constant 17
push constant 48
call Output.moveCursor 2
pop temp 0
push this 1
call Output.printString 1
pop temp 0
push constant 19
push constant 48
call Output.moveCursor 2
pop temp 0
push this 2
call Output.printString 1
pop temp 0
push constant 20
push constant 48
call Output.moveCursor 2
pop temp 0
push this 3
call Output.printString 1
pop temp 0
push constant 21
push constant 48
call Output.moveCursor 2
pop temp 0
push this 4
call Output.printString 1
pop temp 0
label IF_FALSE1
push this 0
call Model.getState 1
call State.StateLevelComplete 0
eq
if-goto IF_TRUE2
goto IF_FALSE2
label IF_TRUE2
push pointer 0
call View.drawBoard 1
pop temp 0
push pointer 0
call View.drawLevelInfo 1
pop temp 0
push this 0
call Model.getTickAccumulator 1
push constant 10
lt
if-goto IF_TRUE3
goto IF_FALSE3
label IF_TRUE3
push constant 12
push constant 48
call Output.moveCursor 2
pop temp 0
push this 7
call Output.printString 1
pop temp 0
goto IF_END3
label IF_FALSE3
push constant 0
call Screen.setColor 1
pop temp 0
push constant 384
push constant 133
push constant 494
push constant 142
call Screen.drawRectangle 4
pop temp 0
label IF_END3
push constant 17
push constant 48
call Output.moveCursor 2
pop temp 0
push this 1
call Output.printString 1
pop temp 0
push constant 19
push constant 48
call Output.moveCursor 2
pop temp 0
push this 5
call Output.printString 1
pop temp 0
push constant 21
push constant 48
call Output.moveCursor 2
pop temp 0
push this 4
call Output.printString 1
pop temp 0
label IF_FALSE2
push this 0
call Model.getState 1
call State.StateGameComplete 0
eq
if-goto IF_TRUE4
goto IF_FALSE4
label IF_TRUE4
push constant 10
push constant 16
call Output.moveCursor 2
pop temp 0
push this 8
call Output.printString 1
pop temp 0
push this 0
call Model.getTickAccumulator 1
push constant 10
lt
if-goto IF_TRUE5
goto IF_FALSE5
label IF_TRUE5
push constant 12
push constant 15
call Output.moveCursor 2
pop temp 0
push this 9
call Output.printString 1
pop temp 0
push pointer 0
push constant 0
call View.drawGameCompleteBorder 2
pop temp 0
goto IF_END5
label IF_FALSE5
push constant 0
call Screen.setColor 1
pop temp 0
push constant 120
push constant 133
push constant 244
push constant 141
call Screen.drawRectangle 4
pop temp 0
push pointer 0
push constant 0
not
call View.drawGameCompleteBorder 2
pop temp 0
label IF_END5
push constant 17
push constant 48
call Output.moveCursor 2
pop temp 0
push this 1
call Output.printString 1
pop temp 0
push constant 19
push constant 48
call Output.moveCursor 2
pop temp 0
push this 6
call Output.printString 1
pop temp 0
push constant 21
push constant 48
call Output.moveCursor 2
pop temp 0
push this 4
call Output.printString 1
pop temp 0
label IF_FALSE4
push constant 0
return
function View.drawLogo 0
push argument 0
pop pointer 0
push constant 56
push constant 0
push constant 512
call Math.multiply 2
add
call Sprites.drawLogoA1 1
pop temp 0
push constant 57
push constant 0
push constant 512
call Math.multiply 2
add
call Sprites.drawLogoB1 1
pop temp 0
push constant 58
push constant 0
push constant 512
call Math.multiply 2
add
call Sprites.drawLogoC1 1
pop temp 0
push constant 59
push constant 0
push constant 512
call Math.multiply 2
add
call Sprites.drawLogoD1 1
pop temp 0
push constant 60
push constant 0
push constant 512
call Math.multiply 2
add
call Sprites.drawLogoE1 1
pop temp 0
push constant 61
push constant 0
push constant 512
call Math.multiply 2
add
call Sprites.drawLogoF1 1
pop temp 0
push constant 62
push constant 0
push constant 512
call Math.multiply 2
add
call Sprites.drawLogoG1 1
pop temp 0
push constant 56
push constant 1
push constant 512
call Math.multiply 2
add
call Sprites.drawLogoA2 1
pop temp 0
push constant 57
push constant 1
push constant 512
call Math.multiply 2
add
call Sprites.drawLogoB2 1
pop temp 0
push constant 58
push constant 1
push constant 512
call Math.multiply 2
add
call Sprites.drawLogoC2 1
pop temp 0
push constant 59
push constant 1
push constant 512
call Math.multiply 2
add
call Sprites.drawLogoD2 1
pop temp 0
push constant 60
push constant 1
push constant 512
call Math.multiply 2
add
call Sprites.drawLogoE2 1
pop temp 0
push constant 61
push constant 1
push constant 512
call Math.multiply 2
add
call Sprites.drawLogoF2 1
pop temp 0
push constant 62
push constant 1
push constant 512
call Math.multiply 2
add
call Sprites.drawLogoG2 1
pop temp 0
push constant 56
push constant 2
push constant 512
call Math.multiply 2
add
call Sprites.drawLogoA3 1
pop temp 0
push constant 57
push constant 2
push constant 512
call Math.multiply 2
add
call Sprites.drawLogoB3 1
pop temp 0
push constant 58
push constant 2
push constant 512
call Math.multiply 2
add
call Sprites.drawLogoC3 1
pop temp 0
push constant 59
push constant 2
push constant 512
call Math.multiply 2
add
call Sprites.drawLogoD3 1
pop temp 0
push constant 60
push constant 2
push constant 512
call Math.multiply 2
add
call Sprites.drawLogoE3 1
pop temp 0
push constant 61
push constant 2
push constant 512
call Math.multiply 2
add
call Sprites.drawLogoF3 1
pop temp 0
push constant 62
push constant 2
push constant 512
call Math.multiply 2
add
call Sprites.drawLogoG3 1
pop temp 0
push constant 0
return
function View.drawBoard 9
push argument 0
pop pointer 0
push this 0
call Model.getBoard 1
pop local 0
push constant 0
pop local 1
push constant 0
pop local 2
push local 0
call Board.getWidth 1
pop local 3
push local 0
call Board.getHeight 1
pop local 4
push constant 22
push local 3
sub
push constant 2
call Math.divide 2
push constant 1
add
pop local 5
push constant 14
push local 4
sub
push constant 2
call Math.divide 2
push constant 1
add
pop local 6
push local 0
call Board.getPlayer 1
pop local 7
label WHILE_EXP0
push local 2
push local 4
lt
not
if-goto WHILE_END0
label WHILE_EXP1
push local 1
push local 3
lt
not
if-goto WHILE_END1
push local 0
push local 1
push local 2
call Board.getCell 3
pop local 8
push local 7
call Player.getX 1
push local 1
eq
push local 7
call Player.getY 1
push local 2
eq
and
if-goto IF_TRUE0
goto IF_FALSE0
label IF_TRUE0
push local 8
call Cell.getTypeOf 1
call CellType.CellTypeGoal 0
eq
if-goto IF_TRUE1
goto IF_FALSE1
label IF_TRUE1
push local 5
push local 1
add
push local 6
push local 2
add
push constant 512
call Math.multiply 2
add
call Sprites.drawGoalPlayer 1
pop temp 0
goto IF_END1
label IF_FALSE1
push local 5
push local 1
add
push local 6
push local 2
add
push constant 512
call Math.multiply 2
add
call Sprites.drawPlayer 1
pop temp 0
label IF_END1
goto IF_END0
label IF_FALSE0
push local 8
call Cell.getTypeOf 1
call CellType.CellTypeNone 0
eq
if-goto IF_TRUE2
goto IF_FALSE2
label IF_TRUE2
push local 8
call Cell.getHasBox 1
if-goto IF_TRUE3
goto IF_FALSE3
label IF_TRUE3
push local 5
push local 1
add
push local 6
push local 2
add
push constant 512
call Math.multiply 2
add
call Sprites.drawBox 1
pop temp 0
goto IF_END3
label IF_FALSE3
push local 5
push local 1
add
push local 6
push local 2
add
push constant 512
call Math.multiply 2
add
call Sprites.drawClear 1
pop temp 0
label IF_END3
label IF_FALSE2
push local 8
call Cell.getTypeOf 1
call CellType.CellTypeGoal 0
eq
if-goto IF_TRUE4
goto IF_FALSE4
label IF_TRUE4
push local 8
call Cell.getHasBox 1
if-goto IF_TRUE5
goto IF_FALSE5
label IF_TRUE5
push local 5
push local 1
add
push local 6
push local 2
add
push constant 512
call Math.multiply 2
add
call Sprites.drawGoalBox 1
pop temp 0
goto IF_END5
label IF_FALSE5
push local 5
push local 1
add
push local 6
push local 2
add
push constant 512
call Math.multiply 2
add
call Sprites.drawGoal 1
pop temp 0
label IF_END5
label IF_FALSE4
push local 8
call Cell.getTypeOf 1
call CellType.CellTypeWall 0
eq
if-goto IF_TRUE6
goto IF_FALSE6
label IF_TRUE6
push local 5
push local 1
add
push local 6
push local 2
add
push constant 512
call Math.multiply 2
add
call Sprites.drawWall 1
pop temp 0
label IF_FALSE6
label IF_END0
push local 1
push constant 1
add
pop local 1
goto WHILE_EXP1
label WHILE_END1
push local 2
push constant 1
add
pop local 2
push constant 0
pop local 1
goto WHILE_EXP0
label WHILE_END0
push constant 0
return
function View.drawLevelInfo 6
push argument 0
pop pointer 0
push this 0
call Model.getLevelManager 1
pop local 0
push constant 14
call String.new 1
push constant 76
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 118
call String.appendChar 2
push constant 101
call String.appendChar 2
push constant 108
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 88
call String.appendChar 2
push constant 88
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 111
call String.appendChar 2
push constant 102
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 89
call String.appendChar 2
push constant 89
call String.appendChar 2
pop local 1
push local 0
call LevelManager.getCurrentLevelNumber 1
pop local 2
push constant 2
call Array.new 1
pop local 3
push constant 0
push local 3
add
push local 2
push constant 10
call Math.divide 2
pop temp 0
pop pointer 1
push temp 0
pop that 0
push constant 1
push local 3
add
push local 2
push constant 10
call Utils.mod 2
pop temp 0
pop pointer 1
push temp 0
pop that 0
push local 1
push constant 6
push constant 48
push constant 0
push local 3
add
pop pointer 1
push that 0
add
call String.setCharAt 3
pop temp 0
push local 1
push constant 7
push constant 48
push constant 1
push local 3
add
pop pointer 1
push that 0
add
call String.setCharAt 3
pop temp 0
push local 0
call LevelManager.getFinalLevelNumber 1
pop local 4
push constant 2
call Array.new 1
pop local 5
push constant 0
push local 5
add
push local 4
push constant 10
call Math.divide 2
pop temp 0
pop pointer 1
push temp 0
pop that 0
push constant 1
push local 5
add
push local 4
push constant 10
call Utils.mod 2
pop temp 0
pop pointer 1
push temp 0
pop that 0
push local 1
push constant 12
push constant 48
push constant 0
push local 5
add
pop pointer 1
push that 0
add
call String.setCharAt 3
pop temp 0
push local 1
push constant 13
push constant 48
push constant 1
push local 5
add
pop pointer 1
push that 0
add
call String.setCharAt 3
pop temp 0
push constant 7
push constant 48
call Output.moveCursor 2
pop temp 0
push local 1
call Output.printString 1
pop temp 0
push local 1
call String.dispose 1
pop temp 0
push local 3
call Array.dispose 1
pop temp 0
push local 5
call Array.dispose 1
pop temp 0
push constant 0
return
function View.drawGameCompleteBorder 2
push argument 0
pop pointer 0
push constant 1
pop local 0
push constant 1
pop local 1
label WHILE_EXP0
push local 1
push constant 15
lt
not
if-goto WHILE_END0
label WHILE_EXP1
push local 0
push constant 23
lt
not
if-goto WHILE_END1
push local 0
push constant 1
eq
push local 0
push constant 22
eq
or
push local 1
push constant 1
eq
or
push local 1
push constant 14
eq
or
if-goto IF_TRUE0
goto IF_FALSE0
label IF_TRUE0
push argument 1
if-goto IF_TRUE1
goto IF_FALSE1
label IF_TRUE1
push local 0
push local 1
push constant 512
call Math.multiply 2
add
call Sprites.drawPlayer 1
pop temp 0
goto IF_END1
label IF_FALSE1
push local 0
push local 1
push constant 512
call Math.multiply 2
add
call Sprites.drawClear 1
pop temp 0
label IF_END1
label IF_FALSE0
push local 0
push constant 1
add
pop local 0
goto WHILE_EXP1
label WHILE_END1
push local 1
push constant 1
add
pop local 1
push constant 1
pop local 0
goto WHILE_EXP0
label WHILE_END0
push constant 0
return
function View.dispose 0
push argument 0
pop pointer 0
push pointer 0
call Memory.deAlloc 1
pop temp 0
push constant 0
return
`

const Square_Main_VM string = `function Main.main 1
call SquareGame.new 0
pop local 0
push local 0
call SquareGame.run 1
pop temp 0
push local 0
call SquareGame.dispose 1
pop temp 0
push constant 0
return
function Main.more 4
push constant 0
if-goto IF_TRUE0
goto IF_FALSE0
label IF_TRUE0
push constant 15
call String.new 1
push constant 115
call String.appendChar 2
push constant 116
call String.appendChar 2
push constant 114
call String.appendChar 2
push constant 105
call String.appendChar 2
push constant 110
call String.appendChar 2
push constant 103
call String.appendChar 2
push constant 32
call String.appendChar 2
push constant 99
call String.appendChar 2
push constant 111
call String.appendChar 2
push constant 110
call String.appendChar 2
push constant 115
call String.appendChar 2
push constant 116
call String.appendChar 2
push constant 97
call String.appendChar 2
push constant 110
call String.appendChar 2
push constant 116
call String.appendChar 2
pop local 2
push constant 0
pop local 2
push constant 1
push local 3
add
push constant 2
push local 3
add
pop pointer 1
push that 0
pop temp 0
pop pointer 1
push temp 0
pop that 0
goto IF_END0
label IF_FALSE0
push local 0
push local 1
neg
call Math.multiply 2
pop local 0
push local 1
push constant 2
neg
call Math.divide 2
pop local 1
push local 0
push local 1
or
pop local 0
label IF_END0
push constant 0
return
`

const Square_Square_VM string = `function Square.new 0
push constant 3
call Memory.alloc 1
pop pointer 0
push argument 0
pop this 0
push argument 1
pop this 1
push argument 2
pop this 2
push pointer 0
call Square.draw 1
pop temp 0
push pointer 0
return
function Square.dispose 0
push argument 0
pop pointer 0
push pointer 0
call Memory.deAlloc 1
pop temp 0
push constant 0
return
function Square.draw 0
push argument 0
pop pointer 0
push constant 0
not
call Screen.setColor 1
pop temp 0
push this 0
push this 1
push this 0
push this 2
add
push this 1
push this 2
add
call Screen.drawRectangle 4
pop temp 0
push constant 0
return
function Square.erase 0
push argument 0
pop pointer 0
push constant 0
call Screen.setColor 1
pop temp 0
push this 0
push this 1
push this 0
push this 2
add
push this 1
push this 2
add
call Screen.drawRectangle 4
pop temp 0
push constant 0
return
function Square.incSize 0
push argument 0
pop pointer 0
push this 1
push this 2
add
push constant 254
lt
push this 0
push this 2
add
push constant 510
lt
and
if-goto IF_TRUE0
goto IF_FALSE0
label IF_TRUE0
push pointer 0
call Square.erase 1
pop temp 0
push this 2
push constant 2
add
pop this 2
push pointer 0
call Square.draw 1
pop temp 0
label IF_FALSE0
push constant 0
return
function Square.decSize 0
push argument 0
pop pointer 0
push this 2
push constant 2
gt
if-goto IF_TRUE0
goto IF_FALSE0
label IF_TRUE0
push pointer 0
call Square.erase 1
pop temp 0
push this 2
push constant 2
sub
pop this 2
push pointer 0
call Square.draw 1
pop temp 0
label IF_FALSE0
push constant 0
return
function Square.moveUp 0
push argument 0
pop pointer 0
push this 1
push constant 1
gt
if-goto IF_TRUE0
goto IF_FALSE0
label IF_TRUE0
push constant 0
call Screen.setColor 1
pop temp 0
push this 0
push this 1
push this 2
add
push constant 1
sub
push this 0
push this 2
add
push this 1
push this 2
add
call Screen.drawRectangle 4
pop temp 0
push this 1
push constant 2
sub
pop this 1
push constant 0
not
call Screen.setColor 1
pop temp 0
push this 0
push this 1
push this 0
push this 2
add
push this 1
push constant 1
add
call Screen.drawRectangle 4
pop temp 0
label IF_FALSE0
push constant 0
return
function Square.moveDown 0
push argument 0
pop pointer 0
push this 1
push this 2
add
push constant 254
lt
if-goto IF_TRUE0
goto IF_FALSE0
label IF_TRUE0
push constant 0
call Screen.setColor 1
pop temp 0
push this 0
push this 1
push this 0
push this 2
add
push this 1
push constant 1
add
call Screen.drawRectangle 4
pop temp 0
push this 1
push constant 2
add
pop this 1
push constant 0
not
call Screen.setColor 1
pop temp 0
push this 0
push this 1
push this 2
add
push constant 1
sub
push this 0
push this 2
add
push this 1
push this 2
add
call Screen.drawRectangle 4
pop temp 0
label IF_FALSE0
push constant 0
return
function Square.moveLeft 0
push argument 0
pop pointer 0
push this 0
push constant 1
gt
if-goto IF_TRUE0
goto IF_FALSE0
label IF_TRUE0
push constant 0
call Screen.setColor 1
pop temp 0
push this 0
push this 2
add
push constant 1
sub
push this 1
push this 0
push this 2
add
push this 1
push this 2
add
call Screen.drawRectangle 4
pop temp 0
push this 0
push constant 2
sub
pop this 0
push constant 0
not
call Screen.setColor 1
pop temp 0
push this 0
push this 1
push this 0
push constant 1
add
push this 1
push this 2
add
call Screen.drawRectangle 4
pop temp 0
label IF_FALSE0
push constant 0
return
function Square.moveRight 0
push argument 0
pop pointer 0
push this 0
push this 2
add
push constant 510
lt
if-goto IF_TRUE0
goto IF_FALSE0
label IF_TRUE0
push constant 0
call Screen.setColor 1
pop temp 0
push this 0
push this 1
push this 0
push constant 1
add
push this 1
push this 2
add
call Screen.drawRectangle 4
pop temp 0
push this 0
push constant 2
add
pop this 0
push constant 0
not
call Screen.setColor 1
pop temp 0
push this 0
push this 2
add
push constant 1
sub
push this 1
push this 0
push this 2
add
push this 1
push this 2
add
call Screen.drawRectangle 4
pop temp 0
label IF_FALSE0
push constant 0
return
`

const Square_SquareGame_VM string = `function SquareGame.new 0
push constant 2
call Memory.alloc 1
pop pointer 0
push constant 0
push constant 0
push constant 30
call Square.new 3
pop this 0
push constant 0
pop this 1
push pointer 0
return
function SquareGame.dispose 0
push argument 0
pop pointer 0
push this 0
call Square.dispose 1
pop temp 0
push pointer 0
call Memory.deAlloc 1
pop temp 0
push constant 0
return
function SquareGame.moveSquare 0
push argument 0
pop pointer 0
push this 1
push constant 1
eq
if-goto IF_TRUE0
goto IF_FALSE0
label IF_TRUE0
push this 0
call Square.moveUp 1
pop temp 0
label IF_FALSE0
push this 1
push constant 2
eq
if-goto IF_TRUE1
goto IF_FALSE1
label IF_TRUE1
push this 0
call Square.moveDown 1
pop temp 0
label IF_FALSE1
push this 1
push constant 3
eq
if-goto IF_TRUE2
goto IF_FALSE2
label IF_TRUE2
push this 0
call Square.moveLeft 1
pop temp 0
label IF_FALSE2
push this 1
push constant 4
eq
if-goto IF_TRUE3
goto IF_FALSE3
label IF_TRUE3
push this 0
call Square.moveRight 1
pop temp 0
label IF_FALSE3
push constant 5
call Sys.wait 1
pop temp 0
push constant 0
return
function SquareGame.run 2
push argument 0
pop pointer 0
push constant 0
pop local 1
label WHILE_EXP0
push local 1
not
not
if-goto WHILE_END0
label WHILE_EXP1
push local 0
push constant 0
eq
not
if-goto WHILE_END1
call Keyboard.keyPressed 0
pop local 0
push pointer 0
call SquareGame.moveSquare 1
pop temp 0
goto WHILE_EXP1
label WHILE_END1
push local 0
push constant 81
eq
if-goto IF_TRUE0
goto IF_FALSE0
label IF_TRUE0
push constant 0
not
pop local 1
label IF_FALSE0
push local 0
push constant 90
eq
if-goto IF_TRUE1
goto IF_FALSE1
label IF_TRUE1
push this 0
call Square.decSize 1
pop temp 0
label IF_FALSE1
push local 0
push constant 88
eq
if-goto IF_TRUE2
goto IF_FALSE2
label IF_TRUE2
push this 0
call Square.incSize 1
pop temp 0
label IF_FALSE2
push local 0
push constant 131
eq
if-goto IF_TRUE3
goto IF_FALSE3
label IF_TRUE3
push constant 1
pop this 1
label IF_FALSE3
push local 0
push constant 133
eq
if-goto IF_TRUE4
goto IF_FALSE4
label IF_TRUE4
push constant 2
pop this 1
label IF_FALSE4
push local 0
push constant 130
eq
if-goto IF_TRUE5
goto IF_FALSE5
label IF_TRUE5
push constant 3
pop this 1
label IF_FALSE5
push local 0
push constant 132
eq
if-goto IF_TRUE6
goto IF_FALSE6
label IF_TRUE6
push constant 4
pop this 1
label IF_FALSE6
label WHILE_EXP2
push local 0
push constant 0
eq
not
not
if-goto WHILE_END2
call Keyboard.keyPressed 0
pop local 0
push pointer 0
call SquareGame.moveSquare 1
pop temp 0
goto WHILE_EXP2
label WHILE_END2
goto WHILE_EXP0
label WHILE_END0
push constant 0
return
`
