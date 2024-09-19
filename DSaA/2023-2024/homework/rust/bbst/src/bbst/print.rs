use crate::bbst::Bbst;
use std::fmt::Display;

impl Bbst {
    pub fn print_inorder(&self) {
        if let Some(head) = &self.head {
            head.print_inorder()
        } else {
            print!("Empty");
        }

        println!();
    }

    pub fn print_postorder(&self) {
        if let Some(head) = &self.head {
            head.print_postorder()
        } else {
            print!("Empty");
        }

        println!();
    }

    pub fn print_preorder(&self) {
        if let Some(head) = &self.head {
            head.print_preorder()
        } else {
            print!("Empty");
        }

        println!();
    }
}

impl Display for Bbst {
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

                writeln!(f, "\nsize: {}, is_full: {}\n", self.size(), self.is_full());
            }
            None => {
                write!(f, "Empty");
            }
        };

        Ok(())
    }
}
