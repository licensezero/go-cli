package inventory

import "bytes"
import "encoding/hex"
import "encoding/json"
import "errors"
import "golang.org/x/crypto/ed25519"

// Verifiable describes a struct that can be serialized for signing.
type Verifiable interface {
	verifyLicensorSignature(keyHex string) error
	verifyAgentSignature(keyHex string) error
}

func (envelope Version1Envelope) verifyLicensorSignature(keyHex string) error {
	return verifyLicensorJSON(
		keyHex,
		envelope.LicensorSignature,
		envelope.Manifest,
	)
}

type version1AgentSignaturePackage struct {
	Manifest          Version1Manifest `json:"license"`
	LicensorSignature string           `json:"licensorSignature"`
}

func (envelope Version1Envelope) verifyAgentSignature(keyHex string) error {
	serialized, err := json.Marshal(version1AgentSignaturePackage{
		Manifest:          envelope.Manifest,
		LicensorSignature: envelope.LicensorSignature,
	})
	if err != nil {
		return errors.New("could not serialize manifest")
	}
	compacted := bytes.NewBuffer([]byte{})
	err = json.Compact(compacted, serialized)
	if err != nil {
		return errors.New("could not serialize agent signature packet")
	}
	err = checkManifestSignature(
		keyHex,
		envelope.AgentSignature,
		compacted.Bytes(),
		"agent",
	)
	if err != nil {
		return err
	}
	return nil
}

func (envelope Version2Envelope) verifyLicensorSignature(keyHex string) error {
	return verifyLicensorJSON(
		keyHex,
		envelope.LicensorSignature,
		envelope.Manifest,
	)
}

type version2AgentSignaturePackage struct {
	Manifest          Version2Manifest `json:"manifest"`
	LicensorSignature string           `json:"licensorSignature"`
}

func (envelope Version2Envelope) verifyAgentSignature(keyHex string) error {
	serialized, err := json.Marshal(version2AgentSignaturePackage{
		Manifest:          envelope.Manifest,
		LicensorSignature: envelope.LicensorSignature,
	})
	if err != nil {
		return errors.New("could not serialize manifest")
	}
	compacted := bytes.NewBuffer([]byte{})
	err = json.Compact(compacted, serialized)
	if err != nil {
		return errors.New("could not serialize agent signature packet")
	}
	err = checkManifestSignature(
		keyHex,
		envelope.AgentSignature,
		compacted.Bytes(),
		"agent",
	)
	if err != nil {
		return err
	}
	return nil
}

func verifyLicensorJSON(keyHex string, signatureHex string, manifest interface{}) error {
	serialized, err := json.Marshal(manifest)
	if err != nil {
		return errors.New("could not serialize manifest")
	}
	compacted := bytes.NewBuffer([]byte{})
	err = json.Compact(compacted, serialized)
	if err != nil {
		return errors.New("could not compact manifest")
	}
	err = checkManifestSignature(
		keyHex,
		signatureHex,
		compacted.Bytes(),
		"licensor",
	)
	return err
}

func checkManifestSignature(publicKey string, signature string, json []byte, source string) error {
	signatureBytes := make([]byte, hex.DecodedLen(len(signature)))
	_, err := hex.Decode(signatureBytes, []byte(signature))
	if err != nil {
		return errors.New("invalid " + source + "signature")
	}
	publicKeyBytes := make([]byte, hex.DecodedLen(len(publicKey)))
	_, err = hex.Decode(publicKeyBytes, []byte(publicKey))
	if err != nil {
		return errors.New("invalid " + source + " public key")
	}
	signatureValid := ed25519.Verify(
		publicKeyBytes,
		json,
		signatureBytes,
	)
	if !signatureValid {
		return errors.New("invalid " + source + " signature")
	}
	return nil
}
