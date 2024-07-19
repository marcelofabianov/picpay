# Desafio Backend PicPay em Go

## Enumeradores

`WalletType` (Tipo de Carteira)
- **WalletTypeCommon:**: USER | Carteira comum do usuário.
- **WalletTypeMerchant**: MERCHANT | Carteira do estabelecimento.

`TransferStatus` (Status da Transferência)
- **TransferStatusPending:** PENDING | Transferência pendente.
- **TransferStatusRejected:** REJECTED | Transferência rejeitada.
- **TransferStatusReserved:** RESERVED | Transferência com valor reservado.
- **TransferStatusCompleted:** COMPLETED | Transferência concluída.
- **TransferStatusReversed:** REVERSED | Transferência revertida.
- **TransferStatusError:** ERROR | Transferência com erro.

`AuthorizationTransferStatus` (Status da Autorização de Transferência)
- **AuthorizationTransferStatusPending:** PENDING | Autorização de transferência pendente.
- **AuthorizationTransferStatusOK:** OK | Autorização de transferência concedida.
- **AuthorizationTransferRejected:** REJECTED | Autorização de transferência rejeitada.
- **AuthorizationTransferStatusError:** ERROR | Autorização de transferência com erro.

## Entidades

**User (Usuário)**

Entidade que representa um usuário.

- **ID**: `PK` `uuid` `required` `unique` | Identificador único do usuário.
- **Name**: `string` `required` | Nome do usuário.
- **Email**: `string` `required` `unique` | E-mail do usuário.
- **Password**: `string` `required` | Senha do usuário.
- **DocumentRegistry**: `string` `required` `unique` | CPF do usuário.
- **Enabled**: `bool` `required` | Indica se o usuário está habilitado ou não.
- **CreatedAt**: `datetime` `required` | Data de criação do usuário.
- **UpdatedAt**: `datetime` `required` | Data da última atualização do usuário.
- **DeletedAt**: `datetime` | Data da exclusão do usuário.
- **Version**: `int` `required` | Versão do registro.

**Wallet (Carteira)**

Entidade que representa a carteira de um usuário.

- **ID**: `PK` `uuid` `required` `unique` | Identificador único da carteira.
- **UserID**: `FK` `uuid` `required` | Identificador do usuário.
- **Amount**: `decimal` `required` | Valor do saldo da carteira.
- **Type**: `WalletType` `required` | Tipo da carteira.
- **Enabled**: `bool` `required` | Indica se a carteira está habilitada ou não.
- **CreatedAt**: `datetime` `required` | Data de criação da carteira.
- **UpdatedAt**: `datetime` `required` | Data da última atualização da carteira.
- **DeletedAt**: `datetime` | Data da exclusão da carteira.
- **Version**: `int` `required` | Versão do registro.

**Transfer (Transferência)**

Entidade que representa uma transferência de valores entre carteiras.

- **ID**: `PK` `uuid` `required` `unique` | Identificador único da transferência.
- **PayerID**: `FK` `uuid` `required` | Identificador do pagador.
- **PayerWalletID**: `FK` `uuid` `required` | Identificador da carteira do pagador.
- **PayeeID**: `FK` `required` | Identificador do recebedor.
- **PayeeWalletID**: `FK` `uuid` `required` | Identificador da carteira do recebedor.
- **Amount**: `decimal` `required` | Valor da transferência.
- **Status**: `TransferStatus` `required` | Status da transferência.
- **Enabled**: `bool` `required` | Indica se a transferência está habilitada ou não.
- **CreatedAt**: `datetime` `required` | Data de criação da transferência.
- **UpdatedAt**: `datetime` `required` | Data da última atualização da transferência.
- **Version**: `int` `required` | Versão do registro.

**AuthorizationTransfer (Autorização de Transferência)**

- **ID**: `PK` `uuid` `required` `unique` | Identificador único da autorização de transferência.
- **TransferID**: `FK` `uuid` `required` | Identificador da transferência.
- **Status**: `AuthorizationTransferStatus` `required` | Status da autorização de transferência.
- **Enabled**: `bool` `required` | Indica se a autorização de transferência está habilitada ou não.
- **CreatedAt**: `datetime` `required` | Data de criação da autorização de transferência.
- **UpdatedAt**: `datetime` `required` | Data da última atualização da autorização de transferência.
- **Version**: `int` `required` | Versão do registro.
