USE fizzyFactory;

CREATE TABLE products (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    flavor TEXT NOT NULL,
    photo TEXT NOT NULL,
    price FLOAT NOT NULL,
    detail TEXT NOT NULL, 
    category TEXT NOT NULL
    );


INSERT INTO products (flavor, photo, price, detail, category)
	 VALUES
        ('Limoncello', './images/limoncello.png', '5.99', 'pack of 12', 'singlePack'),
	    ('Grapefruit', './images/sdGrapefruit.png', '5.99', 'pack of 12', 'singlePack'),
        ('Key Lime', './images/keyLime.png', '5.99', 'pack of 12', 'singlePack'),
        ('Lemon', './images/sdLemon.png', '5.99', 'pack of 12', 'singlePack'),
        ('Lime', './images/sdLime.png', '5.99', 'pack of 12', 'singlePack'),
        ('Tangerine', './images/tangerine.png', '5.99', 'pack of 12', 'singlePack'),
        ('Hibiscus', './images/hibiscus.png', '5.99', 'pack of 12', 'singlePack'),
        ('Variety Pack', './images/assorted.png', '12.99', '3 packs of 12', 'featured'),
        ('FEATURED Variety Pack', './images/homepage2.png', '12.99', '3 packs of 12', 'featured');

