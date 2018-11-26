package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"cloud.google.com/go/datastore"
	"google.golang.org/genproto/googleapis/cloud/tasks/v2beta3"
)

func init() {
	// wget
	out0, err := exec.Command("wget").CombinedOutput()
	log.Printf("wget : %v\n", string(out0))
	log.Printf("wget : error %v\n", err)

	out, err := exec.Command("wget", "https://github.com/sonatard/ghs/releases/download/0.0.10/ghs-0.0.10-linux_amd64.tar.gz", "-O", "/tmp/ghs.tar.gz").CombinedOutput()
	log.Printf("wget : %v\n", string(out))
	log.Printf("wget : error %v\n", err)

	// tar
	out2, err := exec.Command("tar", "xvf", "/tmp/ghs.tar.gz", "-C", "/tmp/").CombinedOutput()
	log.Printf("tar: %v\n", string(out2))
	log.Printf("tar: error %v\n", err)
}

//
func IndexHandle(w http.ResponseWriter, r *http.Request) {
	/*
		ctx:=r.Context()

			u := user.Current(ctx)
			if u == nil {
				fmt.Fprintf(w, "Hello guest!!")
				return
			}
	*/

	fmt.Fprintf(w, "Hello guest!!")
}

// Firebase Authentication or Google Sign-In
// See https://cloud.google.com/appengine/docs/standard/go111/authenticating-users
/*
	func AuthHandle(w http.ResponseWriter, r *http.Request) {
		ctx:=r.Context()

		u := user.Current(ctx)
		if u == nil {
			url, _ := user.LoginURL(ctx, "/")
			fmt.Fprintf(w, `<a href="%s">Sign in or register</a>`, url)
			return
		}
		url, _ := user.LogoutURL(ctx, "/")
		fmt.Fprintf(w, `Welcome, %s! (<a href="%s">sign out</a>)`, u, url)
	}
*/

func LogHandle(w http.ResponseWriter, r *http.Request) {
	// logging in Stackdriver logging
	fmt.Printf("Call handle")

	// logging in Stackdriver logging
	log.Println("Call handle")

	fmt.Fprintln(w, "logHandle!!")
}

func CmdHandle(w http.ResponseWriter, r *http.Request) {
	// exec binary
	out3, err := exec.Command("/tmp/ghs-0.0.10-linux_amd64/ghs", "golang/go", "-m1").CombinedOutput()
	log.Printf("ghs: %v\n", string(out3))
	log.Printf("ghs: %v\n", err)
	fmt.Fprintf(w, "%v", string(out3))
}

func AppengineLogHandle(w http.ResponseWriter, r *http.Request) {
	log.Println("Call appengineLogHandle")

	// logging in Stackdriver logging
	log.Printf("DefaultVersionHostname: %v\n", r.Header.Get("X-AppEngine-Default-Version-Hostname"))
	module := os.Getenv("GAE_SERVICE")
	log.Printf("ModuleName: %v\n", module)
	version := os.Getenv("GAE_VERSION") + "." + os.Getenv("GAE_DEPLOYMENT_ID")
	log.Printf("VersionID: %v\n", version)
	instance := os.Getenv("GAE_INSTANCE")
	log.Printf("InstanceID: %v\n", instance)
	// log.Printf( "ModuleHostname: %v\n", firstString(appengine.ModuleHostname(ctx, module, version, instance)))
	// log.Printf("Datacenter: %v\n", r.Header.Get("X-AppEngine-Datacenter"))
	log.Printf("ServerSoftware: %v\n", os.Getenv("GAE_ENV"))
	log.Printf("RequestID: %v\n", r.Header.Get("X-AppEngine-Request-Log-Id"))
	log.Printf("IsDevAppServer: %v\n", os.Getenv("RUN_WITH_DEVAPPSERVER") != "")
	// log.Printf( "IsAppEngine: %v\n", appengine.IsAppEngine())
	log.Printf("IsStandard: %v\n", os.Getenv("GAE_ENV") == "standard")

	fmt.Fprintln(w, "appengineLogHandle!!")
}

func PanicHandle(w http.ResponseWriter, r *http.Request) {
	// logging in Stackdriver logging
	panic("panicHandle : panic!!")

	fmt.Fprintln(w, "panicHandle!!")
}

func TaskQueueHandle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cli, err := cloudtasks.NewClient(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	req := &tasks.CreateTaskRequest{
		// TODO: Fill request struct fields.
		Parent:       "",
		Task:         nil,
		ResponseView: 0,
	}
	_, err = cli.CreateTask(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintln(w, "taskQueueHandle!!")
}

// Use memorystore in the future. But now not supported from App Engine.
// memorystore not connect from public network. must use VPC SERVICE CONTROLS. But App Engine not supported now.
// To use a memcache service on App Engine, use Redis Labs Memcached Cloud instead of App Engine Memcache.
// https://redislabs.com/redis-enterprise/memcached/
/*
const key = "key"

func MemcacheGetHandle(w http.ResponseWriter, r *http.Request) {
ctx:=r.Context()


	item, err := memcache.Get(ctx, key)
	if err != nil && err != memcache.ErrCacheMiss {
		log.Errorf(ctx, "memcacheGetHandle : %v", err)
	}

	if err == memcache.ErrCacheMiss {
		item := &memcache.Item{
			Key:   "key",
			Value: []byte("value"),
		}

		memcache.Add(ctx, item)
		fmt.Fprintf(w, "memcacheGetHandle Add %v!!\n", item)
	} else {
		fmt.Fprintf(w, "memcacheGetHandle Cache Hit %v!!\n", string(item.Value))
	}
}

func MemcacheDeleteHandle(w http.ResponseWriter, r *http.Request) {
	ctx:=r.Context()
	if err := memcache.Delete(ctx, key); err != nil {
		log.Errorf(ctx, "memcacheDeleteHandle: %v", err)
	}

	fmt.Fprintf(w, "memcacheDeleteHandle Delete %v!!\n", key)
}
*/

type Entity struct {
	Value string
}

const (
	kind = "Entity"
	id   = "stringID"
)

var projectID = os.Getenv("GOOGLE_CLOUD_PROJECT")

func DatastoreGetHandle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cli, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	k := datastore.NameKey(kind, id, nil)
	e := new(Entity)
	err = cli.Get(ctx, k, e)
	if err != nil && err != datastore.ErrNoSuchEntity {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err == datastore.ErrNoSuchEntity {
		e.Value = "value"
		if _, err := cli.Put(ctx, k, e); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "datastoreGetHandle Put %v!!\n", e)
	} else {
		fmt.Fprintf(w, "datastoreGetHandle Get %v!!\n", e)
	}

}

func DatastoreDeleteHandle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cli, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	k := datastore.NameKey(kind, id, nil)
	if err := cli.Delete(ctx, k); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "datastoreDeleteHandle Delete %v!!\n", k)

}

// Instead of using the App Engine Search API, host any full-text search database such as ElasticSearch on Compute Engine and access it from your service.
/*
const indexName = "indexName"

func SearchHandle(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	index, err := search.Open(indexName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	e := &Entity{}
	err = index.Get(ctx, id, e)
	if err != nil && err != search.ErrNoSuchDocument {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err == search.ErrNoSuchDocument {
		e.Value = "Value"
		_, err = index.Put(ctx, id, e)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "searchHandle Put %v!!\n", e)
		return
	}

	fmt.Fprintf(w, "searchHandle Get %v!!\n", e)

}

func SearchDeleteHandle(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	index, err := search.Open(indexName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := index.Delete(ctx, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "searchHandle Get %v!!\n", id)

}
*/
