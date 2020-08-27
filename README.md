# dweb

dweb is a tool for managing your server, with more customizable options in a centralized location.
In particular, it also aims to allow containerization of several common server features.

- [dweb](#dweb)
  - [How it works](#how-it-works)
  - [Configuration](#configuration)
  - [Comparison with IaaS](#comparison-with-iaas)
  - [Why Go?](#why-go)
  - [Features](#features)
    - [HTTP](#http)
    - [SSH](#ssh)
    - [Planned](#planned)
  - [License](#license)

## How it works

The plan is to use the CLI only for basic things such as installing and uninstalling each feature.
However, there may be some configuration options available to add, edit, and remove users or other aspects of each feature.
In any case, each feature will be maintained fairly consistently, with documentation available in the README.

## Configuration

The plan is to use the canonical configuration directory for each platform, with a directory called `dweb` and configuration TOML files for each feature in the directory.
For example, on Linux, there would be files such as `/etc/dweb/ssh.toml`, and so on.

This is subject to change, and as of right now the configuration files haven't been implemented yet.

## Comparison with IaaS

I don't think this project is so much IaaS, honestly.
The tool may eventually allow for IaaS provisioning, but if it does then this still wouldn't be IaaS, so much as a host-managing tool.
We'll see though.

## Why Go?

Honestly, I just want to implement as much things as quickly as possible.
Rust is a great language, and if I have time in the future I'd prefer to RiiR because of the performance and safety bonuses associated.
However, I think Go is still a solid choice for this project, since most features are/will be network-focused.

Another great advantage I think is the fact that Docker has an official Go SDK for their API.
Yes, it is possible to interact with Docker via HTTP, but that would mean writing a ton of code that I don't want to focus on.
Maybe in the future, if there is a solid SDK for Docker in Rust, that would prompt me to just stop everything and rewrite everything here in Rust, but from what I've seen there just isn't enough support right now.

## Features
### HTTP

TODO

### SSH

TODO

### Planned

Right now, these are what I plan on supporting with dweb.
Most things will have Docker options, but some (ie SMTP) just don't make sense to Dockerize.

## License

This repository is MIT-licensed.
Code is copyright Colton Donnelly 2020.
