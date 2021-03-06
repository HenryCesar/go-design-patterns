<details>
<summary> English </summary>

# Intent

**Flyweight** is a structural design pattern that lets you fir more objects into the available amount of RAM by sharing common parts of state between multiple objects instead of keeping all of the data in each object.

<img src=https://refactoring.guru/images/patterns/content/flyweight/flyweight.png width="300" height="200"/>

# Conceptual Example
In a game of Counter-Strike, Terrorist and Counter-Terroist have a different type of dress. For simplicity, let's assume that both Terrorist and Counter-Terrorists have onde dress type each. The dress object is embedded in the player object as below.

Below is the struct for a player. We can see that the dress object is embedded in the player struct:

```go
type player struct
{   
   dress dress
   playerType string //Can be TR or CT
   lat int
   long int
}
```

Let's say there are 5 Terrorist and 5 Counter-Terrorist, so a total of 10 players. Now there are two options concerning dress.

    1. each of the 10 player objects creates a different dress object and embeds them. A total of 10 dress objects will be created.

    2 We create two dress objects: 
    
    - Single Terrorist Dress Object: This will be shared across 5 Terrorists.
    - Single Counter-Terrorist Dress Object: This will be shared across 5 Counter-Terrorist.

As you can see in Approach 1, a total of 10 dress objects are created while in approach 2 only 2 dress objects are created. The second approach is what we follow in the Flyweight design pattern. the two dress objects which we created are called flyweight objects.

The Flyweight pattern takes out the common parts and creates flyweight objects. These flyweight objects (dress) can be shared among multiple objects (player). This drastically reduces the number of dress objects, and the good part is that even if you create more players, only two dress will be sufficient.

In the flyweight pattern, we store the flyweight objects in the map field. Whenever the other objets which share the flyweight objects are created, then flyweight objects are fetched from the map.

Let's see what parts of this arrangement will be intrinsic and extrinsic states:

- *Intrinsic State*: Dress in the intrisic state as it can be shared across multiple Terrorist and Counter-Terrorist objects.
- *Extrinsic State*: Player location and the player weapon are an extrinsic state as they are different for every object.


</details>

<details>
<summary> Portugu??s </summary>

# Prop??sito

O **Flyweight** ?? um padr??o de projeto estrutural que permite a voc?? colocar mais objetos na quantidade de RAM dispon??vel ao compartilhar partes comuns de estado entre os m??ltiplos objetos ao inv??s de manter todos os dados em cada objeto.

<img src=https://refactoring.guru/images/patterns/content/flyweight/flyweight.png width="300" height="200"/>

# Exemplo Conceitual

Em um jogo de Counter-Strike, os Terroristas e Antiterroristas t??m um tipo de uniforme diferente. Para simplificar, vamos supor que tanto os Terroristas quanto os Antiterroristas tenham um tipo de uniforme cada. O objeto uniforme (dress) ?? embutido no objeto jogador (palyer) como abaixo.

Abaixo est?? a struct de um player. Podemos ver que o objeto dress est?? incorporado na struct do player:

```go
type player struct
{   
   dress dress
   playerType string //Pode ser TR ou CT
   lat int
   long int
}
```

Digamos que haja 5 Terroristas e 5 Antiterroristas, ent??o um total de 10 players.
Agora, existem duas op????es de dress.

    1. Cada um dos 10 objetos player cria um objeto dress diferente e os incorpora. Um total de 10 objetos dress ser??o criados.

    2. Criamos dois objetos dress:
        - Objeto Dress Terrorista ??nico: Isso ser?? compartilhado por 5 terroristas.
        - Objeto Dress Antiterrorista ??nico: Isso ser?? compartilhado por 5 antiterroristas.

Como voc?? pode ver na abordagem 1, um total de 10 objetos dress s??o criados, enquanto na abordagem 2 apenas 2 objetos dress s??o criados. A segunda abordagem ?? a que seguimos no padr??o de design Flyweight. Os dois objetos dress que criamos s??o chamados de objetos flyweight.

O padr??o Flyweight remove as partes comuns e cria objetos flyweight. Esses objetos flyweight (dress) podem ent??o ser compartilhados entre v??rios objetos (players). Isso reduz drasticamente o n??mero de dresses, e a parte boa ?? que mesmo se voc?? criar mais players, apenas dois dresses ser??o suficientes.

No padr??o flyweight, armazenamos os objetos flyweight no campo do mapa.
Sempre que os outros objetos que compartilham os objetos flyweight s??o criados, ent??o os objetos flyweight s??o buscados no mapa.

Vamos ver quais partes desse arranjo ser??o intr??secos e extr??nsecos:

- *Estado Intr??seco*: O uniforme est?? no estado intr??nseco, pois pode ser compartilhado entre v??rios objetos Terroristas e Antiterroristas.
- *Estado Extr??nseco*: A localiza????o do jogador e a arma do jogador s??o m estado extr??nseco, pois s??o diferentes para cada objeto.

</details>
