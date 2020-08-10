<?php

class queue
{
    /**
     * @var array
     */
    private $storage = [];

    /**
     * @var int
     */
    private $count = 0;

    /**
     * @param $element
     */
    public function push($element)
    {
        $this->storage[$this->count++] = $element;
    }

    /**
     * @return mixed|null
     */
    public function front()
    {
        if ($this->count === 0) {
            return null;
        }

        $this->count--;
        return array_shift($this->storage);
    }

    /**
     * @return int
     */
    public function size(): int
    {
        return $this->count;
    }
}

$queue = new Queue();
$queue->push('first');
$queue->push('second');

echo $queue->front();
echo $queue->front();
echo $queue->front();