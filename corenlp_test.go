package corenlp

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/nongdenchet/go-corenlp/connector"
)

func TestCoreNLP(t *testing.T) {
	// sample text from https://stanfordnlp.github.io/CoreNLP/
	text := `President Xi Jinping of Chaina, on his first state visit to the United States, showed off his familiarity with American history and pop culture on Tuesday night.`

	// LocalExec connector is responsible to run Stanford CoreNLP process.
	c := connector.NewLocalExec(nil)
	c.JavaArgs = []string{"-Xmx4g"}     // set Java params
	c.ClassPath = os.Getenv("CORE_NLP") // set Java class path
	c.Annotators = []string{"tokenize", "ssplit", "pos", "lemma", "ner"}

	// Annotate text
	doc, err := Annotate(c, text)
	if err != nil {
		panic(err)
	}

	// Output words and pos
	fmt.Println("----- Tokens -----")
	for _, sentence := range doc.Sentences {
		for _, token := range sentence.Tokens {
			fmt.Printf("%s(%s)%s\n", token.Word, token.Pos, token.After)
		}
	}

	// Output entity mentions
	fmt.Println("\n----- Entity Mentions -----")
	for _, sentence := range doc.Sentences {
		for _, token := range sentence.EntityMentions {
			fmt.Printf("%s - %s\n", token.Text, token.Ner)
		}
	}
}

func TestEntityMention(t *testing.T) {
	c := connector.NewLocalExec(nil)
	c.ClassPath = os.Getenv("CORE_NLP")
	c.JavaArgs = []string{"-Xmx8g"}
	c.Annotators = []string{"tokenize", "ssplit", "pos", "lemma", "ner"}

	doc, err := Annotate(c, `Rich Lesser is CEO of Boston Consulting Group`)
	if err != nil {
		panic(err)
	}
	entityMentions := doc.Sentences[0].EntityMentions

	expectText := "Boston Consulting Group"
	if entityMentions[1].Text != expectText {
		t.Errorf("Expect get %s, but got %s", expectText, entityMentions[1].Text)
	}

	expectNer := "ORGANIZATION"
	if entityMentions[1].Ner != expectNer {
		t.Errorf("Expect get %s, but got %s", expectNer, entityMentions[1].Ner)
	}
}

func TestKBP(t *testing.T) {
	c := connector.NewHTTPClient(context.Background(), "http://localhost:9000")
	c.Annotators = []string{"tokenize", "ssplit", "pos", "lemma", "ner", "parse", "coref", "kbp"}

	doc, err := Annotate(c, `Rich Lesser is CEO of Boston Consulting Group`)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", doc.Sentences[0].RawParse)
	entityMentions := doc.Sentences[0].EntityMentions

	expectText := "Boston Consulting Group"
	if entityMentions[1].Text != expectText {
		t.Errorf("Expect get %s, but got %s", expectText, entityMentions[1].Text)
	}

	expectNer := "ORGANIZATION"
	if entityMentions[1].Ner != expectNer {
		t.Errorf("Expect get %s, but got %s", expectNer, entityMentions[1].Ner)
	}
}
