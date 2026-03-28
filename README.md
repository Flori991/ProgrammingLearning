# Programming Learning
A repo for all my learning projects for all kinds of languages.

## Go
This currently includes one project written in Go, called airvpn-api.
It simply takes an API key, obtained through the [UI on their website](https://airvpn.org/apisettings/), as a request header with the key `API-KEY` and then it queries your current sessions, correlates it with server info and gives relevant information back in json format.
The default docker compose file looks like this:
```yaml
services:
  airvpnapi:
    image: ghcr.io/flori991/airvpn-api:latest
    container_name: airvpn-api
    restart: unless-stopped
```

### Configuration
Everything ticked is already available, unticked is still planned.
- [x] `LOG_LEVEL`= `error | warning | info | debug`
- [ ] `PORT` = `3000` (currently only default port 3000)

### Usage
This repo contains a [glance `custom-api`](https://github.com/glanceapp/glance/blob/main/docs/configuration.md#custom-api) widget called [`airvpn-dash`](/Go/air-vpn-api/airvpn-dash.yml) that uses this API to request and show your AIRVPN sessions, feel free to use it and tinker with it. This is not yet fully tested and I only made it for my own consumption, so use at your own risk.
