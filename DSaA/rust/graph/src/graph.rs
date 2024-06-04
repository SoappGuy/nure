#![allow(unused)]

use crate::linked_list::LinkedList;
use std::{fmt::Display, usize};

#[derive(Debug)]
pub struct EdgeTo {
    pub header: usize,
    pub length: i32,
}

#[derive(Debug)]
pub struct Graph {
    pub nodes: Vec<Vec<EdgeTo>>,
    pub edges: Vec<(usize, usize, i32)>,
}

impl Graph {
    pub fn new() -> Self {
        Self {
            nodes: vec![],
            edges: vec![],
        }
    }

    pub fn add(&mut self, vertex1: usize, vertex2: usize, length: i32) {
        while vertex1 >= self.nodes.len() || vertex2 >= self.nodes.len() {
            self.nodes.push(vec![])
        }

        self.nodes[vertex1].push(EdgeTo {
            header: vertex2,
            length,
        });

        self.nodes[vertex2].push(EdgeTo {
            header: vertex1,
            length,
        });

        self.edges.push((vertex1, vertex2, length));
    }

    pub fn dejkstra(&self, start: usize) -> Vec<i32> {
        fn min_dist(distances: &[i32], visited: &[bool]) -> isize {
            let mut min_len: i32 = i32::MAX;
            let mut header: isize = -1;

            for node in (0..(distances.len())) {
                if (!visited[node] && distances[node] <= min_len) {
                    min_len = distances[node];
                    header = node as isize;
                }
            }

            header
        }

        let n: usize = self.nodes.len();

        let mut distances: Vec<i32> = vec![i32::MAX; n];
        let mut visited: Vec<bool> = vec![false; n];

        distances[start] = 0;

        for count in (0..(n - 1)) {
            let curr: usize = min_dist(&distances, &visited) as usize;
            visited[curr] = true;

            for EdgeTo { header, length } in &self.nodes[curr] {
                if (!visited[*header]
                    && distances[curr] != -1
                    && distances[curr] + length < distances[*header])
                {
                    distances[*header] = distances[curr] + length;
                }
            }
        }

        distances
    }

    pub fn kruskal(&mut self) -> Graph {
        fn find_parent(parents: &mut Vec<usize>, x: usize) -> usize {
            if (parents[x] != x) {
                parents[x] = find_parent(parents, parents[x]);
            }

            parents[x]
        }

        let mut mst = Graph::new();
        let mut parents: Vec<usize> = (0..(self.nodes.len())).collect();

        self.edges.sort_by(|a, b| (a.2).cmp(&b.2));

        for (a, b, len) in &self.edges {
            let parent_a = find_parent(&mut parents, *a);
            let parent_b = find_parent(&mut parents, *b);

            if parent_a != parent_b {
                mst.add(*a, *b, *len);
                parents[parent_a] = parent_b;
            }
        }

        mst
    }

    pub fn prim_ull(&self) -> Graph {
        let mut mst = Graph::new();
        let mut edges: LinkedList = LinkedList::new();

        for edge in &self.nodes[0] {
            edges.add((0, edge.header, edge.length));
        }

        let n: usize = self.nodes.len();
        let mut visited: Vec<bool> = vec![false; n];

        visited[0] = true;

        while mst.edges.len() < n - 1 {
            let mut shortest_edge = edges.pop_min().unwrap();

            let from: usize = shortest_edge.0;
            let to: usize = shortest_edge.1;
            let len: i32 = shortest_edge.2;

            if visited[to] {
                continue;
            }

            visited[to] = true;
            mst.add(from, to, len);

            for edge in &self.nodes[to] {
                if !visited[edge.header] {
                    edges.add((to, edge.header, edge.length))
                }
            }
        }

        mst
    }
}

impl Display for Graph {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        for (from, to, len) in &self.edges {
            writeln!(f, "{from} 󰍴󰍴{{{len}}}󰍴󰁔 {to}");
        }

        Ok(())
    }
}
