## Servicio Rest para crear Jwt con llaves RSA256

### Encrypt
```shell
curl --request POST \
  --url http://localhost:90/jwt/encrypt/rsa256 \
  --header 'Content-Type: application/json' \
  --data '{
	"jwt": {
		"nombre": "Juan",
		"edad": 30,
		"activo": true,
		"intereses": [
			"programación",
			"música",
			"deportes"
		],
		"dirección": {
			"calle": "123 Avenida Principal",
			"ciudad": "Ciudad Principal",
			"país": "País Principal",
			"cod_postal": "12345"
		}
	}
}'
```

### Decrypt
```shell
curl --request POST \
  --url http://localhost:90/jwt/decrypt/rsa256 \
  --header 'Content-Type: application/json' \
  --data '{
	"jwt": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY3Rpdm8iOnRydWUsImRpcmVjY2nDs24iOnsiY2FsbGUiOiIxMjMgQXZlbmlkYSBQcmluY2lwYWwiLCJjaXVkYWQiOiJDaXVkYWQgUHJpbmNpcGFsIiwiY29kX3Bvc3RhbCI6IjEyMzQ1IiwicGHDrXMiOiJQYcOtcyBQcmluY2lwYWwifSwiZWRhZCI6MzAsImV4cCI6MTcwODk5Mzg4OCwiaW50ZXJlc2VzIjpbInByb2dyYW1hY2nDs24iLCJtw7pzaWNhIiwiZGVwb3J0ZXMiXSwibm9tYnJlIjoiSnVhbiIsInN1YiI6Ikp3dENyZWF0ZSJ9.dMIMmdg2eO0Gc50Ou9LWOOfwI-z926C98eR2se8SEKeFSOuxJuln969m5gn9xvU0mD8ctICO6eNnqVjKK6QhRhSX5DyQ1mlm6eC-XRQVcK5hvsREL_oiDyBzKgrhIIjFRqy6I-CjFBmMVFBCV3ptOBTjZRhID-AOEU50k0B07ibdXwaiR5_2_YU0PKeOrFF--3QwrW6FK5yFYWvwUniZZuP8lAEqwXwC2ncSmDxs9VJiXsNWcFYWhh7POq-8r2QrP63wMQhuRjAETAoTqOHAsrocsrQNL9sszFdAn2M1YQSLJjpVx3XuKF4qCvs-UHIcq2WrLlfGlwIrI5C-swah-A"
}'
```