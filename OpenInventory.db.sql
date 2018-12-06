BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS `SubCategory` (
	`SubCategoryID`	INTEGER,
	`SubCategoryType`	TEXT NOT NULL,
	PRIMARY KEY(`SubCategoryID`)
);
INSERT INTO `SubCategory` VALUES (1,'Motherboards');
INSERT INTO `SubCategory` VALUES (2,'Monitors');
INSERT INTO `SubCategory` VALUES (3,'Memory Card');
INSERT INTO `SubCategory` VALUES (4,'');
INSERT INTO `SubCategory` VALUES (5,'Memory Card');
INSERT INTO `SubCategory` VALUES (6,'N/A');
INSERT INTO `SubCategory` VALUES (7,'N/A');
INSERT INTO `SubCategory` VALUES (8,'N/A');
INSERT INTO `SubCategory` VALUES (9,'N/A');
INSERT INTO `SubCategory` VALUES (10,'Memory Card');
INSERT INTO `SubCategory` VALUES (11,'N/A');
INSERT INTO `SubCategory` VALUES (12,'N/A');
INSERT INTO `SubCategory` VALUES (13,'Computer Parts');
CREATE TABLE IF NOT EXISTS `Product` (
	`ProductKey`	INTEGER,
	`ProductName`	TEXT NOT NULL,
	`CategoryID`	INTEGER NOT NULL,
	`SubCategoryID`	INTEGER NOT NULL,
	PRIMARY KEY(`ProductKey`),
	FOREIGN KEY(`CategoryID`) REFERENCES `Category`(`CategoryID`),
	FOREIGN KEY(`SubCategoryID`) REFERENCES `SubCategory`(`SubCategoryID`)
);
INSERT INTO `Product` VALUES (1,'ASRock Fatal1ty B250M Performance',1,1);
INSERT INTO `Product` VALUES (2,'SanDisk 128GB Ultra microSDXC UHS-I Memory Card with Adapter',2,3);
INSERT INTO `Product` VALUES (3,'Dell UltraSharp U2715H 27-Inch Screen LED-Lit Monitor',3,2);
INSERT INTO `Product` VALUES (4,'',4,4);
INSERT INTO `Product` VALUES (5,'SanDisk 128GB Ultra microSDXC UHS-I Memory Card with Adapter',5,5);
INSERT INTO `Product` VALUES (6,'N/A',6,6);
INSERT INTO `Product` VALUES (7,'N/A',7,7);
INSERT INTO `Product` VALUES (9,'N/A',9,9);
INSERT INTO `Product` VALUES (10,'SanDisk 128GB Ultra microSDXC UHS-I Memory Card with Adapter',10,10);
INSERT INTO `Product` VALUES (11,'word',11,11);
INSERT INTO `Product` VALUES (12,'N/A',12,12);
INSERT INTO `Product` VALUES (13,'Word',13,13);
CREATE TABLE IF NOT EXISTS `Orders` (
	`OrderID`	INTEGER,
	`ItemID`	INTEGER,
	`Order_Date`	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,,
	`Order_Status`	INTEGER,
	PRIMARY KEY(`OrderID`)
);
CREATE TABLE IF NOT EXISTS `Item` (
	`ItemID`	INTEGER,
	`ProductKey`	INTEGER NOT NULL,
	`ItemStatus`	INTEGER NOT NULL DEFAULT 0,
	`Created_On`	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	`Updated_On`	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY(`ProductKey`) REFERENCES `Product`(`ProductKey`),
	PRIMARY KEY(`ItemID`)
);
INSERT INTO `Item` VALUES (1,1,0,'2018-12-05 00:45:39','2018-12-05 00:45:39');
INSERT INTO `Item` VALUES (2,1,0,'2018-12-05 00:58:07','2018-12-05 00:58:07');
INSERT INTO `Item` VALUES (3,2,0,'2018-12-05 00:58:11','2018-12-05 00:58:11');
INSERT INTO `Item` VALUES (4,3,0,'2018-12-05 00:58:13','2018-12-05 00:58:13');
INSERT INTO `Item` VALUES (5,3,0,'2018-12-05 00:58:15','2018-12-05 00:58:15');
INSERT INTO `Item` VALUES (6,3,0,'2018-12-05 00:58:18','2018-12-05 00:58:18');
INSERT INTO `Item` VALUES (7,2,0,'2018-12-05 19:05:40','2018-12-05 19:05:40');
INSERT INTO `Item` VALUES (8,4,0,'2018-12-05 19:32:36','2018-12-05 19:32:36');
INSERT INTO `Item` VALUES (9,5,0,'2018-12-05 20:01:55','2018-12-05 20:01:55');
INSERT INTO `Item` VALUES (10,10,0,'2018-12-05 22:19:45','2018-12-05 22:19:45');
INSERT INTO `Item` VALUES (11,12,0,'2018-12-05 22:45:22','2018-12-05 22:45:22');
INSERT INTO `Item` VALUES (12,13,0,'2018-12-05 22:45:52','2018-12-05 22:45:52');
CREATE TABLE IF NOT EXISTS `Category` (
	`CategoryID`	INTEGER,
	`CategoryType`	TEXT NOT NULL,
	PRIMARY KEY(`CategoryID`)
);
INSERT INTO `Category` VALUES (1,'Computer Parts');
INSERT INTO `Category` VALUES (2,'Computer Accesories');
INSERT INTO `Category` VALUES (3,'Computer Peripherals');
INSERT INTO `Category` VALUES (4,'');
INSERT INTO `Category` VALUES (5,'Computer Accesories');
INSERT INTO `Category` VALUES (6,'N/A');
INSERT INTO `Category` VALUES (7,'N/A');
INSERT INTO `Category` VALUES (8,'N/A');
INSERT INTO `Category` VALUES (9,'N/A');
INSERT INTO `Category` VALUES (10,'Computer Accesories');
INSERT INTO `Category` VALUES (11,'N/A');
INSERT INTO `Category` VALUES (12,'N/A');
INSERT INTO `Category` VALUES (13,'Computers');
COMMIT;
