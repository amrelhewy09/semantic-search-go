package embeddings

import (
	"OPENAI-GO/embeddings/structs"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

func CreateNewEmbedding(text string, dbclient *structs.MySQL, openAIClient *structs.OpenAIClient) {
	embedding := openAIClient.GetEmbeddingForText(text)
	_, err := dbclient.InsertEmbedding(text, []byte(convertFloatToByte(embedding.Embedding)))
	if err != nil {
		fmt.Println("Error inserting embedding:", err)
	}
}

func GetRelatedEmbeddings(text string, dbclient *structs.MySQL, openAIClient *structs.OpenAIClient) {
	embedding := openAIClient.GetEmbeddingForText(text)
	result := dbclient.GetRelatedEmbeddings(convertFloatToByte(embedding.Embedding))

	fmt.Println("Search Results:")
	for _, headline := range result {
		fmt.Println(headline)
	}
	fmt.Println()

}

func convertFloatToByte(embedding []float32) []byte {
	buf := new(bytes.Buffer)

	for _, v := range embedding {
		err := binary.Write(buf, binary.LittleEndian, v)
		if err != nil {
			log.Fatalf("binary.Write failed: %v", err)
		}
	}

	return buf.Bytes()
}
