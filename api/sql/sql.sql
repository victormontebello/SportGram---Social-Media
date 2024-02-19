CREATE DATABASE IF NOT EXISTS sportgram;
USE sportgram;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    nick VARCHAR(255) NOT NULL unique,
    email VARCHAR(255) NOT NULL,
    senha VARCHAR(100) NOT NULL,
    criadoEm TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    esporte VARCHAR(255) NOT NULL,
    anosExperiencia INT NOT NULL,
    possuiPatrocinio BOOLEAN
) ENGINE=INNODB;

CREATE TABLE seguidores (
    usuario_id INT NOT NULL,
    FOREIGN KEY (usuario_id) 
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    seguidor_id INT NOT NULL,
    FOREIGN KEY (seguidor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE

    PRIMARY KEY (usuario_id, seguidor_id)
) ENGINE=INNODB;
