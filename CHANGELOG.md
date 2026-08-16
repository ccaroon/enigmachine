# CHANGELOG

## 2026.08.16
* Use `rune` for Reflector ID instead of `string`
* Updated the `enigma.Key` struct to be more like the `keysheet.Entry` struct, then reuse `Key` in `Entry` ... Dedupes a bit of code.
* Initial integration of Key Sheet into `encipher` command
* Bit of code refactoring

## 2026.08.15
* `key-sheet` -> `keysheet`
* `keysheet generate` now saves the generated keysheet to a file
* Added `keysheet view` command

## 2026.08.09
* Added `key-sheet generate` command

## 2026.08.03
* Added more of error handling

## 2026.08.01
* Added Error handling to `ParseKeySpec()`
* Added ability to encipher a file
  - `enigmachine encipher @FILE-NAME`
* Added output formatting options to encipher command
  * `--original   | -o`: keep original formatting
  * `--block-size | -b`: group lettings in blocks of this size
  * `--line-size  | -l`: blocks per line

## 2026.07.28
* Updated `NewPlugboard` to accept a list of letter pair strings for the specification instead of a list of individual runes.
    - E.g. ["AX","JR"] instead of ['A','X','J','R']
