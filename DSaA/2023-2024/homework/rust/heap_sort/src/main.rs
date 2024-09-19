#![allow(unused)]

mod list;
mod queue;

use list::*;
use queue::*;

fn main() {
    // let mut lst = List::new();
    // lst.add(2);
    // lst.add(3);
    // lst.add(8);
    // lst.add(1);
    // lst.add(-5);
    // lst.add(2);
    // lst.add(0);
    // lst.add(1);
    //
    // println!("{lst}");
    //
    // lst.heap_sort(true);
    // println!("{lst}");
    //
    // lst.heap_sort(false);
    // println!("{lst}");

    let mut queue = PriorityQueue::new();

    queue.enqueue(10, 0);
    queue.enqueue(100, 6);
    queue.enqueue(1, 6);
    queue.enqueue(-17, 2);

    println!("{}", queue);

    println!("{:?}", queue.dequeue());
    println!("{:?}", queue.dequeue());
    println!("{:?}", queue.dequeue());
    println!("{}", queue);
}
