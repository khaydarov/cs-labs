<?php

$infinity = 1000000000;

$amount = 180;

$notesCount = 3;
$notes = [60, 60, 100];

$notesAmount[0] = 0;
for ($i = 1; $i <= $amount; $i++) {
    $notesAmount[$i] = $infinity;

    for($j = 0; $j < $notesCount; $j++) {
        if ($i >= $notes[$j] && $notesAmount[$i - $notes[$j]] + 1 < $notesAmount[$i]) {
            $notesAmount[$i] = $notesAmount[$i - $notes[$j]] + 1;
        }
    }
}

if ($notesAmount[$amount] === $infinity) {
    echo 'Impossible' . PHP_EOL;
} else {
    echo $notesAmount[$amount] . PHP_EOL;
    while ($amount > 0) {
        for($i = 0; $i < $notesCount; $i++) {
            if ($notesAmount[$amount - $notes[$i]] === $notesAmount[$amount] - 1) {
                echo $notes[$i] . ' ';
                $amount -= $notes[$i];
                break;
            }
        }
    }
}