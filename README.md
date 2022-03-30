![ALERT MAJOR REFACTORING](https://upload.wikimedia.org/wikipedia/commons/thumb/4/4e/OOjs_UI_icon_alert_destructive.svg/480px-OOjs_UI_icon_alert_destructive.svg.png)

#MAJOR REFACTOR IN PROGRESS

# liblsdj

[Little Sound DJ](http://littlesounddj.com) is wonderful tool that transforms your old gameboy into a music making machine. It has a thriving community of users that pushes their old hardware to its limits, in pursuit of new musical endeavours. It can however be cumbersome to manage songs and sounds outside of the gameboy.

In this light *liblsdj* is being developed, a cross-platform and fast Go library for interacting with the LSDJ save format (.sav), song files (.lsdsng) and more. The end goal is to deliver *liblsdj* and afterwards a suite of tools for working with everything LSDJ.

This is a Go port of [liblsdj](https://github.com/stijnfrishert/liblsdj).

## System requirements
You need only a working Go environment installed.

Every architecture supported by Go should work fine.

## Installation

**not ready yet**

## Todo
- [ ] Compress .lsdsng song
- [ ] Decompress .lsdsng song
- [ ] Read song from .lsdsng
- [ ] Write song onto .lsdsng
- [ ] Read from .sav
- [ ] Write to .sav
- [ ] Expose a decent api
- [ ] Write a test suite

## Author

[Savino Pio Liguori](https://8lall0.github.io/), [@twitter](https://twitter.com/imblellow).

## License
This project is licensed under the MIT License - see the LICENSE.md file for details