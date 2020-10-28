# What is boringproxy?

You're using it right now! The website you're reading is hosted on my home
computer, through boringproxy.

boringproxy is a combination reverse proxy and tunnel manager.

What that means is if you have a web service (Nextcloud, Emby, Jellyfin,
personal website, etc) running on a private network (such as behind a [NAT] at
home), boringproxy aims to provide the easiest way to securely (ie HTTPS and
password-protection) expose that server to the internet, so you can access it
from anywhere.

To see how boringproxy compares to other software, see the
[comparison here](/tunneling-comparison/).

boringproxy provides a lightning fast web GUI for managing tunnels, and the
client software works on Linux, Windows, Mac, and ARM (ie raspberry pi). You
can even use the basic functionality with a standard SSH client.

boringproxy is 100% free and open source under the MIT license. It ships as
single executable which contains both the server and client, and requires
extremely little configuration (no config file and only one required command
line parameter).

# TODO demo video

# TODO demo instance signup

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
