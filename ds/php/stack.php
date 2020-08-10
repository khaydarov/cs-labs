<?php

class Stack {
    /**
     * @var int
     */
    private $count = 0;

    /**
     * @var array
     */
    private $storage = [];

    /**
     * @param $element
     */
    public function push($element): void
    {
        $this->storage[$this->count++] = $element;
    }

    /**
     * @return mixed|null
     */
    public function peak()
    {
        if ($this->count === 0) {
            return null;
        }

        return $this->storage[$this->count - 1];
    }

    /**
     * @return int
     */
    public function size(): int
    {
        return $this->count;
    }

    /**
     * @return mixed|null
     */
    public function pop()
    {
        if ($this->count === 0) {
            return null;
        }

        $this->count--;
        return array_pop($this->storage);
    }
}

$stack = new Stack();
$stack->push('first');
$stack->push('second');
$stack->push('third');

echo $stack->size();
echo $stack->pop();
echo $stack->size();
echo $stack->pop();
echo $stack->size();
echo $stack->peak();
