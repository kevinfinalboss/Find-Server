
# FindServer
Uma simples aplicação CLI que retorna algumas informações de um endereço IP.

## Funcionalidades
- IP -> Retorna o endereço de IP
- Servidores -> Retorna o NameServers do endereço
- Portas -> Retorna as portas públicas do endereço
- Cname -> Retorna os Cnames apontados para o endereço

## Aprendizados
Com este projeto pude aprender como utilizar os pacotes de terceiros de forma eficiente e também a tratar erros de uma forma simplificada em GO

## Documentação
[go-md2man](https://link-da-documentação)

[CLI](github.com/urfave/cli)

## Uso/Exemplos no terminal
```
go run .\main.go cname --host kevindev.com.br
Irá retornar todos os cnames desse endereço.

go run .\main.go porta --host kevindev.com.br
Irá retornar todas as portas públicas desse endereço.

```

