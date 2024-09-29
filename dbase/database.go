package dbase

import (
	"context"
	"errors"
	"nihility/logger"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoDBuri = os.Getenv("NYANTAN_URI")
var mongoDBdatabase = os.Getenv("NYANTAN_DATABASE_NAME")
var mongo_client *mongo.Client
var db *mongo.Database

var log = logger.Logger {
    Color: logger.Colors.Purple,
    Pretext: "database",
}

// =====================================================================================================================
// Basic connect and stuff

func check_envs(envs ...string) error {
    for i, e := range envs {
        // FIXME: Remove logs after secrets put in place and debug mode is off
        log.Printf("%02d: Loading DB env: %s=%s\n", i + 1, e, os.Getenv(e))
        if "" == os.Getenv(e) {
            msg := "Environmental variable does not exists: " + e
            log.Println(msg)
            return errors.New(msg)
        }
    }

    return nil
}

func Connect() error {
    var err error
    err = check_envs("NYANTAN_URI", "NYANTAN_DATABASE_NAME")
    if nil != err {
        return err
    }

    // Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoDBuri).SetServerAPIOptions(serverAPI)

    // Create a new client and connect to the server
    mongo_client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
        return err
	}
    db = mongo_client.Database(mongoDBdatabase)

	// Send a ping to confirm a successful connection
	var result bson.M
	if err := db.RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	log.Println("Pinged your deployment. You successfully connected to MongoDB!")

    return nil
}

// =====================================================================================================================
// Users and Auth

func (user *User) Register() {
    db.Collection("users").InsertOne(context.TODO(), user)
}

func (user *User) Find() error {
    filter := bson.D{{"username", user.Id}}
    err := db.Collection("users").FindOne(context.TODO(), filter).Decode(&user)
    if err != nil {
        log.Println(err)
        return err
    }
    return nil
}

func (user *User) Update() error {
    _, err := db.Collection("users").UpdateByID(context.TODO(), user._ID, user)
    if err != nil {
        log.Println(err)
        return err
    }
    return nil
}

func (user *User) Delete() error {
    filter := bson.D{{"_id", user._ID}}
    _, err := db.Collection("users").DeleteOne(context.TODO(), filter)
    if err != nil {
        log.Println(err)
        return err
    }
    return nil
}

func (user *User) Fandoms() []string {
    var results []string
    var tr_roles []TrRole

    filter := bson.D{{"username", user.Id}}
    cursor, err := db.Collection("tr_roles").Find(context.TODO(), filter)
    if err != nil {
        log.Println(err)
        return []string{}
    }
    cursor.All(context.TODO(), &tr_roles)
    for _, tr := range tr_roles {
        results = append(results, tr.Fandom)
    }
    log.Println(results)

    return results
}

// =====================================================================================================================
// Listings

func List_translations(fandoms []string) ([]Translation, error) {
    var translations []Translation
    var filter bson.D

    if 0 == len(fandoms) {
        return translations, errors.New("Cannot search without a fandom!")
    }

    for _, f := range fandoms {
        filter = append(filter, bson.E{"fandom", f})
    }

    cursor, err := db.Collection("translations").Find(context.TODO(), filter)
    if err != nil {
        log.Println(err)
        return translations, err
    }

    err = cursor.All(context.TODO(), &translations)
    if err != nil {
        return translations, err
    }
    log.Println(translations)

    return translations, nil
}

// =====================================================================================================================
// Selects

func Select_translation(id string) (Translation, error) {
    var translation Translation
    var filter primitive.D
    object_id, err := primitive.ObjectIDFromHex(id)
    if nil == err {
        filter = bson.D{{"_id", object_id}}
    } else {
        filter = bson.D{{"title", id}}
    }

    err = db.Collection("translations").FindOne(context.TODO(), filter).Decode(&translation)
    if err != nil {
        log.Println(err)
        return translation, err
    }

    return translation, nil
}

func Select_edit(id primitive.ObjectID, page int) ([]Edit_combined, error) {
    var edits []Edit
    var snippets []Edit_snippet
    var ec []Edit_combined

    filter := bson.D{{"translationid", id}, {"page", page}}
    cursor, err := db.Collection("edits").Find(context.TODO(), filter)
    if nil != err {
        return ec, err
    }
    err = cursor.All(context.TODO(), &edits)
    if nil != err {
        return ec, err
    }

    for _, edit := range edits {
        filter = bson.D{{"edit", edit.Id}}
        cursor, err = db.Collection("edit_snippets").Find(context.TODO(), filter)
        if nil != err {
            return ec, err
        }

        err = cursor.All(context.TODO(), &snippets)
        if nil != err {
            return ec, err
        }

        ec = append(ec, Edit_combined{
            Edit: edit,
            Snippets: snippets,
        })
    }

    return ec, nil
}
