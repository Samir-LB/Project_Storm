ALTER TABLE profesor
DROP COLUMN carrera_id;

ALTER TABLE profesor
DROP CONSTRAINT fk_profesor_carrera;
