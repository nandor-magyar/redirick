# Redirick

A lightweight, by-default rickrolling (configurable), redirect-only webserver that can be used CLI tool or as a container.

## Story
I wanted a simple container that only does redirect and needs the least possible privileges,
I could not find anything that is:
- lightweight
- independent
- rootless
- environment configurable
- easy to deploy
So here it is.

## What it does

This is a basic go program that has a single rule `/*` that catches all incoming requests and (302) redirects to Rick Astley's Never gonna give you up (if not directed otherwise).

Gotta rick-roll 'em fellas!

You can change the `REDIRECT_TARGET` with an environment variable, the bound `PORT` (default: 8080), and the redirect's `STATUS_CODE` (default: 302).

## Status code
https://www.searchenginejournal.com/301-vs-302-redirects-seo/299843/#close
Based on this article as Google stated multiple times that 302 does not hurt link values, since 301 request are cached eternally, 302 should theoretically be better choice. Feel free to comment on this.

## Usage
If your use-case is missing, file a PR I am open to anything.

### go install
If you need short-term temporal redirect, you can just go install this to your target machine.
```
go install github.com/nandor-magyar/redirick@main
redirick help

``` 

### docker-compose
From the docker-compose folder of the project, make sure to checkout environment variables you need, use the `.env.example` file.
The compose file is also armed for a basic traefik usage.
```
docker-compose up -d
```

### kustomize
```
kubectl apply -k kustomize
```

### dyrectorio
Based on the labels declared on the image you have to only fill-in the required env variables and it should work out of the box if you have a node running. No-code, no hassle, just the two variables.

## DIY
Requirements: make, go
```bash
# compile go binary 
make compile
# create custom image
docker build -t your-custom-tag:version .
# use as You like
```
