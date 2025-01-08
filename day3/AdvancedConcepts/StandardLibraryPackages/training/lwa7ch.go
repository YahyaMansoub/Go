package main

import (
	"fmt"
	"net/http"
)

// Handler for the home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to the Home Page</h1><p>This is the home page of the website.</p>")
}
func gameOfLifeHandler(w http.ResponseWriter, r *http.Request) {
	// HTML content for Conway's Game of Life
	const htmlContent = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Conway's Game of Life</title>
    <style>
        body {
            font-family: sans-serif;
            margin: 0;
            padding: 0;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            background-color: #f0f0f0;
        }
        canvas {
            border: 1px solid #000;
        }
        button {
            padding: 10px 20px;
            margin: 10px;
        }
    </style>
</head>
<body>
    <div>
        <canvas id="gameCanvas" width="500" height="500"></canvas><br>
        <button id="startBtn">Start</button>
        <button id="resetBtn">Reset</button>
    </div>
    <script>
        const canvas = document.getElementById('gameCanvas');
        const ctx = canvas.getContext('2d');
        const gridSize = 50;
        const cellSize = canvas.width / gridSize;
        let grid = [];
        let isRunning = false;
        let interval;

        // Initialize grid with random cells
        function initGrid() {
            grid = [];
            for (let i = 0; i < gridSize; i++) {
                grid[i] = [];
                for (let j = 0; j < gridSize; j++) {
                    grid[i][j] = Math.random() < 0.5 ? 0 : 1; // Random start
                }
            }
        }

        // Draw the grid to the canvas
        function drawGrid() {
            for (let i = 0; i < gridSize; i++) {
                for (let j = 0; j < gridSize; j++) {
                    ctx.fillStyle = grid[i][j] ? 'black' : 'white';
                    ctx.fillRect(j * cellSize, i * cellSize, cellSize, cellSize);
                    ctx.strokeStyle = 'gray';
                    ctx.strokeRect(j * cellSize, i * cellSize, cellSize, cellSize);
                }
            }
        }

        // Update grid based on Conway's Game of Life rules
        function updateGrid() {
            const newGrid = grid.map((arr, i) => arr.slice());
            for (let i = 0; i < gridSize; i++) {
                for (let j = 0; j < gridSize; j++) {
                    const aliveNeighbors = countAliveNeighbors(i, j);
                    if (grid[i][j] === 1) {
                        newGrid[i][j] = aliveNeighbors === 2 || aliveNeighbors === 3 ? 1 : 0;
                    } else {
                        newGrid[i][j] = aliveNeighbors === 3 ? 1 : 0;
                    }
                }
            }
            grid = newGrid;
            drawGrid();
        }

        // Count the number of alive neighbors of a cell
        function countAliveNeighbors(x, y) {
            let count = 0;
            for (let i = -1; i <= 1; i++) {
                for (let j = -1; j <= 1; j++) {
                    if (i === 0 && j === 0) continue;
                    const nx = x + i;
                    const ny = y + j;
                    if (nx >= 0 && ny >= 0 && nx < gridSize && ny < gridSize) {
                        count += grid[nx][ny];
                    }
                }
            }
            return count;
        }

        // Toggle the state of a cell
        function toggleCell(x, y) {
            grid[x][y] = grid[x][y] === 1 ? 0 : 1;
            drawGrid();
        }

        // Event listener for clicking on the canvas to toggle cells
        canvas.addEventListener('click', (e) => {
            const x = Math.floor(e.offsetY / cellSize);
            const y = Math.floor(e.offsetX / cellSize);
            toggleCell(x, y);
        });

        // Start or pause the simulation when the Start button is clicked
        document.getElementById('startBtn').addEventListener('click', () => {
            if (isRunning) {
                clearInterval(interval);
                document.getElementById('startBtn').textContent = 'Start';
            } else {
                interval = setInterval(updateGrid, 100);
                document.getElementById('startBtn').textContent = 'Pause';
            }
            isRunning = !isRunning;
        });

        // Reset the grid when the Reset button is clicked
        document.getElementById('resetBtn').addEventListener('click', () => {
            clearInterval(interval);
            isRunning = false;
            document.getElementById('startBtn').textContent = 'Start';
            initGrid();
            drawGrid();
        });

        // Initialize and draw the grid on page load
        initGrid();
        drawGrid();
    </script>
</body>
</html>`

	// Set the Content-Type to "text/html" to let the browser know this is HTML
	w.Header().Set("Content-Type", "text/html")
	// Write the HTML content to the response
	w.Write([]byte(htmlContent))
}
// Handler for the about page
func NicerHandler(w http.ResponseWriter, r *http.Request) {
	const htmlnicer=`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Conway's Game of Life</title>
    <style>
        /* Global Styles */
        body {
            font-family: 'Arial', sans-serif;
            margin: 0;
            padding: 0;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            background-color: #282c34;
            color: #f4f4f4;
            flex-direction: column;
        }

        h1 {
            color: #fff;
            margin-bottom: 20px;
            font-size: 2em;
        }

        /* Canvas Styles */
        canvas {
            border: 2px solid #ddd;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
            transition: box-shadow 0.3s ease;
        }

        canvas:hover {
            box-shadow: 0 0 20px rgba(0, 0, 0, 0.5);
        }

        /* Button Styles */
        button {
            padding: 10px 20px;
            font-size: 1.1em;
            margin: 10px;
            background-color: #4CAF50;
            color: white;
            border: none;
            border-radius: 5px;
            cursor: pointer;
            transition: background-color 0.3s ease, transform 0.3s ease;
        }

        button:hover {
            background-color: #45a049;
            transform: scale(1.1);
        }

        button:active {
            transform: scale(0.95);
        }

        button:focus {
            outline: none;
        }

        /* Container to hold canvas and buttons */
        .controls {
            text-align: center;
        }
    </style>
</head>
<body>

    <h1>Conway's Game of Life</h1>

    <div class="controls">
        <canvas id="gameCanvas" width="500" height="500"></canvas><br>
        <button id="startBtn">Start</button>
        <button id="resetBtn">Reset</button>
    </div>

    <script>
        const canvas = document.getElementById('gameCanvas');
        const ctx = canvas.getContext('2d');
        const gridSize = 50;
        const cellSize = canvas.width / gridSize;
        let grid = [];
        let isRunning = false;
        let interval;

        // Initialize grid with random cells
        function initGrid() {
            grid = [];
            for (let i = 0; i < gridSize; i++) {
                grid[i] = [];
                for (let j = 0; j < gridSize; j++) {
                    grid[i][j] = Math.random() < 0.5 ? 0 : 1; // Random start
                }
            }
        }

        // Draw the grid to the canvas
        function drawGrid() {
            for (let i = 0; i < gridSize; i++) {
                for (let j = 0; j < gridSize; j++) {
                    ctx.fillStyle = grid[i][j] ? '#333' : '#fff';
                    ctx.fillRect(j * cellSize, i * cellSize, cellSize, cellSize);
                    ctx.strokeStyle = '#ddd';
                    ctx.strokeRect(j * cellSize, i * cellSize, cellSize, cellSize);
                }
            }
        }

        // Update grid based on Conway's Game of Life rules
        function updateGrid() {
            const newGrid = grid.map((arr, i) => arr.slice());
            for (let i = 0; i < gridSize; i++) {
                for (let j = 0; j < gridSize; j++) {
                    const aliveNeighbors = countAliveNeighbors(i, j);
                    if (grid[i][j] === 1) {
                        newGrid[i][j] = aliveNeighbors === 2 || aliveNeighbors === 3 ? 1 : 0;
                    } else {
                        newGrid[i][j] = aliveNeighbors === 3 ? 1 : 0;
                    }
                }
            }
            grid = newGrid;
            drawGrid();
        }

        // Count the number of alive neighbors of a cell
        function countAliveNeighbors(x, y) {
            let count = 0;
            for (let i = -1; i <= 1; i++) {
                for (let j = -1; j <= 1; j++) {
                    if (i === 0 && j === 0) continue;
                    const nx = x + i;
                    const ny = y + j;
                    if (nx >= 0 && ny >= 0 && nx < gridSize && ny < gridSize) {
                        count += grid[nx][ny];
                    }
                }
            }
            return count;
        }

        // Toggle the state of a cell
        function toggleCell(x, y) {
            grid[x][y] = grid[x][y] === 1 ? 0 : 1;
            drawGrid();
        }

        // Event listener for clicking on the canvas to toggle cells
        canvas.addEventListener('click', (e) => {
            const x = Math.floor(e.offsetY / cellSize);
            const y = Math.floor(e.offsetX / cellSize);
            toggleCell(x, y);
        });

        // Start or pause the simulation when the Start button is clicked
        document.getElementById('startBtn').addEventListener('click', () => {
            if (isRunning) {
                clearInterval(interval);
                document.getElementById('startBtn').textContent = 'Start';
            } else {
                interval = setInterval(updateGrid, 100);
                document.getElementById('startBtn').textContent = 'Pause';
            }
            isRunning = !isRunning;
        });

        // Reset the grid when the Reset button is clicked
        document.getElementById('resetBtn').addEventListener('click', () => {
            clearInterval(interval);
            isRunning = false;
            document.getElementById('startBtn').textContent = 'Start';
            initGrid();
            drawGrid();
        });

        // Initialize and draw the grid on page load
        initGrid();
        drawGrid();
    </script>
</body>
</html>
`
        w.Header().Set("Content-Type", "text/html")
	    w.Write([]byte(htmlnicer))
}

// Handler for the contact page
func SnakeHandler(w http.ResponseWriter, r *http.Request) {
	const htmlsnake=`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Snake Game</title>
    <style>
        /* Global Styles */
        body {
            font-family: 'Arial', sans-serif;
            margin: 0;
            padding: 0;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            background-color: #282c34;
            color: #f4f4f4;
            flex-direction: column;
        }

        h1 {
            color: #fff;
            margin-bottom: 20px;
            font-size: 2em;
        }

        /* Canvas Styles */
        canvas {
            border: 2px solid #ddd;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
            transition: box-shadow 0.3s ease;
        }

        canvas:hover {
            box-shadow: 0 0 20px rgba(0, 0, 0, 0.5);
        }

        /* Button Styles */
        button {
            padding: 10px 20px;
            font-size: 1.1em;
            margin: 10px;
            background-color: #4CAF50;
            color: white;
            border: none;
            border-radius: 5px;
            cursor: pointer;
            transition: background-color 0.3s ease, transform 0.3s ease;
        }

        button:hover {
            background-color: #45a049;
            transform: scale(1.1);
        }

        button:active {
            transform: scale(0.95);
        }

        button:focus {
            outline: none;
        }

        /* Container to hold canvas and buttons */
        .controls {
            text-align: center;
        }

        .score {
            font-size: 1.5em;
            margin-top: 20px;
        }
    </style>
</head>
<body>

    <h1>Snake Game</h1>

    <div class="controls">
        <canvas id="gameCanvas" width="500" height="500"></canvas><br>
        <button id="startBtn">Start</button>
        <button id="resetBtn">Reset</button>
    </div>
    <div class="score" id="scoreDisplay">Score: 0</div>

    <script>
        const canvas = document.getElementById('gameCanvas');
        const ctx = canvas.getContext('2d');
        const gridSize = 20;
        const canvasSize = canvas.width;
        const cellSize = canvasSize / gridSize;
        let snake = [{ x: 10, y: 10 }];
        let direction = 'RIGHT';
        let food = { x: 5, y: 5 };
        let score = 0;
        let gameInterval;
        let isRunning = false;

        // Draw the game board
        function drawBoard() {
            ctx.clearRect(0, 0, canvas.width, canvas.height);
            // Draw the snake
            snake.forEach((segment, index) => {
                ctx.fillStyle = index === 0 ? 'green' : 'lightgreen';  // Head of snake is green
                ctx.fillRect(segment.x * cellSize, segment.y * cellSize, cellSize, cellSize);
            });

            // Draw the food
            ctx.fillStyle = 'red';
            ctx.fillRect(food.x * cellSize, food.y * cellSize, cellSize, cellSize);

            // Display the score
            document.getElementById('scoreDisplay').textContent = 'Score: ${score}';
        }

        // Update the game state
        function updateGame() {
            const head = { ...snake[0] };

            switch (direction) {
                case 'UP':
                    head.y -= 1;
                    break;
                case 'DOWN':
                    head.y += 1;
                    break;
                case 'LEFT':
                    head.x -= 1;
                    break;
                case 'RIGHT':
                    head.x += 1;
                    break;
            }

            // Check if the snake hits the wall or itself
            if (head.x < 0 || head.x >= gridSize || head.y < 0 || head.y >= gridSize || isCollidingWithSelf(head)) {
                endGame();
                return;
            }

            // Check if the snake eats food
            if (head.x === food.x && head.y === food.y) {
                score += 10;
                food = generateRandomFood();
            } else {
                snake.pop();  // Remove the last segment of the snake
            }

            snake.unshift(head);  // Add the new head to the snake

            drawBoard();
        }

        // Check if the snake collides with itself
        function isCollidingWithSelf(head) {
            for (let i = 1; i < snake.length; i++) {
                if (snake[i].x === head.x && snake[i].y === head.y) {
                    return true;
                }
            }
            return false;
        }

        // Generate random food coordinates
        function generateRandomFood() {
            let foodX, foodY;
            do {
                foodX = Math.floor(Math.random() * gridSize);
                foodY = Math.floor(Math.random() * gridSize);
            } while (isFoodOnSnake(foodX, foodY));

            return { x: foodX, y: foodY };
        }

        // Check if food is on the snake's body
        function isFoodOnSnake(x, y) {
            for (let i = 0; i < snake.length; i++) {
                if (snake[i].x === x && snake[i].y === y) {
                    return true;
                }
            }
            return false;
        }

        // Control the snake with arrow keys
        document.addEventListener('keydown', (e) => {
            if (e.key === 'ArrowUp' && direction !== 'DOWN') direction = 'UP';
            if (e.key === 'ArrowDown' && direction !== 'UP') direction = 'DOWN';
            if (e.key === 'ArrowLeft' && direction !== 'RIGHT') direction = 'LEFT';
            if (e.key === 'ArrowRight' && direction !== 'LEFT') direction = 'RIGHT';
        });

        // Start the game
        function startGame() {
            snake = [{ x: 10, y: 10 }];
            direction = 'RIGHT';
            score = 0;
            food = generateRandomFood();
            isRunning = true;
            gameInterval = setInterval(updateGame, 100);
            document.getElementById('startBtn').textContent = 'Pause';
        }

        // Pause the game
        function pauseGame() {
            clearInterval(gameInterval);
            isRunning = false;
            document.getElementById('startBtn').textContent = 'Start';
        }

        // End the game
        function endGame() {
            clearInterval(gameInterval);
            isRunning = false;
            alert('Game Over! Final Score: ${score}');
            document.getElementById('startBtn').textContent = 'Start';
        }

        // Reset the game
        function resetGame() {
            clearInterval(gameInterval);
            snake = [{ x: 10, y: 10 }];
            direction = 'RIGHT';
            score = 0;
            food = generateRandomFood();
            drawBoard();
        }

        // Event listener for Start/Pause button
        document.getElementById('startBtn').addEventListener('click', () => {
            if (isRunning) {
                pauseGame();
            } else {
                startGame();
            }
        });

        // Event listener for Reset button
        document.getElementById('resetBtn').addEventListener('click', resetGame);

        // Initialize the game board
        drawBoard();
    </script>
</body>
</html>
`
        w.Header().Set("Content-Type", "text/html")
	    w.Write([]byte(htmlsnake))
}

// Handler for the services page
func servicesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Our Services</h1><p>Discover the services we offer.</p>")
}

func main() {
	// Register the handlers for the pages
	http.HandleFunc("/", homeHandler)      // Home page
	http.HandleFunc("/Nicers", NicerHandler) // About page
	http.HandleFunc("/Snake", SnakeHandler) // Contact page
	http.HandleFunc("/services", servicesHandler) // Services page
    http.HandleFunc("/GL",gameOfLifeHandler)
	// Start the web server on port 8080
	fmt.Println("Starting the web server on port 8080...")
	http.ListenAndServe(":8080", nil)
}
