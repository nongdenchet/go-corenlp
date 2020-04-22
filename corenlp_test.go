package corenlp

import (
	"log"
	"os"
	"testing"

	"github.com/nongdenchet/go-corenlp/connector"
)

func TestEntityMention(t *testing.T) {
	c := connector.NewLocalExec(nil)
	c.ClassPath = os.Getenv("CORE_NLP")
	c.JavaArgs = []string{"-Xmx8g"}
	c.Annotators = []string{"tokenize", "ssplit", "pos", "lemma", "ner"}

	doc, err := Annotate(c, "Rich Lesser is CEO of Boston Consulting Group")
	if err != nil {
		log.Fatal(err)
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
