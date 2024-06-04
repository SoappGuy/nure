#![allow(unused)]

use std::{fmt::Display, result};

#[derive(Debug)]
pub struct Node {
    pub data: i32,
    pub priority: usize,
    pub next: Option<Box<Node>>,
}

type LinkNode = Option<Box<Node>>;

impl Node {
    pub fn insert(node: &mut LinkNode, data: i32, priority: usize) {
        if let Some(node_unwrapped) = node {
            if node_unwrapped.priority > priority {
                Node::insert(&mut node_unwrapped.next, data, priority);
            } else {
                *node = Some(Box::new(Node {
                    data,
                    priority,
                    next: node.take(),
                }));
            }
        } else {
            *node = Some(Box::new(Node {
                data,
                priority,
                next: None,
            }))
        }
    }
}

#[derive(Debug)]
pub struct PriorityQueue {
    head: Option<Box<Node>>,
    pub len: usize,
}

impl PriorityQueue {
    pub fn new() -> Self {
        Self { head: None, len: 0 }
    }

    pub fn is_empty(&self) -> bool {
        self.head.is_none()
    }

    pub fn is_full(&self) -> bool {
        self.head.is_some()
    }

    pub fn enqueue(&mut self, data: i32, priority: usize) {
        if let Some(head_unwrapped) = &mut self.head {
            if head_unwrapped.priority > priority {
                Node::insert(&mut head_unwrapped.next, data, priority);
            } else {
                self.head = Some(Box::new(Node {
                    data,
                    priority,
                    next: self.head.take(),
                }));
            }
        } else {
            self.head = Some(Box::new(Node {
                data,
                priority,
                next: None,
            }));
        }

        self.len += 1;
    }

    pub fn dequeue(&mut self) -> Option<i32> {
        if let Some(head) = &mut self.head {
            let data = head.data;
            self.head = head.next.take();

            self.len -= 1;
            Some(data)
        } else {
            None
        }
    }
}

impl Display for PriorityQueue {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let mut result: Vec<String> = vec![];

        let mut node = &self.head;

        while let Some(node_unwrapped) = &node {
            result.push(format!(
                "[{}, p: {}]",
                node_unwrapped.data, node_unwrapped.priority
            ));

            node = &node_unwrapped.next;
        }

        if result.is_empty() {
            write!(f, "Empty")
        } else {
            write!(f, "{}", result.join("󰍴󰍴"))
        }
    }
}
