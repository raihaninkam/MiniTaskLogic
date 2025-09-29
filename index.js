console.log("hello world");

function stringShortener(sentence) {
  let result = [];

  for (let i = 0; i < sentence.length; i++) {
    if (result.length > 0 && result.length - 1 == sentence[i]) {
      result.pop();
      continue;
    } else {
      result.push(sentence[i]);
    }
  }
  return result;
}

console.log(stringShortener("aabbcca"));
