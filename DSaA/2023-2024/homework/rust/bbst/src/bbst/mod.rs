mod node;
mod print;

use self::node::{BoxNode, Node};
use std::cmp::Ordering;
use std::fmt::Display;

use anyhow::{anyhow, Result};
use rand::random;

#[derive(Debug)]
pub struct Bbst {
    head: BoxNode,
    len: usize,
}

impl Bbst {
    pub fn new() -> Self {
        let len = 0;
        let head = None;

        Self { len, head }
    }

    pub fn add(&mut self, data: i32) -> Result<()> {
        let result = Node::add(&mut self.head, data);

        if result.is_ok() {
            self.len += 1;
        }

        result
    }

    pub fn is_full(&self) -> bool {
        if let Some(head) = &self.head {
            head.is_full()
        } else {
            false
        }
    }

    pub fn is_empty(&self) -> bool {
        self.head.is_none()
    }

    pub fn size(&self) -> usize {
        self.len
    }

    pub fn search(&self, data: i32) -> bool {
        if let Some(head) = &self.head {
            head.search(data).is_some()
        } else {
            false
        }
    }

    pub fn delete(&mut self, data: i32) -> Result<()> {
        let result = Node::delete(&mut self.head, data);
        if result.is_ok() {
            self.len -= 1;
        };
        result
    }

    pub fn count_left(&self) -> i32 {
        self.head.as_ref().map_or(0, |head| head.count_left())
    }
}

impl PartialEq for Bbst {
    fn eq(&self, other: &Self) -> bool {
        self.head == other.head && self.len == other.len
    }
}

#[macro_export]
macro_rules! bbst {
    [ $($x:expr),* ] => {{
        let mut bbst = Bbst::new();
        $( bbst.add($x); )*

        bbst
    }};
}

#[macro_export]
macro_rules! rand_bbst {
    [ $count:expr ] => {{
        let mut bbst = Bbst::new();

        for _i in 0..$count {
            bbst.add(random::<i32>() % 10);
        }

        bbst
    }};
}
