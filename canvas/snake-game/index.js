
const  canvas = document.querySelector('.canvas')
const ctx = canvas.getContext('2d')

const scale = 10
const rows = canvas.height / scale
const columns = canvas.width / scale

var snake
var fruit

function init() {

    snake = new Snake()
    fruit = new Fruit()
    fruit.pickLocation()
    console.log(fruit)


    window.setInterval(function() {
        ctx.clearRect(0, 0, canvas.width, canvas.height)
        fruit.draw();
        snake.update();
        snake.draw();

        if (fruit.eated(snake)) {
            //console.log("eating")
            fruit.pickLocation();
        }

    }, 50)
}
init()


function Fruit() {
    this.x
    this.y

    this.pickLocation = function() {
        this.x = (Math.floor(Math.random() * rows-1)+1)*scale 
        this.y = (Math.floor(Math.random() * columns-1)+1)*scale
    }

    this.draw = function() {
        ctx.fillStyle = "yellow"

        ctx.fillRect(this.x, this.y, scale, scale)
    }

    this.eated = function(snake) {
        if (this.x == snake.x && this.y == snake.y ) {
            snake.total ++
            return true
        }
    }
}

// Snake
function Snake() {
    this.x = 0
    this.y = 0
    this.xSpeed = scale * 1;
    this.ySpeed = 0;
    this.total = 0;
    this.tail = [];

    this.draw = function() {

        ctx.fillStyle = '#FFFFFF'

        for (let i = 0; i < this.tail.length; i++) {
            ctx.fillRect(this.tail[i].x, this.tail[i].y, scale, scale)
        }
        ctx.fillRect(this.x, this.y, scale, scale)
    }

    this.update = function() {

        // drow tail
        for (let i = 0; i < this.tail.length-1; i++) {
            this.tail[i] = this.tail[i+1]
        }
        
        this.tail[this.total - 1] = {x: this.x, y: this.y}

        // 
        this.x += this.xSpeed;
        this.y += this.ySpeed;

        if (this.x >= canvas.width) {
            this.x = 0;
        }
        if (this.y >= canvas.height) {
            this.y = 0;
        }
        if (this.x < 0) {
            this.x = canvas.width;
        }
        if (this.y < 0) {
            this.y = canvas.height;
        }
    }

    this.changeDirection = function(direction) {
        switch(direction) {
            case 'Up':
                this.xSpeed = 0;
                this.ySpeed = -scale * 1;
                break
            case 'Down':
                this.xSpeed = 0;
                this.ySpeed = scale * 1;
                break
            case 'Left':
                this.xSpeed = -scale * 1;
                this.ySpeed = 0;
                break
            case 'Right':
                this.xSpeed = scale * 1;
                this.ySpeed = 0;
                break
        }
    }

}



window.addEventListener('keydown', function(e) {
    const direction = e.key.replace('Arrow', '');
    snake.changeDirection(direction)
    console.log(direction)
})
