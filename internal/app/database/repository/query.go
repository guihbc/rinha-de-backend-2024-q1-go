package repository

const clientExistsQuery = "SELECT EXISTS(SELECT 1 FROM clientes WHERE id = $1)"
const debitQuery = `
	UPDATE saldos
	SET valor = valor - $1
	FROM (SELECT limite FROM clientes WHERE id = $2) AS limite_cliente
	WHERE cliente_id = $3 
	AND saldos.valor - $4 >= -limite_cliente.limite
	RETURNING limite, valor;
`
const insertTransactionQuery = "INSERT INTO transacoes(valor, tipo, descricao, cliente_id) values ($1, $2, $3, $4)"
const creditQuery = `
	UPDATE saldos AS s
	SET valor = s.valor + $1
	FROM clientes AS c
	WHERE s.cliente_id = c.id
	AND c.id = $2
	AND s.cliente_id = $3
	RETURNING c.limite, s.valor;
`
