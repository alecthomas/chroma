local application = 'my-app';
local module = 'uwsgi_module';
local dir = '/var/www';
local permission = 644;

{
  'uwsgi.ini': std.manifestIni({
    sections: {
      uwsgi: {
        module: module,
        pythonpath: dir,
        socket: dir + '/uwsgi.sock',
        'chmod-socket': permission,
        callable: application,
        logto: '/var/log/uwsgi/uwsgi.log',
      },
    },
  }),

  'init.sh': |||
    #!/usr/bin/env bash
    mkdir -p %(dir)s
    touch %(dir)s/initialized
    chmod %(perm)d %(dir)s/initialized
  ||| % {dir: dir, perm: permission},

  'cassandra.conf': std.manifestYamlDoc({
    cluster_name: application,
    seed_provider: [
      {
        class_name: 'SimpleSeedProvider',
        parameters: [{ seeds: '127.0.0.1' }],
      },
    ],
  }),
}
