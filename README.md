# Yealink click2dial handler

A simple application to help register a system native click2dial url scheme handler like "tel:" and "callto:" for Yealink desktop phones.

- [Setup](#setup)
- [Installation](#installation)
  - [Linux desktop](#linux-desktop)
  - [Windows desktop](#windows-desktop)
- [Remarks / FAQ](#remarks--faq)
- [Thanks to / related](#thanks-to--related)

## Setup

These instructions are not specific to any OS. For proper OS installation instructions, follow the [installation](#installation)-steps below.

Ensure that your IP is listed or allowed in the `Action URI Allow IP List` setting under  `Features > Remote Control` on your desk phone. It might support wildcards like `*`, but please consult the documentation first.

Connect your phone with one of the two options:

```shell script
# Interactive mode, wizard like
# Example: yealink-click2dial connect 192.168.0.109 admin  
yealink-click2dial connect [phone-ip] [username]

# Non-interactive mode, no-questions-asked
# Example: yealink-click2dial connect 192.168.0.109 admin admin 030123@192.168.0.1
yealink-click2dial connect [phone-ip] [username] [password] [outgoing-uri]
```

After completion, the configuration file should've been created/stored.

For a simple test, try to dial any number through the CLI itself:

```shell script
yealink-click2dial dial [phone-number]
# Example: yealink-click2dial dial +4930123
```

The phone itself might prompt for permission on-screen, please accept it.

The phone might block you for several minutes after too many failed attempts, feel free to re-power the phone to release the lock early.

## Installation

### Linux desktop

This has been tested with KDE based UIs on Kubuntu/Ubuntu. If you use another distro feel free to add to this document and do a PR.

First install the binary from the release page or build it yourself. For simplicity, I chose `~/bin/yealink-click2dial` as the location.

Ensure you followed the [setup](#setup) instructions and have a working configuration.

First we need to register the MIME handler. This is done by adding it to one of the following locations (which should already exist):

```
~/.config/mimeapps.list
~/.local/share/applications/mimeapps.list
```

Look for the `[Added Associations]` section and append/overwrite the handler-line with the following entry:

```
[Added Associations]
x-scheme-handler/tel=yealink-click2dial.desktop;
```

It might be required to comment-out already existing handlers for the `x-scheme-handler/tel` scheme (like the KDE URL handler).

Now create the desktop entry itself. This can be done by creating a file `~/.local/share/applications/yealink-click2dial.desktop` with the following content:

```
[Desktop Entry]
Name=Yealink click2dial
Comment=Dials the given number via the configured desktop phone.
Exec=~/bin/yealink-click2dial dial %u
Terminal=false
Type=Application
```

To test the integration simple trigger it on the shell like this:

```
xdg-open tel:+4930123
```

It should simply dial. If it fails, but the test within the setup-chapter worked, there is a typo/error in the association done in this chapter.

### Windows desktop

This has been tested on Windows 10. If you use any other version or require different steps feel free to document and do a PR.

First install the windows binary into your reachable %PATH% for executables on Windows (for simplicity, I've put it into `C:\Windows`).

Ensure you followed the [setup](#setup) instructions and have a working configuration (and append `.exe` to your calls).

Install the default application listener [windows.reg](windows.reg) into the Windows registry and reboot windows.

Now test the integration on a command line prompt:

```
start "" "tel:+4930123"
```

Windows should prompt you with options of available applications, one of them "yealink-click2dial", which you select.

From now on, every "tel:" click in any kind of app or browser should trigger a dial flow on your phone.

## Remarks / FAQ

### Is the code super-clean?

Nope, this project expands a bit on the phone API like DND feature toggling that is not used. Why? Did it for something else, did not have the time to migrate it to a shared lib.

## Thanks to / related

https://wiki.lug-wr.de/wiki/doku.php?id=user:tstoeber:howto:href-tel-handler:start
https://superuser.com/questions/1112229/cant-change-tel-protocol-handler-in-windows-10

