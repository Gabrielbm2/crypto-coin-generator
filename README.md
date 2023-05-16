# API de Votação de Criptomoedas

Este projeto apresenta uma API para votação de criptomoedas, que permite aos usuários votar em suas criptomoedas favoritas, enviar novas criptomoedas para serem votadas e obter informações sobre as criptomoedas e os votos realizados.

### Funcionalidades

* Criação de votos para as criptomoedas existentes

* Adição de novas criptomoedas para serem votadas

* Exibição das criptomoedas e votos atuais

* Documentação Swagger para fácil acesso aos endpoints

### Arquitetura

A arquitetura do projeto segue uma abordagem modular, onde cada pacote é responsável por uma funcionalidade específica, o que ajuda a manter o código organizado e facilita a manutenção.

### Documentação

A documentação da API pode ser encontrada através do Swagger. Para visualizá-la, acesse a URL (http://ec2-44-214-99-81.compute-1.amazonaws.com/static/docs/.)

### Testes

A API permite criar uma nova criptomoeda especificando seu nome e ID, por exemplo: name: Klever/ID:KLV, além de ver todas as criptomoedas criadas, pesquisar por ID e dar likes e dislikes em alguma criptomoeda que você escolher pelo ID.

### Deploy

A API foi implantada na Amazon Web Services (AWS) usando o serviço Elastic Compute Cloud (EC2).

### Melhorias

* Algumas melhorias possíveis para o projeto são:

* Adição de migrations

* Pool de conexões ao banco de dados

* Adição de um front-end para a interação do usuário com a API

* Inclusão de validações de entrada para evitar erros de usuário

* Melhorias na documentação, incluindo exemplos de requisições e respostas esperadas

### Considerações Finais

Este projeto foi desenvolvido com a intenção de demonstrar as habilidades de programação em Go. Qualquer dúvida ou sugestão pode ser encaminhada por email para gabribmeireles@hotmail.com.
