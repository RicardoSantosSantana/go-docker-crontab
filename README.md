
## CONSEGUIR CODE

  1.Basta rodas a url abaixo com os dados e pronto
  
    https://auth.mercadolivre.com.br/authorization?response_type=code&client_id=5153648960977707&redirect_uri=http://localhost:3000&state=32gfgb.12345

  2.Você será redirecionado para a url cadastrada com o code

    http://localhost:3000/?code=TG-62f2874c32cb3a0001d4f2c6-78101319&state=32gfgb.12345

  ### **client_id**
    É o id do aplicativo que você cria na plataforma do mercado livre
    url: https://developers.mercadolivre.com.br/devcenter
  
  ### **redirect_url**
    É a URI de redirect * que cadastrou no aplicativo
  
  ###  **state**
    É um número aleatorio vc mesmo inventa

  ###  **client_secret**
    a Chave secreta fornecida quando cria ao aplicativo
  

## CONSEGUIR ACCESS_TOKEN E USER_ID
    #Rode no terminal com os dados que conseguiu com os passos acima

    curl -X POST \
    -H 'accept: application/json' \
    -H 'content-type: application/x-www-form-urlencoded' \
    'https://api.mercadolibre.com/oauth/token' \
    -d 'grant_type=authorization_code' \
    -d 'client_id=5153648960977707&' \
    -d 'client_secret=kcm1JZqgWeA66WLf08hxJ6m5vEqCRgMD' \
    -d 'code=TG-62f3c9f6f8e73a0001aeded6-78101319' \
    -d 'redirect_uri=http://localhost:3000'

## RESULTADO ##
  
  ### Esse comando te dará uma resposta como o exemplo abaixo, anote o ***access_token*** e o ***user_id***
    {
      "access_token":"APP_USR-5153648960977707-080912-31a9eb696840d95ccc6a5d7824428289-78101319","token_type":"bearer",
      "expires_in":21600,
      "scope":"read write",
      "user_id":78101319
    }   
## USUÁRIO DE TESTE ##
 
    curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' -H "Content-Type: application/json" -d
    '{
      "site_id":"MLA"
    }'
    https://api.mercadolibre.com/users/test_user

  #### RESULTADO ESPERADO ####
    {
      "id": 1176272578,
      "nickname": "TETE9221698",
      "password": "qatest9755",
      "site_status": "active",
      "email": "test_user_69300454@testuser.com"
    }
  OU 
  {
    "id": 1176951539,
    "nickname": "TESTRBQVQXWZ",
    "password": "qatest495",
    "site_status": "active",
    "email": "test_user_3027888@testuser.com"
}