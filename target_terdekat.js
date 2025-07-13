function targetTerdekat(arr) {
    if (!arr.includes("x")) return 0;

    const temp_result = [];
    let count = 0;
    let currentPos;

    for (const [index, item] of arr.entries()) {
        if (currentPos === item) count = 0;

        if (["o", "x"].includes(item)) {
            if (currentPos && item !== currentPos) {
                temp_result.push(count);
                count = 0;
            }
            currentPos = item;
        }

        if (currentPos) count++;
    }

    return Math.min(...temp_result);
}

console.log(targetTerdekat([" ", " ", "o", " ", " ", "x", " ", "x"])); // 3
console.log(targetTerdekat(["o", " ", " ", " ", "x", "x", "x"])); // 4
console.log(targetTerdekat(["x", " ", " ", " ", "x", "x", "o", " "])); // 1
console.log(targetTerdekat([" ", " ", "o", " "])); // 0
console.log(targetTerdekat([" ", "o", " ", "x", "x", " ", " ", "x"])); // 2
console.log(targetTerdekat([" ", "o", " ", "x", "x", "o", " ", "x"])); // 1
