ALTER TABLE oferta_laboral
DROP COLUMN empresa_id;

ALTER TABLE oferta_laboral
DROP CONSTRAINT fk_oferta_laboral_empresa;