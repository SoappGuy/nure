#![allow(unused)]

use std::usize;

use anyhow::{anyhow, Ok, Result};

#[derive(Debug)]
pub struct Node {
    pub data: (usize, usize, i32),
    pub next: Option<Box<Node>>,
}

type LinkNode = Option<Box<Node>>;

impl Node {
    pub fn find_min(node: &LinkNode, curr_i: usize, curr_min: i32, curr_min_i: usize) -> usize {
        if let Some(node) = &node {
            if node.data.2 < curr_min {
                Node::find_min(&node.next, curr_i + 1, node.data.2, curr_i + 1)
            } else {
                Node::find_min(&node.next, curr_i + 1, curr_min, curr_min_i)
            }
        } else {
            curr_min_i
        }
    }

    pub fn pop_at(node: &mut LinkNode, i: usize) -> Option<(usize, usize, i32)> {
        if let Some(node_unwrapped) = node {
            if i == 0 {
                let data = node_unwrapped.data;
                *node = node_unwrapped.next.take();

                Some(data)
            } else {
                Node::pop_at(&mut node_unwrapped.next, i - 1)
            }
        } else {
            None
        }
    }
}

#[derive(Debug)]
pub struct LinkedList {
    pub head: Option<Box<Node>>,
    pub len: usize,
}

impl LinkedList {
    pub fn new() -> Self {
        Self { head: None, len: 0 }
    }

    pub fn add(&mut self, data: (usize, usize, i32)) {
        let new = Some(Box::new(Node {
            data,
            next: self.head.take(),
        }));

        self.head = new;
        self.len += 1;
    }

    pub fn find_min(&self) -> Option<usize> {
        self.head
            .as_ref()
            .map(|head| Node::find_min(&head.next, 0, head.data.2, 0))
    }

    pub fn pop_at(&mut self, i: usize) -> Option<(usize, usize, i32)> {
        if let Some(head) = &mut self.head {
            if i == 0 {
                let data = head.data;
                self.head = head.next.take();

                Some(data)
            } else {
                Node::pop_at(&mut head.next, i - 1)
            }
        } else {
            None
        }
    }

    pub fn pop_min(&mut self) -> Option<(usize, usize, i32)> {
        self.pop_at(self.find_min()?)
    }
}
