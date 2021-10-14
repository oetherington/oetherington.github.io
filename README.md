# www.etherington.xyz

This repo contains the source and content for www.etherington.xyz. The content
is written dynamically, but is then compiled and served statically.

In order to be served by github pages, the compiled output is in the "docs"
directory, which should therefore not be edited manually.

A sitemap is automatically generated using the git history to get the last
modified date for each file.

#### Install build dependencies

`make install-deps`

#### Build the site

`make`
