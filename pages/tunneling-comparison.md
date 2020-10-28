# How does it compare to other software I may be familiar with?

## Other tunneling solutions (ngrok, localtunnel, etc)

boringproxy is very similar to ngrok and other tunneling software. There are a
lot out there (I maintain a list [here][0]). 

The TL;DR value proposition of boringproxy is that I would consider most of the
others to be tools made for developers which happen to be useful for
self-hosters.  boringproxy is a tool made specifically for self-hosters (though
it's still very useful for developers like myself).

In my opinion, the key features that a solid tunneling solution needs to
provide are:

* 100% open source server and client.
* Easy to self host.
* Built-in reverse proxy.
* Supports custom domains/subdomains.
* Automatically manages HTTPS certificates from Let's Encrypt.

There are only a small handful of options that meet all those requirements. In
addition to these core features, the things that boringproxy does differently
from other solutions are:

* Clients can be remote-controlled through the server web GUI and REST API.
  This means you can start the client up on each machine you want to tunnel out
  of and then forget about it. No more creating a separate ngrok process for
  each tunnel.
* Very simple to run. Single executable download, and essentially no
  configuration needed.


[0]: https://github.com/anderspitman/awesome-tunneling


## Other reverse proxies (Caddy, nginx, traefik, etc)

Reverse proxies handle many different tasks. These typically involve things
like providing HTTPS in front of insecure servers, caching, HTTPS certificate
management, load balancing, automatic gzip compression, HTTP/2 connections,
etc.

However, what these proxies lack for our use case is the ability to tunnel the
web connections into a private network. They can only access IP addresses and
ports that are directly available to the proxy server. So you have to either
run your upstream servers on the same network as the reverse proxy, or set up
your own tunneling system.

boringproxy handles both of these in one service. It includes a minimal, fast
reverse proxy with a robust SSH-based tunneling system.

## Commercial VPN products (NordVPN, ExpressVPN, PIA, etc)

People use commerical [VPN]s for a variety of purposes. Two of the most common
are accessing content that's only available in certain countries, and hiding
your browsing behavior from your ISP.

While the technology involved is very similiar, the goals of boringproxy and
a VPN are somewhat different. boringproxy provides the outside world with
secure access to your private services. VPNs allow you to tunnel out of your
private network so you appear to be connecting from somewhere else.


## Personal VPN (OpenVPN, WireGuard, etc)

Personal VPNs again share a lot of technology with boringproxy. In fact,
if you had a WireGuard network set up and put Caddy on one of your public
servers, that would be a great way to accomplish what boringproxy does. The
downside is you have to manage all the pieces manually.


## Cloudflare

Cloudflare is sort of like a supercharged reverse proxy. In additional to all
of the normal features, it also provides a [CDN] which caches and distributes
your server responses all over the world with low latency.

Cloudflare even includes a service for tunneling, but sadly the server is not
open source.

boringproxy and Cloudflare are actually a great combination. You can use
boringproxy to tunnel your services, and Cloudflare as a free CDN and solid
DNS management service.


## Web hosting services (Wordpress, Netlify, etc)

One nice thing about boringproxy is it makes it extremely easy to host your own
website. All you have to do is start a basic webserver on your home computer,
then use the web GUI to create a tunnel, and start writing HTML files.

This isn't quite as easy as hosting your blog on wordpress.com, but it's very
simple, and let's you see how things are actually working.

It also has the advantage over services like Netlify or GitHub pages in that
you don't have to learn git just to make a website.


[CDN]: https://en.wikipedia.org/wiki/Content_delivery_network

[VPN]: https://en.wikipedia.org/wiki/Virtual_private_network
