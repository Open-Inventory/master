BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS `SubCategory` (
	`SubCategoryID`	INTEGER,
	`SubCategoryType`	TEXT NOT NULL,
	PRIMARY KEY(`SubCategoryID`)
);
CREATE TABLE IF NOT EXISTS `Product` (
	`ProductKey`	INTEGER,
	`ProductName`	TEXT NOT NULL,
	`CategoryID`	INTEGER NOT NULL,
	`SubCategoryID`	INTEGER NOT NULL,
	FOREIGN KEY(`SubCategoryID`) REFERENCES `SubCategory`(`SubCategoryID`),
	FOREIGN KEY(`CategoryID`) REFERENCES `Category`(`CategoryID`),
	PRIMARY KEY(`ProductKey`)
);
CREATE TABLE IF NOT EXISTS `Item` (
	`ItemID`	INTEGER,
	`ProductKey`	INTEGER NOT NULL,
	`ItemStatus`	INTEGER NOT NULL DEFAULT 0,
	PRIMARY KEY(`ItemID`),
	FOREIGN KEY(`ProductKey`) REFERENCES `Product`(`ProductKey`)
);
CREATE TABLE IF NOT EXISTS `Category` (
	`CategoryID`	INTEGER,
	`CategoryType`	TEXT NOT NULL,
	PRIMARY KEY(`CategoryID`)
);
COMMIT;
