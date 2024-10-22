package main


import (
    "fmt" // Mengimpor package fmt untuk format input dan output
    "time" // Mengimpor package time untuk manipulasi tanggal dan waktu
)


const (
    maxPasien = 20 // Jumlah maksimum pasien yang dapat dikelola
    maxPaket  = 10 // Jumlah maksimum paket MCU yang dapat dikelola
)


type PaketMCU struct {
    NamaPaket string // Nama paket MCU
    Harga     int    // Harga paket MCU
}


type Pasien struct {
    Nama       string    // Nama pasien
    JenisPaket string    // Jenis paket MCU yang diambil pasien
    TanggalMCU time.Time // Tanggal MCU
    RekapHasil string    // Rekap hasil pemeriksaan MCU
}


var (
    paketMCUs [maxPaket]PaketMCU // Array untuk menyimpan data paket MCU
    pasiens   [maxPasien]Pasien  // Array untuk menyimpan data pasien
    jumPaket  int                // Jumlah paket MCU yang terdaftar
    jumPasien int                // Jumlah pasien yang terdaftar
)


func main() {
    for {
        // Menampilkan menu utama
        fmt.Println("MENU UTAMA\n1. Kelola Paket Medical Check Up\n2. Tampilkan Seluruh Jenis Paket yang Tersimpan\n3. Kelola Data Pasien MCU\n4. Cari/Tampilkan Data Pasien\n5. Rekap Hasil\n6. Tampilkan Total Pemasukan Layanan MCU pada Periode Waktu Tertentu\n0. Keluar")
        var pilihan int
        fmt.Print("Pilihan: ") // Meminta input pilihan dari pengguna
        fmt.Scan(&pilihan) // Membaca input pilihan dari pengguna
        switch pilihan { // Struktur kontrol untuk mengeksekusi fungsi berdasarkan pilihan pengguna
        case 1:
            menuKelolaPaket() // Memanggil fungsi untuk mengelola paket MCU
        case 2:
            tampilkanPaket() // Memanggil fungsi untuk menampilkan semua paket MCU
        case 3:
            menuKelolaPasien() // Memanggil fungsi untuk mengelola data pasien
        case 4:
            menuCariTampilkanPasien() // Memanggil fungsi untuk mencari dan menampilkan data pasien
        case 5:
            menuRekapHasil() // Memanggil fungsi untuk mengelola rekap hasil
        case 6:
            tampilkanTotalPemasukan() // Memanggil fungsi untuk menampilkan total pemasukan layanan MCU
        case 0:
            return // Keluar dari program
        default:
            fmt.Println("Pilihan tidak valid.") // Menampilkan pesan jika pilihan tidak valid
        }
    }
}


func menuKelolaPaket() {
    // Menampilkan menu untuk mengelola paket MCU
    fmt.Println("1. Tambah Paket\n2. Edit Data Paket\n3. Hapus Paket\n0. Kembali")
    var pilihan int
    fmt.Print("Pilihan: ") // Meminta input pilihan dari pengguna
    fmt.Scan(&pilihan) // Membaca input pilihan dari pengguna
    switch pilihan { // Struktur kontrol untuk mengeksekusi fungsi berdasarkan pilihan pengguna
    case 1:
        tambahPaket() // Memanggil fungsi untuk menambah paket MCU
    case 2:
        editPaket() // Memanggil fungsi untuk mengedit paket MCU
    case 3:
        hapusPaket() // Memanggil fungsi untuk menghapus paket MCU
    case 0:
        return // Kembali ke menu utama
    default:
        fmt.Println("Pilihan tidak valid.") // Menampilkan pesan jika pilihan tidak valid
    }
}


func tambahPaket() {
    if jumPaket >= maxPaket { // Mengecek apakah jumlah paket sudah mencapai batas maksimum
        fmt.Println("Jumlah paket sudah maksimal.")
        return
    }
    fmt.Print("Nama Paket: ") // Meminta input nama paket dari pengguna
    fmt.Scan(&paketMCUs[jumPaket].NamaPaket) // Membaca input nama paket dari pengguna
    fmt.Print("Harga: ") // Meminta input harga paket dari pengguna
    fmt.Scan(&paketMCUs[jumPaket].Harga) // Membaca input harga paket dari pengguna
    jumPaket++ // Menambah jumlah paket yang terdaftar
    fmt.Println("Paket berhasil ditambah.")
}


