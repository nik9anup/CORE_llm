package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/storage"
)

func main() {
	// Documentation: Lists all buckets in Google Cloud Storage.
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	buckets, err := client.Buckets(ctx, "your-project-id")
	if err != nil {
		log.Fatalf("Failed to list buckets: %v", err)
	}

	fmt.Println("Buckets:")
	for _, bucket := range buckets {
		fmt.Println(bucket.Name)
	}
}





package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/storage"
)

func main() {
	// Documentation: Uploads a file to Google Cloud Storage.
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	bucket := client.Bucket("your-bucket-name")
	object := bucket.Object("destination-file-name")
	file, err := os.Open("local-file-path")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	wc := object.NewWriter(ctx)
	if _, err = wc.WriteFrom(file); err != nil {
		log.Fatalf("Failed to write file to bucket: %v", err)
	}
	if err := wc.Close(); err != nil {
		log.Fatalf("Failed to close writer: %v", err)
	}

	fmt.Println("File uploaded successfully.")
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

func main() {
	// Documentation: Creates a new Pub/Sub topic.
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, "your-project-id")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	topic := client.Topic("your-topic-name")
	exists, err := topic.Exists(ctx)
	if err != nil {
		log.Fatalf("Failed to check if topic exists: %v", err)
	}
	if exists {
		fmt.Printf("Topic %s already exists.\n", topic.ID())
		return
	}

	_, err = client.CreateTopic(ctx, "your-topic-name")
	if err != nil {
		log.Fatalf("Failed to create topic: %v", err)
	}

	fmt.Println("Topic created successfully.")
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/storage"
)

func main() {
	// Documentation: Deletes a file from Google Cloud Storage.
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	bucket := client.Bucket("your-bucket-name")
	object := bucket.Object("file-to-delete")

	if err := object.Delete(ctx); err != nil {
		log.Fatalf("Failed to delete object: %v", err)
	}

	fmt.Println("File deleted successfully.")
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

func main() {
	// Documentation: Lists all subscriptions in Google Cloud Pub/Sub.
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, "your-project-id")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	subscriptions, err := client.Subscriptions(ctx)
	if err != nil {
		log.Fatalf("Failed to list subscriptions: %v", err)
	}

	fmt.Println("Subscriptions:")
	for _, sub := range subscriptions {
		fmt.Println(sub.ID())
	}
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

func main() {
	// Documentation: Publishes a message to a Google Cloud Pub/Sub topic.
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, "your-project-id")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	topic := client.Topic("your-topic-name")
	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte("Hello, Cloud Pub/Sub!"),
	})

	// Block until the message is published.
	id, err := result.Get(ctx)
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	fmt.Printf("Message published with ID: %s\n", id)
}





package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

func main() {
	// Documentation: Creates a new Compute Engine instance.
	ctx := context.Background()

	service, err := compute.NewService(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Compute Engine service: %v", err)
	}

	project := "your-project-id"
	zone := "your-zone"
	instanceName := "instance-name"
	image := "projects/debian-cloud/global/images/debian-10-buster-v20220315"

	instance := &compute.Instance{
		Name:        instanceName,
		MachineType: fmt.Sprintf("zones/%s/machineTypes/n1-standard-1", zone),
		Disks: []*compute.AttachedDisk{
			{
				AutoDelete: true,
				Boot:       true,
				Mode:       "READ_WRITE",
				InitializeParams: &compute.AttachedDiskInitializeParams{
					SourceImage: image,
				},
			},
		},
		NetworkInterfaces: []*compute.NetworkInterface{
			{
				AccessConfigs: []*compute.AccessConfig{
					{
						Type: "ONE_TO_ONE_NAT",
						Name: "External NAT",
					},
				},
				Network: "global/networks/default",
			},
		},
	}

	op, err := service.Instances.Insert(project, zone, instance).Do()
	if err != nil {
		log.Fatalf("Failed to create instance: %v", err)
	}

	fmt.Printf("Instance %s is being created.\n", instanceName)
	fmt.Printf("Operation ID: %s\n", op.Name)
}





package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/sqladmin/v1beta4"
)

func main() {
	// Documentation: Lists all Cloud SQL instances.
	ctx := context.Background()

	service, err := sqladmin.NewService(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Cloud SQL service: %v", err)
	}

	instances, err := service.Instances.List("your-project-id").Do()
	if err != nil {
		log.Fatalf("Failed to list instances: %v", err)
	}

	fmt.Println("Cloud SQL Instances:")
	for _, instance := range instances.Items {
		fmt.Printf("%s (%s)\n", instance.Name, instance.Region)
	}
}





