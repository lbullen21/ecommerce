USE fizzyFactory;

CREATE TABLE products (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    flavor TEXT NOT NULL,
    photo TEXT NOT NULL,
    price FLOAT NOT NULL,
    detail TEXT NOT NULL);


INSERT INTO products (flavor, photo, price, detail)
	 VALUES
        ('Limoncello', './images/limoncello.png', '5.99', 'pack of 12'),
	    ('Grapefruit', './images/sdGrapefruit.png', '5.99', 'pack of 12'),
        ('Key Lime', './images/keyLime.png', '5.99', 'pack of 12'),
        ('Lemon', './images/sdLemon.png', '5.99', 'pack of 12'),
        ('Lime', './images/sdLime.png', '5.99', 'pack of 12'),
        ('Tangerine', './images/tangerine.png', '5.99', 'pack of 12'),
        ('Hibiscus', './images/hibiscus.png', '5.99', 'pack of 12'),
        ('Variety Pack', './images/assorted.png', '12.99', '3 packs of 12');

