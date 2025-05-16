# regex_engine

## Description

I created `regex_engine` as a personal project to dive deeper into how regular expressions work under the hood. This project is not just a tool to match patterns; it's a way for me to understand the underlying concepts such as finite automata, state machines, and how regex engines implement pattern matching.

As I began learning about regular expressions, I realized there was a lot going on behind the scenes—like the use of state machines and the conversion of regex patterns into executable algorithms. So, I decided to build my own basic regex engine to solidify my understanding. 

This project has given me a hands-on way to experiment with different regex features and see how they behave in practice. The result is a simple, yet robust engine that implements the core functionalities of regular expressions while being easy to extend for further exploration.

While it may not yet match the feature set of well-known libraries like `PCRE` or `std::regex`, it's a great starting point for anyone wanting to explore the basics of regex processing and state machines.

## Features

This regex engine includes the following features:

- **Quantifiers:**
  - `*` (Zero or more)
  - `?` (Zero or one)
  - `+` (One or more)
  
- **Groups:**
  - `()` for grouping parts of the expression

- **Character Classes:**
  - `[]` for defining custom character sets/ranges
  
- **Escape Characters:**
  - `\` is supported as an escape character for special characters
  
- **Braces for Repetition:**
  - `{n}` matches exactly `n` occurrences
  - `{n,}` matches `n` or more occurrences
  - `{n,m}` matches between `n` and `m` occurrences
  
- **Literals:**
  - Supports defining ranges inside `[]`, e.g., `[a-z]`, `[0-9]`...

