use std::fs;

fn main() {
    let file = fs::read_to_string("../../input/day01.txt").unwrap();

    let mut result = 0;

    for (index, char) in file.chars().enumerate() {
        if char == '(' {
            result += 1;
        } else if char == ')' {
            result -= 1;
        }

        if result == -1 {
            println!("Entered the basement at position: {}", index + 1);
        }
    }

    println!("{:?}", result);
}