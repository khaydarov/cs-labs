<?php

class Node
{
    /**
     * @var mixed
     */
    private $value;

    /**
     * @var Node
     */
    private $next;

    /**
     * Node constructor.
     * @param $value
     */
    public function __construct($value)
    {
        $this->value = $value;
    }

    /**
     * @return mixed
     */
    public function getValue()
    {
        return $this->value;
    }

    /**
     * @return Node
     */
    public function getNext()
    {
        return $this->next;
    }

    /**
     * @param Node $node
     */
    public function setNext(Node $node)
    {
        $this->next = $node;
    }
}

class LinkedList
{
    /**
     * @var int
     */
    private $length = 0;

    /**
     * @var Node
     */
    private $head;

    /**
     * @param $value
     */
    public function add($value)
    {
        if (!$this->head) {
            $this->head = new Node($value);
        } else {
            $current = $this->head;
            while ($current->getNext() !== null) {
                $current = $current->getNext();
            }

            $node = new Node($value);
            $current->setNext($node);
        }

        $this->length++;
    }

    /**
     * @return Node
     */
    public function getHead(): Node
    {
        return $this->head;
    }

    /**
     * @return int
     */
    public function getSize(): int
    {
        return $this->length;
    }

    /**
     * @param $element
     * @return int
     */
    public function indexOf($element): int
    {
        if (!$this->head) {
            return -1;
        }

        $position = 0;
        $current = $this->head;
        while ($current) {
            if ($current->getValue() === $element) {
                return $position;
            }

            $current = $current->getNext();
            $position++;
        }

        return -1;
    }

    /**
     * @param $element
     */
    public function remove($element)
    {
        if (!$this->head) {
            return;
        }

        $removed = false;
        if ($this->head->getValue() === $element) {
            $this->head = $this->head->getNext();
            $removed = true;
        } else {
            $previous = $this->head;
            $current = $previous;

            while ($current) {
                if ($current->getValue() === $element) {
                    $previous->setNext($current->getNext());
                    $current = null;
                    $removed = true;
                    break;
                }

                $previous = $current;
                $current = $current->getNext();
            }
        }

        if ($removed) {
            $this->length--;
        }
    }

    /**
     * @param $index
     * @return mixed|null
     */
    public function elementAt($index)
    {
        if (!$this->head) {
            return null;
        }

        if ($this->length < $index) {
            return null;
        }

        $position = 0;
        $current = $this->head;
        while ($position < $index) {
            $current = $current->getNext();
            $position++;
        }

        return $current->getValue();
    }

    public function addAt($index, $value)
    {
        if ($this->length < $index) {
            return;
        }

        $position = 0;
        $current = $this->head;
        $previous = $current;
        while ($position < $index) {
            $previous = $current;
            $current = $current->getNext();
            $position++;
        }

        $node = new Node($value);
        $node->setNext($current);
        $previous->setNext($node);
        $this->length++;
    }
}

$list = new LinkedList();
$list->add(10);
$list->add(20);
$list->add(25);
$list->add(30);
$list->addAt(1, 15);

var_dump($list->elementAt(2));

