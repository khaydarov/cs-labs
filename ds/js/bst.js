function Node(data, left, right) {
  this.data = data;
  this.left = left;
  this.right = right;
}

function BST() {
  this.root = null;
}

BST.prototype.add = function(data) {
  const node = this.root;

  if (node === null) {
    this.root = new Node(data);
  } else {
    const searchTree = function(node) {
      if (data < node.data) {
        if (!node.left) {
          node.left = new Node(data);
          return;
        } else {
          return searchTree(node.left);
        }
      } else if (data > node.data) {
        if (!node.right) {
          node.right = new Node(data);
          return;
        } else {
          return searchTree(node.right);
        }
      } else {
        return null;
      }

      return searchTree(node);
    }

    searchTree(this.root);
  }
}

BST.prototype.min = function() {
  let current = this.root;
  while (current.left) {
    current = current.left;
  }

  return current.data;
}

BST.prototype.max = function() {
  let current = this.root;
  while (current.right) {
    current = current.right;
  }

  return current.data;
}

BST.prototype.find = function(data) {
  let current = this.root;
  while (current.data !== data) {
    if (data < current.data) {
      current = current.left;
    } else {
      current = current.right;
    }

    if (!current) {
      return null;
    }
  }

  return current;
}

BST.prototype.isPresent = function(data) {
  let current = this.root;
  while (current) {
    if (data === current.data) {
      return true;
    }

    if (data < current.data) {
      current = current.left;
    } else {
      current = current.right;
    }
  }

  return false;
}

BST.prototype.remove = function(data) {
  const removeNode = function(node, data) {
    if (!node) {
      return null;
    }

    if (data === node.data) {
      if (!node.left && !node.right) {
        return null;
      }

      if (!node.left) {
        return node.right;
      }

      if (!node.right) {
        return node.left;
      }

      let tmpNode = node.right;
      while (tmpNode.left) {
        tmpNode = tmpNode.left;
      }

      node.data = tmpNode.data;
      node.right = removeNode(node.right, tmpNode.data);
    } else if (data < node.data) {
      node.left = removeNode(node.left, data);
      return node;
    } else {
      node.right = removeNode(node.right, data);
      return node;
    }
  }
  this.root = removeNode(this.root, data);
}

const bst = new BST();
bst.add(25);
bst.add(21);
bst.add(36);
bst.add(20);
console.log(bst);
console.log(bst.min());
console.log(bst.max());
console.log(bst.find(5));
