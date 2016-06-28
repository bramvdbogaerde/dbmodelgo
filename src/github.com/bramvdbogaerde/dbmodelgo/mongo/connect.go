package mongo

func Open(uri string) MongoDatabase {
	mDB := MongoDatabase{}
	mDB.Open(uri)
	return mDB
}
