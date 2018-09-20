extern crate serde_json;

use std::env;
use std::fs::File;
use std::io::prelude::*;
use std::time::Instant;

use serde_json::Value;

fn main() {
    let fpath = env::args().skip(1).next().unwrap();
    let mut f = File::open(&fpath).expect("no file");
    let mut content = String::new();
    f.read_to_string(&mut content).unwrap();
    let beg = Instant::now();
    for _ in 0..(1e5 as i32) {
        let _v: Value = serde_json::from_str(&content).unwrap();
    }
    let elapsed = beg.elapsed();
    println!("serde_json::from_str: {}.{}s", elapsed.as_secs(), elapsed.subsec_millis());
}
