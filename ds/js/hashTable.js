function HashTable() {
  const storage = [];
  const storageLimit = 10;

  function hash(string, max) {
    let hash = 0;
    for (let i = 0; i < string.length; i++) {
      hash += string.charCodeAt(i);
    }

    return hash % max;
  }

  // Adds new element to the hashMap
  this.add = function(key, value) {
    const index = hash(key, storageLimit);

    if (!storage[index]) {
      storage[index] = [
          [key, value]
      ];
    } else {
      let inserted = false;
      for (let i = 0; i < storage[index].length; i++) {
        if (storage[index][i][0] === key) {
          storage[index][i][1] = value;
          inserted = true;
        }
      }

      if (inserted) {
        storage[index].push([key, value]);
      }
    }
  }

  // Returns value of the key
  this.lookup = function(key) {
    const index = hash(key, storageLimit);

    if (!storage[index]) {
      return undefined;
    }

    for (let i = 0; i < storage[index].length; i++) {
      if (storage[index][i][0] === key) {
        return storage[index][i][1];
      }
    }
  }

  // Removes value with key
  this.remove = function(key) {
    const index = hash(key, storageLimit);

    if (!storage[index]) {
      return false;
    }

    delete storage[index];
  }

  this.values = function() {
    return storage;
  }
}

const map = new HashTable();
map.add('s', 1);
map.add('b', 5);
map.add('g', 2);
map.add('s', 3);
console.log(map.values());
console.log(map.lookup('s'));
