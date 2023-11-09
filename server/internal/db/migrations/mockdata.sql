SET FOREIGN_KEY_CHECKS = 0; 
TRUNCATE table products; 
SET FOREIGN_KEY_CHECKS = 1;

DROP PROCEDURE IF EXISTS insertProducts;

CREATE PROCEDURE insertProducts()
BEGIN 
	DECLARE i INT DEFAULT 1;
	WHILE i < 100 DO
		IF MOD(i,2) = 0 THEN
			SET @brand = IF(MOD(i,4)=0, "Brand Z", "Brand X");
			SET @size = IF(MOD(i,6)=0, "L", "M");
			SET @color = IF(MOD(i,8)=0, "Red", "Black");
			SET @status = IF(MOD(i,10)=0,FALSE, TRUE);
		ELSE
			SET @brand = IF(MOD(i,3)=0, "Brand T", "Brand Y");
			SET @size = IF(MOD(i,9)=0, "XL", "S");
			SET @color = IF(MOD(i,7)=0, "Blue", "White");
			SET @status = IF(MOD(i,5)=0,TRUE, FALSE);
		END IF;
		SET @SKU = CONCAT("PRO0",i,"-",@brand, "-", @color, "-", @size);
		INSERT INTO products (SKU, name, brand, size, color, status) VALUES (@SKU,CONCAT("Product ", i), @brand, @size, @color, @status);
		SET i = i + 1;
	END WHILE;
END;

CALL insertProducts();

SET FOREIGN_KEY_CHECKS = 0; 
TRUNCATE table suppliers;
SET FOREIGN_KEY_CHECKS = 1;

DROP PROCEDURE IF EXISTS insertSuppliers;

CREATE PROCEDURE insertSuppliers()
BEGIN 
	DECLARE i INT DEFAULT 1;
	WHILE i < 100 DO
		SET @email = CONCAT("supplier",i,"@mail.com");
		SET @contact_number = CONCAT("+012345678",IF(i<10, CONCAT("0",i), i));
		SET @status = IF(MOD(i,2)=0,FALSE, TRUE);
		INSERT INTO suppliers (name, email, contact_number, status) VALUES (CONCAT("Supplier ", i), @email, @contact_number, @status);
		SET i = i + 1;
	END WHILE;
END;

CALL insertSuppliers();