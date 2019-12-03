let program = [
  1, 12, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 13, 1, 19, 1, 19, 10, 23, 2, 10, 23, 27, 1, 27, 6, 31, 1, 13, 31, 35, 1, 13, 35, 39, 1, 39, 10, 43, 2, 43, 13, 47, 1, 47, 9, 51, 2, 51, 13, 55, 1, 5, 55, 59, 2, 59, 9, 63, 1, 13, 63, 67, 2, 13, 67, 71, 1, 71, 5, 75, 2, 75, 13, 79, 1, 79, 6, 83, 1, 83, 5, 87, 2, 87, 6, 91, 1, 5, 91, 95, 1, 95, 13, 99, 2, 99, 6, 103, 1, 5, 103, 107, 1, 107, 9, 111, 2, 6, 111, 115, 1, 5, 115, 119, 1, 119, 2, 123, 1, 6, 123, 0, 99, 2, 14, 0, 0
  // 1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50
];

let operators = [
  (a, b) => a + b,
  (a, b) => a * b
];

function solve(program) {
  let i = 0;
  while (program[i] != 99) {
    let op = operators[program[i++] - 1];
    let regA = program[program[i++]];
    let regB = program[program[i++]];
    let regC = program[i++];
    program[regC] = op(regA, regB)
  }
  return program[0];
}


// console.log(solve(program.slice()));
for (let i = 0; i < 100; i++) {
  for (let j = 0; j < 100; j++) {
    let prg = program.slice();
    prg[1] = i;
    prg[2] = j;

    if (solve(prg) == 19690720) {
      console.log(100 * i + j);
      break;
    }
  }
}