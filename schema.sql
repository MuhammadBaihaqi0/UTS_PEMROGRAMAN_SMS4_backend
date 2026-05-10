-- ============================================================
-- SCHEMA: Sistem Booking Lapangan Olahraga
-- Database: PostgreSQL (Supabase)
-- ============================================================

-- Buat tabel bookings
CREATE TABLE IF NOT EXISTS bookings (
    id              SERIAL PRIMARY KEY,
    nama_pemesan    VARCHAR(100)    NOT NULL,
    email           VARCHAR(150)    NOT NULL,
    no_telepon      VARCHAR(15)     NOT NULL,
    nama_lapangan   VARCHAR(100)    NOT NULL,
    jenis_olahraga  VARCHAR(50)     NOT NULL,
    tanggal_main    DATE            NOT NULL,
    jam_mulai       TIME            NOT NULL,
    jam_selesai     TIME            NOT NULL,
    total_harga     DECIMAL(12, 2)  NOT NULL DEFAULT 0,
    status          VARCHAR(20)     NOT NULL DEFAULT 'pending'
                        CHECK (status IN ('pending', 'confirmed', 'cancelled', 'completed')),
    catatan         TEXT            DEFAULT '',
    created_at      TIMESTAMP       DEFAULT NOW(),
    updated_at      TIMESTAMP       DEFAULT NOW()
);

-- ============================================================
-- SEED DATA: 10 data contoh
-- ============================================================

INSERT INTO bookings (nama_pemesan, email, no_telepon, nama_lapangan, jenis_olahraga, tanggal_main, jam_mulai, jam_selesai, total_harga, status, catatan) VALUES
('Budi Santoso',    'budi.santoso@gmail.com',    '081234567890', 'Lapangan A',  'Badminton',  '2026-05-12', '08:00', '10:00', 100000,  'confirmed',  'Bawa raket sendiri'),
('Siti Rahayu',     'siti.rahayu@yahoo.com',     '082345678901', 'Lapangan B',  'Futsal',     '2026-05-12', '10:00', '12:00', 200000,  'pending',    'Butuh bola'),
('Ahmad Fauzi',     'ahmad.fauzi@gmail.com',     '083456789012', 'Lapangan C',  'Basket',     '2026-05-13', '13:00', '15:00', 150000,  'confirmed',  ''),
('Rina Wulandari',  'rina.wulandari@outlook.com','084567890123', 'Lapangan A',  'Badminton',  '2026-05-13', '15:00', '17:00', 100000,  'completed',  'Reguler mingguan'),
('Dedi Kurniawan',  'dedi.kurniawan@gmail.com',  '085678901234', 'Lapangan D',  'Tenis',      '2026-05-14', '07:00', '09:00', 180000,  'confirmed',  'Sewa raket'),
('Maya Putri',      'maya.putri@gmail.com',       '086789012345', 'Lapangan B',  'Futsal',     '2026-05-14', '16:00', '18:00', 200000,  'pending',    ''),
('Hendra Wijaya',   'hendra.wijaya@email.com',   '087890123456', 'Lapangan E',  'Voli',       '2026-05-15', '09:00', '11:00', 120000,  'confirmed',  'Tim kantor'),
('Dewi Anggraini',  'dewi.anggraini@gmail.com',  '088901234567', 'Lapangan A',  'Badminton',  '2026-05-15', '14:00', '16:00', 100000,  'cancelled',  'Dibatalkan karena hujan'),
('Rizky Pratama',   'rizky.pratama@gmail.com',   '089012345678', 'Lapangan C',  'Basket',     '2026-05-16', '10:00', '12:00', 150000,  'pending',    'Turnamen internal'),
('Lestari Dewi',    'lestari.dewi@yahoo.com',    '081122334455', 'Lapangan D',  'Tenis',      '2026-05-16', '15:00', '17:00', 180000,  'confirmed',  'Latihan rutin');
