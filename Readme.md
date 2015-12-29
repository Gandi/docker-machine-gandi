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

<!-- [![CircleCI](https://img.shields.io/circleci/project/gandi/docker-machine-gandi.svg)](https://circleci.com/gh/gandi/docker-machine-gandi/) -->

## Installation

Requirement: [Docker Machine >= 0.5.1](https://github.com/docker/machine)

Download the `docker-machine-driver-gandi` binary from the release page.
Extract the archive and copy the binary to a folder located in your `PATH` and make sure it's executable (e.g. `chmod +x /usr/local/bin/docker-machine-driver-gandi`).

## Usage instructions

Grab your API key from the [Gandi admin](https://www.gandi.net/admin/api_key) and pass that to `docker-machine create` with the `--gandi-api-key` option.


**Example for creating a new machine running default Ubuntu 14.04:**

    docker-machine create --engine-storage-driver devicemapper \
                          --driver gandi \
                          --gandi-api-key=abc123 \
                          ubuntu-machine

Command line flags:

 - `--gandi-api-key`: **required** Your Gandi API key.
 - `--gandi-image`: Image to use to create machine, default Ubuntu 14.04 64 bits LTS (HVM).
 - `--gandi-datacenter`: Datacenter where machine will be created, default Bissen.
 - `--gandi-memory`: machine memory size in MB, default 512.
 - `--gandi-core`: Number of cores for the machine, default 1.
 - `--gandi-url`: url to connect to.
