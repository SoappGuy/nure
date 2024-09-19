#![allow(unused)]

mod graph;
mod linked_list;

use graph::*;

fn main() -> Result<(), ()> {
    // let mut graph = Graph::new();
    //
    // graph.add(0, 1, 4);
    // graph.add(0, 7, 8);
    //
    // graph.add(1, 2, 8);
    // graph.add(1, 7, 11);
    //
    // graph.add(2, 8, 2);
    // graph.add(2, 3, 7);
    // graph.add(2, 5, 4);
    //
    // graph.add(7, 8, 7);
    // graph.add(7, 6, 1);
    //
    // graph.add(6, 8, 6);
    // graph.add(6, 5, 2);
    //
    // graph.add(5, 3, 14);
    // graph.add(5, 4, 10);
    //
    // graph.add(4, 3, 9);
    //
    // println!("{:?}", graph.dejkstra(0));
    //
    let mut graph = Graph::new();

    graph.add(0, 1, 4);
    graph.add(0, 7, 8);

    graph.add(1, 2, 8);
    graph.add(1, 7, 11);

    graph.add(2, 3, 7);
    graph.add(2, 5, 4);
    graph.add(2, 8, 2);

    graph.add(3, 4, 9);
    graph.add(3, 5, 14);

    graph.add(4, 5, 10);

    graph.add(5, 6, 2);

    graph.add(6, 8, 6);
    graph.add(6, 7, 1);

    graph.add(7, 8, 7);

    println!("graph:\n{}\n", graph);

    println!("dejkstra from 0: {:?}\n", graph.dejkstra(0));

    println!("kruskal:\n{}\n", graph.kruskal());

    println!("prim:\n{}\n", graph.prim_ull());

    Ok(())
}
