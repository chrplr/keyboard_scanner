simple HID (keyoard and mouse) event monitor
============================================

`Time-stamp: <2024-12-11 08:19:55 christophe@pallier.org>`

![screenshot](app_small.png)


Usage:

```
python keyboard_scanner.py > event.tsv
```

The first column is the time in ms since the start, the second column is the time difference between between the current event and the previous one.

Installation:

 * This script requires the `pygame` module (`pip install pygame`).
 

