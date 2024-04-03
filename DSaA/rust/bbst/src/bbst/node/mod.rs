mod print;
mod rotate;

use anyhow::{anyhow, Ok, Result};
use std::{
    cmp::{max, Ordering},
    result,
};

pub(super) type BoxNode = Option<Box<Node>>;

#[derive(Debug)]
pub(super) struct Node {
    pub(super) data: i32,
    pub(super) left: BoxNode,
    pub(super) right: BoxNode,
    pub(super) height: isize,
}

impl Node {
    pub(super) fn new(data: i32, left: BoxNode, right: BoxNode) -> Self {
        let left_height = left.as_ref().map_or(0, |left| left.height);
        let right_height = right.as_ref().map_or(0, |right| right.height);

        let height = max(left_height, right_height) + 1;

        Self {
            data,
            left,
            right,
            height,
        }
    }

    pub(super) fn add(root: &mut BoxNode, data: i32) -> Result<()> {
        if let Some(root_node) = root {
            let result = match data.cmp(&root_node.data) {
                Ordering::Less => {
                    if root_node.left.is_some() {
                        Node::add(&mut root_node.left, data)
                    } else {
                        root_node.left = Some(Box::new(Node::new(data, None, None)));
                        Ok(())
                    }
                }
                Ordering::Equal => {
                    // let new_node = Node::new(data, root_node.left.take(), None);
                    // root_node.left = Some(Box::new(new_node));
                    Err(anyhow!("Bst already contain this value"))
                }
                Ordering::Greater => {
                    if root_node.right.is_some() {
                        Node::add(&mut root_node.right, data)
                    } else {
                        root_node.right = Some(Box::new(Node::new(data, None, None)));
                        Ok(())
                    }
                }
            };

            if result.is_ok() {
                let left_height = root_node.left.as_ref().map_or(0, |left| left.height);
                let right_height = root_node.right.as_ref().map_or(0, |right| right.height);

                root_node.height = max(left_height, right_height) + 1;

                Node::balance(root);
                Ok(())
            } else {
                result
            }
        } else {
            *root = Some(Box::new(Node::new(data, None, None)));
            Ok(())
        }
    }

    pub(super) fn is_full(&self) -> bool {
        let sides = (&self.left, &self.right);

        if let (Some(left), Some(right)) = sides {
            left.is_full() && right.is_full()
        } else {
            matches!(sides, (None, None))
        }
    }

    pub(super) fn search(&self, data: i32) -> Option<&Self> {
        match data.cmp(&self.data) {
            Ordering::Equal => Some(self),
            Ordering::Greater => self.right.as_ref().and_then(|right| right.search(data)),
            Ordering::Less => self.left.as_ref().and_then(|left| left.search(data)),
        }
    }

    pub(super) fn delete(root: &mut BoxNode, data: i32) -> Result<()> {
        if let Some(ref mut node) = root {
            let result = match data.cmp(&node.data) {
                Ordering::Less => {
                    let result = Node::delete(&mut node.left, data);
                    if result.is_ok() {
                        node.height -= 1;
                    }

                    result
                }
                Ordering::Greater => {
                    let result = Node::delete(&mut node.right, data);
                    if result.is_ok() {
                        node.height -= 1;
                    };

                    result
                }
                Ordering::Equal => {
                    match (&node.left, &node.right) {
                        (None, None) => {
                            *root = None;
                        }
                        (Some(_), None) => {
                            *root = node.left.take();
                        }
                        (None, Some(_)) => {
                            *root = node.right.take();
                        }
                        (Some(_), Some(_)) => {
                            node.data = Node::delete_min(&mut node.right).unwrap();
                        }
                    };
                    Ok(())
                }
            };

            Node::balance(root);

            result
        } else {
            Err(anyhow!("Element not found"))
        }
    }

    fn delete_min(root: &mut BoxNode) -> Option<i32> {
        if root.as_ref().unwrap().left.is_some() {
            let root = root.as_mut().unwrap();
            if root.right.is_none() {
                root.height -= 1;
            }
            Node::delete_min(&mut root.left)
        } else {
            let node = root.take().unwrap();
            *root = node.left;
            Some(node.data)
        }
    }

    pub(super) fn count_left(&self) -> i32 {
        self.left.as_ref().map_or(0, |left| 1 + left.count_left())
            + self.right.as_ref().map_or(0, |right| right.count_left())
    }
}

impl PartialEq for Node {
    fn eq(&self, other: &Self) -> bool {
        self.data == other.data
    }
}
