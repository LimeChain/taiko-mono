package proposer

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/herumi/bls-eth-go-binary/bls"
	"golang.org/x/crypto/sha3"
)

func GetAddressFromBlsPublikKeyHex(blsPubKeyHex string) common.Address {
	if err := bls.Init(bls.BLS12_381); err != nil {
		log.Crit("Failed to initialize BLS library: %v", err)
	}

	blsPubKeyBytes, err := hex.DecodeString(blsPubKeyHex)
	if err != nil {
		log.Crit("Invalid hex string: %v", err)
	}

	var pubkey bls.PublicKey
	if err := pubkey.Deserialize(blsPubKeyBytes); err != nil {
		log.Crit("Failed to deserialize BLS public key: %v", err)
	}

	pubkeySerialized := pubkey.Serialize()

	hash := sha3.NewLegacyKeccak256()
	hash.Write(pubkeySerialized)
	keccakHash := hash.Sum(nil)

	address := keccakHash[len(keccakHash)-20:]

	return common.BytesToAddress(address)
}
