<?php
$tugas = 80;
$uts = 75;
$uas = 90;

// Menghitung nilai akhir
$nilaiAkhir = ($tugas * 0.30) + ($uts * 0.30) + ($uas * 0.40);

echo "<h3>Perhitungan Nilai</h3>";
echo "Nilai Tugas: $tugas (Bobot 30%)<br>";
echo "Nilai UTS: $uts (Bobot 30%)<br>";
echo "Nilai UAS: $uas (Bobot 40%)<br>";
echo "<hr>";
echo "<b>Nilai Akhir: $nilaiAkhir</b>";
?>