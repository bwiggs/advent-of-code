let min = 246515; // raw min
let realMin = 246666;
let max = 739105; // raw max
let realMax = 699999;


function solve(min, max) {
  let curr = min;
  let results = [];
  // probably a more efficientway to solve this non-linearly, oh well
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
  let seqNums = {};
  for (let i = 1; i < chars.length; i++) {
    // next char must always be bigger or eq to prev
    if (chars[i] < chars[i - 1]) {
      return false
    }
    // track that we did get a run pair
    if (chars[i] == chars[i - 1]) {
      let c = chars[i - 1];
      (seqNums[c] == undefined) ? seqNums[c] = 2: seqNums[c]++;
    }
  }

  let hasSeqPair = Object.values(seqNums).some(v => v == 2);
  return hasSeqPair;
}

console.log(solve(realMin, realMax).length);

// console.log([
//   112233,
//   123444,
//   111122
// ].map(v => {
//   let o = {};
//   o[v] = test(v);
//   return o;
// }));