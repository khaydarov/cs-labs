<?php

class PriorityQueue
{
    /**
     * @var array
     */
    private $storage = [];

    /**
     * @param $element
     */
    public function push($element, $priority)
    {
        if (count($this->storage)) {
            $this->storage[] = [$element, $priority];
        } else {
            $position = 0;
            foreach ($this->storage as $pair) {
                list($elementValue, $elementPriority) = $pair;
                if ($elementPriority < $priority) {
                    break;
                }
                $position++;
            }

            array_splice($this->storage, $position, 0, [[$element, $priority]]);
        }
    }

    /**
     * @return mixed|null
     */
    public function front()
    {
        if (count($this->storage) === 0) {
            return null;
        }

        return array_shift($this->storage);
    }

    /**
     * @return int
     */
    public function size(): int
    {
        return count($this->storage);
    }
}

$queue = new PriorityQueue();
$queue->push('first', 5);
$queue->push('second', 4);
$queue->push('third', 3);

var_dump($queue->front());
var_dump($queue->front());