package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/sqladmin/v1beta4"
	"google.golang.org/api/option"
)

func main() {
	// Documentation: Creates a new database in Cloud SQL.
	ctx := context.Background()

	service, err := sqladmin.NewService(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Cloud SQL service: %v", err)
	}

	instance := "your-instance-id"
	database := "new-database"

	op, err := service.Databases.Insert("your-project-id", instance, &sqladmin.Database{
		Name: database,
	}).Do()
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}

	fmt.Printf("Database %s created successfully.\n", database)
	fmt.Printf("Operation ID: %s\n", op.Name)
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/storage"
)

func main() {
	// Documentation: Lists all objects in a Google Cloud Storage bucket.
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	bucket := client.Bucket("your-bucket-name")
	objects := bucket.Objects(ctx, nil)

	fmt.Println("Objects in Bucket:")
	for {
		objAttrs, err := objects.Next()
		if err == storage.IterateDone {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate objects: %v", err)
		}
		fmt.Println(objAttrs.Name)
	}
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func main() {
	// Documentation: Creates a new document in Cloud Firestore.
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, "your-project-id", option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}
	defer client.Close()

	docRef := client.Collection("cities").Doc("LA")
	data := map[string]interface{}{
		"name":    "Los Angeles",
		"state":   "CA",
		"country": "USA",
	}

	_, err = docRef.Set(ctx, data)
	if err != nil {
		log.Fatalf("Failed to create document: %v", err)
	}

	fmt.Println("Document created successfully.")
}





package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/cloudfunctions/v1"
	"google.golang.org/api/option"
)

func main() {
	// Documentation: Lists all Cloud Functions in a Google Cloud project.
	ctx := context.Background()

	service, err := cloudfunctions.NewService(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Cloud Functions service: %v", err)
	}

	parent := "projects/your-project-id/locations/-"
	functions, err := service.Projects.Locations.Functions.List(parent).Do()
	if err != nil {
		log.Fatalf("Failed to list functions: %v", err)
	}

	fmt.Println("Cloud Functions:")
	for _, function := range functions.Functions {
		fmt.Printf("%s (%s)\n", function.Name, function.Status)
	}
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func main() {
	// Documentation: Creates a new Google Cloud Storage bucket.
	ctx := context.Background()

	client, err := storage.NewClient(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	bucket := "new-bucket-name"
	err = client.Bucket(bucket).Create(ctx, "your-project-id", nil)
	if err != nil {
		log.Fatalf("Failed to create bucket: %v", err)
	}

	fmt.Printf("Bucket %s created successfully.\n", bucket)
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/bigquery"
)

func main() {
	// Documentation: Lists all BigQuery datasets in a Google Cloud project.
	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, "your-project-id")
	if err != nil {
		log.Fatalf("Failed to create BigQuery client: %v", err)
	}
	defer client.Close()

	datasets := client.Datasets(ctx)
	for {
		dataset, err := datasets.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate datasets: %v", err)
		}
		fmt.Println(dataset.DatasetID)
	}
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/spanner/admin/instance/apiv1"
	"google.golang.org/api/option"
	adminpb "google.golang.org/genproto/googleapis/spanner/admin/instance/v1"
)

func main() {
	// Documentation: Creates a new Cloud Spanner instance.
	ctx := context.Background()

	client, err := instance.NewInstanceAdminClient(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Spanner instance client: %v", err)
	}
	defer client.Close()

	instanceID := "your-instance-id"
	instanceName := "projects/your-project-id/instances/" + instanceID
	config := "regional-us-central1"

	req := &adminpb.CreateInstanceRequest{
		Parent:     "projects/your-project-id",
		InstanceId: instanceID,
		Instance: &adminpb.Instance{
			Name:    instanceName,
			Config:  config,
			NodeCount: 1,
		},
	}

	op, err := client.CreateInstance(ctx, req)
	if err != nil {
		log.Fatalf("Failed to create instance: %v", err)
	}

	fmt.Printf("Instance %s is being created.\n", instanceID)
	fmt.Printf("Operation ID: %s\n", op.Name())
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/bigquery"
)

