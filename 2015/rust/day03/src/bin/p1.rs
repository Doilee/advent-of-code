use std::collections::btree_map::Entry;
use std::collections::HashMap;
use std::fs;
use std::iter::Map;
use std::ops::Range;

fn main() {
    let file = fs::read_to_string("../../input/day03.txt").unwrap();


    let mut matrix = HashMap::new();
    let mut pos = (0, 0);

    // starting with 1 because first house needs a present
    let mut amount_of_houses_that_have_atleast_one_present = 1;

    matrix.insert(pos, true);

    for char in file.chars().into_iter() {

        match char {
            '^' => pos = (pos.0, pos.1 - 1),
            '>' => pos = (pos.0 + 1, pos.1),
            '<' => pos = (pos.0 - 1, pos.1),
            'v' => pos = (pos.0, pos.1 + 1),
            _ => panic!("invalid input: {}", char),
        }

        matrix.entry(pos).or_insert_with(|| {
            amount_of_houses_that_have_atleast_one_present += 1;
            return true;
        });
    }

    println!("{:?}", amount_of_houses_that_have_atleast_one_present);
}