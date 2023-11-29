CREATE TABLE Appointments (
    AppointmentID INT PRIMARY KEY,
    StartTime DATETIME NOT NULL,
    EndTime DATETIME NOT NULL,
    UserID INT NOT NULL,
    ResourceID INT NOT NULL,
    FOREIGN KEY (UserID) REFERENCES Users(UserID),
    FOREIGN KEY (ResourceID) REFERENCES Resources(ResourceID),
    CHECK (EndTime > StartTime)
);

-- Insert a record
INSERT INTO Appointments (AppointmentID, StartTime, EndTime, UserID, ResourceID)
VALUES (1, '2023-12-01 09:00:00', '2023-12-01 10:00:00', 1, 1);
