CREATE TABLE banners(
    id serial primary key,
    banner_name varchar(100),
    banner_image varchar(255),
    description varchar(100)
);

INSERT INTO banners (banner_name, banner_image, description) VALUES
('Promo Tahun Baru', 'https://example.com/images/banner1.jpg', 'Diskon hingga 50% untuk produk tertentu'),
('Flash Sale', 'https://example.com/images/banner2.jpg', 'Penawaran terbatas selama 24 jam'),
('Gratis Ongkir', 'https://example.com/images/banner3.jpg', 'Nikmati gratis ongkir untuk pembelian di atas Rp 100.000'),
('Diskon Spesial', 'https://example.com/images/banner4.jpg', 'Hanya hari ini, potongan harga spesial'),
('Produk Baru', 'https://example.com/images/banner5.jpg', 'Produk terbaru telah tersedia di toko kami');
