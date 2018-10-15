BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS `Users` (
	`UserID`	INTEGER NOT NULL UNIQUE,
	`UserName`	TEXT NOT NULL UNIQUE,
	`EmailAddress`	TEXT NOT NULL UNIQUE,
	`UserPassword`	TEXT NOT NULL,
	`DateCreated`	INTEGER NOT NULL,
	PRIMARY KEY(`UserID`)
);
CREATE TABLE IF NOT EXISTS `ProductCategory` (
	`ProductID`	INTEGER NOT NULL,
	`CategoryID`	INTEGER NOT NULL,
	PRIMARY KEY(`ProductID`,`CategoryID`),
	FOREIGN KEY(`CategoryID`) REFERENCES `Category`(`CategoryID`),
	FOREIGN KEY(`ProductID`) REFERENCES `Product`(`ProductID`)
);
CREATE TABLE IF NOT EXISTS `Product` (
	`ProductID`	INTEGER NOT NULL UNIQUE,
	`ProductName`	TEXT NOT NULL,
	`Price`	REAL NOT NULL,
	`UnitsInStock`	INTEGER NOT NULL DEFAULT 1,
	`DateAdded`	INTEGER NOT NULL,
	`ProductWeightPerUnit`	REAL NOT NULL,
	`StorageLocation`	TEXT,
	`Brand`	TEXT,
	`Manufacturer`	TEXT,
	`ItemStatus`	NUMERIC NOT NULL DEFAULT 1,
	PRIMARY KEY(`ProductID`)
);
CREATE TABLE IF NOT EXISTS `Orders` (
	`OrderID`	INTEGER NOT NULL UNIQUE,
	`ProductID`	INTEGER NOT NULL UNIQUE,
	`TotalOrdered`	INTEGER NOT NULL,
	`OrderStatus`	NUMERIC NOT NULL DEFAULT 1,
	`OrderDate`	NUMERIC NOT NULL,
	`ReceivingAddress`	TEXT NOT NULL,
	`ShippedDate`	NUMERIC,
	`TrackingNumber`	INTEGER,
	PRIMARY KEY(`OrderID`),
	FOREIGN KEY(`ProductID`) REFERENCES `Product`(`ProductID`)
);
CREATE TABLE IF NOT EXISTS `Category` (
	`CategoryID`	INTEGER NOT NULL,
	`CategoryName`	TEXT NOT NULL,
	PRIMARY KEY(`CategoryID`)
);
COMMIT;
