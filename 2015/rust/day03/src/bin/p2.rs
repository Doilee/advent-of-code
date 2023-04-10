use std::collections::btree_map::Entry;
use std::collections::HashMap;
use std::fs;
use std::iter::Map;
use std::ops::Range;
use std::slice::SliceIndex;

fn main() {
    let file = fs::read_to_string("../../input/day03.txt").unwrap();


    let mut matrix = HashMap::new();
    let mut santa_pos = (0, 0);
    let mut robo_santa_pos = (0, 0);

    // starting with 1 because first house needs a present
    let mut amount_of_houses_that_have_atleast_one_present = 1;

    matrix.insert(santa_pos, true);

    for (i, char) in file.chars().enumerate() {
        let i = i as i32;

        let mut pos : &mut (i32, i32) = &mut robo_santa_pos;

        if i % 2 == 0 {
            //is even
            pos = &mut santa_pos;
        } else {
            //is odd
            pos = &mut robo_santa_pos;
        }

        match char {
            '^' => *pos = (pos.0, pos.1 - 1),
            '>' => *pos = (pos.0 + 1, pos.1),
            '<' => *pos = (pos.0 - 1, pos.1),
            'v' => *pos = (pos.0, pos.1 + 1),
            _ => panic!("invalid input: {}", char),
        }

        matrix.entry(*pos).or_insert_with(|| {
            amount_of_houses_that_have_atleast_one_present += 1;
            return true;
        });
    }

    println!("{:?}", amount_of_houses_that_have_atleast_one_present);
}