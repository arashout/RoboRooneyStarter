# RoboRooneyStarter
## What is RoboRooney
RoboRooney is a Slack-bot that helps you find available pitches (Powered by MyLocalPitch's API). 
The actual code for the Slack-bot is [here](https://github.com/arashout/roborooney) and the MLP API is [here](https://github.com/arashout/mlpapi).    
This repo is a template that lets you quickly start using RoboRooney.

## Credentials
Before you can start using the bot locally you need to set enviroment variables that will be used for connecting to Slack.   
`API_TOKEN` = This is your api token that you generate from Slack   
`BOT_ID` = In Slack this is called member ID, although this is not strictly necessary it enables the use of @roborooney which makes it super convenient.   

## Heroku    
You can easily run this bot on Heroku by specifying the above credentials in `.env` file.   
Remember to scale your dyno because since the command to run is not `web` but `worker` you need to manually scale the `worker` dyno.
