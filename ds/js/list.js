function Node(el) {
  let element = el;
  let next = null;

  this.getElement = function() {
    return element;
  }

  this.getNext = function() {
    return next;
  }

  this.setNext = function(node) {
    next = node;
  }
}

function LinkedList() {
  let length = 0;
  let head = null;

  this.head = function() {
    return head;
  }

  // Returns list size
  this.size = function() {
    return length;
  }

  // Add new element
  this.add = function(element) {
    const node = new Node(element);

    if (!head) {
      head = node;
    } else {
      let current = head;
      while (current.getNext()) {
        current = current.getNext();
      }

      current.setNext(node);
    }

    length++;
  }

  // Returns index of element
  this.indexOf = function(element) {
    if (!head) {
      return -1;
    }

    let current = head;
    let position = 0;
    while (current) {
      if (current.getElement() === element) {
        return position;
      }

      current = current.getNext();
    }

    return -1;
  }

  // Removes element by value
  this.remove = function(element) {
    if (!head) {
      return;
    }

    let removed = false;
    if (head.getElement() === element) {
      head = head.getNext();
      removed = true;
    } else {
      let current = head;
      let previous = current;

      while (current.getNext()) {
        if (current.getElement() === element) {
          previous.setNext(current.getNext());
          current = null;
          removed = true;
          break;
        }
        previous = current;
        current = current.getNext();
      }
    }

    if (removed) {
      length--;
    }
  }

  // Returns element at position
  this.elementAt = function(index) {
    if (index > this.size()) {
      return -1;
    }

    let position = 0;
    let current = head;
    while (position < index) {
      current = current.getNext();
      position++;
    }

    return current.getElement();
  }

  this.addAt = function (index, element) {
    if (index > this.size()) {
      return false;
    }

    let position = 0;
    let current = head;
    let previous = current;

    while (position < index) {
      previous = current;
      current = current.getNext();
      position++;
    }

    const node = new Node(element);
    node.setNext(current);
    previous.setNext(node);
    length++;
  }
}

const list = new LinkedList();
list.add(123);
list.add(555);
list.add(321);
list.remove(123);
console.log(list.head().getElement());
console.log(list.size());
console.log(list.elementAt(0));
list.addAt(2, 999);
console.log(list.elementAt(2));
