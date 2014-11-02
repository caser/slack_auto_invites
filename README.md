Slack Auto Invites Setup
=========

The Slack API doesn't yet support programmatically inviting members to a community. With public groups, sending everyone manual invitations is a huge pain. This library solves that problem using the following technologies:

  - Google Forms
  - Golang
  - Google App Engine

First, fork and/or clone this repository locally and install [Go](https://golang.org/doc/install) and the [Google App Engine SDK](http://blog.joshsoftware.com/2014/03/12/learn-to-build-and-deploy-simple-go-web-apps-part-one/) and change ```sample_conf.json``` to ```conf.json```.

Set up a site on Google App Engine and put your appspot url in the appropriate place in ```conf.json```. 

Next, find your API token by going to your invite page (i.e. https://yourcommunity.slack.com/admin/invites) and finding the ```api_token``` variable in the source code, then use it to replace the dummy data in ```conf.json```.

Then, change ```appliation: slack-auto-invites``` to whatever your appspot site prefix (i.e., the sample is for slack-auto-invites.appspot.com) and deploy to Google App Engine using ```goapp deploy```. 

Then, create a Google Form which has the following three text fields:

- First name
- Last name
- Email

Once you've created the form, open up the script editor (under the Tools menu) and paste the code from  ```form_google_app_script.js``` into the development environment, changing the URL to your appspot.com url and the admin email to wherever you'd like to receive error logs. 

Then, open up the Resources menu and click on "Current Project's Triggers." Add a new "On form submit" trigger which calls the sendInvite method.

At this point, when your form is submitted, it should send a query to the Go server, which in turn sends an HTTP request to Slack inviting the new member to the community. 

Currently, you have to manually update your app when your token expires. You will receive an email to your adminEmail when requests to Slack start failing due to auth issues.

###TODO:
- Figure out how long it takes for the auth token to expire
- Figure out way to programmatically update auth token
