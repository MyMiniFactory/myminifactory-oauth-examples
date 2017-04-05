<?php
use GuzzleHttp\Client;
require __DIR__ . '/../vendor/autoload.php';
include '../config.php';

$authorizationCode = $_GET['code'];
$redirectURI = PROTOCOL.'://'.HOST.':'.PORT.'/callback.php';
$indexURI = PROTOCOL.'://'.HOST.':'.PORT.'/index.php';

function getTokens($authorizationCode, $redirectURI){
  $client = new Client([
      'base_uri' => AUTH_API_BASE_URI,
  ]);

  $response = $client->request(
    'POST',
    AUTH_API_VERSION.'/oauth/tokens',
    [
      'auth' => [CLIENT_KEY, CLIENT_SECRET],
      'form_params' => [
        'grant_type' => 'authorization_code',
        'redirect_uri' => $redirectURI,
        'code'   => $authorizationCode,
      ]
    ]
  );

  return $response->getBody();;
}

// Request Access Token
$tokens = getTokens($authorizationCode, $redirectURI);
// Send Access Token to client
$accessToken = json_decode($tokens, true)['access_token'];

header("Location: $indexURI?access_token=".$accessToken);
