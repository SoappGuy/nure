mod list;

fn main() {
    let mut lst = list::List::new();
    lst.add(2);
    lst.add(3);
    lst.add(8);
    lst.add(1);
    lst.add(-5);
    lst.add(2);
    lst.add(0);
    lst.add(1);

    println!("{lst}");

    lst.heap_sort(true);
    println!("{lst}");

    lst.heap_sort(false);
    println!("{lst}");
}
