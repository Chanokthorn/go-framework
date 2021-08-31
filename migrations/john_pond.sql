create table pond
(
    PondID      int auto_increment
        primary key,
    PondUUID    text                 not null,
    Location    text                 null,
    CreatedBy   text                 null,
    CreatedDate datetime             null,
    UpdatedBy   text                 null,
    UpdatedDate datetime             null,
    IsDeleted   tinyint(1) default 0 null
);

INSERT INTO john.pond (PondID, PondUUID, Location, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (1, '5371d34e-3af6-34ad-8b4f-fbfe66efa2e4', 'Vitaport', 'system', '2021-08-30 23:33:17', null, null, 0);
INSERT INTO john.pond (PondID, PondUUID, Location, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (2, 'f60b9700-664c-3d93-b785-38659a546cde', 'Kesslerfort', 'system', '2021-08-31 14:27:47', null, null, 0);
INSERT INTO john.pond (PondID, PondUUID, Location, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (3, '9a5ac925-05c7-3c65-a1ba-8a6057d10a82', 'Bettyport', 'system', '2021-08-31 14:27:48', null, null, 0);
INSERT INTO john.pond (PondID, PondUUID, Location, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (4, '85e8064c-371e-317e-9782-54fbbfae5e90', 'Geraldland', 'system', '2021-08-31 14:27:49', null, null, 0);
INSERT INTO john.pond (PondID, PondUUID, Location, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (5, 'd1676435-728f-3553-a732-e42b0ec4aeb0', 'New Becker', 'system', '2021-08-31 14:27:50', null, null, 0);
INSERT INTO john.pond (PondID, PondUUID, Location, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (6, '41a8a16e-75a0-3f8c-90b0-054ecf67e869', 'Leannonshire', 'system', '2021-08-31 14:27:51', null, null, 0);
INSERT INTO john.pond (PondID, PondUUID, Location, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (7, 'cfd0dd02-d529-3039-8b37-0751ed443ca6', 'West Stanton', 'system', '2021-08-31 14:27:52', null, null, 0);
INSERT INTO john.pond (PondID, PondUUID, Location, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (8, '71244784-45a5-3bbb-8656-89ec66102bb4', 'Michaleborough', 'system', '2021-08-31 14:27:53', null, null, 0);
INSERT INTO john.pond (PondID, PondUUID, Location, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (9, 'dda84a69-7be3-3614-8db8-07550e95f368', 'Krisside', 'system', '2021-08-31 14:27:54', null, null, 0);
INSERT INTO john.pond (PondID, PondUUID, Location, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (10, '86d2f703-8d17-380c-a31b-cd586261aa9f', 'Olsontown', 'system', '2021-08-31 14:27:55', null, null, 0);