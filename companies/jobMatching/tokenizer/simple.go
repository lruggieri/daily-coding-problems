package tokenizer

import "strings"

// simply splits input text into words
func GetTokenizerSimple() Tokenizer{
	return func(text string) Tokens{
		stringTokens :=  strings.Split(text," ")
		res := make(Tokens,len(stringTokens))
		for i := range stringTokens{
			res[i] = Token(stringTokens[i])
		}
		return res
	}
}
