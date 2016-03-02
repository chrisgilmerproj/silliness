# Delayed Messages


## Testing

In one window run:

```sh
(env) $ python server.py
```

First create a job:

```sh
$ curl -H "Content-Type: application/json" -XPOST -d '{"url":"example.com","delay":"5"}' http://localhost:8888/send
```

Using the job id that is returned get information about the job:

```sh
$ export JOB_ID=0047dab4-ae60-4c05-a336-7d9d249a20fa
$ curl -XGET http://localhost:8888/$JOB_ID
```

Modify the job:

```sh
$ export JOB_ID=0047dab4-ae60-4c05-a336-7d9d249a20fa
$ curl -H "Content-Type: application/json" -XPUT -d '{"url":"example.com","cb":"/process","delay":"5"}' http://localhost:8888/$JOB_ID
```

delete the job:

```sh
$ export JOB_ID=0047dab4-ae60-4c05-a336-7d9d249a20fa
$ curl -XDELETE http://localhost:8888/$JOB_ID
```

## Instructions

Notes for delayed messages:


HTTP GET/POST/PUT: /send?url={url}&cb={cb}&delay={delay}

Sends a GET/POST/PUT message to a given url after a delay

url is the URL to send the message to

cb [optional] A callback endpoint to send the response to

delay Delay in seconds to delay the message

Return value: Content-Type: "application/json" whose body specifies an id of the message. e.g.

{ "id" : "asdgf53q34" }

HTTP DELETE: /{id}

Cancels the message with id so that it is not sent.


## Personal Notes


POST /send?

url = kevin.com
cb = (none) localhost/service_endpoint (HTTP REST endpoint)
delay = (0s) 5s
action = (GET)/POST only

Return value: Content-Type: "application/json" , { "id" : "asdgf53q34" }

GET /{job_id}

PUT /{job_id}

DELETE /{job_id}


Reactor - send with +20% accuracty

min delay time is 0 seconds


