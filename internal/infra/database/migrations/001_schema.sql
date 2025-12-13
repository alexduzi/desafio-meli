CREATE TABLE products (
        id TEXT PRIMARY KEY,
        title TEXT NOT NULL,
        description TEXT,
        price REAL NOT NULL,
        currency TEXT NOT NULL,
        condition TEXT CHECK(condition IN ('new', 'used', 'refurbished')),
        stock INTEGER DEFAULT 0,
        seller_id TEXT NOT NULL,
        seller_name TEXT NOT NULL,
        category TEXT,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );
    
    CREATE TABLE product_images (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        product_id TEXT NOT NULL,
        image_url TEXT NOT NULL,
        display_order INTEGER DEFAULT 0,
        FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
    );