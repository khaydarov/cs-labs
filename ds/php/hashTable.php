<?php


class HashTable
{
    /**
     * @var array
     */
    private $storage = [];

    /**
     * @var int
     */
    private $limit = 10;

    /**
     * @param $key
     * @param $value
     */
    public function add($key, $value)
    {
        $index = $this->getIndex($key);

        if (!isset($this->storage[$index])) {
            $this->storage[$index] = [
                $key, $value
            ];
        }
    }

    /**
     * @param $key
     * @return mixed|null
     */
    public function get($key)
    {
        $index = $this->getIndex($key);

        if (!isset($this->storage[$index])) {
            return null;
        }

        return $this->storage[$index][1];
    }

    /**
     * @param $key
     */
    public function remove($key)
    {
        $index = $this->getIndex($key);

        if (!isset($this->storage[$index])) {
            return;
        }

        unset($this->storage[$index]);
    }

    private function getIndex($string): int
    {
        $hash = 0;
        for($i = 0; $i < strlen($string); $i++) {
            $hash += ord($string[$i]);
        }

        return $hash % $this->limit;
    }
}

$map = new HashTable();
$map->add('a', 1);
echo $map->get('a');