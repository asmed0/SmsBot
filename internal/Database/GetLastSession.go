package Database

import (
	"github.com/getsentry/raven-go"
	"go.mongodb.org/mongo-driver/bson"
	"smsbot/internal/FiveSim"
)
func GetLastSession(discordID string) *FiveSim.FiveSimLastSession {
	dbsession := getDatabase(false, &DatabaseSession{})

	data := &FiveSim.UserData{}

	err := dbsession.collectionPtr.FindOne(nil, bson.M{"discord_id":discordID}).Decode(data)
	if err != nil{
		raven.CaptureErrorAndWait(err, nil)
	}

	lastSession := &FiveSim.FiveSimLastSession{
		ApiKey:      data.LastSession.ApiKey,
		ID:     data.LastSession.ID,
		Phone:      data.LastSession.Phone,
		Operator:     data.LastSession.Country,
		Product:      data.LastSession.Product,
		Price:     data.LastSession.Price,
		Status:      data.LastSession.Status,
		Expires:     data.LastSession.Country,
		CreatedAt:     data.LastSession.Country,
		Country:     data.LastSession.Country,
		IsDisposed:  data.LastSession.IsDisposed,
	}
	return lastSession
}
