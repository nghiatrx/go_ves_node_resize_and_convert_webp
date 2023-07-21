const sharp = require('sharp');

const input = [
    "./input/input-0.jpg",
    "./input/input-1.jpg",
    "./input/input-2.jpg",
    "./input/input-3.jpg",
    "./input/input-4.jpg",
    "./input/input-5.jpg",
]

async function run() {
    const len = input.length;

    const t1 = Date.now();

    for (let i = 0; i < len; i++) {
        await sharp(input[i])
            .resize(1000)
            .webp({ quality: 80 })
            .toFile(`./output-${i}.webp`);
    }

    const t2 = Date.now();

    console.log("time: ", t2 - t1)

}

run()