import (
	"fmt"
	"go.mongo.org/mongo-driver/mongo"
)

func getCollection(Collection string) *mongo.Collection {
	urli := fmt.Sprintf("mongodb://%s:%s@%s:%d", usr, pwd, host, port)
}