use std::fs;

fn main() {
    let file = fs::read_to_string("../../input/day02.txt").unwrap();

    let mut ribbon_length_in_feet = 0;

    for line in file.split("\n").into_iter() {
        let mut dim : Vec<i32> = line.split("x").map(|l| l.parse::<i32>().unwrap()).collect();

        dim.sort();

        ribbon_length_in_feet +=
            dim[0] + dim[0] + dim[1] + dim[1] +
            dim.into_iter().reduce(|acc, s| acc * s).unwrap();
    }

    println!("{:?}", ribbon_length_in_feet);
}