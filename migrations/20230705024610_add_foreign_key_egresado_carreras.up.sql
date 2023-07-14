-- Agregar columna carrera_id en la tabla egresado
ALTER TABLE students ADD career_id UUID;

-- Establecer la relación entre las tablas
ALTER TABLE students
ADD CONSTRAINT fk_students_career
FOREIGN KEY (career_id) REFERENCES careers(id);