# What is it?

boringproxy is a combination reverse proxy and tunnel manager.

If you have a web service (Nextcloud, Emby, Jellyfin, personal website, etc)
running on a private network (such as behind a [NAT] at home), boringproxy aims
to provide the easiest way to securely (ie HTTPS and password-protection)
expose that server to the internet, so you can access it from anywhere.

boringproxy provides a lightning fast web GUI for managing tunnels, and the
client software works on Linux, Windows, Mac, and ARM (ie raspberry pi).

The site you're reading right now is being hosted through boringproxy.

It is 100% free and open source under the MIT license. The server ships as
single executable, And requires extremely little configuration (no config file
and only a single required command line parameter).


# How does it compare to other software I may be familiar with?

## ngrok/localtunnel/etc

boringproxy is very similar to ngrok and other tunneling software. If you're
already familiar with this sort of thing, I recommend going straight to the
[tunneling comparison] to see how it's different. The TL;DR is that I would
consider most of the others to be tools made for developers which happen to be
useful for self-hosters.  boringproxy is a tool made specifically for
self-hosters (though it's still very useful for developers like myself). See
the full comparison for details.

[tunneling comparison]: #tunnel-comparison


## Reverse proxies (Caddy/nginx/traefik/etc)

Reverse proxies handle many different tasks. These typically involve things
like providing HTTPS in front of insecure servers, caching, load balancing,
automatic gzip compression, HTTP/2 connections, etc.

However, what these proxies lack for our use case is the ability to tunnel
the web connections into a private network. They can only access ports that are
directly available to the proxy server. So you have to either run your upstream
servers on the same network as the reverse proxy, or set up your own tunneling
system.

boringproxy solves both of these in one service. It includes a minimal, fast
reverse proxy with a robust SSH-based tunneling system.

## Commercial VPN products (NordVPN, ExpressVPN, PIA, etc)

People use commerical [VPN]s for a variety of purposes. Two of the most common
are accessing content that's only available in certain countries, and hiding
your browsing behavior from your ISP.

While the technology involved is very similiar, the goals of boringproxy and
a VPN are somewhat different. boringproxy provides the outside world with
secure access to your private services. VPNs allow you to tunnel out of your
private network so you appear to be connecting from somewhere else.


## Personal VPN (OpenVPN/WireGuard/etc)

Personal VPNs again share a lot of technology with boringproxy. In fact,
if you had a WireGuard network set up and put Caddy on one of your public
servers, that would be a great way to accomplish what boringproxy does. The
downside is you have to manage all the pieces separately.


## Cloudflare

Cloudflare is sort of like a supercharged reverse proxy. In additional to all
of the features mentioned above, it also provides a [CDN] which caches and
distributes your server responses all over the world with low latency.

Cloudflare even includes a service for tunneling, which is discussed in the
full [tunneling comparison].

boringproxy and Cloudflare are actually a great combination. You can use
boringproxy to tunnel your services, and Cloudflare as a free CDN.


## Web hosting services (Wordpress/Netlify/etc)


# Comparison with other tunneling options <a name='tunnel-comparison'></a>


## ngrok

ngrok is probably the best-known tunneling software available.

## localtunnel

localtunnel is a popular open source/self-hosted alternative to ngrok.


# What's with the name?

The name has two meanings; one pun and one philosophy. The pun is "bore" as in
bore a tunnel, highlighting the fact that boringproxy is a reverse proxy that
automatically manages making tunnels.

The philosophy is that boring (as in boredom) software is often the most useful
software.  If you have to interact with boringproxy to get something done, I
hope it ends up being the least interesting part of your day. I want it to be a
tool that does its job well and gets out of the way.

This also has implications when it comes to adding features. I want boringproxy
to remain simple and focused. When contemplating adding any feature, the first
question I ask myself is: is it boring enough?


[NAT]: https://en.wikipedia.org/wiki/Network_address_translation

[CDN]: https://en.wikipedia.org/wiki/Content_delivery_network

[VPN]: https://en.wikipedia.org/wiki/Virtual_private_network
