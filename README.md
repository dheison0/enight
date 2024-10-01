# Enight

Este é um projeto feito para uma hamburgueria local criado por apenas um desenvolvedor,
não espere que seja grande coisa e nem que ele consiga fazer de tudo.

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

  - Iniciar o projeto