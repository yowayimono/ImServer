package config

// any approach to require this configuration into your program.
var AppConfig = []byte(`
app:
  cookie_key: 4238uihfieh49r3453kjdfg
  debug_mod: "true"
  port: "8322"
  serve_type: GoServe
  upload_file_path: E:\
redis:
  host: localhost
  port: 6379

mysql:
  dsn: root:0503@tcp(106.52.78.230:3306)/qwm?charset=utf8mb4&parseTime=True&loc=Local
`)
