function Queue() {
  const collection = [];

  this.push = function(element) {
    collection.push(element);
  }

  this.front = function() {
    return collection.shift();
  }

  this.size = function() {
    return collection.length;
  }

  this.isEmpty = function() {
    return collection.length === 0;
  }
}

const queue = new Queue();
queue.push(10);
queue.push(1010);
console.log(queue.front());
console.log(queue.size());
console.log(queue.front());
console.log(queue.size());
