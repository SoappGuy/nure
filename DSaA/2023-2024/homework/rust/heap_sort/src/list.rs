#![allow(unused)]

use std::cmp::Ordering;
use std::fmt::{Display, Formatter};

#[derive(Debug)]
pub struct List {
    pub len: usize,
    cap: usize,
    pub items: Box<[i32]>,
}

impl List {
    pub fn new() -> Self {
        let len = 0;
        let cap = 10;
        let items = vec![0; cap].into_boxed_slice();
        Self { len, cap, items }
    }

    pub fn add(&mut self, item: i32) {
        if self.len == self.cap {
            self.expand();
        }

        self.items[self.len] = item;
        self.len += 1;
    }

    fn expand(&mut self) {
        self.cap *= 2;
        let mut new_items = vec![0; self.cap].into_boxed_slice();

        for i in 0..(self.len) {
            new_items[i] = self.items[i];
        }
        self.items = new_items;
    }

    pub fn is_empty(&self) -> bool {
        self.len == 0
    }

    pub fn is_full(&self) -> bool {
        self.len == self.cap
    }

    fn heapify(arr: &mut [i32], i: usize, is_max_heap: bool) {
        let comparator = if !is_max_heap {
            |a: i32, b: i32| b.cmp(&a)
        } else {
            |a: i32, b: i32| a.cmp(&b)
        };

        let mut largest = i;
        let left = i * 2 + 1;
        let right = left + 1;

        if left < arr.len() && comparator(arr[left], arr[largest]) == Ordering::Greater {
            largest = left;
        }

        if right < arr.len() && comparator(arr[right], arr[largest]) == Ordering::Greater {
            largest = right;
        }

        if largest != i {
            arr.swap(i, largest);
            Self::heapify(arr, largest, is_max_heap);
        }
    }

    fn build_heap(&mut self, is_max_heap: bool) {
        let mut i = (self.len - 1) / 2;

        while i > 0 {
            Self::heapify(&mut self.items, i, is_max_heap);
            i -= 1;
        }

        Self::heapify(&mut self.items, 0, is_max_heap);
    }

    pub fn heap_sort(&mut self, ascending: bool) {
        if self.len <= 1 {
            return;
        }

        self.build_heap(ascending);
        let mut end = self.len - 1;
        while end > 0 {
            self.items.swap(0, end);
            Self::heapify(&mut self.items[..end], 0, ascending);
            end -= 1;
        }
    }
}

impl Display for List {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        writeln!(
            f,
            "[{}]",
            self.items[0..self.len]
                .iter()
                .map(|item| { item.to_string() })
                .collect::<Vec<String>>()
                .join(", ")
        );
        Ok(())
    }
}
