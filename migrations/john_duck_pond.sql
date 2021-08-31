create table duck_pond
(
    DuckPondID  int auto_increment
        primary key,
    DuckID      int                  null,
    PondID      int                  null,
    CreatedBy   text                 null,
    CreatedDate datetime             null,
    UpdatedBy   text                 null,
    UpdatedDate datetime             null,
    IsDeleted   tinyint(1) default 0 null
);

INSERT INTO john.duck_pond (DuckPondID, DuckID, PondID, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (1, 30, 6, 'system', '2021-08-31 15:12:56', 'system', '2021-08-31 15:22:36', 1);
INSERT INTO john.duck_pond (DuckPondID, DuckID, PondID, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (2, 30, 7, 'system', '2021-08-31 15:12:56', 'system', '2021-08-31 15:22:36', 1);
INSERT INTO john.duck_pond (DuckPondID, DuckID, PondID, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (3, 30, 8, 'system', '2021-08-31 15:12:56', 'system', '2021-08-31 15:22:36', 1);
INSERT INTO john.duck_pond (DuckPondID, DuckID, PondID, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (4, 32, 6, 'system', '2021-08-31 15:18:33', null, null, 0);
INSERT INTO john.duck_pond (DuckPondID, DuckID, PondID, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (5, 32, 7, 'system', '2021-08-31 15:18:33', null, null, 0);
INSERT INTO john.duck_pond (DuckPondID, DuckID, PondID, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (6, 32, 8, 'system', '2021-08-31 15:18:33', null, null, 0);
INSERT INTO john.duck_pond (DuckPondID, DuckID, PondID, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (7, 30, 1, 'system', '2021-08-31 15:20:50', 'system', '2021-08-31 15:22:36', 1);
INSERT INTO john.duck_pond (DuckPondID, DuckID, PondID, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (8, 30, 2, 'system', '2021-08-31 15:20:50', 'system', '2021-08-31 15:22:36', 1);
INSERT INTO john.duck_pond (DuckPondID, DuckID, PondID, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (9, 30, 3, 'system', '2021-08-31 15:20:50', 'system', '2021-08-31 15:22:36', 1);
INSERT INTO john.duck_pond (DuckPondID, DuckID, PondID, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (10, 30, 4, 'system', '2021-08-31 15:20:50', 'system', '2021-08-31 15:22:36', 1);
INSERT INTO john.duck_pond (DuckPondID, DuckID, PondID, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (11, 33, 1, 'system', '2021-08-31 15:21:50', 'system', '2021-08-31 15:22:12', 1);
INSERT INTO john.duck_pond (DuckPondID, DuckID, PondID, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (12, 33, 2, 'system', '2021-08-31 15:21:50', 'system', '2021-08-31 15:22:12', 1);
INSERT INTO john.duck_pond (DuckPondID, DuckID, PondID, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (13, 33, 3, 'system', '2021-08-31 15:21:50', 'system', '2021-08-31 15:22:12', 1);
INSERT INTO john.duck_pond (DuckPondID, DuckID, PondID, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (14, 33, 4, 'system', '2021-08-31 15:21:50', 'system', '2021-08-31 15:22:12', 1);
INSERT INTO john.duck_pond (DuckPondID, DuckID, PondID, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (15, 30, 1, 'system', '2021-08-31 15:22:36', null, null, 0);
INSERT INTO john.duck_pond (DuckPondID, DuckID, PondID, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (16, 30, 2, 'system', '2021-08-31 15:22:36', null, null, 0);
INSERT INTO john.duck_pond (DuckPondID, DuckID, PondID, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (17, 30, 3, 'system', '2021-08-31 15:22:36', null, null, 0);
INSERT INTO john.duck_pond (DuckPondID, DuckID, PondID, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (18, 30, 4, 'system', '2021-08-31 15:22:36', null, null, 0);