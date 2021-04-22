package generators

import (
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

type uuidBinaryGenerator struct {
	base
}

func newUUIDBinaryGenerator(base base) (Generator, error) {
	return &uuidBinaryGenerator{base: base}, nil
}

func (g *uuidBinaryGenerator) EncodeValue() {
	uuid, _ := uuid.NewV4()

	uuidBytes := uuid.Bytes()
	binLen := uint32(len(uuidBytes))

	g.buffer.Write(uint32Bytes(binLen))
	g.buffer.WriteSingleByte(bsontype.BinaryUUID)
	g.buffer.Write(uuidBytes)
}

func (g *uuidBinaryGenerator) EncodeValueAsString() {
	uuid, _ := uuid.NewV4()
	g.buffer.Write([]byte(uuid.String()))
}
