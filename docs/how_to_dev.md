
## Local Dev Setup

Current Go Version: [1.18](https://go.dev/dl/)

### Creating Discord Bot Playground
**Note:** This is not necessary if using the existing ARCH playground or if a personal playground has already been set up.
1. Create a new Discord server naming it whatever (e.g. "Bot Playground Server")
2. Go to the [Discord Developer Portal](https://discord.com/developers/applications).
3. Create a new application naming it whatever (e.g. "Bot Playground").
4. On the "General Information" tab, make note of "APPLICATION ID".
5. On the "Bot" tab, add a new bot naming it whatever (e.g. "The Speaker").
6. Click "Reset Token" and make note of the bot's token.
7. Toggle off the "PUBLIC BOT" setting.
8. Toggle on the "PRESENCE INTENT", "SERVER MEMBERS INTENT", and "MESSAGE CONTENT INTENT" settings.
9. Save Changes. 
10. Navigate to "OAth2" >> "URL Generator".
11. Select the following "SCOPES" (minimal permissions, adjust as necessary):
    - bot
    - applications.commands
12. Select the following "BOT PERMISSIONS" (minimal permissions, adjust as necessary):
    - Read Messages/View Channels
    - Send Messages
13. Navigate to the generated URL.
14. Select the desired server and "Continue".
15. Confirm the selected permissions and "Authorize".

Congrats! The bot has now been added to the server.

### Creating the `.env` File
This application expects a `.env` file at the base directory which contains the following variables:
* **DISCORD_BOT_TOKEN** - Token of the bot created at the Discord Developer Portal, prepended by "Bot ".
* **DISCORD_APP_ID** - "APPLICATION ID" from Discord Developer Portal.
* **DISCORD_GUILD_ID** - Discord server ID (right click on server >> Copy ID).

Example file:
```
DISCORD_BOT_TOKEN="Bot MTA0NTQwNjQ0NDU3MDI4NDE3Mw.G4Z2uG.OdAzvzhywprUorR1UvHV1uQIQt8_nja8jjqjVY"
DISCORD_APP_ID=1045406444570284173
DISCORD_GUILD_ID=1045407987159797860
```

## Run Command
`go run cmd/main.go`

---

Thanks [Ryan Anderson](https://github.com/randerson8907) for the write-up!