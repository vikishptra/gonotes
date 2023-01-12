
# GO-NOTES

GO-NOTES adalah sebuah service simple notes yang di dalam nya terdapat operasi create,read,update,delete yang sudah pasti terdapat error handling sesuai dengan case saya sendiri

- runCreateTodo 
- getAllTodo
- getTodoByID
- runTodoByChecked
- runTodoDeleteByID
- runUpdateTodoByID


## API GONOTES

#### Get all items with pagination

```http
  GET /api/v1/todo/?page=1&size=3
```
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `page` | `int` | **Required** halaman yang mau di request|
| `size` | `int` | **Required** hasil data yang mau di tampilkan dari halaman|


#### Get item By ID

```http
  GET /api/v1/todo/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. mengambil data dari paramater id |


#### PUT item checked(bool) By ID

```http
  PUT /api/v1/todo/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. mengambil data dari paramater id |




#### POST create item todo

```http
  POST /api/v1/todo/message
```

```
{ 
    "message":"agenda hari ini"
}
```

| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `message`      | `string` | **Required**. membuat message todo|



#### DELETE item todo By ID

```http
  DELETE /api/v1/todo/message/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. mengambil data dari paramater id|

```
notes : 
data notes checked harus bernilai true jika mau terhapus
```

#### PUT update item todo By ID

```http
  PUT /api/v1/todo/message/:id
```

```
{ 
    "message":"update agenda"
}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. mengambil data dari paramater id|


| Body | Type     | Description                       | 
| :-------- | :------- | :-------------------------------- | 
| `message`      | `string` | **Required**. membuat message todo|


