<details>
<summary> English </summary>

# Intent
**Bridge** is a structural design pattern that lets you split a large class or a set of closely related classes into two separate hierarchies -abstraction and implementation- which can be developed independently of each other.

# Conceptual Example
Say, you have two types of computers: Mac and Windows. Also, two types of printers:
Epson and HP. Both computers and printers need to work with each other in any combination. The client doesn't want to worry about the details of connecting printers to computers.

If we introduce new printers, we don't want our code to grow exponentially. Instead pf creating four structs for the 2*2 combination, we create two hierarchies:

- Abstraction hierarchy: this will be our computers
- Implementation hierarchy: this will be our printers

These two hierarchies communicate with each other via a Bridge, where the Abstraction (computer) contains a reference to the Implementation (printer). Both the abstraction and implementation can be developed independently without affecting each other.

</details>

<details>
<summary> Português </summary>

# Propósito
O **Bridge** é um padrão de projeto estrutural que permite que você divida uma classe grande ou um conjunto de classes intimamente ligadas em duas hierarquias separadas -abstração e implementação-  que podem ser desenvolvidas independentemente umas das outras.

# Exemplo conceitual
Digamos que você tenha dois tipos de computador: Mac e Windows. Além disso, dois tipos de impressoras: Epson e HP. Tanto os computadores quanto as impressoras precisam funcionar entre si em qualquer combinação. O cliente não quer se preocupar com os detalhes da conexão de impressoras a computadores.

Se introduzirmos novas impressoras, não queremos que nosso código cresça exponencialmente. Em vez de criar quatro structs para a combinação 2 * 2, criamos duas hierarquias:

- Hierarquia de abstração: serão nossos computadores
- Hierarquia de implementação: serão as nossas impressoras

Essas duas hierarquias se comunicam por meio de uma Bridge, onde a Abstração (computador) contém uma referência à Implementação (impressora). Tanto a abstração quanto a implementação podem ser desenvolvidas independentemente sem afetar uma à outra.

</details>