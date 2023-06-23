use std::fs::File;
use std::io::{BufRead, BufReader};
use std::time::Instant;

fn main() {
    let file = File::open("/home/yrsdi/Downloads/1000000-sales-records.txt").expect("Failed to open file");
    let reader = BufReader::new(file);

    let start_time = Instant::now();

    for line in reader.lines().take(1_000_000) {
        if let Ok(record) = line {
            // Process the record here
            println!("{}", record);
        }
    }

    let elapsed_time = start_time.elapsed();
    println!("Elapsed time: {:?}", elapsed_time);
}
