# Usage

`boringproxy` consists of two kinds of long running processes that act as server and client.

They can be invoked directly on the shell or in a tmux session, or be configured to run as system daemons.

It is managed via an administrative web interface.

## Command line interface

The command line interface is used to create server and client instances.

## Server

The server requires you to pass the domain it will serve. It is invoked with:

```bash
./boringproxy server -admin-domain <FQDN>
```

* `FQDN` - the main domain configured in DNS

> Example:
> 
> `boringproxy server -admin-domain bpdemo.brng.pro`

Upon first run a configuration file `boringproxy_db.json` with admin credentials will be created and a login link is displayed.

* `https://<FQDN>/login?access_token=<TOKEN>`

> Example:
> 
> `https://bpdemo.brng.pro/login?access_token=yJaicLl6zj48zZItXvtQGb4CH5m5fId5`

## Client

The client needs to know which server to connect to, a valid combination of user and access token plus a unique name for the current client. You can start it with:

```bash
./boringproxy client -server <FQDN> -user <USER> -token <TOKEN> -client-name <CLIENT>
```

* `FQDN` - the domain of the remote server
* `USER` - username, min. six characters, if not admin
* `TOKEN` - access token to authenticate against the service
* `CLIENT` - identifier for the client, used in the interface to distinguish the available tunnel origins

> Example:
> 
> `boringproxy client -server bpdemo.brng.pro -token fKFIjefKDFLEFijKDFJKELJF -client-name demo-client -user demo-user`

## Web interface

The web interface is used to manage users, access tokens and tunnels.

Access it at `https://<FQDN>/` where you will be presented with a prompt for a token, or with the direct login link from above.

> ![](./../screenshot.png)
>
> Web interface after logging in as admin.

### Tunnels

The tunnels pane allows to create and remove tunnels. It offers you to specify the settings of a tunnel:

- **Domain** - 
- **Client Name** - 
- **Client Address** - 
- **Client Port** - 
- **Allow External TCP** - 
- **Password Protect** - 

### Tokens

The tokens pane allows to create and remove tokens for your current user.  
Multiple tokens per user can be issued.  
Administrators can manage the tokens of all users.

### Users

The users pane allows users with administrative priviledge to add and remove users.

## HTTP API

There is an experimental HTTP API, that is being used internally by the web interface. Authentication details and routes will become documented when its API stabilises.
