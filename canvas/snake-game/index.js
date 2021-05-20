const  canvas = document.querySelector('.canvas')
const ctx = canvas.getContext('2d')
const scale = 10
const rows = canvas.height / scale
const columns = canvas.width / scale

var snake

(function init() {
    snake = new Snake()
    window.setInterval(function() {
        ctx.clearRect(0, 0, canvas.width, canvas.height)
        snake.update();
        snake.draw();
    }, 50)
}())

window.addEventListener('keydown', function(e) {
    const direction = e.key.replace('Arrow', '');
    snake.changeDirection(direction)
    console.log(direction)
})

function Snake() {
    this.x = 0
    this.y = 0
    this.xSpeed = scale * 1;
    this.ySpeed = 0;

    this.draw = function() {
        ctx.fillStyle = '#FFFFFF'
        ctx.fillRect(this.x, this.y, scale, scale)
    }

    this.update = function() {
        this.x += this.xSpeed;
        this.y += this.ySpeed;

        if (this.x > canvas.width) {
            this.x = 0;
        }
        if (this.y > canvas.height) {
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

