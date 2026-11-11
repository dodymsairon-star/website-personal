<?php
// Membuat variabel
$nama = "Budi Santoso";
$nim = "24001234";
$prodi = "Rekayasa Perangkat Lunak";
$semester = 2;
$ipk = 3.75;

// Menampilkan dengan echo
echo "<h3>Data Mahasiswa</h3>";
echo "Nama: " . $nama . "<br>";
echo "NIM: " . $nim . "<br>";
echo "Prodi: " . $prodi . "<br>";
echo "Semester: " . $semester . "<br>";
echo "IPK: " . $ipk . "<hr>";

// Memeriksa tipe data
echo "<h3>Tipe Data</h3>";
echo "Nama: "; var_dump($nama); echo "<br>";
echo "NIM: "; var_dump($nim); echo "<br>";
echo "Prodi: "; var_dump($prodi); echo "<br>";
echo "Semester: "; var_dump($semester); echo "<br>";
echo "IPK: "; var_dump($ipk);
?>