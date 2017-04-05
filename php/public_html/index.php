<?php
require __DIR__ . '/../vendor/autoload.php';
include '../config.php';

$indexURI = PROTOCOL.'://'.HOST.':'.PORT.'/index.php';
$loginRedirectURI = $indexURI;
$loginRedirectURI = '/web/authorize';
$redirectURI = PROTOCOL.'://'.HOST.':'.PORT.'/callback.php';
$responseType = 'code';

$state = 'code';

$loginURL = AUTH_API_BASE_URI
.'/web/login'
.'?client_id='.CLIENT_KEY
.'&login_redirect_uri='.$loginRedirectURI
.'&redirect_uri='.$redirectURI
.'&response_type='.$responseType
.'&response_type='.$responseType
.'&state='.$state
;

?><!DOCTYPE html>
<html>
<head>
  <title>MyMiniFactory</title>
</head>

<?php

if(empty($_GET['access_token'])){
  ?>
  <a href="<?php echo $loginURL; ?>">Login With MyMiniFactory</a>
</html>
  <?php
} else {
?>
  <p>Confgratulation you are logged in to MyMiniFactory API !</p>
  <p>Access Token: <?php $_GET['access_token'] ?></p>
  <p><a href="<?php echo $indexURI ?>" >Log out</a></p>
  </html>
<?php
}
