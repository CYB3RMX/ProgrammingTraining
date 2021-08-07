fn main(){
    println!("Hello world!!"); // Prints Hello world!!
    println!("My name is: {}", "CYB3RMX"); /* Formatted print */

    let integer1:i32 = 10; // Signed 32 bit integer [default]
    let integer2:u32 = 20; // Unsigned 32 bit integer
    let mystr:&str = "I am a string!!"; // String
    let boolean:bool = true; // Boolean => true or false
    let onechar:char = 'A'; // Char
    let floaaa:f64 = 10.00;  // Float number [f64 by default]
    
    print!("Number1: {} || Number2: {} || String: {}", integer1, integer2, mystr);

    let mut mutable_variable = 154;
    println!("Before mut: {}", mutable_variable);
    mutable_variable = 200;
    println!("After mut: {}", mutable_variable);
}