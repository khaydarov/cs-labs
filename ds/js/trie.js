function Node() {
  let keys = new Map();
  let end = false;

  this.setEnd = function() {
    end = true;
  }

  this.isEnd = function() {
    return end;
  }

  this.getKeys = function() {
    return keys;
  }
}

function Trie() {
  const root = new Node();

  // Adds word to the prefix tree
  this.add = function(input, node = root) {
    if (input.length === 0) {
      node.setEnd();
    } else if (!node.getKeys().has(input[0])) {
      node.getKeys().set(input[0], new Node());
      return this.add(input.substr(1), node.getKeys().get(input[0]));
    } else {
      return this.add(input.substr(1), node.getKeys().get(input[0]));
    }
  }

  // checks if word is exists
  this.isWord = function(word) {
    let node = root;
    while (word.length > 1) {
      if (!node.getKeys().has(word[0])) {
        return false;
      }

      node = node.getKeys().get(word[0]);
      word = word.substr(1);
    }

    return (node.getKeys().has(word) && node.getKeys().get(word).isEnd());
  }

  // prints dictionary
  this.print = function() {
    let words = [];

    function search(node = root, string) {
      if (node.getKeys().size !== 0) {
        for (let letter of node.getKeys().keys()) {
          search(node.getKeys().get(letter), string.concat(letter));
        }

        if (node.isEnd()) {
          words.push(string);
        }
      } else {
        string.length > 0 ? words.push(string) : undefined;
      }
    }

    search(root, '');
    return words;
  }
}

const trie = new Trie();
trie.add("privet")
trie.add("poka");
console.log(trie.isWord("privet"));
console.log(trie.print());
