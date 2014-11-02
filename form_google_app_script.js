/**
 * Uses the Forms API to create a simple quiz.
 * For more information on using the Forms API, see
 * https://developers.google.com/apps-script/reference/forms
 */
function sendInvite(e) {
  // INPUT YOUR INFORMATION HERE
  var adminEmail = 'YOUR_EMAIL'
  var url = 'YOUR_APP_ENGINE_SERVER';

  var formResponse = e.response;
  var itemResponses = formResponse.getItemResponses();
  
  var fname = itemResponses[0].getResponse();
  var lname = itemResponses[1].getResponse();
  var email = itemResponses[2].getResponse();

  var payload = {
     "fname": fname,
     "lname": lname,
     "email": email
   };
  var headers = {
    // gets around App Engine bug with parsing url-encoded forms
    "Content-Type": "application/octet-stream"
  };
  var options = {
    'method': 'post',
    'headers': headers,
    'payload': payload
  };
  
  var response = UrlFetchApp.fetch(url, options);
  
  var responseCode = response.getResponseCode();
  
  if (responseCode != 200) {
    if (MailApp.getRemainingDailyQuota() > 0) {
      var message = "Response code was " + responseCode + ". URL was " + url + " and payload was " + JSON.stringify(payload) + ".";
      MailApp.sendEmail(adminEmail,
                        'Error with HTTP request to Gophers Slack',
                        message);
    }
  }
         
  var jsonResponse = Utilities.jsonParse(response.getContentText());
  
  if (jsonResponse.ok != true) {
    if (MailApp.getRemainingDailyQuota() > 0) {
      var message = "Response was: " + response.getContentText() + ". Payload was: " + JSON.stringify(payload) + ". URL was: " + url + ".";
      MailApp.sendEmail(adminEmail,
                        'Error with HTTP request to Gophers Slack.',
                        message);
    }
  }
}