func editPaket() {
    var nama string
    fmt.Print("Nama Paket yang akan diedit: ") // Meminta input nama paket yang akan diedit dari pengguna
    fmt.Scan(&nama) // Membaca input nama paket dari pengguna
    idx := sequentialSearchPaket(nama) // Mencari indeks paket berdasarkan nama
    if idx == -1 { // Mengecek apakah paket ditemukan
        fmt.Println("Paket tidak ditemukan.")
        return
    }
    fmt.Print("Nama Paket baru: ") // Meminta input nama paket baru dari pengguna
    fmt.Scan(&paketMCUs[idx].NamaPaket) // Membaca input nama paket baru dari pengguna
    fmt.Print("Harga baru: ") // Meminta input harga baru dari pengguna
    fmt.Scan(&paketMCUs[idx].Harga) // Membaca input harga baru dari pengguna
    fmt.Println("Paket berhasil diedit.")
}


func hapusPaket() {
    var nama string
    fmt.Print("Nama Paket yang akan dihapus: ") // Meminta input nama paket yang akan dihapus dari pengguna
    fmt.Scan(&nama) // Membaca input nama paket dari pengguna
    idx := sequentialSearchPaket(nama) // Mencari indeks paket berdasarkan nama
    if idx == -1 { // Mengecek apakah paket ditemukan
        fmt.Println("Paket tidak ditemukan.")
        return
    }
    for i := idx; i < jumPaket-1; i++ { // Menggeser paket yang ada setelah paket yang dihapus
        paketMCUs[i] = paketMCUs[i+1]
    }
    jumPaket-- // Mengurangi jumlah paket yang terdaftar
    fmt.Println("Paket berhasil dihapus.")
}


func sequentialSearchPaket(nama string) int {
    // Mencari paket berdasarkan nama menggunakan sequential search
    for i := 0; i < jumPaket; i++ {
        if paketMCUs[i].NamaPaket == nama { // Mengecek apakah nama paket sesuai
            return i // Mengembalikan indeks jika ditemukan
        }
    }
    return -1 // Mengembalikan -1 jika tidak ditemukan
}


func tampilkanPaket() {
    // Menampilkan semua paket MCU yang terdaftar
    for i := 0; i < jumPaket; i++ {
        fmt.Printf("Nama Paket: %s, Harga: %d\n", paketMCUs[i].NamaPaket, paketMCUs[i].Harga)
    }
}


func menuKelolaPasien() {
    // Menampilkan menu untuk mengelola data pasien
    fmt.Println("1. Tambah Pasien\n2. Edit Data Pasien\n3. Hapus Pasien\n0. Kembali")
    var pilihan int
    fmt.Print("Pilihan: ") // Meminta input pilihan dari pengguna
    fmt.Scan(&pilihan) // Membaca input pilihan dari pengguna
    switch pilihan { // Struktur kontrol untuk mengeksekusi fungsi berdasarkan pilihan pengguna
    case 1:
        tambahPasien() // Memanggil fungsi untuk menambah pasien
    case 2:
        editPasien() // Memanggil fungsi untuk mengedit pasien
    case 3:
        hapusPasien() // Memanggil fungsi untuk menghapus pasien
    case 0:
        return // Kembali ke menu utama
    default:
        fmt.Println("Pilihan tidak valid.") // Menampilkan pesan jika pilihan tidak valid
    }
}


