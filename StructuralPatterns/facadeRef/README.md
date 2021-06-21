<details>
<summary> English </summary>

# Intent

**Facade** is a structural pattern that provides a simplified interface to a library, a framework, or any other complex set of classes.

<img src=https://refactoring.guru/images/patterns/content/facade/facade.png width="300" height="200"/>

# Conceptual Example

It's easy to underestimate the complexities that happen behind the scenes when you order a pizza using your credit card. There are dozens of subsystems that are acting in this process. Here's just a shorlist of them:

- Check account
- Check security PIN
- Credit/debit balance
- Make ledger entry
- Send notification

In a complex system like this, it's easy to get lost and easy to break sutff if you're doing something wrong. That's why there's a concept of tha Face pattern: a thing that lets the client work with dozens of components using a simples interface. The client only needs to enter the card details, the security pin, the amount to pay, and the operation type. The Facade directs further communications with various components without exposing the client to internal complexities.

</details>

<details>
<summary> Português </summary>

# Propósito

O **Facade** (Também conhecido como: *Fachada*) é um padrão de projeto estrutural que fornece uma interface simplificada para uma biblioteca, um framework, ou qualquer conjunto complexo de classes.

<img src=https://refactoring.guru/images/patterns/content/facade/facade.png width="300" height="200"/>

# Exemplo Conceitual

É fácil subestimar as complexidades que acontecem nos bastidores quando você pede uma pizza com cartão de crédito. Existem dezenas de subsistemas que atuam nesse processo. Aqui está apenas uma lista deles:

- Checar conta
- Verificar o PIN de segurança
- Saldo do crédito/débito
- Fazer uma entrada no livro-caixa
- Enviar notificação

Em um sistema complexo como este, é fácil se perder e quebrar coisas se você estiver fazendo algo errado. É por isso que existe o conceito do padrão Facade: uma coisa que permite o cliente trabalhar com dezenas de componentes usando uma interface simples. O cliente só precisa inserir os dados do cartão, o PIN de segurança, o valor a pagar, e o tipo de operação. O Facade direciona comunicações adicionais com vários componentes sem expor o cliente a complexidades internas.

</details>
