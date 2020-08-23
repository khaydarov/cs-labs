<?php

/**
 * Class MaximumSubarrayProblem
 */
class MaximumSubarrayProblem
{
    /**
     * @var array
     */
    private $array = [];

    /**
     * @var MSPImplementation
     */
    private $algorithm;

    /**
     * MaximumSubarrayProblem constructor.
     * @param array $array
     * @param MSPImplementation $algorithm
     */
    public function __construct(array $array, MSPImplementation $algorithm)
    {
        $this->array = $array;
        $this->algorithm = $algorithm;
    }

    /**
     * @return array
     */
    public function getSubarrayWithMaxSum(): array
    {
        return $this->algorithm->execute($this->array);
    }
}

/**
 * Interface MSPImplementation
 */
interface MSPImplementation
{
    public function execute(array $array): array;
}

/**
 * Class KADR
 */
class KADR implements MSPImplementation
{
    /**
     * @param array $array
     * @return array
     */
    public function execute(array $array): array
    {
        $ans = $array[0];
        $ans_l = 0;
        $ans_r = 0;
        $sum = 0;
        $min_sum = 0;
        $min_pos = -1;

        for($r = 0; $r < count($array); $r++) {
            $sum += $array[$r];
            $cur = $sum - $min_sum;

            if ($cur > $ans) {
                $ans = $cur;
                $ans_l = $min_pos + 1;
                $ans_r = $r;
            }

            if ($sum < $min_sum) {
                $min_sum = $sum;
                $min_pos = $r;
            }
        }

        return [
            $ans_l,
            $ans_r,
            $ans
        ];
    }
}

/**
 * Class Naive
 */
class Naive implements MSPImplementation
{
    /**
     * @param array $array
     * @return array
     */
    public function execute(array $array): array
    {
        $s = [];
        $s[0] = 0;
        for($i = 0; $i < count($array); $i++) {
            $s[$i + 1] = $s[$i] + $array[$i];
        }

        $left = 0;
        $right = 0;
        $max = $array[0];
        for ($r = 0; $r < count($s); $r++) {
            for ($l = 0; $l < $r; $l++) {
                if ($s[$r] - $s[$l] > $max) {
                    $max = $s[$r] - $s[$l];
                    $left = $l;
                    $right = $r;
                }
            }
        }

        return [
            $left,
            $right - 1,
            $max
        ];
    }
}

$msp = new MaximumSubarrayProblem([2, -1, 3, -9, 4], new Naive());
var_dump($msp->getSubarrayWithMaxSum());

$msp1 = new MaximumSubarrayProblem([2, -1, 3, -9, 4], new KADR());
var_dump($msp1->getSubarrayWithMaxSum());
