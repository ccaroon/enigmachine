# CHANGELOG

## 2026.08.01
* Added Error handling to ParseKeySpec()
* Added ability to encipher a file
* Added output formatting options to encipher cmd
  * keep original formatting
  * Let group/block size
  * Blocks per line

## 2026.07.28
* Updated `NewPlugboard` to accept a list of letter pair strings for the specification instead of a list of individual runes.
    - E.e. ["AX","JR"] instead of ['A','X','J','R']
