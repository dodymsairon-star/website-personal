-- Buat Database
CREATE DATABASE IF NOT EXISTS fintrack_db;
USE fintrack_db;

-- Tabel Transactions
CREATE TABLE IF NOT EXISTS transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    ref VARCHAR(20) UNIQUE NOT NULL,
    date DATE NOT NULL,
    description VARCHAR(255) NOT NULL,
    category VARCHAR(100) NOT NULL,
    amount DECIMAL(15, 2) NOT NULL,
    type ENUM('income', 'expense') NOT NULL,
    status ENUM('success', 'pending', 'failed') DEFAULT 'success',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_date (date),
    INDEX idx_type (type),
    INDEX idx_category (category),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Insert Data Dummy (Optional)
INSERT INTO transactions (ref, date, description, category, amount, type, status) VALUES
('TRX001', '2026-07-25', 'Gaji Bulanan', 'Gaji', 25000000.00, 'income', 'success'),
('TRX002', '2026-07-24', 'Belanja Groceries', 'Makanan', 850000.00, 'expense', 'success'),
('TRX003', '2026-07-23', 'Proyek Freelance', 'Freelance', 5000000.00, 'income', 'success'),
('TRX004', '2026-07-22', 'Bensin & Tol', 'Transportasi', 450000.00, 'expense', 'success'),
('TRX005', '2026-07-21', 'Belanja Online', 'Belanja', 1200000.00, 'expense', 'pending'),
('TRX006', '2026-07-20', 'Dividen Saham', 'Investasi', 2500000.00, 'income', 'success'),
('TRX007', '2026-07-19', 'Nonton Bioskop', 'Hiburan', 150000.00, 'expense', 'success'),
('TRX008', '2026-07-18', 'Listrik & Air', 'Tagihan', 750000.00, 'expense', 'success'),
('TRX009', '2026-07-17', 'Bonus Proyek', 'Freelance', 7500000.00, 'income', 'success'),
('TRX010', '2026-07-16', 'Makan Siang', 'Makanan', 125000.00, 'expense', 'failed'),
('TRX011', '2026-07-15', 'Internet Bulanan', 'Tagihan', 350000.00, 'expense', 'success'),
('TRX012', '2026-07-14', 'Return Investasi', 'Investasi', 2750000.00, 'income', 'success');