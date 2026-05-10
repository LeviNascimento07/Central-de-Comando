# SGI Back-end

API REST em Go (Gin) para o Sistema de Gestão de Incidentes.

## Configuração

1. Copie o arquivo `.env` e preencha a senha:
```
DB_PASSWORD=sua_senha_aqui
PORT=8080
```

2. Instale as dependências:
```bash
go mod tidy
```

3. Execute o servidor:
```bash
go run main.go
```

O servidor sobe em `http://localhost:8080`.

---

## Endpoints

### Usuarios — `/usuarios`

| Método | Rota             | Descrição               | Body                                      |
|--------|------------------|-------------------------|-------------------------------------------|
| GET    | /usuarios        | Lista todos os usuários | —                                         |
| GET    | /usuarios/:id    | Busca usuário por ID    | —                                         |
| POST   | /usuarios        | Cria novo usuário       | `{ nome, login, senha, atualizado_por }` |
| PUT    | /usuarios/:id    | Atualiza usuário        | `{ nome, login, senha, atualizado_por }` |
| DELETE | /usuarios/:id    | Remove usuário          | —                                         |

**Exemplo POST /usuarios:**
```json
{
  "nome": "João Silva",
  "login": "joao.silva",
  "senha": "senha123",
  "atualizado_por": "admin"
}
```

---

### Equipamentos — `/equipamentos`

| Método | Rota                | Descrição                  | Body                            |
|--------|---------------------|----------------------------|---------------------------------|
| GET    | /equipamentos       | Lista todos os equipamentos| —                               |
| GET    | /equipamentos/:id   | Busca equipamento por ID   | —                               |
| POST   | /equipamentos       | Cria novo equipamento      | `{ descricao, valor_diaria }`  |
| PUT    | /equipamentos/:id   | Atualiza equipamento       | `{ descricao, valor_diaria }`  |
| DELETE | /equipamentos/:id   | Remove equipamento         | —                               |

**Exemplo POST /equipamentos:**
```json
{
  "descricao": "Notebook Dell XPS",
  "valor_diaria": 150.00
}
```

---

### Incidentes — `/incidentes`

| Método | Rota              | Descrição                | Body                                                            |
|--------|-------------------|--------------------------|-----------------------------------------------------------------|
| GET    | /incidentes       | Lista todos os incidentes| —                                                               |
| GET    | /incidentes/:id   | Busca incidente por ID   | —                                                               |
| POST   | /incidentes       | Cria novo incidente      | `{ data, descricao, equipamento_id, pessoa_id, status_id }`   |
| PUT    | /incidentes/:id   | Atualiza incidente       | `{ data, descricao, equipamento_id, pessoa_id, status_id }`   |
| DELETE | /incidentes/:id   | Remove incidente         | —                                                               |

**Exemplo POST /incidentes:**
```json
{
  "data": "2026-05-09T10:00:00Z",
  "descricao": "Equipamento com defeito na tela",
  "equipamento_id": 1,
  "pessoa_id": 2,
  "status_id": 1
}
```

> Se `data` for omitido no POST, o campo é preenchido automaticamente com a data/hora atual.

---

## Banco de Dados

| Campo         | Valor              |
|---------------|--------------------|
| Host          | m23-0t.h.filess.io |
| Port          | 3307               |
| User          | sgi_marketdead     |
| Database      | sgi_marketdead     |
| Password      | variável DB_PASSWORD |

## Tabelas esperadas

```sql
CREATE TABLE tbUsuarios (
  usuario_id     INT AUTO_INCREMENT PRIMARY KEY,
  nome           VARCHAR(100),
  login          VARCHAR(50),
  senha          VARCHAR(255),
  atualizado_em  DATETIME,
  atualizado_por VARCHAR(50)
);

CREATE TABLE tbEquipamento (
  equipamento_id INT AUTO_INCREMENT PRIMARY KEY,
  descricao      VARCHAR(200),
  valor_diaria   DECIMAL(10,2)
);

CREATE TABLE tbIncidente (
  incidente_id   INT AUTO_INCREMENT PRIMARY KEY,
  data           DATETIME,
  descricao      TEXT,
  equipamento_id INT,
  pessoa_id      INT,
  status_id      INT
);
```