func tambahPasien() {
    if jumPasien >= maxPasien { // Mengecek apakah jumlah pasien sudah mencapai batas maksimum
        fmt.Println("Jumlah pasien sudah maksimal.")
        return
    }
    fmt.Print("Nama: ") // Meminta input nama pasien dari pengguna
    fmt.Scan(&pasiens[jumPasien].Nama) // Membaca input nama pasien dari pengguna
    fmt.Print("Jenis Paket: ") // Meminta input jenis paket dari pengguna
    fmt.Scan(&pasiens[jumPasien].JenisPaket) // Membaca input jenis paket dari pengguna
    fmt.Print("Tanggal MCU (YYYY-MM-DD): ") // Meminta input tanggal MCU dari pengguna
    var tgl string
    fmt.Scan(&tgl) // Membaca input tanggal MCU dari pengguna
    pasiens[jumPasien].TanggalMCU, _ = time.Parse("2006-01-02", tgl) // Mengonversi string ke tipe time.Time
    jumPasien++ // Menambah jumlah pasien yang terdaftar
    fmt.Println("Pasien berhasil ditambah.")
}


func editPasien() {
    var nama string
    fmt.Print("Nama Pasien yang akan diedit: ") // Meminta input nama pasien yang akan diedit dari pengguna
    fmt.Scan(&nama) // Membaca input nama pasien dari pengguna
    idx := sequentialSearchPasien(nama) // Mencari indeks pasien berdasarkan nama
    if idx == -1 { // Mengecek apakah pasien ditemukan
        fmt.Println("Pasien tidak ditemukan.")
        return
    }
    fmt.Print("Nama baru: ") // Meminta input nama baru dari pengguna
    fmt.Scan(&pasiens[idx].Nama) // Membaca input nama baru dari pengguna
    fmt.Print("Jenis Paket baru: ") // Meminta input jenis paket baru dari pengguna
    fmt.Scan(&pasiens[idx].JenisPaket) // Membaca input jenis paket baru dari pengguna
    fmt.Print("Tanggal MCU baru (YYYY-MM-DD): ") // Meminta input tanggal MCU baru dari pengguna
    var tgl string
    fmt.Scan(&tgl) // Membaca input tanggal MCU dari pengguna
    pasiens[idx].TanggalMCU, _ = time.Parse("2006-01-02", tgl) // Mengonversi string ke tipe time.Time
    fmt.Println("Data pasien berhasil diedit.")
}


func hapusPasien() {
    var nama string
    fmt.Print("Nama Pasien yang akan dihapus: ") // Meminta input nama pasien yang akan dihapus dari pengguna
    fmt.Scan(&nama) // Membaca input nama pasien dari pengguna
    idx := sequentialSearchPasien(nama) // Mencari indeks pasien berdasarkan nama
    if idx == -1 { // Mengecek apakah pasien ditemukan
        fmt.Println("Pasien tidak ditemukan.")
        return
    }
    for i := idx; i < jumPasien-1; i++ { // Menggeser pasien yang ada setelah pasien yang dihapus
        pasiens[i] = pasiens[i+1]
    }
    jumPasien-- // Mengurangi jumlah pasien yang terdaftar
    fmt.Println("Pasien berhasil dihapus.")
}


func sequentialSearchPasien(nama string) int {
    // Mencari pasien berdasarkan nama menggunakan sequential search
    for i := 0; i < jumPasien; i++ {
        if pasiens[i].Nama == nama { // Mengecek apakah nama pasien sesuai
            return i // Mengembalikan indeks jika ditemukan
        }
    }
    return -1 // Mengembalikan -1 jika tidak ditemukan
}


func menuCariTampilkanPasien() {
    // Menampilkan menu untuk mencari dan menampilkan data pasien
    fmt.Println("1. Cari pasien berdasarkan nama\n2. Tampilkan Data Pasien Terurut berdasarkan Waktu MCU\n3. Tampilkan Data Pasien Terurut berdasarkan waktu MCU dan jenis paket\n4. Sortir pasien berdasarkan paket MCU\n5. Sortir pasien berdasarkan periode MCU\n0. Kembali")
    var pilihan int
    fmt.Print("Pilihan: ") // Meminta input pilihan dari pengguna
    fmt.Scan(&pilihan) // Membaca input pilihan dari pengguna
    switch pilihan { // Struktur kontrol untuk mengeksekusi fungsi berdasarkan pilihan pengguna
    case 1:
        cariPasien() // Memanggil fungsi untuk mencari pasien berdasarkan nama
    case 2:
        tampilkanPasienBerdasarkanWaktuMCU() // Memanggil fungsi untuk menampilkan pasien berdasarkan waktu MCU
    case 3:
        tampilkanPasienBerdasarkanWaktuDanPaketMCU() // Memanggil fungsi untuk menampilkan pasien berdasarkan waktu dan paket MCU
    case 4:
        sortirPasienBerdasarkanPaketMCU() // Memanggil fungsi untuk menyortir pasien berdasarkan paket MCU
    case 5:
        sortirPasienBerdasarkanPeriodeMCU() // Memanggil fungsi untuk menyortir pasien berdasarkan periode MCU
    case 0:
        return // Kembali ke menu utama
    default:
        fmt.Println("Pilihan tidak valid.") // Menampilkan pesan jika pilihan tidak valid
    }
}


