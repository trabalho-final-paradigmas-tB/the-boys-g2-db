DROP TABLE IF EXISTS crimes;

CREATE TABLE CRIMES (
    ID SERIAL PRIMARY KEY,
    NOME_CRIME VARCHAR(100),
    DESCRICAO TEXT,
    DATA_CRIME DATE,
    HEROI_RESPONSAVEL INT,
    SEVERIDADE INT
);
