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
<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.2.1/jquery.min.js"></script>
</head>
<body>
<?php

if(empty($_GET['access_token'])){
  ?>
  <a href="<?php echo $loginURL; ?>">Login With MyMiniFactory</a>
  <?php
} else {
?>
  <p>Confgratulation you are logged in to MyMiniFactory API !</p>
  <p>Access Token: <?php echo $_GET['access_token']; ?></p>
  <div class="user-info" style="display:none;">
    <p>Connected User Info:</p>
  </div>
  <p><a href="<?php echo $indexURI; ?>" >Log out</a></p>
  <script>
  var accessToken = "<?php echo $_GET['access_token']; ?>";
  $.ajax({
         url: "https://www.myminifactory.com/api/v2/user",
         headers: { 'Authorization': 'Bearer ' + accessToken },
         type: "GET",
         success: function(response) {
           $('.user-info').append('<img src="'+response.avatar_thumbnail_url+'" width=100 /><p><a href="'+response.profile_url+'" >username: ' + response.username + '</a></p>');
           $('.user-info').show();
         }
      });
  // $.ajax({
  //        url: "https://www.myminifactory.com/api/v2/user?access_token="+accessToken,
  //        type: "GET",
  //        success: function(response) {
  //          $('.user-info').append('<img src="'+response.avatar_thumbnail_url+'" width=100 /><p><a href="'+response.profile_url+'" >username: ' + response.username + '</a></p>');
  //          $('.user-info').show();
  //        }
  //     });
  </script>
<?php
}
?>
</body></html>
