use std::collections::HashMap;

const ROCK_SCORE: i32 = 1;
const PAPER_SCORE: i32 = 2;
const SCISSORS_SCORE: i32 = 3;

pub fn part1() {
    let input = include_str!("input/day2.txt")
        .split("\n");

    let rps: HashMap<char, &str> = HashMap::from([
        ('A', "rock"),
        ('B', "paper"),
        ('C', "scissors"),
        ('X', "rock"),
        ('Y', "paper"),
        ('Z', "scissors"),
    ]);

    let mut score = 0;

    for part in input {
        let opponent_hand = part.chars().next().unwrap();
        let my_hand = part.chars().last().unwrap();

        let o = rps.get(&opponent_hand).unwrap();
        let m = rps.get(&my_hand).unwrap();

        score += match o.to_owned() {
            "rock" => match m.to_owned() {
                "rock" => 3 + ROCK_SCORE,
                "paper" => 6 + PAPER_SCORE,
                _ => SCISSORS_SCORE
            },
            "paper" => match m.to_owned() {
                "paper" => 3 + PAPER_SCORE,
                "scissors" => 6 + SCISSORS_SCORE,
                _ => ROCK_SCORE
            },
            "scissors" => match m.to_owned() {
                "scissors" => 3 + SCISSORS_SCORE,
                "rock" => 6 + ROCK_SCORE,
                _ => PAPER_SCORE
            },
            _ => 0
        }
    }

    println!("{}", score);
}

pub fn part2() {
    let input = include_str!("input/day2.txt")
        .split("\n");

    let meaning: HashMap<char, &str> = HashMap::from([
        ('A', "rock"),
        ('B', "paper"),
        ('C', "scissors"),
        ('X', "lose"),
        ('Y', "draw"),
        ('Z', "win"),
    ]);

    let mut score = 0;

    for part in input {
        let opponent_hand = part.chars().next().unwrap();
        let my_hand = part.chars().last().unwrap();

        let o = meaning.get(&opponent_hand).unwrap();
        let strategy = meaning.get(&my_hand).unwrap();

        score += match o.to_owned() {
            "rock" => match strategy.to_owned() {
                "draw" => 3 + ROCK_SCORE,
                "win" => 6 + PAPER_SCORE,
                "lose" => 0 + SCISSORS_SCORE,
                _ => 0
            },
            "paper" => match strategy.to_owned() {
                "draw" => 3 + PAPER_SCORE,
                "win" => 6 + SCISSORS_SCORE,
                "lose" => 0 + ROCK_SCORE,
                _ => 0
            },
            "scissors" => match strategy.to_owned() {
                "draw" => 3 + SCISSORS_SCORE,
                "win" => 6 + ROCK_SCORE,
                "lose" => 0 + PAPER_SCORE,
                _ => 0
            },
            _ => 0
        }
    }

    println!("{}", score);
}