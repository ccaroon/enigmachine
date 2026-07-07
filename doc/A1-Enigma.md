# A1: Enigma

In this assignment, you will develop a software replica of the Enigma
encryption machine used by the German military in World War II.  Efforts
to break the Enigma cipher led to the development of the first
electronic computers at Bletchley Park, England: the Bombe (designed by
Alan Turing) and its successor, Colossus (designed by Tommy Flowers).
Using computers, the Allies were eventually able to break the Enigma
code, giving them an intelligence edge that changed the balance of the
war.

**Table of Contents:**

- [A1: Enigma](#a1-enigma)
  - [Introduction to the Enigma ](#introduction-to-the-enigma-)
  - [Assignment information ](#assignment-information-)
  - [Part 0:  Sanity check ](#part-0--sanity-check-)
  - [Part 1:  Index ](#part-1--index-)
  - [Part 2:  Rotors ](#part-2--rotors-)
  - [Interlude:  Test sweep ](#interlude--test-sweep-)
  - [Part 3: Reflector ](#part-3-reflector-)
  - [Part 4:  Plugboard ](#part-4--plugboard-)
  - [Part 5:  Ciphering a character ](#part-5--ciphering-a-character-)
  - [Part 6:  Stepping ](#part-6--stepping-)
  - [Part 7: Ciphering a string ](#part-7-ciphering-a-string-)
  - [What to turn in ](#what-to-turn-in-)
  - [Assessment ](#assessment-)
  - [Karma ](#karma-)
  - [](#)

## Introduction to the Enigma <a name="intro"></a>

As a visual introduction to the Enigma, please watch this video:
https://www.youtube.com/watch?v=G2_Q9FoD-oQ.  Then continue with reading this
writeup.

Have you watched the video?  Good.  :)

The Enigma was not just a single machine, but a family of closely related
machines.  So keep in mind that there might be details that differ from
one video to another that you watch online, or even between a video and
this writeup.

**Substitution ciphers.**
The Enigma machine implemented a *substitution cipher*, which encrypts a
message by substituting one character for another.  Such ciphers go back
at least as far as Julius Caesar, who used a simple substitution cipher
to encrypt military orders.  The Caesar cipher specified that each
letter in the alphabet would be encrypted using the letter three
positions later in the alphabet. For example, 'A' would be encrypted as
'D', 'B' would be encrypted as 'E', 'C' would be encrypted as 'F', and
so on. The code wraps around at the end of the alphabet, so 'X', 'Y' and
'Z' would be encrypted as 'A', 'B', and 'C', respectively.
Although the Caesar cipher was effective in its time (when very few
people could read at all), its simple pattern of encrypting letters
seems pretty obvious today.

A *keyed* substitution cipher is more effective.  The *key* is a secret
presumably known only to the message sender and receiver; it defines the mapping
between a *plaintext* letter and a *ciphertext* letter.  For the Caesar cipher,
we can write the key as a 26-character string, in which the first character is
the encryption of 'A', the second character is the encryption of 'B', and so on.
The key for the Caesar cipher would thus be
"DEFGHIJKLMNOPQRSTUVWXYZABC":
```
Plaintext letter:  ABCDEFGHIJKLMNOPQRSTUVWXYZ
                   |
                   V
Ciphertext letter: DEFGHIJKLMNOPQRSTUVWXYZABC
```
Of course, there are many other possible
keys:  since there are 26! (roughly 4 \* 10^26) different arrangements
of the 26 letters in the alphabet, there are 26! different keys and so
26! different keyed substitution ciphers.  Nonetheless, this simple type
of encryption is vulnerable to statistical attacks, as anyone who has
solved cryptogram puzzles can attest.

**The Enigma machine.**
https://en.wikipedia.org/wiki/Enigma_machine

The key for the Enigma cipher involved many aspects of the physical
configuration of the machine; we describe these below.  After the machine was
configured with the key, an operator pressed a letter on a typewriter-like
keyboard,  which caused a circuit to close and light up a letter on a
*lampboard* aka *lightboard*, a kind of simple display screen.  (Note the two
different uses of "key" in the previous sentence:  the first refers to a
*cryptographic key*, which is a secret used as an input to encryption and
decryption algorithms;  the second is in the word "keyboard", which refers to a
button that is pressed.) To encrypt a message, the operator typed the plaintext
letters, the ciphertext letters would light up, and the operator would record
each ciphertext letter on paper. Likewise, to decrypt a message, the operator
typed the ciphertext letters, the plaintext letters would light up, and the
operator would record each plaintext letter on paper.

What the diagram calls a *wheel* is also called a *rotor* in literature
about the Enigma.  Henceforth, we will use the term *rotor*.  The rotors
were installed next to each other on a *spindle*, around which they rotated.

You'll see the alphabet, in order, on each rotor in that picture.  Sometimes
rotors were labeled with the alphabet, and sometimes (as in the video linked
above) they were instead labeled with the integers 1..26.  It makes no
difference to the operation of the machine which labeling scheme was used:  'A'
corresponds to 1, 'B' to 2, and so forth.

In the wiring diagram above, there is another disc at the far left of the
spindle that looks something like a rotor but in fact does not rotate; it is
called the *reflector*.

As the diagram shows, when the operator typed a key on the keyboard, electrical
current would flow through the plugboard to the rotors, through the reflector,
back through the rotors and plugboard, and light up a letter on the lightboard.
The reflector, wheels, and plugboard are all mechanisms that could be installed
in different ways into the machine.  The way in which they all were collectively
installed determined the key that was being used for encryption or decryption.

Here is a brief description of each of the mechanisms; we go into greater
detail in later parts of this writeup:

* Each **rotor** implements its own, distinct substitution cipher by wires that
  connect the 26 contacts on its left side to the 26 contacts on its right side.
  The combination of many rotors together on the spindle creates an even more
  sophisticated substitution cipher. The original Enigma machines used at the
  beginning of World War II had three rotors, known by the Roman-numeral
  designations I, II, and III, which could be placed on the spindle in any
  order.  Later, other rotors with different wirings were manufactured, as were
  machines that could accept more than three rotors on the spindle.

* The **reflector** also implements a substitution cipher, though a less
  sophisticated one.  It has 26 contacts only on its right side, but none on its
  left.  Current that it receives from the right side is wired to be reflected
  back out that side, but in a different position. The reflector was always
  inserted at the leftmost side of the spindle. A keypress caused current to
  flow right-to-left through the rotors, through the reflector, then back
  through the rotors, left-to-right.  The reflector's substitution pattern
  caused the entire circuit to be symmetric, meaning that if 'E' mapped to 'Q'
  in a particular machine state, then 'Q' would likewise be mapped to 'E' in
  that same state. The minor advantage of this design was that the machine did
  not need separate modes of operation for decryption and encryption. But it
  turned out to be a major cryptographic weakness, enabling the code to be
  cracked.

* The **plugboard** implements yet another substitution cipher, which
  simply swaps pairs of letters.  Though the wiring of the rotors and reflector
  was pre-determined by how they were manufactured, the wiring of the plugboard
  could easily be changed by the operator by plugging cables into ports. The
  plugboard could have up to 13 cables inserted in it.  A cable inserted between
  two letters caused those letters to be swapped.  Even though it seems simpler
  than a rotor, the plugboard was a major source of cryptographic strength for
  the Enigma cipher.

* The **static wheel** shown above is of no interest to us.  (It is effectively
  a no-op.)

What made the Enigma cipher especially difficult to break was that just after
the operator typed a letter, and just before that letter was enciphered, the
rotors moved, changing the substitution pattern. Thus, the letter 'E' might get
mapped to a 'Q' the first time it was enciphered, but to a completely different
letter each successive time. The rightmost rotor stepped for every letter typed,
and the other rotors stepped conditionally based on *notches* on the rotors:
when a rotor stepped past a notch, not only did that rotor step, but also the
rotor to its left.  This stepping mechanism made the rotors work somewhat like
an old-style mechanical car odometer.

When a rotor was installed, the operator could manually rotate it to be in any
of the 26 possible orientations, based on the 26 labels appearing on the rotor
(either 'A'..'Z' or 01..26).   There was a window in the machine through which
the operator could observe the top letter showing on each rotor. The letter that
would be on top of the rotor at the point in time when the notch engaged was
known as the *turnover*.

*(Note:  we are omitting one more Enigma mechanism, the *ring setting*. For
aficionados of the Enigma, in this assignment you should assume the ring
setting is always 'A' for each rotor.  Anyone who doesn't understand what that
means can safely ignore it.)*

**Simulators.**  As you proceed with building a simulation of the Enigma machine
in this assignment, it could be useful to have a completed simulator at your
side.  The best Enigma simulator we can recommend is [this one][heath].
It provides excellent insight into the physical machine, especially the wiring
of the rotors and the reflector.  To assemble it, you need to buy a tube of
Pringles potato chips, print out the paper at 100% scaling (which can be
tricky because it's A4; to help with that we provide these close-ups [[1]][p1]
[[2]][p2] of the important parts that should be printable on letter paper),
then cut out the paper pieces and wrap them around the tube.  **It's
totally worth your while to assemble this simulator.**

There are also many online simulators, such as [this one][enigmaco] or
[this one][palloks]. They are useful if you want to construct your own
test cases, but they are not as useful as the Pringles-tube model if
your goal is understanding the operation of the machine.

[heath]: http://wiki.franklinheath.co.uk/index.php/Enigma/Paper_Enigma
[enigmaco]: http://enigmaco.de/
[palloks]: http://people.physik.hu-berlin.de/~palloks/js/enigma/enigma-u_v20_en.html
[p1]: p1.pdf
[p2]: p2.pdf

## Assignment information <a name="info"></a>

**Objectives.**

* Gain familiarity with basic OCaml language features, including
  built-in data types, lists, and functions.
* Practice writing programs in the functional style using immutable data.
* Improve skills in computational thinking by learning about and modeling
  a mildly complex, real-world computational artifact.

**Recommended reading.**

* [Lectures and recitations 1&ndash;4][web]
* The [CS 3110 Style Guide][style]
* The policies about deadlines, extensions, and late submissions
  in the [course syllabus][late]

[web]: http://www.cs.cornell.edu/Courses/cs3110/2017fa/
[style]: http://www.cs.cornell.edu/Courses/cs3110/2017fa/handouts/style.html
[syllabus]: http://www.cs.cornell.edu/Courses/cs3110/2017fa/syllabus.php
[ai]: http://www.cs.cornell.edu/Courses/cs3110/2017fa/syllabus.php#ai
[late]: http://www.cs.cornell.edu/Courses/cs3110/2017fa/syllabus.php#late

**Requirements.**

Your task in this assignment is to simulate the operation of
an Enigma machine.

* Your solution must provide the functions described below,
  with exactly those names and types.
* Your solution must correctly implement the Enigma cipher,
  including the spindle, rotors, reflector, turnovers, starting
  positions, and plugboard of the Enigma machine.
* Your code must be written with good style and be well documented.
* You must submit an OUnit test suite, as described below.

**What we provide.**

In the release code on the course website you will find these files:

* A template file `enigma.ml` with some function stubs.
* A template file `enigma_test.ml` with some OUnit stubs.
* A template file `a1.txt` for submitting your written feedback.
* Some additional scripts for compiling your code.

**Grading issues.**

* **Late submissions:** Carefully review the course policies on [submission and
  late assignments][late].  Verify before the deadline on CMS that you have
  submitted the correct version.
* **Environment, names, and types:** Your solution must compile and run under
  OCaml 4.05.0 with OUnit 2.0.0. You are required to adhere to the names and
  types of the functions specified below. Your solution must pass the `make
  check` described below in Part 0.  Otherwise, your solution will receive
  minimal credit.
* **Code style:** Refer to the [CS 3110 style guide][style]. Ugly code that is
  functionally correct will nonetheless be penalized. Take extra time to think
  and find elegant solutions.

**Prohibited OCaml features:**
You may not use imperative data structures, including refs, arrays, mutable
fields, and the `Bytes` and `Hashtbl` modules. (We haven't covered any of those,
so it's okay if they sound mysterious.) Strings are allowed, but the deprecated
functions on them in the [`String` module][string] are not, because those
functions provide imperative features. Your solutions may not require linking
any additional libraries/packages beyond the standard library, which is linked
by default.

[string]: https://caml.inria.fr/pub/docs/manual-ocaml/libref/String.html

## Part 0:  Sanity check <a name="check"></a>

Download the release code.  Run `make check` from the folder containing the
release code.  You should get output that includes the following (other lines
might be interspersed):
```
===========================================================
Your OCaml environment looks good to me.  Congratulations!
===========================================================
Your function names and types look good to me.
Congratulations!
===========================================================
```
If instead you get warnings about your OCaml environment, seek help immediately
from a consultant:  something is wrong with the version of OCaml you have
installed.

Assuming your environment is good, you should not get warnings about about your
function names and types at this point, because you haven't modified the
provided file `enigma.ml` yet.  As you complete the assignment, you will be
modifying that file.  As you solve each part of the assignment, you should
re-run `make check` to ensure that you haven't changed the names or types of any
of the required functions.

Running simply `make` will build your OUnit test suite in `enigma_test.ml` and
run it.  The test cases in the suite we ship to you will currently fail, because
you haven't yet implemented the functions that they test.  Running `make clean`
will delete generated files produced by the make and build systems.

## Part 1:  Index <a name="index"></a>

Write a function `index` that gives the 0-based index of a letter in the
alphabet.  Below is a specification comment for the function.  The line below
the comment is an OCaml *value specification* giving the name and type of a
value; OCaml would respond with that specification if you entered your function
at the toplevel.
```
(* [index c] is the 0-based index of [c] in the alphabet.
 * requires: [c] is an uppercase letter in A..Z *)
val index : char -> int
```

Your solution should need only about one line of code and should make
use of the [`Char` module][char]. We provide two OUnit unit tests for your
function, one to check that the index of 'A' is 0, and another to check
that the index of 'Z' is 25. Compile and run your test suite with `make`.
Make sure that both the tests pass. Re-run `make check` to double-check
that you haven't changed the names and types of any required functions;
continue to do that after solving each part.

[char]: http://caml.inria.fr/pub/docs/manual-ocaml/libref/Char.html

## Part 2:  Rotors <a name="rotors"></a>

Our goal in this part is to implement functions that model how current
passes through rotors.  This part gets rather tricky, so read the
writeup very closely!

**Rotor wiring.**
Here is a picture of a disassembled rotor showing some of the internal
wiring:

<img style="display: block; margin-left: auto; margin-right: auto; max-width:50%"
  src="internal.jpg"/>

<p style="text-align:center;">image source:
  https://www.flickr.com/photos/duanewessels/7063367271</p>

Rather than using complicated diagrams, the wiring of a rotor is often
specified in literature on the Enigma as a 26-character string that is a
permutation of the upper-case letters of the alphabet.  For example,
`EKMFLGDQVZNTOWYHXUSPAIBRCJ` is a wiring specification.  The specification has
nothing to do with the order in which letters are printed on the exterior of the
rotor.  (Indeed, you will note that all the pictures above show the letters in
alphabetical order.) Rather, the specification describes the interior wiring.

To understand a wiring specification, first forget about the letters.  They are
really just shorthand for a number&mdash;specifically, the index of that letter
in the alphabet.  So wiring specification `EKMFLGDQVZNTOWYHXUSPAIBRCJ` is better
understood as the following list of numbers:  4, 10, 12, 5, 11, ...  (Now
you know why we had you implement Part 1.)

Next, let's associate each of those numbers with its index in the list:
```
  0     4
  1    10
  2    12
  3     5
  4    11
...   ...
```
Finally, think of that table as defining a function (w), such that (w(0) =
4), (w(1) = 10), (w(2) = 12), (w(3) = 5), (w(4) = 11), etc.

The function (w) we have thus constructed is what the wiring specification
really denotes. A rotor has a left side and a right side, and each side has 26
*contacts* from which electrical current can enter or leave the rotor. If
current enters the rotor from the right side at contact (i), it leaves the
rotor from the left side at contact (w(i)).  And if current enters the rotor
from the left side at contact (j), it leaves the rotor from the right side
at contact (w^{-1}(j)), where (w^{-1}) denotes the inverse of (w).

**Rotor orientation.**
There are 26 fixed *positions* at which current can enter a rotor. (For example,
if the plugboard is empty, typing 'A' causes current to enter the right-most
rotor at position 0, 'B' at 1, etc.) But current entering at position (i)
does not necessarily connect with contact (i) on the rotor, because rotors
turn around the spindle.  That rotation causes an *offset* between the positions
and the contacts.

Recall that the operator can see on the top of the machine, for each rotor, one
of the letters that appears as exterior label on the rotor.  Also recall that
those letters are printed on the rotor in alphabetical order.  So the offset
of the rotor can be can be determined by looking at the letter that is showing.
We call that the *top letter* of the rotor.  Here, for example, is a closeup
of a three-rotor Enigma machine in which the top letters are currently RDK:

<img style="display: block; margin-left: auto; margin-right: auto; max-width:25%"
  src="topletters.jpg"/>

<p style="text-align:center;">image source:
  https://de.wikipedia.org/wiki/Bomba</p>

When a rotor is in its *default* orientation, with 'A' showing as the top
letter, the offset of the rotor is 0:  current entering at
position 0 touches contact 0, position 1 touches contact 1, and so forth.  But
when 'B' is showing as the top letter, the offset of the rotor is 1:  current
entering from at position 0 touches contact 1, position 1 touches
contact 2, and so forth.

Likewise, as current exits from a contact on the opposite side of the rotor,
that contact is offset from the actual exit position.  For example, when
the offset is 1, current flowing to contact 1 exits at position 0,
contact 2 exits at position 1, and so forth.

Note that the way the offset works is thus independent of the direction that
current is flowing&mdash;which makes sense, as the offset is determined
by the movement of the rotor around the spindle, not by the flow of current
through the rotor.

**Examples.**
Now that we know about how wiring and orientation affect how current flows
through a rotor, let's look at some examples.   Here is where a 3D model is
worth a thousand words:  the Pringles-can model linked above makes these
examples considerably easier to follow.

*Example 1.* Suppose the rotor has wiring specification
`EKMFLGDQVZNTOWYHXUSPAIBRCJ`, its top letter is 'A', current is flowing from
right to left, and current enters at position 0:

* Current entering at right-hand position 0 flows to right-hand contact 0 of the
  rotor, because the rotor's offset is 0.

* Current entering right-hand contact 0 flows to left-hand contact 4,
  because (w(0) = 4) for this rotor's wiring specification.

* Current exiting left-hand contact 4 flows to left-hand position 4, because the
  rotor's offset is 0.

*Example 2.* But now suppose that the rotor's top letter is 'B'.  How would
Example 1 change?

* Current entering at right-hand position 0 flows to contact 1 of the rotor,
  because the rotor's offset is now 1.

* Current entering right-hand contact 1 flows to left-hand contact 10,
  because (w(1) = 10).

* Current exiting left-hand contact 10 flows to left-hand position 9, because
  the rotor's offset is 1.

*Example 3.* Now let's redo Example 1, but with current flowing from left to
right.

* Current entering at left-hand position 0 flows to left-hand contact 0 of the
  rotor, because the rotor's offset is 0.

* Current entering left-hand contact 0 flows to right-hand contact 20,
  because (w^{-1}(0) = 20) for this rotor's wiring specification.
  (That's a fact we haven't previously established, but you should be able
  to work it out yourself from the fact that 'A', whose index in the alphabet
  is 0, appears at position 20 in the wiring specification.)

* Current exiting right-hand contact 20 flows to right-hand position 20, because
  the rotor's offset is 0.

*Example 4.* Finally let's redo Example 2, but with current flowing from
left to right.

* Current entering at left-hand position 0 flows to left-hand contact 1 of the
  rotor, because the rotor's offset is 1.

* Current entering left-hand contact 1 flows to right-hand contact 22,
  because (w^{-1}(1) = 22) for this rotor's wiring specification.

* Current exiting right-hand contact 22 flows to right-hand position 21, because
  the rotor's offset is 1.

**Functions to implement.**
Write a function `map_r_to_l` that computes how a rotor maps current
from right to left, and a corresponding function `map_l_to_r` for left
to right.  Note that these functions do not perform any stepping
of the rotor; rather, they model how current passes through the rotor
when it's in a particular orientation.  We'll implement stepping later.

We define a *valid* wiring specification as one that is a 26-character
upper-case string and that is a permutation of the letters of the alphabet.
Using that definition, here are specifications for the two functions:
```
(* [map_r_to_l wiring top_letter input_pos] is the left-hand output position
 * at which current would appear when current enters at right-hand input
 * position [input_pos] to a rotor whose wiring specification is given by
 * [wiring].  The orientation of the rotor is given by [top_letter],
 * which is the top letter appearing to the operator in the rotor's
 * present orientation.
 * requires:
 *  - [wiring] is a valid wiring specification.
 *  - [top_letter] is in 'A'..'Z'
 *  - [input_pos] is in 0..25
 *)
val map_r_to_l : string -> char -> int -> int

(* [map_l_to_r] computes the same function as [map_r_to_l], except
 * for current flowing left to right. *)
val map_l_to_r : string -> char -> int -> int
```

Your solution should need only a few lines of code per function and should make
use of the [`String` module][string] (but, as usual, not any of the deprecated
functions in it, because those are imperative).  You will likely discover that
you want to implement a modulo operator that differs from the built-in operator
`mod`, and that you want to implement a function that is the inverse of `index`,
which you wrote in Part I.

**Unit tests to implement.**
Write OUnit unit tests for your `map_r_to_l` and `map_l_to_r` functions. Add
those unit tests (and all others you create in future Parts of this assignment)
to `enigma_test.ml`.  Here are some ideas to get you started, a couple of which
we've already included in the release version of `enigma_test.ml`:

* The wiring specification `ABCDEFGHIJKLMNOPQRSTUVWXYZ` with top letter 'A'
  should be the identity map, and other top letters should have no
  effect on the map.  You could even test this before you've written
  any code at all to account for the orientation.

* Examples 1 and 3 above would be good to test next, since they involve
  a more complicated wiring specification but not the orientation.

* The wiring specification `BACDEFGHIJKLMNOPQRSTUVWXYZ` with top letters 'A',
  'B', and 'C' would be useful to test next, to see whether your code
  for handling the orientation is correct.

* Then you could test Examples 2 and 4 above.

After that, you could try some other wiring specifications.  For example,
here are the wiring specifications for the three standard rotors used by the
1930 Enigma I:
```
rotor I:   EKMFLGDQVZNTOWYHXUSPAIBRCJ
rotor II:  AJDKSIRUXBLHWTMCQGZNPYFVOE
rotor III: BDFHJLCPRTXVZNYEIWGAKMUSQO
```
The Roman numerals above are the *names* of the rotors, not the order
they are installed on the spindle.  Indeed, they could be installed in any
order.  You'll recognize rotor I as the rotor we have used in our examples
above.  As test cases with those, you could try the following:

* If `rotor_III = "BDFHJLCPRTXVZNYEIWGAKMUSQO"`, then
  `map_r_to_l rotor_III 'O' 14` should be `17`.

* If `rotor_I = "EKMFLGDQVZNTOWYHXUSPAIBRCJ"`, then
  `map_l_to_r rotor_I 'F' 10` should be `14`.

But we don't recommend trying those test cases until you have Examples 1 through
4 working. Please do not assume that rotors I, II and III are the only rotors.
We will test your solution with [historical rotors][wiring] and shiny new
rotors.

[wiring]: http://www.cryptomuseum.com/crypto/enigma/wiring.htm

## Interlude:  Test sweep <a name="sweep"></a>

We will put increasing emphasis on unit test suites as the semester goes along.
Recall that we did not require you to submit any test suite for A0.  By
the time we get to A3, we will do automated testing of your test suite itself
to see how capable it is of finding bugs.

But for this assignment, we'll be in the middle of those two extremes.  As you
know by now, you do need to submit an OUnit test suite.  For the sake of testing
the correctness of your own solution code, your test suite should at a minimum
contain all the test cases suggested in this writeup.  The graders, however,
will not specifically check for that.  Instead, we ask you to do the following.

For each function (other than `index`) that is required by this assignment, you
should identify four to six *representative test cases*.  Think of this as a
small, self-curated set of tests extracted from your larger set of tests for
each function.  We call this subset of your test suite a *test sweep*:  it's a
quick overview of some important parts of the full suite.

You should also write a comment in your test suite file explaining
why you chose the cases you included in your sweep.  That comment should
explain what makes those test cases interesting.  Here is a non-exhaustive,
generic list of reasons why you might deem a test case interesting:

* It illustrates a common, simple input or output.

* It explores a *boundary case*: an input or output that is extremal in
  some way.

* It captures an example provided by the writeup.

* It tests an especially tricky part of an algorithm.

* It causes a particular helper function or small, important piece
  of the primary function to be evaluated.

* ...

When you use any of those explanations as part of your comment, though, you
should make the explanation specific to the function, input, and/or output
that it involves.

In the released version of `enigma_test.ml`, we have provided a template
for how to present your test sweep.  The grader will review your test sweep and
its explanation as part of grading your assignment, but will not grade the rest
of your test suite.  You do not need to stress over picking the four to six
absolute best test cases for each function: the grader will not be looking
to see whether the cases you excluded are somehow better than the cases you
included.  Rather, what matters is that (i) you have several test cases,
(ii) each one is important, and (iii) there isn't signficant overlap between
the test cases in the sweep of each function.

Note that, as your test sweep is part of your solution, posting it on Piazza
is not allowed.  You are, of course, welcome to bring questions about your
test sweep to office hours.  Simply asking, "Is this sweep good/right?" is
not a profitable way to begin a discussion, though, much like asking,
"Is this code good/right?" is not.  Instead, try to begin by articulating what
your particular concern about it might be.

## Part 3: Reflector <a name="reflector"></a>

The wiring of a reflector is specified with the same kind of 26-character string
as a rotor, and that string is used to define a function (w) as before.
Unlike rotors, the reflector does not rotate, so we don't have to worry about
the orientation or offset of the reflector.  Also, current flows only one
direction through a reflector, so we don't have to worry about right-to-left vs.
left-to-right. When current enters a reflector at input position (i), it
exits at output position w(i).


A *valid reflector specification* is a valid wiring specification that meets
one additional criterion:  if (w(i) = j) then
(w(j) = i).  It is this property that makes the wiring a reflection, in
the sense that it swaps (i) and (j).


The wiring specifications of the standard reflectors B and C from the 1930
Enigma I are as follows:
```
reflector B: YRUHQSLDPXNGOKMIEBFZCWVJAT
reflector C: FVPJIAOYEDRZXWGCTKUQSBNMHL
```

For example, current entering reflector B at input position 0 would exit at
output position 24, because (w(0) = 24) for that reflector.

Write a function `map_refl` that computes how a reflector maps current.
Here is a specification for the function:
```
(* [map_refl wiring input_pos] is the output position at which current would
 * appear when current enters at input position [input_pos] to a reflector
 * whose wiring specification is given by [wiring].
 * requires:
 *  - [wiring] is a valid reflector specification.
 *  - [input_pos] is in 0..25
 *)
val map_refl : string -> int -> int
```

*Hint: you should be able to implement this function rather simply in terms of
functions you've already implemented.*

Write OUnit unit tests for your `map_refl` function.  Add those unit tests to
`enigma_test.ml`.  Please do not assume that reflectors B and C are the only
reflectors. We will test your solution with [historical reflectors][wiring] and
shiny new reflectors.

## Part 4:  Plugboard <a name="plugboard"></a>

The plugboard swaps pairs of letters before and after current passes through the
rest of the machine.  The operator configured the plugboard by inserting cables
between pairs of letters.  We'll represent a plugboard by a list of pairs of
characters. For example, here's a plugboard in which two cables have been
inserted; one between 'A' and 'Z', the other between 'X' and 'Y': `[('A','Z');
('X','Y')]`. The order of characters in the pairs should not matter and the
order of pairs in the list should not matter.  So that plugboard has an effect
that is equivalent to this one: `[('Y','X'); ('Z','A')]`. The plugboard with no
cables inserted would be represented by the empty list, `[]`. We define a
*valid* plugboard as one in which only the letters 'A'..'Z' appear, and no
letter appears more than once.  A valid plugboard therefore may have anywhere
from 0 to 13 elements in its list.

Write a function `map_plug` that computes how a plugboard maps letters.
Here is a specification for the function:
```
(* [map_plug plugs c] is the letter to which [c] is transformed
 * by the plugboard [plugs].
 * requires:
 *  - [plugs] is a valid plugboard
 *  - [c] is in 'A'..'Z'
 *)
val map_plug : (char * char) list -> char -> char
```

*Hint: write a recursive function that pattern-matches against the input list.*

Write OUnit unit tests for your `map_plug` function.  Add those
unit tests to `enigma_test.ml`.  The empty plugboard and
a plugboard that has 13 cables inserted into it would be good bases for test cases.

## Part 5:  Ciphering a character <a name="cipherchar"></a>

Now it's time to assemble all the functions you've written and tested. We'll
first implement the ciphering of just a single character based on the current
configuration of the Enigma machine, including the plugboard, the reflector, and
the rotors and their orientations. But for now, we continue to leave out any
stepping of the rotors.

**Input and output.**
The operator types a letter *X* as input.  That letter is potentially
transformed to a different letter *Y* because of the plugboard.  The result of
that transformation becomes an electrical current at position *p*, where *p* is
the index of *Y*.  That current passes through the rotors, the reflector, and
back through the rotors.  At each of those, the output position from one rotor
becomes the input position to the next.  Eventually current exits the rotors at
position *q*. Let *Z* be the letter whose index is *q*.  The plugboard
potentially transforms *Z* into an output letter *W*, which is the letter that
lights up on the lampboard.

The design of Enigma limits it to the upper-case letters A..Z. Lower-case
letters, numerals, punctuation, spaces, etc. are not supported.

**An example.**
Consider the following configuration: reflector B is installed; rotors I, II,
and III are installed in that order from left to right; the top letter of
every rotor is 'A'; and the plugboard is empty.  Then input letter 'G' would be
enciphered to output letter 'P', as follows:

* The input letter is 'G'.  The plugboard is empty, so no transformation
  occurs.  Current begins flowing into the rotors at the position of the index
  of 'G', which is 6.

* Rotor III with top letter 'A' maps current from right-hand position
  6 to left-hand position 2.

* Rotor II with top letter 'A' maps current from right-hand position 2
  to left-hand position 3.

* Rotor I with top letter 'A' maps current from right-hand position 3
  to left-hand position 5.

* Reflector B maps current from position 5 to position 18.

* Rotor I with top letter 'A' maps current from left-hand position 18
  to right-hand position 18.

* Rotor II with top letter 'A' maps current from left-hand position 18
  to right-hand position 4.

* Rotor III with top letter 'A' maps current from left-hand position 4
  to right-hand position 15.

* Current flows out of the rotors at position 15, which is the index
  of 'P'.  The plugboard is empty, so no transformation occurs.
  The output letter is 'P'.

Expanding on that example, the following table shows how any individual
character would be enciphered in that machine configuration.  You'll see, for
example, that input letter 'G' lines up with output letter 'P', as we just
explained in detail:
```
input letter:  ABCDEFGHIJKLMNOPQRSTUVWXYZ
                     |
                     V
output letter: UEJOBTPZWCNSRKDGVMLFAQIYXH
```

**Machine configuration.**
Let's introduce the following type definitions to represent the configuration of
a machine:
```
type rotor = {
	wiring : string;
	turnover : char;
}

type oriented_rotor = {
	rotor : rotor;
	top_letter : char;
}

type config = {
	refl : string;
	rotors : oriented_rotor list;
	plugboard : (char*char) list;
}
```

The order of elements in the `rotors` list in the `config` type is the same as
the order in which the rotors are installed on the spindle, from left to right.
For example, if there are two rotors installed, and if the left-most rotor on
the spindle (i.e., the rotor next to the reflector) is `r`, and if the
right-most rotor on the spindle (i.e., the rotor next to the static wheel) is
`q`, then that configuration would be represented by the list `[r; q]`,
assuming that `r` and `q` have been appropriately defined.

We define a *valid* configuration as one in which all the wiring specifications
are valid, the plugboard is valid, and the turnovers and top letters are all in
'A'..'Z'.  Note that there might be any number of rotors in a valid
configuration: perhaps 0, perhaps the standard 3, perhaps more.

The `turnover` field of the `rotor` type won't actually be used in any
meaningful way in this part of the assignment.  We'll use it when we get
to the next part, in which we implement stepping.

**Function to implement.**
Write a function `cipher_char` that computes how the Enigma machine
ciphers a single letter.  Note that this function still does not perform
any stepping of the rotors; rather, it models how the machine transforms
an input letter into an output letter when all the rotors are in a
particular orientation.

Here is a specification for the function:
```
(* [cipher_char config c] is the letter to which the Enigma machine
 * ciphers input [c] when it is in configuration [config].
 * requires:
 *  - [config] is a valid configuration
 *  - [c] is in 'A'..'Z'
 *)
val cipher_char : config -> char -> char
```

*Hint: consider writing two helper functions,*
```
map_rotors_r_to_l : oriented_rotor list -> int -> int
map_rotors_l_to_r : oriented_rotor list -> int -> int
```
*which compute the output position that results when current enters and passes
through an entire list of rotors.  They will be recursive functions over the
list.  Can you factor out a common helper function from them so that there isn't
any duplicated code?*

**Unit tests to implement.**
Write OUnit unit tests for your `cipher_char` function.  Add those
unit tests to `enigma_test.ml`.  Here are some ideas:

* When the plugboard is empty, the list of rotors is empty, and the reflector
  is wired as `ABCDEFGHIJKLMNOPQRSTUVWXYZ`, then `cipher_char` will behave
  like the identity function:  'A' ciphers to 'A, 'B' to 'B', etc.

* The example worked above gives you many possible test cases.

* Using a simulator, or simply pencil and paper, you could work out many more
  test cases yourself.

**Caution:** A potential point of confusion arises if you are trying to use
an online simulator to construct test cases for your `cipher_char`.  Those
simulators will implement stepping, which you will implement in the next part
of the assignment.  As you'll learn below, the rotors step just before
ciphering a character, which means their top letters change.  So you might
have to use different top letters to get the results you expect with those
simulators.  For the details, you'll have to keep reading.

## Part 6:  Stepping <a name="stepping"></a>

The intuition of *stepping* is that the rotors behave mostly like the numbers on
an odometer:  the rightmost completes a full revolution, and in so doing, causes
the next rotor to its left to take a single step, and so forth. But the stepping
system on the Enigma is a little more complicated.  Recall that each rotor has a
*turnover*, which is the top letter that is showing when a notch is reached in
the rotor.  Those notches cause the rotors to step according to the following
rules:

*Rule 1:* The rightmost rotor always takes a single step just before enciphering
each letter.

*Rule 2:* Just before every letter is enciphered, if the top letter of any rotor
*except the leftmost* is its turnover, then that rotor and the rotor to its left
step.

*Rule 3:* No rotor steps more than once per letter enciphered, even if the above
rules could be construed as suggesting that a rotor would step twice.

Some Enigma machines or rotors were built with different stepping rules; in case
of discrepancy, the rules above are in effect for this assignment.

Here are the turnovers of some of the historical rotors:
```
turnover of I:   Q
turnover of II:  E
turnover of III: V
```

**Example 1.**  Suppose the installed rotors are III-II-I (as always, leftmost
to rightmost), starting in orientations where the top letters are KDO.  Then as
each character is enciphered, they would step as follows:  KDO, KDP, KDQ, KER,
LFS, LFT, LFU, ...

**Example 2.** Suppose the installed rotors are III-II-I, starting with top
letters VDP.  Then as each character is enciphered, they would step as follows:
VDP, VDQ, VER, WFS, WFT, ...

Write a function `step` that computes how the Enigma machine steps just
before enciphering a letter.  Here is a specification for the function:
```
(* [step config] is the new configuration to which the Enigma machine
 * transitions when it steps beginning in configuration [config].
 * requires: [config] is a valid configuration
 *)
val step : config -> config
```
*Observation: the type of this function is a hallmark of the functional style of
programming. Instead of mutating state, we produce a new value out of an old
one.*

*Hint: This is the most algorithmically complex part of the assignment.  If
you're having difficulty, implement only Rule 1, then move on to the rest of the
assignment. Come back and implement Rules 2 and 3 later.  Actually, that might
be a good strategy in any case.*

Write OUnit unit tests for your `step` function.  Add those unit tests to
`enigma_test.ml`.  The two examples given above would be a good source of tests.

## Part 7: Ciphering a string <a name="cipherstring"></a>

At last, it's time to encipher an entire string.  The order of operations as
they occur on the Enigma machine is as follows:

* The operator types a letter.

* The machine immediately steps to a new configuration.

* The letter is enciphered according to that new configuration.

* The enciphered letter is lit on the lampboard.

* (The operator repeats until finished.)

Write a function `cipher` that computes how the Enigma machine ciphers a string.
Here is a specification for the function:
```
(* [cipher config s] is the string to which [s] enciphers
 * when the Enigma machine begins in configuration [config].
 * requires:
 *   - [config] is a valid configuration
 *   - [s] contains only uppercase letters
 *)
val cipher : config -> string -> string
```

Write OUnit unit tests for your `cipher` function.  Add those unit tests to
`enigma_test.ml`.

Here is one case for you: use reflector B, rotors I, II, and III (from left to
right, as always), with starting top letters F, U, and N, and a plugboard with a
single cable connecting 'A' and 'Z'.  Encipher YNGXQ; it should produce the name
of a pretty good programming language.

## What to turn in <a name="turnin"></a>

Submit files with these names on [CMS][cms]:

* `enigma.ml`, containing your solution code.
* `enigma_test.ml`, containing your OUnit test suite.
* `a1.txt`, containing your written feedback.

**REMINDER:** ensure that your solution passes the `make check` test described
in Part 0; if it does not, your solution will receive minimal credit.

[cms]: https://cms.csuglab.cornell.edu/

## Assessment <a name="assessment"></a>

The course staff will assess four aspects of your solution, as discussed below.
The percentages shown are the approximate weight we expect each to receive.

* **Correctness.** [60%] Does your solution compile against and correctly pass
  the course staff's test cases?  This is assessed entirely automatically by an
  autograder, which is why it is essential for your solution to pass `make
  check`.  Each test case you pass gets you more points. So to maximize your
  partial credit, do your best to submit code that will pass as many test cases
  as possible.  As a rough guide, here is a further breakdown of the approximate
  weight we expect to assign to each part of the implementation: part 1: 5%;
  parts 2 and 3 combined: 25%; part 4: 10%; part 5: 10%; part 6: 25%; part 7:
  25%.

* **Testing.** [15-20%] Does your test sweep satisfy the criteria discussed above?

* **Code quality.** [15-20%] How readable and elegant is your source code?  How
  readable and informative are your specification comments?  All functions,
  including helper functions, must have specification comments.

* **Feedback.** [5%] Did you complete the written feedback document?

## Karma <a name="karma"></a>

*Karma problems* are a way of going above and beyond the call of 3110 duty.
Solving them is not required, nor does it earn any explicit extra credit. But
the course staff will keep track of the karma problems that you solve
throughout the semester.  The number of orange camels, below, is an indicator
of how hard the karma problem might be.

Be mindful that you still need to provide the required functions, above, so make
sure your implementation of any karma functionality does not violate the
specifications/names/types of those functions.

<img src="camel_orange.png" height="40" width="40">
<img src="camel_black.png" height="40" width="40">
<img src="camel_black.png" height="40" width="40">
-
<img src="camel_orange.png" height="40" width="40">
<img src="camel_orange.png" height="40" width="40">
<img src="camel_black.png" height="40" width="40">
**Codebreaker:**

Implement a tool that breaks Enigma encryptions without having the key.
A brute force attack against a three-rotor machine is probably not that hard.
Doing it for a five-rotor machine with plugboard, however, could be an interesting
challenge.

<img src="camel_orange.png" height="40" width="40">
<img src="camel_orange.png" height="40" width="40">
<img src="camel_orange.png" height="40" width="40">
**Sims:**

Implement a virtual simulator using an [OCaml text interface library][text]
in which the user can install rotors, adjust the plugboard, encrypt strings, etc.

[text]: http://caml.inria.fr/cgi-bin/hump.en.cgi?sort=0&browse=64

* * *

**Acknowledgement:** Adapted from Prof. David Reed (Creighton University).
