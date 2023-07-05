-- Agregar columna carrera_id en la tabla egresado
ALTER TABLE egresado ADD carrera_id UUID;

-- Establecer la relación entre las tablas
ALTER TABLE egresado
ADD CONSTRAINT fk_egresado_carrera
FOREIGN KEY (carrera_id) REFERENCES carrera(id);