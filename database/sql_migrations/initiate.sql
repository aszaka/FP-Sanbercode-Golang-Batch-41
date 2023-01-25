-- +migrate Down

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS master_ps;
DROP TABLE IF EXISTS peminjaman;
DROP TABLE IF EXISTS harga;

-- +migrate Up
-- +migrate statementBegin

CREATE TABLE users(
                      id_user SERIAL NOT NULL PRIMARY KEY,
                      username VARCHAR(256),
                      password VARCHAR(256),
                      level INT NOT NULL
);

CREATE TABLE master_ps(
                          id_ps SERIAL NOT NULL PRIMARY KEY,
                          tipe_ps VARCHAR(256),
                          daftar_game VARCHAR(500),
                          status_peminjaman INT NOT NULL
);

CREATE TABLE harga(
                      id_harga SERIAL NOT NULL PRIMARY KEY,
                      waktu INT NOT NULL,
                      satuan VARCHAR NOT NULL,
                      harga INT NOT NULL
);

CREATE TABLE peminjaman(
                           id_peminjaman SERIAL NOT NULL PRIMARY KEY,
                           id_ps BIGINT NOT NULL,
                           id_user BIGINT NOT NULL,
                           id_harga BIGINT NOT NULL,
                           peminjaman TIMESTAMP,
                           pengembalian TIMESTAMP,
                           keterangan VARCHAR(500)
);

-- +migrate statementEnd