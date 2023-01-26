
# Aplikasi Rental PS

#### Oleh : Abdurrahman Shalahudin Zaka

#### Link Railway : https://fp-sanbercode-golang-batch-41-production.up.railway.app




## Endpoints

#### User

| Method   | Path          | Description                |
|:---------|:--------------|:---------------------------|
| `GET`    | `/user`       | Menampilkan daftar user    |
| `POST`   | `/user/user`  | Menambahkan pelanggan baru |
| `POST`   | `/user/admin` | Menambahkan admin baru     |
| `PUT`    | `/user/:id`   | Mengupdate data user       |
| `DELETE` | `/user/:id`   | Menghapus data user        |

#### PS

| Method   | Path      | Description           |
|:---------|:----------|:----------------------|
| `GET`    | `/ps`     | Menampilkan daftar PS |
| `POST`   | `/ps`     | Menambahkan PS baru   |
| `PUT`    | `/ps/:id` | Mengupdate data PS    |
| `DELETE` | `/ps/:id` | Menghapus data PS     |

#### Harga

| Method   | Path      | Description              |
|:---------|:----------|:-------------------------|
| `GET`    | `/ps`     | Menampilkan daftar Harga |
| `POST`   | `/ps`     | Menambahkan Harga baru   |
| `PUT`    | `/ps/:id` | Mengupdate data Harga    |
| `DELETE` | `/ps/:id` | Menghapus data Harga     |

#### Peminjaman

| Method   | Path              | Description                   |
|:---------|:------------------|:------------------------------|
| `GET`    | `/peminjaman`     | Menampilkan daftar Peminjaman |
| `POST`   | `/peminjaman`     | Menambahkan Peminjaman baru   |
| `PUT`    | `/peminjaman/:id` | Mengupdate data Peminjaman    |
| `DELETE` | `/peminjaman/:id` | Menghapus data Peminjaman     |