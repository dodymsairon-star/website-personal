<?php
// Array asosiatif 4 mahasiswa
$mahasiswa = [
    ["nim" => "24001", "nama" => "Alya", "nilai" => 85],
    ["nim" => "24002", "nama" => "Bima", "nilai" => 70],
    ["nim" => "24003", "nama" => "Citra", "nilai" => 92],
    ["nim" => "24004", "nama" => "Danu", "nilai" => 60]
];

echo "<h3>Daftar Mahasiswa</h3>";

// Menggunakan foreach
foreach ($mahasiswa as $mhs) {
    echo "NIM: " . $mhs["nim"] . " | ";
    echo "Nama: " . $mhs["nama"] . " | ";
    echo "Nilai: " . $mhs["nilai"] . "<br>";
}
?>