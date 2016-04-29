<!--[metadata]>
+++
title = "Gandi"
description = "Gandi driver for docker machine"
keywords = ["machine, Gandi, driver, docker"]
[menu.main]
parent="smn_machine_drivers"
+++
<![end-metadata]-->

# Docker Machine driver plugin for Gandi

This plugin adds support for [Gandi](https://www.gandi.net/) cloud instances to the `docker-machine` command line tool.

[![CircleCI](https://img.shields.io/circleci/project/Gandi/docker-machine-gandi.svg)](https://circleci.com/gh/Gandi/docker-machine-gandi/)

## Installation

Requirement: [Docker Machine >= 0.5.1](https://github.com/docker/machine)

Download the `docker-machine-driver-gandi` binary from the release page.
Extract the archive and copy the binary to a folder located in your `PATH` and make sure it's executable (e.g. `chmod +x /usr/local/bin/docker-machine-driver-gandi`).

## Usage instructions

Grab your API key from the [Gandi admin](https://www.gandi.net/admin/api_key) and pass it to `docker-machine create` with the `--gandi-api-key` option.

### Example

    $ docker-machine create --driver gandi \
                            --gandi-api-key=abc123 \
                            ubuntu-machine

[Read more about Docker support on Gandi servers](https://wiki.gandi.net/iaas/references/server/docker).

### Command line flags:

 - `--gandi-api-key`: **required** Your Gandi API key.
 - `--gandi-image`: Image used to create the machine. Default is "Ubuntu 16.04 64 bits LTS (HVM)".
 - `--gandi-datacenter`: Datacenter where machine will be created. Default is Bissen, Luxembourg (LU-BI1).
 - `--gandi-memory`: Memory size in MB. Default is 512 MB.
 - `--gandi-core`: Number of cores for the machine. Default is 1 CPU core.
 - `--gandi-url`: API url to connect to. Default is the production endpoint URL.

