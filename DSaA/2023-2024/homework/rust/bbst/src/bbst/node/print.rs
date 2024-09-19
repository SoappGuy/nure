use super::Node;

impl Node {
    pub(in crate::bbst) fn print(
        &self,
        f: &mut std::fmt::Formatter<'_>,
        padding: &str,
        ptr: &str,
        nested: bool,
    ) {
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

    pub(in crate::bbst) fn print_inorder(&self) {
        if let Some(left) = &self.left {
            left.print_inorder();
        }

        print!("{} ", self.data);

        if let Some(right) = &self.right {
            right.print_inorder();
        }
    }

    pub(in crate::bbst) fn print_postorder(&self) {
        if let Some(left) = &self.left {
            left.print_postorder();
        }

        if let Some(right) = &self.right {
            right.print_postorder();
        }

        print!("{} ", self.data);
    }

    pub(in crate::bbst) fn print_preorder(&self) {
        print!("{} ", self.data);

        if let Some(left) = &self.left {
            left.print_preorder();
        }

        if let Some(right) = &self.right {
            right.print_preorder();
        }
    }
}
