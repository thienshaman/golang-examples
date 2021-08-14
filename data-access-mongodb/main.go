package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Student struct {
	Name  string
	Email string
	Rank  string
}

var client *mongo.Client

// var ctx context.Context

func main() {
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.TODO())
	fmt.Println("Connnected!")

	// find document by _id
	studentById, err := findStudentById("61161eaedce1823acfecc761")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(studentById)

	// find all document
	students, err := findAllStudent()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(students)

	// find document by rank
	studentByRank, err := findStudentByRank("C")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(studentByRank)

	// add new document
	newStudent := Student{"Le Huu Hieu", "hieu@test.com", "C"}
	idInserted, err := addStudent(newStudent)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(idInserted)

	// update document
	studentUpdate := Student{"Le Huu Hieu", "hieule@test.com", "B"}
	updateCount, err := updateStudent("611686543c3f235c407bc743", studentUpdate)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated %v Documents!\n", updateCount)

	// delete document
	deleteCount, err := deleteStudent("611686543c3f235c407bc743")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v Documents!\n", deleteCount)

}

func findStudentById(id string) (Student, error) {
	mydb := client.Database("mydb")
	studentCollection := mydb.Collection("student")
	var student Student
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return student, err
	}
	filter := bson.D{{"_id", objID}}
	err = studentCollection.FindOne(context.TODO(), filter).Decode(&student)
	if err != nil {
		return student, err
	}
	return student, nil
}

func findAllStudent() ([]Student, error) {
	mydb := client.Database("mydb")
	studentCollection := mydb.Collection("student")
	var students []Student
	cur, err := studentCollection.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		return students, err
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var item Student
		err := cur.Decode(&item)
		if err != nil {
			return students, err
		}
		students = append(students, item)

	}

	return students, nil
}

func findStudentByRank(rank string) ([]Student, error) {
	mydb := client.Database("mydb")
	studentCollection := mydb.Collection("student")
	var students []Student
	filter := bson.D{{"rank", rank}}
	cur, err := studentCollection.Find(context.TODO(), filter, options.Find())
	if err != nil {
		return students, err
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var item Student
		err := cur.Decode(&item)
		if err != nil {
			return students, err
		}
		students = append(students, item)

	}

	return students, nil
}

func addStudent(student Student) (string, error) {
	mydb := client.Database("mydb")
	studentCollection := mydb.Collection("student")

	result, err := studentCollection.InsertOne(context.TODO(), student)

	if err != nil {
		return "", err
	}

	idInserted := result.InsertedID.(primitive.ObjectID).String()

	return idInserted, nil
}

func updateStudent(id string, student Student) (int64, error) {
	mydb := client.Database("mydb")
	studentCollection := mydb.Collection("student")

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}
	filter := bson.D{{"_id", objId}}
	result, err := studentCollection.UpdateOne(context.TODO(), filter, bson.D{{"$set", student}})
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil

}

func deleteStudent(id string) (int64, error) {
	mydb := client.Database("mydb")
	studentCollection := mydb.Collection("student")

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}
	filter := bson.D{{"_id", objId}}
	result, err := studentCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}
