# GO
Learning GO (basic)
<img src="/assets/climb.png">

# Sobre GO

> Golang é uma linguagem de programação de código aberto, ou seja, qualquer pessoa pode inspecionar, modificar e aprimorar. ELa é uma linguagem *compilada* e *estaticamente tipada*
  * Compilada: é quando o código-fonte é executado diretamente pelo sistema operacional ou pelo processador. Isso ocorre após ele ser traduzido por um processo de compilação que utiliza um programa chamado compilador
  * Estaticamente tipada: define no código os tipos das variáveis de um programa. Assim, eles são conhecidos durante a compilação.

# Ainda sobre GO, mas olhando para o código - CRIANDO UMA API

> Primeiramente, uma API é utilizada para trocar dados entre diferentes tipos de software a fim de automatizar procedimentos e desenvolver novas funcionalidades.

Vamos utilizar o protocolo HTTP, então é importante saber que ele define um conjunto de métodos de requisição responsáveis por indicar a ação a ser executada para um dado recurso. Embora esses métodos possam ser descritos como substantivos, eles também são comumente referenciados como Verbos HTTP, como ilustram os códigos desenvolvidos nesse diretório

## pt.1 - JSON, Request, Response e GO

Import "fmt" (format) > é utilizado para formatar os dados de E/S, usado para debugs, formatação, interação entre usuário, entre outros.

Exportando arquivos > As funções que você deseja chamar de outros pacotes precisam começar com uma letra maiúscula.

Importando arquivos *sem* estar conectado ao github > para isso precisamos colocar o nome do projeto que foi inciado, por exemplo "API-GO-REST/nome-do-arquivo"

## pt.2 - Roteador, recursos por ID e docker

O pacote gorilla/mux implementa um roteador de requisições e respostas para corresponder às solicitações de entrada ao seu respectivo manipulador ou handler.

## pt.3 - Conexão com o banco e exibindo os dados

NewEncoder > Estamos recebendo um JSON e requeremos decodificando para que o ORM do GO saiba trabalhar
NewDecoder > Estamos pegando a informação, criando um ResponseWriter, criando uma informação para exibir encodando essa nova personalidade que estamos mostrando.

## pt.4 - Criando, editando e deletando com Gorm

Usamos o Post, Put e Delete


## pt.4 - Middleware e integração com front-end

Queremos setar um cabeçalho indicando que o "Content-type" que será devolvido para o nosso ResponseWriter é "application/json" e uma maneira de organizar uma determinada funcionalidade e compartilhá-la com outras partes do nosso código é através dos Middlewares. O middleware vai receber a solicitação do roteador e, antes de passar para o nosso controller, ele vai inserir algum fluxo ou funcionalidade.
Resumidamente, usar um middleware é muito útil para conseguirmos evitar duplicidade de código, para não precisarmos ficar copiando e colando a mesma linha para todas as funções que queremos. 

Aqui nessa etapa rolou um problema de CORS, no front-end, mas que precisa ser resolvido no back-end

* Cross-Origin Resource Sharing (CORS) -> Ele considera um compartilhamento que não pode acontecer, por estarmos acessando em portas diferentes. Isso acontece por questões de segurança, a gente não pode permitir que uma aplicação recebe requisições de qualquer porta ou domínio, tem que ser da mesma origem. 
  [ler mais sobre CORS](https://developer.mozilla.org/pt-BR/docs/Web/HTTP/CORS)

1. Para resolvermos isso precisamos ir em routes e fazer a configuração, fizemos isso usando a biblioteca gorilla.
2. Colocamos que "*" pois assumimos que qualquer um pode acessar e consumir os dados da API, mas poderíamos restringir para um localhost específico
3. Por último passamos nosso roteador (r)
