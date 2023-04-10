use std::fs;

fn main() {
    let file = fs::read_to_string("../../input/day02.txt").unwrap();

    let mut square_feet = 0;

    for line in file.split("\n").into_iter() {

        let dim : Vec<i32> = line.split("x").map(|l| l.parse::<i32>().unwrap()).collect();

        let mut sides = [dim[0] * dim[1], dim[1] * dim[2], dim[2] * dim[0]];

        square_feet +=
            sides.iter().map(|s| s * 2).into_iter().sum::<i32>() +
            sides.iter().min().unwrap();
    }

    println!("{:?}", square_feet);
}