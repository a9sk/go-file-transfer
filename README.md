# Cadabra

Secure File Transfer in Golang (project used to learn Go, i have never used it before)

## Installation

To install Cadabra, follow this steps:

1. Clone this repository on your machine: 
   ```bash
   git clone https://github.com/a9sk/cadabra
   ```
2. Change directory into the one you just downloaded:
   ```bash
   cd cadabra
   ```
3. Run the setup.sh script:
   ```bash
   bash setup.sh
   ```

## Usage

    Usage: cdbr [options] ...

    Options:
            -mode (client or server)
            -host (default: localhost)
            -port

### Generate TLS certificates

To use the server option, it is necessary to generate [TLS certificates](https://www.openssl.org/docs/manmaster/man7/ossl-guide-tls-introduction.html).
This repository provides a script to generate and store them.

To generate the certificates run:
```bash
   bash certificates.sh
```

NOTE: It is recommended to change the certificates when changing the server, but also to change them from time to time.

#### Example of localhost server hosting command:
```bash
   cdbr -mode server -host localhost -port 2222
```

#### Example of localhost client connection command:
```bash
   cdbr -mode client -host localhost -port 2222
```

## Licence

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).

## Contacts

To report bugs, request new features, or ask questions, contact the project author:

- Email: 920a9sk42f76c765@proton.me
- GitHub: [@a9sk](https://github.com/a9sk)