func main() {
	// Documentation: Inserts a row into a BigQuery table.
	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, "your-project-id")
	if err != nil {
		log.Fatalf("Failed to create BigQuery client: %v", err)
	}
	defer client.Close()

	datasetID := "your-dataset-id"
	tableID := "your-table-id"

	inserter := client.Dataset(datasetID).Table(tableID).Inserter()
	row := map[string]interface{}{
		"column1": "value1",
		"column2": 123,
		"column3": true,
	}

	if err := inserter.Put(ctx, row); err != nil {
		log.Fatalf("Failed to insert row: %v", err)
	}

	fmt.Println("Row inserted successfully.")
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/option"
)

type Task struct {
	Description string
	Completed   bool
}

func main() {
	// Documentation: Creates a new entity in Cloud Datastore.
	ctx := context.Background()

	client, err := datastore.NewClient(ctx, "your-project-id", option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Datastore client: %v", err)
	}
	defer client.Close()

	taskKey := datastore.IncompleteKey("Task", nil)
	task := &Task{
		Description: "Sample task",
		Completed:   false,
	}

	if _, err := client.Put(ctx, taskKey, task); err != nil {
		log.Fatalf("Failed to create task: %v", err)
	}

	fmt.Println("Task created successfully.")
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/iam"
	"google.golang.org/api/option"
)

func main() {
	// Documentation: Lists IAM policies for a Google Cloud resource.
	ctx := context.Background()

	client, err := iam.NewIamClient(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create IAM client: %v", err)
	}

	resource := "//cloudresourcemanager.googleapis.com/projects/your-project-id"
	policy, err := client.GetPolicy(ctx, &iam.GetPolicyRequest{
		Resource: resource,
	})
	if err != nil {
		log.Fatalf("Failed to get IAM policy: %v", err)
	}

	fmt.Printf("IAM Policy for %s:\n", resource)
	for _, binding := range policy.Bindings {
		fmt.Printf("- Role: %s\n", binding.Role)
		fmt.Println("  Members:")
		for _, member := range binding.Members {
			fmt.Printf("  - %s\n", member)
		}
	}
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/kms/apiv1"
	kmspb "google.golang.org/genproto/googleapis/cloud/kms/v1"
	"google.golang.org/api/option"
)

func main() {
	// Documentation: Creates a new Cloud KMS keyring and key.
	ctx := context.Background()

	client, err := kms.NewKeyManagementClient(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create KMS client: %v", err)
	}
	defer client.Close()

	location := "global"
	keyRing := "your-keyring"
	keyID := "your-key"

	keyRingPath := fmt.Sprintf("projects/your-project-id/locations/%s/keyRings/%s", location, keyRing)
	req := &kmspb.CreateCryptoKeyRequest{
		Parent:      keyRingPath,
		CryptoKeyId: keyID,
		CryptoKey: &kmspb.CryptoKey{
			Purpose: kmspb.CryptoKey_ENCRYPT_DECRYPT,
		},
	}

	key, err := client.CreateCryptoKey(ctx, req)
	if err != nil {
		log.Fatalf("Failed to create key: %v", err)
	}

	fmt.Printf("Key %s created in keyring %s.\n", key.Name, keyRingPath)
}





package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/cloudcdn/v1"
	"google.golang.org/api/option"
)

func main() {
	// Documentation: Lists Cloud CDN services in a Google Cloud project.
	ctx := context.Background()

	service, err := cloudcdn.NewService(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Cloud CDN service: %v", err)
	}

	project := "projects/your-project-id"
	response, err := service.Projects.Locations.Services.List(project, "global").Do()
	if err != nil {
		log.Fatalf("Failed to list CDN services: %v", err)
	}

	fmt.Println("Cloud CDN Services:")
	for _, cdn := range response.Services {
		fmt.Printf("%s (%s)\n", cdn.Name, cdn.DisplayName)
	}
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/scheduler/apiv1"
	"github.com/golang/protobuf/ptypes/duration"
	"google.golang.org/api/option"
	schedulerpb "google.golang.org/genproto/googleapis/cloud/scheduler/v1"
)

