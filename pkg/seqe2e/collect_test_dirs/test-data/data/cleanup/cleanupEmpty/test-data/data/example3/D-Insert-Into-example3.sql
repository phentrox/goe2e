-- +dbevo Up
INSERT INTO example3
VALUES (1, 'name-1');

-- +dbevo Down
DELETE FROM example3
WHERE id = 1;