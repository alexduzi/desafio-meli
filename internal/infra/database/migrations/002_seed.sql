-- Product 1: iPhone 15 Pro Max
INSERT INTO products (id, title, description, price, currency, condition, stock, seller_id, seller_name, category)
VALUES (
    'MLB001',
    'iPhone 15 Pro Max 256GB - Titanium Blue',
    'Latest Apple flagship smartphone with A17 Pro chip, titanium design, and advanced camera system. Includes original box, charger cable, and documentation. Factory unlocked, works with all carriers.',
    1299.99,
    'USD',
    'new',
    45,
    'SELLER001',
    'TechWorld Store',
    'Electronics > Smartphones'
);

INSERT INTO product_images (product_id, image_url, display_order) VALUES
('MLB001', 'https://images.unsplash.com/photo-1696446702230-a8ff49103cd1?w=800', 0),
('MLB001', 'https://images.unsplash.com/photo-1695048133142-1a20484d2569?w=800', 1),
('MLB001', 'https://images.unsplash.com/photo-1695048133082-1a02e074f08f?w=800', 2);

-- Product 2: Gaming Laptop
INSERT INTO products (id, title, description, price, currency, condition, stock, seller_id, seller_name, category)
VALUES (
    'MLB002',
    'ASUS ROG Strix G16 Gaming Laptop - RTX 4070',
    'High-performance gaming laptop with Intel i9-13980HX, NVIDIA RTX 4070 8GB, 16GB DDR5 RAM, 1TB NVMe SSD, 16" QHD 240Hz display. Perfect for gaming and content creation. Barely used, like new condition.',
    1899.50,
    'USD',
    'used',
    8,
    'SELLER002',
    'Gaming Paradise',
    'Electronics > Computers'
);

INSERT INTO product_images (product_id, image_url, display_order) VALUES
('MLB002', 'https://images.unsplash.com/photo-1603302576837-37561b2e2302?w=800', 0),
('MLB002', 'https://images.unsplash.com/photo-1593642632823-8f785ba67e45?w=800', 1),
('MLB002', 'https://images.unsplash.com/photo-1588872657578-7efd1f1555ed?w=800', 2),
('MLB002', 'https://images.unsplash.com/photo-1625019030820-e4ed970a6c95?w=800', 3);

-- Product 3: Mechanical Keyboard
INSERT INTO products (id, title, description, price, currency, condition, stock, seller_id, seller_name, category)
VALUES (
    'MLB003',
    'Keychron Q1 Pro Mechanical Keyboard - Wireless',
    'Premium 75% layout mechanical keyboard with hot-swappable switches, RGB backlighting, aluminum frame. Includes Gateron Pro Red switches, USB-C cable, and keycap puller. Wireless connectivity via Bluetooth 5.1.',
    189.99,
    'USD',
    'new',
    120,
    'SELLER003',
    'Keyboard Kingdom',
    'Electronics > Computer Accessories'
);

INSERT INTO product_images (product_id, image_url, display_order) VALUES
('MLB003', 'https://images.unsplash.com/photo-1595225476474-87563907a212?w=800', 0),
('MLB003', 'https://images.unsplash.com/photo-1587829741301-dc798b83add3?w=800', 1);

-- Product 4: Nike Air Jordan 1 Retro
INSERT INTO products (id, title, description, price, currency, condition, stock, seller_id, seller_name, category)
VALUES (
    'MLB004',
    'Nike Air Jordan 1 Retro High OG "Chicago" - Size 10',
    'Iconic Air Jordan 1 in the classic Chicago colorway (White/Black/Red). Authentic, never worn, includes original box with all accessories. Premium leather construction. A must-have for sneaker collectors.',
    450.00,
    'USD',
    'new',
    3,
    'SELLER004',
    'Sneaker Vault',
    'Fashion > Shoes > Sneakers'
);

INSERT INTO product_images (product_id, image_url, display_order) VALUES
('MLB004', 'https://images.unsplash.com/photo-1556906781-9a412961c28c?w=800', 0),
('MLB004', 'https://images.unsplash.com/photo-1600269452121-4f2416e55c28?w=800', 1),
('MLB004', 'https://images.unsplash.com/photo-1552346154-21d32810aba3?w=800', 2);

-- Product 5: Sony WH-1000XM5 Headphones
INSERT INTO products (id, title, description, price, currency, condition, stock, seller_id, seller_name, category)
VALUES (
    'MLB005',
    'Sony WH-1000XM5 Wireless Noise Cancelling Headphones - Black',
    'Industry-leading noise cancellation with Sony''s HD Noise Canceling Processor QN1. 30-hour battery life, multipoint connection, premium sound quality. Used for 2 months, excellent condition. Includes carrying case and cables.',
    329.99,
    'USD',
    'used',
    15,
    'SELLER001',
    'TechWorld Store',
    'Electronics > Audio > Headphones'
);

INSERT INTO product_images (product_id, image_url, display_order) VALUES
('MLB005', 'https://images.unsplash.com/photo-1546435770-a3e426bf472b?w=800', 0),
('MLB005', 'https://images.unsplash.com/photo-1484704849700-f032a568e944?w=800', 1),
('MLB005', 'https://images.unsplash.com/photo-1505740420928-5e560c06d30e?w=800', 2),
('MLB005', 'https://images.unsplash.com/photo-1524678606370-a47ad25cb82a?w=800', 3);