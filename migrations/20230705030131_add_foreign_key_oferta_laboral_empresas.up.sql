-- Agregar columna empresa_id en la tabla oferta_laboral
ALTER TABLE oferta_laboral
ADD empresa_id UUID;

-- Establecer la relación entre las tablas
ALTER TABLE oferta_laboral
ADD CONSTRAINT fk_oferta_laboral_empresa
FOREIGN KEY (empresa_id) REFERENCES empresa(id);