func cariPasien() {
    var nama string
    fmt.Print("Masukkan nama pasien: ") // Meminta input nama pasien dari pengguna
    fmt.Scan(&nama) // Membaca input nama pasien dari pengguna
    idx := sequentialSearchPasien(nama) // Mencari indeks pasien berdasarkan nama
    if idx == -1 { // Mengecek apakah pasien ditemukan
        fmt.Println("Pasien tidak ditemukan.")
        return
    }
    // Menampilkan data pasien jika ditemukan
    fmt.Printf("Nama: %s, Paket: %s, Tanggal MCU: %s, Rekap Hasil: %s\n", pasiens[idx].Nama, pasiens[idx].JenisPaket, pasiens[idx].TanggalMCU.Format("2006-01-02"), pasiens[idx].RekapHasil)
}


func tampilkanPasienBerdasarkanWaktuMCU() {
    var order string
    fmt.Print("Masukkan urutan (asc/desc): ") // Meminta input urutan (asc/desc) dari pengguna
    fmt.Scan(&order) // Membaca input urutan dari pengguna
    if order == "asc" { // Mengecek apakah urutan adalah ascending
        insertionSortPasienByTanggalMCUAsc() // Memanggil fungsi untuk mengurutkan pasien berdasarkan tanggal MCU secara ascending
    } else if order == "desc" { // Mengecek apakah urutan adalah descending
        insertionSortPasienByTanggalMCUDesc() // Memanggil fungsi untuk mengurutkan pasien berdasarkan tanggal MCU secara descending
    } else {
        fmt.Println("Urutan tidak valid.") // Menampilkan pesan jika urutan tidak valid
        return
    }
    tampilkanPasien() // Menampilkan data pasien yang sudah diurutkan
}


func insertionSortPasienByTanggalMCUAsc() {
    // Mengurutkan pasien berdasarkan tanggal MCU secara ascending menggunakan insertion sort
    for i := 1; i < jumPasien; i++ {
        key := pasiens[i]
        j := i - 1
        for j >= 0 && pasiens[j].TanggalMCU.After(key.TanggalMCU) {
            pasiens[j+1] = pasiens[j]
            j--
        }
        pasiens[j+1] = key
    }
}


func insertionSortPasienByTanggalMCUDesc() {
    // Mengurutkan pasien berdasarkan tanggal MCU secara descending menggunakan insertion sort
    for i := 1; i < jumPasien; i++ {
        key := pasiens[i]
        j := i - 1
        for j >= 0 && pasiens[j].TanggalMCU.Before(key.TanggalMCU) {
            pasiens[j+1] = pasiens[j]
            j--
        }
        pasiens[j+1] = key
    }
}


func tampilkanPasien() {
    // Menampilkan data semua pasien yang terdaftar
    for i := 0; i < jumPasien; i++ {
        fmt.Printf("Nama: %s, Paket: %s, Tanggal MCU: %s, Rekap Hasil: %s\n", pasiens[i].Nama, pasiens[i].JenisPaket, pasiens[i].TanggalMCU.Format("2006-01-02"), pasiens[i].RekapHasil)
    }
}


