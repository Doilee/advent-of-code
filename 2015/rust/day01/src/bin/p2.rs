use std::fs;

fn main() {
    let file = fs::read_to_string("../../input/day01.txt").unwrap();

    let mut result = 0;

    for char in file.chars().into_iter() {
        if char == '(' {
            result += 1;
        } else if char == ')' {
            result -= 1;
        }
    }

    println!("{:?}", result);
}