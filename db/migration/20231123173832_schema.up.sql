CREATE TABLE "produk" (
  "id_produk" varchar PRIMARY KEY,
  "nama_produk" varchar NOT NULL,
  "harga" varchar NOT NULL,
  "kategori_id" varchar NOT NULL,
  "status_id" varchar NOT NULL
);

CREATE TABLE "kategori" (
  "id_kategori" varchar PRIMARY KEY ,
  "nama_kategori" varchar NOT NULL
);

CREATE TABLE "status" (
  "id_status" varchar PRIMARY KEY ,
  "nama_status" varchar NOT NULL
);

ALTER TABLE "produk" ADD FOREIGN KEY ("kategori_id") REFERENCES "kategori" ("id_kategori");

ALTER TABLE "produk" ADD FOREIGN KEY ("status_id") REFERENCES "status" ("id_status");
