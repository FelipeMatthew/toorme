# TOORME API - NESTJS

APIs para controle e gerenciamento da base de dados e todas as interações para ser consumida no app_toorme

## Ambientes de desenvolvimento 

URL: 
Teste:   

#### Variáveis de ambiente

	DB_HOST=""
	DB_PORT=
	DB_NAME=""
	DB_USER=""
	DB_PASS=""


<!-- REMOVE HERE -->

  TypeOrmModule.forRoot({
      type: 'postgres',
      host: 'localhost',
      port: 5432,
      username: 'postgres',
      password: 'postgres',
      database: 'toorme',
      entities: [__dirname + "/../**/*.entity.{js,ts}"],
      synchronize: true, // todo: remove later
    })
    ,