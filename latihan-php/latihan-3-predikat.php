<?php
$nilai = 74; // Coba ubah: 49, 50, 64, 65, 74, 75, 84, 85

// Menentukan predikat
if ($nilai >= 85) {
    $predikat = "A";
} elseif ($nilai >= 75) {
    $predikat = "B";
} elseif ($nilai >= 65) {
    $predikat = "C";
} elseif ($nilai >= 50) {
    $predikat = "D";
} else {
    $predikat = "E";
}

// Menentukan status
if ($nilai >= 70) {
    $status = "Lulus";
} else {
    $status = "Belum Lulus";
}

echo "<h3>Hasil Penilaian</h3>";
echo "Nilai Anda: <b>$nilai</b><br>";
echo "Predikat: <b>$predikat</b><br>";
echo "Status: <b>$status</b>";
?>