![FIAPCAR](https://github.com/user-attachments/assets/c5ce03c8-74d7-4296-a76a-10e2d108fa5f)

# HACKATON - FIAPCAR AUTOMÓVEIS

### SOBRE:

Backend para uma plataforma para revenda de veículos, implementada utilizando Golang, Postgres e Kubernetes. Toda a parte de autenticação foi desenvolvida de forma customizada e separada das funcionalidades do negócio, implementada utilizando JWT. Nesta fase, o serviço foi migrado para a nuvem Oracle Cloud Infrastructure (OCI).

### COMO TESTAR:

Importar o arquivo 'HACKATON_FIAPCAR.postman_collection.json' para o Postman.

Para utilizar as APIs, é necessário cadastrar um usuário na API '/usuario/cadastra-usuario' e em seguida gerar um token jwt utilizando a API '/usuario/cria-token?login=LOGINUSUARIO&senha=SENHAUSUARIO', preenchendo o login e senha com os dados do usuário criado anteriormente.

Para testar as APIs, é necessário informar o token gerado no campo 'Authorization'. 

Apenas usuários criados com a flag "admin = true" tem permissão para utilizar APIs de cadastro, atualização, deleção e consulta de veículos vendidos.

\
**CASO QUEIRA TESTAR EM AMBIENTE LOCAL, REALIZAR O DEPLOYMENT:**

Pré requisitos:\
Kubectl v1.29.0\
Minikube v1.32.0\
Go 1.21.3\
Docker 24.0.6

Executar o script deploy_local.sh
Ao executar o script, será deletado o cluster existente do Minikube e criado um novo, ajuste nas permissões do docker.sock, habilita o metrics-server, realiza o deploy da aplicação no cluster e configuração do HPA.

Importante aguardar alguns minutos para conclusão da criação dos pods antes de testar.

Comando para acompanhar o status da implantação:\
watch minikube kubectl get pods
