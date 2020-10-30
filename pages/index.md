# What is boringproxy?

You're using it right now! The website you're reading is hosted on my home
computer, through boringproxy.

boringproxy is a combination reverse proxy and tunnel manager.

What that means is if you have a self-hosted web service (Nextcloud, Emby,
Jellyfin, personal website, etc) running on a private network (such as behind a
[NAT] at home), boringproxy aims to provide the easiest way to securely (ie
HTTPS and optional password-protection) expose that server to the internet, so you can
access it from anywhere.

To see how boringproxy compares to other similar software, see the
[comparison here](/tunneling-comparison/).

The main features are:

* 100% free and open source under the MIT license.
* Designed from the ground up with self-hosters in mind.
* No more port forwarding, NAT traversal, firewall rules, HTTPS certificate
  management, etc etc. Let boringproxy handle it all for you.
* No config files. Seriously, none. It has nice defaults and the few knobs are
  easily adjusted with simple CLI parameters.
* Lightning fast web GUI for managing tunnels from one central place. It even
  works great on mobile browsers.
* The client software works on Linux, Windows, Mac, and ARM (ie raspberry pi).
* Ships as single executable which contains both the server and client.
* SSH under the hood. You can use a standard SSH client if you prefer.

# TODO demo video

# Demo Instance

There is a demo instance at `https://brng.pro`. If you submit your email
address using the form below, it will create an account for you and send you
a login link to play with the demo.

<form action='https://demo-signup.boringproxy.io/request' method='POST'>
  <label for='email-input'>Email:</label>
  <input type='text' id='email-input' name='email'>
  <input type='submit' class='button'>
</form>


# What's with the name?

The name has two meanings; one pun and one philosophy. The pun is "bore" as in
bore a hole/tunnel, highlighting the fact that boringproxy is a reverse proxy
that automatically manages making tunnels.

The philosophy is that boring (as in boredom) software is often the most useful
software.  If you have to interact with boringproxy to get something done, I
hope it ends up being the least interesting part of your day. I want it to be a
tool that does its job well and gets out of the way.

This also has implications when it comes to adding features. I want boringproxy
to remain simple and focused. When contemplating adding any feature, the first
question I ask myself is: is it boring enough?


[NAT]: https://en.wikipedia.org/wiki/Network_address_translation

