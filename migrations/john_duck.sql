create table duck
(
    DuckID      int auto_increment
        primary key,
    DuckUUID    text             not null,
    Name        text             null,
    Color       text             null,
    IsActive    bit default b'1' null,
    CreatedBy   text             null,
    CreatedDate datetime         null,
    UpdatedBy   text             null,
    UpdatedDate datetime         null,
    IsDeleted   bit default b'0' null
);

INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (1, '123123123', 'First', 'white', 'bone', '2021-08-26 23:45:55', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (2, '523523432', 'Second', 'red', 'bone', '2021-08-27 00:08:50', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (3, '323232323', 'Third', 'white', 'bone', '2021-08-27 00:30:30', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (4, 'c3f4bf31-356b-3d16-9de8-801fdfc28928', 'Darrin Hermiston', 'DarkOrchid', 'system', '2021-08-27 00:42:24', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (5, 'f4739e82-913f-38fc-9764-46a2f66e73e3', 'Michaela West', 'Plum', 'system', '2021-08-27 00:45:33', 'system', '2021-08-27 11:52:00', true);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (6, '8387496c-2c44-3d5b-8701-4ea487acab3f', 'Willard Hansen', 'Azure', 'system', '2021-08-27 11:07:30', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (7, '30076911-25ec-3b66-b424-5cf8fdc726fa', 'Sharon Hahn', 'OldLace', 'system', '2021-08-27 11:14:43', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (8, '0817e1b9-4f83-3549-8645-2b3126d843f0', 'Wallace Kirlin', 'IndianRed ', 'system', '2021-08-27 11:15:17', 'system', '2021-08-30 10:56:19', false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (9, '2adeb1ca-562b-3975-ae6a-1ff09dd90e46', 'Forrest Ratke', 'CornflowerBlue', 'system', '2021-08-27 11:15:54', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (10, 'c60faf70-e9fa-3654-8f15-dc68f50e24aa', 'Demetrius Pfannerstill', 'SeaShell', 'system', '2021-08-27 11:17:49', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (11, 'a9489690-352e-3694-a4ba-97ba6d5e41af', 'Randi Prohaska', 'HotPink', 'system', '2021-08-27 11:18:50', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (12, '09370ecf-7095-3391-aacd-d13d4fa676d0', 'Guy Kessler', 'GhostWhite', 'system', '2021-08-27 11:21:54', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (13, '4c382d53-3edb-34a1-8289-c6155ddaf890', 'Lonnie Brakus', 'Bisque', 'system', '2021-08-27 11:41:16', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (14, '96ae1799-8468-3a79-a332-666560aee516', 'Theresia Hartmann', 'DarkCyan', 'system', '2021-08-27 11:42:08', 'system', '2021-08-27 11:55:24', true);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (15, 'fe9dfe9a-2969-3605-8060-86ac3bf2123a', 'Andreane Heller', 'Aquamarine', 'system', '2021-08-27 11:42:42', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (16, '164dee31-e33c-3812-8f43-18e64ba9a90b', 'Terrell Johns', 'PaleTurquoise', 'system', '2021-08-27 11:43:07', 'system', '2021-08-27 11:54:59', true);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (17, '3a26299d-a8fb-3fa0-94aa-b38b7ce80083', 'Dominique Farrell', 'Cyan', 'system', '2021-08-30 11:41:48', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (18, 'ce385835-495b-3e48-a364-63c040a61477', 'Keyon Towne', 'LightSalmon', 'system', '2021-08-30 11:42:05', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (19, 'c68bc67d-56ba-3a6e-b7bf-4c9a8ea5de9b', 'Mike Tremblay', 'GreenYellow', 'system', '2021-08-30 11:47:19', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (20, '3288c81d-fd96-34b1-808b-5704bd55cf32', 'Elta Berge', 'Crimson', 'system', '2021-08-30 11:48:59', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (21, '38bd49c4-461b-3026-980e-afea0bc434d5', 'Alexanne Wuckert', 'DimGrey', 'system', '2021-08-30 11:49:49', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (22, '5b35e0d5-fdc6-3049-8a6a-6f815c200d00', 'Geo Hettinger', 'Turquoise', 'system', '2021-08-30 11:51:51', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (23, '886cc026-e53d-38ad-8d55-cb565d849a7b', 'Mylene Johnston', 'SaddleBrown', 'system', '2021-08-30 11:52:03', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (24, 'ba50297a-d103-3044-a27e-5f07d24c87cf', 'Tia Blanda', 'Gainsboro', 'system', '2021-08-30 15:31:09', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (25, 'ef445ecc-e848-37dd-a4e6-fd284aedb429', 'Tracey DuBuque', 'IndianRed ', 'system', '2021-08-30 15:32:39', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (26, '5164f8ae-8e4a-3c22-9f36-7c8a385ebdb6', 'Ignacio Jewess', 'Sienna', 'system', '2021-08-30 15:33:57', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (27, '35698f21-32dd-37a6-8828-a483dec40c13', 'Jacques Cronin', 'Beige', 'system', '2021-08-30 17:28:08', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (28, 'dbb90ee9-8d45-340c-8f28-9496a7f3aefe', 'Aletha Kuvalis', 'FireBrick', 'system', '2021-08-30 23:02:21', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (29, '23123462-f076-3017-89d4-635be9b90d6f', 'Liana Renner', 'DarkGreen', 'system', '2021-08-30 23:05:26', 'system', '2021-08-30 23:06:43', false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (30, 'be48d516-ef6c-3ac9-a0e6-5571acf95921', 'Abbigail Lesch', 'Chocolate', 'system', '2021-08-31 15:12:56', 'system', '2021-08-31 15:22:36', false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (31, '245ead5b-578a-3648-bde8-f9ff42db0336', 'Nicole Nienow', 'DarkSlateGray', 'system', '2021-08-31 15:17:26', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (32, '69ede490-ca12-3ec7-9f3d-d896d3f82cc8', 'Chadrick Cole', 'LightGreen', 'system', '2021-08-31 15:18:33', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (33, 'd02dd2c0-a49e-3a90-87c6-c3f71d020b29', 'Dewayne Parisian', 'LemonChiffon', 'system', '2021-08-31 15:19:07', 'system', '2021-08-31 15:22:12', false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (34, '4a7bdf09-6744-357f-a0d8-08705e23fe73', 'Shana Barton', 'Cornsilk', '1212312121', '2021-08-31 22:10:42', '1212312121', '2021-08-31 22:13:19', true);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (35, '37cb0cbc-00d3-3a2b-a7c6-2589dd58882d', 'John Schmitt', 'MediumVioletRed', '1212312121', '2021-08-31 22:28:46', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (36, 'faa6ac27-3fa1-3cbb-abdd-96b7d64600ab', 'Cheyanne McCullough', 'Gray', '1212312121', '2021-08-31 22:29:49', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (37, '952d2a4d-b4c6-370f-b4e7-c70afbb5a523', 'Alexandra Leannon', 'DarkSeaGreen', '1212312121', '2021-08-31 22:38:15', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (38, '9ee30fb7-2266-3487-b598-6fb6e875ce4d', 'Howell Wilderman', 'Olive', '1212312121', '2021-09-01 10:28:33', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (39, '1d51839e-b417-3680-9897-a7a3f0e008f9', 'Scottie Hand', 'Darkorange', '1212312121', '2021-09-01 10:46:55', null, null, false);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (40, '33c37a0c-2c7f-3862-89bf-63bacd2f499a', 'Kallie Kemmer', 'Blue', '1212312121', '2021-09-01 11:38:14', '1212312121', '2021-09-01 11:40:50', true);
INSERT INTO john.duck (DuckID, DuckUUID, Name, Color, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (41, 'e8c3a5d7-afa5-384c-b3e0-987c7f0a2e45', 'Aidan Casper', 'DarkGray', '1212312121', '2021-09-01 21:11:46', null, null, false);