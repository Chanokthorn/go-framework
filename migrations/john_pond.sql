create table pond
(
    PondID      int auto_increment
        primary key,
    PondUUID    text             not null,
    Location    text             null,
    IsActive    bit default b'1' null,
    CreatedBy   text             null,
    CreatedDate datetime         null,
    UpdatedBy   text             null,
    UpdatedDate datetime         null,
    IsDeleted   bit default b'0' null
);

INSERT INTO john.pond (PondID, PondUUID, Location, IsActive, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (1, '4cf4e1ba-d16d-3009-b1c1-2a67d0085a58', 'Cambodia', true, '1212312121', '2021-09-03 09:58:48', null, null, false);