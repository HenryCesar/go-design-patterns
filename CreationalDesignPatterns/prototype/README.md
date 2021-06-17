<details>
<summary> English </summary>

# Conceptual Example
Let’s try to figure out the Prototype pattern using an example based on the operating system’s file system. The OS file system is recursive: the folders contain files and folders, which may also include files and folders, and so on.

Each file and folder can be represented by an ``inode`` interface. `inode` interface also has the `clone` function.

Both `file` and `folder` structs implement the print and `clone` functions since they are of the `inode` type. Also, notice the `clone` function in both `file` and `folder`. The `clone` function in both of them returns a copy of the respective `file` or `folder`. During the cloning, we append the keyword “_clone” for the name field.

</details>

<details>
<summary> Português </summary>

# Exemplo conceitual

Vamos tentar descobrir o padrão Prototype usando um exemplo baseado no sistema de arquivos do sistema operacional. O sistema de arquivos do SO é recursivo: as pastas contêm arquivos e pastas, que também podem incluir arquivos e pastas e assim por diante.

Cada arquivo e pasta pode ser representado por uma interface `inode`. A interface `inode` também possui a função `clone`.

Ambas as structs `file` e `folder` implementam as funções print e `clone`, uma vez que são do tipo `inode`. Além disso, observe a função `clone` em `file` e `folder`. A função `clone` em ambos retorna uma cópia do respectivo arquivo ou pasta. Durante a clonagem, acrescentamos a palavra-chave “_clone” ao campo de nome.

</details>