/**
 * @param {string[]} A
 * @return {string[]}
 */
var commonChars = function(A) {
    let globalMap = {};
    for (let i = 0; i < A.length; i++) {
        const map = {};
        const string = A[i];
        for (let j = 0; j < string.length; j++) {
            if (!map[string[j]]) {
                map[string[j]] = 0;
            }

            map[string[j]]++;
        }

        if (i === 0) {
            globalMap = map;
        } else {
            for (let char in globalMap) {
                if (!map[char]) {
                    globalMap[char] = 0;
                } else {
                    globalMap[char] = Math.min(globalMap[char], map[char]);
                }
            }
        }
    }

    const result = [];
    for (let char in globalMap) {
        const count = globalMap[char];

        for (let j = 0; j < count; j++) {
            result.push(char);
        }
    }

    return result;
};

const strings = ["bella","label","roller"];
// const strings = ["acabcddd","bcbdbcbd","baddbadb","cbdddcac","aacbcccd","ccccddda","cababaab","addcaccd"];
// const strings = ["blll", "ord", "eve"];
// const strings = ["cool","lock","cook"];
commonChars(strings)
