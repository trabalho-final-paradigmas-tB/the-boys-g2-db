-- Tabela dos heróis (Sexo: F/M)

DROP TABLE IF EXISTS herois;

CREATE TABLE herois (
    CODIGO_HEROI        INT AUTO_INCREMENT,
    NOME_REAL           VARCHAR(80)     NULL,
    NOME_HEROI          VARCHAR(80)     NOT NULL,
    SEXO                VARCHAR(1)      NOT NULL,
    ALTURA_HEROI        FLOAT           NOT NULL,
    PESO_HEROI          FLOAT           NOT NULL,
    DATA_NASCIMENTO     DATE            NOT NULL,
    LOCAL_NASCIMENTO    VARCHAR(80)     NULL,
    PODERES             VARCHAR(80)     NOT NULL,
    NIVEL_FORCA         INT             NOT NULL,
    POPULARIDADE        INT             NOT NULL,
    STATUS              VARCHAR(7)      NOT NULL,
    HISTORICO_BATALHAS  VARCHAR(80)     NOT NULL,
    PRIMARY KEY (CODIGO_HEROI)
);

SELECT * FROM herois;

CREATE TRIGGER banir_heroi
BEFORE INSERT OR UPDATE ON herois
FOR EACH ROW
BEGIN
    IF NEW.POPULARIDADE <= 20 THEN
        SET NEW.STATUS = 'BANIDO';
    END IF;
END;