func tampilkanPasienBerdasarkanWaktuDanPaketMCU() {
    var order string
    fmt.Print("Masukkan urutan (asc/desc): ") // Meminta input urutan (asc/desc) dari pengguna
    fmt.Scan(&order) // Membaca input urutan dari pengguna
    if order == "asc" { // Mengecek apakah urutan adalah ascending
        insertionSortPasienByTanggalDanPaketMCUAsc() // Memanggil fungsi untuk mengurutkan pasien berdasarkan tanggal dan paket MCU secara ascending
    } else if order == "desc" { // Mengecek apakah urutan adalah descending
        insertionSortPasienByTanggalDanPaketMCUDesc() // Memanggil fungsi untuk mengurutkan pasien berdasarkan tanggal dan paket MCU secara descending
    } else {
        fmt.Println("Urutan tidak valid.") // Menampilkan pesan jika urutan tidak valid
        return
    }
    tampilkanPasien() // Menampilkan data pasien yang sudah diurutkan
}


func insertionSortPasienByTanggalDanPaketMCUAsc() {
    // Mengurutkan pasien berdasarkan tanggal dan paket MCU secara ascending menggunakan insertion sort
    for i := 1; i < jumPasien; i++ {
        key := pasiens[i]
        j := i - 1
        for j >= 0 && (pasiens[j].TanggalMCU.After(key.TanggalMCU) || (pasiens[j].TanggalMCU.Equal(key.TanggalMCU) && pasiens[j].JenisPaket > key.JenisPaket)) {
            pasiens[j+1] = pasiens[j]
            j--
        }
        pasiens[j+1] = key
    }
}


func insertionSortPasienByTanggalDanPaketMCUDesc() {
    // Mengurutkan pasien berdasarkan tanggal dan paket MCU secara descending menggunakan insertion sort
    for i := 1; i < jumPasien; i++ {
        key := pasiens[i]
        j := i - 1
        for j >= 0 && (pasiens[j].TanggalMCU.Before(key.TanggalMCU) || (pasiens[j].TanggalMCU.Equal(key.TanggalMCU) && pasiens[j].JenisPaket < key.JenisPaket)) {
            pasiens[j+1] = pasiens[j]
            j--
        }
        pasiens[j+1] = key
    }
}


func sortirPasienBerdasarkanPaketMCU() {
    var paket string
    fmt.Print("Masukkan jenis paket MCU: ") // Meminta input jenis paket MCU dari pengguna
    fmt.Scan(&paket) // Membaca input jenis paket dari pengguna
    for i := 0; i < jumPasien; i++ {
        // Menampilkan data pasien yang sesuai dengan jenis paket MCU
        if pasiens[i].JenisPaket == paket {
            fmt.Printf("Nama: %s, Paket: %s, Tanggal MCU: %s, Rekap Hasil: %s\n", pasiens[i].Nama, pasiens[i].JenisPaket, pasiens[i].TanggalMCU.Format("2006-01-02"), pasiens[i].RekapHasil)
        }
    }
}


func sortirPasienBerdasarkanPeriodeMCU() {
    var start, end string
    fmt.Print("Masukkan tanggal mulai (YYYY-MM-DD): ") // Meminta input tanggal mulai dari pengguna
    fmt.Scan(&start) // Membaca input tanggal mulai dari pengguna
    fmt.Print("Masukkan tanggal akhir (YYYY-MM-DD): ") // Meminta input tanggal akhir dari pengguna
    fmt.Scan(&end) // Membaca input tanggal akhir dari pengguna
    startDate, _ := time.Parse("2006-01-02", start) // Mengonversi string tanggal mulai ke tipe time.Time
    endDate, _ := time.Parse("2006-01-02", end) // Mengonversi string tanggal akhir ke tipe time.Time
    for i := 0; i < jumPasien; i++ {
        // Menampilkan data pasien yang tanggal MCU-nya berada dalam periode yang ditentukan
        if pasiens[i].TanggalMCU.After(startDate) && pasiens[i].TanggalMCU.Before(endDate) {
            fmt.Printf("Nama: %s, Paket: %s, Tanggal MCU: %s, Rekap Hasil: %s\n", pasiens[i].Nama, pasiens[i].JenisPaket, pasiens[i].TanggalMCU.Format("2006-01-02"), pasiens[i].RekapHasil)
        }
    }
}


