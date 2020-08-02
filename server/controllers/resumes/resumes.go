package resumes

import (
  "log"
  "strconv"
  "context"
  "net/http"
  "encoding/json"
  "go.mongodb.org/mongo-driver/bson"
  "github.com/trivektor/svelte-resume-builder/db"
  "go.mongodb.org/mongo-driver/bson/primitive"
  "github.com/gorilla/mux"
  . "github.com/gobeam/mongo-go-pagination"
)

const DATABASE_NAME = "svelte_resumes_builder"
const COLLECTION_NAME = "resumes"

type Section struct {
  Name string `json:"name"`
  Description string `json:"description"`
}

type Resume struct {
  ID primitive.ObjectID `bson:"_id" json:"id"`
  Name string `json:"name"`
  Description string `json:"description"`
  Sections []Section `json:"sections"`
}

func Index(w http.ResponseWriter, r *http.Request) {
  page := 1

  if pageQuery := r.URL.Query().Get("page"); pageQuery != "" {
    page, _ = strconv.Atoi(pageQuery)
  }

  var resumes []Resume
  collection := db.Client.Database(DATABASE_NAME).Collection(COLLECTION_NAME)
  paginatedData, err := New(collection).Limit(100).Page(int64(page)).Filter(bson.M{}).Find()

  if err != nil {
    log.Fatal(err)
  }

  for _, raw := range paginatedData.Data {
    var resume *Resume

    if marshalErr := bson.Unmarshal(raw, &resume); marshalErr == nil {
      resumes = append(resumes, *resume)
    }
  }

  response := make(map[string]interface{})
  response["resumes"] = resumes
  response["total_count"] = paginatedData.Pagination.Total
  response["page"] = paginatedData.Pagination.Page

  responseJson, jsonErr := json.Marshal(response)

  if jsonErr != nil {
    log.Fatal(jsonErr)
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(responseJson)
}

func Create(w http.ResponseWriter, r *http.Request) {
  var resume Resume

  json.NewDecoder(r.Body).Decode(&resume)

  collection := db.Client.Database(DATABASE_NAME).Collection(COLLECTION_NAME)

  res, err := collection.InsertOne(context.TODO(), bson.M{"name": resume.Name, "description": resume.Description})

  if err != nil {
    log.Fatal(err)
  }
  id := res.InsertedID.(primitive.ObjectID)

  m := make(map[string]string)
  m["id"] = id.Hex()
  result, _ := json.Marshal(m)

  w.Header().Set("Content-Type", "application/json")
  w.Write(result)
}

func Show(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  id := params["id"]
  var resume Resume
  docId, _ := primitive.ObjectIDFromHex(id)

  collection := db.Client.Database(DATABASE_NAME).Collection(COLLECTION_NAME)
  collection.FindOne(context.TODO(), bson.M{"_id": docId}).Decode(&resume)

  json, _ := json.Marshal(resume)

  w.Header().Set("Content-Type", "application/json")
  w.Write(json)
}

func Update(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  id := params["id"]
  var resume Resume

  json.NewDecoder(r.Body).Decode(&resume)

  docId, _ := primitive.ObjectIDFromHex(id)

  collection := db.Client.Database(DATABASE_NAME).Collection(COLLECTION_NAME)

  update := bson.M{
    "$set": bson.M{
      "name": resume.Name,
      "description": resume.Description,
      "sections": resume.Sections,
    },
  }

  collection.UpdateOne(context.TODO(), bson.M{"_id": docId}, update)
  json, _ := json.Marshal(map[string]string{"status": "ok"})

  w.Header().Set("Content-Type", "application/json")
  w.Write(json)
}

func Delete(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  id := params["id"]
  docId, _ := primitive.ObjectIDFromHex(id)
  collection := db.Client.Database(DATABASE_NAME).Collection(COLLECTION_NAME)
  _, err := collection.DeleteOne(context.TODO(), bson.M{"_id": docId})

  if err != nil {
    log.Fatal(err)
  }

  json, _ := json.Marshal(map[string]string{"status": "ok"})

  w.Header().Set("Content-Type", "application/json")
  w.Write(json)
}
