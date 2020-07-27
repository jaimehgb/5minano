package crypto

type Crypto interface {
	// Sign will sign a message with the given private key.
	Sign([]byte, *KeyPair) (*Signature, error)

	// CheckSig will determine if a signature is valid given the message and the signer public key.
	CheckSig([]byte, *Signature, *PublicKey)

	// NewSeed will generate a pseudorandom seed which can be used for a nano wallet.
	NewSeed() (*Seed, error)

	// PublicSeed will return the public seed where to derive non-hardened public keys from,
	// without revealing the private keys of those public keys.
	//
	// Even though there is no risk of loss of fund if this seed gets compromised it's still
	// recommended to be kept secret as having access to it will allow whoever does to determine
	// which addresses are owned by you.
	PublicSeed(*Seed) (*PublicSeed, error)

	// DeriveKeyPair will derive the public and private keys for a given index from the seed.
	DeriveKeyPair(*Seed, int) (*KeyPair, error)

	// DerivePublicKey will derive only a public key from a given PublicSeed.
	DerivePublicKey(*PublicSeed, int) (*PublicKey, error)
}
