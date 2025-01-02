# Enight

Este é um projeto feito para uma hamburgueria local criado por apenas um
desenvolvedor, não espere que seja grande coisa e nem que ele consiga fazer
de tudo.

## Como o sistema deve funcionar

Todas as informações serão guardadas em um banco de dados SQLite comum - não
vejo necessidade de um usar um banco de dados gigantes em um projeto pequeno
desse - e a API será desenvolvida usando a linguagem de programação Go por sua
simplicidade e eficacial no uso de recursos, podendo colocar o sistema todo
para rodar em um container de 128MB de RAM, a interface será desenvolvida
usando o framework Vite.JS + ReactJS e para criar o layout e estilo das páginas
o TailwindCSS vai entrar nessa brincadeira.

O cliente deve ser capaz de conversar com um bot no whatsapp onde ele vai se
registrar definindo onde mora, apartir disso ele podera solicitar um novo
pedido onde será gerado uma URL para acessar a página de escolha dos produtos,
nessa página vai ter de tudo incluindo uma barra de pesquisa, quando um produto
for selecionado sera aberta uma página onde possui as imagens do produto, sua
descrição e tamanhos disponiveis para venda, o cliente poderá selecionar o
tamanho e quantidade de cada produto, ao finalizar ele verá uma lista com tudo
incluido e o preço total já com o frete, se aceitar o pedido sera enviado para
o dono da hamburgueria através do sistema em uma página web dedicada a isso, o
dono do sistema podera definir se vai aceitar o pedido ou recusar, caso recuse
ele podera enviar uma mensagem explicando o motivo.

## TODO

  - Criar sistema de configurações;
  - Criar todo o sistema da web e do bot no whatsapp.


## Como rodar o sistema?

O sistema pode ser iniciado usando o script `run.sh` ou então você pode
compilar o servidor que está na pasta `server/` usando o Go 1.23 e a interface
web que está na pasta `web/` usando o NodeJS 20.18, depois disso é preciso
definir as váriaveis de ambiente citadas abaixo e rodar as migrações dos modelos
de dados que estão na pasta `server/database/migrations`, isso pode ser feito
usando a interface CLI do SQLite 3 usando o comando `sqlite3`.

As seguintes váriaveis de ambiente são carregadas na inicialização de todo
sistema, elas podem ser definidas em um arquivo `.env` que sera carregado:

  - `SERVER_PORT` A porta onde o servidor vai ficar escutando;
  - `WEB_FILES_PATH` Localização onde os arquivos das páginas web estão;
  - `DB_PATH` Localização do arquivo de banco de dados SQLite3;
  - `BOT_DB_PATH` A localização do banco de dados onde vai ficar armazenado as
    informações de login go bot;
  - `DEBUG` Define se os sistema está no modo de depuração(ativo por padrão, 
    "false" desativa);
  - `JWT_TOKEN` Pode ser usado durante o desenvolvimento para evitar a
    nessecidade de refazer o login toda vez que o sistema é reiniciado já que
    se não definido é usado um token aleatório;
  - `VITE_WHATSAPP_PHONE` Número de telefone para contato no WhatsApp(normalmente o do bot);
  - `VITE_API_BASE` URL base da API, incluindo protocolo e página.
