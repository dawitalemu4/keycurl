# Docker Setup

All you need to do run this project with docker is to:

1. Download this folder using the keycurl.zip in this folder and extract it in your preferred location

2. Rename `.env.example` to `.env` and put your own custom values (you can just use the provided values)

3. Run `docker-compose up` in your keycurl directory while your docker engine is running

4. Go to `localhost:YOURPORT` in your browser

Check out [keycurl.github.io/shortcuts](https://keycurl.github.io/shortcuts) or [shortcuts.md](https://github.com/dawitalemu4/keycurl.github.io/blob/main/src/assets/docs/shortcuts.md) for tutorials on how to run the [startup script](https://github.com/dawitalemu4/keycurl/tree/main/.docker-setup/startup.sh) in this folder from a shortcut on your taskbar to easily start up keycurl!

> Note for mac/linux users: If you are running into platform related errors, try to use this flag in the docker file: `FROM --platform=linux/amd64 dawitalemu4/keycurl:latest`
