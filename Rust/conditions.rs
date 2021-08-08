fn main(){
    // Basic if statement
    let cond_var:i32 = 100;
    if cond_var > 100 {
        println!("Greater than 100");
    } else if cond_var < 100 {
        println!("Less than 100");
    } else {
        println!("100");
    }

    // Match statement
    let _variable = match cond_var {
        50 => {
            println!("50");
        },
        90 => {
            println!("90");
        },
        100 => {
            println!("100");
        },
        _ => {
            println!("Default");
        }
    };
}