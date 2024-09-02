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

    println!("{bbst}");
    println!("{}", bbst.count_left());

    Ok(())
}
