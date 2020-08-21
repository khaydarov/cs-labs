<?php

$infinity = 1000000000;

$amount = 120;

$notesCount = 2;
$notes = [100, 60];

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
            if ($amount >= $notes[$i] && $notesAmount[$amount - $notes[$i]] === $notesAmount[$amount] - 1) {
                echo $notes[$i] . ' ';
                $amount -= $notes[$i];
                break;
            }
        }
    }
}

//function iWantToGet($amount) {
//    $notes = [100, 60];
//    $result = [];
//
//    for ($i = 0; $i < count($notes); $i++) {
//        $note = $notes[$i];
//
//        while ($amount - $note >= 0) {
//            $amount -= $note;
//            $result[] = $note;
//        }
//    }
//
//    return $result;
//}
//
//var_dump(iWantToGet(120));