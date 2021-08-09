use std::io::{self, Write};

fn keyboard_input() -> String {
    let mut buffer = String::new();
    let _outputs = io::stdin().read_line(&mut buffer).unwrap();
    return buffer.trim().to_string();
}

fn main() {
    print!("Enter input: ");
    let _ = io::stdout().flush();
    let input = keyboard_input().to_string();
    println!("\nInput: {}", input);
}
