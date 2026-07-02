#!/usr/bin/env python3
"""
Simple historically accurate Enigma I simulator.

*** THIS CODE WAS "WRITTEN" BY AN AI ***
*** DON'T TRUST IT ***
"""


class Util:
    @classmethod
    def c2i(cls, char):
        return ord(char) - 65

    @classmethod
    def i2c(cls, idx):
        return chr(idx % 26 + 65)


class Rotor:
    ROTORS = {
        "I": "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
        "II": "AJDKSIRUXBLHWTMCQGZNPYFVOE",
        "III": "BDFHJLCPRTXVZNYEIWGAKMUSQO",
        "IV": "ESOVPZJAYQUIRHXLNFTGKDCMWB",
        "V": "VZBRGITYUPSDNHLXAWMJQOFECK",
    }

    NOTCH = {
        "I": "Q",
        "II": "E",
        "III": "V",
        "IV": "J",
        "V": "Z",
    }

    def __init__(self, name, pos="A", ring="A"):
        self.w = self.ROTORS[name]
        self.inv = [""] * 26

        print(f"N: {self.w}")
        for i, ch in enumerate(self.w):
            self.inv[Util.c2i(ch)] = Util.i2c(i)

        print(f"I: {"".join(self.inv)}")

        self.pos = Util.c2i(pos)
        self.ring = Util.c2i(ring)
        self.notch = Util.c2i(self.NOTCH[name])

    def step(self):
        self.pos = (self.pos + 1) % 26

    def at_notch(self):
        return self.pos == self.notch

    def fwd(self, c):
        x = (Util.c2i(c) + self.pos - self.ring) % 26
        y = Util.c2i(self.w[x])
        return Util.i2c((y - self.pos + self.ring) % 26)

    def rev(self, c):
        x = (Util.c2i(c) + self.pos - self.ring) % 26
        y = Util.c2i(self.inv[x])
        return Util.i2c((y - self.pos + self.ring) % 26)


class Enigma:
    ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

    REFLECTORS = {
        "B": "YRUHQSLDPXNGOKMIEBFZCWVJAT",
        "C": "FVPJIAOYEDRZXWGCTKUQSBNMHL",
    }

    def __init__(
        self, rotors=("I", "II", "III"), rings="AAA", pos="AAA", refl="B", pairs=()
    ):
        self.l = Rotor(rotors[0], pos[0], rings[0])
        self.m = Rotor(rotors[1], pos[1], rings[1])
        self.r = Rotor(rotors[2], pos[2], rings[2])
        self.ref = self.REFLECTORS[refl]
        self.pb = {c: c for c in self.ALPHABET}
        for p in pairs:
            a, b = p[0], p[1]
            self.pb[a] = b
            self.pb[b] = a

    def reflect(self, c):
        return self.ref[Util.c2i(c)]

    def step(self):
        if self.m.at_notch():
            self.l.step()
            self.m.step()
        elif self.r.at_notch():
            self.m.step()
        self.r.step()

    def enc_char(self, ch):
        if ch not in self.ALPHABET:
            return ch

        self.step()
        c = self.pb[ch]
        c = self.r.fwd(c)
        c = self.m.fwd(c)
        c = self.l.fwd(c)
        c = self.reflect(c)
        c = self.l.rev(c)
        c = self.m.rev(c)
        c = self.r.rev(c)
        return self.pb[c]

    def encrypt(self, text):
        return "".join(
            self.enc_char(c) if c in self.ALPHABET else c for c in text.upper()
        )


if __name__ == "__main__":
    rotor = Rotor("I")
    # e = Enigma(
    #     rotors=("III", "II", "I"),
    #     rings="AAA",
    #     pos="AAA",
    #     pairs=["AM", "FI", "NV", "PS", "TU", "WZ"],
    # )
    # while True:
    #     try:
    #         s = input("Text: ")
    #     except EOFError:
    #         break
    #     print(e.encrypt(s))