func menuRekapHasil() {
    var nama string
    fmt.Print("Masukkan nama pasien: ") // Meminta input nama pasien dari pengguna
    fmt.Scan(&nama) // Membaca input nama pasien dari pengguna
    idx := sequentialSearchPasien(nama) // Mencari indeks pasien berdasarkan nama
    if idx == -1 { // Mengecek apakah pasien ditemukan
        fmt.Println("Pasien tidak ditemukan.")
        return
    }
    fmt.Printf("Nama: %s, Rekap Hasil: %s\n", pasiens[idx].Nama, pasiens[idx].RekapHasil) // Menampilkan data rekap hasil pasien jika ditemukan
    fmt.Println("1. Tambah rekap hasil\n2. Edit rekap hasil\n3. Hapus rekap hasil\n0. Kembali")
    var pilihan int
    fmt.Print("Pilihan: ") // Meminta input pilihan dari pengguna
    fmt.Scan(&pilihan) // Membaca input pilihan dari pengguna
    switch pilihan { // Struktur kontrol untuk mengeksekusi fungsi berdasarkan pilihan pengguna
    case 1:
        tambahRekapHasil(idx) // Memanggil fungsi untuk menambah rekap hasil
    case 2:
        editRekapHasil(idx) // Memanggil fungsi untuk mengedit rekap hasil
    case 3:
        hapusRekapHasil(idx) // Memanggil fungsi untuk menghapus rekap hasil
    case 0:
        return // Kembali ke menu sebelumnya
    default:
        fmt.Println("Pilihan tidak valid.") // Menampilkan pesan jika pilihan tidak valid
    }
}


func tambahRekapHasil(idx int) {
    fmt.Print("Masukkan rekap hasil: ") // Meminta input rekap hasil dari pengguna
    fmt.Scan(&pasiens[idx].RekapHasil) // Membaca input rekap hasil dari pengguna
    fmt.Println("Rekap hasil berhasil ditambah.") // Menampilkan pesan bahwa rekap hasil berhasil ditambah
}


func editRekapHasil(idx int) {
    fmt.Print("Masukkan rekap hasil baru: ") // Meminta input rekap hasil baru dari pengguna
    fmt.Scan(&pasiens[idx].RekapHasil) // Membaca input rekap hasil baru dari pengguna
    fmt.Println("Rekap hasil berhasil diedit.") // Menampilkan pesan bahwa rekap hasil berhasil diedit
}


func hapusRekapHasil(idx int) {
    pasiens[idx].RekapHasil = "" // Mengosongkan rekap hasil pasien
    fmt.Println("Rekap hasil berhasil dihapus.") // Menampilkan pesan bahwa rekap hasil berhasil dihapus
}


func tampilkanTotalPemasukan() {
    var start, end string
    var total int
    fmt.Print("Masukkan tanggal mulai (YYYY-MM-DD): ") // Meminta input tanggal mulai dari pengguna
    fmt.Scan(&start) // Membaca input tanggal mulai dari pengguna
    fmt.Print("Masukkan tanggal akhir (YYYY-MM-DD): ") // Meminta input tanggal akhir dari pengguna
    fmt.Scan(&end) // Membaca input tanggal akhir dari pengguna
    startDate, _ := time.Parse("2006-01-02", start) // Mengonversi string tanggal mulai ke tipe time.Time
    endDate, _ := time.Parse("2006-01-02", end) // Mengonversi string tanggal akhir ke tipe time.Time
    for i := 0; i < jumPasien; i++ {
        // Mengecek apakah tanggal MCU pasien berada dalam periode yang ditentukan
        if pasiens[i].TanggalMCU.After(startDate) && pasiens[i].TanggalMCU.Before(endDate) {
            for j := 0; j < jumPaket; j++ {
                // Menambahkan harga paket MCU ke total pemasukan
                if pasiens[i].JenisPaket == paketMCUs[j].NamaPaket {
                    total += paketMCUs[j].Harga
                }
            }
        }
    }
    fmt.Printf("Total pemasukan: %d\n", total) // Menampilkan total pemasukan
}