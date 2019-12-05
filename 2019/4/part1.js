let min = 246515; // raw min
let realMin = 246666;
let max = 739105; // raw max
let realMax = 699999;


function solve(min, max) {
  let curr = min;
  let results = [];
  // probably a way to make this more efficient, but oh well
  while (curr <= max) {
    if (test(curr)) {
      results.push(curr);
    }
    curr++;
  }
  return results;
}

function test(pass) {
  let chars = pass.toString().split('');
  let hasSeqPair = false;
  for (let i = 1; i < chars.length; i++) {
    // next char must always be bigger or eq to prev
    if (chars[i] < chars[i - 1]) {
      return false
    }
    // track that we did get a seq pair
    if (chars[i] == chars[i - 1]) {
      hasSeqPair = true;
    }
  }
  return hasSeqPair;
}

results = solve(realMin, realMax);
console.log(results.length);

// console.log([122345, 111123, 135679, 111111, 223450, 123789].map(v => {
//   let o = {};
//   o[v] = test(v);
//   return o;
// }));