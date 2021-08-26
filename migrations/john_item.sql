create table item
(
    id          int auto_increment
        primary key,
    uuid        text                 null,
    name        text                 null,
    CreatedBy   text                 null,
    CreatedDate datetime             null,
    UpdatedBy   text                 null,
    UpdatedDate datetime             null,
    IsDeleted   tinyint(1) default 0 null
);

INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (-6, '5d6cb2ab-f02f-4a0e-9cdf-6cfc0173230d', 'Dashawn Tillman', null, null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (1, '123123123', 'john', 'john', null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (2, '585647342', 'jane', 'john', null, null, '2021-08-23 22:46:05', 1);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (120, '726d6cf7-4118-487d-9e2b-a3d7b0dc7e8d', 'john2', null, null, null, '2021-08-25 20:50:43', 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (121, '47c76b1d-1a20-4978-a1dc-49cfe846f3e7', 'Earnest Green', 'John', null, null, '2021-08-26 11:07:00', 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (122, 'a0269be3-4b79-4d00-935a-d98c73a79ae9', 'Edwin Bechtelar', null, null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (123, '27534f15-1213-46b3-8280-2a336d01360a', 'Rey Heller', null, null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (124, '2479ad02-2ab5-4818-aea5-40d1cbdcf27d', 'Emmy Abshire', null, null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (125, 'd5121081-2564-4d52-a9b7-6295f87aa8e4', 'Billie Gottlieb', null, null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (126, '858dfdcc-9e74-44d3-903a-52533e9fc7c9', 'Garett Pagac', null, null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (127, '6b93c263-74ee-4717-a442-674f048b6dfd', 'Cortez Ritchie', null, null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (128, '844f0501-beca-401c-b302-3c3d6d37a9ef', 'Howell McLaughlin', null, null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (129, 'e6683551-c11c-4632-88f4-c72b849810c0', 'Lolita Carter', null, null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (130, '6862caae-b292-40fd-bd2d-578da15c6c08', 'Burley Wolff', null, null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (131, 'd41b3a1e-2c11-4bba-88ce-c4dfa552cf72', 'Tracey Bergstrom', null, null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (132, '674138a2-b26c-44a9-987a-48924c0014a2', 'Myrtis Prohaska', null, null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (133, '0e395f81-6614-4674-97ca-fc6eac3ff200', 'Aidan Feest', null, null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (134, '389cc4a6-d219-4284-8075-336b086669a4', 'Else Jenkins', null, null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (135, 'e4ca7d98-3d06-4eac-8025-c8e39af0a721', 'Katarina Lueilwitz', null, null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (136, 'ccb2a0a6-8ca8-4be0-8627-d75a247e4bc4', 'Maya Wilkinson', null, null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (137, '66ea360e-48ea-4014-9efb-d00337459759', 'Adeline Sanford', null, null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (138, 'c28dee33-5b15-44e9-859e-aa60ff69cde0', 'Gus Wilderman', null, null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (139, '35b4b001-f830-4ce4-b901-a85e3e8856e2', 'Giovani Crooks', null, null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (140, '4f0d1661-dbb0-4913-a7b6-5fdc37696efc', 'Elwin Kozey', null, null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (141, '94286db4-fe60-41a0-aa1f-4d2d8ccd3c6a', 'Arnaldo Nienow', null, null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (142, '1bad14b3-cfdc-4049-9d82-51d5955c4b0f', 'Hellen Koss', null, null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (143, 'eb16b51c-4a5c-433a-bbcc-d886c376105d', 'Jolie Harris', null, null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (144, '0397eb5f-01ce-459d-9750-f02c339794a3', 'Astrid Herzog', 'system', null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (145, '46440b74-635f-4f2b-8807-e1ea41233ee5', 'Thad OReilly', 'system', null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (146, '50513889-e64d-4ce8-9fd8-b673b8d1d352', 'Meredith Howe', 'system', null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (147, '33f47b71-9fd2-441c-a738-851e3568fed5', 'Cade Schulist', 'system', null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (148, 'ecc03df0-b474-427d-938e-e2c082b7a56d', 'Esta Crooks', 'system', null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (149, 'e67693a3-9b68-443e-b1b5-bc8172930d39', 'Alta Schinner', 'system', null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (150, '75449ad7-9b60-4477-8560-4e89b963413b', 'Constantin Romaguera', 'system', null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (151, '34f6b889-71f1-4816-8867-54c26fb102c3', 'Judy Hettinger', 'system', null, null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (152, 'c6c904f1-d7ee-4f0e-93d2-2acd7c49dbb6', 'Heidi Beier', 'system', '2021-08-26 13:53:06', null, null, 0);
INSERT INTO john.item (id, uuid, name, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate, IsDeleted) VALUES (153, '17ab0e1b-87de-46ed-84ce-62453be249b5', 'updated name', 'system', '2021-08-26 13:57:34', 'system', '2021-08-26 14:19:36', 1);