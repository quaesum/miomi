INSERT INTO animals
(age, name, sex, description, sterilized, vaccinated, onrainbow, onhappines)
VALUES(12, 'Василиса',1 , 'Кошечка Василиса ищет новых хозяев', 1, 0, 0, 0),
      (5, 'Ибрагим', 0, 'Кот Ибрагим ищет новых хозяев', 1, 0, 0, 0),
      (7, 'Рыжуля', 1, 'Кошечка Рыжуля ищет новых хозяев', 0, 1, 0, 0),
      (2, 'Кнопка', 1, 'Кошечка ищет новых хозяев', 1, 0, 0, 0),
      (1, 'Масяня', 1, 'Кошечка Масяня ищет новых хозяев', 0, 0, 0, 0),
      (3, 'Гарфилд', 0, 'Кот Гарфилд ищет новых хозяев', 1, 0, 0, 0),
      (5, 'Наоми', 1, 'Кошечка Наоми ищет новых хозяев', 0, 0, 0, 0),
      (11, 'Изольда', 1, 'Кошечка Изольда ищет новых хозяев', 1, 0, 0, 0),
      (9, 'Гавриил', 0, 'Кот Гавриил ищет новых хозяев', 0, 0, 0, 0),
      (2, 'Бьянка', 1, 'Кошечка Бьянка ищет новых хозяев', 0, 0, 0, 0),
      (2, 'Аскель', 0, 'Аскель ищет новых хозяев', 0, 1, 0, 0),
      (5, 'Локи', 0, 'Локи ищет новых хозяев', 0, 0, 0, 0),
      (7, 'Тор', 0, 'Тор ищет новых хозяев', 1, 1, 0, 0),
      (10, 'Френк', 0, 'Френк ищет новых хозяев', 0, 1, 0, 0),
      (8, 'Флоренс', 0, 'Флоренс ищет новых хозяев', 0, 0, 0, 0),
      (4, 'Маверик', 0, 'Маверик ищет новых хозяев', 0, 0, 0, 0),
      (3, 'Пинки', 1, 'Пинки ищет новых хозяев', 0, 1, 0, 0),
      (8, 'Дино', 0, 'Дино ищет новых хозяев', 1, 0, 0, 0),
      (1, 'Ви', 1, 'Ви ищет новых хозяев', 0, 0, 0, 0),
      (1, 'Брумхильда', 1, 'Брумхильда ищет новых хозяев', 0, 0, 0, 0);

INSERT INTO animal_shelters
(shelter_name, description, logo, phone, adress)
VALUES('Фауна Города', 'Самый крупный центр передержки животных в Минске', '', '+375(29)123-12-12', 'Тышкевичей 1'),
      ('Крутой приют в Минске', 'Крутой приют в минске', '', '+375(29)123-21-21', 'Пушкина 12'),
      ('Большой приют в Минске', 'Большой приют в Минске', '', '+375(29)123-13-32', 'Центральная 33');

INSERT INTO volunteers
(firstName, lastName, password, email, deactevated, createdAt)
VALUES('Максим', 'Пожарицкий', '', 'pozaritskiy@post.com', 0, 0),
      ('Алиса', 'Смелянская', '', 'smelyanskaya@post.com', 0, 0);

INSERT INTO volunteers_on_shelters
(volunteerID, shelterID)
VALUES(1, 2),
      (2,3),
      (1,1),
      (2,1);

INSERT INTO animals_on_shelters
(animalID, shelterID)
VALUES(1, 1),
      (2, 1),
      (3, 3),
      (4, 2),
      (5, 1),
      (6, 3),
      (7, 2),
      (8, 2),
      (9, 3),
      (10, 3),
      (11, 3),
      (12, 3),
      (13, 1),
      (14, 3),
      (15, 2),
      (16, 3),
      (17, 1),
      (18, 2),
      (19, 1),
      (20, 3);

INSERT INTO miomi.animal_types
(name)
VALUES('Кот'),
      ('Собака'),
      ('Птица'),
      ('Другое');

INSERT INTO animals_on_types
(animal_typeID, animalID)
VALUES(1, 1),
      (1, 2),
      (1, 3),
      (1, 4),
      (1, 5),
      (1, 6),
      (1, 7),
      (1, 8),
      (1, 9),
      (1, 10),
      (2, 11),
      (2, 12),
      (2, 13),
      (2, 14),
      (2, 15),
      (2, 16),
      (2, 17),
      (2, 18),
      (2, 19),
      (2, 20);

