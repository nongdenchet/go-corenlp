package document

type EntityMentions []EntityMention

// EntityMention represents entity mentioned in the sentence https://stanfordnlp.github.io/CoreNLP/entitymentions.html
type EntityMention struct {
	DocTokenBegin int    `json:"docTokenBegin"`
	DocTokenEnd   int    `json:"docTokenEnd"`
	TokenBegin    int    `json:"tokenBegin"`
	TokenEnd      int    `json:"tokenEnd"`
	Text          string `json:"text"`
	Ner           string `json:"ner"`

	CharacterOffsetBegin int `json:"characterOffsetBegin"`
	CharacterOffsetEnd   int `json:"characterOffsetEnd"`
}
