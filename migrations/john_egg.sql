create table egg
(
    EggID       int auto_increment
        primary key,
    DuckID      int                  not null,
    Name        text                 null,
    Age         int                  null,
    CreatedBy   text                 null,
    CreatedDate datetime             null,
    UpdatedBy   text                 null,
    UpdatedDate datetime             null,
    IsDeleted   tinyint(1) default 0 null
);

INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (1, 16, 'Eduardo Runolfsson', 1, 'system', '2021-08-27 11:43:07', 'system', '2021-08-27 11:54:59', 1);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (2, 16, 'Gerardo Pouros', 35, 'system', '2021-08-27 11:43:07', 'system', '2021-08-27 11:54:59', 1);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (3, 16, 'Davonte Sauer', 39, 'system', '2021-08-27 11:43:07', 'system', '2021-08-27 11:54:59', 1);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (4, 16, 'Layne Dooley', 27, 'system', '2021-08-27 11:43:07', 'system', '2021-08-27 11:54:59', 1);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (5, 16, 'Edna Schuster', 45, 'system', '2021-08-27 11:43:07', 'system', '2021-08-27 11:54:59', 1);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (6, 16, 'Samanta Turner', 28, 'system', '2021-08-27 11:43:07', 'system', '2021-08-27 11:54:59', 1);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (7, 16, 'Holly Emard', 43, 'system', '2021-08-27 11:52:41', 'system', '2021-08-27 11:54:59', 1);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (8, 16, 'Cyril Gaylord', 37, 'system', '2021-08-27 11:52:41', 'system', '2021-08-27 11:54:59', 1);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (9, 16, 'Liliane Gleason', 34, 'system', '2021-08-27 11:52:41', 'system', '2021-08-27 11:54:59', 1);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (10, 16, 'Rowan Weissnat', 18, 'system', '2021-08-27 11:52:41', 'system', '2021-08-27 11:54:59', 1);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (11, 16, 'Uriah Gutmann', 43, 'system', '2021-08-27 11:52:41', 'system', '2021-08-27 11:54:59', 1);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (12, 16, 'Evie Jakubowski', 40, 'system', '2021-08-27 11:52:41', 'system', '2021-08-27 11:54:59', 1);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (13, 16, 'Brooke OReilly', 21, 'system', '2021-08-27 11:52:41', 'system', '2021-08-27 11:54:59', 1);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (14, 16, 'Kobe Kunze', 22, 'system', '2021-08-27 11:52:41', 'system', '2021-08-27 11:54:59', 1);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (15, 16, 'Hilda Glover', 29, 'system', '2021-08-27 11:52:41', 'system', '2021-08-27 11:54:59', 1);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (16, 16, 'Marianne Kulas', 16, 'system', '2021-08-27 11:52:41', 'system', '2021-08-27 11:54:59', 1);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (17, 8, 'Felicity Koch', 28, 'system', '2021-08-27 11:56:11', 'system', '2021-08-30 10:56:19', 1);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (18, 8, 'Abagail Pfeffer', 37, 'system', '2021-08-27 11:56:11', 'system', '2021-08-30 10:56:19', 1);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (19, 8, 'Natalia Nader', 50, 'system', '2021-08-27 11:56:11', 'system', '2021-08-30 10:56:19', 1);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (20, 8, 'Reese Schultz', 25, 'system', '2021-08-27 11:56:11', 'system', '2021-08-30 10:56:19', 1);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (21, 8, 'Eliezer Emmerich', 38, 'system', '2021-08-30 10:56:19', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (22, 8, 'Ryley Watsica', 8, 'system', '2021-08-30 10:56:19', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (23, 8, 'Elliot Harvey', 40, 'system', '2021-08-30 10:56:19', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (24, 8, 'Rylee Stroman', 14, 'system', '2021-08-30 10:56:19', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (25, 8, 'Werner Robel', 21, 'system', '2021-08-30 10:56:19', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (26, 8, 'Alessandro Greenfelder', 36, 'system', '2021-08-30 10:56:19', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (27, 8, 'Jackson Franecki', 48, 'system', '2021-08-30 10:56:19', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (28, 17, 'Emerald Hermiston', 1, 'system', '2021-08-30 11:41:48', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (29, 18, 'Hassie Kessler', 47, 'system', '2021-08-30 11:42:05', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (30, 18, 'Candelario McGlynn', 46, 'system', '2021-08-30 11:42:05', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (31, 18, 'Zola Kohler', 27, 'system', '2021-08-30 11:42:05', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (32, 19, 'Alicia Muller', 35, 'system', '2021-08-30 11:47:19', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (33, 19, 'Reese Okuneva', 13, 'system', '2021-08-30 11:47:19', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (34, 19, 'Nikita Schroeder', 22, 'system', '2021-08-30 11:47:19', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (35, 20, 'Tristin Nicolas', 3, 'system', '2021-08-30 11:48:59', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (36, 20, 'Petra Legros', 20, 'system', '2021-08-30 11:48:59', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (37, 20, 'Glen Hoppe', 39, 'system', '2021-08-30 11:48:59', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (38, 20, 'Destinee Kemmer', 27, 'system', '2021-08-30 11:48:59', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (39, 20, 'Mariana Lehner', 48, 'system', '2021-08-30 11:48:59', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (40, 20, 'Joaquin Yost', 49, 'system', '2021-08-30 11:48:59', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (41, 20, 'Broderick Ledner', 25, 'system', '2021-08-30 11:48:59', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (42, 20, 'Raul Russel', 42, 'system', '2021-08-30 11:48:59', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (43, 20, 'Queenie Thompson', 4, 'system', '2021-08-30 11:48:59', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (44, 21, 'Elvie Kiehn', 6, 'system', '2021-08-30 11:49:49', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (45, 21, 'Murl Romaguera', 27, 'system', '2021-08-30 11:49:49', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (46, 21, 'Rafaela Lind', 44, 'system', '2021-08-30 11:49:49', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (47, 21, 'Mylene Cronin', 3, 'system', '2021-08-30 11:49:49', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (48, 21, 'Fritz Wilkinson', 23, 'system', '2021-08-30 11:49:49', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (49, 21, 'Emerson Schulist', 37, 'system', '2021-08-30 11:49:49', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (50, 22, 'Dulce Swift', 42, 'system', '2021-08-30 11:51:51', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (51, 22, 'Kaden Bernier', 40, 'system', '2021-08-30 11:51:51', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (52, 22, 'Filomena Kuvalis', 24, 'system', '2021-08-30 11:51:51', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (53, 22, 'Kimberly Barrows', 21, 'system', '2021-08-30 11:51:51', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (54, 22, 'Darian Fahey', 39, 'system', '2021-08-30 11:51:51', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (55, 22, 'Raymundo Bogan', 44, 'system', '2021-08-30 11:51:51', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (56, 22, 'Verona Littel', 47, 'system', '2021-08-30 11:51:51', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (57, 23, 'Loy DuBuque', 4, 'system', '2021-08-30 11:52:03', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (58, 23, 'April Ward', 44, 'system', '2021-08-30 11:52:03', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (59, 23, 'Baron Schneider', 38, 'system', '2021-08-30 11:52:03', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (60, 23, 'Charity Casper', 29, 'system', '2021-08-30 11:52:03', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (61, 23, 'Shanny Spencer', 12, 'system', '2021-08-30 11:52:03', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (62, 23, 'Diana Leffler', 24, 'system', '2021-08-30 11:52:03', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (63, 23, 'Daphne Parker', 2, 'system', '2021-08-30 11:52:03', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (64, 23, 'Fay Fahey', 12, 'system', '2021-08-30 11:52:03', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (65, 23, 'Savanna Franecki', 41, 'system', '2021-08-30 11:52:03', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (66, 24, 'Myrtle Lemke', 6, 'system', '2021-08-30 15:31:09', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (67, 24, 'Matilda Gibson', 2, 'system', '2021-08-30 15:31:09', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (68, 24, 'Maeve Casper', 12, 'system', '2021-08-30 15:31:09', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (69, 24, 'Lenora Runte', 46, 'system', '2021-08-30 15:31:09', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (70, 24, 'Adele Sanford', 18, 'system', '2021-08-30 15:31:09', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (71, 25, 'Dusty Pagac', 1, 'system', '2021-08-30 15:32:39', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (72, 25, 'Madaline Murphy', 48, 'system', '2021-08-30 15:32:39', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (73, 25, 'Connor Kshlerin', 16, 'system', '2021-08-30 15:32:39', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (74, 25, 'Arlie Renner', 11, 'system', '2021-08-30 15:32:39', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (75, 25, 'Arch Heidenreich', 21, 'system', '2021-08-30 15:32:39', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (76, 25, 'Robin Pollich', 2, 'system', '2021-08-30 15:32:39', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (77, 25, 'Georgianna Hackett', 10, 'system', '2021-08-30 15:32:39', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (78, 25, 'Kennith Aufderhar', 19, 'system', '2021-08-30 15:32:39', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (79, 25, 'Eudora Kautzer', 39, 'system', '2021-08-30 15:32:39', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (80, 25, 'Estel Barrows', 37, 'system', '2021-08-30 15:32:39', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (81, 26, 'Darian Osinski', 36, 'system', '2021-08-30 15:33:57', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (82, 26, 'Marion Yost', 48, 'system', '2021-08-30 15:33:57', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (83, 26, 'Diego Mayert', 42, 'system', '2021-08-30 15:33:57', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (84, 26, 'Evert Ledner', 41, 'system', '2021-08-30 15:33:57', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (85, 26, 'Frances Durgan', 24, 'system', '2021-08-30 15:33:57', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (86, 26, 'Tamia Eichmann', 25, 'system', '2021-08-30 15:33:57', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (87, 26, 'Nora Adams', 43, 'system', '2021-08-30 15:33:57', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (88, 26, 'Alisa Tillman', 5, 'system', '2021-08-30 15:33:57', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (89, 26, 'Everette Mosciski', 27, 'system', '2021-08-30 15:33:57', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (90, 27, 'Vilma McClure', 40, 'system', '2021-08-30 17:28:08', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (91, 27, 'Lisandro Sanford', 3, 'system', '2021-08-30 17:28:08', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (92, 27, 'Garett Rohan', 25, 'system', '2021-08-30 17:28:08', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (93, 27, 'Elvera Balistreri', 2, 'system', '2021-08-30 17:28:08', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (94, 27, 'Matilde Ritchie', 27, 'system', '2021-08-30 17:28:08', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (95, 27, 'Name Grimes', 19, 'system', '2021-08-30 17:28:08', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (96, 27, 'Dessie Emard', 13, 'system', '2021-08-30 17:28:08', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (97, 27, 'Vinnie Emard', 22, 'system', '2021-08-30 17:28:08', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (98, 28, 'Hayley Murphy', 18, 'system', '2021-08-30 23:02:21', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (99, 28, 'Kieran Pouros', 24, 'system', '2021-08-30 23:02:21', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (100, 28, 'Jamar Dare', 30, 'system', '2021-08-30 23:02:21', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (101, 28, 'Dan Jaskolski', 44, 'system', '2021-08-30 23:02:21', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (102, 28, 'Araceli Lind', 15, 'system', '2021-08-30 23:02:21', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (103, 28, 'Emie Schuster', 33, 'system', '2021-08-30 23:02:21', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (104, 28, 'Jordan Hamill', 32, 'system', '2021-08-30 23:02:21', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (105, 28, 'Louvenia Rice', 12, 'system', '2021-08-30 23:02:21', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (106, 28, 'Philip Shanahan', 20, 'system', '2021-08-30 23:02:21', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (107, 29, 'Al Schroeder', 34, 'system', '2021-08-30 23:05:26', 'system', '2021-08-30 23:06:43', 1);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (108, 29, 'Kiley Hirthe', 25, 'system', '2021-08-30 23:05:26', 'system', '2021-08-30 23:06:43', 1);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (109, 29, 'Ismael Metz', 27, 'system', '2021-08-30 23:05:26', 'system', '2021-08-30 23:06:43', 1);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (110, 29, 'Dustin Konopelski', 34, 'system', '2021-08-30 23:05:26', 'system', '2021-08-30 23:06:43', 1);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (111, 29, 'Roman Reichert', 48, 'system', '2021-08-30 23:05:26', 'system', '2021-08-30 23:06:43', 1);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (112, 29, 'Glen Mills', 32, 'system', '2021-08-30 23:05:26', 'system', '2021-08-30 23:06:43', 1);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (113, 29, 'Gayle Mante', 35, 'system', '2021-08-30 23:05:26', 'system', '2021-08-30 23:06:43', 1);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (114, 29, 'Gloria Glover', 43, 'system', '2021-08-30 23:06:43', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (115, 29, 'Shyanne Lockman', 1, 'system', '2021-08-30 23:06:43', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (116, 29, 'Colin Schmitt', 35, 'system', '2021-08-30 23:06:43', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (117, 29, 'Ashley Cummerata', 8, 'system', '2021-08-30 23:06:43', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (118, 29, 'Odell Ferry', 35, 'system', '2021-08-30 23:06:43', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (119, 29, 'Albert Harris', 12, 'system', '2021-08-30 23:06:43', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (120, 29, 'Jesus Stokes', 21, 'system', '2021-08-30 23:06:43', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (121, 29, 'Gussie Quitzon', 21, 'system', '2021-08-30 23:06:43', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (122, 29, 'Orie Hermiston', 6, 'system', '2021-08-30 23:06:43', null, null, 0);
INSERT INTO john.egg (EggID, DuckID, Name, Age, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (123, 29, 'Justen Haley', 43, 'system', '2021-08-30 23:06:43', null, null, 0);