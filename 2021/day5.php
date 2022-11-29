<?php

function count_overlap($matrix) {
    $amount_of_overlap = 0;

    foreach ($matrix as $row) {
        foreach ($row as $cell) {
            if ($cell > 1) {
                $amount_of_overlap++;
            }
        }
    }

    return $amount_of_overlap;
}

$input = fopen('day5_input.txt', 'r');

$data = [];
$matrix = [];

for ($i = 0; $i <= 1000; $i++) {
    $matrix[] = [];
    for ($u = 0; $u <= 1000; $u++) {
        $matrix[$i][] = 0;
    }
}

while (!feof($input)) {
    fscanf($input, '%d,%d -> %d,%d', $x1, $y1, $x2, $y2);
    $data[] = [$x1, $y1, $x2, $y2];
}

foreach ($data as $data_line) {
    $from_x = min($data_line[0], $data_line[2]);
    $from_y = min($data_line[1], $data_line[3]);
    $to_x = max($data_line[0], $data_line[2]);
    $to_y = max($data_line[1], $data_line[3]);

    // Draw a horizontal line
    if ($from_y === $to_y) {
        for ($i = $from_x; $i <= $to_x; $i++) {
            $matrix[$from_y][$i]++;
        }
    }

    // Draw a vertical line
    else if ($from_x === $to_x) {
        for ($i = $from_y; $i <= $to_y; $i++) {
            $matrix[$i][$from_x]++;
        }
    }

    // Draw a diagonal line
    else {
        // y = 50
        // x = 100
        // difference = 50, so it will be 50 + iteration

        $from_x = $data_line[0];
        $from_y = $data_line[1];
        $to_x = $data_line[2];
        $to_y = $data_line[3];

        $y_direction = $to_y <=> $from_y;
        $x_direction = $to_x <=> $from_x;

        for ($i = 0; $i <= abs($to_x - $from_x); $i++) {
            $matrix[$from_x + ($i * $x_direction)][$from_y + ($i * $y_direction)]++;
            $matrix[$from_y + ($i * $y_direction)][$from_x + ($i * $x_direction)]++;
        }
    }
}

//foreach ($matrix as $row) {
//    foreach ($row as $cell) {
//        if ($cell) echo $cell;
//        else echo '.';
//    }
//    echo "\n";
//}

echo count_overlap($matrix);


//var_dump(count_overlap($matrix));