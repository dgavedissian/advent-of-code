use std::fs::File;
use std::io::{BufReader, BufRead, Error};
use std::collections::HashSet;

fn parse_operation(operation: &str) -> i32 {
    match operation.chars().next() {
        Some('+') => operation[1..].parse::<i32>().unwrap(),
        Some('-') => -operation[1..].parse::<i32>().unwrap(),
        _ => 0
    }
}

fn part1() -> Result<i32, Error> {
    let input = File::open("input.txt")?;
    let buffered = BufReader::new(input);

    let mut acc: i32 = 0;
    for line in buffered.lines() {
        acc += parse_operation(line?.as_str());
    }
    Ok(acc)
}

fn part2() -> Result<i32, Error> {
    let input = File::open("input.txt")?;
    let buffered = BufReader::new(input);

    let operations: Vec<_> = buffered.lines().collect();
    let mut seen_frequencies: HashSet<i32> = HashSet::new();
    let mut acc: i32 = 0;
    loop {
        for line in &operations {
            acc += parse_operation(line.as_ref().unwrap().as_str());
            // If the number fails to insert (as it already exists), then we've seen it before.
            if !seen_frequencies.insert(acc) {
                return Ok(acc);
            }
        }
    }
}

fn main() {
    println!("{:?}", part1());
    println!("{:?}", part2());
}
