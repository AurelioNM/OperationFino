# Qual o nome dos usuários cadastrados na base de dados?
SELECT c.nm_cliente AS nomes FROM cliente c;

# Qual a quantidade de usuários cadastrados na tabela cliente?
SELECT COUNT(c.nm_cliente) FROM cliente c;


# Qual a maior venda realizada?
SELECT MAX(p.valor_pedido) FROM pedido p;

# Gerar um relatório com o valor total de compras feita por cada cliente ordenado pelos que mais gastaram.
explain analyse SELECT c.nm_cliente AS nome,
	SUM(p.valor_pedido) AS total_valor_pedido
FROM cliente c
JOIN pedido p ON c.id_cliente=p.id_cliente
GROUP BY nome
ORDER BY total_valor_pedido DESC;

# Usando a tabela de produtos, qual o produto mais caro, o mais barato, qual a méia de preço?
SELECT MAX(p.valor_pedido) AS mais_caro,
	MIN(p.valor_pedido) AS mais_barato,
	AVG(p.valor_pedido) AS media
FROM pedido p;

# Usando a tabela de produtos, quais são os top 10 mais caros ordernados do maior para o menor?
SELECT p.nm_produto AS nome,
	p.valor_produto AS valor
FROM produto p
ORDER BY valor DESC
LIMIT 10;

# Usando a tabela de produtos, quais são os top 10 mais baratos ordernados do menor para o maior?
SELECT p.nm_produto AS nome,
	p.valor_produto AS valor
FROM produto p
ORDER BY valor ASC
LIMIT 10;

# Usando a tabela de pedidos, apresente de forma acendente quais foram os pedidos e seus produtos
SELECT pe.id_pedido AS pedido_id,
	pe.valor_pedido AS pedido_valor,
	pp.id_produto AS produto_id,
	pr.nm_produto AS produto_nome,
	pr.valor_produto AS produto_valor
FROM pedido pe
JOIN pedido_produto pp ON pe.id_pedido=pp.id_pedido
JOIN produto pr ON pr.id_produto=pp.id_produto
ORDER BY pedido_id ASC;

# Usando a tabela de pedidos, apresente de forma acendente quais foram os pedidos e seus produtos e nome do comprador
SELECT pe.id_pedido AS pedido_id,
	pr.nm_produto AS produto_nome,
	c.nm_cliente AS cliente_nome
from pedido pe
JOIN pedido_produto pp ON pp.id_pedido = pe.id_pedido
JOIN produto pr ON pr.id_produto = pp.id_produto
JOIN cliente c ON c.id_cliente = pe.id_cliente
ORDER BY pedido_id ASC;

# Apresente apenas o primeiro produto com maior margem de ganho com os seguintes itens: nome do comprador, data da venda, qual o valor liquido faturado (custo - valor de venda), qual o percentual de lucro (valor de compra / valor de venda), id_produto, nm_produto

