<?php

class Set
{
    /**
     * @var array
     */
    private $storage = [];

    /**
     * @param $element
     * @return bool
     */
    public function has($element): bool
    {
        return in_array($element, $this->storage);
    }

    /**
     * @param $element
     */
    public function add($element)
    {
        if ($this->has($element)) {
            return;
        }

        $this->storage[] = $element;
    }

    /**
     * @return int
     */
    public function size(): int
    {
        return count($this->storage);
    }

    /**
     * @param $element
     */
    public function remove($element)
    {
        if (!$this->has($element)) {
            return;
        }

        $position = array_search($element, $this->storage);
        array_splice($this->storage, $position, 1);
    }

    /**
     * @param Set $set
     */
    public function union(Set $set)
    {
        foreach ($set->getValues() as $element) {
            if ($this->has($element)) {
                continue;
            }

            $this->add($element);
        }
    }

    /**
     * @param Set $set
     * @return Set
     */
    public function intersection(Set $set)
    {
        $intersection = new Set();
        foreach ($set->getValues() as $element) {
            if ($this->has($element)) {
                $intersection->add($element);
            }
        }

        return $intersection;
    }

    /**
     * @param Set $set
     * @return Set
     */
    public function difference(Set $set)
    {
        $difference = new Set();
        foreach ($set->getValues() as $element) {
            if (!$this->has($element)) {
                $difference->add($element);
            }
        }

        return $difference;
    }

    /**
     * @return array
     */
    public function getValues(): array
    {
        return $this->storage;
    }
}

$set = new Set();
$set->add(1);
$set->add(2);
$set->add(3);
$set->add(2);

$anotherSet = new Set();
$anotherSet->add(5);
$anotherSet->add(3);

$set->remove(1);

//$set->union($anotherSet);

var_dump($set->getValues());

$intersection = $set->intersection($anotherSet);
var_dump($intersection->getValues());