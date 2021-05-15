# Blind decryption using hashed ElGamal

WARNING: THIS IS A PROOF OF CONCEPT AND SHOULD NOT BE USED IN PRACTICE

Use-case: Corona-certificates that prevents tracing of users and sharing of medical data.

Players:
- Healt authority H signing a certificate
- User U that get and show their certificate
- Verifier V that verify a user's certificate

We use a blind decryption protocol with hashed ElGamal based on a paper by [Sakurai and Yamane](https://link.springer.com/chapter/10.1007/3-540-61996-8_45).

Step 1:
- H encrypts U's ID using **Encrypt** to produce a ciphertext C
- The public key used is wrt the exipiration date of the certificate
- H signs the ciphertext C creating a signature S (not included)
- H creates a certificate for user ID containing C and S

Step 2:
- U receives a certificate from H and show it to V

Step 3:
- V receive a certificate from U containing C and S
- V verifies thas S is correct (not included)
- V applies **Blind** on C to get a blinded value R
- V sends R to H to get a blind decryption

Step 4:
- H receive R and applies **BlindDecrypt** with multiple keys to get D
- The secret keys used is wrt possible expiration date of a certificate
- H sends the set of D values back to V without knowing C or ID

Step 5:
- V receives the set of D values from H and applies **Unblind** to each D to get ID
- V verifies that one of the decrypted ID's is identical to U's physical ID
