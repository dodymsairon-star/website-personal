<?php
// Fungsi 1: Luas Persegi Panjang
function hitungLuasPersegiPanjang($panjang, $lebar) {
    return $panjang * $lebar;
}

// Fungsi 2: Kelulusan
function tentukanKelulusan($nilai) {
    if ($nilai >= 70) {
        return "Lulus";
    } else {
        return "Belum Lulus";
    }
}

echo "<h3>1. Hitung Luas Persegi Panjang</h3>";
echo "Luas (5 x 10): " . hitungLuasPersegiPanjang(5, 10) . "<br>";
echo "Luas (3 x 7): " . hitungLuasPersegiPanjang(3, 7) . "<br>";
echo "Luas (10 x 10): " . hitungLuasPersegiPanjang(10, 10) . "<hr>";

echo "<h3>2. Status Kelulusan</h3>";
echo "Nilai 80: " . tentukanKelulusan(80) . "<br>";
echo "Nilai 65: " . tentukanKelulusan(65) . "<br>";
echo "Nilai 70: " . tentukanKelulusan(70);
?>