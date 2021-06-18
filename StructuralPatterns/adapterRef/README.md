<details>
<summary> English </summary>

# Intent
**Adapter** is a structural design pattern that allows objects with incompatible interfaces to collaborate.

# Conceptual Example

We have a client code that expects some features of an object (Lightining port), but we have another object called adaptee (Windows laptop) wich offers the same functionality but through a different interface (USB port).

This is where the Adapter pattern comes into the picture. We create a struct type know as adapter that will:

- Adhere to the same interface which the client expercts (Lightning port).
- Translate the request from the client to the adaptee in the form that the adaptee expects. The adapter accepts a Lightning connector and then translates its signals into a USB format and passes them to the USB port in windows laptop.

</details>

<details>
<summary> Português </summary>

# Propósito
O **Adapter** é um padrão de projeto estrutural que permite objetos com interfaces incompatíveis colaborarem entre si.

# Exemplo Conceitual

Temos um código cliente que espera alguns recursos de um objeto (Lightning port), mas temos outro objeto chamado *adaptee* (laptop Windows) que oferece a mesma funcionalidade, mas por meio de uma interface diferente (porta USB).

É aqui que o padrão Adapter entra em cena. Criamos um tipo de struct conhecido como adapter que irá:

- Seguir a mesma interface que o cliente espera (Lightning port).
- Traduzir a solicitação do cliente par ao adapter na forma que o adpter espera. O adapter aceita um conector Lightning e, em seguida, converte seus sinais em um formato USB e os passa para a porta USB no laptop com Windows.

</details>