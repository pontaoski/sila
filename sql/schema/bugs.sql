CREATE TABLE Products (
    ID BIGSERIAL NOT NULL PRIMARY KEY,
    Name TEXT,
    Product_Description TEXT,
    Active boolean
);

CREATE TABLE Components (
    ID BIGSERIAL NOT NULL PRIMARY KEY,
    Name TEXT,
    Component_Description TEXT,
    Product_ID BIGSERIAL NOT NULL,
    FOREIGN KEY (Product_ID) REFERENCES Products(ID)
);

CREATE TABLE Users (
    ID BIGSERIAL NOT NULL PRIMARY KEY,
    Email TEXT,
    FullName TEXT
);

CREATE TABLE Bugs (
    ID BIGSERIAL NOT NULL PRIMARY KEY,
    Title TEXT,
    Created_At TIMESTAMP NOT NULL,
    Edited_At TIMESTAMP,
    Bug_Description TEXT,
    Bug_Severity Severity,
    Bug_Status Status,
    Component_ID BIGSERIAL NOT NULL,
    FOREIGN KEY (Component_ID) REFERENCES Components(ID),
    Author_ID BIGSERIAL NOT NULL,
    FOREIGN KEY (Author_ID) REFERENCES Users(ID)
);


CREATE TABLE Comments (
    ID BIGSERIAL NOT NULL PRIMARY KEY,
    Comment_Text TEXT,
    Created_At TIMESTAMP NOT NULL,
    Edited_At TIMESTAMP,
    Redacted boolean DEFAULT FALSE,
    Bug_ID BIGSERIAL NOT NULL,
    FOREIGN KEY (Bug_ID) REFERENCES Bugs(ID),
    Author_ID BIGSERIAL NOT NULL,
    FOREIGN KEY (Author_ID) REFERENCES Users(ID)
);

CREATE TABLE Bug_Dependencies (
    Required_By BIGSERIAL NOT NULL,
    Depends_On  BIGSERIAL NOT NULL,
    FOREIGN KEY (Required_By) REFERENCES Bugs(ID),
    FOREIGN KEY (Depends_On) REFERENCES Bugs(ID)
);

CREATE TYPE severity AS ENUM (
    'Wishlist', 'Minor', 'Normal', 'Grave', 'Major', 'Critical', 'Crash',
    'Task',
    'MinorSecurityIssue', 'MediumSecurityIssue', 'MajorSecurityIssue', 'CriticalSecurityIssue'
);

CREATE TYPE status AS ENUM (
    'Reported', 'Confirmed', 'Assigned', 'Resolved', 'NeedInfo'
);
