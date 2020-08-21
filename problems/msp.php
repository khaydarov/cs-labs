<?php

$arr = [-2, 1, 3, -9, 4];
$s = [];

foreach ($arr as $i => $item) {
    if ($i === 0) {
        $s[$i] = $item;
    } else {
        $s[$i] = $s[$i - 1] + $item;
    }
}

$s[0] = 0;
for($i = 0; $i < count($arr); $i++) {
    $s[$i + 1] = $s[$i] + $arr[$i];
}

$left = 0;
$right = 0;
$max = 0;
for ($r = 0; $r < count($s); $r++) {
    for ($l = 0; $l < $r; $l++) {
        if ($s[$r] - $s[$l] > $max) {
            $max = $s[$r] - $s[$l];
            $left = $l;
            $right = $r;
        }
    }
}

echo 'Left: ' . $left . PHP_EOL;
echo 'Right: ' . $right . PHP_EOL;
echo 'Max: ' . $max;