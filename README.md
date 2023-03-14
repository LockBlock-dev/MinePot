# MinePot

[![GitHub stars](https://img.shields.io/github/stars/LockBlock-dev/MinePot.svg)](https://github.com/LockBlock-dev/MinePot/stargazers)

MinePot is a Minecraft Server Honeypot made in Golang. Its goal is to catch Minecraft Server Scanners by listening for [Handshake](https://wiki.vg/Protocol#Handshake) and [Ping](https://wiki.vg/Protocol#Status) packets.

See the [changelog](/CHANGELOG.md) for the latest updates.

## Table of content

-   [**Features**](#features)
-   [**Installation**](#installation)
-   [**Compiling from source**](#compiling-from-source)
-   [**Configuring MinePot**](#configuring-minepot)
-   [**Config details**](#config-details)
-   [**FAQ**](#faq)
-   [**Credits**](#credits)
-   [**Copyright**](#copyright)

## Features

-   Listen on any TCP port for incoming Minecraft packets
-   Answer [Handshake](https://wiki.vg/Protocol#Handshake) packets
-   Answer [Ping](https://wiki.vg/Protocol#Status) packets
-   Custom Status Response :
    -   Custom version or version mirroring (send the received protocol version)
    -   Fake players
    -   Custom MOTD
    -   Custom favicon
-   IP reporting to [Abuse IP DB](https://www.abuseipdb.com/)
-   IP reporting to a Discord Webhook
-   History as a [CSV](https://en.wikipedia.org/wiki/Comma-separated_values) formatted .history file

## Installation

-   Download [go](https://go.dev/dl/) (go 1.20 required).
-   Download or clone the project.
-   Download the binary from the [Releases](../../releases) or [build it](#compiling-from-source) yourself.
-   [Configure MinePot](#configuring-minepot).
-   Install MinePot by using [`install.sh`](/install.sh). It will setup the tool and start it as a service for you.

## Compiling from source

-   Use [`build.sh`](/build.sh) or use `go build` in [`src`](/src)

## Configuring MinePot

If you already used [`install.sh`](/install.sh), the config can be found in `/etc/minepot/config.json`.

-   Open the [`config`](/config.json) in your favorite editor.
-   Enable the features you want to use. See [Config details](#config-details) for in-depth explanations.
-   Edit the Status Response as you want. You can use [mctools MOTD creator](https://mctools.org/motd-creator) for the MOTD.
-   Change the `faviconPath` to any PNG image you want to use.

## Config details

| Item               | Values                                                     | Meaning                                            |
| ------------------ | ---------------------------------------------------------- | -------------------------------------------------- |
| debug              | `boolean`                                                  | Enable debug logs                                  |
| writeLogs          | `boolean`                                                  | Enable logs file                                   |
| logFile            | `text`                                                     | Path to the logs file                              |
| writeHistory       | `boolean`                                                  | Enable history file                                |
| historyFile        | `text`                                                     | Path to the history file                           |
| port               | `number`                                                   | TCP port to listen on                              |
| pollIntervalMs     | `number`                                                   | Time to wait between each packet (in milliseconds) |
| idleTimeoutS       | `number`                                                   | Time to wait before the connection times out       |
| reportThreshold    | `number`                                                   | Amount of packets before being reported            |
| abuseIPDBReport    | `boolean`                                                  | Enable Abuse IP DB reports                         |
| abuseIPDBKey       | `text`                                                     | Abuse IP DB API key                                |
| abuseIPDBCooldownH | `number`                                                   | Cooldown between each reports (in hours)           |
| webhookReport      | `boolean`                                                  | Enable Discord webhook reports                     |
| webhookUrl         | `text`                                                     | Discord webhook URL                                |
| webhookCooldownH   | `number`                                                   | Cooldown between each reports (in hours)           |
| webhookEmbedColor  | `text`                                                     | Embed hex color                                    |
| statusResponse     | `boolean`                                                  | Enable Status Response                             |
| statusResponseData | [`JSON`](https://wiki.vg/Server_List_Ping#Status_Response) | Minecraft Status Reponse data                      |
| faviconPath        | `text`                                                     | Path to the favicon PNG image                      |

## FAQ

-   Q: Do you plan to release a Windows version?  
    A: No.

## Credits

-   [Wiki.vg](https://wiki.vg) Minecraft protocol documentation

## Copyright

See the [license](/LICENSE).
