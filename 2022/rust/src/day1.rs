pub fn execute() {
    let data = include_str!("input/day1.txt")
        .split("\n\n")
        .map(|e| e.lines().map(|c| c.parse::<u32>().unwrap()).sum::<u32>())
        .max()
        .unwrap();


     println!("{:?}", data);
}
