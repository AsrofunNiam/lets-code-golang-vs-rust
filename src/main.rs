use std::time::Instant;
use std::mem;
fn main() { 

    // Inisialisasi timer untuk memantau waktu
    let start = Instant::now();
    
    // Mengalokasikan sejumlah besar data
    let data = LargeData {
        data: vec![0; 10_000_000], // Alokasikan 10 juta elemen
        };
    let data_size = data.data.len() * std::mem::size_of::<u8>();

    // Ukur ukuran objek data menggunakan size_of_val
    let size = mem::size_of_val(&data);
    let vec_size = mem::size_of_val(&data.data);


    println!("Ukuran data yang dialokasikan (struktur): {} bytes", size); 
    println!("Ukuran data dalam Vec<u8>: {} bytes", vec_size);
    println!("Data dialokasikan setelah {:?}", start.elapsed());
    println!("Ukuran memori yang digunakan oleh elemen-elemen Vec<u8>: {} bytes", data_size);



}

#[derive(Debug)]
struct LargeData {
    data: Vec<u8>,
}
