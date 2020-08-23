<?php

/**
 * Class ATM
 */
class ATM
{
    /**
     * @var array
     */
    private $banknotes = [];

    /**
     * @var ATMAlgorithm
     */
    private $algorithm;

    /**
     * ATM constructor.
     * @param array $banknotes
     * @param ATMAlgorithm $algorithm
     */
    public function __construct(array $banknotes, ATMAlgorithm $algorithm)
    {
        $this->banknotes = $banknotes;
        $this->algorithm = $algorithm;
    }

    /**
     * @param int $amount
     * @return array
     */
    public function getBankNotes(int $amount): array
    {
        return $this->algorithm->getBanknotes($this->banknotes, $amount);
    }
}

/**
 * Interface ATMAlgorithm
 */
interface ATMAlgorithm
{
    /**
     * @param array $banknotes
     * @param int $amount
     * @return array
     */
    public function getBanknotes(array $banknotes, int $amount): array;
}

/**
 * Class DpImplementation
 */
class DpImplementation implements ATMAlgorithm
{
    private $inf = 1000000000;

    /**
     * @param array $banknotes
     * @param int $amount
     * @return array
     */
    public function getBanknotes(array $banknotes, int $amount): array
    {
        $dp[0] = 0;
        for ($i = 1; $i <= $amount; $i++) {
            $dp[$i] = $this->inf;

            for($j = 0; $j < count($banknotes); $j++) {
                $banknote = $banknotes[$j];

                if ($i >= $banknote && $dp[$i - $banknote] + 1 < $dp[$i]) {
                    $dp[$i] = $dp[$i - $banknote] + 1;
                }
            }
        }

        if ($dp[$amount] === $this->inf) {
            return [];
        }

        $result = [];
        while ($amount > 0) {
            for($i = 0; $i < count($banknotes); $i++) {
                $banknote = $banknotes[$i];
                if ($amount >= $banknote && $dp[$amount - $banknote] === $dp[$amount] - 1) {
                    $result[] = $banknote;
                    $amount -= $banknote;
                    break;
                }
            }
        }

        return $result;
    }
}

/**
 * Class NaiveImplementation
 */
class NaiveImplementation implements ATMAlgorithm
{
    /**
     * @param array $banknotes
     * @param int $amount
     * @return array
     */
    public function getBanknotes(array $banknotes, int $amount): array
    {
        $result = [];
        for ($i = 0; $i < count($banknotes); $i++) {
            $note = $banknotes[$i];

            while ($amount - $note >= 0) {
                $amount -= $note;
                $result[] = $note;
            }
        }

        return $result;
    }
}

//$atm = new ATM([100, 60], new NaiveImplementation());
//var_dump($atm->getBanknotes(120));

echo PHP_EOL;

$atm1 = new ATM([100, 60], new DpImplementation());
var_dump(count($atm1->getBankNotes(1000000)));