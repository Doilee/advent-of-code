<?php

$string = file_get_contents('../input/06.txt');

const LAST_4 = 4;
const LAST_14 = 14;

$lastFour = [];

foreach (str_split($string) as $index => $character) {
    $lastFour[] = $character;

    if (count($lastFour) > LAST_14) array_shift($lastFour);

    if (arrayHasDupes($lastFour)) continue;

    if (count($lastFour) == LAST_14) {
        var_dump($index + 1);
        die();
    }
}

function arrayHasDupes(array $array) : bool {
    return count(array_unique($array)) !== count($array);
}
