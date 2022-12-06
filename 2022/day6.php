<?php

$string = file_get_contents('day6.txt');

const PART_ONE = 4;
const PART_TWO = 14;

$lastFour = [];

foreach (str_split($string) as $index => $character) {
    $lastFour[] = $character;

    if (count($lastFour) > PART_TWO) {
        array_shift($lastFour);
    }

    if (arrayHasDupes($lastFour)) {
        var_dump(array_unique($lastFour));
        var_dump($lastFour);
        continue;
    }

    if (count($lastFour) == PART_TWO) {
        var_dump($index + 1);
        die();
    }
}

function arrayHasDupes(array $array) : bool {
    return count(array_unique($array)) !== count($array);
}
