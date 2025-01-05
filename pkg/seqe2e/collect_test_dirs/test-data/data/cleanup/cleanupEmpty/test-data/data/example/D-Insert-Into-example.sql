-- +dbevo Up
INSERT INTO example
VALUES
(1, 'name-1', 'user-name-1'),
(2, 'name-3', 'user-name-3')
;

-- +dbevo Down
DELETE FROM example
WHERE id in (1, 2)