package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Block struct {
	Pos       int
	Timestamp string
	Data      BookCheckout
	Hash      string
	PrevHash  string
}

type BookCheckout struct {
	BookID       string `json:"book_id"`
	User         string `json:"user"`
	CheckoutDate string `json:"checkout_date"`
	IsGenesis    bool   `json:"is_genesis"`
}
type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishDate string `json:"publish_date"`
	ISBN        string `json:"isbn"`
}
type Blockchain struct {
	blocks []*Block
}

var BlockChain *Blockchain

func (b *Block) generateHash() {
	bytes, _ := json.Marshal(b.Data)

	// FIX: use strconv.Itoa instead of string(b.Pos)
	data := strconv.Itoa(b.Pos) + b.Timestamp + b.PrevHash + string(bytes)

	hash := sha256.Sum256([]byte(data))

	// FIX: assign final hash to block
	b.Hash = fmt.Sprintf("%x", hash)
}

func CreateBlock(prevBlock *Block, checkoutitem BookCheckout) *Block {
	block := &Block{}
	block.Pos = prevBlock.Pos + 1
	block.Timestamp = time.Now().String()
	block.Data = checkoutitem
	block.PrevHash = prevBlock.Hash
	block.generateHash()
	return block
}

func (bc *Blockchain) AddBlock(data BookCheckout) {
	prevBlock := bc.blocks[len(bc.blocks)-1]

	block := CreateBlock(prevBlock, data)

	if validBlock(block, prevBlock) {
		bc.blocks = append(bc.blocks, block)
	}
}

func validBlock(Block, prevBlock *Block) bool {
	if prevBlock.Pos+1 != Block.Pos {
		return false
	}
	if prevBlock.Hash != Block.PrevHash {
		return false
	}
	if !Block.validateHash(Block.Hash) {
		return false
	}
	return true

}

func (b *Block) validateHash(hash string) bool {
	b.generateHash()
	if b.Hash != hash {
		return false
	}
	return true
}
func writeBlock(w http.ResponseWriter, r *http.Request) {
	var checkoutitems BookCheckout
	if err := json.NewDecoder(r.Body).Decode(&checkoutitems); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not write block:%v", err)
		w.Write([]byte("could not write block"))
		return
	}
	BlockChain.AddBlock(checkoutitems)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("block added"))
}

func newBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not create:%v", err)
		w.Write([]byte("could not create new book"))
		return

	}

	h := md5.New()
	io.WriteString(h, book.ISBN+book.PublishDate)
	book.ID = fmt.Sprintf("%x", h.Sum(nil))

	resp, err := json.MarshalIndent(book, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not marshal payload:%v", err)
		w.Write([]byte("could not save book data"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}
func GenesisBlock() *Block {
	return CreateBlock(&Block{}, BookCheckout{IsGenesis: true})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}

func getBlockchain(w http.ResponseWriter, r *http.Request) {
	jbytes, err := json.MarshalIndent(BlockChain.blocks, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not get blockchain:%v", err)
		json.NewEncoder(w).Encode(err)
		return
	}

	io.WriteString(w, string(jbytes))
}
func main() {

	BlockChain = NewBlockchain()
	r := mux.NewRouter()
	r.HandleFunc("/", getBlockchain).Methods("GET")
	r.HandleFunc("/", writeBlock).Methods("POST")
	r.HandleFunc("/new", newBook).Methods("POST")

	go func() {
		for _, block := range BlockChain.blocks {
			fmt.Printf("Prev.hash:%s\n", block.PrevHash)
			bytes, _ := json.MarshalIndent(block.Data, "", " ")
			fmt.Printf("%v\n", string(bytes))
			fmt.Printf("Hash:%x\n", block.Hash)
			fmt.Println()

		}
	}()

	log.Println("Listening on Port 3000")

	log.Fatal("ListenAndServe: ", http.ListenAndServe(":3000", r))
}
