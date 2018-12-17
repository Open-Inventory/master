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
INSERT INTO `SubCategory` VALUES (10,'Memory Card');
INSERT INTO `SubCategory` VALUES (13,'Computer Parts');
INSERT INTO `SubCategory` VALUES (14,'Memory Card');
CREATE TABLE IF NOT EXISTS `Product` (
	`ProductKey`	INTEGER,
	`ProductName`	TEXT NOT NULL,
	`CategoryID`	INTEGER NOT NULL,
	`SubCategoryID`	INTEGER NOT NULL,
	FOREIGN KEY(`SubCategoryID`) REFERENCES `SubCategory`(`SubCategoryID`),
	PRIMARY KEY(`ProductKey`),
	FOREIGN KEY(`CategoryID`) REFERENCES `Category`(`CategoryID`)
);
INSERT INTO `Product` VALUES (1,'ASRock Fatal1ty B250M Performance',1,1);
INSERT INTO `Product` VALUES (2,'SanDisk 128GB Ultra microSDXC UHS-I Memory Card with Adapter',2,3);
INSERT INTO `Product` VALUES (3,'Dell UltraSharp U2715H 27-Inch Screen LED-Lit Monitor',3,2);
INSERT INTO `Product` VALUES (6,'SanDisk 128GB Ultra microSDXC UHS-I Memory Card with Adapter',14,14);
CREATE TABLE IF NOT EXISTS `Orders` (
	`OrderID`	INTEGER,
	`ProductName`	TEXT,
	`Order_Date`	INTEGER,
	`Order_Status`	INTEGER,
	`Order_Quantity`	INTEGER,
	`ItemID`	INTEGER,
	PRIMARY KEY(`OrderID`)
);
INSERT INTO `Orders` VALUES (1,'ASRock Fatal1ty B250M Performance',1,0,2,1);
INSERT INTO `Orders` VALUES (2,'SanDisk 128GB Ultra microSDXC UHS-I Memory Card with Adapter',1,0,1,3);
INSERT INTO `Orders` VALUES (3,'Dell UltraSharp U2715H 27-Inch Screen LED-Lit Monitor',1,0,1,4);
INSERT INTO `Orders` VALUES (4,'Dell UltraSharp U2715H 27-Inch Screen LED-Lit Monitor',1,0,1,5);
INSERT INTO `Orders` VALUES (5,'ASRock Fatal1ty B250M Performance',1,0,1,2);
INSERT INTO `Orders` VALUES (6,'SanDisk 128GB Ultra microSDXC UHS-I Memory Card with Adapter',1,0,1,8);
CREATE TABLE IF NOT EXISTS `Item` (
	`ItemID`	INTEGER,
	`ProductKey`	INTEGER NOT NULL,
	`ItemStatus`	INTEGER NOT NULL DEFAULT 0,
	`Created_On`	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	`Updated_On`	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	`OrderID`	INTEGER DEFAULT 0,
	FOREIGN KEY(`ProductKey`) REFERENCES `Product`(`ProductKey`),
	PRIMARY KEY(`ItemID`)
);
INSERT INTO `Item` VALUES (1,1,0,'2018-12-05 00:45:39','2018-12-05 00:45:39',1);
INSERT INTO `Item` VALUES (2,1,0,'2018-12-05 00:58:07','2018-12-05 00:58:07',5);
INSERT INTO `Item` VALUES (3,2,0,'2018-12-05 00:58:11','2018-12-05 00:58:11',2);
INSERT INTO `Item` VALUES (4,3,0,'2018-12-05 00:58:13','2018-12-05 00:58:13',3);
INSERT INTO `Item` VALUES (5,3,0,'2018-12-05 00:58:15','2018-12-05 00:58:15',4);
INSERT INTO `Item` VALUES (6,3,0,'2018-12-05 00:58:18','2018-12-05 00:58:18',1);
INSERT INTO `Item` VALUES (7,2,0,'2018-12-05 19:05:40','2018-12-05 19:05:40',0);
INSERT INTO `Item` VALUES (8,6,0,'2018-12-07 20:12:47','2018-12-07 20:12:47',6);
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
INSERT INTO `Category` VALUES (10,'Computer Accesories');
INSERT INTO `Category` VALUES (13,'Computers');
INSERT INTO `Category` VALUES (14,'Computer Accesories');
COMMIT;
