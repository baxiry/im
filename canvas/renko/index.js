
console.log("hello canvas")
const canvas = document.getElementById('canvas');
canvas.width = window.innerWidth - 20;
canvas.height  = window.innerHeight - 100;

const ctx = canvas.getContext('2d');
ctx.fillStyle = 'blue';

for (let i = 0; i<= 200; i+=20) {
    ctx.fillRect(i/2, i+i/2, 10, 30)// .fillStyle = 'blue';
    i++;
}

ctx.fillStyle = 'green';
for (let i = 210; i<=400; i+=20) {
    ctx.fillRect(i/2, i+i/2, 10, 30)// .fillStyle = 'blue';
    i++;
}
