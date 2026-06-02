# Spectra Demo App (CryptoShop)

This repository serves as a demonstration target for the Spectra Cryptographic Intelligence platform.

**DO NOT USE THIS CODE IN PRODUCTION.** 

This repository intentionally contains deprecated, broken, and quantum-vulnerable cryptography (MD5, SHA-1, RSA-2048, ECC P-256) designed to trigger high Quantum Risk Scores (QRS) during a Spectra scan.

## Usage

To test Spectra against this repository:

```bash
$ git clone https://github.com/HarshalPatel1972/spectra-demo-app.git
$ cd spectra-demo-app
$ spectra scan .
```
