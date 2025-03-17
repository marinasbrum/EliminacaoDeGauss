use std::env;
use std::time::{Instant, SystemTime, UNIX_EPOCH};
use std::process::exit;

const MAXN: usize = 2000; // Valor máximo de N

// Função de eliminação de Gauss
fn gauss(n: usize, a: &mut Vec<Vec<f32>>, b: &mut Vec<f32>, x: &mut Vec<f32>) {
    println!("Computing Serially.");

    // Eliminação de Gauss
    for norm in 0..n - 1 {
        for row in norm + 1..n {
            let multiplier = a[row][norm] / a[norm][norm];
            for col in norm..n {
                a[row][col] -= a[norm][col] * multiplier;
            }
            b[row] -= b[norm] * multiplier;
        }
    }

    // Substituição retroativa
    for row in (0..n).rev() {
        x[row] = b[row];
        for col in (row + 1..n).rev() {
            x[row] -= a[row][col] * x[col];
        }
        x[row] /= a[row][row];
    }
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

fn parameters(argc: usize, argv: Vec<String>) -> usize {
    let mut n: usize = 0;

    // Se o número de argumentos for suficiente
    if argc >= 2 {
        // Tenta converter o primeiro argumento para o tamanho da matriz
        n = argv[1].parse().unwrap_or_else(|_| {
            println!("N é inválido, usando o valor padrão de 5.");
            5 // valor padrão para n
        });
        // Se N estiver fora dos limites, ajusta para 5
        if n < 1 || n > MAXN {
            println!("N = {} está fora do intervalo permitido. Usando N = 5.", n);
            n = 5;
        }
    } else {
        println!("Usage: <matrix_dimension>");
        exit(0);
    }

    println!("\nMatrix dimension N = {}.", n);
    n
}

fn main() {
    let args: Vec<String> = env::args().collect();
    let argc = args.len();

    // Processa os parâmetros e obtém o tamanho da matriz
    let n = parameters(argc, args);

    // Inicializa as variáveis com os valores fixos
    let mut a: Vec<Vec<f32>> = vec![vec![0.0; n]; n];
    let mut b: Vec<f32> = vec![0.0; n];
    let mut x: Vec<f32> = vec![0.0; n];

    // Preenche a matriz A e o vetor B com valores fixos para a demonstração
    if n == 5 {
        a = vec![
            vec![2.0, 1.0, 1.0, 0.0, 0.0],
            vec![1.0, 3.0, 0.0, 1.0, 0.0],
            vec![1.0, 0.0, 2.0, 1.0, 0.0],
            vec![0.0, 1.0, 1.0, 3.0, 1.0],
            vec![0.0, 0.0, 0.0, 1.0, 2.0],
        ];
        b = vec![5.0, 7.0, 6.0, 8.0, 4.0];
    }

    // Imprime as matrizes de entrada
    print_inputs(&a, &b, n);

    // Inicia a medição de tempo
    println!("\nStarting clock.");
    let etstart = SystemTime::now();
    let etstart2 = Instant::now();

    // Chama a função de eliminação de Gauss
    gauss(n, &mut a, &mut b, &mut x);

    // Para a medição de tempo
    let etstop = SystemTime::now();
    let etstop2 = Instant::now();
    println!("Stopped clock.");

    // Calcular o tempo de execução
    let usecstart = etstart.duration_since(UNIX_EPOCH).unwrap().as_micros();
    let usecstop = etstop.duration_since(UNIX_EPOCH).unwrap().as_micros();

    // Imprime o vetor solução X
    print_x(&x, n);

    // Exibe o tempo decorrido
    println!("\nElapsed time = {} ms.", (usecstop - usecstart) as f32 / 1000.0);

    // Exibe o tempo de CPU
    let cpu_time_ms = (etstop2 - etstart2).as_millis();
    println!("(CPU times are accurate to the nearest {} ms)", 1.0 / 1000.0);
    println!("My total CPU time = {} ms.", cpu_time_ms);

    println!("--------------------------------------------");
}