func main() {
	// Documentation: Creates a new Cloud Scheduler job.
	ctx := context.Background()

	client, err := scheduler.NewCloudSchedulerClient(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Cloud Scheduler client: %v", err)
	}
	defer client.Close()

	project := "your-project-id"
	location := "your-location" // e.g., "us-central1"
	jobName := "your-job-name"

	parent := fmt.Sprintf("projects/%s/locations/%s", project, location)
	job := &schedulerpb.Job{
		Name: fmt.Sprintf("%s/jobs/%s", parent, jobName),
		Target: &schedulerpb.Job_HttpTarget{
			HttpTarget: &schedulerpb.HttpTarget{
				Uri: "https://example.com",
				HttpMethod: schedulerpb.HttpMethod_GET,
			},
		},
		Schedule: "*/5 * * * *", // Run every 5 minutes
		TimeZone: "UTC",
		RetryConfig: &schedulerpb.RetryConfig{
			RetryCount: 3,
		},
		AttemptDeadline: &duration.Duration{
			Seconds: 600,
		},
	}

	createdJob, err := client.CreateJob(ctx, &schedulerpb.CreateJobRequest{
		Parent: parent,
		Job:    job,
	})
	if err != nil {
		log.Fatalf("Failed to create job: %v", err)
	}

	fmt.Printf("Job %s created successfully.\n", createdJob.Name)
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/redis/apiv1"
	"google.golang.org/api/option"
	redispb "google.golang.org/genproto/googleapis/cloud/redis/v1"
)

func main() {
	// Documentation: Lists Cloud Memorystore Redis instances.
	ctx := context.Background()

	client, err := redis.NewCloudRedisClient(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Memorystore Redis client: %v", err)
	}
	defer client.Close()

	parent := "projects/your-project-id/locations/-"
	response, err := client.ListInstances(ctx, &redispb.ListInstancesRequest{
		Parent: parent,
	})
	if err != nil {
		log.Fatalf("Failed to list Redis instances: %v", err)
	}

	fmt.Println("Cloud Memorystore Redis Instances:")
	for _, instance := range response.Instances {
		fmt.Printf("%s (%s)\n", instance.Name, instance.DisplayName)
	}
}





package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/sqladmin/v1beta4"
	"google.golang.org/api/option"
)

func main() {
	// Documentation: Creates a backup for a Cloud SQL instance.
	ctx := context.Background()

	service, err := sqladmin.NewService(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Cloud SQL service: %v", err)
	}

	instance := "your-instance-id"
	backupConfig := &sqladmin.BackupConfiguration{
		StartTime: "04:00", // HH:MM format
		Location:  "us-central1",
	}

	op, err := service.BackupRuns.Insert("your-project-id", instance, backupConfig).Do()
	if err != nil {
		log.Fatalf("Failed to create backup: %v", err)
	}

	fmt.Printf("Backup operation ID: %s\n", op.Name)
}





package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/cloudfunctions/v1"
	"google.golang.org/api/option"
)

func main() {
	// Documentation: Lists triggers for Cloud Functions.
	ctx := context.Background()

	service, err := cloudfunctions.NewService(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Cloud Functions service: %v", err)
	}

	parent := "projects/your-project-id/locations/-"
	triggers, err := service.Projects.Locations.Triggers.List(parent).Do()
	if err != nil {
		log.Fatalf("Failed to list triggers: %v", err)
	}

	fmt.Println("Cloud Functions Triggers:")
	for _, trigger := range triggers.Triggers {
		fmt.Printf("%s (%s)\n", trigger.Name, trigger.EventType)
	}
}





package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func main() {
	// Documentation: Creates a signed URL for a Cloud Storage object.
	ctx := context.Background()

	client, err := storage.NewClient(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	bucket := "your-bucket-name"
	object := "your-object-name"
	expiration := time.Now().Add(1 * time.Hour) // URL expires in 1 hour

	url, err := storage.SignedURL(bucket, object, &storage.SignedURLOptions{
		GoogleAccessID: "your-service-account@your-project-id.iam.gserviceaccount.com",
		PrivateKey:     []byte("-----BEGIN PRIVATE KEY-----\nYOUR_PRIVATE_KEY\n-----END PRIVATE KEY-----\n"),
		Method:         "GET",
		Expires:        expiration,
	})
	if err != nil {
		log.Fatalf("Failed to create signed URL: %v", err)
	}

	fmt.Printf("Signed URL for %s/%s:\n%s\n", bucket, object, url)
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/logging"
	"google.golang.org/api/option"
)

func main() {
	// Documentation: Lists Cloud Logging entries.
	ctx := context.Background()

	client, err := logging.NewClient(ctx, "your-project-id", option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Logging client: %v", err)
	}
	defer client.Close()

	filter := "severity=ERROR"
	iter := client.Entries(ctx, logging.Filter(filter))
	for {
		entry, err := iter.Next()
		if err == logging.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate entries: %v", err)
		}
		fmt.Printf("[%s] %s\n", entry.Timestamp.Format(time.RFC3339), entry.Payload)
	}
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/vision/apiv1"
	"google.golang.org/api/option"
)

