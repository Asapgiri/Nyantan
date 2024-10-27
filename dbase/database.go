package dbase

import (
	"context"
	"errors"
	"nyantan/logger"
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

var dbUSERS           *mongo.Collection
var dbTR_ROLES        *mongo.Collection
var dbTRANSLATIONS    *mongo.Collection
var dbEDITS           *mongo.Collection
var dbEDIT_SNIPPETS   *mongo.Collection

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

    dbUSERS =         db.Collection("users")
    dbTR_ROLES =      db.Collection("tr_roles")
    dbTRANSLATIONS =  db.Collection("translations")
    dbEDITS =         db.Collection("edits")
    dbEDIT_SNIPPETS = db.Collection("edit_snippets")

    return nil
}

// =====================================================================================================================
// Users and Auth

func (user *User) Register() {
    dbUSERS.InsertOne(context.TODO(), user)
}

func (user *User) Find() error {
    filter := bson.D{{"username", user.Id}}
    err := dbUSERS.FindOne(context.TODO(), filter).Decode(&user)
    if err != nil {
        log.Printf("Trying to find user %s, fails with: %s", user.Id, err)
        return err
    }
    return nil
}

func (user *User) Update() error {
    _, err := dbUSERS.UpdateByID(context.TODO(), user._ID, user)
    if err != nil {
        log.Println(err)
        return err
    }
    return nil
}

func (user *User) Delete() error {
    filter := bson.D{{"_id", user._ID}}
    _, err := dbUSERS.DeleteOne(context.TODO(), filter)
    if err != nil {
        log.Println(err)
        return err
    }
    return nil
}

func (user *User) Fandoms() []TrRole {
    var tr_roles []TrRole

    filter := bson.D{{"username", user.Id}}
    cursor, err := dbTR_ROLES.Find(context.TODO(), filter)
    if err != nil {
        log.Println(err)
        return []TrRole{}
    }
    cursor.All(context.TODO(), &tr_roles)
    log.Println(tr_roles)

    return tr_roles
}

// =====================================================================================================================
// Translations

func (Translation) List(fandoms []string) ([]Translation, error) {
    var translations []Translation
    var filter bson.D

    if 0 == len(fandoms) {
        return translations, errors.New("Cannot search without a fandom!")
    }

    for _, f := range fandoms {
        filter = append(filter, bson.E{"fandom", f})
    }

    cursor, err := dbTRANSLATIONS.Find(context.TODO(), filter)
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

func (tr *Translation) Select(id string) error {
    var filter primitive.D
    object_id, err := primitive.ObjectIDFromHex(id)
    if nil == err {
        filter = bson.D{{"_id", object_id}}
    } else {
        filter = bson.D{{"title", id}}
    }

    err = dbTRANSLATIONS.FindOne(context.TODO(), filter).Decode(tr)
    if err != nil {
        log.Println(err)
        return err
    }

    return nil
}

func (tr *Translation) Add() error {
    _, err := dbTRANSLATIONS.InsertOne(context.TODO(), tr)
    return err
}

func (tr *Translation) Update() error {
    _, err := dbTRANSLATIONS.UpdateByID(context.TODO(), tr.Id, tr)
    if err != nil {
        log.Println(err)
        return err
    }
    return nil

}

func (tr *Translation) Delete() error {
    filter := bson.D{{"_id", tr.Id}}
    _, err := dbTRANSLATIONS.DeleteOne(context.TODO(), filter)
    if err != nil {
        log.Println(err)
        return err
    }
    return nil
}

// =====================================================================================================================
// Edits

func Select_edit(id primitive.ObjectID, page int) ([]Edit_combined, error) {
    var edits []Edit
    var ec []Edit_combined

    filter := bson.D{{"translationid", id}, {"page", page}}
    cursor, err := dbEDITS.Find(context.TODO(), filter)
    if nil != err {
        return ec, err
    }
    err = cursor.All(context.TODO(), &edits)
    if nil != err {
        return ec, err
    }

    for _, edit := range edits {
        snip := Edit_snippet{}
        snippets, _ := snip.SelectAll(edit.Id)
        ec = append(ec, Edit_combined{
            Edit: edit,
            Snippets: snippets,
        })
    }

    return ec, nil
}

func (edit *Edit) Add() error {
    _, err := dbEDITS.InsertOne(context.TODO(), edit)
    return err
}

// =====================================================================================================================
// Snippets

func (snip *Edit_snippet) SelectAll(editId primitive.ObjectID) ([]Edit_snippet, error) {
    var snippets []Edit_snippet

    filter := bson.D{{"edit", editId}}
    cursor, err := dbEDIT_SNIPPETS.Find(context.TODO(), filter)
    if nil != err {
        return snippets, err
    }

    err = cursor.All(context.TODO(), &snippets)

    return snippets, err
}

func (snip *Edit_snippet) Select(id primitive.ObjectID) error {
    filter := bson.D{{"_id", id}}
    err := dbEDIT_SNIPPETS.FindOne(context.TODO(), filter).Decode(snip)
    return err
}

func (snip *Edit_snippet) Add() error {
    _, err := dbEDIT_SNIPPETS.InsertOne(context.TODO(), snip)
    return err
}

func (snip *Edit_snippet) Update() error {
    _, err := dbEDIT_SNIPPETS.UpdateByID(context.TODO(), snip.Id, snip)
    return err

}

func (snip *Edit_snippet) Delete() error {
    filter := bson.D{{"_id", snip.Id}}
    _, err := dbEDIT_SNIPPETS.DeleteOne(context.TODO(), filter)
    return err
}
