const closestPairSum = (numbers: Array<number>): Array<number> => {
    if (numbers.length < 2) return [];

    let temp: Array<number> = [];
    let smallSum: number = 0;
    let tempNum: number = 0;

    const numbersLength: number = numbers.length;

    for (let i = 0; i < numbersLength; i++) {
        const fIndex = numbers[0];
        for (const [j, el] of numbers.entries()) {
            if (j === numbers.length - 1) break;

            const right = numbers[j + 1];

            if (i === 0 && j === 0) {
                temp = [el, right];
                tempNum = Math.abs(el + right);
                smallSum = Math.abs(el) + Math.abs(right);
            } else if (
                Math.abs(fIndex + right) <= tempNum &&
                Math.abs(fIndex) + Math.abs(right) <= smallSum
            ) {
                temp = [fIndex, right];
                tempNum = Math.abs(fIndex + right);
                smallSum = Math.abs(fIndex) + Math.abs(right);
            }
        }
        numbers.splice(0, 1);
    }

    return temp;
};

const arr = [
    100, -101, 30, -5, 1, -1, 99, -100, 3, -3, 2, -2, 7, -7, 50, -49, -51
];
console.log(closestPairSum(arr));
