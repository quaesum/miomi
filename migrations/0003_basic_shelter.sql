-- +goose Up
INSERT INTO animal_shelters
    (shelter_name, description, logo, phone, adress)
    VALUES ('Фауна города','КОММУНАЛЬНОЕ ПРОИЗВОДСТВЕННОЕ УНИТАРНОЕ ПРЕДПРИЯТИЕ "ФАУНА ГОРОДА" (ГОСУДАРСТВЕННОЕ ПРЕДПРИЯТИЕ "ФАУНА ГОРОДА")', '', '+375173771793', 'Гурского, 42');

-- +goose Down