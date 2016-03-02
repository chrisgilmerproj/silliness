import collections
import datetime
import json
import uuid

import tornado.gen
import tornado.httpclient
import tornado.ioloop
import tornado.web

PROCESS_DELAY = 10000

bucket_storage = collections.defaultdict(list)
job_storage = {}


def get_timestamp():
    return int(datetime.datetime.utcnow().strftime('%s'))


def get_bucket(ts, delay=0):
    bucket = (ts + delay) / 10
    return bucket


def save_data(data):
    bucket = get_bucket(data['ts'], data['delay'])
    bucket_storage[bucket].append(data['id'])
    job_storage[data['id']] = data


def process_data():
    http_client = tornado.httpclient.AsyncHTTPClient()
    ts = get_timestamp()
    bucket = get_bucket(ts)
    job_id_list = bucket_storage[bucket]

    # TODO: Parallelize this step
    for job_id in job_id_list:
        job = job_storage[job_id]
        if 'response' not in job:
            response = yield http_client.fetch(job['url'],
                                               method=job['method'])
            job_storage[job_id]['response'] = response
            if job['cb']:
                response = yield http_client.fetch(job['cb'])


class SendHandler(tornado.web.RequestHandler):

    @tornado.gen.coroutine
    def post(self):
        data = json.loads(self.request.body)
        url = data.get('url', None)
        if not url:
            raise tornado.web.HTTPError(400)
        cb = data.get('cb', None)
        try:
            delay = int(data.get('delay', '0'))
        except ValueError:
            raise tornado.web.HTTPError(400)
        method = data.get('method', 'GET')
        if method not in ['GET', 'POST']:
            raise tornado.web.HTTPError(400)
        job_id = str(uuid.uuid4())

        data = {
            'id': job_id,
            'url': url,
            'cb': cb,
            'delay': delay,
            'method': method,
            'ts': get_timestamp(),
        }
        save_data(data)
        self.set_header("Content-type", "application/json")
        self.write(json.dumps({'id': job_id}))


class IDHandler(tornado.web.RequestHandler):

    @tornado.gen.coroutine
    def get(self, job_id):
        data = job_storage[job_id]
        self.set_header("Content-type", "application/json")
        self.write(json.dumps(data))

    @tornado.gen.coroutine
    def put(self):
        new_data = json.loads(self.request.body)
        job_id = new_data.pop('id')

        old_data = job_storage[job_id]
        old_bucket = get_bucket(old_data['ts'], old_data['delay'])
        bucket_storage[old_bucket].pop(job_id)

        job_storage[job_id].update(new_data)
        old_data.update(new_data)
        save_data(old_data)

        self.set_header("Content-type", "application/json")
        self.write(json.dumps(job_storage[job_id]))

    @tornado.gen.coroutine
    def delete(self):
        new_data = json.loads(self.request.body)
        job_id = new_data.pop('id')

        old_data = job_storage[job_id]
        old_bucket = get_bucket(old_data['ts'], old_data['delay'])
        bucket_storage[old_bucket].pop(job_id)
        job_storage.pop(job_id)

        self.set_header("Content-type", "application/json")
        self.write(json.dumps({}))


def main():
    app = tornado.web.Application(
        [
            (r"/send", SendHandler),
            (r"/([0-9a-zA-z-]+)", IDHandler),
            ],
        )
    app.listen(8888)
    tornado.ioloop.PeriodicCallback(process_data, PROCESS_DELAY).start()
    tornado.ioloop.IOLoop.current().start()


if __name__ == "__main__":
    main()