func main() {
	// Documentation: Creates a new Cloud Vision API client.
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Vision API client: %v", err)
	}
	defer client.Close()

	fmt.Println("Vision API client created successfully.")
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

func main() {
	// Documentation: Lists all Cloud Pub/Sub topics in a Google Cloud project.
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, "your-project-id", option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Pub/Sub client: %v", err)
	}
	defer client.Close()

	topics, err := client.Topics(ctx)
	if err != nil {
		log.Fatalf("Failed to list topics: %v", err)
	}

	fmt.Println("Cloud Pub/Sub Topics:")
	for _, topic := range topics {
		fmt.Println(topic.ID())
	}
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/language/apiv1"
	"google.golang.org/api/option"
)

func main() {
	// Documentation: Creates a new Cloud Natural Language API client.
	ctx := context.Background()

	client, err := language.NewClient(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Natural Language API client: %v", err)
	}
	defer client.Close()

	fmt.Println("Natural Language API client created successfully.")
}





package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/sqladmin/v1beta4"
	"google.golang.org/api/option"
)

func main() {
	// Documentation: Lists all Cloud SQL instances in a Google Cloud project.
	ctx := context.Background()

	service, err := sqladmin.NewService(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Cloud SQL service: %v", err)
	}

	instances, err := service.Instances.List("your-project-id").Do()
	if err != nil {
		log.Fatalf("Failed to list instances: %v", err)
	}

	fmt.Println("Cloud SQL Instances:")
	for _, instance := range instances.Items {
		fmt.Printf("%s (%s)\n", instance.Name, instance.Region)
	}
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/iot/apiv1"
	"google.golang.org/api/option"
	iotpb "google.golang.org/genproto/googleapis/cloud/iot/v1"
)

func main() {
	// Documentation: Creates a new Cloud IoT Core device.
	ctx := context.Background()

	client, err := iot.NewDeviceManagerClient(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create IoT Core client: %v", err)
	}
	defer client.Close()

	parent := "projects/your-project-id/locations/your-region"
	deviceID := "your-device-id"
	device := &iotpb.Device{
		Id:   deviceID,
		Type: iotpb.Device_GATEWAY,
	}

	createdDevice, err := client.CreateDevice(ctx, &iotpb.CreateDeviceRequest{
		Parent: parent,
		Device: device,
	})
	if err != nil {
		log.Fatalf("Failed to create device: %v", err)
	}

	fmt.Printf("Device %s created successfully.\n", createdDevice.Name)
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/scheduler/apiv1"
	"google.golang.org/api/option"
	schedulerpb "google.golang.org/genproto/googleapis/cloud/scheduler/v1"
)

func main() {
	// Documentation: Lists Cloud Scheduler jobs in a Google Cloud project.
	ctx := context.Background()

	client, err := scheduler.NewCloudSchedulerClient(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Cloud Scheduler client: %v", err)
	}
	defer client.Close()

	project := "your-project-id"
	location := "your-location" // e.g., "us-central1"

	parent := fmt.Sprintf("projects/%s/locations/%s", project, location)
	response, err := client.ListJobs(ctx, &schedulerpb.ListJobsRequest{
		Parent: parent,
	})
	if err != nil {
		log.Fatalf("Failed to list jobs: %v", err)
	}

	fmt.Println("Cloud Scheduler Jobs:")
	for _, job := range response.Jobs {
		fmt.Printf("%s (%s)\n", job.Name, job.Schedule)
	}
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/cloudbuild/apiv1"
	"google.golang.org/api/option"
	cloudbuildpb "google.golang.org/genproto/googleapis/devtools/cloudbuild/v1"
)