INSERT INTO photos
(filename, origin, file_type)
VALUES ('/news/news1.jpeg', '', 'jpeg'),
       ('/news/news2.jpeg', '', 'jpeg'),
       ('/news/news3.jpeg', '', 'jpeg'),
       ('/news/news4.jpeg', '', 'jpeg'),
       ('/news/news5.jpeg', '', 'jpeg'),
       ('/animals/cat1.jpeg', '', 'jpeg'),
       ('/animals/cat2.jpg', '', 'jpg'),
       ('/animals/cat3.jpg', '', 'jpg'),
       ('/animals/cat4.jpeg', '', 'jpeg'),
       ('/animals/cat5.jpg', '', 'jpg'),
       ('/animals/cat6.jpg', '', 'jpg'),
       ('/animals/cat7.jpg', '', 'jpg'),
       ('/animals/cat8.jpeg', '', 'jpeg'),
       ('/animals/cat9.jpg', '', 'jpg'),
       ('/animals/cat10.jpg', '', 'jpg'),
       ('/animals/dog1.jpeg', '', 'jpeg'),
       ('/animals/dog2.jpeg', '', 'jpeg'),
       ('/animals/dog3.jpeg', '', 'jpeg'),
       ('/animals/dog4.jpeg', '', 'jpeg'),
       ('/animals/dog5.jpeg', '', 'jpeg'),
       ('/animals/dog6.jpeg', '', 'jpeg'),
       ('/animals/dog7.jpeg', '', 'jpeg'),
       ('/animals/dog8.jpeg', '', 'jpeg'),
       ('/animals/dog9.jpeg', '', 'jpeg'),
       ('/animals/dog10.jpeg', '', 'jpeg'),
       ('/animals/cat11.jpeg', '', 'jpeg'),
       ('/animals/cat12.jpg', '', 'jpg'),
       ('/animals/cat13.jpeg', '', 'jpeg'),
       ('/animals/cat14.jpeg', '', 'jpeg'),
       ('/animals/cat15.jpeg', '', 'jpeg'),
       ('/animals/dog11.jpeg', '', 'jpeg'),
       ('/animals/dog12.jpeg', '', 'jpeg'),
       ('/animals/dog13.jpeg', '', 'jpeg'),
       ('/animals/dog14.jpeg', '', 'jpeg'),
       ('/animals/dog15.jpeg', '', 'jpeg');

INSERT INTO animals_photos
(animalID, photoID)
VALUES(1,1),
      (2, 2),
      (3, 3),
      (4, 4),
      (5, 5),
      (6, 6),
      (7, 7),
      (8, 8),
      (9, 9),
      (10, 10),
      (1, 4),
      (2, 5),
      (3, 6),
      (4, 7);

INSERT INTO news_photos
(newsID, photoID)
VALUES(1, 61),
      (2, 62),
      (3, 63),
      (4, 64),
      (5, 65);
INSERT INTO miomi.news
(label, description)
VALUES('Приглашаем на мероприятие!', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin condimentum purus eget ligula pharetra, ut sodales magna scelerisque. Aenean egestas, felis non semper consequat, nulla justo gravida arcu, at consectetur odio turpis id leo. In vel nisi a nibh molestie tempor ut sed nisi. Vestibulum eu pellentesque magna. Phasellus molestie finibus lacus, a pellentesque libero tincidunt eu. Proin eget fermentum diam. Sed pharetra dictum tortor, in sagittis purus gravida ut. Pellentesque aliquam pulvinar urna, at pharetra lectus maximus vitae. Suspendisse tristique mattis interdum. '),
      ('Не приглашаем на мероприятие', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin condimentum purus eget ligula pharetra, ut sodales magna scelerisque. Aenean egestas, felis non semper consequat, nulla justo gravida arcu, at consectetur odio turpis id leo. In vel nisi a nibh molestie tempor ut sed nisi. Vestibulum eu pellentesque magna. Phasellus molestie finibus lacus, a pellentesque libero tincidunt eu. Proin eget fermentum diam. Sed pharetra dictum tortor, in sagittis purus gravida ut. Pellentesque aliquam pulvinar urna, at pharetra lectus maximus vitae. Suspendisse tristique mattis interdum. '),
      ('Возьми питомца и получи шоколадку', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin condimentum purus eget ligula pharetra, ut sodales magna scelerisque. Aenean egestas, felis non semper consequat, nulla justo gravida arcu, at consectetur odio turpis id leo. In vel nisi a nibh molestie tempor ut sed nisi. Vestibulum eu pellentesque magna. Phasellus molestie finibus lacus, a pellentesque libero tincidunt eu. Proin eget fermentum diam. Sed pharetra dictum tortor, in sagittis purus gravida ut. Pellentesque aliquam pulvinar urna, at pharetra lectus maximus vitae. Suspendisse tristique mattis interdum. '),
      ('Международный день собак', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin condimentum purus eget ligula pharetra, ut sodales magna scelerisque. Aenean egestas, felis non semper consequat, nulla justo gravida arcu, at consectetur odio turpis id leo. In vel nisi a nibh molestie tempor ut sed nisi. Vestibulum eu pellentesque magna. Phasellus molestie finibus lacus, a pellentesque libero tincidunt eu. Proin eget fermentum diam. Sed pharetra dictum tortor, in sagittis purus gravida ut. Pellentesque aliquam pulvinar urna, at pharetra lectus maximus vitae. Suspendisse tristique mattis interdum. '),
      ('Астрономы педсказывают лучший день, чтобы взять питомца из приюта', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin condimentum purus eget ligula pharetra, ut sodales magna scelerisque. Aenean egestas, felis non semper consequat, nulla justo gravida arcu, at consectetur odio turpis id leo. In vel nisi a nibh molestie tempor ut sed nisi. Vestibulum eu pellentesque magna. Phasellus molestie finibus lacus, a pellentesque libero tincidunt eu. Proin eget fermentum diam. Sed pharetra dictum tortor, in sagittis purus gravida ut. Pellentesque aliquam pulvinar urna, at pharetra lectus maximus vitae. Suspendisse tristique mattis interdum. ');
