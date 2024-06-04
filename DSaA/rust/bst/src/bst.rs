#![allow(unused)]

use anyhow::{anyhow, Result};
use rand::random;
use std::{cmp::Ordering, fmt::Display, iter};

type LinkNode = Option<Box<Node>>;

#[derive(Debug)]
struct Node {
    data: i32,
    left: LinkNode,
    right: LinkNode,
}

impl Node {
    fn print(&self, f: &mut std::fmt::Formatter<'_>, padding: &str, ptr: &str, nested: bool) {
        if let Some(right) = &self.right {
            right.print(
                f,
                &(padding.to_owned() + if nested { "│   " } else { "    " }),
                "┌──",
                false,
            );
        }

        writeln!(f, "{}{}({})", padding, ptr, self.data);

        if let Some(left) = &self.left {
            left.print(
                f,
                &(padding.to_owned() + if !nested { "│   " } else { "    " }),
                "└──",
                true,
            );
        }
    }

    fn print_preorder(&self) {
        print!("{} ", self.data);

        if let Some(left) = &self.left {
            left.print_preorder();
        }

        if let Some(right) = &self.right {
            right.print_preorder();
        }
    }

    fn add(root: &mut LinkNode, data: i32) -> Result<()> {
        if let Some(root_node) = root {
            match data.cmp(&root_node.data) {
                Ordering::Less => {
                    if root_node.left.is_some() {
                        Node::add(&mut root_node.left, data)
                    } else {
                        root_node.left = Some(Box::new(Node {
                            data,
                            left: None,
                            right: None,
                        }));
                        Ok(())
                    }
                }
                Ordering::Equal => Err(anyhow!("Bst already contain this value")),
                Ordering::Greater => {
                    if root_node.right.is_some() {
                        Node::add(&mut root_node.right, data)
                    } else {
                        root_node.right = Some(Box::new(Node {
                            data,
                            left: None,
                            right: None,
                        }));
                        Ok(())
                    }
                }
            }
        } else {
            *root = Some(Box::new(Node {
                data,
                left: None,
                right: None,
            }));
            Ok(())
        }
    }

    fn is_equal(a: &LinkNode, b: &LinkNode) -> bool {
        if a.is_none() && b.is_none() {
            true
        } else if let (Some(a), Some(b)) = (a, b) {
            a.data == b.data
                && Node::is_equal(&a.left, &b.left)
                && Node::is_equal(&a.right, &b.right)
        } else {
            false
        }
    }
}

#[derive(Debug)]
pub struct Bst {
    head: LinkNode,
    len: usize,
}

impl Bst {
    pub fn new() -> Self {
        Self { head: None, len: 0 }
    }

    pub fn print_preorder(&self) {
        if let Some(head) = &self.head {
            head.print_preorder()
        } else {
            print!("Empty");
        }

        println!();
    }

    pub fn add(&mut self, data: i32) -> Result<()> {
        let result = Node::add(&mut self.head, data);

        if result.is_ok() {
            self.len += 1;
        }

        result
    }

    pub fn is_equal(a: &Bst, b: &Bst) -> bool {
        Node::is_equal(&a.head, &b.head)
    }
}

impl Display for Bst {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match &self.head {
            Some(head) => {
                if let Some(right) = &head.right {
                    right.print(f, "    ", "┌──", false);
                }

                writeln!(f, "-->({})", head.data);

                if let Some(left) = &head.left {
                    left.print(f, "    ", "└──", true);
                }
            }
            None => {
                write!(f, "Empty");
            }
        };

        Ok(())
    }
}

#[macro_export]
macro_rules! bst {
    [ $($x:expr),* ] => {{
        let mut bbst = Bst::new();
        $( bbst.add($x); )*

        bbst
    }};
}

#[macro_export]
macro_rules! rand_bst {
    [ $count:expr ] => {{
        let mut bbst = Bst::new();

        for _i in 0..$count {
            bbst.add(random::<i32>() % 10);
        }

        bbst
    }};
}
