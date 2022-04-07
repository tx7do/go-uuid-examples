package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/chilts/sid"
	goSnowflake "github.com/godruoyi/go-snowflake"
	guuid "github.com/google/uuid"
	"github.com/kjk/betterguid"
	"github.com/lithammer/shortuuid"
	"github.com/oklog/ulid"
	pborman "github.com/pborman/uuid"
	"github.com/rs/xid"
	goUUID "github.com/satori/go.uuid"
	"github.com/segmentio/ksuid"
	"github.com/sony/sonyflake"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func genShortUUID() {
	id := shortuuid.New()
	fmt.Printf("github.com/lithammer/shortuuid: %s\n", id)
}

func genXid() {
	id := xid.New()
	fmt.Printf("github.com/rs/xid:              %s\n", id.String())
}

func genKsuid() {
	id := ksuid.New()
	fmt.Printf("github.com/segmentio/ksuid:     %s\n", id.String())
}

func genBetterGUID() {
	id := betterguid.New()
	fmt.Printf("github.com/kjk/betterguid:      %s\n", id)
}

func genUlid() {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	fmt.Printf("github.com/oklog/ulid:          %s\n", id.String())
}

func genSonyflake() {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		log.Fatalf("flake.NextID() failed with %s\n", err)
	}
	// Note: this is base16, could shorten by encoding as base62 string
	fmt.Printf("github.com/sony/sonyflake:      %d\n", id)
}

func genSnowflake() {
	nodeId := rand.Int63() % 1023
	node, err := snowflake.NewNode(nodeId)
	if err != nil {
		log.Fatalf("snowflake.NewNode() failed with %s\n", err)
	}

	id := node.Generate().Int64()
	fmt.Printf("github.com/bwmarrin/snowflake:      %d\n", id)
}

func genGoSnowflake() {
	id := goSnowflake.ID()
	fmt.Printf("github.com/godruoyi/go-snowflake:      %d\n", id)
}

func genSid() {
	id := sid.Id()
	fmt.Printf("github.com/chilts/sid:          %s\n", id)
}

func genUUIDv4() {
	id := goUUID.NewV4()
	fmt.Printf("github.com/satori/go.uuid:      %s\n", id)
}

func genUUID() {
	id := guuid.New()
	fmt.Printf("github.com/google/uuid:         %s\n", id.String())
}

func genPbormanUUID() {
	id := pborman.NewRandom()
	fmt.Printf("github.com/pborman/uuid:         %s\n", id.String())
}

func genMongoDBObjectID() {
	id := primitive.NewObjectID()
	fmt.Printf("go.mongodb.org/mongo-driver/bson/primitive:         %s\n", id.Hex())
}

func main() {
	genXid()
	genKsuid()
	genBetterGUID()
	genUlid()
	genSonyflake()
	genSnowflake()
	genGoSnowflake()
	genSid()
	genShortUUID()
	genUUIDv4()
	genUUID()
	genPbormanUUID()
	genMongoDBObjectID()
}
