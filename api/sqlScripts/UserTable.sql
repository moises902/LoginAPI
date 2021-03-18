-- Query to create Users table in the Medicare DB
-- Must already be in Medicare DB

DROP TABLE IF EXISTS Users;

CREATE TABLE Users (
    autoID INT NOT NULL AUTO_INCREMENT,
    email VARCHAR(60) NOT NULL UNIQUE,
    pass VARCHAR(60) NOT NULL,
    PRIMARY KEY (autoID)
);