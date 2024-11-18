DROP TABLE IF EXISTS crimes;

CREATE TABLE CRIMES (
    id SERIAL PRIMARY KEY,
    nome_crime VARCHAR(100),
    descricao TEXT,
    data_crime DATE,
    heroi_responsavel INT REFERENCES HEROI(codigo_heroi) ON DELETE CASCADE,
    severidade INT,
    oculto BOOLEAN DEFAULT FALSE
);

SELECT * FROM CRIMES;

ALTER TABLE CRIMES ALTER COLUMN SEVERIDADE TYPE TEXT;

ALTER TABLE crimes DROP CONSTRAINT crimes_heroi_responsavel_fkey;

ALTER TABLE crimes
ADD CONSTRAINT crimes_heroi_responsavel_fkey
FOREIGN KEY (heroi_responsavel) REFERENCES heroi(codigo_heroi) ON DELETE CASCADE;

delete from crimes;
