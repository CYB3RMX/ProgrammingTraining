fn say_name(name:String){
    println!("Hello: {}", name);
}

fn return_number(num1:i32, num2:i32) -> i32 {
    return num1+num2;
}

fn print_all(array:[i32;5]){
    for element in array.iter(){
        println!("Index: {}", element);
    }
}

fn main(){
    // Basic sum of two numbers
    let number1 = 5;
    let number2 = 6;
    let number3 = return_number(number1, number2);
    println!("Sum: {}", number3);

    // Say name
    say_name("CYB3RMX".to_string());

    // Print all
    let myarr:[i32;5] = [10, 20, 30, 40, 50];
    print_all(myarr);
}