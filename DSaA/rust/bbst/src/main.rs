#![allow(unused)]

#[macro_use]
mod bbst;

use crate::bbst::*;
use anyhow::Result;
use rand::random;

fn main() -> Result<()> {
    // let mut bbst = bbst![0, 2, 3, 1, -2, -3, -1];
    // let mut bbst = bbst![10, 5, 11, 4, 7, 3, 2, 6];
    let mut bbst = rand_bbst![10];

    // let mut bbst = bbst![5, 9, 2, 1, 8, 4, 11];

    // let mut bbst = Bbst::new();
    // bbst.add(2);
    // bbst.add(3);
    // println!("{:?}", bbst.is_empty());
    // bbst.delete(6);
    // bbst.delete(2);
    println!("{bbst}");
    println!("{}", bbst.count_left());

    Ok(())
}
