CREATE EXTENSION IF NOT EXISTS 'pgcrypto';

CREATE TABLE USUARIO(
    USU_CODIGO SERIAL NOT NULL PRIMARY KEY,
    USU_EMAIL VARCHAR(100) NOT NULL,
    USU_NOME VARCHAR(100) NOT NULL,
    USU_SENHA VARCHAR(100) NOT NULL,
    USU_DATACRIACAO TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    USU_RECEBERALERTA BOOL NOT NULL DEFAULT FALSE,
);

CREATE UNIQUE INDEX UI_USU_EMAIL ON USUARIO (USU_EMAIL);