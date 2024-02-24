package repository

const clientExistsQuery = "SELECT EXISTS(SELECT 1 FROM clientes WHERE id = $1)"
const debitQuery = `
	UPDATE saldos
	SET valor = valor - $1
	FROM (SELECT limite FROM clientes WHERE id = $2) AS limite_cliente
	WHERE cliente_id = $3
		AND abs(saldos.valor - $4) <= limite_cliente.limite
	RETURNING limite, valor;
`
const insertTransactionQuery = "INSERT INTO transacoes(valor, tipo, descricao, cliente_id) values ($1, $2, $3, $4)"
