package strand

type RNA string

var dnaToRnaNucleotideMap = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	rna := ""
	for _, r := range dna {
		rna += string(dnaToRnaNucleotideMap[r])
	}
	return rna
}
