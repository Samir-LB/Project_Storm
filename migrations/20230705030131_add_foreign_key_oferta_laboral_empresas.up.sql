-- Agregar columna empresa_id en la tabla oferta_laboral
ALTER TABLE work_offers
ADD company_id UUID;

-- Establecer la relación entre las tablas
-- ALTER TABLE work_offers
-- ADD CONSTRAINT fk_work_offer_company
-- FOREIGN KEY (company_id) REFERENCES companies(id);
