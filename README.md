# Desafio Backend PicPay em Go

https://github.com/PicPay/picpay-desafio-backend

## Descrição

O desafio consiste em criar uma API REST para simular transações financeiras entre usuários.

## Regras de Negócio

- Dever conter 2 tipos de wallet (Usuário e Lojista).
- Ambos usuários precisam ter nome completo, CPF/CNPJ, e-mail e senha.
- Usuário PF podem realizar transferências para Pessoa Física e Lojistas.
- O CPF/CNPJ, devem ser válidos e únicos.
- Usuário com wallet do tipo Lojista só pode receber transferências, não pode transferir para outros usuários.
- Validar se o usuário tem saldo antes da transferência.
- Antes de finalizar a transferência, deve-se consultar um serviço autorizador externo. (crie um mock)
- A operação de transferência deverá ser uma transação com opção de rollback em caso de falha.
- No recebimento de pagamento o usuário precisa receber uma notificação enviada por um serviço de terceiros que eventualmente poderá estar indisponível.

## Requisitos não funcionais

- A API Rest deverá seguir os princípios RESTFul.
- Deverá conter autenticação por token JWT.
- Criar fila para processar envio para o serviço de notificação.
- Criar fila para validar autorização de transferência com serviço externo.
- Definir timeout de 10s para autorização de transferência com 3 tentativas de reenvio.
- Definir timeout de 5s para processamento de fila de notificação com 3 tentativas de reenvio.

## Fluxo de Transferência

1. Usuário faz o pedido de transferência.
2. Valida se o usuário tem autorização para a operação.
3. Valida se o usuário tem saldo suficiente.
4. Reserva o valor da transferência.
5. Consulta serviço de autorização externo. (3 tentativas com timeout de 10s)
6. Realiza a transferência.
7. Envia notificação para o usuário de destino. (3 tentativas com timeout de 5s)

## Arquiteturas

A aplicação foi desenvolvida usando conceitos de arquitetura Ports and Adapters, EDA e Event Sourcing.

## Documentação

- [Entities](_docs/entities.md)

## Endpoints

...

## Tecnologias

...

## Requirements

...
