<br>
 <div align="center" class="imgs">
 <img style="margin: 0 50px;" height="100em" src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/800px-Go_Logo_Blue.svg.png"> 

 <img height="100em" src="https://cdn.worldvectorlogo.com/logos/twilio.svg"> 
</div>
<br>

<br>
<hr>

<style>

a {
    color: red;
    transition: all .6s;
    text: none;
}

a:hover {
    color: #b61212;
    transition: all .6s;
    text-decoration: none;
}

.imgs {
    display: flex
}

</style>

## SMS

O serviço de sms é escrito em codigo Go, localizado na package
<b><i>services</i></b>

- Para rodar a aplicação certifiquice de ter o Go instalado em sua maquina, ou , execute um codigo Docker para poder compilar o codigo

- Acesse a pasta <b><i>services</i></b> com o comando no terminal

        cd services

- Apos acessar o diretorio do arquivo, rode o arquivo no seu terminal com o comando

        go run Sms.go

- Para se realizar o <b><i>Build</i></b> use o comando

        go build Sms.go

        
    <br>
<hr>


##  API


A Api escolhida foi Api <a href="https://www.twilio.com/"><b><i>Twilio</i></b></a> pois possui facil integração com mais diversos sistemas, suporte e uma documentação completa, alem de ser ultilisada por grandes empresas como
<br>
- Nubank
- Mercado Livre
- Dasa
- Inter
- Rappi

Não é um serviço qualquer sem confiança, por isso sua escolha
<hr>

## Variaveis de Ambiente

Para poder rodar a aplicação é de estrema importancia

- Criar um arquivo <b><i>.env</i></b>
(Windows)

        type nul > ".env"

- Criar um arquivo <b><i>.env</i></b>
(Linux)

        cat > .env

- Criar um arquivo <b><i>.env</i></b>
(Mac OS)

        touch .env

- Escrever as Variaveis da seguinte forma

        ACCOUNT_SID=SeuSidGeradoDaApi
        AUTH_TOKEN=SeuTockenGeradoDaApi
        TWILIO_TO=+NumRemetente (Codigo Nacional + Numero)
        TWILIO_FROM=+SeuNumeroGeradoDaApi

