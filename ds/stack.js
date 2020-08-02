function Stack() {
  const storage = [];

  let count = 0;

  this.push = function(value) {
    storage[count++] = value;
  }

  this.peak = function() {
    return storage[count - 1];
  }

  this.pop = function () {
    if (count === 0) {
      return undefined;
    }

    count--;
    const result = storage[count];
    delete storage[count];

    return result;
  }

  this.size = function() {
    return count;
  }
}

const stack = new Stack();
stack.push(10);
stack.push(15);
console.log(stack.size());
console.log(stack.peak());
console.log(stack.pop());
console.log(stack.pop());
stack.push(20);
console.log(stack.size());
