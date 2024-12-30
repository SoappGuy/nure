use std::io;

fn main() {
    let mut input = String::new();
    io::stdin().read_line(&mut input).unwrap();

    let mut parts = input.trim().split_whitespace();
    let n: usize = parts.next().unwrap().parse().unwrap();
    let m: usize = parts.next().unwrap().parse().unwrap();

    let mut result: Vec<usize> = (1..=n).collect();

    for _ in 0..m {
        input.clear();
        io::stdin().read_line(&mut input).unwrap();

        let mut parts = input.trim().split_whitespace();
        let l: usize = parts.next().unwrap().parse::<usize>().unwrap() - 1;
        let r: usize = parts.next().unwrap().parse::<usize>().unwrap() - 1;

        let segment = result[l..=r].to_vec();
        let mut rest = result[..l].to_vec();
        rest.extend_from_slice(&result[r + 1..]);
        result = segment;
        result.extend(rest);
    }

    for x in result {
        print!("{} ", x.to_string());
    }
}
