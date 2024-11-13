DROP  TABLE IF missoes EXISTS

CREATE TABLE missoes (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    descricao TEXT NOT NULL,
    classificacao VARCHAR(50),
    dificuldade INT CHECK (dificuldade >= 1 AND dificuldade <= 10)
);