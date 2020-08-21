function Node(el) {
    let element = el;
    let next, previous = null;

    this.getElement = function() {
        return element;
    }

    this.setNext = function(node) {
        next = node;
    }

    this.getNext = function() {
        return next;
    }

    this.setPrevious = function(node) {
        previous = node;
    }

    this.getPrevious = function() {
        return previous;
    }
}

function DoubleLinkedList() {
    let counter = 0;
    let head = null;

    this.getHead = function() {
        return head;
    }

    this.add = function(el) {
        if (!head) {
            head = new Node(el);
        } else {
            let current = head;
            while (current.getNext()) {
                current = current.getNext();
            }

            const node = new Node(el);
            current.setNext(node);
            node.setPrevious(current);
        }
    }
}

const list = new DoubleLinkedList();
list.add(5);
list.add(6);