<details>
<summary> English </summary>

# Conceptual Example

It’s impossible to implement the classic Factory Method pattern in Go due to lack of OOP features such as classes and inheritance. However, we can still implement the basic version of the pattern, the Simple Factory.

In this example, we’re going to build various types of weapons using a factory struct.

First, we create the `iGun` interface, which defines all methods a gun should have. There is a `gun` struct type that implements the iGun interface. Two concrete guns — `ak47` and `musket` — both embed gun struct and indirectly implement all `iGun` methods.

The `gunFactory` struct serves as a factory, which creates guns of the desired type based on an incoming argument. The main.go acts as a client. Instead of directly interacting with `ak47` or `musket`, it relies on `gunFactory` to create instances of various guns, only using string parameters to control the production.

</details>

<details>
<summary> Português </summary>

# Exemplo Conceitual

É impossível implementar o padrão Factory Method clássico no Go devido à falta de recursos OOP, como classes e herança. No entanto, ainda podemos implementar a versão básica do padrão, o Factory Simples.

Neste exemplo, vamos construir vários tipos de armas usando uma struct factory.

Primeiro, criamos a interface `iGun`, que define todos os métodos que uma arma deve ter. Existe um tipo de struct gun que implementa a interface `iGun`. Duas armas concretas — `ak47` e `musket` — ambas incorporam a struct da arma e indiretamente implementam todos os métodos `iGun`.

A struct gunFactory serve como um factory, que cria armas do tipo desejado com base em um argumento de entrada. O main.go atua como o cliente. Em vez de interagir diretamente com o `ak47` ou `musket`, ele conta com o gunFactory para criar instâncias de várias armas, usando apenas parâmetros de tipo string para controlar a produção.
</details>