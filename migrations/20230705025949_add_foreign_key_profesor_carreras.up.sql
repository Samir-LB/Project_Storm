-- Agregar columna carrera_id en la tabla profesor
ALTER TABLE teachers
ADD career_id UUID;

-- Establecer la relación entre las tablas
ALTER TABLE teachers
ADD CONSTRAINT fk_teacher_career
FOREIGN KEY (career_id) REFERENCES careers(id);