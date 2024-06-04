#![allow(unused)]

mod bst;

#[macro_use]
use bst::*;
use rand::random;

fn main() {
    let mut bst1 = bst![1, 2, 0, 4];

    let mut bst2 = bst![1, 2, 0, 4];
    let mut bst3 = bst![1, 3, 0, 4];
    let mut bst4 = rand_bst!(4);

    println!("bst1:\n{bst1}\n");
    println!("bst2:\n{bst2}\n");
    println!("bst3:\n{bst3}\n");
    println!("bst4:\n{bst4}\n");

    println!("bst1 == bst2: {}", Bst::is_equal(&bst1, &bst2));
    println!("bst1 == bst3: {}", Bst::is_equal(&bst1, &bst3));
    println!("bst1 == bst4: {}", Bst::is_equal(&bst1, &bst4));
}
