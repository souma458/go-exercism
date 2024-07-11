package protein

import (
	"errors"
)

var (
	ErrStop        error = errors.New("stop")
	ErrInvalidBase error = errors.New("invalid codon")
)

var codonProteinMap = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func FromRNA(rna string) ([]string, error) {
	proteins := make([]string, 0)

	for i := 0; i < len(rna)/3; i++ {
		protein, codonError := FromCodon(string(rna[3*i : 3*i+3]))
		if codonError != nil {
			return handleError(proteins, codonError)
		}
		proteins = append(proteins, protein)
	}

	return proteins, nil
}

func FromCodon(codon string) (string, error) {
	protein := codonProteinMap[codon]
	if protein == "" {
		return "", ErrInvalidBase
	}
	if protein == "STOP" {
		return "", ErrStop
	}
	return codonProteinMap[codon], nil
}

func handleError(proteins []string, fromCodonError error) ([]string, error) {
	if fromCodonError == ErrStop && len(proteins) != 0 {
		return proteins, nil
	}
	return proteins, fromCodonError
}
