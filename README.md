# Blind decryption using hashed ElGamal

Use-case: Corona-certificates that prevents tracing and sharing of medical data

Players:
- Healt authorities H signing certificates
- Users U that get and view their certificate
- Verifiers V that verify user's certificates

We use a blind decryption protocol with hashed ElGamal based on the work by [Sakurai and Yamane from 1996](https://link.springer.com/chapter/10.1007/3-540-61996-8_45).

The protocol goes as following:

- H encrypt a user's ID using Encrypt to produce a ciphertext C
- H signs the ciphertext C creating a signature S (not included)
- H creates a ceritificate for user ID inclduding C and S

- U receive certificate from H and show it to V

- V receive certificate from U including C and S
- V verifies thas S is correct
- V apply Blind on C to get a blinded value R
- V send R to H to get a blind decryption

- H receive R and apply BlindDecrypt to get D
- H send D back to V without knowing C or ID*

- V receive D from H and compute Unblind to get ID*
- V verifies that ID* is identical to user's ID
