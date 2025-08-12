const readline = await import("readline");

const arg: string = `*[Math]*

Jawablah pertanyaan di bawah ini
`;

const arg2: string = `

Waktu: 30 detik
Hadiah:
- $300 balance
- 250 XP`;

function solveMath(params: string): number {
    const splittedParams: string[] = params
        .split(arg)
        .join("")
        .split(arg2)
        .join("")
        .split(" ");

    let temp: number = 0;
    let currentOperator: string = "";

    for (const [index, item] of splittedParams.entries()) {
        if (index === 0) {
            temp = Number(item);
            continue;
        } else if (index % 2 !== 0) {
            currentOperator = item;
        } else {
            switch (currentOperator) {
                case "÷":
                    temp /= Number(item);
                    break;
                case "×":
                    temp *= Number(item);
                    break;
                case "+":
                    temp += Number(item);
                    break;
                case "-":
                    temp -= Number(item);
                    break;
            }
        }
    }

    return temp;
}

function getMathQuestion() {
    const rl = readline.createInterface({
        input: process.stdin,
        output: process.stdout,
        prompt: ""
    });

    rl.prompt();
    const lines: string[] = [];

    rl.on("line", line => {
        lines.push(line);
        rl.prompt();
    });

    rl.on("close", () => {
        if (lines.length > 0) {
            const fullText = lines.join("\n");
            if (fullText === "exit") {
                return;
            }
            console.log("HASILNYA : ", solveMath(fullText));
        }
        getMathQuestion();
    });
}

console.log("=============================================");
console.log("MATH QUESTION");
console.log("type 'exit' and press CTRL + C TO EXIT");
console.log("=============================================");

getMathQuestion();

export {};