func main() {
	// Documentation: Creates a new Cloud Build trigger.
	ctx := context.Background()

	client, err := cloudbuild.NewClient(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Cloud Build client: %v", err)
	}
	defer client.Close()

	project := "your-project-id"
	triggerID := "your-trigger-id"

	trigger := &cloudbuildpb.BuildTrigger{
		TriggerTemplate: &cloudbuildpb.BuildTrigger_GitHub{
			GitHub: &cloudbuildpb.GitHubEventsConfig{
				Owner:      "your-github-owner",
				Repo:       "your-github-repo",
				PullRequest: true,
			},
		},
		Substitutions: map[string]string{
			"_YOUR_VAR": "value",
		},
	}

	createdTrigger, err := client.CreateBuildTrigger(ctx, &cloudbuildpb.CreateBuildTriggerRequest{
		ProjectId: project,
		Trigger:   trigger,
		TriggerId: triggerID,
	})
	if err != nil {
		log.Fatalf("Failed to create build trigger: %v", err)
	}

	fmt.Printf("Build trigger created: %s\n", createdTrigger.Name)
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/option"
)

type Task struct {
	Description string
	Completed   bool
}

func main() {
	// Documentation: Lists entities in Cloud Datastore.
	ctx := context.Background()

	client, err := datastore.NewClient(ctx, "your-project-id", option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Datastore client: %v", err)
	}
	defer client.Close()

	query := datastore.NewQuery("Task")
	tasks := client.Run(ctx, query)

	fmt.Println("Tasks:")
	for {
		var task Task
		_, err := tasks.Next(&task)
		if err == datastore.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate tasks: %v", err)
		}
		fmt.Printf("- %s (Completed: %v)\n", task.Description, task.Completed)
	}
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

func main() {
	// Documentation: Creates a new Cloud Pub/Sub subscription.
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, "your-project-id", option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Pub/Sub client: %v", err)
	}
	defer client.Close()

	topicID := "your-topic-id"
	subID := "your-subscription-id"

	topic := client.Topic(topicID)
	sub, err := client.CreateSubscription(ctx, subID, pubsub.SubscriptionConfig{
		Topic: topic,
	})
	if err != nil {
		log.Fatalf("Failed to create subscription: %v", err)
	}

	fmt.Printf("Subscription %s created for topic %s.\n", sub.ID(), topicID)
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/redis/apiv1"
	"google.golang.org/api/option"
	redispb "google.golang.org/genproto/googleapis/cloud/redis/v1"
)

func main() {
	// Documentation: Lists Cloud Memorystore Redis instances.
	ctx := context.Background()

	client, err := redis.NewCloudRedisClient(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Memorystore Redis client: %v", err)
	}
	defer client.Close()

	parent := "projects/your-project-id/locations/-"
	response, err := client.ListInstances(ctx, &redispb.ListInstancesRequest{
		Parent: parent,
	})
	if err != nil {
		log.Fatalf("Failed to list Redis instances: %v", err)
	}

	fmt.Println("Cloud Memorystore Redis Instances:")
	for _, instance := range response.Instances {
		fmt.Printf("%s (%s)\n", instance.Name, instance.DisplayName)
	}
}





package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/cloudfunctions/v1"
	"google.golang.org/api/option"
)

func main() {
	// Documentation: Creates a new HTTP-triggered Cloud Function.
	ctx := context.Background()

	service, err := cloudfunctions.NewService(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Cloud Functions service: %v", err)
	}

	parent := "projects/your-project-id/locations/-"
	functionID := "your-function-id"

	function := &cloudfunctions.CloudFunction{
		Name:        fmt.Sprintf("projects/your-project-id/locations/-/functions/%s", functionID),
		Description: "HTTP function triggered by HTTP request",
		EntryPoint:  "yourEntryPoint",
		Runtime:     "go113",
		Timeout:     "60s",
		AvailableMemoryMb: 256,
		SourceArchiveUrl: "gs://your-bucket/function-source.zip",
		HttpsTrigger: &cloudfunctions.HttpsTrigger{
			Url: "https://your-cloud-function-url",
		},
	}

	createdFunction, err := service.Projects.Locations.Functions.Create(parent, function).Do()
	if err != nil {
		log.Fatalf("Failed to create function: %v", err)
	}

	fmt.Printf("Cloud Function %s created.\n", createdFunction.Name)
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/cloudtasks/apiv2"
	"google.golang.org/api/option"
	taskspb "google.golang.org/genproto/googleapis/cloud/tasks/v2"
)

