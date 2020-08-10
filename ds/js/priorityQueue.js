function PriorityQueue() {
  const collection = [];

  this.isEmpty = function() {
    return collection.length === 0;
  }

  this.push = function(element) {
    if (this.isEmpty()) {
      collection.push(element);
    } else {
      let added = false;

      for (let i = 0; i < collection.length; i++) {
        if (element.priority < collection[i].priority) {
          collection.splice(i, 0, element);
          added = true;
          break;
        }
      }

      if (!added) {
        collection.push(element);
      }
    }
  }

  this.front = function() {
    return collection.shift();
  }

  this.size = function() {
    return collection.length;
  }

  this.values = function() {
    return collection;
  }
}

const queue = new PriorityQueue();

queue.push({
  task: 'Last',
  priority: 4
})

queue.push({
  task: 'First',
  priority: 2
});

queue.push({
  task: 'Second',
  priority: 1
});

console.log(queue.values());
