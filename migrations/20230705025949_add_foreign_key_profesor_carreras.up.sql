-- Agregar columna carrera_id en la tabla profesor
ALTER TABLE profesor
ADD carrera_id UUID;

-- Establecer la relación entre las tablas
ALTER TABLE profesor
ADD CONSTRAINT fk_profesor_carrera
FOREIGN KEY (carrera_id) REFERENCES carrera(id);