func main() {
	// Documentation: Lists Cloud Tasks queues in a Google Cloud project.
	ctx := context.Background()

	client, err := cloudtasks.NewClient(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Cloud Tasks client: %v", err)
	}
	defer client.Close()

	project := "your-project-id"
	parent := fmt.Sprintf("projects/%s/locations/-", project)

	response, err := client.ListQueues(ctx, &taskspb.ListQueuesRequest{
		Parent: parent,
	})
	if err != nil {
		log.Fatalf("Failed to list queues: %v", err)
	}

	fmt.Println("Cloud Tasks Queues:")
	for _, queue := range response.Queues {
		fmt.Printf("%s (%s)\n", queue.Name, queue.State)
	}
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/iam/admin/apiv1"
	"google.golang.org/api/option"
	iampb "google.golang.org/genproto/googleapis/iam/admin/v1"
)

func main() {
	// Documentation: Creates a new IAM policy binding for a Google Cloud resource.
	ctx := context.Background()

	client, err := admin.NewIamClient(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create IAM client: %v", err)
	}
	defer client.Close()

	resource := "//cloudresourcemanager.googleapis.com/projects/your-project-id"
	member := "user:your-email@example.com"
	role := "roles/editor"

	policy, err := client.SetIamPolicy(ctx, &iampb.SetIamPolicyRequest{
		Resource: resource,
		Policy: &iampb.Policy{
			Bindings: []*iampb.Binding{
				{
					Role:    role,
					Members: []string{member},
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("Failed to set IAM policy: %v", err)
	}

	fmt.Println("IAM policy binding created successfully.")
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {
	// Documentation: Lists Cloud Spanner instances in a Google Cloud project.
	ctx := context.Background()

	client, err := spanner.NewInstanceAdminClient(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Spanner client: %v", err)
	}
	defer client.Close()

	projectID := "your-project-id"
	parent := fmt.Sprintf("projects/%s", projectID)

	iter := client.Instances(ctx, &spannerpb.ListInstancesRequest{
		Parent: parent,
	})
	for {
		instance, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate instances: %v", err)
		}
		fmt.Printf("Instance ID: %s, Display Name: %s\n", instance.Instance, instance.DisplayName)
	}
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/bigtable"
	"google.golang.org/api/option"
)

func main() {
	// Documentation: Creates a new table in Cloud Bigtable.
	ctx := context.Background()

	adminClient, err := bigtable.NewAdminClient(ctx, "your-project-id", "your-instance-id", option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Bigtable admin client: %v", err)
	}
	defer adminClient.Close()

	tableName := "your-table-name"
	err = adminClient.CreateTable(ctx, tableName)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	fmt.Printf("Table %s created successfully in Bigtable.\n", tableName)
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/dataflow"
	"google.golang.org/api/option"
	dataflowpb "google.golang.org/genproto/googleapis/dataflow/v1"
)

func main() {
	// Documentation: Lists Cloud Dataflow jobs in a Google Cloud project.
	ctx := context.Background()

	client, err := dataflow.NewClient(ctx, "your-project-id", option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Dataflow client: %v", err)
	}
	defer client.Close()

	response, err := client.ListJobs(ctx, &dataflowpb.ListJobsRequest{
		ProjectId: "your-project-id",
	})
	if err != nil {
		log.Fatalf("Failed to list jobs: %v", err)
	}

	fmt.Println("Cloud Dataflow Jobs:")
	for _, job := range response.Jobs {
		fmt.Printf("%s (%s)\n", job.Name, job.CurrentState)
	}
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/monitoring/dashboard/apiv1"
	"google.golang.org/api/option"
	dashboardpb "google.golang.org/genproto/googleapis/monitoring/dashboard/v1"
)

func main() {
	// Documentation: Creates a new Cloud Monitoring dashboard.
	ctx := context.Background()

	client, err := dashboard.NewDashboardsClient(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Monitoring client: %v", err)
	}
	defer client.Close()

	parent := "projects/your-project-id"
	dashboard := &dashboardpb.Dashboard{
		Name: "projects/your-project-id/dashboards/your-dashboard-id",
		GridLayout: &dashboardpb.GridLayout{
			Columns: 2,
			Rows:    2,
		},
		Title: "Example Dashboard",
	}

	createdDashboard, err := client.CreateDashboard(ctx, &dashboardpb.CreateDashboardRequest{
		Parent:    parent,
		Dashboard: dashboard,
	})
	if err != nil {
		log.Fatalf("Failed to create dashboard: %v", err)
	}

	fmt.Printf("Dashboard %s created successfully.\n", createdDashboard.Name)
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/composer/apiv1"
	"google.golang.org/api/option"
	composerpb "google.golang.org/genproto/googleapis/cloud/composer/v1"
)

func main() {
	// Documentation: Lists Cloud Composer environments in a Google Cloud project.
	ctx := context.Background()

	client, err := composer.NewEnvironmentsClient(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Composer client: %v", err)
	}
	defer client.Close()

	parent := "projects/your-project-id/locations/-"
	response, err := client.ListEnvironments(ctx, &composerpb.ListEnvironmentsRequest{
		Parent: parent,
	})
	if err != nil {
		log.Fatalf("Failed to list environments: %v", err)
	}

	fmt.Println("Cloud Composer Environments:")
	for _, environment := range response.Environments {
		fmt.Printf("%s (%s)\n", environment.Name, environment.State)
	}
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/speech/apiv1"
	"google.golang.org/api/option"
)

func main() {
	// Documentation: Creates a new Cloud Speech-to-Text API client.
	ctx := context.Background()

	client, err := speech.NewClient(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Speech-to-Text API client: %v", err)
	}
	defer client.Close()

	fmt.Println("Speech-to-Text API client created successfully.")
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {
	// Documentation: Lists Cloud Firestore databases in a Google Cloud project.
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, "your-project-id", option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}
	defer client.Close()

	iter := client.Databases(ctx)
	for {
		db, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate databases: %v", err)
		}
		fmt.Printf("Database ID: %s, Project ID: %s\n", db.ID, db.ProjectID)
	}
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/translate"
	"google.golang.org/api/option"
)

func main() {
	// Documentation: Creates a new Cloud Translation API client.
	ctx := context.Background()

	client, err := translate.NewClient(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Translation API client: %v", err)
	}
	defer client.Close()

	fmt.Println("Translation API client created successfully.")
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {
	// Documentation: Lists all Cloud Storage buckets in a Google Cloud project.
	ctx := context.Background()

	client, err := storage.NewClient(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Storage client: %v", err)
	}
	defer client.Close()

	iter := client.Buckets(ctx, "your-project-id")
	for {
		bucketAttrs, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate buckets: %v", err)
		}
		fmt.Println(bucketAttrs.Name)
	}
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/dataproc/apiv1"
	"google.golang.org/api/option"
	dataprocpb "google.golang.org/genproto/googleapis/cloud/dataproc/v1"
)

func main() {
	// Documentation: Creates a new Cloud Dataproc cluster.
	ctx := context.Background()

	client, err := dataproc.NewClusterControllerClient(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Dataproc client: %v", err)
	}
	defer client.Close()

	projectID := "your-project-id"
	region := "your-region"
	clusterName := "your-cluster-name"

	cluster := &dataprocpb.Cluster{
		ProjectId: projectID,
		ClusterName: clusterName,
		Config: &dataprocpb.ClusterConfig{
			MasterConfig: &dataprocpb.InstanceGroupConfig{
				NumInstances: 1,
				MachineTypeUri: "n1-standard-4",
			},
			WorkerConfig: &dataprocpb.InstanceGroupConfig{
				NumInstances: 2,
				MachineTypeUri: "n1-standard-4",
			},
		},
	}

	op, err := client.CreateCluster(ctx, &dataprocpb.CreateClusterRequest{
		ProjectId: projectID,
		Region: region,
		Cluster: cluster,
	})
	if err != nil {
		log.Fatalf("Failed to create cluster: %v", err)
	}

	fmt.Printf("Cluster operation ID: %s\n", op.Name)
}





package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/aiplatform/operations"
	"google.golang.org/api/option"
	aiplatformpb "google.golang.org/genproto/googleapis/cloud/aiplatform/v1"
)

func main() {
	// Documentation: Lists AI Platform models in a Google Cloud project.
	ctx := context.Background()

	client, err := operations.NewClient(ctx, option.WithCredentialsFile("path-to-service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create AI Platform client: %v", err)
	}
	defer client.Close()

	parent := "projects/your-project-id/locations/-"
	response, err := client.ListModels(ctx, &aiplatformpb.ListModelsRequest{
		Parent: parent,
	})
	if err != nil {
		log.Fatalf("Failed to list models: %v", err)
	}

	fmt.Println("AI Platform Models:")
	for _, model := range response.Models {
		fmt.Printf("%s (%s)\n", model.Name, model.DeploymentUri)
	}
}
