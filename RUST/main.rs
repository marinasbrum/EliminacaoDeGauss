use std::env;
use std::time::{Instant, SystemTime, UNIX_EPOCH};
use rand::{Rng, SeedableRng};
use rand::rngs::StdRng;
use std::process::exit;

const MAXN: usize = 2000; 

fn time_seed() -> u32 {
    let start = Instant::now();
    start.elapsed().as_micros() as u32
}

fn parameters(argc: usize, argv: Vec<String>) -> (usize, u64) {
    let mut seed: u64 = 0;
    let mut n: usize = 0;

    if argc >= 2 {
        n = argv[1].parse().unwrap_or_else(|_| {
            println!("N é inválido, usando o valor padrão de 10.");
            10 
        });
        if n < 1 || n > MAXN {
            println!("N = {} está fora do intervalo permitido. Usando N = 10.", n);
            n = 10;
        }
    } else {
        println!("Usage: <matrix_dimension> [random seed]");
        exit(0);
    }
    println!("Matrix dimension N = {}", n);

    if argc == 3 {
        seed = argv[2].parse().unwrap_or_else(|_| {
            let generated_seed = time_seed() as u64;
            println!("Seed inválida, usando seed gerada automaticamente: {}", generated_seed);
            generated_seed
        });
        println!("Random seed = {}", seed);
    } else {
        seed = time_seed() as u64; 
    }

    (n, seed)
}

fn initialize_inputs(n: usize, seed: u64) -> (Vec<Vec<f32>>, Vec<f32>) {
    let mut rng = StdRng::seed_from_u64(seed); 
    let mut a: Vec<Vec<f32>> = vec![vec![0.0; n]; n];
    let mut b: Vec<f32> = vec![0.0; n];

    println!("\nInitializing...");

    for col in 0..n {
        for row in 0..n {
            a[row][col] = rng.gen_range(0.0..1.0); 
        }
        b[col] = rng.gen_range(0.0..1.0); 
    }
    (a, b)
}

fn print_inputs(a: &Vec<Vec<f32>>, b: &Vec<f32>, n: usize) {
    if n < 10 {
        println!("\nA =");
        for row in 0..n {
            for col in 0..n {
                print!("{:5.2}{}", a[row][col], if col < n - 1 { ", " } else { ";\n\t" });
            }
        }
        print!("\nB = [");
        for col in 0..n {
            print!("{:5.2}{}", b[col], if col < n - 1 { "; " } else { "]\n" });
        }
    }
}

fn print_x(x: &Vec<f32>, n: usize) {
    if n < 100 {
        print!("\nX = [");
        for row in 0..n {
            print!("{:5.2}{}", x[row], if row < n - 1 { "; " } else { "]\n" });
        }
    }
}

fn gauss(n: usize, a: &mut Vec<Vec<f32>>, b: &mut Vec<f32>, x: &mut Vec<f32>) {
    println!("Computing Serially.");

    for norm in 0..n - 1 {
        for row in norm + 1..n {
            let multiplier = a[row][norm] / a[norm][norm];
            for col in norm..n {
                a[row][col] -= a[norm][col] * multiplier;
            }
            b[row] -= b[norm] * multiplier;
        }
    }
    for row in (0..n).rev() {
        x[row] = b[row];
        for col in (row + 1..n).rev() {
            x[row] -= a[row][col] * x[col];
        }
        x[row] /= a[row][row];
    }
}

fn main() {
    let args: Vec<String> = env::args().collect();
    let argc = args.len();
    let (n, seed) = parameters(argc, args);
    let (mut a, mut b) = initialize_inputs(n, seed);
    let mut x = vec![0.0; n]; 

    print_inputs(&a, &b, n);
    println!("\nStarting clock.");

    let etstart = SystemTime::now();
    let etstart2 = Instant::now();

    gauss(n, &mut a, &mut b, &mut x);

    let etstop = SystemTime::now();
    let etstop2 = Instant::now();

    println!("Stopped clock.");

    let usecstart = etstart.duration_since(UNIX_EPOCH).unwrap().as_micros();
    let usecstop = etstop.duration_since(UNIX_EPOCH).unwrap().as_micros();

    print_x(&x, n);

    println!("\nElapsed time = {} ms.", (usecstop - usecstart) as f32 / 1000.0);

    let cpu_time_ms = (etstop2 - etstart2).as_millis();
    println!("(CPU times are accurate to the nearest {} ms)", 1.0 / 1000.0);
    println!("My total CPU time = {} ms.", cpu_time_ms);

    println!("--------------------------------------------");
}
