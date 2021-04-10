<?php

class TArray
{
    private $storage;
    private $length = 0;
    private $capacity;

    public function __construct($size)
    {
        $this->capacity = $size;
        $this->storage = new SplFixedArray($size);
    }

    public function putAt($index, $value)
    {
        if ($this->length === $this->capacity) {
            $this->increaseCapacity();
        }
        $this->storage[$index] = $value;
        $this->length++;
    }

    public function getCapacity()
    {
        return $this->capacity;
    }

    public function getLength()
    {
        return $this->length;
    }

    private function increaseCapacity()
    {
        $this->capacity *= 2;
        $storage = new SplFixedArray($this->capacity);

        foreach ($this->storage as $index => $value) {
            $storage[$index] = $value;
        }

        $this->storage = $storage;
    }
}

$arr = new TArray(2);
$arr->putAt(0, 0);
$arr->putAt(1, 1);
//$arr->putAt(2, 2);
//$arr->putAt(3, 3);
//$arr->putAt(4, 4);

echo $arr->getCapacity();