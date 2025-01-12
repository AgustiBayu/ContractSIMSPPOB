CREATE TABLE layanans(
    id serial primary key,
    service_code varchar(50) not null,
    service_name varchar(50) not null,
    service_icon varchar(255) not null,
    service_tarif int not null
);

INSERT INTO layanans (service_code, service_name, service_icon, service_tarif) VALUES
('SVC001', 'Cuci Mobil', 'https://example.com/icons/car-wash.png', 50000),
('SVC002', 'Service Motor', 'https://example.com/icons/motor-service.png', 75000),
('SVC003', 'Laundry Kiloan', 'https://example.com/icons/laundry.png', 30000),
('SVC004', 'Jasa Cleaning', 'https://example.com/icons/cleaning.png', 100000),
('SVC005', 'Potong Rambut', 'https://example.com/icons/haircut.png', 40000),
('SVC006', 'Pijat Tradisional', 'https://example.com/icons/massage.png', 120000),
('SVC007', 'Catering Harian', 'https://example.com/icons/catering.png', 50000),
('SVC008', 'Jasa Kurir', 'https://example.com/icons/courier.png', 25000),
('SVC009', 'Rental Mobil', 'https://example.com/icons/car-rental.png', 300000),
('SVC010', 'Jasa Fotografi', 'https://example.com/icons/photography.png', 200000);