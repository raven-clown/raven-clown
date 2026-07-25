# Embedded & Hardware

How the hardware side of a project actually works.

## Arduino Programming

- An Arduino runs a single compiled C/C++ program in a loop — no operating system underneath, just a `setup()` function that runs once and a `loop()` function that runs continuously, reading inputs and driving outputs.
- Digital pins read or write a simple on/off signal; analog pins can read a continuous range of voltage, which matters for a sensor that reports something more granular than just "present or not."

## IR Sensors

- An IR object-detection sensor works by emitting infrared light and measuring what bounces back — when an object (like a vehicle) is in range, the reflected signal crosses a threshold and the sensor's output pin flips state.
- The signal needs debouncing in software — a raw sensor reading can flicker rapidly right at the detection threshold, so the code has to require a stable reading over a short window before treating it as a real detection, or it ends up reacting to noise.

## Circuit Prototyping

- Breadboarding lets a circuit be wired and rewired without soldering, which makes it the right first step for testing whether a design actually works before committing to a permanent build.
- The usual cycle: wire the circuit, write the minimum code needed to test one behavior at a time, confirm it, then add the next piece — rather than wiring the whole circuit and writing all the code before testing anything.
