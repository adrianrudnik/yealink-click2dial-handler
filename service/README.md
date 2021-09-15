# Yealink Click to Dial scheme handler

## Installation

### Linux desktop

Append the MIME additions in your profile file. It might be in one of these locations:

```
~/.config/mimeapps.list
~/.local/share/applications/mimeapps.list
```

and add the following entry:

```
[Added Associations]
x-scheme-handler/tel=yealink-click2dial.desktop;
```

and ensure that no other handler for the same is already configured.

Then create the desktop entry itself in `~/.local/share/applications/yealink-click2dial.desktop`.

```
[Desktop Entry]
Name=Yealink Click2Dial
Comment=Dials the given number via the configured desktop phone.
Exec=~/bin/yealink-click2dial dial %u
Terminal=false
Type=Application
```

## Thanks to

https://wiki.lug-wr.de/wiki/doku.php?id=user:tstoeber:howto:href-tel-handler:start
