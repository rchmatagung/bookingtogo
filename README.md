
# Technical Test Booking To Go

Berikut merupakan hasil technical test yang saya buat menggunakan Golang... Silahkan lihat dokumentasi dibawah ini



## Menjalankan di Local Device

Clone Project ke Local Device

```bash
  git clone https://github.com/rchmatagung/bookingtogo.git
```

Masuk kedalam repository yang sudah di clone

```bash
  cd bookingtogo
```

melakukan install depedency pada go mod

```bash
  go mod tidy
```

Sebelum melakukan run di local silahkan import terlebih dahulu sql yang sudah diberikan di email yang sudah saya kirimkan

Start the server

```bash
  go run cmd/api/main.go
```


## Referensi API

### Nationality (Kewarganegaraan)

#### Membuat Kewarganegaraan

```http
  POST /nationality
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `nationality_name` | `string` | **Required** |
| `nationality_code` | `string` | **Required** |

#### Mendapatkan Semua Kewarganegaraan

```http
  GET /nationality
```

### Customer (Pelanggan)

#### Membuat Customer

```http
  POST /customer
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `nationality_id` | `integer` | **Required** |
| `customer_name` | `string` | **Required** |
| `customer_dob` | `string` | **Required** |
| `customer_phone` | `string` | **Required** |
| `customer_email` | `string` | **Required** |
| `family_list` | `Array Object` | **Required** |

Parameter Array Object family_list

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `family_list_relation` | `string` | **Required** |
| `family_list_name` | `string` | **Required** |
| `family_list_dob` | `string` | **Required** |

#### Mendapatkan Semua Customer

```http
  GET /customer
```

#### Mendapatkan Customer Berdasarkan Customer ID

```http
  GET /customer/{customerId}
```

#### Update Customer & Family List

```http
  PUT /customer
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `customer_id` | `integer` | **Required** |
| `nationality_id` | `integer` | **Required** |
| `customer_name` | `string` | **Required** |
| `customer_dob` | `string` | **Required** |
| `customer_phone` | `string` | **Required** |
| `customer_email` | `string` | **Required** |
| `family_list` | `Array Object` | **Required** |

Parameter Array Object family_list

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `family_list_id` | `integer` | **Required** |
| `family_list_relation` | `string` | **Required** |
| `family_list_name` | `string` | **Required** |
| `family_list_dob` | `string` | **Required** |
| `is_delete` | `boolean` | **Required** |

#### Catatan
Apabila parameter family_list_id = 0 maka akan menambahkan data keluarga, jika bukan 0 atau terisi sesuai dengan family_list_id yang ada di database maka akan melakukan update data... Tetapi jika is_delete itu = true maka akan menghapus data keluarga dari customer tersebut



