<details>
<summary> English </summary>

# Intent
**Composite** is a structural design pattern that lets you compose objects into tree structures and then work with these structures as if they were individual objects.

# Conceptual Example
Let’s try to understand the Composite pattern with an example of an operating system’s file system. In the file system, there are two types of objects: files and folders. There are cases when files and folders should be treated to be the same way. This is where the Composite pattern comes in handy.

Imagine that you need to run a search for a particular keyword in your file system. This search operation applies to both files and folders. For a file, it will just look into the contents of the file; for a folder, it will go through all files of that folder to find that keyword.

</details>

<details>
<summary> Português </summary>

# Propósito
O **Composite** é um padrão de projeto estrutural que permite que você componha objetos em estruturas de árvores e então trabalhe com essas estruturas como se elas fossem objetos individuais.

# Exemplo conceitual

Vamos tentar entender o padrão Composite com um exemplo de sistema de arquivos de um sistema operacional. No sistema de arquivos, existem dois tipos de objetos: arquivos e pastas. Há casos em que arquivos e pastas devem ser tratados da mesma maneira. É aqui que o padrão Composite se torna útil.

Imagine que você precise fazer uma pesquisa por uma determinada palavra-chave em seu sistema de arquivos. Esta operação de pesquisa se aplica a arquivos e pastas. Para um arquivo, ele apenas examinará o conteúdo do arquivo; para uma pasta, ele percorrerá todos os arquivos dessa pasta para encontrar essa palavra-chave.

</details>