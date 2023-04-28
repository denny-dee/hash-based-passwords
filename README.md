# Password Hash Generator

This is a Go command-line application that generates a password hash for a given set of inputs. The generated hash can be used as a password for apps and websites.

Inspired by https://github.com/gustavomondron/twik and https://github.com/alexandrepossebom/Twik.

## Features

- Generates a hash from the given arguments.
- The hash length can be 4-32 characters.
- The user can choose to use one of the following character sets for the hash:
  - Numbers only
  - Numbers and alphabet (lowercase only)
  - Numbers and alphabet (lowercase and uppercase)
  - Numbers, alphabet, and symbols that are available on any keyboard (i.e., !@#$%^&*():;',.)
- The hash is deterministic, i.e., similar input produces similar hashes.

## Installation

To install the application, you need to have Go installed on your system.

1. Clone this repository.
```bash
git clone https://github.com/denny-dee/hash-based-passwords.git
cd hash-based-passwords
```

2. Build the application using the following command:
```bash
    go build
```

## Usage

To use the application, run the following command:

```bash
./hash-based-passwords -private-key <private-key> -master-key <master-key> -site-name <site-name> -salt <salt> -length <length> -charset <charset>
```

where:

- `<private-key>` is the user's private key (may be a long string of any characters).
- `<master-key>` is the user's master key (a string of up to 32 characters).
- `<site-name>` is the name of the site or app for which the password hash is being generated (e.g., "github.com", "bank of georgia", "apple id", "Google", etc.).
- `<salt>` is an additional tag used as "salt" (e.g., update date, "persona" or "work" tags).
- `<length>` is the desired length of the hash (4, 8, 16, or 32 characters).
- `<charset>` is the character set to use for hash generation (num, lower, upper, alnum, or all).

The application will generate a hash of the desired length using the selected character set and print it to the console.

### Examples

Here are some examples of using the application:

```bash
./hash-based-passwords -private-key "my secret key" -master-key "my master key" -site-name "github.com" -salt "work" -length 12 -charset alnum
```
This will generate a 12-character hash using the alphanumeric character set.

```bash
./hash-based-passwords -private-key "my secret key" -master-key "my master key" -site-name "bank of georgia" -salt "2023-04-28" -length 16 -charset all
```

This will generate a 16-character hash using the character set that includes numbers, lowercase and uppercase letters, and symbols.

## License

This code is released under the MIT license.
