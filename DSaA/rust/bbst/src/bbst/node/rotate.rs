use super::{BoxNode, Node};

impl Node {
    pub fn rotate_left(mut node: BoxNode) -> BoxNode {
        if let Some(mut root_node) = node.take() {
            if let Some(mut right_child) = root_node.right.take() {
                root_node.right = right_child.left.take();
                Node::update_height(&mut root_node);

                right_child.left = Some(root_node);
                Node::update_height(&mut right_child);

                Some(right_child)
            } else {
                Some(root_node)
            }
        } else {
            None
        }
    }

    pub fn rotate_right(mut node: BoxNode) -> BoxNode {
        if let Some(mut root_node) = node.take() {
            if let Some(mut left_child) = root_node.left.take() {
                root_node.left = left_child.right.take();
                Node::update_height(&mut root_node);

                left_child.right = Some(root_node);
                Node::update_height(&mut left_child);

                Some(left_child)
            } else {
                Some(root_node)
            }
        } else {
            None
        }
    }

    pub fn rotate_left_right(mut node: BoxNode) -> BoxNode {
        if let Some(mut node) = node {
            if let Some(left_child) = node.left.take() {
                node.left = Node::rotate_left(Some(left_child));
                return Node::rotate_right(Some(node));
            }
            Some(node)
        } else {
            None
        }
    }

    pub fn rotate_right_left(mut node: BoxNode) -> BoxNode {
        if let Some(mut node) = node {
            if let Some(right_child) = node.right.take() {
                node.right = Node::rotate_right(Some(right_child));
                return Node::rotate_left(Some(node));
            }
            Some(node)
        } else {
            None
        }
    }

    fn update_height(node: &mut Node) {
        let left_height = node.left.as_ref().map_or(0, |left| left.height);
        let right_height = node.right.as_ref().map_or(0, |right| right.height);
        node.height = std::cmp::max(left_height, right_height) + 1;
    }

    fn balance_factor(node: &BoxNode) -> isize {
        node.as_ref().map_or(0, |node| {
            let left_height = node.left.as_ref().map_or(0, |node| node.height);
            let right_height = node.right.as_ref().map_or(0, |node| node.height);
            // println!("{}: {}", node.data, left_height - right_height);
            right_height - left_height
        })
    }

    pub fn balance(node: &mut BoxNode) {
        let balance_factor = Node::balance_factor(node);

        if balance_factor == -2 {
            let left_balance = Node::balance_factor(&node.as_mut().unwrap().left);
            // let right_balance = Node::balance_factor(&node.as_mut().unwrap().right);

            if left_balance <= 0 {
                *node = Node::rotate_right(node.take());
            } else {
                *node = Node::rotate_left_right(node.take());
            }
        } else if balance_factor == 2 {
            let right_balance = Node::balance_factor(&node.as_mut().unwrap().right);

            if right_balance >= 0 {
                *node = Node::rotate_left(node.take());
            } else {
                *node = Node::rotate_right_left(node.take());
            }
        }

        if let Some(node) = node {
            Node::update_height(node);
        }
    }

    // pub fn balance(node: &mut BoxNode) {
    //     let balance_factor = Node::balance_factor(node);
    //
    //     if balance_factor > 1 {
    //         if let Some(node) = node {
    //             if Node::balance_factor(&node.right) < 0 {
    //                 node.right = Node::rotate_left(node.right.take());
    //             }
    //         }
    //         *node = Node::rotate_right(node.take());
    //     } else if balance_factor < -1 {
    //         if let Some(node) = node {
    //             if Node::balance_factor(&node.left) > 0 {
    //                 node.left = Node::rotate_right(node.left.take());
    //             }
    //         }
    //         *node = Node::rotate_left(node.take());
    //     }
    //
    //     Node::update_height(node.as_mut().unwrap());
    // }
}
