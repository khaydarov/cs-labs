function MySet() {
  const collection = [];

  this.has = function(element) {
    return collection.indexOf(element) !== -1;
  }

  this.add = function(element) {
    if (this.has(element)) {
      return false;
    }

    collection.push(element);
    return true;
  }

  this.remove = function(element) {
    if (!this.has(element)) {
      return false;
    }

    const position = collection.indexOf(element);
    collection.splice(position, 1);

    return true;
  }

  this.values = function () {
    return collection;
  }

  this.size = function() {
    return collection.length;
  }

  this.union = function(otherSet) {
    const unionSet = new MySet();
    const currentSetValues = this.values();
    const otherSetValues = otherSet.values();

    currentSetValues.forEach((el) => {
      unionSet.add(el);
    });

    otherSetValues.forEach((el) => {
      unionSet.add(el);
    });

    return unionSet;
  }

  this.intersection = function (otherSet) {
    const intersectionSet = new MySet();
    const otherSetValues = otherSet.values();

    otherSetValues.forEach((el) => {
      if (this.has(el)) {
        intersectionSet.add(el);
      }
    });

    return intersectionSet;
  }

  this.difference = function(otherSet) {
    const differenceSet = new MySet();
    const currentSetValues = this.values();

    currentSetValues.forEach((el) => {
      if (!otherSet.has(el)) {
        differenceSet.add(el);
      }
    });

    return differenceSet;
  }
}

const set = new MySet();
set.add(10);
set.add(15);
set.add(25);

const anotherSet = new MySet();
anotherSet.add(10);
anotherSet.add(30);

console.log(set.values());
console.log(set.union(anotherSet).values());
console.log(set.intersection(anotherSet).values());
console.log(set.difference(anotherSet).values());
