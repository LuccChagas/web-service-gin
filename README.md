# web-service-gin

API Simples que realiza um CRUD tendo como o Objeto *Albuns de Musica*. Os dados são armazenados em um banco de dados local.

### Dependencias ultilizadas

- Gin (https://gin-gonic.com/)
- Sqlite3 (https://www.sqlite.org/index.html)
- Gorm (https://gorm.io/index.html)


### Objeto principal

```golang
type Album struct {
	ID     uint    `json:"id" gorm:"primary_key"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
```
