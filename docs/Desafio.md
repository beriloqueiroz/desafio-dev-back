# Desafio backend

Olá, Queremos conhecer melhor suas habilidades
técnicas e como você gerencia situações de incertezas.

## Introdução
Este exercício será dividido em duas etapas:
1. Desenvolvimento da Solução
Você deverá desenhar e desenvolver uma solução para o cenário dado. Inclua no README do seu repositório um diagrama mostrando a arquitetura proposta e uma breve
explicação de como chegou a ela. Utilize o Docker Compose para facilitar a execução dos diferentes componentes da sua solução. O código deve ser disponibilizado em um repositório do Github. Após finalizar o
desenvolvimento, enviar o link do repositório para podermos avaliar.

2. Apresentação da Solução
Você será convidado a nos demonstrar a sua solução. Queremos entender como você chegou a ela, então esteja preparado para responder muitas perguntas! Existem inúmeras soluções corretas para este desafio. Lembre-se, escalabilidade e resiliência são fundamentais!
Caso tenha alguma dúvida, esteja muito à vontade para nos perguntar!
Esperamos que divirta-se com este desafio

## Desafio
O envio de notificações se tornou uma feature bastante comum para muitas aplicações nos
últimos tempos. Além disso, em um mundo cada vez mais conectado é muito comum a
integração com APIs externas. Neste desafio, propomos que você desenhe e implemente
uma solução que permita o envio de notificações para os usuários sobre informações de
clima e tempo que serão lidas do CPTEC (Centro de Previsão de Tempo e Estudos Climáticos) ligada ao INPE.
- Requisitos funcionais:
  - A previsão climática e de ondas deverá ser lida usando a API do CPTEC. Mais
  informações consulte o link. Esses são os endpoints que iremos consumir:
    - Busca de Localidades http://servicos.cptec.inpe.br/XML/listaCidades?city={nome_cidade}
    - Retorna a previsão climática de uma cidade http://servicos.cptec.inpe.br/XML/cidade/{id_cidade}/previsao.xml
    - Retorna a previsão ondas para localidades litorâneas http://servicos.cptec.inpe.br/XML/cidade/{id_cidade}/dia/{dia_param}/ondas.xml
  - A notificação deve conter as temperaturas previstas para os próximos 4 dias
  - Se for uma localidade litorânea também deverá mostrar a previsão de ondas para o dia atual.
  - Inicialmente será suportado apenas o envio de notificações para uma aplicação web. 
  - Suporte a novos tipos de notificações, tais como Push, SMS e email, serão incluídos no futuro;
  - Deverá ser permitido o agendamento para o envio das notificações;
  - Usuários que solicitaram o opt-out não deverão mais receber notificações;
- Requisitos Não Funcionais:
  - As notificações devem ser entregues o mais rápido possível no agendamento previsto, porém pequenos delays são aceitáveis;
  - A solução deve ser escalável e resiliente;

Obs: Não é necessário construir o frontend para gestão de usuários. Podemos utilizar diretamente as APIs Rest publicadas pelas aplicações.