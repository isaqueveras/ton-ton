# tón-tón
tón-tón is a simple headless cms in go

![image](https://github.com/isaqueveras/ton-ton/assets/46972789/0c3986e6-3705-40ac-a5b2-bbca1496d926)

---

`Get Article >> /v1/article/:id [GET]`
```json
{
  "title": "Historia da cidade de Mombaça",
  "content": "Mombaça é um município brasileiro do Estado do Ceará [...]",
  "status": "Draft",
  "slug": "historia-da-cidade-de-mombaca",
  "creator_id": "a5bc8e67",
  "created_at": "2023-10-04T20:59:49.728882-03:00"
}
```

`Add Article >> /v1/article [POST]`
```json
{
  "title": "Historia da cidade de Mombaça",
  "content": "Mombaça é um município brasileiro do Estado do Ceará [...]",
  "status": "Draft",
  "creator_id": "a5bc8e67"
}
```

`Edit Article >> /v1/article/:id [PUT]`
```json
{
  "content": "Mombaça é um município brasileiro [...]",
  "status": "Publish"
}
```

`Delete Article >> /v1/article/:id [DELETE]`
