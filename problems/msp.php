<?php

$arr = [2, -1, 3, -9, 4];
$s = [];
$s[0] = 0;
for($i = 0; $i < count($arr); $i++) {
    $s[$i + 1] = $s[$i] + $arr[$i];
}

$left = 0;
$right = 0;
$max = $arr[0];
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
echo 'Max: ' . $max . PHP_EOL;
echo '------' . PHP_EOL;

$ans = $arr[0];
$ans_l = 0;
$ans_r = 0;
$sum = 0;
$min_sum = 0;
$min_pos = -1;

for($r = 0; $r < count($arr); $r++) {
    $sum += $arr[$r];
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

echo 'Left: ' . $ans_l . PHP_EOL;
echo 'Right: ' . $ans_r . PHP_EOL;
echo 'Max: ' . $ans;
