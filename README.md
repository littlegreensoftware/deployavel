[![Build Status][ci-img]][ci]

## deployavel 

deployavel is a CLI application written in GO to
easily provision and manage Digital Ocean servers for Laravel based applications.

Usage:
  
  ```deployavel [command]```

Available Commands:

  - ```create``` Create a single resource in Forge
  - ```get``` Get a single resource in Forge
  - ```help``` Help about any command
  - ```list``` List all resources of a specific type in Forge

Flags:

  - ```--config path/to/config.yml``` config file (default is $PWD/config.yml)
  - ```-h, --help``` help for deployavel

Use ```deployavel [command] --help``` for more information about a command.

[Forge API Docs](https://forge.laravel.com/api-documentation)

[ci-img]: https://travis-ci.org/littlegreensoftware/deployavel.svg?branch=master
[ci]: https://travis-ci.org/littlegreensoftware/deployavel