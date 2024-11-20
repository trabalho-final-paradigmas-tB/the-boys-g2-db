DROP TABLE IF EXISTS HEROI;

CREATE TABLE HEROI (
    CODIGO_HEROI        BIGSERIAL PRIMARY KEY,  
    NOME_REAL           VARCHAR(80),
    NOME_HEROI          VARCHAR(80)     NOT NULL,
    SEXO                VARCHAR(1)      NOT NULL CHECK (SEXO IN ('F', 'M')),
    ALTURA       		FLOAT           NOT NULL,
    PESO          		FLOAT           NOT NULL,
    DATA_NASCIMENTO     DATE            NOT NULL,
    LOCAL_NASCIMENTO    VARCHAR(80),
    PODERES             TEXT            NOT NULL,
    NIVEL_FORCA         INT             NOT NULL CHECK (NIVEL_FORCA BETWEEN 0 AND 100),
    POPULARIDADE        INT             NOT NULL CHECK (POPULARIDADE BETWEEN 0 AND 100),
    STATUS              VARCHAR(7)      NOT NULL CHECK (STATUS IN ('Ativo', 'Inativo', 'Banido')),
    HISTORICO_BATALHAS  INTEGER[]       NOT NULL
);

-- Trigger para atualizar o status dos heróis com base na popularidade

CREATE OR REPLACE FUNCTION atualizar_status_heroi()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.POPULARIDADE < 20 THEN
        NEW.STATUS := 'Banido';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER banir_heroi
BEFORE INSERT OR UPDATE ON HEROI
FOR EACH ROW
EXECUTE FUNCTION atualizar_status_heroi();

-- Consulta para verificar a tabela criada

delete from heroi;

select * from heroi;

SELECT * FROM information_schema.tables;

SELECT column_name, data_type 
FROM information_schema.columns
WHERE table_name = 'heroi';