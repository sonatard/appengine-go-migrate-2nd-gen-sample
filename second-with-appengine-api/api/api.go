package api

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"google.golang.org/api/iterator"

	"cloud.google.com/go/spanner"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
	"google.golang.org/appengine/search"
	"google.golang.org/appengine/taskqueue"
	"google.golang.org/appengine/user"
)

var spannerCli *spanner.Client

func init() {
	// Spanner
	/*
		ctx := context.Background()
		spannerCli, err := spanner.NewClientWithConfig(
			ctx,
			fmt.Sprintf("projects/%v/instances/%v/databases/%v", projectID, projectID, projectID),
			spanner.ClientConfig{
				SessionPoolConfig: spanner.SessionPoolConfig{
					MinOpened: 1,
				},
			},
		)
		if err != nil {
			panic(err)
		}

		q := `CREATE TABLE "Table" (
						"ID" STRING(MAX) NOT NULL,
						"Value" STRING(MAX) NOT NULL,
					) PRIMARY KEY ("ID");`
	*/
}

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.WithContext(r.Context(), r)

	u := user.Current(ctx)
	if u == nil {
		fmt.Fprintf(w, "Hello guest!!")
		return
	}

	fmt.Fprintf(w, "Hello %v!!", u)

}

func AuthHandle(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.WithContext(r.Context(), r)

	u := user.Current(ctx)
	if u == nil {
		url, _ := user.LoginURL(ctx, "/")
		fmt.Fprintf(w, `<a href="%s">Sign in or register</a>`, url)
		return
	}
	url, _ := user.LogoutURL(ctx, "/")
	fmt.Fprintf(w, `Welcome, %s! (<a href="%s">sign out</a>)`, u, url)
}

func LogHandle(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.WithContext(r.Context(), r)

	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("%v", err)
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Printf("%v", err)
		}
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			fmt.Printf("%v\n", ip)
		}
	}

	// logging in Stackdriver logging
	fmt.Printf("Call handle")

	// logging in Stackdriver logging
	log.Infof(ctx, "Call handle")

	fmt.Fprintln(w, "logHandle!!")
}

func AppengineLogHandle(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.WithContext(r.Context(), r)
	firstString := func(args ...interface{}) string {
		return args[0].(string)
	}

	log.Infof(ctx, "Call appengineLogHandle")

	// logging in Stackdriver logging
	log.Infof(ctx, "DefaultVersionHostname: %v", appengine.DefaultVersionHostname(ctx))
	module := appengine.ModuleName(ctx)
	log.Infof(ctx, "ModuleName: %v", module)
	version := appengine.VersionID(ctx)
	log.Infof(ctx, "VersionID: %v", version)
	instance := appengine.InstanceID()
	log.Infof(ctx, "InstanceID: %v", instance)
	log.Infof(ctx, "ModuleHostname: %v", firstString(appengine.ModuleHostname(ctx, module, version, instance)))
	log.Infof(ctx, "Datacenter: %v", appengine.Datacenter(ctx))
	log.Infof(ctx, "ServerSoftware: %v", appengine.ServerSoftware())
	log.Infof(ctx, "RequestID: %v", appengine.RequestID(ctx))
	log.Infof(ctx, "IsDevAppServer: %v", appengine.IsDevAppServer())
	log.Infof(ctx, "IsAppEngine: %v", appengine.IsAppEngine())
	log.Infof(ctx, "IsStandard: %v", appengine.IsStandard())

	fmt.Fprintln(w, "appengineLogHandle!!")
}

func PanicHandle(w http.ResponseWriter, r *http.Request) {
	// logging in Stackdriver logging
	panic("panicHandle : panic!!")

	fmt.Fprintln(w, "panicHandle!!")
}

func TaskQueueHandle(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.WithContext(r.Context(), r)

	t := taskqueue.NewPOSTTask("/log", nil)
	if _, err := taskqueue.Add(ctx, t, "logQueue"); err != nil {
		log.Errorf(ctx, "taskQueueHandle : %v", err)
	}

	fmt.Fprintln(w, "taskQueueHandle!!")
}

const key = "key"

func MemcacheGetHandle(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.WithContext(r.Context(), r)

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
	ctx := appengine.WithContext(r.Context(), r)
	if err := memcache.Delete(ctx, key); err != nil {
		log.Errorf(ctx, "memcacheDeleteHandle: %v", err)
	}

	fmt.Fprintf(w, "memcacheDeleteHandle Delete %v!!\n", key)
}

type Entity struct {
	Value string
}

const (
	kind = "Entity"
	id   = "stringID"
)

func DatastoreGetHandle(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	k := datastore.NewKey(ctx, kind, id, 0, nil)
	e := new(Entity)
	err := datastore.Get(ctx, k, e)
	if err != nil && err != datastore.ErrNoSuchEntity {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err == datastore.ErrNoSuchEntity {
		e.Value = "value"
		if _, err := datastore.Put(ctx, k, e); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "datastoreGetHandle Put %v!!\n", e)
	} else {
		fmt.Fprintf(w, "datastoreGetHandle Get %v!!\n", e)
	}

}

func DatastoreDeleteHandle(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	k := datastore.NewKey(ctx, kind, id, 0, nil)
	if err := datastore.Delete(ctx, k); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "datastoreDeleteHandle Delete %v!!\n", k)

}

var (
	table = "table"
)

func SpannerGetHandle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if _, err := spannerCli.ReadWriteTransaction(ctx, func(ctx context.Context, tx *spanner.ReadWriteTransaction) error {
		params := Entity{
			Value: "value",
		}
		stmt := spanner.Statement{
			SQL: fmt.Sprintf("SELECT * FROM %s WHERE Value=@Params.Value", table),
			Params: map[string]interface{}{
				"Params": params,
			},
		}

		iter := tx.Query(ctx, stmt)
		defer iter.Stop()

		row, err := iter.Next()
		if err == iterator.Done {
			// Not Found
			e := &Entity{
				Value: "value",
			}
			m, err := spanner.InsertStruct(table, e)
			if err != nil {
				return err
			}

			if err := tx.BufferWrite([]*spanner.Mutation{m}); err != nil {
				return err
			}

			fmt.Fprintf(w, "spannerGetHandle Put %v!!\n", e)

			return nil
		}

		if err != nil {
			return err
		}

		var e Entity
		if err := row.ToStruct(&e); err != nil {
			return err
		}

		fmt.Fprintf(w, "spannerGetHandle Get %v!!\n", e)
		return nil
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

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
