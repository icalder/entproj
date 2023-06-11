package examples

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/icalder/entproj/util"
	_ "github.com/lib/pq"
)

func ExampleRegistry() {
	ctx := context.Background()
	client, tx := util.OpenSQLiteAndStartTransaction(ctx)
	defer client.Close()
	defer tx.Rollback()

	// Run the automatic migration tool to create all schema resources.
	if err := tx.Client().Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	registry1, err := tx.Registry.Create().
		SetID(uuid.MustParse("1e325c56-c609-4c17-b68b-474eb9097684")).
		SetName("my-registry").
		Save(ctx)

	if err != nil {
		log.Fatalf("failed creating a Registry: %v", err)
	}

	registry2, err := tx.Registry.Create().
		SetID(uuid.MustParse("10d5308c-6a54-463c-adb6-7c9cef932cea")).
		SetName("test-registry").
		Save(ctx)

	if err != nil {
		log.Fatalf("failed creating a Registry: %v", err)
	}

	repo1, err := tx.Repository.Create().
		SetName("ubuntu").
		SetRegistry(registry1).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating a Repository: %v", err)
	}

	artifact1, err := tx.Artifact.Create().
		SetRepository(repo1).
		SetDigest("sha256:5b0bcabd1ed22e9fb1310cf6c2dec7cdef19f0ad69efa1f392e94a4333501270").
		SetMediaType("application/vnd.oci.image.manifest.v1+json").
		SetTags([]string{"22.04", "latest"}).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating an Artifact: %v", err)
	}

	fmt.Println(registry1)
	fmt.Println(registry2)
	fmt.Println(repo1.Name)
	fmt.Println(artifact1.Tags)
	// Output:
	// Registry(id=1e325c56-c609-4c17-b68b-474eb9097684, name=my-registry)
	// Registry(id=10d5308c-6a54-463c-adb6-7c9cef932cea, name=test-registry)
	// ubuntu
	// [22.04 latest]
}
