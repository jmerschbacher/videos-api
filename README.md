# videos-api
CRUD de vídeos destinado ao Alura Challenge #1 - Semana 1.

## Rotas

Obs.: Collection do Postman na pasta raíz do projeto

- GET /videos -> Listar todos os vídeos cadastrados -> status 200
- GET /videos/{id} -> Listar vídeo por ID -> status 200
- POST /videos (request body) -> Cadastrar vídeo -> status 201
- PATCH /videos/{id} (request body) -> Atualizar um ou mais dados de vídeos -> status 200
- DELETE /videos/{id} -> Excluir vídeo -> status 204

## Pendências
- Ajustar configurações do MySQL no Docker -> `docker-compose.yml`
- Configurar properties do Gorm para fazer a entity `video.go` para gerar a coluna ID autoincrementável (hoje há necessidade de abrir a tabela e fazer o ajuste manualmente)
- Criar testes unitários
- Criar handler de injeção de dependências
