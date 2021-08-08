fn main(){
    let test_string:&str = &String::new();
    println!("String length: {}", test_string.len());

    let not_empty_str:&str = &String::from("Not a empty string!");
    println!("String length: {}", not_empty_str.len());

    // Convert to string
    let someint:i32 = 100;
    someint.to_string();
    println!("Integer: {}", someint);

    // String concat
    let mystr1 = "Hello";
    let mystr2 = "World!!";
    let mystr3 = format!("{} {}", mystr1, mystr2);
    println!("New string: {}", mystr3);

    // Tokens with vectors
    let splitter = "One,two,three,four";
    let tokens:Vec<&str> = splitter.split(",").collect();
    for tok in tokens{
        println!("Token: {}", tok);
    }
}